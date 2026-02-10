package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples/rest"
	"github.com/ikabirov/go-okx/rest/api/market"
)

func main() {
	param := &market.GetFundingRateParam{
		InstId: "BTC-USD-SWAP",
	}
	req, resp := market.NewGetFundingRate(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}

	res := resp.(*market.GetFundingRateResponse)

	log.Println(req, res)
}
