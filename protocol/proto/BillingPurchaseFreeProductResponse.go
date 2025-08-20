package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BillingPurchaseFreeProductResponse struct {
	*ResponsePacket
	ParcelResult    *ParcelResultDB
	MailDB          *MailDB
	PurchaseProduct *PurchaseCountDB
}

func (x *BillingPurchaseFreeProductResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BillingPurchaseFreeProductResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
