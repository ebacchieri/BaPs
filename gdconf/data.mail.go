package gdconf

import (
	"encoding/json"
	ht "html/template"
	"os"
	tt "text/template"

	"github.com/gucooing/BaPs/pkg/logger"
)

type MailInfo struct {
	Type   string `json:"type"`
	Header string `json:"header"`
	Body   string `json:"body"`
	Tpl    any    `json:"-"`
}

func (g *GameConfig) loadMailInfo() {
	g.GetGPP().MailInfoMap = make(map[string]*MailInfo)
	name := "Mail.json"
	file, err := os.ReadFile(g.dataPath + name)
	if err != nil {
		logger.Error("文件:%s 读取失败,err:%s", name, err)
		return
	}
	if err := json.Unmarshal(file, &g.GetGPP().MailInfoMap); err != nil {
		logger.Error("文件:%s 解析失败,err:%s", name, err)
		return
	}
	for k, v := range g.GetGPP().MailInfoMap {
		switch v.Type {
		case "text":
			tmpl, err := tt.New(k).Parse(v.Body)
			if err != nil {
				logger.Error("错误的模板配置:%s ,err:%s", k, err)
				continue
			}
			v.Tpl = tmpl
		case "html":
			tmpl, err := ht.ParseFiles(v.Body)
			if err != nil {
				logger.Error("错误的模板配置:%s ,err:%s", k, err)
				continue
			}
			v.Tpl = tmpl
		default:
			logger.Error("未知的邮件模板:%s", v.Type)
			continue
		}
	}

	logger.Info("邮件配置读取成功文件:%s 读取成功,解析数量:%v", name, len(g.GetGPP().MailInfoMap))
}

func GetMailInfo(name string) *MailInfo {
	return GC.GPP.MailInfoMap[name]
}
