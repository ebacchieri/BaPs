package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEquipmentExcel() {
	g.GetExcel().EquipmentExcel = make([]*sro.EquipmentExcel, 0)
	name := "EquipmentExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().EquipmentExcel)
}

type EquipmentExcel struct {
	EquipmentExcelMap map[int64]*sro.EquipmentExcel
}

func (g *GameConfig) gppEquipmentExcel() {
	g.GetGPP().EquipmentExcel = &EquipmentExcel{
		EquipmentExcelMap: make(map[int64]*sro.EquipmentExcel, 0),
	}
	for _, v := range g.GetExcel().GetEquipmentExcel() {
		g.GetGPP().EquipmentExcel.EquipmentExcelMap[v.Id] = v
	}
	logger.Info("装备配置表完成数量:%v个", len(g.GetGPP().EquipmentExcel.EquipmentExcelMap))
}

func GetEquipmentExcelMap() []*sro.EquipmentExcel {
	return GC.GetExcel().GetEquipmentExcel()
}

func GetEquipmentExcel(id int64) *sro.EquipmentExcel {
	return GC.GetGPP().EquipmentExcel.EquipmentExcelMap[id]
}
