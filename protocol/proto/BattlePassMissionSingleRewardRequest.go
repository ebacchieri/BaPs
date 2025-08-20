package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassMissionSingleRewardRequest struct {
	*RequestPacket
	BattlePassId    int64
	MissionUniqueId int64
}

func (x *BattlePassMissionSingleRewardRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassMissionSingleRewardRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
