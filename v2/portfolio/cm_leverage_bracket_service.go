package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMLeverageBracketService get CM notional and leverage brackets
type GetCMLeverageBracketService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *GetCMLeverageBracketService) Symbol(symbol string) *GetCMLeverageBracketService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetCMLeverageBracketService) Do(ctx context.Context, opts ...RequestOption) (res []*CMLeverageBracket, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/leverageBracket",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CMLeverageBracket{}, err
	}
	res = make([]*CMLeverageBracket, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CMLeverageBracket{}, err
	}
	return res, nil
}

// CMLeverageBracket define CM leverage bracket
type CMLeverageBracket struct {
	Symbol   string      `json:"symbol"`
	Brackets []CMBracket `json:"brackets"`
}

// CMBracket define CM bracket info
type CMBracket struct {
	Bracket          int     `json:"bracket"`          // bracket level
	InitialLeverage  int     `json:"initialLeverage"`  // the maximum leverage
	QtyCap           float64 `json:"qtyCap"`           // upper edge of base asset quantity
	QtyFloor         float64 `json:"qtyFloor"`         // lower edge of base asset quantity
	MaintMarginRatio float64 `json:"maintMarginRatio"` // maintenance margin rate
	Cum              float64 `json:"cum"`              // Auxiliary number for quick calculation
}
