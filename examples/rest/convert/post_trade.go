package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples/rest"
	"github.com/ikabirov/go-okx/rest/api/convert"
)

func main() {
	param := &convert.PostTradeParam{
		BaseCcy:  "BTC",
		QuoteCcy: "USDT",
		Side:     "buy",
		Sz:       "1",
		SzCcy:    "USDT",
		QuoteId:  "123456789",
	}
	req, resp := convert.NewPostTrade(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.PostTradeResponse))
}
