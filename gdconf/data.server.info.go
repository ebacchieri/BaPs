package gdconf

import (
	"encoding/json"
	"fmt"
	"github.com/gucooing/BaPs/config"
	"github.com/gucooing/BaPs/pkg/logger"
	"os"
	"sync"
	"time"
)

type ServerInfo struct {
	LoadSync         sync.RWMutex       `json:"-"`
	UpTime           time.Time          `json:"-"`
	ConnectionGroups []*ConnectionGroup `json:"ConnectionGroups"`
}
type ConnectionGroup struct {
	Name                       string                    `json:"Name"`
	ManagementDataUrl          string                    `json:"ManagementDataUrl"`
	IsProductionAddressables   bool                      `json:"IsProductionAddressables"`
	ApiUrl                     string                    `json:"ApiUrl"`
	GatewayUrl                 string                    `json:"GatewayUrl"`
	KibanaLogUrl               string                    `json:"KibanaLogUrl"`
	ProhibitedWordBlackListUri string                    `json:"ProhibitedWordBlackListUri"`
	ProhibitedWordWhiteListUri string                    `json:"ProhibitedWordWhiteListUri"`
	CustomerServiceUrl         string                    `json:"CustomerServiceUrl"`
	OverrideConnectionGroups   []OverrideConnectionGroup `json:"OverrideConnectionGroups"`
	BundleVersion              string                    `json:"BundleVersion"`
}

type OverrideConnectionGroup struct {
	Name                       string `json:"Name"`
	AddressablesCatalogUrlRoot string `json:"AddressablesCatalogUrlRoot"`
}

func (g *GameConfig) loadServerInfo() {
	g.GetGPP().ServerInfo = &ServerInfo{
		LoadSync: sync.RWMutex{},
	}
	name := "ServerInfo.json"
	file, err := os.ReadFile(g.dataPath + name)
	if err != nil {
		logger.Error("文件:%s 读取失败,err:%s", name, err)
		return
	}
	if err := json.Unmarshal(file, &g.GetGPP().ServerInfo); err != nil {
		logger.Error("文件:%s 解析失败,err:%s", name, err)
		return
	}
	g.GetGPP().ServerInfo.ConnectionGroups = GenServerInfo(g.GetGPP().ServerInfo.ConnectionGroups)

	logger.Info("服务器配置读取成功文件:%s ", name)
}

func GetServerInfo() *ServerInfo {
	return GC.GPP.ServerInfo
}

func GenServerInfo(list []*ConnectionGroup) []*ConnectionGroup {
	connectionGroups := make([]*ConnectionGroup, 0)
	for _, v := range list {
		connectionGroups = append(connectionGroups, &ConnectionGroup{
			Name:                       v.Name,
			ManagementDataUrl:          config.GetOtherAddr().GetManagementDataUrl(),
			IsProductionAddressables:   true,
			ApiUrl:                     fmt.Sprintf("%s/api/", config.GetHttpNet().GetOuterAddr()),
			GatewayUrl:                 fmt.Sprintf("%s/getEnterTicket/", config.GetHttpNet().GetOuterAddr()),
			KibanaLogUrl:               "https://prod-logcollector.bluearchiveyostar.com:5300", //    fmt.Sprintf("%s/client/log/", s.GetOuterAddr()),
			ProhibitedWordBlackListUri: v.ProhibitedWordBlackListUri,
			ProhibitedWordWhiteListUri: v.ProhibitedWordWhiteListUri,
			CustomerServiceUrl:         v.CustomerServiceUrl,
			OverrideConnectionGroups:   v.OverrideConnectionGroups,
			BundleVersion:              v.BundleVersion,
		})
	}
	return connectionGroups
}
