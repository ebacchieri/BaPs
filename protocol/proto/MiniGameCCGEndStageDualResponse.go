package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGEndStageDualResponse struct {
	*ResponsePacket
	StageDB *MiniGameCCGStagePlayDB
	SaveDB  *MiniGameCCGSaveDB
}

func (x *MiniGameCCGEndStageDualResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGEndStageDualResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
