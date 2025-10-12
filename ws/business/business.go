package business

import (
	"context"

	"github.com/ikabirov/go-okx/ws"
)

type Business struct {
	C *ws.Client
}

func NewBusiness() *Business {
	business := &Business{
		C: ws.DefaultClientBusiness,
	}
	return business
}

// subscribe
func (p *Business) Subscribe(ctx context.Context, args any, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(ctx, subscribe, nil)
}
