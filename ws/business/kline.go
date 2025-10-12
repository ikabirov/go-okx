package business

import (
	"context"
	"encoding/json"

	"github.com/ikabirov/go-okx/ws"
)

type HandlerKline func(EventKline)

type EventKline struct {
	Arg  ws.Args    `json:"arg"`
	Data [][]string `json:"data"`
}

// default subscribe
func SubscribeKline(ctx context.Context, args *ws.Args, handler HandlerKline, handlerError ws.HandlerError, simulated bool) error {
	h := func(message []byte) {
		var event EventKline
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewBusiness().Subscribe(ctx, args, h, handlerError)
}
