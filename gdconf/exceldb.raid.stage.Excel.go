package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadRaidStageExcel() {
	g.GetExcel().RaidStageExcel = make([]*sro.RaidStageExcel, 0)
	name := "RaidStageExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RaidStageExcel)
}

type RaidStageExcel struct {
	RaidStageExcelMap map[int64]*sro.RaidStageExcel
}

func (g *GameConfig) gppRaidStageExcel() {
	g.GetGPP().RaidStageExcel = &RaidStageExcel{
		RaidStageExcelMap: make(map[int64]*sro.RaidStageExcel),
	}

	for _, v := range g.GetExcel().GetRaidStageExcel() {
		g.GetGPP().RaidStageExcel.RaidStageExcelMap[v.Id] = v
	}

	logger.Info("处理总力战关卡配置表完成,关卡配置:%v个",
		len(g.GetGPP().RaidStageExcel.RaidStageExcelMap))
}

func GetRaidStageExcel(id int64) *sro.RaidStageExcel {
	return GC.GetGPP().RaidStageExcel.RaidStageExcelMap[id]
}
