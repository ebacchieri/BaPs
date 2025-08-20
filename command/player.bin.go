package command

import (
	"encoding/base64"
	"fmt"

	"github.com/gucooing/BaPs/game"

	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/cdq"
)

var (
	efJson   = "json"
	efBase64 = "base64"
)

const (
	getPlayerUnknown               = -1
	getPlayerMarshalErr            = -2
	getPlayerEncodingFormatUnknown = -3
)

func (c *Command) ApplicationCommandGetPlayer() {
	getPlayer := &cdq.Command{
		Name:        "getPlayer",
		AliasList:   []string{"gp"},
		Description: "获取指定玩家的db数据",
		Permissions: cdq.Admin,
		Options: []*cdq.CommandOption{
			{
				Name:        "uid",
				Description: "玩家游戏id",
				Required:    true,
				Alias:       "u",
			},
			{
				Name:        "encodingFormat",
				Description: "编码格式",
				Required:    false,
				Alias:       "ef",
				Default:     efBase64,
				ExpectedS:   []string{efJson, efBase64},
			},
			{
				Name:        "basis",
				Description: "获取玩家基础数据",
				Required:    false,
				Alias:       "b",
			},
		},
		Handlers: cdq.AddHandlers(syncGateWay, c.getPlayer),
	}

	c.C.ApplicationCommand(getPlayer)
}

func (c *Command) getPlayer(ctx *cdq.Context) {
	uid := ctx.GetFlags().Int64("uid")
	session := enter.GetSessionByUid(uid)
	if session == nil {
		ctx.Return(getPlayerUnknown, fmt.Sprintf("玩家未注册 UID:%v", uid))
		return
	}
	var info any
	if ctx.GetFlags().Bool("basis") {
		info = game.GetPlayerBin(session).GetBaseBin()
	} else {
		info = game.GetPlayerBin(session)
	}

	jsonInfo, err := sonic.Marshal(info)
	if err != nil {
		ctx.Return(getPlayerMarshalErr, fmt.Sprintf("玩家数据序列化失败 UID:%v", uid))
		return
	}

	switch ctx.GetFlags().String("encodingFormat") {
	case efJson:
		ctx.Return(cdq.ApiCodeOk, string(jsonInfo))
	case efBase64:
		ctx.Return(cdq.ApiCodeOk, base64.StdEncoding.EncodeToString(jsonInfo))
	default:
		ctx.Return(getPlayerEncodingFormatUnknown, fmt.Sprintf("未知的编码方式"))
	}
}
