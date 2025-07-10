package pack

import (
	"github.com/ebacchieri/BaPs/common/enter"
	"github.com/ebacchieri/BaPs/game"
	"github.com/ebacchieri/BaPs/protocol/mx"
	"github.com/ebacchieri/BaPs/protocol/proto"
)

func MemoryLobbyList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.MemoryLobbyListResponse)

	rsp.MemoryLobbyDBs = game.GetMemoryLobbyDBs(s)
}
