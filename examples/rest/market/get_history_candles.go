package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples/rest"
	"github.com/ikabirov/go-okx/rest/api/market"
)

func main() {
	param := &market.GetCandlesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetHistoryCandles(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetCandlesResponse))
}
