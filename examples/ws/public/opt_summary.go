package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventOptSummary) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribeOptSummary(context.Background(), "BTC-USDT", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
