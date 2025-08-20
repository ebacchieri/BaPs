package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type BattlePassInfoDB struct {
	BattlePassId                      int64
	PassLevel                         int32
	PassExp                           int64
	PurchaseGroupId                   int64
	ReceiveRewardLevel                int32
	ReceivePurchaseRewardLevel        int32
	WeeklyPassExp                     int64
	LastWeeklyPassExpLimitRefreshDate mx.MxTime
}

func (x *BattlePassInfoDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
