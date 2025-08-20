package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassReceiveRewardRequest struct {
	*RequestPacket
	BattlePassId int64
}

func (x *BattlePassReceiveRewardRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassReceiveRewardRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
