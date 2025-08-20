package game

import (
	"errors"
	"time"

	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/protocol/proto"
)

func NewMemoryLobbyBin() *sro.MemoryLobbyBin {
	return &sro.MemoryLobbyBin{}
}

func GetMemoryLobbyBin(s *enter.Session) *sro.MemoryLobbyBin {
	bin := GetPlayerBin(s)
	if bin == nil {
		return nil
	}
	if bin.MemoryLobbyBin == nil {
		bin.MemoryLobbyBin = NewMemoryLobbyBin()
	}
	return bin.MemoryLobbyBin
}

func GetMemoryLobbyInfoList(s *enter.Session) map[int64]*sro.MemoryLobbyInfo {
	bin := GetMemoryLobbyBin(s)
	if bin == nil {
		return nil
	}
	if bin.MemoryLobbyInfoList == nil {
		bin.MemoryLobbyInfoList = make(map[int64]*sro.MemoryLobbyInfo)
	}
	return bin.MemoryLobbyInfoList
}

func GetMemoryLobbyInfo(s *enter.Session, memoryLobbyId int64) *sro.MemoryLobbyInfo {
	bin := GetMemoryLobbyInfoList(s)
	if bin == nil {
		return nil
	}
	return bin[memoryLobbyId]
}

func GetMemoryLobbyDBs(s *enter.Session) []*proto.MemoryLobbyDB {
	list := make([]*proto.MemoryLobbyDB, 0)
	if s == nil {
		return list
	}
	for _, info := range GetMemoryLobbyInfoList(s) {
		list = append(list, &proto.MemoryLobbyDB{
			Type:                proto.ParcelType_MemoryLobby,
			MemoryLobbyUniqueId: info.MemoryLobbyId,
		})
	}

	return list
}

func AddMemoryLobby(s *enter.Session, memoryLobbyUniqueId int64) error {
	list := GetMemoryLobbyInfoList(s)
	if list == nil {
		return errors.New("数据错误")
	}
	if _, ok := list[memoryLobbyUniqueId]; ok {
		return errors.New("重复添加")
	}
	list[memoryLobbyUniqueId] = &sro.MemoryLobbyInfo{
		MemoryLobbyId: memoryLobbyUniqueId,
		ChosenDate:    time.Now().Unix(),
	}
	return nil
}

func GetMemoryLobbyDB(s *enter.Session, memoryLobbyId int64) *proto.MemoryLobbyDB {
	bin := GetMemoryLobbyInfo(s, memoryLobbyId)
	if bin == nil {
		return nil
	}
	return &proto.MemoryLobbyDB{
		Type:                proto.ParcelType_MemoryLobby,
		MemoryLobbyUniqueId: bin.MemoryLobbyId,
	}
}
