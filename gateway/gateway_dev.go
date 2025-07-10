//go:build dev
// +build dev

package gateway

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/config"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/protocol/cmd"
	"github.com/gucooing/BaPs/protocol/mx"
)

func logPlayerMsg(logType int, msg mx.Message) {
	cmdId := cmd.Get().GetCmdIdByProtoObj(msg)
	if _, ok := config.GetBlackCmd()[cmdId.String()]; ok ||
		!config.GetIsLogMsgPlayer() {
		return
	}
	var a string
	switch logType {
	case Client:
		a = "@LogTag(player_msg)@ gateway c--->s cmd id:"
	case Server:
		a = "@LogTag(player_msg)@ gateway s--->c cmd id:"
	case NoRoute:
		a = "@LogTag(player_no_route)@ c --> s no route for msg, cmd id:"
	}
	b, _ := sonic.MarshalIndent(msg, "", "  ")

	logger.Debug("%s%s :%s", a, cmdId.String(), string(b))
}
