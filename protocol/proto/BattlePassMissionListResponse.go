package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassMissionListResponse struct {
	*ResponsePacket
	MissionHistoryUniqueIds []int64
	ProgressDBs             []*MissionProgressDB
}

func (x *BattlePassMissionListResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassMissionListResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
