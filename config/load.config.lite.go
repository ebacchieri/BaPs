//go:build lite

package config

import (
	"fmt"
	"net"

	"github.com/gucooing/BaPs/pkg/alg"
)

func (c *Config) check() {
	if c == nil || mx.CheckDev() {
		return
	}
	// 检查是否开启lite
	c.IsLite = true
	// 检查是否关闭api key
	c.GucooingApiKey = ""
	// 检查是否开启自动注册
	c.AutoRegistration = true
	// 强制玩家数为1
	c.GateWay.MaxPlayerNum = 1
	// 检查是否为外网配置
	if !alg.IsPrivateIP(net.ParseIP(c.HttpNet.InnerIp)) {
		panic("Only private addresses are allowed")
	}
	// 强制http和固定地址
	c.HttpNet.Tls = false
	c.HttpNet.OuterAddr = fmt.Sprintf("http://%s:%s", c.HttpNet.InnerIp, c.HttpNet.InnerPort)
}
