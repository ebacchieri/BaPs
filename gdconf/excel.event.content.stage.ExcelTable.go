package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEventContentStageExcelTable() {
	g.GetExcel().EventContentStageExcelTable = make([]*sro.EventContentStageExcelTable, 0)
	name := "EventContentStageExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EventContentStageExcelTable)
}

type EventContentStageExcel struct {
	EventContentStageExcelById map[int64]*sro.EventContentStageExcelTable
}

func (g *GameConfig) gppEventContentStageExcelTable() {
	g.GetGPP().EventContentStageExcel = &EventContentStageExcel{
		EventContentStageExcelById: make(map[int64]*sro.EventContentStageExcelTable),
	}
	for _, v := range g.GetExcel().GetEventContentStageExcelTable() {
		g.GetGPP().EventContentStageExcel.EventContentStageExcelById[v.Id] = v
	}
	logger.Info("处理活动关卡详情配置表完成数量:%v个", len(g.GetGPP().EventContentStageExcel.EventContentStageExcelById))
}

func GetEventContentStageExcelTable(id int64) *sro.EventContentStageExcelTable {
	return GC.GetGPP().EventContentStageExcel.EventContentStageExcelById[id]
}
