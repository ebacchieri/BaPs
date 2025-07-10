package gdconf

import (
	sro "github.com/ebacchieri/BaPs/common/server_only"
	"github.com/ebacchieri/BaPs/pkg/logger"
)

func (g *GameConfig) loadCharacterLevelExcelTable() {
	g.GetExcel().CharacterLevelExcelTable = make([]*sro.CharacterLevelExcelTable, 0)
	name := "CharacterLevelExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().CharacterLevelExcelTable)
}

type CharacterLevelExcel struct {
	CharacterLevelExcelTableMap map[int32]*sro.CharacterLevelExcelTable
}

func (g *GameConfig) gppCharacterLevelExcelTable() {
	g.GetGPP().CharacterLevelExcel = &CharacterLevelExcel{
		CharacterLevelExcelTableMap: make(map[int32]*sro.CharacterLevelExcelTable, 0),
	}
	for _, v := range g.GetExcel().GetCharacterLevelExcelTable() {
		g.GetGPP().CharacterLevelExcel.CharacterLevelExcelTableMap[v.Level] = v
	}
	logger.Info("处理角色等级配置表完成数量:%v个", len(g.GetGPP().CharacterLevelExcel.CharacterLevelExcelTableMap))
}

func GetCharacterLevelExcelTable(level int32) *sro.CharacterLevelExcelTable {
	return GC.GetGPP().CharacterLevelExcel.CharacterLevelExcelTableMap[level]
}

func UpCharacterLevel(level int32, exp int64) (int32, int64) {
	for {
		conf := GetCharacterLevelExcelTable(level)
		if conf == nil {
			return level - 1, exp
		}
		if exp < conf.Exp {
			return level, exp
		}
		exp -= conf.Exp
		level++
	}
}
