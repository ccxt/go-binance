package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type AssetIndexService struct {
	c      *Client
	symbol *string // for example, BTCUSD
}

type AssetIndex struct {
	Symbol                string `json:"symbol"`
	Time                  uint64 `json:"time"`
	Index                 string `json:"index"`
	BidBuffer             string `json:"bidBuffer"`
	AskBuffer             string `json:"askBuffer"`
	BidRate               string `json:"bidRate"`
	AskRate               string `json:"askRate"`
	AutoExchangeBidBuffer string `json:"autoExchangeBidBuffer"`
	AutoExchangeAskBuffer string `json:"autoExchangeAskBuffer"`
	AutoExchangeBidRate   string `json:"autoExchangeBidRate"`
	AutoExchangeAskRate   string `json:"autoExchangeAskRate"`
}

func (s *AssetIndexService) Symbol(symbol string) *AssetIndexService {
	s.symbol = &symbol
	return s
}

func (s *AssetIndexService) Do(ctx context.Context, opts ...RequestOption) (res []*AssetIndex, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/assetIndex",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*AssetIndex, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
