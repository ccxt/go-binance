package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeUMInitialLeverageService change user's initial leverage of specific symbol in UM
type ChangeUMInitialLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeUMInitialLeverageService) Symbol(symbol string) *ChangeUMInitialLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeUMInitialLeverageService) Leverage(leverage int) *ChangeUMInitialLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeUMInitialLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *UMLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("leverage", s.leverage)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMLeverage)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMLeverage define leverage info
type UMLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}
