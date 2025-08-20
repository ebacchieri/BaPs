package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGSelectRewardCardRequest struct {
	*RequestPacket
	EventContentId int64
	SelectedIndex  int32
	RewardIndex    int32
}

func (x *MiniGameCCGSelectRewardCardRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGSelectRewardCardRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
