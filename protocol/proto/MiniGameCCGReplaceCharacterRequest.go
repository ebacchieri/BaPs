package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGReplaceCharacterRequest struct {
	*RequestPacket
	EventContentId int64
	SlotIndex      int32
	CharacterId    int64
	IsStriker      bool
}

func (x *MiniGameCCGReplaceCharacterRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGReplaceCharacterRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
