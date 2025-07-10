package enter

import (
	"errors"
	"github.com/bytedance/sonic"
	"./common/check"
	dbstruct "./db/struct"
	"./protocol/mx"
	"time"

	"./db"
	"./pkg/alg"
	"./pkg/logger"
	"./protocol/proto"
)

var MaxCacheYostarClanTime = 30 // 最大玩家缓存时间 单位:分
const (
	ClanMaxMemberCount = 50 // 社团最大成员数量
)

type YostarClan struct {
	ServerId         int64                  `json:"serverId"`
	ActiveTime       time.Time              `json:"-"` // 上次活跃时间
	LastUpTime       time.Time              `json:"-"` // 上次保存时间
	ClanName         string                 `json:"clanName"`
	AllAccount       map[int64]*ClanAccount `json:"allAccount"` // 账号id
	President        int64                  `json:"president"`  // 主席账号id
	JoinOption       int32                  `json:"joinOption"`
	Notice           string                 `json:"notice"`
	ApplicantAccount map[int64]*ClanAccount `json:"applicantAccount"` // 申请者
}

type ClanAccount struct {
	Uid             int64 `json:"uid"`
	SocialGrade     int32 `json:"socialGrade"`     //  职位
	JoinTime        int64 `json:"joinTime"`        // 加入时间
	LastLoginTime   int64 `json:"lastLoginTime"`   // 上次访问时间
	AttendanceCount int64 `json:"attendanceCount"` // 出席天数
	ApplicantTime   int64 `json:"applicantTime"`   // 申请时间
}

// 每天4点检查一次是否有用户长时间离线然后离线掉好友数据
func (e *EnterSet) checkYostarClan() {
	yostarClanList := make([]*dbstruct.YostarClan, 0)
	for serverId, info := range GetAllYostarClan() {
		if time.Now().After(info.ActiveTime.
			Add(time.Minute * time.Duration(MaxCacheYostarClanTime))) {
			bin := info.GetYostarClan()
			if bin != nil {
				yostarClanList = append(yostarClanList, bin)
			}
			DelYostarClan(serverId)
			logger.Debug("YostarClan:%v,超时离线", serverId)
		}
	}
	if db.GetDBGame().UpAllYostarClan(yostarClanList) != nil {
		logger.Error("社团数据保存失败")
	} else {
		logger.Info("社团数据保存完毕")
	}
}

// GetAllYostarClan 获取全部缓存社团-有锁
func GetAllYostarClan() map[int64]*YostarClan {
	list := make(map[int64]*YostarClan)
	e := getEnterSet()
	check.GateWaySync.Lock()
	defer check.GateWaySync.Unlock()
	for v, k := range e.YostarClan {
		list[v] = k
	}
	return list
}

// GetAllYostarClanList 获取全部缓存社团
func GetAllYostarClanList() []*YostarClan {
	list := make([]*YostarClan, 0)
	e := getEnterSet()
	for _, k := range e.YostarClan {
		list = append(list, k)
	}
	return list
}

// GetYostarClanByServerId 拉取社团消息
func GetYostarClanByServerId(ycId int64) (yc *YostarClan) {
	defer func() {
		if yc != nil {
			yc.ActiveTime = time.Now()
		}
	}()
	var err error
	var ok bool
	if ycId == 0 {
		return
	}
	s := getEnterSet()
	if yc, ok = s.YostarClan[ycId]; ok {
		return
	}
	yc, err = DbGetYostarClan(ycId)
	if err != nil {
		return
	}
	if s.YostarClan == nil {
		s.YostarClan = make(map[int64]*YostarClan)
	}
	if s.YostarClanHash == nil {
		s.YostarClanHash = make(map[string]int64)
	}
	s.YostarClan[ycId] = yc
	s.YostarClanHash[yc.ClanName] = ycId
	return
}

func GetYostarClanByClanName(clanName string) *YostarClan {
	s := getEnterSet()
	if serverId, ok := s.YostarClanHash[clanName]; ok {
		return GetYostarClanByServerId(serverId)
	}

	yc, err := DbGetYostarClanByClanName(clanName)
	if err != nil {
		return nil
	}
	if s.YostarClan == nil {
		s.YostarClan = make(map[int64]*YostarClan)
	}
	if s.YostarClanHash == nil {
		s.YostarClanHash = make(map[string]int64)
	}
	s.YostarClan[yc.ServerId] = yc
	s.YostarClanHash[yc.ClanName] = yc.ServerId
	return yc
}

// DbGetYostarClanByClanName 从db拉取数据
func DbGetYostarClanByClanName(clanName string) (*YostarClan, error) {
	yc := new(YostarClan)
	bin := db.GetDBGame().GetYostarClanByClanName(clanName)
	if bin == nil {
		return nil, errors.New("sql err")
	}
	sonic.Unmarshal([]byte(bin.ClanInfo), yc)
	yc.ServerId = bin.ServerId
	yc.ClanName = bin.ClanName
	yc.ActiveTime = time.Now()
	yc.LastUpTime = time.Now()
	return yc, nil
}

// DbGetYostarClan 从db拉取数据
func DbGetYostarClan(ycId int64) (*YostarClan, error) {
	yc := new(YostarClan)
	bin := db.GetDBGame().GetYostarClanByServerId(ycId)
	if bin == nil {
		return nil, errors.New("sql err")
	}
	sonic.Unmarshal([]byte(bin.ClanInfo), yc)
	yc.ServerId = ycId
	yc.ClanName = bin.ClanName
	yc.ActiveTime = time.Now()
	yc.LastUpTime = time.Now()
	return yc, nil
}

func DelYostarClan(serverId int64) bool {
	check.GateWaySync.Lock()
	defer check.GateWaySync.Unlock()
	s := getEnterSet()
	if info, ok := s.YostarClan[serverId]; ok {
		info.UpDate()
		delete(s.YostarClan, serverId)
		delete(s.YostarClanHash, info.ClanName)
		return true
	}
	return false
}

// GetYostarClan 预处理db数据
func (x *YostarClan) GetYostarClan() *dbstruct.YostarClan {
	if x == nil {
		return nil
	}
	bin := &dbstruct.YostarClan{
		ServerId: x.ServerId,
		ClanName: x.ClanName,
	}
	ycInfo, err := sonic.Marshal(x)
	if err != nil {
		return nil
	}
	bin.ClanInfo = string(ycInfo)
	return bin
}

// UpDate 将玩家数据保存到数据库
func (x *YostarClan) UpDate() error {
	if x == nil {
		return errors.New("YostarClan is nil")
	}
	bin := &dbstruct.YostarClan{
		ServerId: x.ServerId,
		ClanName: x.ClanName,
	}
	ycInfo, err := sonic.Marshal(x)
	if err != nil {
		return err
	}
	bin.ClanInfo = string(ycInfo)
	err = db.GetDBGame().UpdateYostarClan(bin)
	return err
}

func (x *YostarClan) GetMemberCount() int64 {
	if x == nil {
		return 0
	}
	return int64(len(x.AllAccount))
}

func (x *YostarClan) GetAllAccount() map[int64]*ClanAccount {
	if x == nil {
		return nil
	}
	account := make(map[int64]*ClanAccount)
	for k, v := range x.AllAccount {
		account[k] = v
	}
	return account
}

func (x *YostarClan) GetClanAccount(uid int64) *ClanAccount {
	if x == nil {
		return nil
	}
	info, ok := x.AllAccount[uid]
	if !ok {
		return nil
	}
	return info
}

func (x *YostarClan) AddAccount(uid int64, socialGrade int32) bool {
	if x == nil {
		return false
	}
	if x.AllAccount == nil {
		x.AllAccount = make(map[int64]*ClanAccount)
	}
	if x.AllAccount[uid] == nil {
		x.AllAccount[uid] = &ClanAccount{
			Uid:           uid,
			SocialGrade:   socialGrade,
			JoinTime:      time.Now().Unix(),
			LastLoginTime: time.Now().Unix(),
		}
		return true
	}
	return false
}

func (x *YostarClan) RemoveAccount(uid int64) bool {
	if x == nil {
		return true
	}
	if x.AllAccount == nil {
		x.AllAccount = make(map[int64]*ClanAccount)
	}
	info := x.AllAccount[uid]
	if info == nil {
		return true
	}
	if info.SocialGrade == int32(proto.ClanSocialGrade_President) {
		// 禁止首领退出
		return false
	}
	delete(x.AllAccount, uid)
	return true
}

func (x *YostarClan) SetPresident(uid int64) bool {
	if x == nil {
		return false
	}
	if x.AllAccount == nil {
		x.AllAccount = make(map[int64]*ClanAccount)
	}
	if oldPresident, ok := x.AllAccount[x.President]; ok {
		oldPresident.SocialGrade = int32(proto.ClanSocialGrade_Member)
	}
	if newPresident, ok := x.AllAccount[uid]; ok {
		x.President = uid
		newPresident.SocialGrade = int32(proto.ClanSocialGrade_President)
		return true
	}

	return false
}

func (x *YostarClan) SetNotice(notice string) bool {
	if x == nil {
		return false
	}
	x.Notice = notice
	return true
}

func (x *YostarClan) SetJoinOption(joinOption int32) bool {
	if x == nil {
		return false
	}
	x.JoinOption = joinOption
	return true
}

func (x *YostarClan) GetAllApplicantAccount() map[int64]*ClanAccount {
	if x == nil {
		return nil
	}
	account := make(map[int64]*ClanAccount)
	for k, v := range x.ApplicantAccount {
		account[k] = v
	}
	return account
}

func (x *YostarClan) AddApplicantAccount(uid int64) bool {
	if x == nil {
		return false
	}
	if x.GetMemberCount() >= ClanMaxMemberCount {
		return false
	}
	if x.ApplicantAccount == nil {
		x.ApplicantAccount = make(map[int64]*ClanAccount)
	}
	x.ApplicantAccount[uid] = &ClanAccount{
		Uid:           uid,
		SocialGrade:   int32(proto.ClanSocialGrade_Applicant),
		ApplicantTime: time.Now().Unix(),
	}
	return true
}

func (x *YostarClan) RemoveApplicantAccount(uid int64) {
	if x == nil {
		return
	}
	if x.ApplicantAccount == nil {
		x.ApplicantAccount = make(map[int64]*ClanAccount)
	}
	delete(x.ApplicantAccount, uid)
	return
}

func (x *ClanAccount) GetSocialGrade() int32 {
	if x == nil {
		return 0
	}
	return x.SocialGrade
}

func (x *ClanAccount) GetJoinDate() mx.MxTime {
	if x == nil {
		return mx.MxTime{}
	}
	return mx.Unix(x.JoinTime, 0)
}

func (x *ClanAccount) GetLastLoginTime() mx.MxTime {
	if x == nil {
		return mx.MxTime{}
	}
	return mx.Unix(x.LastLoginTime, 0)
}

func (x *ClanAccount) SetLastLoginTime() {
	if x == nil {
		return
	}
	if time.Unix(x.LastLoginTime, 0).Before(alg.GetLastDayH(4)) {
		x.AttendanceCount++ // 隔了一天就加1
	}
	x.LastLoginTime = time.Now().Unix()
}

func (x *ClanAccount) GetAttendanceCount() int64 {
	if x == nil {
		return 0
	}
	return x.AttendanceCount
}
