package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyRewardExcel() {
	g.GetExcel().AcademyRewardExcel = make([]*sro.AcademyRewardExcel, 0)
	name := "AcademyRewardExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyRewardExcel)
}

type AcademyRewardExcel struct {
	AcademyRewardExcelMap map[int64]map[int64]*sro.AcademyRewardExcel
}

func (g *GameConfig) gppAcademyRewardExcel() {
	g.GetGPP().AcademyRewardExcel = &AcademyRewardExcel{
		AcademyRewardExcelMap: make(map[int64]map[int64]*sro.AcademyRewardExcel),
	}
	for _, v := range g.GetExcel().GetAcademyRewardExcel() {
		if g.GetGPP().AcademyRewardExcel.AcademyRewardExcelMap[v.ScheduleGroupId] == nil {
			g.GetGPP().AcademyRewardExcel.AcademyRewardExcelMap[v.ScheduleGroupId] = make(map[int64]*sro.AcademyRewardExcel)
		}
		g.GetGPP().AcademyRewardExcel.AcademyRewardExcelMap[v.ScheduleGroupId][v.LocationRank] = v
	}

	logger.Info("处理课程表奖励配置完成,数量:%v个",
		len(g.GetGPP().AcademyRewardExcel.AcademyRewardExcelMap))
}

func GetAcademyRewardExcel(gId int64, rank int64) *sro.AcademyRewardExcel {
	confList := GC.GetGPP().AcademyRewardExcel.AcademyRewardExcelMap[gId]
	if confList == nil {
		return nil
	}
	loadRank := int64(0)
	for _, conf := range confList {
		if conf.LocationRank <= rank &&
			loadRank < conf.LocationRank {
			loadRank = conf.LocationRank
		}
	}
	return confList[loadRank]
}
