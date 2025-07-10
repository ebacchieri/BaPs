package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadStickerPageContentExcel() {
	g.GetExcel().StickerPageContentExcel = make([]*sro.StickerPageContentExcel, 0)
	name := "StickerPageContentExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().StickerPageContentExcel)
}

type StickerPageContentExcel struct {
	StickerPageContentExcelList map[int64]*sro.StickerPageContentExcel
}

func (g *GameConfig) gppStickerPageContentExcel() {
	g.GetGPP().StickerPageContentExcel = &StickerPageContentExcel{
		StickerPageContentExcelList: make(map[int64]*sro.StickerPageContentExcel),
	}
	for _, v := range g.GetExcel().GetStickerPageContentExcel() {
		g.GetGPP().StickerPageContentExcel.StickerPageContentExcelList[v.Id] = v
	}
	logger.Info("处理贴纸配置完成,贴纸:%v个", len(g.GetGPP().StickerPageContentExcel.StickerPageContentExcelList))
}

func GetStickerPageContentExcelList() map[int64]*sro.StickerPageContentExcel {
	return GC.GetGPP().StickerPageContentExcel.StickerPageContentExcelList
}
