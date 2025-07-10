package enter

import (
	"fmt"
	"github.com/gucooing/BaPs/common/check"
	dbstruct "github.com/gucooing/BaPs/db/struct"
	"github.com/gucooing/BaPs/protocol/mx"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/db"
	"github.com/gucooing/BaPs/pkg/alg"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

var MaxCachePlayerTime = 10 // 最大玩家缓存时间 单位:分钟
var MaxPlayerNum int64 = 0  // 最大在线玩家

type Session struct {
	AccountServerId int64
	YostarUID       int64
	MxToken         string
	ActiveTime      time.Time // 上次活跃时间
	LastUpTime      time.Time // 上次保存时间
	AccountState    proto.AccountState
	PlayerBin       *sro.PlayerBin // 玩家数据
	Actions         map[proto.ServerNotificationFlag]bool
	AccountFriend   *AccountFriend
	Mission         *Mission
	Toast           []*Toast
	PlayerHash      map[int64]any
	arenaInfo       *ArenaInfo // 竞技场临时数据
	Error           proto.WebAPIErrorCode
}

type Toast struct {
	Text      string
	BeginDate mx.MxTime
	EndDate   mx.MxTime
}

// 定时检查一次是否有用户长时间离线
func (e *EnterSet) checkSession() {
	yostarGameList := make([]*dbstruct.YostarGame, 0)
	yostarFriendList := make([]*dbstruct.YostarFriend, 0)
	for accountServerId, info := range GetAllSession() {
		if time.Now().After(info.ActiveTime.Add(time.Duration(MaxCachePlayerTime) * time.Minute)) {
			bin := info.GetYostarGame()
			if bin != nil {
				yostarGameList = append(yostarGameList, bin)
				yostarFriendList = append(yostarFriendList, info.GetFriendInfo().GetYostarFriend())
			}
			info.LastUpTime = time.Now()
			DelSession(accountServerId)
			logger.Debug("AccountId:%v,超时离线", accountServerId)
		}
	}
	if len(yostarGameList) == 0 {
		return
	}

	if db.GetDBGame().UpAllYostarGame(yostarGameList) != nil {
		logger.Error("玩家数据保存失败")
	} else {
		logger.Info("玩家数据保存完毕,num:%v", len(yostarGameList))
	}
	if db.GetDBGame().UpAllYostarFriend(yostarFriendList) != nil {
		logger.Error("好友数据保存失败")
	} else {
		logger.Info("好友数据保存完毕,num:%v", len(yostarFriendList))
	}
}

// GetSessionBySessionKey 获取指定在线玩家
func GetSessionBySessionKey(sessionKey *proto.SessionKey) *Session {
	if sessionKey == nil ||
		sessionKey.AccountServerId == 0 ||
		sessionKey.MxToken == "" {
		return nil
	}
	e := getEnterSet()
	if info, ok := e.SessionMap[sessionKey.AccountServerId]; ok {
		if info.MxToken != sessionKey.MxToken {
			return nil
		}
		return info
	}
	return nil
}

// GetSessionByAccountServerId 获取指定在线玩家
func GetSessionByAccountServerId(accountServerId int64) *Session {
	if accountServerId == 0 {
		return nil
	}
	e := getEnterSet()
	if info, ok := e.SessionMap[accountServerId]; ok {
		return info
	}
	return nil
}

// GetSessionByUid 此接口的用处是拉取玩家数据,包括在数据库中的
func GetSessionByUid(uid int64) *Session {
	if uid == 0 {
		return nil
	}
	e := getEnterSet()
	info, ok := e.SessionMap[uid]
	if ok {
		return info
	}
	bin := db.GetDBGame().GetYostarGameByAccountServerId(uid)
	if bin == nil || bin.BinData == nil {
		return nil
	}
	info = NewSession(uid)
	info.AccountServerId = uid
	info.ActiveTime = time.Now()
	err := pb.Unmarshal(bin.BinData, info.PlayerBin)
	if err != nil || info.PlayerBin.GetBaseBin().GetAccountId() != uid {
		return nil
	}
	AddSession(info)
	return info
}

func NewSession(accountServerId int64) *Session {
	return &Session{
		AccountServerId: accountServerId,
		Actions:         make(map[proto.ServerNotificationFlag]bool),
		PlayerBin:       new(sro.PlayerBin),
		AccountFriend:   newAccountFriend(accountServerId),
	}
}

// GetAllSession 获取全部在线玩家-有锁
func GetAllSession() map[int64]*Session {
	allSession := make(map[int64]*Session)
	e := getEnterSet()
	check.GateWaySync.Lock()
	defer check.GateWaySync.Unlock()
	for k, v := range e.SessionMap {
		allSession[k] = v
	}
	return allSession
}

// GetAllSessionList 获取全部在线玩家-列表
func GetAllSessionList() []*Session {
	allSession := make([]*Session, 0)
	e := getEnterSet()
	for _, v := range e.SessionMap {
		allSession = append(allSession, v)
	}
	return allSession
}

// DelSession 删除指定在线玩家-有锁
func DelSession(accountServerId int64) bool {
	e := getEnterSet()
	if e.SessionMap == nil {
		e.SessionMap = make(map[int64]*Session)
	}
	check.GateWaySync.Lock()
	defer check.GateWaySync.Unlock()
	if _, ok := e.SessionMap[accountServerId]; ok {
		delete(e.SessionMap, accountServerId)
		atomic.AddInt64(&check.SessionNum, -1)
		return true
	}
	return false
}

// AddSession 添加Session
func AddSession(x *Session) bool {
	if x == nil ||
		x.AccountServerId == 0 {
		return false
	}
	e := getEnterSet()
	if e.SessionMap == nil {
		e.SessionMap = make(map[int64]*Session)
	}
	if _, ok := e.SessionMap[x.AccountServerId]; ok {
		e.SessionMap[x.AccountServerId] = x
		return false
	}
	e.SessionMap[x.AccountServerId] = x
	x.ActiveTime = time.Now()
	atomic.AddInt64(&check.SessionNum, 1)
	return true
}

func (x *Session) GetFriendInfo() *AccountFriend {
	return x.AccountFriend
}

// Close 保存全部玩家数据
func Close() {
	UpAllPlayerBin()
}

// UpAllPlayerBin 保存全部玩家数据
func UpAllPlayerBin() {
	// 保存玩家主要数据
	yostarGameList := make([]*dbstruct.YostarGame, 0)
	yostarFriendList := make([]*dbstruct.YostarFriend, 0)
	for _, info := range GetAllSession() {
		if info.LastUpTime.After(info.ActiveTime) {
			continue
		}
		bin := info.GetYostarGame()
		if bin != nil {
			yostarGameList = append(yostarGameList, bin)
			yostarFriendList = append(yostarFriendList, info.GetFriendInfo().GetYostarFriend())
		}
		info.LastUpTime = time.Now()
	}
	if len(yostarGameList) != 0 {
		if db.GetDBGame().UpAllYostarGame(yostarGameList) != nil {
			logger.Error("全部玩家数据保存失败")
		} else {
			logger.Info("全部玩家数据保存完毕,num:%v", len(yostarGameList))
		}
		if db.GetDBGame().UpAllYostarFriend(yostarFriendList) != nil {
			logger.Error("全部好友数据保存失败")
		} else {
			logger.Info("全部好友数据保存完毕,num:%v", len(yostarFriendList))
		}
	}

	// 保存社团数据
	yostarClanList := make([]*dbstruct.YostarClan, 0)
	for _, info := range GetAllYostarClan() {
		if info.LastUpTime.After(info.ActiveTime) {
			continue
		}
		bin := info.GetYostarClan()
		if bin != nil {
			yostarClanList = append(yostarClanList, bin)
		}
	}
	if len(yostarClanList) != 0 {
		if db.GetDBGame().UpAllYostarClan(yostarClanList) != nil {
			logger.Error("全部社团数据保存失败")
		} else {
			logger.Info("全部社团数据保存完毕,num:%v", len(yostarClanList))
		}
	}
}

// GetPbBinData 将玩家pb数据转二进制数据
func (x *Session) GetPbBinData() []byte {
	bin, err := pb.Marshal(x.PlayerBin)
	if err != nil {
		return nil
	}
	return bin
}

// GetYostarGame 将玩家数据转成数据库格式
func (x *Session) GetYostarGame() *dbstruct.YostarGame {
	if x == nil {
		return nil
	}
	var fin = true
	defer func() {
		if !fin {
			logger.Debug("玩家:%v,数据保存失败,数据保存将保存到服务端硬盘,下次启动时将尝试写入数据库", x.AccountServerId)
			if err := x.upDataDisk(); err != nil {
				logger.Debug("玩家:%v,数据保存将到服务端硬盘失败,失败原因:%s", x.AccountServerId, err.Error())
			}
		}
	}()
	bin, err := pb.Marshal(x.PlayerBin)
	if err != nil {
		fin = false
		return nil
	}
	data := &dbstruct.YostarGame{
		AccountServerId: x.AccountServerId,
		BinData:         bin,
	}

	return data
}

// 将玩家二进制数据写入磁盘中
func (x *Session) upDataDisk() error {
	bin, err := pb.Marshal(x.PlayerBin)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("./player/%v.bin", x.AccountServerId), bin, 0644)
	return err
}

// TaskUpDiskPlayerData 将磁盘中的二进制玩家数据写入数据库中
func TaskUpDiskPlayerData() bool {
	files, err := filepath.Glob(filepath.Join("./player/", "*.bin"))
	if err != nil {
		return true
	}
	logger.Info("尝试保存本地玩家数据")
	for _, file := range files {
		bin, err := os.ReadFile(file)
		if err != nil {
			logger.Error("读取本地玩家数据失败:%s", err.Error())
			return false
		}
		accountServerId := alg.S2I64(strings.TrimSuffix(filepath.Base(file), ".bin"))
		if accountServerId == 0 {
			logger.Error("本地玩家数据文件名错误,文件:%s", file)
			return false
		}
		data := &dbstruct.YostarGame{
			AccountServerId: accountServerId,
			BinData:         bin,
		}
		if err = db.GetDBGame().UpdateYostarGame(data); err != nil {
			return false
		}
		if os.Remove(file) != nil {
			logger.Warn("删除本地玩家数据文件失败,文件:%s,可能是权限不足导致的,请手动删除,避免下次启动时数据被覆盖", file)
		}
	}
	logger.Info("保存本地玩家数据成功")
	return true
}

func (x *Session) getPlayerHash() map[int64]any {
	if x.PlayerHash == nil {
		x.PlayerHash = make(map[int64]any)
	}
	return x.PlayerHash
}

// AddPlayerHash 写入数据到哈希表中
func (x *Session) AddPlayerHash(k int64, v any) bool {
	list := x.getPlayerHash()
	if _, ok := list[k]; ok {
		return false
	}
	list[k] = v
	return true
}

func (x *Session) DelPlayerHash(k int64) bool {
	list := x.getPlayerHash()
	if _, ok := list[k]; ok {
		delete(list, k)
		return true
	}

	return false
}

func (x *Session) getPlayerHashByKeyId(k int64) any {
	list := x.getPlayerHash()
	return list[k]
}

func (x *Session) GetCharacterByKeyId(k int64) *sro.CharacterInfo {
	v := x.getPlayerHashByKeyId(k)
	switch info := v.(type) {
	case *sro.CharacterInfo:
		return info
	default:
		return nil
	}
}

func (x *Session) GetItemByKeyId(k int64) *sro.ItemInfo {
	v := x.getPlayerHashByKeyId(k)
	switch info := v.(type) {
	case *sro.ItemInfo:
		return info
	default:
		return nil
	}
}

func (x *Session) GetEquipmentByKeyId(k int64) *sro.EquipmentInfo {
	v := x.getPlayerHashByKeyId(k)
	switch info := v.(type) {
	case *sro.EquipmentInfo:
		return info
	default:
		return nil
	}
}
