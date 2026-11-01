package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMLeverageBracketService get UM notional and leverage brackets
type GetUMLeverageBracketService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *GetUMLeverageBracketService) Symbol(symbol string) *GetUMLeverageBracketService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *GetUMLeverageBracketService) Do(ctx context.Context, opts ...RequestOption) (res []*LeverageBracket, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/leverageBracket",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*LeverageBracket{}, err
	}
	res = make([]*LeverageBracket, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LeverageBracket{}, err
	}
	return res, nil
}

// LeverageBracket define leverage bracket
type LeverageBracket struct {
	Symbol       string    `json:"symbol"`
	NotionalCoef string    `json:"notionalCoef"`
	Brackets     []Bracket `json:"brackets"`
}

// Bracket define bracket info
type Bracket struct {
	Bracket          int     `json:"bracket"`          // Notional bracket
	InitialLeverage  int     `json:"initialLeverage"`  // Max initial leverage for this bracket
	NotionalCap      float64 `json:"notionalCap"`      // Cap notional of this bracket
	NotionalFloor    float64 `json:"notionalFloor"`    // Notional threshold of this bracket
	MaintMarginRatio float64 `json:"maintMarginRatio"` // Maintenance ratio for this bracket
	Cum              float64 `json:"cum"`              // Auxiliary number for quick calculation
}
