package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadFurnitureExcel() {
	g.GetExcel().FurnitureExcel = make([]*sro.FurnitureExcel, 0)
	name := "FurnitureExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().FurnitureExcel)
}

type FurnitureExcel struct {
	FurnitureExcelMap map[int64]*sro.FurnitureExcel
}

func (g *GameConfig) gppFurnitureExcel() {
	g.GetGPP().FurnitureExcel = &FurnitureExcel{
		FurnitureExcelMap: make(map[int64]*sro.FurnitureExcel, 0),
	}
	for _, v := range g.GetExcel().GetFurnitureExcel() {
		g.GetGPP().FurnitureExcel.FurnitureExcelMap[v.Id] = v
	}
	logger.Info("处理家具配置表完成数量:%v个", len(g.GetGPP().FurnitureExcel.FurnitureExcelMap))
}

func GetFurnitureExcel(id int64) *sro.FurnitureExcel {
	return GC.GetGPP().FurnitureExcel.FurnitureExcelMap[id]
}

func GetFurnitureExcelMap() []*sro.FurnitureExcel {
	return GC.GetExcel().GetFurnitureExcel()
}
