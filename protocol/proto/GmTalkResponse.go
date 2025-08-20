package proto

import (
	"github.com/bytedance/sonic"
	"github.com/ebacchieri/BaPs/protocol/mx"
)

type GmTalkResponse struct {
	*ResponsePacket
}

func (x *GmTalkResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *GmTalkResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
