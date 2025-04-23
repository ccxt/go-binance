package futures

import (
	"context"
	"encoding/json"
)

// SymbolConfigService get futures symbol configuration
type SymbolConfigService struct {
	c      *Client
	symbol *string
}

// SymbolConfig define futures symbol configuration
type SymbolConfig struct {
	Symbol           string `json:"symbol"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  bool   `json:"isAutoAddMargin"`
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
}

// Symbol set symbol
func (s *SymbolConfigService) Symbol(symbol string) *SymbolConfigService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *SymbolConfigService) Do(ctx context.Context) ([]*SymbolConfig, error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v1/symbolConfig",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*SymbolConfig
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
