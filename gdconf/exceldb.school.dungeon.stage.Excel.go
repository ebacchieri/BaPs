package gdconf

import (
	sro "./common/server_only"
	"./pkg/logger"
)

func (g *GameConfig) loadSchoolDungeonStageExcel() {
	g.GetExcel().SchoolDungeonStageExcel = make([]*sro.SchoolDungeonStageExcel, 0)
	name := "SchoolDungeonStageExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().SchoolDungeonStageExcel)
}

type SchoolDungeonStage struct {
	SchoolDungeonStageMap map[int64]*sro.SchoolDungeonStageExcel
}

func (g *GameConfig) gppSchoolDungeonStageExcel() {
	g.GetGPP().SchoolDungeonStage = &SchoolDungeonStage{
		SchoolDungeonStageMap: make(map[int64]*sro.SchoolDungeonStageExcel, 0),
	}

	for _, v := range g.GetExcel().GetSchoolDungeonStageExcel() {
		g.GetGPP().SchoolDungeonStage.SchoolDungeonStageMap[v.StageId] = v
	}

	logger.Info("学院交流会关卡信息数量完成:%v个", len(g.GetGPP().SchoolDungeonStage.SchoolDungeonStageMap))
}

func GetSchoolDungeonStageExcel(stageId int64) *sro.SchoolDungeonStageExcel {
	return GC.GetGPP().SchoolDungeonStage.SchoolDungeonStageMap[stageId]
}
