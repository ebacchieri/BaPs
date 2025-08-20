package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameCCGReplaceCharacterResponse struct {
	*ResponsePacket
	SaveDB         *MiniGameCCGSaveDB
	CCGCharacterDB *MiniGameCCGCharacterDB
}

func (x *MiniGameCCGReplaceCharacterResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameCCGReplaceCharacterResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
