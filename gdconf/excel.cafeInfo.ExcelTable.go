package gdconf

import (
	sro "github.com/ebacchieri/BaPs/common/server_only"
	"github.com/ebacchieri/BaPs/pkg/logger"
)

func (g *GameConfig) loadCafeInfoExcelTable() {
	g.GetExcel().CafeInfoExcelTableInfo = make([]*sro.CafeInfoExcelTableInfo, 0)
	name := "CafeInfoExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CafeInfoExcelTableInfo)
}

type CafeInfoExcel struct {
	CafeInfoExcelTableMap map[int64]*sro.CafeInfoExcelTableInfo
}

func (g *GameConfig) gppCafeInfoExcelTable() {
	g.GetGPP().CafeInfoExcel = &CafeInfoExcel{
		CafeInfoExcelTableMap: make(map[int64]*sro.CafeInfoExcelTableInfo, 0),
	}

	for _, v := range g.GetExcel().GetCafeInfoExcelTableInfo() {
		g.GetGPP().CafeInfoExcel.CafeInfoExcelTableMap[v.CafeId] = v
	}

	logger.Info("处理咖啡厅数量完成:%v个", len(g.GetGPP().CafeInfoExcel.CafeInfoExcelTableMap))
}

func GetCafeInfoExcelTables() map[int64]*sro.CafeInfoExcelTableInfo {
	if g := GC.GetGPP(); g == nil {
		return nil
	} else {
		return g.CafeInfoExcel.CafeInfoExcelTableMap
	}
}

func GetCafeInfoExcelTableInfo(id int64) *sro.CafeInfoExcelTableInfo {
	return GC.GetGPP().CafeInfoExcel.CafeInfoExcelTableMap[id]
}
