package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/ws"
	"github.com/ikabirov/go-okx/ws/public"
)

func main() {
	args := &ws.Args{
		Channel: "books5",
		InstId:  "BTC-USDT",
	}
	handler := func(c public.EventBooks) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeBooks(context.Background(), args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
