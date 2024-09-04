package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples/rest"
	"github.com/ikabirov/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostCancelOrderParam{
		InstId: "XRP-USDT",
		OrdId:  "515101542723186689",
	}
	req, resp := trade.NewPostCancelOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostCancelOrderResponse))
}
