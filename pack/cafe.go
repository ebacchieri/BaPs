package pack

import (
	"./protocol/mx"
	"time"

	"./common/enter"
	sro "./common/server_only"
	"./game"
	"./gdconf"
	"./protocol/proto"
)

func CafeGetInfo(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.CafeGetInfoResponse)

	rsp.CafeDBs = game.GetPbCafeDBs(s)
	rsp.FurnitureDBs = game.GetFurnitureDBs(s) // 已获得家具数据
}

func CafeAck(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeAckRequest)
	rsp := response.(*proto.CafeAckResponse)

	cafeInfo := game.GetCafeInfo(s, req.CafeDBId)
	if cafeInfo == nil {
		s.Error = 0
		return
	}

	rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)

	cafeInfo.IsNew = false
}

func CafeOpen(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeOpenRequest)
	rsp := response.(*proto.CafeOpenResponse)

	bin := game.NewCafeBin(s, game.GetCafeBin(s), req.CafeId)
	rsp.OpenedCafeDB = game.GetCafeDB(s, bin.ServerId)
	rsp.FurnitureDBs = game.GetFurnitureDBs(s)
}

func CafeRemove(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeRemoveFurnitureRequest)
	rsp := response.(*proto.CafeRemoveFurnitureResponse)

	defer func() {
		rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)
		rsp.FurnitureDBs = game.GetFurnitureDBs(s)
	}()

	for _, serverId := range req.FurnitureServerIds {
		game.RemoveFurniture(s, serverId, req.CafeDBId)
	}
	game.UpCafeComfortValue(s, req.CafeDBId)
}

func CafeRemoveAll(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeRemoveAllFurnitureRequest)
	rsp := response.(*proto.CafeRemoveAllFurnitureResponse)

	defer func() {
		rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)
		rsp.FurnitureDBs = game.GetFurnitureDBs(s)
	}()

	cafeInfo := game.GetCafeInfo(s, req.CafeDBId)
	if cafeInfo == nil {
		return
	}
	for serverId, ok := range cafeInfo.FurnitureList {
		if ok {
			game.RemoveFurniture(s, serverId, req.CafeDBId)
		}
	}
	game.UpCafeComfortValue(s, req.CafeDBId)
}

func CafeDeploy(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeDeployFurnitureRequest)
	rsp := response.(*proto.CafeDeployFurnitureResponse)

	rsp.ChangedFurnitureDBs = make([]*proto.FurnitureDB, 0)
	defer func() {
		rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)
	}()

	sid := game.DeployRelocateFurniture(s, req.FurnitureDB, req.CafeDBId)
	rsp.ChangedFurnitureDBs = append(rsp.ChangedFurnitureDBs,
		game.GetFurnitureDB(s, sid))
}

func CafeRelocate(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeRelocateFurnitureRequest)
	rsp := response.(*proto.CafeRelocateFurnitureResponse)

	defer func() {
		rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)
	}()

	sid := game.DeployRelocateFurniture(s, req.FurnitureDB, req.CafeDBId)
	rsp.RelocatedFurnitureDB = game.GetFurnitureDB(s, sid)
}

func CafeInteract(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeInteractWithCharacterRequest)
	rsp := response.(*proto.CafeInteractWithCharacterResponse)

	parcelResultList := make([]*game.ParcelResult, 0)
	defer func() {
		rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)
		rsp.ParcelResultDB = game.ParcelResultDB(s, parcelResultList)
		rsp.CharacterDB = game.GetCharacterDB(s, req.CharacterId)
	}()

	cafeInfo := game.GetCafeInfo(s, req.CafeDBId)
	if cafeInfo == nil {
		return
	}
	if visitCharacterInfo, ok := cafeInfo.VisitCharacterList[req.CharacterId]; ok {
		visitCharacterInfo.LastInteractTime = time.Now().Unix()
		// 添加学生好感度
		parcelResultList = append(parcelResultList, &game.ParcelResult{
			ParcelType: proto.ParcelType_FavorExp,
			ParcelId:   visitCharacterInfo.CharacterId,
			Amount:     15,
		})
	}
}

func CafeSummonCharacter(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeSummonCharacterRequest)
	rsp := response.(*proto.CafeSummonCharacterResponse)

	defer func() {
		rsp.CafeDBs = game.GetPbCafeDBs(s)
	}()
	characterInfo := s.GetCharacterByKeyId(req.CharacterServerId)
	cafeInfo := game.GetCafeInfo(s, req.CafeDBId)
	if cafeInfo == nil || characterInfo == nil {
		return
	}

	if cafeInfo.VisitCharacterList == nil {
		cafeInfo.VisitCharacterList = make(map[int64]*sro.VisitCharacterInfo)
	}
	cafeInfo.SummonUpdate = time.Now().Unix()
	cafeInfo.VisitCharacterList[characterInfo.CharacterId] = &sro.VisitCharacterInfo{
		CharacterId: characterInfo.CharacterId,
		IsSummon:    true,
	}
}

func CafeRankUp(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeRankUpRequest)
	rsp := response.(*proto.CafeRankUpResponse)

	defer func() {
		rsp.CafeDB = game.GetCafeDB(s, req.CafeDBId)
	}()

	rsp.ConsumeResultDB = &proto.ConsumeResultDB{
		UsedItemServerIdAndRemainingCounts: make(map[int64]int64),
	}

	cafeInfo := game.GetCafeInfo(s, req.CafeDBId)
	if cafeInfo == nil {
		return
	}
	cafeRankConf := gdconf.GetCafeRankExcelTable(cafeInfo.CafeId, cafeInfo.CafeRank)
	if cafeRankConf == nil {
		return
	}
	recConf := gdconf.GetRecipeIngredientExcelTable(cafeRankConf.RecipeId)
	if recConf == nil {
		return
	}
	// 根据配方计算需要的东西
	parcelResultList := game.GetParcelResultList(recConf.CostParcelType,
		recConf.CostId, recConf.CostAmount, true)
	parcelResultList = append(parcelResultList, game.GetParcelResultList(recConf.IngredientParcelType,
		recConf.IngredientId, recConf.IngredientAmount, true)...)

	// 扣除
	rsp.ParcelResultDB = game.ParcelResultDB(s, parcelResultList)
	// 构造回复
	cafeInfo.CafeRank++
	for itemServerId, _ := range req.ConsumeRequestDB.ConsumeItemServerIdAndCounts {
		// 这里采用从配方里取物品id判断物品唯一id 节省性能 什么逆天物品服务端唯一id设定?
		for _, id := range recConf.IngredientId {
			itemInfo := game.GetItemInfo(s, id)
			if itemInfo != nil &&
				itemInfo.ServerId == itemServerId {
				rsp.ConsumeResultDB.UsedItemServerIdAndRemainingCounts[itemServerId] = int64(itemInfo.StackCount)
				break
			}
		}
	}
}

func CafeReceiveCurrency(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.CafeReceiveCurrencyResponse)

	parcelResultList := make([]*game.ParcelResult, 0)
	for _, bin := range game.GetCafeInfoList(s) {
		game.UpCafeVisitCharacterDB(bin)
		for _, prodBin := range bin.ProductionList {
			parcelResultList = append(parcelResultList, &game.ParcelResult{
				ParcelType: proto.ParcelType(prodBin.ParcelType),
				ParcelId:   prodBin.ParcelId,
				Amount:     prodBin.Amount / 100,
			})
			// prodBin.Amount = 0
		}
		bin.ProductionAppliedTime = time.Now().Unix()
	}

	rsp.CafeDBs = game.GetPbCafeDBs(s)
	rsp.ParcelResultDB = game.ParcelResultDB(s, parcelResultList)
}

func CafeListPreset(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.CafeListPresetResponse)

	rsp.CafePresetDBs = make([]*proto.CafePresetDB, 0)
}

func CafeTravel(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeTravelRequest)
	rsp := response.(*proto.CafeTravelResponse)

	bin := game.GetFriendBin(s)
	if bin == nil {
		return
	}
	if _, ok := bin.FriendList[req.TargetAccountId]; !ok &&
		enter.GetYostarClanByServerId(game.GetServerId(s)).GetClanAccount(req.TargetAccountId) == nil {
		return
	}

	friendS := enter.GetSessionByUid(req.TargetAccountId)
	if friendS == nil {
		return
	}

	rsp.FriendDB = game.GetFriendDB(friendS)
	rsp.CafeDBs = game.GetPbCafeDBs(friendS)
}

func CafeGiveGift(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.CafeGiveGiftRequest)
	rsp := response.(*proto.CafeGiveGiftResponse)

	parcelResultList := make([]*game.ParcelResult, 0)

	defer func() {
		rsp.ParcelResultDB = game.ParcelResultDB(s, parcelResultList)
	}()
	consumeResultDB := &proto.ConsumeResultDB{
		RemovedItemServerIds:                    make([]int64, 0),
		RemovedEquipmentServerIds:               make([]int64, 0),
		RemovedFurnitureServerIds:               make([]int64, 0),
		UsedItemServerIdAndRemainingCounts:      make(map[int64]int64),
		UsedEquipmentServerIdAndRemainingCounts: make(map[int64]int64),
		UsedFurnitureServerIdAndRemainingCounts: make(map[int64]int64),
	}
	rsp.ConsumeResultDB = consumeResultDB

	addExp := int64(0)
	for sid, num := range req.ConsumeRequestDB.ConsumeItemServerIdAndCounts {
		itemDb := s.GetItemByKeyId(sid)
		if itemDb == nil || itemDb.StackCount < int32(num) {
			continue
		}
		addExp += num * 60
		parcelResultList = append(parcelResultList, &game.ParcelResult{
			ParcelType: proto.ParcelType_Item,
			ParcelId:   itemDb.UniqueId,
			Amount:     -num,
		})
	}
	// 添加学生好感度
	parcelResultList = append(parcelResultList, &game.ParcelResult{
		ParcelType: proto.ParcelType_FavorExp,
		ParcelId:   req.CharacterUniqueId,
		Amount:     addExp,
	})
}
