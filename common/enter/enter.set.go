package enter

import (
	dbstruct "github.com/gucooing/BaPs/db/struct"
	"time"

	"github.com/gucooing/BaPs/pkg/alg"
)

var es *EnterSet

type EnterSet struct {
	SessionMap     map[int64]*Session             // 玩家信息-缓存
	EnterTicketMap map[string]*TicketInfo         // 登录通行证字典-缓存
	MailMap        map[int64]*dbstruct.YostarMail // 全服邮件
	YostarClan     map[int64]*YostarClan          // 全部缓存社团
	YostarClanHash map[string]int64
}

func InitEnterSet() {
	e := getEnterSet()
	e.checkMail()
}

func getEnterSet() *EnterSet {
	if es == nil {
		es = &EnterSet{}
		go es.Check()
	}
	return es
}

func (e *EnterSet) Check() {
	ticker := time.NewTicker(time.Minute * 3)          // 三分钟验证一次
	friendTicker := time.NewTimer(alg.GetEveryDayH(4)) // 每天四点
	upAllTicker := time.NewTicker(time.Minute * 30)    // 半小时一次
	for {
		select {
		case <-ticker.C:
			e.checkSession()
			e.checkMail()
		case <-friendTicker.C:
			e.checkYostarClan()
			friendTicker.Reset(alg.GetEveryDayH(4))
		case <-upAllTicker.C:
			UpAllPlayerBin()
		}
	}
}
