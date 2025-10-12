package main

import (
	"context"
	"log"

	"github.com/ikabirov/go-okx/examples"
	"github.com/ikabirov/go-okx/ws"
	"github.com/ikabirov/go-okx/ws/private"
)

func main() {
	args := &ws.Args{
		InstType: "SPOT",
	}
	handler := func(c private.EventOrders) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeOrders(context.Background(), args, examples.Auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
