package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGBuyPerkResponse struct {
	*ResponsePacket
	Perks                     []int64
	ParcelResultDB            *ParcelResultDB
	EventContentCollectionDBs []*EventContentCollectionDB
}

func (x *MiniGameCCGBuyPerkResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGBuyPerkResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
