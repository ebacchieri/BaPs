package game

import (
	"time"

	"github.com/gucooing/BaPs/protocol/mx"

	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/gdconf"
	"github.com/gucooing/BaPs/pkg/alg"
	"github.com/gucooing/BaPs/protocol/proto"
)

func NewMomoTalkBin() *sro.MomoTalkBin {
	return &sro.MomoTalkBin{}
}

func GetMomoTalkBin(s *enter.Session) *sro.MomoTalkBin {
	bin := GetPlayerBin(s)
	if bin == nil {
		return nil
	}
	if bin.MomoTalkBin == nil {
		bin.MomoTalkBin = NewMomoTalkBin()
	}
	return bin.MomoTalkBin
}

func GetFavorScheduleInfoList(s *enter.Session) map[int64]*sro.FavorScheduleInfo {
	bin := GetMomoTalkBin(s)
	if bin == nil {
		return nil
	}
	if bin.FavorScheduleInfoList == nil {
		bin.FavorScheduleInfoList = make(map[int64]*sro.FavorScheduleInfo)
	}
	return bin.FavorScheduleInfoList
}

func GetFavorScheduleInfo(s *enter.Session, characterId int64) *sro.FavorScheduleInfo {
	bin := GetFavorScheduleInfoList(s)
	if bin[characterId] == nil {
		bin[characterId] = &sro.FavorScheduleInfo{
			ScheduleGroupList: make([]int64, 0),
			MomoTalkInfoList:  make(map[int64]*sro.MomoTalkInfo),
			CurMessageGroupId: 0,
		}
	}
	return bin[characterId]
}

func UpMomoTalkInfo(s *enter.Session, characterId, gid, messageId int64) {
	// 验证是否有这个id
	conf := gdconf.GetAcademyMessangerExcel(gid)
	if conf == nil {
		return
	}
	// 简单验证是否允许
	if conf.MessageCondition == "FavorRankUp" {
		characterInfo := GetCharacterInfo(s, characterId)
		if characterInfo == nil ||
			int64(characterInfo.FavorRank) < conf.ConditionValue {
			return
		}
	}
	bin := GetFavorScheduleInfo(s, characterId)
	if bin == nil {
		return
	}
	if bin.MomoTalkInfoList == nil {
		bin.MomoTalkInfoList = make(map[int64]*sro.MomoTalkInfo)
	}
	bin.CurMessageGroupId = gid
	if info, ok := bin.MomoTalkInfoList[gid]; ok {
		info.MessageGroupId = alg.MaxInt64(info.MessageGroupId, gid)
		info.ChosenMessageId = alg.MaxInt64(info.ChosenMessageId, messageId)
		info.ChosenDate = time.Now().Unix()
	}
	bin.MomoTalkInfoList[gid] = &sro.MomoTalkInfo{
		MessageGroupId:  gid,
		ChosenMessageId: messageId,
		ChosenDate:      time.Now().Unix(),
	}
}

func UpScheduleGroup(s *enter.Session, scheduleId int64) []*ParcelResult {
	conf := gdconf.GetAcademyFavorScheduleExcel(scheduleId)
	if conf == nil || s == nil {
		return nil
	}
	bin := GetFavorScheduleInfo(s, conf.CharacterId)
	if bin == nil {
		return nil
	}
	if bin.ScheduleGroupList == nil {
		bin.ScheduleGroupList = make([]int64, 0)
	}
	for _, id := range bin.ScheduleGroupList {
		if id == scheduleId {
			return nil
		}
	}
	bin.ScheduleGroupList = append(bin.ScheduleGroupList, scheduleId)

	return GetParcelResultList(conf.RewardParcelType, conf.RewardParcelId, conf.RewardAmount, false)
}

func GetMomoTalkChoiceDBs(s *enter.Session, characterId int64) []*proto.MomoTalkChoiceDB {
	list := make([]*proto.MomoTalkChoiceDB, 0)
	characterInfo := GetCharacterInfo(s, characterId)
	bin := GetFavorScheduleInfo(s, characterId)
	if bin == nil || characterInfo == nil {
		return list
	}
	for _, info := range bin.MomoTalkInfoList {
		list = append(list, &proto.MomoTalkChoiceDB{
			CharacterDBId:   characterInfo.ServerId,
			MessageGroupId:  info.MessageGroupId,
			ChosenMessageId: info.ChosenMessageId,
			ChosenDate:      mx.Unix(info.ChosenDate, 0),
		})
	}
	return list
}
