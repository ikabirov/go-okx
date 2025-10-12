package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventOpenInterest) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeOpenInterest(context.Background(), "BTC-USDT-SWAP", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
