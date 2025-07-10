package gdconf

import (
	"encoding/json"
	"github.com/gucooing/BaPs/pkg/logger"
	"os"
	"net/http"
	"io"
)

type ProdIndex struct {
	Notices                       []*Notice                      `json:"Notices"`
	Events                        []*Event                       `json:"Events"`
	Issues                        []interface{}                  `json:"Issues"`
	Maintenance                   *Maintenance                   `json:"Maintenance"`
	Banners                       []*Banner                      `json:"Banners"`
	ServerStatus                  int                            `json:"ServerStatus"`
	LatestClientVersion           string                         `json:"LatestClientVersion"`
	GachaProbabilityDisplay       []*GachaProbabilityDisplay     `json:"GachaProbabilityDisplay"`
	MarketAppleId                 interface{}                    `json:"MarketAppleId"`
	NotificationBeforeMaintenance *NotificationBeforeMaintenance `json:"NotificationBeforeMaintenance"`
	ContentLock                   []interface{}                  `json:"ContentLock"`
	GachaPeriodDisplay            []*GachaPeriodDisplay          `json:"GachaPeriodDisplay"`
	RaidPeriodDisplay             interface{}                    `json:"RaidPeriodDisplay"`
	Survey                        *Survey                        `json:"Survey"`
	GuidePopup                    []*GuidePopup                  `json:"GuidePopup"`
	DownloadLimitTime             float64                        `json:"DownloadLimitTime"`
	EnableSQLite                  bool                           `json:"EnableSQLite"`
}

type Banner struct {
	BannerId            int      `json:"BannerId"`
	StartDate           string   `json:"StartDate"`
	EndDate             string   `json:"EndDate"`
	Url                 string   `json:"Url"`
	FileName            []string `json:"FileName"`
	LinkedLobbyBannerId int      `json:"LinkedLobbyBannerId"`
	BannerType          int      `json:"BannerType"`
	BannerDisplayType   int      `json:"BannerDisplayType"`
}

type Event struct {
	NoticeId     int    `json:"NoticeId"`
	StartDate    string `json:"StartDate"`
	EndDate      string `json:"EndDate"`
	Url          string `json:"Url"`
	Title        string `json:"Title"`
	DisplayOrder int    `json:"DisplayOrder"`
}

type GachaPeriodDisplay struct {
	GachaPeriodDisplayId int    `json:"GachaPeriodDisplayId"`
	Text                 string `json:"Text"`
}

type GachaProbabilityDisplay struct {
	GachaProbabilityDisplayId int    `json:"GachaProbabilityDisplayId"`
	Url                       string `json:"Url"`
	LinkedLobbyBannerId       int    `json:"LinkedLobbyBannerId"`
	GachaDisplayTag           int    `json:"GachaDisplayTag"`
}

type GuidePopup struct {
	GuidePopupId   int         `json:"GuidePopupId"`
	GuidePopupType int         `json:"GuidePopupType"`
	PopupType      int         `json:"PopupType"`
	StartDate      string      `json:"StartDate"`
	EndDate        string      `json:"EndDate"`
	FileName       string      `json:"FileName"`
	Url            string      `json:"Url"`
	Message        string      `json:"Message"`
	SurveyId       int         `json:"SurveyId"`
	NotifyUrl      interface{} `json:"NotifyUrl"`
	GotoUrl        *string     `json:"GotoUrl"`
	DisplayOrder   int         `json:"DisplayOrder"`
	PopupOKText    string      `json:"PopupOKText"`
}

type Maintenance struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Text      string `json:"Text"`
}

type Notice struct {
	NoticeId     int    `json:"NoticeId"`
	StartDate    string `json:"StartDate"`
	EndDate      string `json:"EndDate"`
	Url          string `json:"Url"`
	Title        string `json:"Title"`
	DisplayOrder int    `json:"DisplayOrder"`
}

type NotificationBeforeMaintenance struct {
	PopupType int    `json:"PopupType"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Text      string `json:"Text"`
}

type Survey struct {
	SurveyId  int         `json:"SurveyId"`
	PopupType int         `json:"PopupType"`
	StartDate string      `json:"StartDate"`
	EndDate   string      `json:"EndDate"`
	FileName  string      `json:"FileName"`
	Url       string      `json:"Url"`
	Text      string      `json:"Text"`
	NotifyUrl interface{} `json:"NotifyUrl"`
}

func (g *GameConfig) loadProdIndex() {
load:
	g.GetGPP().ProdIndex = new(ProdIndex)
	name := "ProdIndex.json"
	file, err := os.ReadFile(g.dataPath + name)
	if err != nil {
		if os.IsNotExist(err) {
			err := g.downloadManagementData(g.dataPath + name)
			if err == nil {
				logger.Info("ProdIndex.json自动下载成功！")
				goto load
			}
		} else {
			logger.Error("文件:%s 读取失败,err:%s", name, err)
			return
		}
	}
	if err := json.Unmarshal(file, &g.GetGPP().ProdIndex); err != nil {
		logger.Error("文件:%s 解析失败,err:%s", name, err)
		return
	}
	logger.Info("公告配置读取成功文件:%s ", name)
}

func (g *GameConfig) downloadManagementData(path string) error {
	resp, err := http.Get(g.managementDataUrl)
	if err != nil {
		logger.Error("下载ProdIndex.json失败,请手动下载")
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}

func GetProdIndex() *ProdIndex {
	return GC.GPP.ProdIndex
}
