package pack

import (
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/protocol/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func StickerLogin(s *enter.Session, request, response mx.Message) {
	//req := request.(*proto.StickerLoginRequest)
	rsp := response.(*proto.StickerLoginResponse)

	rsp.StickerBookDB = game.GetStickerBookDB(s)
}

func StickerLobby(s *enter.Session, request, response mx.Message) {
	//req := request.(*proto.StickerLobbyRequest)
	rsp := response.(*proto.StickerLobbyResponse)

	rsp.ReceivedStickerDBs = make([]*proto.StickerDB, 0)
	rsp.StickerBookDB = game.GetStickerBookDB(s)
}

func StickerUseSticker(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.StickerUseStickerRequest)
	rsp := response.(*proto.StickerUseStickerResponse)

	game.UseSticker(s, req.StickerUniqueId)
	rsp.StickerBookDB = game.GetStickerBookDB(s)
}
