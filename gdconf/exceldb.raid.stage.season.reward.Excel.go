package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadRaidStageSeasonRewardExcel() {
	g.GetExcel().RaidStageSeasonRewardExcel = make([]*sro.RaidStageSeasonRewardExcel, 0)
	name := "RaidStageSeasonRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RaidStageSeasonRewardExcel)
}

type RaidStageSeasonRewardExcel struct {
	RaidStageSeasonRewardExcelMap map[int64]*sro.RaidStageSeasonRewardExcel
}

func (g *GameConfig) gppRaidStageSeasonRewardExcel() {
	g.GetGPP().RaidStageSeasonRewardExcel = &RaidStageSeasonRewardExcel{
		RaidStageSeasonRewardExcelMap: make(map[int64]*sro.RaidStageSeasonRewardExcel),
	}

	for _, v := range g.GetExcel().GetRaidStageSeasonRewardExcel() {
		g.GetGPP().RaidStageSeasonRewardExcel.RaidStageSeasonRewardExcelMap[v.SeasonRewardId] = v
	}

	logger.Info("处理总力战总分奖励配置表完成,总力战总分奖励配置:%v个",
		len(g.GetGPP().RaidStageSeasonRewardExcel.RaidStageSeasonRewardExcelMap))
}

func GetRaidStageSeasonRewardExcel(id int64) *sro.RaidStageSeasonRewardExcel {
	return GC.GetGPP().RaidStageSeasonRewardExcel.RaidStageSeasonRewardExcelMap[id]
}
