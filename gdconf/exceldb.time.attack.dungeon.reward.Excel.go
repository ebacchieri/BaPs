package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadTimeAttackDungeonRewardExcel() {
	g.GetExcel().TimeAttackDungeonRewardExcel = make([]*sro.TimeAttackDungeonRewardExcel, 0)
	name := "TimeAttackDungeonRewardExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().TimeAttackDungeonRewardExcel)
}

type TimeAttackDungeonRewardExcel struct {
	TimeAttackDungeonRewardExcelMap map[int64]*sro.TimeAttackDungeonRewardExcel
}

func (g *GameConfig) gppTimeAttackDungeonRewardExcel() {
	g.GetGPP().TimeAttackDungeonRewardExcel = &TimeAttackDungeonRewardExcel{
		TimeAttackDungeonRewardExcelMap: make(map[int64]*sro.TimeAttackDungeonRewardExcel),
	}

	for _, v := range g.GetExcel().GetTimeAttackDungeonRewardExcel() {
		g.GetGPP().TimeAttackDungeonRewardExcel.TimeAttackDungeonRewardExcelMap[v.Id] = v
	}

	logger.Info("处理综合战术考试关卡配置完成,技能配置:%v个",
		len(g.GetGPP().TimeAttackDungeonRewardExcel.TimeAttackDungeonRewardExcelMap))
}

func GetTimeAttackDungeonRewardExcel(id int64) *sro.TimeAttackDungeonRewardExcel {
	return GC.GetGPP().TimeAttackDungeonRewardExcel.TimeAttackDungeonRewardExcelMap[id]
}
