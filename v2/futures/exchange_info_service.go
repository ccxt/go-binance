package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// ExchangeInfoService exchange info service
type ExchangeInfoService struct {
	c *Client
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/exchangeInfo",
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
	Timezone        string      `json:"timezone"`
	ServerTime      int64       `json:"serverTime"`
	RateLimits      []RateLimit `json:"rateLimits"`
	ExchangeFilters []any       `json:"exchangeFilters"`
	Symbols         []Symbol    `json:"symbols"`
}

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
}

// Symbol market symbol
type Symbol struct {
	Symbol                string            `json:"symbol"`
	Pair                  string            `json:"pair"`
	ContractType          ContractType      `json:"contractType"`
	DeliveryDate          int64             `json:"deliveryDate"`
	OnboardDate           int64             `json:"onboardDate"`
	Status                string            `json:"status"`
	MaintMarginPercent    string            `json:"maintMarginPercent"`
	RequiredMarginPercent string            `json:"requiredMarginPercent"`
	PricePrecision        int               `json:"pricePrecision"`
	QuantityPrecision     int               `json:"quantityPrecision"`
	BaseAssetPrecision    int               `json:"baseAssetPrecision"`
	QuotePrecision        int               `json:"quotePrecision"`
	UnderlyingType        string            `json:"underlyingType"`
	UnderlyingSubType     []string          `json:"underlyingSubType"`
	SettlePlan            int64             `json:"settlePlan"`
	TriggerProtect        string            `json:"triggerProtect"`
	OrderType             []OrderType       `json:"orderType"`
	TimeInForce           []TimeInForceType `json:"timeInForce"`
	Filters               []map[string]any  `json:"filters"`
	QuoteAsset            string            `json:"quoteAsset"`
	MarginAsset           string            `json:"marginAsset"`
	BaseAsset             string            `json:"baseAsset"`
	LiquidationFee        string            `json:"liquidationFee"`
	MarketTakeBound       string            `json:"marketTakeBound"`
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

// PercentPriceFilter define percent price filter of symbol
type PercentPriceFilter struct {
	MultiplierDecimal string `json:"multiplierDecimal"`
	MultiplierUp      string `json:"multiplierUp"`
	MultiplierDown    string `json:"multiplierDown"`
}

// MarketLotSizeFilter define market lot size filter of symbol
type MarketLotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// MaxNumOrdersFilter define max num orders filter of symbol
type MaxNumOrdersFilter struct {
	Limit int64 `json:"limit"`
}

// MaxNumAlgoOrdersFilter define max num algo orders filter of symbol
type MaxNumAlgoOrdersFilter struct {
	Limit int64 `json:"limit"`
}

// MinNotionalFilter define min notional filter of symbol
type MinNotionalFilter struct {
	Notional string `json:"notional"`
}

// LotSizeFilter return lot size filter of symbol
func (s *Symbol) LotSizeFilter() *LotSizeFilter {
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
func (s *Symbol) PriceFilter() *PriceFilter {
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

// PercentPriceFilter return percent price filter of symbol
func (s *Symbol) PercentPriceFilter() *PercentPriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePercentPrice) {
			f := &PercentPriceFilter{}
			if i, ok := filter["multiplierDecimal"]; ok {
				f.MultiplierDecimal = i.(string)
			}
			if i, ok := filter["multiplierUp"]; ok {
				f.MultiplierUp = i.(string)
			}
			if i, ok := filter["multiplierDown"]; ok {
				f.MultiplierDown = i.(string)
			}
			return f
		}
	}
	return nil
}

// MarketLotSizeFilter return market lot size filter of symbol
func (s *Symbol) MarketLotSizeFilter() *MarketLotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMarketLotSize) {
			f := &MarketLotSizeFilter{}
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

// MaxNumOrdersFilter return max num orders filter of symbol
func (s *Symbol) MaxNumOrdersFilter() *MaxNumOrdersFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMaxNumOrders) {
			f := &MaxNumOrdersFilter{}
			if i, ok := filter["limit"]; ok {
				if limit, okk := common.ToInt64(i); okk == nil {
					f.Limit = limit
				}
			}
			return f
		}
	}
	return nil
}

// MaxNumAlgoOrdersFilter return max num orders filter of symbol
func (s *Symbol) MaxNumAlgoOrdersFilter() *MaxNumAlgoOrdersFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMaxNumAlgoOrders) {
			f := &MaxNumAlgoOrdersFilter{}
			if i, ok := filter["limit"]; ok {
				if limit, okk := common.ToInt64(i); okk == nil {
					f.Limit = limit
				}
			}
			return f
		}
	}
	return nil
}

// MinNotionalFilter return min notional filter of symbol
func (s *Symbol) MinNotionalFilter() *MinNotionalFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMinNotional) {
			f := &MinNotionalFilter{}
			if i, ok := filter["notional"]; ok {
				f.Notional = i.(string)
			}
			return f
		}
	}
	return nil
}
