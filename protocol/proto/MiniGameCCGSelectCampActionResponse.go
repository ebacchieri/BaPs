package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGSelectCampActionResponse struct {
	*ResponsePacket
	StageDB *MiniGameCCGStagePlayDB
	SaveDB  *MiniGameCCGSaveDB
}

func (x *MiniGameCCGSelectCampActionResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGSelectCampActionResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
