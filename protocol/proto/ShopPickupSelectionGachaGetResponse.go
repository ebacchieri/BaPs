package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopPickupSelectionGachaGetResponse struct {
	*ResponsePacket
	PickupCharacterSelection map[int64]int64
}

func (x *ShopPickupSelectionGachaGetResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopPickupSelectionGachaGetResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
