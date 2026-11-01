package options

import (
	"context"
	"encoding/json"
	"net/http"
)

type Mark struct {
	Symbol           string `json:"symbol"`
	MarkPrice        string `json:"markPrice"`
	BidIV            string `json:"bidIV"`
	AskIV            string `json:"askIV"`
	MarkIV           string `json:"markIV"`
	Delta            string `json:"delta"`
	Theta            string `json:"theta"`
	Gamma            string `json:"gamma"`
	Vega             string `json:"vega"`
	HighPriceLimit   string `json:"highPriceLimit"`
	LowPriceLimit    string `json:"lowPriceLimit"`
	RiskFreeInterest string `json:"riskFreeInterest"`
}

// MarkService list recent trades in orderbook
type MarkService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *MarkService) Symbol(symbol string) *MarkService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarkService) Do(ctx context.Context, opts ...RequestOption) (res []*Mark, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/mark",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Mark{}, err
	}
	res = make([]*Mark, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Mark{}, err
	}
	return res, nil
}
