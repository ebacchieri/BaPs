package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassBuyLevelResponse struct {
	*ResponsePacket
	BattlePassInfo    *BattlePassInfoDB
	AccountCurrencyDB *AccountCurrencyDB
}

func (x *BattlePassBuyLevelResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *BattlePassBuyLevelResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
