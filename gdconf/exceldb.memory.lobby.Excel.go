package gdconf

import (
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadMemoryLobbyExcel() {
	g.GetExcel().MemoryLobbyExcel = make([]*sro.MemoryLobbyExcel, 0)
	name := "MemoryLobbyExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().MemoryLobbyExcel)
}

type MemoryLobbyExcel struct {
	MemoryLobbyExcelList map[int64]*sro.MemoryLobbyExcel
}

func (g *GameConfig) gppMemoryLobbyExcel() {
	g.GetGPP().MemoryLobbyExcel = &MemoryLobbyExcel{
		MemoryLobbyExcelList: make(map[int64]*sro.MemoryLobbyExcel),
	}
	for _, v := range g.GetExcel().GetMemoryLobbyExcel() {
		g.GetGPP().MemoryLobbyExcel.MemoryLobbyExcelList[v.Id] = v
	}
	logger.Info("处理记忆大厅配置完成,记忆大厅:%v个", len(g.GetGPP().MemoryLobbyExcel.MemoryLobbyExcelList))
}

func GetMemoryLobbyExcelList() map[int64]*sro.MemoryLobbyExcel {
	return GC.GetGPP().MemoryLobbyExcel.MemoryLobbyExcelList
}
