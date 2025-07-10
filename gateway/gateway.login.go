package gateway

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/common/check"
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/db"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/pkg/alg"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/protocol/mx"
	"github.com/gucooing/BaPs/protocol/proto"
	pb "google.golang.org/protobuf/proto"
	"sync"
	"sync/atomic"
)

var loginNum int64              // 登录玩家数量
var loginUidMap map[int64]int64 // 登录map
var upLoginNum int64            // 成功登录玩家数量
var loginSync sync.Mutex

func GetAllowedSequence(uid int64) int64 {
	loginSync.Lock()
	defer loginSync.Unlock()
	if loginUidMap == nil {
		loginUidMap = make(map[int64]int64)
	}
	if sep, ok := loginUidMap[uid]; ok {
		return atomic.LoadInt64(&loginNum) - sep
	}
	sep := atomic.AddInt64(&loginNum, 1)
	loginUidMap[uid] = sep
	return atomic.LoadInt64(&loginNum) - sep
}

func DelAllowedSequence(uid int64) {
	loginSync.Lock()
	defer loginSync.Unlock()
	if _, ok := loginUidMap[uid]; ok {
		delete(loginUidMap, uid)
		atomic.AddInt64(&upLoginNum, 1)
	}
}

func GetTicketSequence() int64 {
	return atomic.LoadInt64(&loginNum) - atomic.LoadInt64(&upLoginNum)
}

func (g *Gateway) getEnterTicket(c *gin.Context) {
	if !alg.CheckGateWay(c) {
		errBestHTTP(c, proto.WebAPIErrorCode_ClientSuspectedAsCheater)
		return
	}
	bin, err := mx.GetFormMx(c)
	if err != nil {
		return
	}
	rsp := &proto.QueuingGetTicketResponse{}
	defer g.send(c, rsp)
	req := new(proto.QueuingGetTicketRequest)
	err = sonic.Unmarshal(bin, req)
	if err != nil {
		logger.Debug("request err:%s c--->s:%s", err.Error(), string(bin))
		return
	}

	rsp.RequiredSecondsPerUser = 1
	if check.SessionNum >= enter.MaxPlayerNum &&
		enter.MaxPlayerNum > 0 {
		rsp.TicketSequence = GetTicketSequence()                // 排队的玩家数量
		rsp.AllowedSequence = GetAllowedSequence(req.YostarUID) // 你的顺序-倒序
		logger.Debug("在线玩家满")
		return
	}

	yoStarUserLogin := db.GetDBGame().GetYoStarUserLoginByYostarUid(req.YostarUID)
	if yoStarUserLogin == nil {
		errBestHTTP(c, proto.WebAPIErrorCode_DBError)
		return
	}
	if yoStarUserLogin.YostarLoginToken != req.YostarToken ||
		yoStarUserLogin.YostarLoginToken == "" {
		errBestHTTP(c, proto.WebAPIErrorCode_InvalidToken)
		return
	}
	if yoStarUserLogin.Ban {
		errBestHTTP(c, proto.WebAPIErrorCode_AccountBanned)
		return
	}
	// 适配便携码登录提升生命周期
	//yoStarUserLogin.YostarLoginToken = ""
	//if err = db.GetDBGame().UpdateYoStarUserLogin(yoStarUserLogin); err != nil {
	//	errBestHTTP(c, proto.WebAPIErrorCode_DBError)
	//	return
	//}
	enterTicket := mx.GetMxToken(alg.RandCodeInt64(), 30)
	if !enter.AddEnterTicket(yoStarUserLogin.AccountServerId, req.YostarUID, enterTicket) {
		errBestHTTP(c, proto.WebAPIErrorCode_DBError)
		return
	}
	rsp.EnterTicket = enterTicket
	rsp.Birth = "19000101" // 百岁老登玩ba不过分吧
	responsePacket := &proto.ResponsePacket{
		BasePacket: &proto.BasePacket{
			Protocol: req.Protocol,
		},
	}
	rsp.SetPacket(responsePacket)
	logger.Debug("EnterTicket交换成功:%s", rsp.EnterTicket)
	DelAllowedSequence(req.YostarUID)
}

func AccountCheckYostar(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AccountCheckYostarRequest)
	rsp := response.(*proto.AccountCheckYostarResponse)
	var err error

	tickInfo := enter.GetEnterTicketInfo(req.EnterTicket)
	if tickInfo == nil {
		rsp.ResultMessag = "EnterTicket验证失败"
		logger.Debug("EnterTicket验证失败")
		s.Error = proto.WebAPIErrorCode_InvalidToken
		return
	}
	enter.DelEnterTicket(req.EnterTicket)
	s = enter.GetSessionByAccountServerId(tickInfo.AccountServerId)
	if s == nil {
		yostarGame := db.GetDBGame().GetYostarGameByAccountServerId(tickInfo.AccountServerId)
		if yostarGame == nil {
			// new Game Player
			yostarGame, err = db.GetDBGame().AddYostarGameByYostarUid(tickInfo.AccountServerId)
			if err != nil {
				s.Error = proto.WebAPIErrorCode_DBError
				logger.Debug("账号创建失败:%s", err.Error())
				return
			}
		}
		s = enter.NewSession(tickInfo.AccountServerId)
		s.YostarUID = tickInfo.YostarUID
		if yostarGame.BinData != nil {
			pb.Unmarshal(yostarGame.BinData, s.PlayerBin)
		} else {
			s.PlayerBin = game.NewYostarGame(tickInfo.AccountServerId)
			logger.Debug("AccountServerId:%v,新玩家登录Game,创建新账号中", tickInfo.AccountServerId)
		}
	}
	// 更新一次账号缓存
	mxToken := mx.GetMxToken(tickInfo.AccountServerId, 64)
	s.MxToken = mxToken
	if !enter.AddSession(s) {
		logger.Info("AccountServerId:%v,重复上线账号,如果老客户端在线则会被离线", tickInfo.AccountServerId)
	} else {
		logger.Info("AccountServerId:%v,上线账号", tickInfo.AccountServerId)
	}
	rsp.ResultState = 1
	responsePacket := &proto.ResponsePacket{
		BasePacket: &proto.BasePacket{
			SessionKey: &proto.SessionKey{
				AccountServerId: tickInfo.AccountServerId,
				MxToken:         s.MxToken,
			},
			Protocol:  req.Protocol,
			AccountId: tickInfo.AccountServerId,
		},
		MissionProgressDBs:         nil,
		EventMissionProgressDBDict: nil,
		StaticOpenConditions:       nil,
		ServerNotification:         game.GetServerNotification(s),
		ServerTimeTicks:            game.GetServerTime(),
	}
	response.SetPacket(responsePacket)
	// 初始化玩家数据

	newPlayerHash(s)
}

// NewPlayerHash 初始化哈希表
func newPlayerHash(s *enter.Session) {
	// 初始化角色哈希表
	for _, info := range game.GetCharacterInfoList(s) {
		s.AddPlayerHash(info.GetServerId(), info)
	}
	// 初始化物品哈希表
	for _, info := range game.GetItemList(s) {
		s.AddPlayerHash(info.GetServerId(), info)
	}
	// 初始化装备哈希表
	for _, info := range game.GetEquipmentInfoList(s) {
		conf := gdconf.GetEquipmentExcel(info.GetUniqueId())
		if conf == nil || conf.MaxLevel >= 10 {
			continue
		}
		s.AddPlayerHash(info.GetUniqueId(), info)
	}
	////初始化家具哈希表
	//for _, info := range game.GetFurnitureInfoList(s) {
	//	if info.IsBase {
	//		s.AddPlayerHash(info.FurnitureId, info)
	//	}
	//}
}
