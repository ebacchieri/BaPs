package sdk

import (
	"encoding/json"
	"github.com/gucooing/BaPs/gdconf"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/config"
)

func (s *SDK) GetServerInfo(c *gin.Context) *gdconf.ServerInfo {
	conf := gdconf.GetServerInfo()
	data := &gdconf.ServerInfo{
		ConnectionGroups: make([]*gdconf.ConnectionGroup, 0),
	}
	switch config.GetOtherAddr().ServerInfoUrl {
	case "local":
		return conf
	}
	getServerInfo := func() {
		conf.LoadSync.Lock()
		defer conf.LoadSync.Unlock()
		url := c.Request.URL.String()
		resp, err := http.Get(config.GetOtherAddr().GetServerInfoUrl() + url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		bin, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(bin, &data)

		conf.ConnectionGroups = gdconf.GenServerInfo(data.ConnectionGroups)
		conf.UpTime = time.Now()
	}

	conf.LoadSync.RLock()
	defer conf.LoadSync.RUnlock()
	if conf.UpTime.Add(15 * time.Minute).Before(time.Now()) {
		conf.LoadSync.RUnlock()
		getServerInfo()
		conf.LoadSync.RLock()
	}

	return conf
}

func (s *SDK) connectionGroups(c *gin.Context) {
	c.JSON(200, s.GetServerInfo(c))
}
