package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ReceiveAccountLevelRewardRequest struct {
	*RequestPacket
}

func (x *ReceiveAccountLevelRewardRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ReceiveAccountLevelRewardRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
