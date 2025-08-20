package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BillingPurchaseFreeProductRequest struct {
	*RequestPacket
	ShopCashId int64
}

func (x *BillingPurchaseFreeProductRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BillingPurchaseFreeProductRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
