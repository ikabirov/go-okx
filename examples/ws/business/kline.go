package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/ws"
	"github.com/ikabirov/go-okx/ws/business"
)

func main() {
	args := &ws.Args{
		Channel: "candle1m",
		InstId:  "BTC-USDT",
	}
	handler := func(c business.EventKline) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := business.SubscribeKline(context.Background(), args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
