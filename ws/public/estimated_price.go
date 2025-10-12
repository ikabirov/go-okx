package public

import (
	"context"
	"encoding/json"

	"github.com/ikabirov/go-okx/ws"
)

type HandlerEstimatedPrice func(EventEstimatedPrice)

type EventEstimatedPrice struct {
	Arg  ws.Args          `json:"arg"`
	Data []EstimatedPrice `json:"data"`
}

type EstimatedPrice struct {
	InstType string `json:"instType"`
	InstId   string `json:"instId"`
	SettlePx string `json:"settlePx"`
	Ts       int64  `json:"ts,string"`
}

// default subscribe
func SubscribeEstimatedPrice(ctx context.Context, args *ws.Args, handler HandlerEstimatedPrice, handlerError ws.HandlerError, simulated bool) error {
	args.Channel = "estimated-price"

	h := func(message []byte) {
		var event EventEstimatedPrice
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPublic(simulated).Subscribe(ctx, args, h, handlerError)
}
