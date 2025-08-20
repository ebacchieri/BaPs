package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ReceiveAccountLevelRewardResponse struct {
	*ResponsePacket
	ReceivedAccountLevelRewardIds []int64
	ParcelResultDB *ParcelResultDB
}

func (x *ReceiveAccountLevelRewardResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ReceiveAccountLevelRewardResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
