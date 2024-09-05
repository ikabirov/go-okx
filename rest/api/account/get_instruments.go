package account

import "github.com/ikabirov/go-okx/rest/api"

func NewGetInstruments(param *GetInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetInstrumentsResponse{}
}

type GetInstrumentsParam struct {
	InstType   string `url:"instType"`
	InstId     string `url:"instId,omitempty"`
	InstFamily string `url:"instFamily,omitempty"`
	Uly        string `url:"uly,omitempty"`
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
