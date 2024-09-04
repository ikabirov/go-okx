package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples"
	"github.com/ikabirov/go-okx/ws/private"
)

func main() {
	handler := func(c private.EventBalanceAndPosition) {
		log.Println(c)
	}
	handlerError := func(err error) {
		panic(err)
	}
	if err := private.SubscribeBalanceAndPosition(examples.Auth, handler, handlerError); err != nil {
		panic(err)
	}
	select {}
}
