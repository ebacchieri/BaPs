package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/protocol/proto"
)

func (g *GameConfig) loadEliminateRaidSeasonManageExcel() {
	g.GetExcel().EliminateRaidSeasonManageExcel = make([]*sro.EliminateRaidSeasonManageExcel, 0)
	name := "EliminateRaidSeasonManageExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EliminateRaidSeasonManageExcel)
}

type EliminateRaidSeasonManageExcel struct {
	EliminateRaidSeasonManageExcelMap map[int64]*sro.EliminateRaidSeasonManageExcel
}

func (g *GameConfig) gppEliminateRaidSeasonManageExcel() {
	g.GetGPP().EliminateRaidSeasonManageExcel = &EliminateRaidSeasonManageExcel{
		EliminateRaidSeasonManageExcelMap: make(map[int64]*sro.EliminateRaidSeasonManageExcel),
	}

	for _, v := range g.GetExcel().GetEliminateRaidSeasonManageExcel() {
		g.GetGPP().EliminateRaidSeasonManageExcel.EliminateRaidSeasonManageExcelMap[v.SeasonId] = v
	}

	logger.Info("处理大决战赛季配置表完成,赛季配置:%v个",
		len(g.GetGPP().EliminateRaidSeasonManageExcel.EliminateRaidSeasonManageExcelMap))
}

func GetEliminateRaidSeasonManageExcel(seasonId int64) *sro.EliminateRaidSeasonManageExcel {
	return GC.GetGPP().EliminateRaidSeasonManageExcel.EliminateRaidSeasonManageExcelMap[seasonId]
}

func GetEliminateRaidTier(seasonId, ranking int64) int32 {
	conf := GetEliminateRaidSeasonManageExcel(seasonId)
	if conf == nil || ranking < 0 {
		return 1
	}
	return GetEliminateRaidRankingRewardExcel(conf.RankingRewardGroupId, ranking).GetTier()
}

func GetEliminateRaidRankingRewardExcelBySeasonId(seasonId, ranking int64) *sro.EliminateRaidRankingRewardExcel {
	conf := GetEliminateRaidSeasonManageExcel(seasonId)
	if conf == nil {
		return nil
	}
	return GetEliminateRaidRankingRewardExcel(conf.RankingRewardGroupId, ranking)
}

func GetEliminateRaidEchelonType(seasonId int64, raidBossGroup string) proto.EchelonType {
	conf := GetEliminateRaidSeasonManageExcel(seasonId)
	if conf == nil {
		return proto.EchelonType_None
	}
	switch raidBossGroup {
	case conf.OpenRaidBossGroup01:
		return proto.EchelonType_EliminateRaid01
	case conf.OpenRaidBossGroup02:
		return proto.EchelonType_EliminateRaid02
	case conf.OpenRaidBossGroup03:
		return proto.EchelonType_EliminateRaid03
	default:
		return proto.EchelonType_None
	}
}
