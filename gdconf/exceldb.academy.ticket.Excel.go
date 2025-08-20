package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadAcademyTicketExcel() {
	g.GetExcel().AcademyTicketExcel = make([]*sro.AcademyTicketExcel, 0)
	name := "AcademyTicketExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().AcademyTicketExcel)
}

type AcademyTicketExcel struct {
	AcademyTicketExcelMap map[int64]*sro.AcademyTicketExcel
}

func (g *GameConfig) gppAcademyTicketExcel() {
	g.GetGPP().AcademyTicketExcel = &AcademyTicketExcel{
		AcademyTicketExcelMap: make(map[int64]*sro.AcademyTicketExcel),
	}
	for _, v := range g.GetExcel().GetAcademyTicketExcel() {
		g.GetGPP().AcademyTicketExcel.AcademyTicketExcelMap[v.ScheduleTicktetMax] = v
	}

	logger.Info("处理课程表最大票信息完成,数量:%v个",
		len(g.GetGPP().AcademyTicketExcel.AcademyTicketExcelMap))
}

func GetScheduleTicktetMax(level int64) int64 {
	for i := int64(3); ; i++ {
		conf := GC.GetGPP().AcademyTicketExcel.AcademyTicketExcelMap[i]
		if conf == nil {
			return i - 1
		}
		if conf.LocationRankSum > level {
			return i - 1
		}
	}
}
