package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadGoodsExcel() {
	g.GetExcel().GoodsExcel = make([]*sro.GoodsExcel, 0)
	name := "GoodsExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().GoodsExcel)
}

type GoodsExcel struct {
	GoodsExcelMap map[int64]*sro.GoodsExcel
}

func (g *GameConfig) gppGoodsExcel() {
	g.GetGPP().GoodsExcel = &GoodsExcel{
		GoodsExcelMap: make(map[int64]*sro.GoodsExcel),
	}
	for _, v := range g.GetExcel().GetGoodsExcel() {
		g.GetGPP().GoodsExcel.GoodsExcelMap[v.Id] = v
	}

	logger.Info("处理商品配置完成,成就:%v个",
		len(g.GetGPP().GoodsExcel.GoodsExcelMap))
}

func GetGoodsExcel(id int64) *sro.GoodsExcel {
	return GC.GetGPP().GoodsExcel.GoodsExcelMap[id]
}
