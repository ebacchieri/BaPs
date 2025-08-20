package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGCompleteGameResponse struct {
	*ResponsePacket
	OldSaveDB      *MiniGameCCGSaveDB
	ParcelResultDB *ParcelResultDB
	RewardParcels  []*ParcelInfo
}

func (x *MiniGameCCGCompleteGameResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGCompleteGameResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
