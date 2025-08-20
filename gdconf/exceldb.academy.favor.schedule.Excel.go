package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyFavorScheduleExcel() {
	g.GetExcel().AcademyFavorScheduleExcel = make([]*sro.AcademyFavorScheduleExcel, 0)
	name := "AcademyFavorScheduleExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyFavorScheduleExcel)
}

type AcademyFavorScheduleExcel struct {
	AcademyFavorScheduleExcelMap map[int64]*sro.AcademyFavorScheduleExcel
}

func (g *GameConfig) gppAcademyFavorScheduleExcel() {
	g.GetGPP().AcademyFavorScheduleExcel = &AcademyFavorScheduleExcel{
		AcademyFavorScheduleExcelMap: make(map[int64]*sro.AcademyFavorScheduleExcel),
	}
	for _, v := range g.GetExcel().GetAcademyFavorScheduleExcel() {
		g.GetGPP().AcademyFavorScheduleExcel.AcademyFavorScheduleExcelMap[v.Id] = v
	}

	logger.Info("处理MomoTalk剧情配置完成,剧情:%v个",
		len(g.GetGPP().AcademyFavorScheduleExcel.AcademyFavorScheduleExcelMap))
}

func GetAcademyFavorScheduleExcel(id int64) *sro.AcademyFavorScheduleExcel {
	return GC.GetGPP().AcademyFavorScheduleExcel.AcademyFavorScheduleExcelMap[id]
}
