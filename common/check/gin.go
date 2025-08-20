package check

import (
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gucooing/BaPs/config"
	"github.com/gucooing/BaPs/pkg/alg"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/pkg/logger"
)

var TPS int64
var RT time.Duration
var OLDTPS int64
var OLDRT time.Duration
var SessionNum int64

func GinNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		tps := atomic.LoadInt64(&TPS)
		OLDTPS = tps / 60
		rt := atomic.LoadInt64((*int64)(&RT))
		if tps == 0 || rt == 0 {
			OLDRT = 0
			continue
		} else {
			OLDRT = time.Duration(rt / tps)
		}
		logger.Info("SessionNum: %v", SessionNum)
		logger.Info("TPS: %v", OLDTPS)
		logger.Info("RT: %s", OLDRT)
		atomic.StoreInt64(&TPS, 0)
		atomic.StoreInt64((*int64)(&RT), 0)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

var GateWaySync *sync.Mutex

func GateSync() gin.HandlerFunc {
	return func(c *gin.Context) {
		GateWaySync.Lock()
		c.Next()
		GateWaySync.Unlock()
	}
}

func GinNoPublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.GetIsLite() {
			if !alg.IsPrivateIP(net.ParseIP(c.ClientIP())) {
				c.Abort()
			}
			//host := c.Request.Host
			//srts := strings.Split(host, ":")
			//if len(srts) < 1 {
			//	c.Abort()
			//}
			//if !alg.IsPrivateIP(net.ParseIP(srts[0])) {
			//	c.Abort()
			//}
		}
	}
}

func GinNoLite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 暂不启用
		//if config.GetIsLite() {
		//	c.Abort()
		//}
	}
}
