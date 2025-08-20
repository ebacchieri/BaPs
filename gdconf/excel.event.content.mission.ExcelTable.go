package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEventContentMissionExcelTable() {
	g.GetExcel().EventContentMissionExcelTable = make([]*sro.EventContentMissionExcelTable, 0)
	name := "EventContentMissionExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EventContentMissionExcelTable)
}

type EventContentMissionExcel struct {
	EventContentMissionExcelById             map[int64]*sro.EventContentMissionExcelTable
	EventContentMissionExcelByEventContentId map[int64]*sro.EventContentMissionExcelTable
}

func (g *GameConfig) gppEventContentMissionExcelTable() {
	g.GetGPP().EventContentMissionExcel = &EventContentMissionExcel{
		EventContentMissionExcelById:             make(map[int64]*sro.EventContentMissionExcelTable),
		EventContentMissionExcelByEventContentId: make(map[int64]*sro.EventContentMissionExcelTable),
	}
	for _, v := range g.GetExcel().GetEventContentMissionExcelTable() {
		g.GetGPP().EventContentMissionExcel.EventContentMissionExcelById[v.Id] = v
		g.GetGPP().EventContentMissionExcel.EventContentMissionExcelById[v.EventContentId] = v
	}
	logger.Info("处理活动关卡总表完成数量:%v个", len(g.GetGPP().EventContentMissionExcel.EventContentMissionExcelById))
}
