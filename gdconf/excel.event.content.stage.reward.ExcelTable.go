package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEventContentStageRewardExcelTable() {
	g.GetExcel().EventContentStageRewardExcelTable = make([]*sro.EventContentStageRewardExcelTable, 0)
	name := "EventContentStageRewardExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EventContentStageRewardExcelTable)
}

type EventContentStageRewardExcel struct {
	EventContentStageRewardExcelList map[int64][]*sro.EventContentStageRewardExcelTable
}

func (g *GameConfig) gppEventContentStageRewardExcelTable() {
	info := make(map[int64][]*sro.EventContentStageRewardExcelTable)
	g.GetGPP().EventContentStageRewardExcel = &EventContentStageRewardExcel{
		EventContentStageRewardExcelList: info,
	}
	for _, v := range g.GetExcel().GetEventContentStageRewardExcelTable() {
		if info[v.GroupId] == nil {
			info[v.GroupId] = make([]*sro.EventContentStageRewardExcelTable, 0)
		}
		info[v.GroupId] = append(info[v.GroupId], v)
	}
	logger.Info("处理活动关卡奖励配置表完成数量:%v个", len(g.GetGPP().EventContentStageRewardExcel.EventContentStageRewardExcelList))
}

func GetEventContentStageRewardExcelList(eventContentStageRewardId int64) []*sro.EventContentStageRewardExcelTable {
	return GC.GetGPP().EventContentStageRewardExcel.EventContentStageRewardExcelList[eventContentStageRewardId]
}
