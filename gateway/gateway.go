package gateway

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/common/check"
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/config"
	"github.com/gucooing/BaPs/pkg/alg"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/protocol"
	"github.com/gucooing/BaPs/protocol/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

type Gateway struct {
	router *gin.Engine
}

func NewGateWay(router *gin.Engine) *Gateway {
	g := &Gateway{
		router: router,
	}
	enter.MaxCachePlayerTime = alg.MinInt(config.GetGateWay().MaxCachePlayerTime, 30)
	enter.MaxPlayerNum = config.GetGateWay().MaxPlayerNum
	g.initRouter()

	return g
}

func (g *Gateway) initRouter() {
	g.router.POST("/getEnterTicket/gateway", check.GateSync(), g.getEnterTicket) // 这个地方要加个限速器,不然会被dos
	api := g.router.Group("/api")
	{
		api.POST("/gateway", check.GinNoLite(), g.gateWay)
	}
}

func (g *Gateway) send(c *gin.Context, n mx.Message) {
	rsp, err := protocol.MarshalResponse(n)
	if err != nil {
		logger.Debug("marshal err:", err)
		return
	}
	var str string
	if config.GetHttpNet().Encoding {
		byt, _ := sonic.Marshal(rsp)
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		gz.Write(byt)
		gz.Close()
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")
		str = b.String()
	} else {
		str, _ = sonic.MarshalString(rsp)
	}

	c.String(http.StatusOK, str)
}

func (g *Gateway) gateWay(c *gin.Context) {
	if !alg.CheckGateWay(c) {
		return
	}
	bin, err := mx.GetFormMx(c)
	if err != nil {
		logger.Warn("get form mx error:", err)
		return
	}
	packet, base, err := protocol.UnmarshalRequest(bin)
	if err != nil {
		errBestHTTP(c, proto.WebAPIErrorCode_MailBoxFull) // 临时解决方案-避免客户端被弹出
		logger.Debug("unmarshal c--->s err:%s,json:%s", err, string(bin))
		return
	}

	g.registerMessage(c, packet, base)
}

func errBestHTTP(c *gin.Context, errorCode proto.WebAPIErrorCode) {
	msg := &protocol.NetworkProtocolResponse{
		Packet:   fmt.Sprintf("{\"Protocol\":-1,\"ErrorCode\":%d}", errorCode),
		Protocol: "Error",
	}
	c.JSON(200, msg)
}
