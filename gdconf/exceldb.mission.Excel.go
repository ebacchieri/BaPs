package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadMissionExcelTable() {
	g.GetExcel().MissionExcel = make([]*sro.MissionExcel, 0)
	name := "MissionExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().MissionExcel)
}

type MissionExcel struct {
	MissionExcelTableMap      map[int64]*sro.MissionExcel
	MissionExcelTableCategory map[string][]*sro.MissionExcel
}

func (g *GameConfig) gppMissionExcelTable() {
	g.GetGPP().MissionExcel = &MissionExcel{
		MissionExcelTableMap:      make(map[int64]*sro.MissionExcel, 0),
		MissionExcelTableCategory: make(map[string][]*sro.MissionExcel),
	}
	for _, v := range g.GetExcel().GetMissionExcel() {
		g.GetGPP().MissionExcel.MissionExcelTableMap[v.Id] = v
		if g.GetGPP().MissionExcel.MissionExcelTableCategory[v.Category] == nil {
			g.GetGPP().MissionExcel.MissionExcelTableCategory[v.Category] = make([]*sro.MissionExcel, 0)
		}
		g.GetGPP().MissionExcel.MissionExcelTableCategory[v.Category] =
			append(g.GetGPP().MissionExcel.MissionExcelTableCategory[v.Category], v)
	}
	logger.Info("处理任务配置表完成数量:%v个", len(g.GetGPP().MissionExcel.MissionExcelTableMap))
}

func GetMissionExcelTableCategoryList(category string) []*sro.MissionExcel {
	return GC.GetGPP().MissionExcel.MissionExcelTableCategory[category]
}

func GetMissionExcelTable(id int64) *sro.MissionExcel {
	return GC.GetGPP().MissionExcel.MissionExcelTableMap[id]
}
