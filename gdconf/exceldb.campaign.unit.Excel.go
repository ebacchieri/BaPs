package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadCampaignUnitExcel() {
	g.GetExcel().CampaignUnitExcel = make([]*sro.CampaignUnitExcel, 0)
	name := "CampaignUnitExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CampaignUnitExcel)
}

type CampaignUnitExcel struct {
	CampaignUnitExcelMap      map[int64]*sro.CampaignUnitExcel
	CampaignUnitExcelStageMap map[int64]*CampaignUnitExcelGrade
}

type CampaignUnitExcelGrade struct {
	Boss      *sro.CampaignUnitExcel
	GradeList map[string]*sro.CampaignUnitExcel
}

func (g *GameConfig) gppCampaignUnitExcel() {
	info := &CampaignUnitExcel{
		CampaignUnitExcelMap:      make(map[int64]*sro.CampaignUnitExcel, 0),
		CampaignUnitExcelStageMap: make(map[int64]*CampaignUnitExcelGrade),
	}

	for _, v := range g.GetExcel().GetCampaignUnitExcel() {
		info.CampaignUnitExcelMap[v.Id] = v
		stageId := v.Id / 100
		if info.CampaignUnitExcelStageMap[stageId] == nil {
			info.CampaignUnitExcelStageMap[stageId] = &CampaignUnitExcelGrade{
				GradeList: make(map[string]*sro.CampaignUnitExcel, 0),
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
