package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGCreateGameRequest struct {
	*RequestPacket
	EventContentId   int64
	ForceDiscardSave bool
	DisablePerk      bool
}

func (x *MiniGameCCGCreateGameRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGCreateGameRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
