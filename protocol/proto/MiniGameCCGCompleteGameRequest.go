package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGCompleteGameRequest struct {
	*RequestPacket
	EventContentId int64
}

func (x *MiniGameCCGCompleteGameRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGCompleteGameRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
