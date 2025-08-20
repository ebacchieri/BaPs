package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassMissionSingleRewardResponse struct {
	*ResponsePacket
	AddedHistoryDB *MissionHistoryDB
	ParcelResultDB *ParcelResultDB
	BattlePassInfo *BattlePassInfoDB
}

func (x *BattlePassMissionSingleRewardResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassMissionSingleRewardResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
