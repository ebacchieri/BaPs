package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadSkillExcel() {
	g.GetExcel().SkillExcel = make([]*sro.SkillExcel, 0)
	name := "SkillExcel.json"
	loadExcelFile(excelPath+name, &g.GetExcel().SkillExcel)
}

type SkillExcel struct {
	SkillExcelMap map[string]map[int32]*sro.SkillExcel
}

func (g *GameConfig) gppSkillExcel() {
	g.GetGPP().SkillExcel = &SkillExcel{
		SkillExcelMap: make(map[string]map[int32]*sro.SkillExcel),
	}

	for _, v := range g.GetExcel().GetSkillExcel() {
		if g.GetGPP().SkillExcel.SkillExcelMap[v.GroupId] == nil {
			g.GetGPP().SkillExcel.SkillExcelMap[v.GroupId] = make(map[int32]*sro.SkillExcel)
		}
		g.GetGPP().SkillExcel.SkillExcelMap[v.GroupId][v.Level] = v
	}

	logger.Info("处理技能配置完成,技能配置:%v个",
		len(g.GetGPP().SkillExcel.SkillExcelMap))
}

func GetSkillExcel(groupId string, level int32) *sro.SkillExcel {
	list, ok := GC.GetGPP().SkillExcel.SkillExcelMap[groupId]
	if !ok {
		return nil
	}
	return list[level]
}
