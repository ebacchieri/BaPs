package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyLocationExcel() {
	g.GetExcel().AcademyLocationExcel = make([]*sro.AcademyLocationExcel, 0)
	name := "AcademyLocationExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyLocationExcel)
}

type AcademyLocationExcel struct {
	AcademyLocationExcelMap map[int64]*sro.AcademyLocationExcel
}

func (g *GameConfig) gppAcademyLocationExcel() {
	g.GetGPP().AcademyLocationExcel = &AcademyLocationExcel{
		AcademyLocationExcelMap: make(map[int64]*sro.AcademyLocationExcel),
	}
	for _, v := range g.GetExcel().GetAcademyLocationExcel() {
		g.GetGPP().AcademyLocationExcel.AcademyLocationExcelMap[v.Id] = v
	}

	logger.Info("处理课程表学院信息完成,数量:%v个",
		len(g.GetGPP().AcademyLocationExcel.AcademyLocationExcelMap))
}

func GetAcademyLocationExcelList() []*sro.AcademyLocationExcel {
	return GC.GetExcel().GetAcademyLocationExcel()
}
