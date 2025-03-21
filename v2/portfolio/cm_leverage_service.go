package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeCMInitialLeverageService change user's initial leverage of specific symbol in CM
type ChangeCMInitialLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeCMInitialLeverageService) Symbol(symbol string) *ChangeCMInitialLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeCMInitialLeverageService) Leverage(leverage int) *ChangeCMInitialLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeCMInitialLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *CMLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/cm/leverage",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("leverage", s.leverage)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CMLeverage)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMLeverage define leverage info
type CMLeverage struct {
	Leverage int    `json:"leverage"`
	MaxQty   string `json:"maxQty"`
	Symbol   string `json:"symbol"`
}
