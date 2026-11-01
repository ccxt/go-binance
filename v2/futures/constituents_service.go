package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type ConstituentsService struct {
	c      *Client
	symbol string // for example, BTCUSDT
}

type ConstituentsServiceRsp struct {
	Symbol       string          `json:"symbol"`
	Time         uint64          `json:"time"`
	Constituents []*Constituents `json:"constituents"`
}

type Constituents struct {
	Exchange string `json:"exchange"`
	Symbol   string `json:"symbol"`
}

func (s *ConstituentsService) Symbol(symbol string) *ConstituentsService {
	s.symbol = symbol
	return s
}

func (s *ConstituentsService) Do(ctx context.Context, opts ...RequestOption) (res *ConstituentsServiceRsp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/constituents",
	}
	r.setParam("symbol", s.symbol)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ConstituentsServiceRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
