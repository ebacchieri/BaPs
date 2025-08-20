package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadGachaElementExcel() {
	g.GetExcel().GachaElementExcel = make([]*sro.GachaElementExcel, 0)
	name := "GachaElementExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().GachaElementExcel)
}

type GachaElementExcel struct {
	GachaElementExcelMap    map[int64]*sro.GachaElementExcel
	GachaElementGroupIdList map[int64]*GachaElementGroupId // GroupId
}

type GachaElementGroupId struct {
	GachaGroupId          int64
	Rarity                string
	GachaElementExcelList []*sro.GachaElementExcel
}

func (g *GameConfig) gppGachaElementExcel() {
	g.GetGPP().GachaElementExcel = &GachaElementExcel{
		GachaElementExcelMap:    make(map[int64]*sro.GachaElementExcel),
		GachaElementGroupIdList: make(map[int64]*GachaElementGroupId),
	}
	for _, v := range g.GetExcel().GetGachaElementExcel() {
		g.GetGPP().GachaElementExcel.GachaElementExcelMap[v.Id] = v

		if g.GetGPP().GachaElementExcel.GachaElementGroupIdList[v.GachaGroupId] == nil {
			g.GetGPP().GachaElementExcel.GachaElementGroupIdList[v.GachaGroupId] = &GachaElementGroupId{
				GachaGroupId:          v.GachaGroupId,
				Rarity:                v.Rarity,
				GachaElementExcelList: make([]*sro.GachaElementExcel, 0),
			}
		}
		g.GetGPP().GachaElementExcel.GachaElementGroupIdList[v.GachaGroupId].GachaElementExcelList = append(
			g.GetGPP().GachaElementExcel.GachaElementGroupIdList[v.GachaGroupId].GachaElementExcelList,
			v,
		)
	}

	logger.Info("处理随机组池配置完成,成就:%v个",
		len(g.GetGPP().GachaElementExcel.GachaElementExcelMap))
}

func GetGachaElementExcel(id int64) *sro.GachaElementExcel {
	return GC.GetGPP().GachaElementExcel.GachaElementExcelMap[id]
}

func GetGachaElementGroupIdByGachaGroupId(gachaGroupId int64) *GachaElementGroupId {
	return GC.GetGPP().GachaElementExcel.GachaElementGroupIdList[gachaGroupId]
}
