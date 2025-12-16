package options

import (
	"context"
	"encoding/json"
	"net/http"
)

// ExchangeInfoService exchange info service
type ExchangeInfoService struct {
	c *Client
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/exchangeInfo",
		secType:  secTypeNone,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ExchangeInfo exchange info
type ExchangeInfo struct {
	Timezone        string           `json:"timezone"`
	ServerTime      int64            `json:"serverTime"`
	OptionContracts []OptionContract `json:"optionContracts"`
	OptionAssets    []OptionAsset    `json:"optionAssets"`
	OptionSymbols   []OptionSymbol   `json:"optionSymbols"`
	RateLimits      []RateLimit      `json:"rateLimits"`
}

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
}

// Option Contract
type OptionContract struct {
	Id          int64  `json:"id"`
	BaseAsset   string `json:"baseAsset"`
	QuoteAsset  string `json:"quoteAsset"`
	Underlying  string `json:"underlying"`
	SettleAsset string `json:"settleAsset"`
}

// Option Asset
type OptionAsset struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Option Symbol
type OptionSymbol struct {
	ContractId           int64                    `json:"contractId"`
	ExpiryDate           int64                    `json:"expiryDate"`
	Filters              []map[string]any `json:"filters"`
	Id                   int64                    `json:"id"`
	Symbol               string                   `json:"symbol"`
	Side                 string                   `json:"side"`
	StrikePrice          string                   `json:"strikePrice"`
	Underlying           string                   `json:"underlying"`
	Unit                 int64                    `json:"unit"`
	MakerFeeRate         string                   `json:"makerFeeRate"`
	TakerFeeRate         string                   `json:"takerFeeRate"`
	MinQty               string                   `json:"minQty"`
	MaxQty               string                   `json:"maxQty"`
	InitialMargin        string                   `json:"initialMargin"`
	MaintenanceMargin    string                   `json:"maintenanceMargin"`
	MinInitialMargin     string                   `json:"minInitialMargin"`
	MinMaintenanceMargin string                   `json:"minMaintenanceMargin"`
	PriceScale           int                      `json:"priceScale"`
	QuantityScale        int                      `json:"quantityScale"`
	QuoteAsset           string                   `json:"quoteAsset"`
}

// LotSizeFilter define lot size filter of symbol
type LotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// PriceFilter define price filter of symbol
type PriceFilter struct {
	MaxPrice string `json:"maxPrice"`
	MinPrice string `json:"minPrice"`
	TickSize string `json:"tickSize"`
}

// LotSizeFilter return lot size filter of symbol
func (s *OptionSymbol) LotSizeFilter() *LotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeLotSize) {
			f := &LotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PriceFilter return price filter of symbol
func (s *OptionSymbol) PriceFilter() *PriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePrice) {
			f := &PriceFilter{}
			if i, ok := filter["maxPrice"]; ok {
				f.MaxPrice = i.(string)
			}
			if i, ok := filter["minPrice"]; ok {
				f.MinPrice = i.(string)
			}
			if i, ok := filter["tickSize"]; ok {
				f.TickSize = i.(string)
			}
			return f
		}
	}
	return nil
}
