package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadCafeRankExcel() {
	g.GetExcel().CafeRankExcel = make([]*sro.CafeRankExcel, 0)
	name := "CafeRankExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CafeRankExcel)
}

type CafeRankExcel struct {
	CafeRankExcelMap map[int64]map[int32]*sro.CafeRankExcel
}

func (g *GameConfig) gppCafeRankExcel() {
	g.GetGPP().CafeRankExcel = &CafeRankExcel{
		CafeRankExcelMap: make(map[int64]map[int32]*sro.CafeRankExcel),
	}

	for _, v := range g.GetExcel().GetCafeRankExcel() {
		if g.GetGPP().CafeRankExcel.CafeRankExcelMap[v.CafeId] == nil {
			g.GetGPP().CafeRankExcel.CafeRankExcelMap[v.CafeId] = make(map[int32]*sro.CafeRankExcel)
		}
		g.GetGPP().CafeRankExcel.CafeRankExcelMap[v.CafeId][v.Rank] = v
	}

	logger.Info("处理咖啡厅等级配置表完成,数量:%v个",
		len(g.GetGPP().CafeRankExcel.CafeRankExcelMap))
}

func GetCafeRankExcel(cafeId int64, rank int32) *sro.CafeRankExcel {
	if GC.GetGPP().CafeRankExcel.CafeRankExcelMap[cafeId] == nil {
		return nil
	}
	return GC.GetGPP().CafeRankExcel.CafeRankExcelMap[cafeId][rank]
}
