package main

import (
	"log"

	"github.com/ikabirov/go-okx/examples/rest"
	"github.com/ikabirov/go-okx/rest/api"
	"github.com/ikabirov/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetFillsParam{
		InstType: api.InstTypeSPOT,
		PagingParam: api.PagingParam{
			Limit: 2,
		},
	}
	req, resp := trade.NewGetFillsHistory(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetFillsResponse))
}
