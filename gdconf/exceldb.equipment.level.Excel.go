package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadEquipmentLevelExcel() {
	g.GetExcel().EquipmentLevelExcel = make([]*sro.EquipmentLevelExcel, 0)
	name := "EquipmentLevelExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().EquipmentLevelExcel)
}

type EquipmentLevelExcel struct {
	EquipmentLevelExcelMap map[int32]*sro.EquipmentLevelExcel
}

func (g *GameConfig) gppEquipmentLevelExcel() {
	g.GetGPP().EquipmentLevelExcel = &EquipmentLevelExcel{
		EquipmentLevelExcelMap: make(map[int32]*sro.EquipmentLevelExcel, 0),
	}
	for _, v := range g.GetExcel().GetEquipmentLevelExcel() {
		g.GetGPP().EquipmentLevelExcel.EquipmentLevelExcelMap[v.Level] = v
	}
	logger.Info("装备等级配置表完成数量:%v个", len(g.GetGPP().EquipmentLevelExcel.EquipmentLevelExcelMap))
}

func GetEquipmentLevelExcel(level int32) *sro.EquipmentLevelExcel {
	return GC.GetGPP().EquipmentLevelExcel.EquipmentLevelExcelMap[level]
}
