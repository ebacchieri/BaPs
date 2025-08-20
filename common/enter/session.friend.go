package enter

import (
	"errors"

	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/db"
	dbstruct "github.com/gucooing/BaPs/db/struct"
)

type AccountFriend struct {
	Uid                     int64          `json:"uid"`
	AutoAcceptFriendRequest bool           `json:"auto_accept_friend_request"` // 是否自动同意好友申请
	FriendList              map[int64]bool `json:"friend_list"`                // 好友列表
	ReceivedList            map[int64]bool `json:"received_list"`              // 被申请好友列表
	SendReceivedList        map[int64]bool `json:"send_received_list"`         // 发送的好友申请列表
	BlockedList             map[int64]bool `json:"blocked_list"`               // 黑名单列表
}

// NewAccountFriend 拉取好友关系
func newAccountFriend(uid int64) *AccountFriend {
	af, err := dbGetAccountFriend(uid)
	if err != nil {
		return nil
	}
	return af
}

// DbGetAccountFriend 从db拉取数据
func dbGetAccountFriend(uid int64) (*AccountFriend, error) {
	af := new(AccountFriend)
	bin := db.GetDBGame().GetYostarFriendByAccountServerId(uid)
	if bin == nil {
		return nil, errors.New("sql err")
	}
	sonic.UnmarshalString(bin.FriendInfo, af)
	af.Uid = uid
	return af, nil
}

// GetYostarFriend 预处理db数据
func (x *AccountFriend) GetYostarFriend() *dbstruct.YostarFriend {
	if x == nil {
		return nil
	}
	bin := &dbstruct.YostarFriend{
		AccountServerId: x.Uid,
	}
	friendInfo, err := sonic.MarshalString(x)
	if err != nil {
		return nil
	}
	bin.FriendInfo = friendInfo

	return bin
}

func (x *AccountFriend) GetFriendList() map[int64]bool {
	if x == nil {
		return nil
	}
	friendList := make(map[int64]bool)
	for k, v := range x.FriendList {
		friendList[k] = v
	}
	return friendList
}
