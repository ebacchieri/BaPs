package command

import (
	"fmt"
	"time"

	"github.com/gucooing/BaPs/config"
	"github.com/gucooing/BaPs/protocol/mx"

	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/cdq"
)

const (
	setPlayerUnknown = -1
	setTypeUnknown   = -2
)

var (
	setAccountLevel = "AccountLevel"
	setToast        = "Toast"
	setGetConfig    = "GetConfig"
)

func (c *Command) ApplicationCommandSet() {
	set := &cdq.Command{
		Name:        "set",
		AliasList:   make([]string, 0),
		Description: "直接设置一个值",
		Permissions: cdq.User,
		Options: []*cdq.CommandOption{
			{
				Name:        "uid",
				Description: "玩家游戏id",
				Required:    true,
				Alias:       "u",
			},
			{
				Name:        "type",
				Description: "设置类型",
				Required:    true,
				Alias:       "t",
				ExpectedS:   []string{setAccountLevel, setToast, setGetConfig},
			},
			{
				Name:        "sub1",
				Description: "设置参数1",
				Required:    true,
				Alias:       "s1",
			},
		},
		Handlers: cdq.AddHandlers(syncGateWay, c.set),
	}

	c.C.ApplicationCommand(set)
}

func (c *Command) set(ctx *cdq.Context) {
	// 玩家验证
	uid := ctx.GetFlags().Int64("uid")
	s := enter.GetSessionByAccountServerId(uid)
	if s == nil {
		ctx.Return(setPlayerUnknown, fmt.Sprintf("玩家不在线或未注册 UID:%v", uid))
		return
	}
	switch ctx.GetFlags().String("type") {
	case setAccountLevel:
		game.SetAccountLevel(s, ctx.GetFlags().Int32("sub1"))
		game.AddToast(s, &enter.Toast{
			Text:      "已设置账号等级,请重新登录以刷新",
			BeginDate: mx.Now(),
			EndDate:   mx.Now().Add(30 * time.Second),
		})
		ctx.Return(cdq.ApiCodeOk, fmt.Sprintf("已设置玩家等级:%v级", game.GetAccountLevel(s)))
	case setToast:
		game.AddToast(s, &enter.Toast{
			Text:      ctx.GetFlags().String("sub1"),
			BeginDate: mx.Now(),
			EndDate:   mx.Now().Add(30 * time.Second),
		})
		ctx.Return(cdq.ApiCodeOk, "已设置玩家通知")
	case setGetConfig:
		ctx.Return(cdq.ApiCodeOk, config.GetConfig().String())
	default:
		ctx.Return(setTypeUnknown, "未知的Set Type")
	}
}
