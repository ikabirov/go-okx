package market

import "github.com/ikabirov/go-okx/rest/api"

func NewGetFundingRate(param *GetFundingRateParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/public/funding-rate",
		Method: api.MethodGet,
		Param:  param,
	}, &GetFundingRateResponse{}
}

type GetFundingRateParam struct {
	InstId string `url:"instId"`
}

type GetFundingRateResponse struct {
	api.Response
	Data []FundingRateData `json:"data"`
}

type FundingRateData struct {
	FundingRate     string `json:"fundingRate"`
	FundingTime     string `json:"fundingTime"`
	InstId          string `json:"instId"`
	InstType        string `json:"instType"`
	NextFundingRate string `json:"nextFundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
}
