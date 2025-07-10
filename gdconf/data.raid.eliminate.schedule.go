package gdconf

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gucooing/BaPs/pkg/logger"
)

type RaidEliminateSchedule struct {
	CurRaidEliminateSchedule *RaidEliminateScheduleInfo
	RaidEliminateScheduleMap map[int64]*RaidEliminateScheduleInfo
}

type RaidEliminateScheduleInfo struct {
	SeasonId     int64
	StartTime    Time `json:"SeasonStartData"`
	EndTime      Time `json:"SeasonEndData"`
	NextSeasonId int64
}

func (g *GameConfig) loadRaidEliminateSchedule() {
	g.GetGPP().RaidEliminateSchedule = &RaidEliminateSchedule{
		RaidEliminateScheduleMap: make(map[int64]*RaidEliminateScheduleInfo),
	}
	name := "RaidEliminateSchedule.json"
	file, err := os.ReadFile(g.dataPath + name)
	if err != nil {
		logger.Error("文件:%s 读取失败,err:%s", name, err)
		return
	}
	if err := json.Unmarshal(file, &g.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap); err != nil {
		logger.Error("文件:%s 解析失败,err:%s", name, err)
		return
	}
	for _, v := range g.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap {
		if v.NextSeasonId != 0 &&
			g.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap[v.NextSeasonId] == nil {
			panic(fmt.Sprintf("缺少下一个大决战排期,NextSeasonId:%v", v.NextSeasonId))
		}
		// 检查排期时间是否满足7天以上
		if v.StartTime.Time().Add(7 * 24 * time.Hour).After(v.EndTime.Time()) {
			panic(fmt.Sprintf("大决战排期时间错误 排期不足7天,SeasonId:%v", v.SeasonId))
		}
	}
	logger.Info("大决战排期读取成功文件:%s 读取成功,解析数量:%v", name, len(g.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap))
}

func GetCurRaidEliminateSchedule() *RaidEliminateScheduleInfo {
	cur := GC.GetGPP().RaidEliminateSchedule.CurRaidEliminateSchedule
	if cur != nil {
		if cur.EndTime.Time().After(time.Now()) {
			return GC.GetGPP().RaidEliminateSchedule.CurRaidEliminateSchedule
		}
		next := GC.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap[cur.NextSeasonId]
		if next != nil && next.StartTime.Time().After(time.Now()) {
			return GC.GetGPP().RaidEliminateSchedule.CurRaidEliminateSchedule
		}
	}
	cur = nil
	for _, v := range GC.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap {
		// 排期中
		if v.EndTime.Time().After(time.Now()) && time.Now().After(v.StartTime.Time()) {
			cur = v
		} else { // 排期间隔中
			next := GC.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap[v.NextSeasonId]
			if next != nil &&
				time.Now().After(v.EndTime.Time()) && next.StartTime.Time().After(time.Now()) {
				cur = v
			}
		}
		if cur != nil {
			GC.GetGPP().RaidEliminateSchedule.CurRaidEliminateSchedule = cur
			return cur
		}
	}
	return cur
}

func GetNextRaidEliminateSchedule() *RaidEliminateScheduleInfo {
	cur := GetCurRaidEliminateSchedule()
	if cur == nil {
		return nil
	}
	return GC.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap[cur.NextSeasonId]
}

func GetRaidEliminateScheduleMap() map[int64]*RaidEliminateScheduleInfo {
	return GC.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap
}

func GetRaidEliminateScheduleInfo(seasonId int64) *RaidEliminateScheduleInfo {
	return GC.GetGPP().RaidEliminateSchedule.RaidEliminateScheduleMap[seasonId]
}
