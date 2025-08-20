package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGSelectRewardCardResponse struct {
	*ResponsePacket
	StageDB           *MiniGameCCGStagePlayDB
	SaveDB            *MiniGameCCGSaveDB
	ReceivedRewardIds []int64
}

func (x *MiniGameCCGSelectRewardCardResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGSelectRewardCardResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
