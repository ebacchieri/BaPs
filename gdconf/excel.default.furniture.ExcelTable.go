package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
)

func (g *GameConfig) loadDefaultFurnitureExcelTable() {
	g.GetExcel().DefaultFurnitureExcelTable = make([]*sro.DefaultFurnitureExcelTable, 0)
	name := "DefaultFurnitureExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().DefaultFurnitureExcelTable)
}

func GetDefaultFurnitureExcelList() []*sro.DefaultFurnitureExcelTable {
	return GC.GetExcel().GetDefaultFurnitureExcelTable()
}
