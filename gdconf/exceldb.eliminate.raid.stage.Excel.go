package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEliminateRaidStageExcel() {
	g.GetExcel().EliminateRaidStageExcel = make([]*sro.EliminateRaidStageExcel, 0)
	name := "EliminateRaidStageExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EliminateRaidStageExcel)
}

type EliminateRaidStageExcel struct {
	EliminateRaidStageExcelMap map[int64]*sro.EliminateRaidStageExcel
}

func (g *GameConfig) gppEliminateRaidStageExcel() {
	g.GetGPP().EliminateRaidStageExcel = &EliminateRaidStageExcel{
		EliminateRaidStageExcelMap: make(map[int64]*sro.EliminateRaidStageExcel),
	}

	for _, v := range g.GetExcel().GetEliminateRaidStageExcel() {
		g.GetGPP().EliminateRaidStageExcel.EliminateRaidStageExcelMap[v.Id] = v
	}

	logger.Info("处理大决战关卡配置表完成,关卡配置:%v个",
		len(g.GetGPP().EliminateRaidStageExcel.EliminateRaidStageExcelMap))
}

func GetEliminateRaidStageExcel(id int64) *sro.EliminateRaidStageExcel {
	return GC.GetGPP().EliminateRaidStageExcel.EliminateRaidStageExcelMap[id]
}
