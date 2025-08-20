package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopPickupSelectionGachaBuyResponse struct {
	*ShopBuyGacha2Response
	FreeRecruitHistoryDB *ShopFreeRecruitHistoryDB
}

func (x *ShopPickupSelectionGachaBuyResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopPickupSelectionGachaBuyResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ShopBuyGacha2Response = packet.(*ShopBuyGacha2Response)
}
