package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEquipmentStatExcel() {
	g.GetExcel().EquipmentStatExcel = make([]*sro.EquipmentStatExcel, 0)
	name := "EquipmentStatExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().EquipmentStatExcel)
}

type EquipmentStatExcel struct {
	EquipmentStatExcelMap map[int64]*sro.EquipmentStatExcel
}

func (g *GameConfig) gppEquipmentStatExcel() {
	g.GetGPP().EquipmentStatExcel = &EquipmentStatExcel{
		EquipmentStatExcelMap: make(map[int64]*sro.EquipmentStatExcel, 0),
	}
	for _, v := range g.GetExcel().GetEquipmentStatExcel() {
		g.GetGPP().EquipmentStatExcel.EquipmentStatExcelMap[v.EquipmentId] = v
	}
	logger.Info("装备详情表完成数量:%v个", len(g.GetGPP().EquipmentStatExcel.EquipmentStatExcelMap))
}

func GetEquipmentStatExcel(eId int64) *sro.EquipmentStatExcel {
	return GC.GetGPP().EquipmentStatExcel.EquipmentStatExcelMap[eId]
}
