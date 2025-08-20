package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassCheckResponse struct {
	*ResponsePacket
	HasNotReceiveReward bool
	HasCompleteMission  bool
}

func (x *BattlePassCheckResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassCheckResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
