package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEliminateRaidStageSeasonRewardExcel() {
	g.GetExcel().EliminateRaidStageSeasonRewardExcel = make([]*sro.EliminateRaidStageSeasonRewardExcel, 0)
	name := "EliminateRaidStageSeasonRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EliminateRaidStageSeasonRewardExcel)
}

type EliminateRaidStageSeasonRewardExcel struct {
	EliminateRaidStageSeasonRewardExcelMap map[int64]*sro.EliminateRaidStageSeasonRewardExcel
}

func (g *GameConfig) gppEliminateRaidStageSeasonRewardExcel() {
	g.GetGPP().EliminateRaidStageSeasonRewardExcel = &EliminateRaidStageSeasonRewardExcel{
		EliminateRaidStageSeasonRewardExcelMap: make(map[int64]*sro.EliminateRaidStageSeasonRewardExcel),
	}

	for _, v := range g.GetExcel().GetEliminateRaidStageSeasonRewardExcel() {
		g.GetGPP().EliminateRaidStageSeasonRewardExcel.EliminateRaidStageSeasonRewardExcelMap[v.SeasonRewardId] = v
	}

	logger.Info("处理大决战总分奖励配置表完成,大决战总分奖励配置:%v个",
		len(g.GetGPP().EliminateRaidStageSeasonRewardExcel.EliminateRaidStageSeasonRewardExcelMap))
}

func GetEliminateRaidStageSeasonRewardExcel(id int64) *sro.EliminateRaidStageSeasonRewardExcel {
	return GC.GetGPP().EliminateRaidStageSeasonRewardExcel.EliminateRaidStageSeasonRewardExcelMap[id]
}
