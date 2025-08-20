package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadRaidRankingRewardExcel() {
	g.GetExcel().RaidRankingRewardExcel = make([]*sro.RaidRankingRewardExcel, 0)
	name := "RaidRankingRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RaidRankingRewardExcel)
}

type RaidRankingRewardExcel struct {
	RaidRankingRewardExcelMap map[int64][]*sro.RaidRankingRewardExcel
}

func (g *GameConfig) gppRaidRankingRewardExcel() {
	g.GetGPP().RaidRankingRewardExcel = &RaidRankingRewardExcel{
		RaidRankingRewardExcelMap: make(map[int64][]*sro.RaidRankingRewardExcel),
	}

	for _, v := range g.GetExcel().GetRaidRankingRewardExcel() {
		if g.GetGPP().RaidRankingRewardExcel.RaidRankingRewardExcelMap[v.RankingRewardGroupId] == nil {
			g.GetGPP().RaidRankingRewardExcel.RaidRankingRewardExcelMap[v.RankingRewardGroupId] =
				make([]*sro.RaidRankingRewardExcel, 0)
		}
		g.GetGPP().RaidRankingRewardExcel.RaidRankingRewardExcelMap[v.RankingRewardGroupId] =
			append(g.GetGPP().RaidRankingRewardExcel.RaidRankingRewardExcelMap[v.RankingRewardGroupId], v)
	}

	logger.Info("处理总力战结算奖励配置表完成,奖励配置:%v个",
		len(g.GetGPP().RaidRankingRewardExcel.RaidRankingRewardExcelMap))
}

func GetRaidRankingRewardExcel(gid, ranking int64) *sro.RaidRankingRewardExcel {
	for _, conf := range GC.GetGPP().RaidRankingRewardExcel.RaidRankingRewardExcelMap[gid] {
		if conf.RankStart <= ranking && (conf.RankEnd >= ranking || conf.RankEnd == 0) {
			return conf
		}
		if ranking <= 0 && conf.RankEnd == 0 {
			return conf
		}
	}
	return nil
}
