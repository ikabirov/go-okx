package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/ws/public"
)

func main() {
	handler := func(c public.EventPriceLimit) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := public.SubscribePriceLimit(context.Background(), "LTC-USD-190628", handler, handlerError, false); err != nil {
		panic(err)
	}
	select {}
}
