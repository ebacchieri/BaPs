package command

import (
	"database/sql"
	"encoding/json"
	"fmt"
	dbstruct "github.com/gucooing/BaPs/db/struct"
	"time"

	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/cdq"
)

const (
	gmDelAllMailErr       = -1
	gmDelMailErr          = -2
	gmGenParcelInfoErr    = -3
	gmSendMailErr         = -4
	gmPlayerUnknown       = -5
	gmDelAllPlayerMailErr = -6
	gmDelPlayerMailErr    = -7
	gmSendPlayerMailErr   = -8
)

func (c *Command) ApplicationCommandGameMail() {
	mail := &cdq.Command{
		Name:        "gameMail",
		AliasList:   []string{"gm"},
		Description: "发送一封游戏邮件,对象可以是全局或单个玩家",
		Permissions: cdq.Admin,
		Options: []*cdq.CommandOption{
			{
				Name:        "recipient",
				Description: "玩家游戏id或all全局对象",
				Required:    true,
				Alias:       "r",
			},
			{
				Name:        "sender",
				Description: "发件人",
				Required:    false,
				Default:     "gucooing",
				Alias:       "s",
			},
			{
				Name:        "comment",
				Description: "邮件文本内容",
				Required:    false,
				Default:     "请查收邮件",
				Alias:       "c",
			},
			{
				Name:        "sendDate",
				Description: "领取开始时间",
				Required:    false,
				Default:     "946656000",
				Alias:       "sd",
			},
			{
				Name:        "expireDate",
				Description: "邮件有效截至时间",
				Required:    false,
				Default:     "4070880000",
				Alias:       "ed",
			},
			{
				Name:        "parcelInfoList",
				Description: "附件 json格式",
				Required:    false,
				Alias:       "ps",
			},
			{
				Name:        "del",
				Description: "删除邮件 可传入邮件id或all",
				Required:    false,
				Alias:       "d",
			},
		},
		Handlers: cdq.AddHandlers(syncGateWay, c.gameMail),
	}

	c.C.ApplicationCommand(mail)
}

func (c *Command) gameMail(ctx *cdq.Context) {
	switch ctx.GetFlags().String("recipient") {
	case "all", "All": // 全局邮件操作
		gmYostarMail(ctx)
	default:
		gmPlayerMail(ctx)
	}
}

func gmYostarMail(ctx *cdq.Context) {
	if strDel := ctx.GetFlags().String("del"); strDel == "all" || strDel == "All" {
		if enter.DelAllYostarMail() != nil {
			ctx.Return(gmDelAllMailErr, "删除所有全局邮件失败")
		} else {
			ctx.Return(cdq.ApiCodeOk, "删除所有全局邮件成功")
		}
	} else if delId := ctx.GetFlags().Int64("del"); delId != 0 {
		if enter.DelYostarMail(delId) != nil {
			ctx.Return(gmDelMailErr, "删除全局邮件失败")
		} else {
			ctx.Return(cdq.ApiCodeOk, "删除全局邮件成功")
		}
	} else {
		parcelInfoList, err := genParcelInfo[dbstruct.ParcelInfo](ctx.GetFlags().String("parcelInfoList"))
		if err != nil {
			ctx.Return(gmGenParcelInfoErr, fmt.Sprintf("解析邮件附件失败:%s", err.Error()))
			return
		}
		sendMail := &dbstruct.YostarMail{
			Sender:         ctx.GetFlags().String("sender"),
			Comment:        ctx.GetFlags().String("comment"),
			SendDate:       sql.NullTime{Time: time.Unix(ctx.GetFlags().Int64("sendDate"), 0), Valid: true},
			ExpireDate:     sql.NullTime{Time: time.Unix(ctx.GetFlags().Int64("expireDate"), 0), Valid: true},
			ParcelInfoList: parcelInfoList,
		}
		if enter.AddYostarMail(sendMail) {
			ctx.Return(cdq.ApiCodeOk, "全局邮件发送成功")
		} else {
			ctx.Return(gmSendMailErr, "全局邮件发送失败")
		}
	}
}

func gmPlayerMail(ctx *cdq.Context) {
	uid := ctx.GetFlags().Int64("recipient")
	s := enter.GetSessionByAccountServerId(uid)
	if s == nil {
		ctx.Return(gmPlayerUnknown, fmt.Sprintf("玩家不在线或未注册 UID:%v", uid))
		return
	}
	if strDel := ctx.GetFlags().String("del"); strDel == "all" || strDel == "All" {
		if game.DelAllMail(s) {
			ctx.Return(cdq.ApiCodeOk, "删除玩家全部邮件成功")
		} else {
			ctx.Return(gmDelAllPlayerMailErr, "删除玩家全部邮件失败,这是一个神奇的bug")
		}
	} else if delId := ctx.GetFlags().Int64("del"); delId != 0 {
		if game.DelMail(s, delId) {
			ctx.Return(cdq.ApiCodeOk, "删除玩家邮件成功")
		} else {
			ctx.Return(gmDelPlayerMailErr, "删除玩家邮件失败,邮件可能不存在")
		}
	} else {
		parcelInfoList, err := genParcelInfo[sro.ParcelInfo](ctx.GetFlags().String("parcelInfoList"))
		if err != nil {
			ctx.Return(gmGenParcelInfoErr, fmt.Sprintf("解析邮件附件失败:%s", err.Error()))
			return
		}
		sendMail := &sro.MailInfo{
			Sender:         ctx.GetFlags().String("sender"),
			Comment:        ctx.GetFlags().String("comment"),
			SendDate:       ctx.GetFlags().Int64("sendDate"),
			ExpireDate:     ctx.GetFlags().Int64("expireDate"),
			ParcelInfoList: parcelInfoList,
		}
		if game.AddMail(s, sendMail) {
			ctx.Return(cdq.ApiCodeOk, "玩家邮件发送成功")
		} else {
			ctx.Return(gmSendPlayerMailErr, "玩家邮件发送失败")
		}
	}
}

func genParcelInfo[T any](str string) ([]*T, error) {
	list := make([]*T, 0)
	if str == "" {
		return list, nil
	}
	err := json.Unmarshal([]byte(str), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
