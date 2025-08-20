package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEliminateRaidStageRewardExcel() {
	g.GetExcel().EliminateRaidStageRewardExcel = make([]*sro.EliminateRaidStageRewardExcel, 0)
	name := "EliminateRaidStageRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().EliminateRaidStageRewardExcel)
}

type EliminateRaidStageRewardExcel struct {
	EliminateRaidStageRewardExcelMap map[int64][]*sro.EliminateRaidStageRewardExcel
}

func (g *GameConfig) gppEliminateRaidStageRewardExcel() {
	g.GetGPP().EliminateRaidStageRewardExcel = &EliminateRaidStageRewardExcel{
		EliminateRaidStageRewardExcelMap: make(map[int64][]*sro.EliminateRaidStageRewardExcel),
	}

	for _, v := range g.GetExcel().GetEliminateRaidStageRewardExcel() {
		if g.GetGPP().EliminateRaidStageRewardExcel.EliminateRaidStageRewardExcelMap[v.GroupId] == nil {
			g.GetGPP().EliminateRaidStageRewardExcel.EliminateRaidStageRewardExcelMap[v.GroupId] = make([]*sro.EliminateRaidStageRewardExcel, 0)
		}
		g.GetGPP().EliminateRaidStageRewardExcel.EliminateRaidStageRewardExcelMap[v.GroupId] =
			append(g.GetGPP().EliminateRaidStageRewardExcel.EliminateRaidStageRewardExcelMap[v.GroupId], v)
	}

	logger.Info("处理大决战关卡通过奖励配置表完成,总力战关卡通过奖励配置:%v个",
		len(g.GetGPP().EliminateRaidStageRewardExcel.EliminateRaidStageRewardExcelMap))
}

func GetEliminateRaidStageRewardExcel(gid int64) []*sro.EliminateRaidStageRewardExcel {
	return GC.GetGPP().EliminateRaidStageRewardExcel.EliminateRaidStageRewardExcelMap[gid]
}
