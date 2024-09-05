package account

import "github.com/ikabirov/go-okx/rest/api"

func NewGetInstruments() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/instruments",
		Method: api.MethodGet,
	}, &GetInstrumentsResponse{}
}

type GetInstrumentsResponse struct {
	api.Response
	Data []Instrument `json:"data"`
}

type Instrument struct {
	InstId                string `json:"instId"`
	ContractValue         string `json:"ctVal"`
	ContractMult          string `json:"ctMult"`
	ContractValueCurrency string `json:"ctValCcy"`
	LotSize               string `json:"lotSz"`
}
