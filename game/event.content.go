package game

import (
	"time"

	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/protocol/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func NewEventContentBin() *sro.EventContentBin {
	info := &sro.EventContentBin{}
	return info
}

func GetEventContentBin(s *enter.Session) *sro.EventContentBin {
	info := GetPlayerBin(s)
	if info.EventContentBin == nil {
		info.EventContentBin = NewEventContentBin()
	}
	return info.EventContentBin
}

func GetEventContentInfoList(s *enter.Session) map[int64]*sro.EventContentInfo {
	info := GetEventContentBin(s)
	if info.EventContentInfoList == nil {
		info.EventContentInfoList = make(map[int64]*sro.EventContentInfo)
	}
	return info.EventContentInfoList
}

func GetEventContentInfo(s *enter.Session, eventContentId int64) *sro.EventContentInfo {
	list := GetEventContentInfoList(s)
	if list[eventContentId] == nil {
		list[eventContentId] = &sro.EventContentInfo{
			EventContentId:   eventContentId,
			BoxGachaShopInfo: nil,
			StageHistoryList: make(map[int64]*sro.EventContentStageHistory),
		}
	}
	return list[eventContentId]
}

func UpEventContentInfo(s *enter.Session, eventContentId, stageUniqueId int64) *sro.EventContentStageHistory {
	info := GetEventContentInfo(s, eventContentId)

	if bin, ok := info.StageHistoryList[stageUniqueId]; ok {
		return bin
	}
	bin := &sro.EventContentStageHistory{
		StageUniqueId: stageUniqueId,
		RewardTime:    0,
		LastPlayer:    time.Now().Unix(),
	}
	info.StageHistoryList[stageUniqueId] = bin
	return bin
}

func GetStageHistoryDBs(s *enter.Session, eventContentId int64) (list []*proto.CampaignStageHistoryDB) {
	info := GetEventContentInfo(s, eventContentId)
	list = make([]*proto.CampaignStageHistoryDB, 0)
	for _, bin := range info.StageHistoryList {
		list = append(list, GetCampaignStageHistoryDB(bin))
	}

	return
}

func GetCampaignStageHistoryDB(bin *sro.EventContentStageHistory) *proto.CampaignStageHistoryDB {
	info := &proto.CampaignStageHistoryDB{
		StageUniqueId:           bin.GetStageUniqueId(),
		TodayPlayCount:          1,
		ClearTurnRecord:         1,
		LastPlay:                mx.Unix(bin.GetLastPlayer(), 0),
		FirstClearRewardReceive: mx.Unix(bin.GetRewardTime(), 0),

		StoryUniqueId:                   0,
		ChapterUniqueId:                 0,
		TacticClearCountWithRankSRecord: 0,
		BestStarRecord:                  0,
		Star1Flag:                       false,
		Star2Flag:                       false,
		Star3Flag:                       false,
		TodayPurchasePlayCountHardStage: 0,
		StarRewardReceive:               mx.MxTime{},
		IsClearedEver:                   false,
		TodayPlayCountForUI:             0,
	}
	return info
}
