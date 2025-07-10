package pack

import (
	"./common/enter"
	"./protocol/mx"
	"./protocol/proto"
)

func CraftInfoList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.CraftInfoListResponse)

	rsp.CraftInfos = make([]*proto.CraftInfoDB, 0)
	rsp.ShiftingCraftInfos = make([]*proto.ShiftingCraftInfoDB, 0)
}
