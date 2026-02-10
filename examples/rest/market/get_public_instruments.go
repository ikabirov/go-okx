package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples/rest"
	"github.com/ikabirov/go-okx/rest/api/market"
)

func main() {
	param := &market.GetPublicInstrumentsParam{
		InstType: "SWAP",
		InstId:   "BTC-USDT-SWAP",
	}
	req, resp := market.NewGetPublicInstruments(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}

	res := resp.(*market.GetPublicInstrumentsResponse)

	log.Println(req, res)
}
