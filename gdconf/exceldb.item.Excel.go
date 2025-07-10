package gdconf

import (
	"sync"
	"time"

	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
)

func (g *GameConfig) loadItemExcel() {
	g.GetExcel().ItemExcel = make([]*sro.ItemExcel, 0)
	name := "ItemExcel.json"
	loadExcelFile(excelDbPath+name, &g.GetExcel().ItemExcel)
}

type ItemExcel struct {
	RecruitCoin          *RecruitCoin
	recruitCoinSync      sync.Mutex
	ItemExcelMap         map[int64]*sro.ItemExcel
	ItemExcelCategoryMap map[string][]*sro.ItemExcel
}

type RecruitCoin struct {
	Item   *sro.ItemExcel
	EnTime time.Time
}

func (g *GameConfig) gppItemExcel() {
	g.GetGPP().ItemExcel = &ItemExcel{
		ItemExcelMap:         make(map[int64]*sro.ItemExcel),
		ItemExcelCategoryMap: make(map[string][]*sro.ItemExcel),
		recruitCoinSync:      sync.Mutex{},
	}

	for _, v := range g.GetExcel().GetItemExcel() {
		if v.ExpirationDateTime != "" {
			enTime, err := time.Parse("2006-01-02 15:04:05", v.ExpirationDateTime)
			if err != nil {
				continue
			}
			if time.Now().After(enTime) {
				continue
			}
		}
		g.GetGPP().ItemExcel.ItemExcelMap[v.Id] = v
		if g.GetGPP().ItemExcel.ItemExcelCategoryMap[v.ItemCategory] == nil {
			g.GetGPP().ItemExcel.ItemExcelCategoryMap[v.ItemCategory] = make([]*sro.ItemExcel, 0)
		}
		g.GetGPP().ItemExcel.ItemExcelCategoryMap[v.ItemCategory] = append(
			g.GetGPP().ItemExcel.ItemExcelCategoryMap[v.ItemCategory], v)
	}

	logger.Info("处理道具配置完成,道具:%v个,类型:%v个", len(g.GetGPP().ItemExcel.ItemExcelMap),
		len(g.GetGPP().ItemExcel.ItemExcelCategoryMap))
}

// 弃用,已经有更好的方法了
// func GetRecruitCoin() *sro.ItemExcel {
// 	bin := GC.GetGPP().ItemExcel
// 	bin.recruitCoinSync.Lock()
// 	defer bin.recruitCoinSync.Unlock()
// 	if bin.RecruitCoin != nil &&
// 		!time.Now().After(bin.RecruitCoin.EnTime) {
// 		return bin.RecruitCoin.Item
// 	}
// 	confList := GC.GetGPP().ItemExcel.ItemExcelCategoryMap["RecruitCoin"]
// 	for _, conf := range confList {
// 		if conf.ExpirationDateTime == "2099-12-31 23:59:59" {
// 			continue
// 		}
// 		enTime, err := time.Parse("2006-01-02 15:04:05", conf.ExpirationDateTime)
// 		if err != nil {
// 			continue
// 		}
// 		if !time.Now().After(enTime) {
// 			bin.RecruitCoin = &RecruitCoin{
// 				Item:   conf,
// 				EnTime: enTime,
// 			}
// 			return conf
// 		}
//
// 	}
// 	return nil
// }

func GetItemExcelCategoryMap() map[string][]*sro.ItemExcel {
	return GC.GetGPP().ItemExcel.ItemExcelCategoryMap
}

func GetItemExcelCategoryMapByCategory(itemCategory string) []*sro.ItemExcel {
	return GC.GetGPP().ItemExcel.ItemExcelCategoryMap[itemCategory]
}

func GetItemExcel(id int64) *sro.ItemExcel {
	return GC.GetGPP().ItemExcel.ItemExcelMap[id]
}

func IsItem(id int64) bool {
	return GC.GetGPP().ItemExcel.ItemExcelMap[id] != nil
}
