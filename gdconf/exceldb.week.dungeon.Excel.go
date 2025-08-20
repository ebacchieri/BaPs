package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadWeekDungeonExcel() {
	g.GetExcel().WeekDungeonExcel = make([]*sro.WeekDungeonExcel, 0)
	name := "WeekDungeonExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().WeekDungeonExcel)
}

type WeekDungeonExcel struct {
	WeekDungeonExcelMap map[int64]*sro.WeekDungeonExcel
}

func (g *GameConfig) gppWeekDungeonExcel() {
	g.GetGPP().WeekDungeonExcel = &WeekDungeonExcel{
		WeekDungeonExcelMap: make(map[int64]*sro.WeekDungeonExcel, 0),
	}

	for _, v := range g.GetExcel().GetWeekDungeonExcel() {
		if len(v.StarGoal) != len(v.StarGoalAmount) {
			logger.Warn("WeekDungeonExcel.json StarGoal和StarGoalAmount不对应")
		}
		g.GetGPP().WeekDungeonExcel.WeekDungeonExcelMap[v.StageId] = v
	}

	logger.Info("悬赏通缉关卡信息数量完成:%v个", len(g.GetGPP().WeekDungeonExcel.WeekDungeonExcelMap))
}

func GetWeekDungeonExcel(stageId int64) *sro.WeekDungeonExcel {
	return GC.GetGPP().WeekDungeonExcel.WeekDungeonExcelMap[stageId]
}
