package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/ws"
	"github.com/ikabirov/go-okx/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "mark-price-candle1m",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventMarkPriceKline) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeMarkPriceKline(context.Background(), args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
