package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassGetInfoResponse struct {
	*ResponsePacket
	BattlePassInfo *BattlePassInfoDB
}

func (x *BattlePassGetInfoResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassGetInfoResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
