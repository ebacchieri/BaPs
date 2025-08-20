package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGEndStageDualRequest struct {
	*RequestPacket
	EventContentId int64
	Summary        *MiniGameCCGSummary
}

func (x *MiniGameCCGEndStageDualRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGEndStageDualRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
