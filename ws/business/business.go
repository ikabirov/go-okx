package business

import (
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
func (p *Business) Subscribe(args interface{}, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}
