package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopPickupSelectionGachaGetRequest struct {
	*RequestPacket
	ShopRecruitId int64
}

func (x *ShopPickupSelectionGachaGetRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopPickupSelectionGachaGetRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
