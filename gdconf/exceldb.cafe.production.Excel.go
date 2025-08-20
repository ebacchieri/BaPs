package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadCafeProductionExcel() {
	g.GetExcel().CafeProductionExcel = make([]*sro.CafeProductionExcel, 0)
	name := "CafeProductionExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CafeProductionExcel)
}

type CafeProductionExcel struct {
	CafeProductionExcelMap map[int64]map[int32]map[int64]*sro.CafeProductionExcel
}

func (g *GameConfig) gppCafeProductionExcel() {
	g.GetGPP().CafeProductionExcel = &CafeProductionExcel{
		CafeProductionExcelMap: make(map[int64]map[int32]map[int64]*sro.CafeProductionExcel),
	}

	for _, v := range g.GetExcel().GetCafeProductionExcel() {
		if g.GetGPP().CafeProductionExcel.CafeProductionExcelMap[v.CafeId] == nil {
			g.GetGPP().CafeProductionExcel.CafeProductionExcelMap[v.CafeId] = make(map[int32]map[int64]*sro.CafeProductionExcel)
		}
		if g.GetGPP().CafeProductionExcel.CafeProductionExcelMap[v.CafeId][v.Rank] == nil {
			g.GetGPP().CafeProductionExcel.CafeProductionExcelMap[v.CafeId][v.Rank] = make(map[int64]*sro.CafeProductionExcel)
		}
		g.GetGPP().CafeProductionExcel.CafeProductionExcelMap[v.CafeId][v.Rank][v.CafeProductionParcelId] = v
	}

	logger.Info("处理咖啡厅生产配置表完成,数量:%v个",
		len(g.GetGPP().CafeProductionExcel.CafeProductionExcelMap))
}

func GetCafeProductionExcelList(cafeId int64, rank int32) map[int64]*sro.CafeProductionExcel {
	if GC.GetGPP().CafeProductionExcel.CafeProductionExcelMap[cafeId] == nil {
		return nil
	}
	return GC.GetGPP().CafeProductionExcel.CafeProductionExcelMap[cafeId][rank]
}
