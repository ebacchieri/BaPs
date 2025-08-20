package gdconf

import (
	"time"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadTimeAttackDungeonSeasonManageExcelTable() {
	g.GetExcel().TimeAttackDungeonSeasonManageExcel = make([]*sro.TimeAttackDungeonSeasonManageExcel, 0)
	name := "TimeAttackDungeonSeasonManageExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().TimeAttackDungeonSeasonManageExcel)
}

type TimeAttackDungeonSeasonManageExcel struct {
	Cur                                   *sro.TimeAttackDungeonSeasonManageExcel
	TimeAttackDungeonSeasonManageExcelMap map[int64]*sro.TimeAttackDungeonSeasonManageExcel
}

func (g *GameConfig) gppTimeAttackDungeonSeasonManageExcelTable() {
	g.GetGPP().TimeAttackDungeonSeasonManageExcel = &TimeAttackDungeonSeasonManageExcel{
		TimeAttackDungeonSeasonManageExcelMap: make(map[int64]*sro.TimeAttackDungeonSeasonManageExcel, 0),
	}

	for _, v := range g.GetExcel().GetTimeAttackDungeonSeasonManageExcel() {
		g.GetGPP().TimeAttackDungeonSeasonManageExcel.TimeAttackDungeonSeasonManageExcelMap[v.Id] = v
	}

	logger.Info("处理综合战术考试排期表完成,数量:%v个", len(g.GetGPP().TimeAttackDungeonSeasonManageExcel.TimeAttackDungeonSeasonManageExcelMap))
}

func GetTimeAttackDungeonSeasonManageExcelById(id int64) *sro.TimeAttackDungeonSeasonManageExcel {
	return GC.GetGPP().TimeAttackDungeonSeasonManageExcel.TimeAttackDungeonSeasonManageExcelMap[id]
}

func GetCurTimeAttackDungeonSeasonManageExcelTable() *sro.TimeAttackDungeonSeasonManageExcel {
	conf := GC.GetGPP().TimeAttackDungeonSeasonManageExcel
	if conf == nil {
		return nil
	}
	var nextStartTime time.Time
	if conf.Cur != nil {
		next := GetTimeAttackDungeonSeasonManageExcelById(conf.Cur.Id + 1)
		if next != nil {
			nextStartTime, _ = time.Parse("2006-01-02 15:04:05", next.StartDate)
		}
	}
	getCur := func() {
		for index, v := range GC.GetExcel().GetTimeAttackDungeonSeasonManageExcel() { // 读取原始文件,保证顺序
			startTime, err := time.Parse("2006-01-02 15:04:05", v.StartDate)
			endTime, err := time.Parse("2006-01-02 15:04:05", v.EndDate)
			if err != nil {
				logger.Error("综合战术考试排期表时间格式错误")
				return
			}
			next := GC.GetExcel().GetTimeAttackDungeonSeasonManageExcel()[index+1]
			if next != nil {
				nextStartTime, _ = time.Parse("2006-01-02 15:04:05", next.StartDate)
			}

			if (time.Now().After(startTime) && time.Now().Before(endTime)) ||
				(time.Now().After(endTime) && time.Now().Before(nextStartTime)) { // 上期结束且下期未开启
				conf.Cur = v
				return
			}
		}
		logger.Warn("找不到当前综合战术考试排期")
	}

	if conf.Cur == nil || nextStartTime.After(time.Now()) {
		getCur()
	}
	return conf.Cur
}
