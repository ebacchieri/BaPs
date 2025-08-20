package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassMissionMultipleRewardResponse struct {
	*ResponsePacket
	AddedHistoryDBs []*MissionHistoryDB
	ParcelResultDB  *ParcelResultDB
	BattlePassInfo  *BattlePassInfoDB
}

func (x *BattlePassMissionMultipleRewardResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassMissionMultipleRewardResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
