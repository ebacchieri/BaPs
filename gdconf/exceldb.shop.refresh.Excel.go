package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadShopRefreshExcel() {
	g.GetExcel().ShopRefreshExcel = make([]*sro.ShopRefreshExcel, 0)
	name := "ShopRefreshExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().ShopRefreshExcel)
}

type ShopRefreshExcel struct {
	ShopRefreshExcelMap  map[string][]*sro.ShopRefreshExcel
	ShopRefreshExcelList map[int64]*sro.ShopRefreshExcel
}

func (g *GameConfig) gppShopRefreshExcel() {
	g.GetGPP().ShopRefreshExcel = &ShopRefreshExcel{
		ShopRefreshExcelMap:  make(map[string][]*sro.ShopRefreshExcel),
		ShopRefreshExcelList: make(map[int64]*sro.ShopRefreshExcel),
	}
	for _, v := range g.GetExcel().GetShopRefreshExcel() {
		g.GetGPP().ShopRefreshExcel.ShopRefreshExcelList[v.Id] = v
		if g.GetGPP().ShopRefreshExcel.ShopRefreshExcelMap[v.CategoryType] == nil {
			g.GetGPP().ShopRefreshExcel.ShopRefreshExcelMap[v.CategoryType] = make([]*sro.ShopRefreshExcel, 0)
		}
		g.GetGPP().ShopRefreshExcel.ShopRefreshExcelMap[v.CategoryType] = append(
			g.GetGPP().ShopRefreshExcel.ShopRefreshExcelMap[v.CategoryType],
			v,
		)
	}
	logger.Info("处理可刷新商品配置完成,商店类型:%v个", len(g.GetGPP().ShopRefreshExcel.ShopRefreshExcelMap))
}

func GetShopRefreshExcelMap(categoryType string) []*sro.ShopRefreshExcel {
	return GC.GetGPP().ShopRefreshExcel.ShopRefreshExcelMap[categoryType]
}

func GetShopRefreshExcel(shopId int64) *sro.ShopRefreshExcel {
	return GC.GetGPP().ShopRefreshExcel.ShopRefreshExcelList[shopId]
}
