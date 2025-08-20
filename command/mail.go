package command

import (
	"fmt"
	"strings"

	"github.com/gucooing/BaPs/common/mail"
	"github.com/gucooing/cdq"
)

const (
	mailSendErr = -1
)

func (c *Command) ApplicationCommandMail() {
	apiMail := &cdq.Command{
		Name:        "mail",
		AliasList:   []string{"em"},
		Description: "发送一封邮件,对象可以是多个用户",
		Permissions: cdq.Admin,
		Options: []*cdq.CommandOption{
			{
				Name:        "header",
				Description: "邮件标题",
				Required:    true,
				Alias:       "h",
			},
			{
				Name:        "body",
				Description: "邮件内容,必须是string格式",
				Required:    true,
				Alias:       "b",
			},
			{
				Name:        "usernames",
				Description: "送达邮箱,可以是多个邮箱,用;隔开",
				Required:    true,
				Alias:       "u",
			},
		},
		Handlers: cdq.AddHandlers(c.mail),
	}

	c.C.ApplicationCommand(apiMail)
}

func (c *Command) mail(ctx *cdq.Context) {
	usernames := strings.Split(ctx.GetFlags().String("usernames"), ";")
	err := mail.SendTextMail(
		ctx.GetFlags().String("header"),
		ctx.GetFlags().String("body"),
		usernames...)
	if err != nil {
		ctx.Return(mailSendErr, fmt.Sprintf("邮件发送失败err:%s", err))
	} else {
		ctx.Return(cdq.ApiCodeOk, "邮件发送成功")
	}
}
