package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadWeekDungeonRewardExcel() {
	g.GetExcel().WeekDungeonRewardExcel = make([]*sro.WeekDungeonRewardExcel, 0)
	name := "WeekDungeonRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().WeekDungeonRewardExcel)
}

type WeekDungeonRewardExcel struct {
	WeekDungeonRewardExcelMap map[int64][]*sro.WeekDungeonRewardExcel
}

func (g *GameConfig) gppWeekDungeonRewardExcel() {
	g.GetGPP().WeekDungeonRewardExcel = &WeekDungeonRewardExcel{
		WeekDungeonRewardExcelMap: make(map[int64][]*sro.WeekDungeonRewardExcel, 0),
	}

	for _, v := range g.GetExcel().GetWeekDungeonRewardExcel() {
		if g.GetGPP().WeekDungeonRewardExcel.WeekDungeonRewardExcelMap[v.GroupId] == nil {
			g.GetGPP().WeekDungeonRewardExcel.WeekDungeonRewardExcelMap[v.GroupId] = make([]*sro.WeekDungeonRewardExcel, 0)
		}
		g.GetGPP().WeekDungeonRewardExcel.WeekDungeonRewardExcelMap[v.GroupId] = append(
			g.GetGPP().WeekDungeonRewardExcel.WeekDungeonRewardExcelMap[v.GroupId], v)
	}

	logger.Info("处理悬赏通缉关卡奖励配置完成,数量:%v个", len(g.GetGPP().WeekDungeonRewardExcel.WeekDungeonRewardExcelMap))
}

func GetWeekDungeonRewardExcelList(stageId int64) []*sro.WeekDungeonRewardExcel {
	return GC.GetGPP().WeekDungeonRewardExcel.WeekDungeonRewardExcelMap[stageId]
}
