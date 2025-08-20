package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassBuyLevelRequest struct {
	*RequestPacket
	BattlePassId            int64
	BattlePassBuyLevelCount int32
}

func (x *BattlePassBuyLevelRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassBuyLevelRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
