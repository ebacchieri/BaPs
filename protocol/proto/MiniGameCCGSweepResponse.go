package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGSweepResponse struct {
	*ResponsePacket
	Rewards        []*[]*ParcelInfo
	ParcelResultDB *ParcelResultDB
}

func (x *MiniGameCCGSweepResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGSweepResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
