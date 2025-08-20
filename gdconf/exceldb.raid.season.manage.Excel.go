package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadRaidSeasonManageExcel() {
	g.GetExcel().RaidSeasonManageExcel = make([]*sro.RaidSeasonManageExcel, 0)
	name := "RaidSeasonManageExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RaidSeasonManageExcel)
}

type RaidSeasonManageExcel struct {
	RaidSeasonManageExcelMap map[int64]*sro.RaidSeasonManageExcel
}

func (g *GameConfig) gppRaidSeasonManageExcel() {
	g.GetGPP().RaidSeasonManageExcel = &RaidSeasonManageExcel{
		RaidSeasonManageExcelMap: make(map[int64]*sro.RaidSeasonManageExcel),
	}

	for _, v := range g.GetExcel().GetRaidSeasonManageExcel() {
		g.GetGPP().RaidSeasonManageExcel.RaidSeasonManageExcelMap[v.SeasonId] = v
	}

	logger.Info("处理总力战赛季配置表完成,赛季配置:%v个",
		len(g.GetGPP().RaidSeasonManageExcel.RaidSeasonManageExcelMap))
}

func GetRaidSeasonManageExcel(seasonId int64) *sro.RaidSeasonManageExcel {
	return GC.GetGPP().RaidSeasonManageExcel.RaidSeasonManageExcelMap[seasonId]
}

func GetRaidTier(seasonId, ranking int64) int32 {
	conf := GetRaidSeasonManageExcel(seasonId)
	if conf == nil || ranking < 0 {
		return 1
	}
	return GetRaidRankingRewardExcel(conf.RankingRewardGroupId, ranking).GetTier()
}

func GetRaidRankingRewardExcelBySeasonId(seasonId, ranking int64) *sro.RaidRankingRewardExcel {
	conf := GetRaidSeasonManageExcel(seasonId)
	if conf == nil {
		return nil
	}
	return GetRaidRankingRewardExcel(conf.RankingRewardGroupId, ranking)
}
