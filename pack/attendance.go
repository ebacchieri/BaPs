package pack

import (
	"github.com/ebacchieri/BaPs/protocol/mx"
	"time"

	"github.com/ebacchieri/BaPs/common/enter"
	sro "github.com/ebacchieri/BaPs/common/server_only"
	"github.com/ebacchieri/BaPs/game"
	"github.com/ebacchieri/BaPs/gdconf"
	"github.com/ebacchieri/BaPs/protocol/proto"
)

func AttendanceReward(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AttendanceRewardRequest)
	rsp := response.(*proto.AttendanceRewardResponse)

	rsp.AttendanceBookRewards = make([]*proto.AttendanceBookReward, 0)
	rsp.AttendanceHistoryDBs = make([]*proto.AttendanceHistoryDB, 0)
	for id, day := range req.DayByBookUniqueId {
		bin := game.GetAttendanceInfo(s, id)
		conf := gdconf.GetAttendanceInfo(id)
		if conf == nil || bin == nil {
			continue
		}
		if bin.AttendedDay == nil {
			bin.AttendedDay = make(map[int64]int64)
		}
		if _, ok := bin.AttendedDay[day]; ok ||
			conf.AttendanceReward[day] == nil {
			continue
		}
		if rewardConf, ok := conf.AttendanceReward[day]; ok {
			parcelInfoList := []*sro.ParcelInfo{
				{
					Type: rewardConf.RewardParcelType,
					Id:   rewardConf.RewardId,
					Num:  rewardConf.RewardAmoun,
				},
			}
			game.AddMailBySystem(s, conf.MailType, parcelInfoList)
		}

		bin.AttendedDay[day] = time.Now().Unix()
		rsp.AttendanceBookRewards = append(rsp.AttendanceBookRewards,
			game.GetAttendanceBookReward(s, id))
		rsp.AttendanceHistoryDBs = append(rsp.AttendanceHistoryDBs,
			game.GetAttendanceHistoryDB(s, id))
		// Do not change the serial number
		bin.LastReward = time.Now().Unix()
	}
}
