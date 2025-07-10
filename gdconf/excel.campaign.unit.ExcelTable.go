package gdconf

import (
	sro "github.com/ebacchieri/BaPs/common/server_only"
	"github.com/ebacchieri/BaPs/pkg/logger"
)

func (g *GameConfig) loadCampaignUnitExcelTable() {
	g.GetExcel().CampaignUnitExcelTable = make([]*sro.CampaignUnitExcelTable, 0)
	name := "CampaignUnitExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CampaignUnitExcelTable)
}

type CampaignUnitExcel struct {
	CampaignUnitExcelMap      map[int64]*sro.CampaignUnitExcelTable
	CampaignUnitExcelStageMap map[int64]*CampaignUnitExcelGrade
}

type CampaignUnitExcelGrade struct {
	Boss      *sro.CampaignUnitExcelTable
	GradeList map[string]*sro.CampaignUnitExcelTable
}

func (g *GameConfig) gppCampaignUnitExcelTable() {
	info := &CampaignUnitExcel{
		CampaignUnitExcelMap:      make(map[int64]*sro.CampaignUnitExcelTable, 0),
		CampaignUnitExcelStageMap: make(map[int64]*CampaignUnitExcelGrade),
	}

	for _, v := range g.GetExcel().GetCampaignUnitExcelTable() {
		info.CampaignUnitExcelMap[v.Id] = v
		stageId := v.Id / 100
		if info.CampaignUnitExcelStageMap[stageId] == nil {
			info.CampaignUnitExcelStageMap[stageId] = &CampaignUnitExcelGrade{
				GradeList: make(map[string]*sro.CampaignUnitExcelTable, 0),
			}
		}
		if v.Grade == "Boss" {
			info.CampaignUnitExcelStageMap[stageId].Boss = v
		} else {
			info.CampaignUnitExcelStageMap[stageId].GradeList[v.Grade] = v
		}
	}

	g.GetGPP().CampaignUnitExcel = info

	logger.Info("任务关卡怪物信息关卡数量完成:%v个", len(g.GetGPP().CampaignUnitExcel.CampaignUnitExcelStageMap))
}
