package portfolio

import (
	"context"
	"encoding/json"
)

// UMSymbolConfigService get UM futures symbol configuration
type UMSymbolConfigService struct {
	c      *Client
	symbol *string
}

// UMSymbolConfig define UM futures symbol configuration
type UMSymbolConfig struct {
	Symbol           string `json:"symbol"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  bool   `json:"isAutoAddMargin"`
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
}

// Symbol set symbol
func (s *UMSymbolConfigService) Symbol(symbol string) *UMSymbolConfigService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *UMSymbolConfigService) Do(ctx context.Context) ([]*UMSymbolConfig, error) {
	r := &request{
		method:   "GET",
		endpoint: "/papi/v1/um/symbolConfig",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*UMSymbolConfig
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
