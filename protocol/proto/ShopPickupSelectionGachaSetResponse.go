package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopPickupSelectionGachaSetResponse struct {
	*ResponsePacket
}

func (x *ShopPickupSelectionGachaSetResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopPickupSelectionGachaSetResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
