package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassMissionMultipleRewardRequest struct {
	*RequestPacket
	MissionCategory MissionCategory
	BattlePassId    int64
}

func (x *BattlePassMissionMultipleRewardRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassMissionMultipleRewardRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
