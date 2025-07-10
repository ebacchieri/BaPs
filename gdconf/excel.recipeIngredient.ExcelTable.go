package gdconf

import (
	sro "github.com/ebacchieri/BaPs/common/server_only"
	"github.com/ebacchieri/BaPs/pkg/logger"
)

func (g *GameConfig) loadRecipeIngredientExcelTable() {
	g.GetExcel().RecipeIngredientExcelTable = make([]*sro.RecipeIngredientExcelTable, 0)
	name := "RecipeIngredientExcelTable.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RecipeIngredientExcelTable)
}

type RecipeIngredientExcel struct {
	RecipeIngredientExcelMap map[int64]*sro.RecipeIngredientExcelTable
}

func (g *GameConfig) gppRecipeIngredientExcelTable() {
	g.GetGPP().RecipeIngredientExcel = &RecipeIngredientExcel{
		RecipeIngredientExcelMap: make(map[int64]*sro.RecipeIngredientExcelTable),
	}

	for _, v := range g.GetExcel().GetRecipeIngredientExcelTable() {
		g.GetGPP().RecipeIngredientExcel.RecipeIngredientExcelMap[v.Id] = v
	}

	logger.Info("处理材料配置表完成,技能配置:%v个",
		len(g.GetGPP().RecipeIngredientExcel.RecipeIngredientExcelMap))
}

func GetRecipeIngredientExcelTable(id int64) *sro.RecipeIngredientExcelTable {
	return GC.GetGPP().RecipeIngredientExcel.RecipeIngredientExcelMap[id]
}
