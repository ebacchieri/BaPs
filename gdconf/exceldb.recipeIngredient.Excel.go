package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadRecipeIngredientExcel() {
	g.GetExcel().RecipeIngredientExcel = make([]*sro.RecipeIngredientExcel, 0)
	name := "RecipeIngredientExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().RecipeIngredientExcel)
}

type RecipeIngredientExcel struct {
	RecipeIngredientExcelMap map[int64]*sro.RecipeIngredientExcel
}

func (g *GameConfig) gppRecipeIngredientExcel() {
	g.GetGPP().RecipeIngredientExcel = &RecipeIngredientExcel{
		RecipeIngredientExcelMap: make(map[int64]*sro.RecipeIngredientExcel),
	}

	for _, v := range g.GetExcel().GetRecipeIngredientExcel() {
		g.GetGPP().RecipeIngredientExcel.RecipeIngredientExcelMap[v.Id] = v
	}

	logger.Info("处理材料配置表完成,技能配置:%v个",
		len(g.GetGPP().RecipeIngredientExcel.RecipeIngredientExcelMap))
}

func GetRecipeIngredientExcel(id int64) *sro.RecipeIngredientExcel {
	return GC.GetGPP().RecipeIngredientExcel.RecipeIngredientExcelMap[id]
}
