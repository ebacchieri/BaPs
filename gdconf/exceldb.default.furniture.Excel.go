package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
)

func (g *GameConfig) loadDefaultFurnitureExcel() {
	g.GetExcel().DefaultFurnitureExcel = make([]*sro.DefaultFurnitureExcel, 0)
	name := "DefaultFurnitureExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().DefaultFurnitureExcel)
}

func GetDefaultFurnitureExcelList() []*sro.DefaultFurnitureExcel {
	return GC.GetExcel().GetDefaultFurnitureExcel()
}
