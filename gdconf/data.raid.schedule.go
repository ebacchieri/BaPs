package gdconf

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gucooing/BaPs/pkg/logger"
)

type RaidSchedule struct {
	CurRaidSchedule *RaidScheduleInfo
	RaidScheduleMap map[int64]*RaidScheduleInfo
}

type RaidScheduleInfo struct {
	SeasonId     int64
	StartTime    Time `json:"SeasonStartData"`
	EndTime      Time `json:"SeasonEndData"`
	NextSeasonId int64
}

func (g *GameConfig) loadRaidSchedule() {
	g.GetGPP().RaidSchedule = &RaidSchedule{
		RaidScheduleMap: make(map[int64]*RaidScheduleInfo),
	}
	name := "RaidSchedule.json"
	file, err := os.ReadFile(g.dataPath + name)
	if err != nil {
		logger.Error("文件:%s 读取失败,err:%s", name, err)
		return
	}
	if err := json.Unmarshal(file, &g.GetGPP().RaidSchedule.RaidScheduleMap); err != nil {
		logger.Error("文件:%s 解析失败,err:%s", name, err)
		return
	}
	for _, v := range g.GetGPP().RaidSchedule.RaidScheduleMap {
		if v.NextSeasonId != 0 &&
			g.GetGPP().RaidSchedule.RaidScheduleMap[v.NextSeasonId] == nil {
			panic(fmt.Sprintf("缺少下一个总力战排期,NextSeasonId:%v", v.NextSeasonId))
		}
		// 检查排期时间是否满足7天以上
		if v.StartTime.Time().Add(7 * 24 * time.Hour).After(v.EndTime.Time()) {
			panic(fmt.Sprintf("总力战排期时间错误 排期不足7天,SeasonId:%v", v.SeasonId))
		}
	}
	logger.Info("总力战排期读取成功文件:%s 读取成功,解析数量:%v", name, len(g.GetGPP().RaidSchedule.RaidScheduleMap))
}

func GetCurRaidSchedule() *RaidScheduleInfo {
	cur := GC.GetGPP().RaidSchedule.CurRaidSchedule
	if cur != nil {
		if cur.EndTime.Time().After(time.Now()) {
			return GC.GetGPP().RaidSchedule.CurRaidSchedule
		}
		next := GC.GetGPP().RaidSchedule.RaidScheduleMap[cur.NextSeasonId]
		if next != nil && next.StartTime.Time().After(time.Now()) {
			return GC.GetGPP().RaidSchedule.CurRaidSchedule
		}
	}
	cur = nil
	for _, v := range GC.GetGPP().RaidSchedule.RaidScheduleMap {
		// 排期中
		if v.EndTime.Time().After(time.Now()) && time.Now().After(v.StartTime.Time()) {
			cur = v
		} else { // 排期间隔中
			next := GC.GetGPP().RaidSchedule.RaidScheduleMap[v.NextSeasonId]
			if next != nil &&
				time.Now().After(v.EndTime.Time()) && next.StartTime.Time().After(time.Now()) {
				cur = v
			}
		}
		if cur != nil {
			GC.GetGPP().RaidSchedule.CurRaidSchedule = cur
			return cur
		}
	}
	return cur
}

func GetNextRaidSchedule() *RaidScheduleInfo {
	cur := GetCurRaidSchedule()
	if cur == nil {
		return nil
	}
	return GC.GetGPP().RaidSchedule.RaidScheduleMap[cur.NextSeasonId]
}

func GetRaidScheduleMap() map[int64]*RaidScheduleInfo {
	return GC.GetGPP().RaidSchedule.RaidScheduleMap
}

func GetRaidScheduleInfo(seasonId int64) *RaidScheduleInfo {
	return GC.GetGPP().RaidSchedule.RaidScheduleMap[seasonId]
}
