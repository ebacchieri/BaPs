package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadCampaignStageExcel() {
	g.GetExcel().CampaignStageExcel = make([]*sro.CampaignStageExcel, 0)
	name := "CampaignStageExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CampaignStageExcel)
}

type CampaignStageExcel struct {
	CampaignStageExcelMap map[int64]*sro.CampaignStageExcel
}

func (g *GameConfig) gppCampaignStageExcel() {
	g.GetGPP().CampaignStageExcel = &CampaignStageExcel{
		CampaignStageExcelMap: make(map[int64]*sro.CampaignStageExcel, 0),
	}

	for _, v := range g.GetExcel().GetCampaignStageExcel() {
		g.GetGPP().CampaignStageExcel.CampaignStageExcelMap[v.Id] = v
	}

	logger.Info("处理任务关卡信息数量完成:%v个", len(g.GetGPP().CampaignStageExcel.CampaignStageExcelMap))
}

func GetCampaignStageExcel(id int64) *sro.CampaignStageExcel {
	return GC.GetGPP().CampaignStageExcel.CampaignStageExcelMap[id]
}

func GetCampaignStageExcelMap() map[int64]*sro.CampaignStageExcel {
	return GC.GetGPP().CampaignStageExcel.CampaignStageExcelMap
}
