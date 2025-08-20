package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGSelectCampActionRequest struct {
	*RequestPacket
	EventContentId  int64
	SelectedOption  MiniGameCCGCampOption
	RemoveCardDBIds []int32
}

func (x *MiniGameCCGSelectCampActionRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGSelectCampActionRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
