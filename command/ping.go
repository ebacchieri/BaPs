package command

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/common/check"
	"github.com/gucooing/BaPs/pkg"
	"github.com/gucooing/cdq"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"sync/atomic"
)

const (
	pingMarshalErr = -1
)

type Ping struct {
	PlayerNum     int64   `json:"playerNum"`     // 在线玩家数量
	Tps           int64   `json:"tps"`           // 上一分钟请求量
	Rt            string  `json:"rt"`            // 上一分钟每一个请求平均处理时间
	ServerVersion string  `json:"serverVersion"` // 服务端版本
	ApiVersion    string  `json:"apiVersion"`    // api版本
	CpuOc         float64 `json:"cpuOc"`         // cpu占用
	MemoryOc      string  `json:"memoryOc"`      // 内存占用
	BaPsMemoryOc  string  `json:"baPsMemoryOc"`  // BaPs内存占用
}

func (c *Command) ApplicationCommandPing() {
	ping := &cdq.Command{
		Name:        "ping",
		AliasList:   make([]string, 0),
		Description: "检查服务是否存活",
		Permissions: cdq.Guest,
		Options:     nil,
		Handlers:    cdq.AddHandlers(c.ping),
	}
	c.C.ApplicationCommand(ping)
}

func (c *Command) ping(ctx *cdq.Context) {
	response := Ping{
		PlayerNum:     atomic.LoadInt64(&check.SessionNum),
		Tps:           check.OLDTPS,
		Rt:            check.OLDRT.String(),
		ServerVersion: pkg.ServerVersion,
		ApiVersion:    apiVersion,
		CpuOc:         GetCpuOc(),
		MemoryOc:      MemoryOc(),
		BaPsMemoryOc:  BaPsMemoryOc(),
	}
	str, err := sonic.MarshalString(response)
	if err != nil {
		ctx.Return(pingMarshalErr, "序列化失败")
	}
	ctx.Return(cdq.ApiCodeOk, str)
}

func GetCpuOc() float64 {
	percents, err := cpu.Percent(0, false)
	if err != nil {
		return 0
	}
	return percents[0]
}

func MemoryOc() string {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return "0/0"
	}
	used := float64(memInfo.Used)
	total := float64(memInfo.Total)

	if used/1024/1024 > 1024 {
		return fmt.Sprintf("%.2f/%.2fGB", used/1024/1024/1024, total/1024/1024/1024)
	}
	return fmt.Sprintf("%.2f/%.2fMB", used/1024/1024, total/1024/1024)
}

func BaPsMemoryOc() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	used := float64(m.Alloc)
	if used/1024/1024 > 1024 {
		return fmt.Sprintf("%.2fGB", used/1024/1024/1024)
	}
	return fmt.Sprintf("%.2fMB", used/1024/1024)
}
