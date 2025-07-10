package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadGuideMissionExcel() {
	g.GetExcel().GuideMissionExcel = make([]*sro.GuideMissionExcel, 0)
	name := "GuideMissionExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().GuideMissionExcel)
}

type GuideMissionExcel struct {
	GuideMissionExcelMap map[int64]*sro.GuideMissionExcel
}

func (g *GameConfig) gppGuideMissionExcel() {
	g.GetGPP().GuideMissionExcel = &GuideMissionExcel{
		GuideMissionExcelMap: make(map[int64]*sro.GuideMissionExcel),
	}
	for _, v := range g.GetExcel().GetGuideMissionExcel() {
		g.GetGPP().GuideMissionExcel.GuideMissionExcelMap[v.Id] = v
	}

	logger.Info("处理成就配置完成,成就:%v个",
		len(g.GetGPP().GuideMissionExcel.GuideMissionExcelMap))
}

func GetGuideMissionExcel(id int64) *sro.GuideMissionExcel {
	return GC.GetGPP().GuideMissionExcel.GuideMissionExcelMap[id]
}
