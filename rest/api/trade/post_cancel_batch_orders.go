package trade

import "github.com/ikabirov/go-okx/rest/api"

func NewPostCancelBatchOrders(param []*PostCancelOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/cancel-batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostCancelOrderResponse{}
}
