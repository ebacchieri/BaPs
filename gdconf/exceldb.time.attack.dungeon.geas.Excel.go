package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadTimeAttackDungeonGeasExcel() {
	g.GetExcel().TimeAttackDungeonGeasExcel = make([]*sro.TimeAttackDungeonGeasExcel, 0)
	name := "TimeAttackDungeonGeasExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().TimeAttackDungeonGeasExcel)
}

type TimeAttackDungeonGeasExcel struct {
	TimeAttackDungeonGeasExcelMap map[int64]*sro.TimeAttackDungeonGeasExcel
}

func (g *GameConfig) gppTimeAttackDungeonGeasExcel() {
	g.GetGPP().TimeAttackDungeonGeasExcel = &TimeAttackDungeonGeasExcel{
		TimeAttackDungeonGeasExcelMap: make(map[int64]*sro.TimeAttackDungeonGeasExcel),
	}

	for _, v := range g.GetExcel().GetTimeAttackDungeonGeasExcel() {
		g.GetGPP().TimeAttackDungeonGeasExcel.TimeAttackDungeonGeasExcelMap[v.Id] = v
	}

	logger.Info("处理综合战术考试关卡配置完成,技能配置:%v个",
		len(g.GetGPP().TimeAttackDungeonGeasExcel.TimeAttackDungeonGeasExcelMap))
}

func GetTimeAttackDungeonGeasExcel(id int64) *sro.TimeAttackDungeonGeasExcel {
	return GC.GetGPP().TimeAttackDungeonGeasExcel.TimeAttackDungeonGeasExcelMap[id]
}
