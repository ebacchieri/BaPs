package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadRaidStageRewardExcel() {
	g.GetExcel().RaidStageRewardExcel = make([]*sro.RaidStageRewardExcel, 0)
	name := "RaidStageRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RaidStageRewardExcel)
}

type RaidStageRewardExcel struct {
	RaidStageRewardExcelMap map[int64][]*sro.RaidStageRewardExcel
}

func (g *GameConfig) gppRaidStageRewardExcel() {
	g.GetGPP().RaidStageRewardExcel = &RaidStageRewardExcel{
		RaidStageRewardExcelMap: make(map[int64][]*sro.RaidStageRewardExcel),
	}

	for _, v := range g.GetExcel().GetRaidStageRewardExcel() {
		if g.GetGPP().RaidStageRewardExcel.RaidStageRewardExcelMap[v.GroupId] == nil {
			g.GetGPP().RaidStageRewardExcel.RaidStageRewardExcelMap[v.GroupId] = make([]*sro.RaidStageRewardExcel, 0)
		}
		g.GetGPP().RaidStageRewardExcel.RaidStageRewardExcelMap[v.GroupId] =
			append(g.GetGPP().RaidStageRewardExcel.RaidStageRewardExcelMap[v.GroupId], v)
	}

	logger.Info("处理总力战关卡通过奖励配置表完成,总力战关卡通过奖励配置:%v个",
		len(g.GetGPP().RaidStageRewardExcel.RaidStageRewardExcelMap))
}

func GetRaidStageRewardExcel(gid int64) []*sro.RaidStageRewardExcel {
	return GC.GetGPP().RaidStageRewardExcel.RaidStageRewardExcelMap[gid]
}
