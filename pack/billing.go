package pack

import (
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/protocol/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func BillingPurchaseListByYostar(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.BillingPurchaseListByYostarResponse)

	rsp.CountList = make([]*proto.PurchaseCountDB, 0)
	rsp.OrderList = make([]*proto.PurchaseOrderDB, 0)
	rsp.MonthlyProductList = make([]*proto.MonthlyProductPurchaseDB, 0)
	rsp.BlockedProductDBs = make([]*proto.BlockedProductDB, 0)
}

func BillingTransactionStartByYostar(s *enter.Session, request, response mx.Message) {
	//req := request.(*proto.BillingTransactionStartByYostarRequest)
	rsp := response.(*proto.BillingTransactionStartByYostarResponse)

	rsp.PurchaseServerTag = proto.PurchaseServerTag_TestIn
	rsp.PurchaseServerCallbackUrl = ""
}
