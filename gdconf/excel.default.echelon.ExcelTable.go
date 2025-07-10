package gdconf

import (
	sro "github.com/ebacchieri/BaPs/common/server_only"
)

func (g *GameConfig) loadDefaultEchelonExcelTable() {
	g.GetExcel().DefaultEchelonExcelTable = make([]*sro.DefaultEchelonExcelTable, 0)
	name := "DefaultEchelonExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().DefaultEchelonExcelTable)
}

func GetDefaultEchelonExcelList() []*sro.DefaultEchelonExcelTable {
	return GC.GetExcel().GetDefaultEchelonExcelTable()
}
