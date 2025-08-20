package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassMissionListRequest struct {
	*RequestPacket
	BattlePassId int64
}

func (x *BattlePassMissionListRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassMissionListRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
