package pack

import (
	"./common/enter"
	"./game"
	"./protocol/mx"
	"./protocol/proto"
)

func MemoryLobbyList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.MemoryLobbyListResponse)

	rsp.MemoryLobbyDBs = game.GetMemoryLobbyDBs(s)
}
