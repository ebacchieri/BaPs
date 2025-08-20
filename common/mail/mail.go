package mail

import (
	"bytes"
	"errors"
	ht "html/template"
	tt "text/template"

	"github.com/gucooing/BaPs/config"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/wneessen/go-mail"
)

var mc *Mail

type Mail struct {
	*mail.Client
	from string
}

func NewMail() {
	conf := config.GetMail()
	if conf == nil || !conf.Enable {
		logger.Info("不启用邮件服务配置")
		return
	}
	logger.Info("启用邮件服务配置,正在尝试连接")

	client, err := mail.NewClient(
		conf.Host,
		mail.WithPort(conf.Port),              // 端口
		mail.WithSMTPAuth(mail.SMTPAuthPlain), // 认证方式
		mail.WithUsername(conf.Username),      // 邮箱账号
		mail.WithPassword(conf.Password),      // 邮箱密码或应用专用密码
		mail.WithTLSPolicy(mail.TLSMandatory), // 强制 TLS
	)
	if err != nil {
		logger.Error("连接邮件服务失败")
		return
	}
	logger.Info("邮件服务连接成功")

	mc = &Mail{
		Client: client,
		from:   conf.Username,
	}
}

func SendTextMail(header, body string, usernames ...string) error {
	if !enableMail() {
		return nil
	}
	var messages []*mail.Msg
	for _, username := range usernames {
		message, err := newMailMessage(username)
		if err != nil {
			return err
		}
		message.Subject(header)
		message.SetBodyString(mail.TypeTextPlain, body)
		messages = append(messages, message)
	}
	return mc.DialAndSend(messages...)
}

func SendTemplateMail(username string, conf *gdconf.MailInfo, data interface{}) error {
	if !enableMail() {
		return nil
	}
	switch conf.Type {
	case "text":
		return sendTextTemplateMail(username, conf, data)
	case "html":
		return sendHtmlTemplateMail(username, conf, data)
	default:
		return errors.New("未知的邮件模板")
	}
}

func sendTextTemplateMail(username string, conf *gdconf.MailInfo, data interface{}) error {
	message, err := newMailMessage(username)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(nil)
	if err := conf.Tpl.(*tt.Template).Execute(buffer, data); err != nil {
		return err
	}
	message.Subject(conf.Header)
	message.SetBodyString(mail.TypeTextPlain, buffer.String())
	return mc.DialAndSend(message)
}

func sendHtmlTemplateMail(username string, conf *gdconf.MailInfo, data interface{}) error {
	message, err := newMailMessage(username)
	if err != nil {
		return err
	}
	message.Subject(conf.Header)
	if err := message.SetBodyHTMLTemplate(conf.Tpl.(*ht.Template), data); err != nil {
		return err
	}
	return mc.DialAndSend(message)
}

func newMailMessage(username string) (*mail.Msg, error) {
	message := mail.NewMsg()
	if err := message.From(mc.from); err != nil {
		return nil, errors.New("邮件服务配置错误")
	}
	if err := message.To(username); err != nil {
		return nil, errors.New("收件人不合法")
	}
	return message, nil
}

func enableMail() bool {
	if mc == nil {
		return false
	}
	return config.GetMail().Enable
}
