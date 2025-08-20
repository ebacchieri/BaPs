package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassCheckRequest struct {
	*RequestPacket
	BattlePassId int64
}

func (x *BattlePassCheckRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassCheckRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
