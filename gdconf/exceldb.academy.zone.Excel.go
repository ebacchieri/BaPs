package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyZoneExcel() {
	g.GetExcel().AcademyZoneExcel = make([]*sro.AcademyZoneExcel, 0)
	name := "AcademyZoneExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyZoneExcel)
}

type AcademyZoneExcel struct {
	AcademyZoneExcelMap map[int64]*sro.AcademyZoneExcel
}

func (g *GameConfig) gppAcademyZoneExcel() {
	g.GetGPP().AcademyZoneExcel = &AcademyZoneExcel{
		AcademyZoneExcelMap: make(map[int64]*sro.AcademyZoneExcel),
	}
	for _, v := range g.GetExcel().GetAcademyZoneExcel() {
		g.GetGPP().AcademyZoneExcel.AcademyZoneExcelMap[v.Id] = v
	}

	logger.Info("处理课程表教室信息完成,数量:%v个",
		len(g.GetGPP().AcademyZoneExcel.AcademyZoneExcelMap))
}

func GetAcademyZoneExcelList() []*sro.AcademyZoneExcel {
	return GC.GetExcel().GetAcademyZoneExcel()
}

func GetAcademyZoneExcel(zoneId int64) *sro.AcademyZoneExcel {
	return GC.GetGPP().AcademyZoneExcel.AcademyZoneExcelMap[zoneId]
}
