package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMPositionRiskService get UM position risk information
type GetUMPositionRiskService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *GetUMPositionRiskService) Symbol(symbol string) *GetUMPositionRiskService {
	s.symbol = &symbol
	return s
}

type UMPositionRisk struct {
	Symbol                 string `json:"symbol"`                     // symbol name
	PositionAmt            string `json:"positionAmt"`                // position amount
	EntryPrice             string `json:"entryPrice"`                 // average entry price
	MarkPrice              string `json:"markPrice,omitempty"`        // mark price (only in position risk endpoint)
	UnRealizedProfit       string `json:"unRealizedProfit"`           // unrealized profit
	LiquidationPrice       string `json:"liquidationPrice,omitempty"` // liquidation price (only in position risk endpoint)
	Leverage               string `json:"leverage"`                   // current initial leverage
	MaxNotionalValue       string `json:"maxNotionalValue,omitempty"` // maximum notional value (position risk)
	PositionSide           string `json:"positionSide"`               // position side
	Notional               string `json:"notional,omitempty"`         // notional value (only in position risk endpoint)
	UpdateTime             int64  `json:"updateTime"`                 // last update time
}

// Do send request
func (s *GetUMPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*UMPositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*UMPositionRisk{}, err
	}
	res = make([]*UMPositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UMPositionRisk{}, err
	}
	return res, nil
}
