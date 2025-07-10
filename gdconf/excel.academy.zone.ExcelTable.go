package gdconf

import (
	sro "./common/server_only"
	"./pkg/logger"
)

func (g *GameConfig) loadAcademyZoneExcelTable() {
	g.GetExcel().AcademyZoneExcelTable = make([]*sro.AcademyZoneExcelTable, 0)
	name := "AcademyZoneExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyZoneExcelTable)
}

type AcademyZoneExcel struct {
	AcademyZoneExcelMap map[int64]*sro.AcademyZoneExcelTable
}

func (g *GameConfig) gppAcademyZoneExcelTable() {
	g.GetGPP().AcademyZoneExcel = &AcademyZoneExcel{
		AcademyZoneExcelMap: make(map[int64]*sro.AcademyZoneExcelTable),
	}
	for _, v := range g.GetExcel().GetAcademyZoneExcelTable() {
		g.GetGPP().AcademyZoneExcel.AcademyZoneExcelMap[v.Id] = v
	}

	logger.Info("处理课程表教室信息完成,数量:%v个",
		len(g.GetGPP().AcademyZoneExcel.AcademyZoneExcelMap))
}

func GetAcademyZoneExcelTableList() []*sro.AcademyZoneExcelTable {
	return GC.GetExcel().GetAcademyZoneExcelTable()
}

func GetAcademyZoneExcelTable(zoneId int64) *sro.AcademyZoneExcelTable {
	return GC.GetGPP().AcademyZoneExcel.AcademyZoneExcelMap[zoneId]
}
