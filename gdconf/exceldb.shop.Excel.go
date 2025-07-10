package gdconf

import (
	"time"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadShopExcel() {
	g.GetExcel().ShopExcel = make([]*sro.ShopExcel, 0)
	name := "ShopExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().ShopExcel)
}

type ShopExcel struct {
	ShopExcelMap     map[int64]*sro.ShopExcel
	ShopExcelTypeMap map[string][]*sro.ShopExcel
}

func (g *GameConfig) gppShopExcel() {
	g.GetGPP().ShopExcel = &ShopExcel{
		ShopExcelMap:     make(map[int64]*sro.ShopExcel),
		ShopExcelTypeMap: make(map[string][]*sro.ShopExcel),
	}
	for _, v := range g.GetExcel().GetShopExcel() {
		g.GetGPP().ShopExcel.ShopExcelMap[v.Id] = v
		salePeriodTo, _ := time.Parse("2006-01-02 15:04:05", v.SalePeriodTo)
		if v.SalePeriodTo != "" && time.Now().After(salePeriodTo) {
			continue
		}
		if g.GetGPP().ShopExcel.ShopExcelTypeMap[v.CategoryType] == nil {
			g.GetGPP().ShopExcel.ShopExcelTypeMap[v.CategoryType] = make([]*sro.ShopExcel, 0)
		}
		g.GetGPP().ShopExcel.ShopExcelTypeMap[v.CategoryType] = append(
			g.GetGPP().ShopExcel.ShopExcelTypeMap[v.CategoryType], v)
	}
	logger.Info("处理商品配置完成,商店类型:%v个", len(g.GetGPP().ShopExcel.ShopExcelTypeMap))
}

func GetShopExcelType(categoryType string) []*sro.ShopExcel {
	if GC.GetGPP().ShopExcel.ShopExcelTypeMap == nil {
		return nil
	}
	return GC.GetGPP().ShopExcel.ShopExcelTypeMap[categoryType]
}

func GetShopExcel(shopId int64) *sro.ShopExcel {
	return GC.GetGPP().ShopExcel.ShopExcelMap[shopId]
}
