package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadCafeInfoExcel() {
	g.GetExcel().CafeInfoExcel = make([]*sro.CafeInfoExcel, 0)
	name := "CafeInfoExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CafeInfoExcel)
}

type CafeInfoExcel struct {
	CafeInfoExcelMap map[int64]*sro.CafeInfoExcel
}

func (g *GameConfig) gppCafeInfoExcel() {
	g.GetGPP().CafeInfoExcel = &CafeInfoExcel{
		CafeInfoExcelMap: make(map[int64]*sro.CafeInfoExcel, 0),
	}

	for _, v := range g.GetExcel().GetCafeInfoExcel() {
		g.GetGPP().CafeInfoExcel.CafeInfoExcelMap[v.CafeId] = v
	}

	logger.Info("处理咖啡厅数量完成:%v个", len(g.GetGPP().CafeInfoExcel.CafeInfoExcelMap))
}

func GetCafeInfoExcels() map[int64]*sro.CafeInfoExcel {
	if g := GC.GetGPP(); g == nil {
		return nil
	} else {
		return g.CafeInfoExcel.CafeInfoExcelMap
	}
}

func GetCafeInfoExcel(id int64) *sro.CafeInfoExcel {
	return GC.GetGPP().CafeInfoExcel.CafeInfoExcelMap[id]
}
