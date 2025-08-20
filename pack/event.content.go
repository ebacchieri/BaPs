package pack

import (
	"time"

	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/protocol/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func EventContentAdventureList(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentAdventureListRequest)
	rsp := response.(*proto.EventContentAdventureListResponse)

	rsp.AlreadyReceiveRewardId = make([]int64, 0)
	rsp.EventContentBonusRewardDBs = make([]*proto.EventContentBonusRewardDB, 0)
	rsp.StrategyObjecthistoryDBs = make([]*proto.StrategyObjectHistoryDB, 0)
	rsp.StagePoint = 0

	rsp.StageHistoryDBs = game.GetStageHistoryDBs(s, req.EventContentId)
}

func EventContentBoxGachaShopList(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentBoxGachaShopListRequest)
	rsp := response.(*proto.EventContentBoxGachaShopListResponse)

	rsp.BoxGachaDB = &proto.EventContentBoxGachaDB{
		EventContentId: req.EventContentId,
		Seed:           0,
		Round:          1, // 回合数 默认1
		PurchaseCount:  0, // 购买数量
	}
	rsp.BoxGachaGroupIdByCount = make(map[int64]int64)
}

func EventContentScenarioGroupHistoryUpdate(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentScenarioGroupHistoryUpdateRequest)
	rsp := response.(*proto.EventContentScenarioGroupHistoryUpdateResponse)

	defer func() {
		rsp.ScenarioGroupHistoryDBs = game.GetScenarioGroupHistoryDBs(s)
	}()

	game.FinishScenarioGroupHistoryInfo(s, req.ScenarioGroupUniqueId, req.ScenarioType, req.EventContentId)
	// 奖励

}

func EventContentEnterStoryStage(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentEnterStoryStageRequest)
	rsp := response.(*proto.EventContentEnterStoryStageResponse)

	rsp.ParcelResultDB = game.ParcelResultDB(s, nil)
	rsp.SaveDataDB = &proto.EventContentStoryStageSaveDB{
		CampaignSubStageSaveDB: &proto.CampaignSubStageSaveDB{
			ContentSaveDB: &proto.ContentSaveDB{
				ContentType:              0,
				AccountServerId:          s.AccountServerId,
				CreateTime:               mx.Now(),
				StageUniqueId:            req.StageUniqueId,
				StageEntranceFee:         make([]*proto.ParcelInfo, 0),
				EnemyKillCountByUniqueId: make(map[int64]int64),

				LastEnterStageEchelonNumber: 0,
				TacticClearTimeMscSum:       0,
				AccountLevelWhenCreateDB:    0,
				BIEchelon:                   "",
				BIEchelon1:                  "",
				BIEchelon2:                  "",
				BIEchelon3:                  "",
				BIEchelon4:                  "",
			},
		},
		ContentType: proto.ContentType_EventContentStoryStage,
	}
}

func EventContentStoryStageResult(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentStoryStageResultRequest)
	rsp := response.(*proto.EventContentStoryStageResultResponse)

	rsp.FirstClearReward = make([]*proto.ParcelInfo, 0)
	parcelResultList := make([]*game.ParcelResult, 0)
	defer func() {
		rsp.ParcelResultDB = game.ParcelResultDB(s, parcelResultList)
		rsp.FirstClearReward = rsp.ParcelResultDB.ParcelForMission
	}()

	conf := gdconf.GetEventContentStageExcelTable(req.StageUniqueId)
	if conf == nil {
		return
	}
	// up
	bin := game.UpEventContentInfo(s, req.EventContentId, req.StageUniqueId)
	if bin.GetRewardTime() == 0 {
		for _, re := range gdconf.GetEventContentStageRewardExcelList(conf.EventContentStageRewardId) {
			switch re.RewardTag {
			case "FirstClear":
				parcelResultList = append(parcelResultList, &game.ParcelResult{
					ParcelType: proto.ParcelType_None.Value(re.RewardParcelType),
					ParcelId:   re.RewardId,
					Amount:     re.RewardAmount,
				})
			}
		}
		bin.RewardTime = time.Now().Unix()
	}
	rsp.CampaignStageHistoryDB = game.GetCampaignStageHistoryDB(bin)
}

func EventContentEnterMainGroundStage(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentEnterMainGroundStageRequest)
	rsp := response.(*proto.EventContentEnterMainGroundStageResponse)

	bin := game.UpEventContentInfo(s, req.EventContentId, req.StageUniqueId)
	rsp.CampaignStageHistoryDB = game.GetCampaignStageHistoryDB(bin)
}

func EventContentMainGroundStageResult(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.EventContentMainGroundStageResultRequest)
	rsp := response.(*proto.EventContentMainGroundStageResultResponse)

	parcelResultList := make([]*game.ParcelResult, 0)
	rsp.FirstClearReward = make([]*proto.ParcelInfo, 0)
	defer func() {
		rsp.ParcelResultDB = game.ParcelResultDB(s, parcelResultList)
		rsp.FirstClearReward = rsp.ParcelResultDB.ParcelForMission
	}()

	conf := gdconf.GetEventContentStageExcelTable(req.Summary.StageId)
	if conf == nil {
		return
	}
	bin := game.UpEventContentInfo(s, req.EventContentId, req.Summary.StageId)
	if bin.RewardTime == 0 && req.Summary.Winner == proto.GroupTag_Group01.String() {
		for _, re := range gdconf.GetEventContentStageRewardExcelList(conf.EventContentStageRewardId) {
			switch re.RewardTag {
			case "FirstClear":
				parcelResultList = append(parcelResultList, &game.ParcelResult{
					ParcelType: proto.ParcelType_None.Value(re.RewardParcelType),
					ParcelId:   re.RewardId,
					Amount:     re.RewardAmount,
				})
			}
		}
		bin.RewardTime = time.Now().Unix()
	}
	rsp.CampaignStageHistoryDB = game.GetCampaignStageHistoryDB(bin)
}

func EventContentCollectionList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.EventContentCollectionListResponse)

	rsp.EventContentUnlockCGDBs = make([]*proto.EventContentCollectionDB, 0)
}
