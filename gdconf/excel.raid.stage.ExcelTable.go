package gdconf

import (
	sro "./common/server_only"
	"./pkg/logger"
)

func (g *GameConfig) loadRaidStageExcelTable() {
	g.GetExcel().RaidStageExcelTable = make([]*sro.RaidStageExcelTable, 0)
	name := "RaidStageExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RaidStageExcelTable)
}

type RaidStageExcel struct {
	RaidStageExcelMap map[int64]*sro.RaidStageExcelTable
}

func (g *GameConfig) gppRaidStageExcelTable() {
	g.GetGPP().RaidStageExcel = &RaidStageExcel{
		RaidStageExcelMap: make(map[int64]*sro.RaidStageExcelTable),
	}

	for _, v := range g.GetExcel().GetRaidStageExcelTable() {
		g.GetGPP().RaidStageExcel.RaidStageExcelMap[v.Id] = v
	}

	logger.Info("处理总力战关卡配置表完成,关卡配置:%v个",
		len(g.GetGPP().RaidStageExcel.RaidStageExcelMap))
}

func GetRaidStageExcelTable(id int64) *sro.RaidStageExcelTable {
	return GC.GetGPP().RaidStageExcel.RaidStageExcelMap[id]
}
