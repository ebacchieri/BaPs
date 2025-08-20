package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopPickupSelectionGachaSetRequest struct {
	*RequestPacket
	ShopRecruitID            int64
	PickupCharacterSelection map[int64]int64
}

func (x *ShopPickupSelectionGachaSetRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopPickupSelectionGachaSetRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
