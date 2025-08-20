package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGEnterStageResponse struct {
	*ResponsePacket
	StageDB *MiniGameCCGStagePlayDB
}

func (x *MiniGameCCGEnterStageResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGEnterStageResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
