package pack

import (
	"github.com/ebacchieri/BaPs/common/enter"
	"github.com/ebacchieri/BaPs/game"
	"github.com/ebacchieri/BaPs/protocol/mx"
	"github.com/ebacchieri/BaPs/protocol/proto"
)

func AttachmentGet(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.AttachmentGetResponse)

	rsp.AccountAttachmentDB = game.GetAccountAttachmentDB(s)
}

func AttachmentEmblemList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.AttachmentEmblemListResponse)

	rsp.EmblemDBs = game.GetEmblemDBs(s)
}

func AttachmentEmblemAcquire(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AttachmentEmblemAcquireRequest)
	rsp := response.(*proto.AttachmentEmblemAcquireResponse)

	game.UpEmblemInfoList(s, req.UniqueIds)
	rsp.EmblemDBs = game.GetEmblemDBs(s)
}

func AttachmentEmblemAttach(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AttachmentEmblemAttachRequest)
	rsp := response.(*proto.AttachmentEmblemAttachResponse)

	game.SetEmblemUniqueId(s, req.UniqueId)
	rsp.AttachmentDB = game.GetAccountAttachmentDB(s)
}
