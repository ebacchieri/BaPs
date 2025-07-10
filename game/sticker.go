package game

import (
	"errors"
	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/protocol/proto"
)

func NewStickerBin() *sro.StickerBin {
	return &sro.StickerBin{}
}

func GetStickerBin(s *enter.Session) *sro.StickerBin {
	bin := GetPlayerBin(s)
	if bin == nil {
		return nil
	}
	if bin.StickerBin == nil {
		bin.StickerBin = NewStickerBin()
	}
	return bin.StickerBin
}

func GetStickerMap(s *enter.Session) map[int64]*sro.StickerInfo {
	bin := GetStickerBin(s)
	if bin == nil {
		return nil
	}
	if bin.StickerMap == nil {
		bin.StickerMap = make(map[int64]*sro.StickerInfo)
	}
	return bin.StickerMap
}

func GetStickerInfo(s *enter.Session, stickerId int64) *sro.StickerInfo {
	bin := GetStickerMap(s)
	if bin == nil {
		return nil
	}
	return bin[stickerId]
}

func GetStickerBookDB(s *enter.Session) *proto.StickerBookDB {
	info := &proto.StickerBookDB{
		AccountId:        s.AccountServerId,
		UnusedStickerDBs: make([]*proto.StickerDB, 0),
		UsedStickerDBs:   make([]*proto.StickerDB, 0),
	}
	for _, v := range GetStickerMap(s) {
		stickerDB := &proto.StickerDB{
			Type:            proto.ParcelType_Sticker,
			StickerUniqueId: v.StickerId,
		}
		if v.Used {
			info.UsedStickerDBs = append(info.UsedStickerDBs, stickerDB)
		} else {
			info.UnusedStickerDBs = append(info.UnusedStickerDBs, stickerDB)
		}
	}

	return info
}

func AddSticker(s *enter.Session, stickerId int64) error {
	list := GetStickerMap(s)
	if _, ok := list[stickerId]; ok {
		return errors.New("重复添加")
	}
	list[stickerId] = &sro.StickerInfo{StickerId: stickerId}
	return nil
}

func UseSticker(s *enter.Session, stickerId int64) error {
	bin := GetStickerInfo(s, stickerId)
	if bin == nil {
		return errors.New("贴纸不存在")
	}
	bin.Used = true
	return nil
}

func GetStickerDBById(s *enter.Session, stickerId int64) *proto.StickerDB {
	bin := GetStickerInfo(s, stickerId)
	return &proto.StickerDB{
		Type:            proto.ParcelType_Sticker,
		StickerUniqueId: bin.StickerId,
	}
}
