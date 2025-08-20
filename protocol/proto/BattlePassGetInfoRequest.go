package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassGetInfoRequest struct {
	*RequestPacket
	BattlePassId int64
}

func (x *BattlePassGetInfoRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassGetInfoRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
