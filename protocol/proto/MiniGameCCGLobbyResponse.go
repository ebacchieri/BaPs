package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGLobbyResponse struct {
	*ResponsePacket
	CCGSaveDB   *MiniGameCCGSaveDB
	Perks       []int64
	RewardPoint int32
	CanSweep    bool
}

func (x *MiniGameCCGLobbyResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGLobbyResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
