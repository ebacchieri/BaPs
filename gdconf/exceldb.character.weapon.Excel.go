package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadCharacterWeaponExcel() {
	g.GetExcel().CharacterWeaponExcel = make([]*sro.CharacterWeaponExcel, 0)
	name := "CharacterWeaponExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CharacterWeaponExcel)
}

type CharacterWeaponExcel struct {
	CharacterWeaponExcelMap map[int64]*sro.CharacterWeaponExcel
}

func (g *GameConfig) gppCharacterWeaponExcel() {
	g.GetGPP().CharacterWeaponExcel = &CharacterWeaponExcel{
		CharacterWeaponExcelMap: make(map[int64]*sro.CharacterWeaponExcel),
	}

	for _, v := range g.GetExcel().GetCharacterWeaponExcel() {
		g.GetGPP().CharacterWeaponExcel.CharacterWeaponExcelMap[v.Id] = v
	}

	logger.Info("角色武器配置完成,角色武器:%v个",
		len(g.GetGPP().CharacterWeaponExcel.CharacterWeaponExcelMap))
}

func GetCharacterWeaponExcel(characterId int64) *sro.CharacterWeaponExcel {
	return GC.GetGPP().CharacterWeaponExcel.CharacterWeaponExcelMap[characterId]
}
