package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type IndexInfoService struct {
	c      *Client
	symbol *string // for example, FOOTBALLUSDT
}

type IndexInfo struct {
	Symbol        string           `json:"symbol"`
	Time          uint64           `json:"time"`
	Component     string           `json:"component"`
	BaseAssetList []*BaseAssetList `json:"baseAssetList"`
}

type BaseAssetList struct {
	BaseAsset          string `json:"baseAsset"`
	QuoteAsset         string `json:"quoteAsset"`
	WeightInQuantity   string `json:"weightInQuantity"`
	WeightInPercentage string `json:"weightInPercentage"`
}

func (s *IndexInfoService) Symbol(symbol string) *IndexInfoService {
	s.symbol = &symbol
	return s
}

func (s *IndexInfoService) Do(ctx context.Context, opts ...RequestOption) (res []*IndexInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/v1/indexInfo",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*IndexInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
