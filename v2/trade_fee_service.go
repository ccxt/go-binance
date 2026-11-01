package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// TradeFeeService shows current trade fee for all symbols available
type TradeFeeService struct {
	c      *Client
	symbol *string
}

// Symbol set the symbol parameter for the request
func (s *TradeFeeService) Symbol(symbol string) *TradeFeeService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *TradeFeeService) Do(ctx context.Context) (res []*TradeFeeDetails, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/tradeFee",
		secType:  secTypeSigned,
	}

	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return res, err
	}
	res = make([]*TradeFeeDetails, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TradeFeeDetails represents details about fees
type TradeFeeDetails struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
