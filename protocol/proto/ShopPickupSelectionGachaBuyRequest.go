package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopPickupSelectionGachaBuyRequest struct {
	*ShopBuyGacha2Request
	FreeRecruitId int64
	Cost          *ParcelCost
}

func (x *ShopPickupSelectionGachaBuyRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopPickupSelectionGachaBuyRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ShopBuyGacha2Request = packet.(*ShopBuyGacha2Request)
}
