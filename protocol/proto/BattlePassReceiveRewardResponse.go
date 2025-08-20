package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassReceiveRewardResponse struct {
	*ResponsePacket
	BattlePassInfo *BattlePassInfoDB
	ParcelResult   *ParcelResultDB
}

func (x *BattlePassReceiveRewardResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassReceiveRewardResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
