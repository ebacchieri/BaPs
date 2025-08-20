package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyMessangerExcel() {
	g.GetExcel().AcademyMessangerExcel = make([]*sro.AcademyMessangerExcel, 0)
	name := "AcademyMessangerExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyMessangerExcel)
}

type AcademyMessangerExcel struct {
	AcademyMessangerExcelMap map[int64]*sro.AcademyMessangerExcel
}

func (g *GameConfig) gppAcademyMessangerExcel() {
	g.GetGPP().AcademyMessangerExcel = &AcademyMessangerExcel{
		AcademyMessangerExcelMap: make(map[int64]*sro.AcademyMessangerExcel),
	}

	for _, v := range g.GetExcel().GetAcademyMessangerExcel() {
		g.GetGPP().AcademyMessangerExcel.AcademyMessangerExcelMap[v.MessageGroupId] = v
	}

	logger.Info("处理MomoTalk对话配置完成,MomoTalk对话:%v个",
		len(g.GetGPP().AcademyMessangerExcel.AcademyMessangerExcelMap))
}

func GetAcademyMessangerExcel(gid int64) *sro.AcademyMessangerExcel {
	return GC.GetGPP().AcademyMessangerExcel.AcademyMessangerExcelMap[gid]
}
