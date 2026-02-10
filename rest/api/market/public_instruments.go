package market

import "github.com/ikabirov/go-okx/rest/api"

func NewGetPublicInstruments(param *GetPublicInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/public/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetPublicInstrumentsResponse{}
}

type GetPublicInstrumentsParam struct {
	InstType   string `url:"instType"`
	InstId     string `url:"instId,omitempty"`
	Uly        string `url:"uly,omitempty"`
	InstFamily string `url:"instFamily,omitempty"`
}

type GetPublicInstrumentsResponse struct {
	api.Response
	Data []PublicInstrument `json:"data"`
}

type PublicInstrument struct {
	InstType                string        `json:"instType"`
	InstId                  string        `json:"instId"`
	Uly                     string        `json:"uly,omitempty"`
	GroupId                 string        `json:"groupId,omitempty"`
	InstFamily              string        `json:"instFamily,omitempty"`
	BaseCcy                 string        `json:"baseCcy,omitempty"`
	QuoteCcy                string        `json:"quoteCcy,omitempty"`
	SettleCcy               string        `json:"settleCcy,omitempty"`
	CtVal                   string        `json:"ctVal,omitempty"`
	CtMult                  string        `json:"ctMult,omitempty"`
	CtValCcy                string        `json:"ctValCcy,omitempty"`
	OptType                 string        `json:"optType,omitempty"`
	Stk                     string        `json:"stk,omitempty"`
	ListTime                string        `json:"listTime,omitempty"`
	AuctionEndTime          string        `json:"auctionEndTime,omitempty"`
	ContTdSwTime            string        `json:"contTdSwTime,omitempty"`
	PreMktSwTime            string        `json:"preMktSwTime,omitempty"`
	OpenType                string        `json:"openType,omitempty"`
	ExpTime                 string        `json:"expTime,omitempty"`
	Lever                   string        `json:"lever,omitempty"`
	TickSz                  string        `json:"tickSz"`
	LotSz                   string        `json:"lotSz"`
	MinSz                   string        `json:"minSz"`
	CtType                  string        `json:"ctType,omitempty"`
	Alias                   string        `json:"alias,omitempty"`
	State                   string        `json:"state"`
	RuleType                string        `json:"ruleType,omitempty"`
	MaxLmtSz                string        `json:"maxLmtSz"`
	MaxMktSz                string        `json:"maxMktSz"`
	MaxLmtAmt               string        `json:"maxLmtAmt,omitempty"`
	MaxMktAmt               string        `json:"maxMktAmt,omitempty"`
	MaxTwapSz               string        `json:"maxTwapSz,omitempty"`
	MaxIcebergSz            string        `json:"maxIcebergSz,omitempty"`
	MaxTriggerSz            string        `json:"maxTriggerSz,omitempty"`
	MaxStopSz               string        `json:"maxStopSz,omitempty"`
	FutureSettlement        bool          `json:"futureSettlement,omitempty"`
	TradeQuoteCcyList       []string      `json:"tradeQuoteCcyList,omitempty"`
	InstIdCode              int           `json:"instIdCode,omitempty"`
	UpcChg                  []UpcomingChange `json:"upcChg,omitempty"`
}

type UpcomingChange struct {
	Param    string `json:"param"`
	NewValue string `json:"newValue"`
	EffTime  string `json:"effTime"`
}
