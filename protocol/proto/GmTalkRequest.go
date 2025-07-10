package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type GmTalkRequest struct {
	*RequestPacket

	TestType int
}

func (x *GmTalkRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *GmTalkRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
