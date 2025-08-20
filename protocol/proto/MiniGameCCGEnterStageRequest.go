package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGEnterStageRequest struct {
	*RequestPacket
	EventContentId int64
	NodeId         int64
}

func (x *MiniGameCCGEnterStageRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGEnterStageRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
