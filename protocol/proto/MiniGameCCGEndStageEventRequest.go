package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGEndStageEventRequest struct {
	*RequestPacket
	EventContentId int64
}

func (x *MiniGameCCGEndStageEventRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGEndStageEventRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
