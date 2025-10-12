package public

import (
	"context"
	"encoding/json"

	"github.com/ikabirov/go-okx/ws"
)

type HandlerMarkPrice func(EventMarkPrice)

type EventMarkPrice struct {
	Arg  ws.Args     `json:"arg"`
	Data []MarkPrice `json:"data"`
}

type MarkPrice struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	MarkPx   string `json:"markPx"`
	Ts       int64  `json:"ts,string"`
}

// default subscribe
func SubscribeMarkPrice(ctx context.Context, instId string, handler HandlerMarkPrice, handlerError ws.HandlerError, simulated bool) error {
	args := &ws.Args{
		Channel: "mark-price",
		InstId:  instId,
	}

	h := func(message []byte) {
		var event EventMarkPrice
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(ctx, args, h, handlerError)
}
