package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyLocationRankExcel() {
	g.GetExcel().AcademyLocationRankExcel = make([]*sro.AcademyLocationRankExcel, 0)
	name := "AcademyLocationRankExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyLocationRankExcel)
}

type AcademyLocationRankExcel struct {
	AcademyLocationRankExcelMap map[int64]*sro.AcademyLocationRankExcel
}

func (g *GameConfig) gppAcademyLocationRankExcel() {
	g.GetGPP().AcademyLocationRankExcel = &AcademyLocationRankExcel{
		AcademyLocationRankExcelMap: make(map[int64]*sro.AcademyLocationRankExcel),
	}
	for _, v := range g.GetExcel().GetAcademyLocationRankExcel() {
		g.GetGPP().AcademyLocationRankExcel.AcademyLocationRankExcelMap[v.Rank] = v
	}

	logger.Info("处理课程表等级配置完成,数量:%v个",
		len(g.GetGPP().AcademyLocationRankExcel.AcademyLocationRankExcelMap))
}

func GetAcademyLocationRankExcel(rank int64) *sro.AcademyLocationRankExcel {
	return GC.GetGPP().AcademyLocationRankExcel.AcademyLocationRankExcelMap[rank]
}
