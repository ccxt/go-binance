package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type exchangeInfoServiceTestSuite struct {
	baseTestSuite
}

func TestExchangeInfoService(t *testing.T) {
	suite.Run(t, new(exchangeInfoServiceTestSuite))
}

func (s *exchangeInfoServiceTestSuite) TestExchangeInfo() {
	data := []byte(`{
		"exchangeFilters": [],
		"rateLimits": [ 
			{
				"interval": "MINUTE", 
				"intervalNum": 1, 
				"limit": 6000, 
				"rateLimitType": "REQUEST_WEIGHT" 
			},
			{
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 6000,
				"rateLimitType": "ORDERS"
			}
		],
		"serverTime": 1565613908500,
		"symbols": [
			{
				"filters": [
					{
						"filterType": "PRICE_FILTER", 
						"maxPrice": "100000", 
						"minPrice": "0.1", 
						"tickSize": "0.1" 
					},
					{
						"filterType": "LOT_SIZE", 
						"maxQty": "100000", 
						"minQty": "1", 
						"stepSize": "1" 
					},
					{
						"filterType": "MARKET_LOT_SIZE", 
						"maxQty": "100000", 
						"minQty": "1", 
						"stepSize": "1" 
					},
					{
						"filterType": "MAX_NUM_ORDERS", 
						"limit": 200
					},
					{
						"limit": 20,
						"filterType": "MAX_NUM_ALGO_ORDERS"
					},
					{
						"filterType": "PERCENT_PRICE", 
						"multiplierUp": "1.0500", 
						"multiplierDown": "0.9500", 
						"multiplierDecimal": "4"
					}
				],
				"OrderType": [ 
					"LIMIT", 
					"MARKET", 
					"STOP",
					"TAKE_PROFIT",
					"TRAILING_STOP_MARKET"
				],
				"timeInForce": [
					"GTC",
					"IOC",
					"FOK",
					"GTX"
				],
				"symbol": "BTCUSD_200925",
				"pair": "BTCUSD",
				"contractType": "CURRENT_QUARTER", 
				"deliveryDate": 1601020800000,
				"onboardDate": 1590739200000,
				"contractStatus": "TRADING", 
				"contractSize": 100,    
				"quoteAsset": "USD",
				"baseAsset": "BTC",   
				"marginAsset": "BTC",
				"pricePrecision": 1,
				"quantityPrecision": 0,
				"baseAssetPrecision": 8,
				"quotePrecision": 8,
				"equalQtyPrecision": 4,
				"triggerProtect": "0.0500",
				"maintMarginPercent": "2.5000",
				"requiredMarginPercent": "5.0000",
				"underlyingType": "COIN", 
				"underlyingSubType": [] 
			}
		],
		"timezone": "UTC"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewExchangeInfoService().Do(newContext())
	s.r().NoError(err)
	ei := &ExchangeInfo{
		Timezone:   "UTC",
		ServerTime: 1565613908500,
		RateLimits: []RateLimit{
			{RateLimitType: "REQUEST_WEIGHT", Interval: "MINUTE", IntervalNum: 1, Limit: 6000},
			{RateLimitType: "ORDERS", Interval: "MINUTE", IntervalNum: 1, Limit: 6000},
		},
		ExchangeFilters: []any{},
		Symbols: []Symbol{
			{
				Symbol:                "BTCUSD_200925",
				Pair:                  "BTCUSD",
				ContractType:          "CURRENT_QUARTER",
				DeliveryDate:          1601020800000,
				OnboardDate:           1590739200000,
				ContractStatus:        "TRADING",
				ContractSize:          100,
				QuoteAsset:            "USD",
				BaseAsset:             "BTC",
				MarginAsset:           "BTC",
				PricePrecision:        1,
				QuantityPrecision:     0,
				BaseAssetPrecision:    8,
				QuotePrecision:        8,
				EqualQtyPrecision:     4,
				TriggerProtect:        "0.0500",
				MaintMarginPercent:    "2.5000",
				RequiredMarginPercent: "5.0000",
				UnderlyingType:        "COIN",
				UnderlyingSubType:     []any{},
				OrderType:             []OrderType{OrderTypeLimit, OrderTypeMarket, OrderTypeStop, OrderTypeTakeProfit, OrderTypeTrailingStopMarket},
				TimeInForce:           []TimeInForceType{TimeInForceTypeGTC, TimeInForceTypeIOC, TimeInForceTypeFOK, TimeInForceTypeGTX},
				Filters: []map[string]any{
					{"filterType": "PRICE_FILTER", "minPrice": "0.1", "maxPrice": "100000", "tickSize": "0.1"},
					{"filterType": "LOT_SIZE", "minQty": "1", "maxQty": "100000", "stepSize": "1"},
					{"filterType": "MARKET_LOT_SIZE", "maxQty": "100000", "minQty": "1", "stepSize": "1"},
					{"filterType": "MAX_NUM_ORDERS", "limit": 200},
					{"filterType": "MAX_NUM_ALGO_ORDERS", "limit": 20},
					{"filterType": "PERCENT_PRICE", "multiplierUp": "1.0500", "multiplierDown": "0.9500", "multiplierDecimal": "4"},
				},
			},
		},
	}
	s.assertExchangeInfoEqual(ei, res)
	s.r().Len(ei.Symbols[0].Filters, 6, "Filters")

	ePriceFilter := &PriceFilter{
		MaxPrice: "100000",
		MinPrice: "0.1",
		TickSize: "0.1",
	}
	s.assertPriceFilterEqual(ePriceFilter, res.Symbols[0].PriceFilter())

	eLotSizeFilter := &LotSizeFilter{
		MaxQuantity: "100000",
		MinQuantity: "1",
		StepSize:    "1",
	}
	s.assertLotSizeFilterEqual(eLotSizeFilter, res.Symbols[0].LotSizeFilter())

	eMarketLotSizeFilter := &MarketLotSizeFilter{
		MaxQuantity: "100000",
		MinQuantity: "1",
		StepSize:    "1",
	}
	s.assertMarketLotSizeFilterEqual(eMarketLotSizeFilter, res.Symbols[0].MarketLotSizeFilter())

	eMaxNumOrdersFilter := &MaxNumOrdersFilter{
		Limit: 200,
	}
	s.assertMaxNumOrdersFilterEqual(eMaxNumOrdersFilter, res.Symbols[0].MaxNumOrdersFilter())

	eMaxNumAlgoOrdersFilter := &MaxNumAlgoOrdersFilter{
		Limit: 20,
	}
	s.assertMaxNumAlgoOrdersFilterEqual(eMaxNumAlgoOrdersFilter, res.Symbols[0].MaxNumAlgoOrdersFilter())

	ePercentPriceFilter := &PercentPriceFilter{
		MultiplierDecimal: "4",
		MultiplierUp:      "1.0500",
		MultiplierDown:    "0.9500",
	}
	s.assertPercentPriceFilterEqual(ePercentPriceFilter, res.Symbols[0].PercentPriceFilter())
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()

	r.Equal(e.Timezone, a.Timezone, "Timezone")
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")
	r.Len(a.RateLimits, len(e.RateLimits), "RateLimits")
	for i := range a.RateLimits {
		r.Equal(e.RateLimits[i].RateLimitType, a.RateLimits[i].RateLimitType, "RateLimitType")
		r.Equal(e.RateLimits[i].Limit, a.RateLimits[i].Limit, "Limit")
		r.Equal(e.RateLimits[i].Interval, a.RateLimits[i].Interval, "Interval")
		r.Equal(e.RateLimits[i].IntervalNum, a.RateLimits[i].IntervalNum, "IntervalNum")
	}
	r.Equal(e.ExchangeFilters, a.ExchangeFilters, "ExchangeFilters")
	r.Len(a.Symbols, len(e.Symbols), "Symbols")

	for i := range a.Symbols {
		r.Equal(e.Symbols[i].Symbol, a.Symbols[i].Symbol, "Symbol")
		r.Equal(e.Symbols[i].BaseAsset, a.Symbols[i].BaseAsset, "BaseAsset")
		r.Equal(e.Symbols[i].QuoteAsset, a.Symbols[i].QuoteAsset, "QuoteAsset")
		r.Equal(e.Symbols[i].Pair, a.Symbols[i].Pair, "Pair")
		r.Equal(e.Symbols[i].ContractType, a.Symbols[i].ContractType, "ContractType")
		r.Equal(e.Symbols[i].DeliveryDate, a.Symbols[i].DeliveryDate, "DeliveryDate")
		r.Equal(e.Symbols[i].OnboardDate, a.Symbols[i].OnboardDate, "OnboardDate")
		r.Equal(e.Symbols[i].ContractStatus, a.Symbols[i].ContractStatus, "ContractStatus")
		r.Equal(e.Symbols[i].ContractSize, a.Symbols[i].ContractSize, "ContractSize")
		r.Equal(e.Symbols[i].MarginAsset, a.Symbols[i].MarginAsset, "MarginAsset")
		r.Equal(e.Symbols[i].BaseAssetPrecision, a.Symbols[i].BaseAssetPrecision, "BaseAssetPrecision")
		r.Equal(e.Symbols[i].QuotePrecision, a.Symbols[i].QuotePrecision, "QuotePrecision")
		r.Equal(e.Symbols[i].EqualQtyPrecision, a.Symbols[i].EqualQtyPrecision, "EqualQtyPrecision")
		r.Equal(e.Symbols[i].TriggerProtect, a.Symbols[i].TriggerProtect, "TriggerProtect")
		r.Equal(e.Symbols[i].UnderlyingType, a.Symbols[i].UnderlyingType, "UnderlyingType")
		r.Equal(e.Symbols[i].UnderlyingSubType, a.Symbols[i].UnderlyingSubType, "UnderlyingSubType")
		r.Equal(e.Symbols[i].MaintMarginPercent, a.Symbols[i].MaintMarginPercent, "MaintMarginPercent")
		r.Equal(e.Symbols[i].PricePrecision, a.Symbols[i].PricePrecision, "PricePrecision")
		r.Equal(e.Symbols[i].QuantityPrecision, a.Symbols[i].QuantityPrecision, "QuantityPrecision")
		r.Equal(e.Symbols[i].RequiredMarginPercent, a.Symbols[i].RequiredMarginPercent, "RequiredMarginPercent")

		for fi, currentFilter := range a.Symbols[i].Filters {
			r.Len(currentFilter, len(e.Symbols[i].Filters[fi]))
			switch currentFilter["filterType"] {
			case "PRICE_FILTER":
				r.Equal(e.Symbols[i].PriceFilter().MinPrice, a.Symbols[i].PriceFilter().MinPrice, "MinPrice")
				r.Equal(e.Symbols[i].PriceFilter().MaxPrice, a.Symbols[i].PriceFilter().MaxPrice, "MaxPrice")
				r.Equal(e.Symbols[i].PriceFilter().TickSize, a.Symbols[i].PriceFilter().TickSize, "TickSize")
			case "LOT_SIZE":
				r.Equal(e.Symbols[i].LotSizeFilter().MinQuantity, a.Symbols[i].LotSizeFilter().MinQuantity, "MinQuantity")
				r.Equal(e.Symbols[i].LotSizeFilter().MaxQuantity, a.Symbols[i].LotSizeFilter().MaxQuantity, "MaxQuantity")
				r.Equal(e.Symbols[i].LotSizeFilter().StepSize, a.Symbols[i].LotSizeFilter().StepSize, "StepSize")
			case "MARKET_LOT_SIZE":
				r.Equal(e.Symbols[i].MarketLotSizeFilter().MinQuantity, a.Symbols[i].MarketLotSizeFilter().MinQuantity, "MinQuantity")
				r.Equal(e.Symbols[i].MarketLotSizeFilter().MaxQuantity, a.Symbols[i].MarketLotSizeFilter().MaxQuantity, "MaxQuantity")
				r.Equal(e.Symbols[i].MarketLotSizeFilter().StepSize, a.Symbols[i].MarketLotSizeFilter().StepSize, "StepSize")
			case "MAX_NUM_ORDERS":
				r.Equal(e.Symbols[i].MaxNumOrdersFilter().Limit, a.Symbols[i].MaxNumOrdersFilter().Limit, "Limit")
			case "MAX_NUM_ALGO_ORDERS":
				r.Equal(e.Symbols[i].MaxNumAlgoOrdersFilter().Limit, a.Symbols[i].MaxNumAlgoOrdersFilter().Limit, "Limit")
			case "PERCENT_PRICE":
				r.Equal(e.Symbols[i].PercentPriceFilter().MultiplierDecimal, a.Symbols[i].PercentPriceFilter().MultiplierDecimal, "MultiplierDecimal")
				r.Equal(e.Symbols[i].PercentPriceFilter().MultiplierUp, a.Symbols[i].PercentPriceFilter().MultiplierUp, "MultiplierUp")
				r.Equal(e.Symbols[i].PercentPriceFilter().MultiplierDown, a.Symbols[i].PercentPriceFilter().MultiplierDown, "MultiplierDown")
			}
		}

		r.Len(a.Symbols[i].OrderType, len(e.Symbols[i].OrderType))
		for j, orderType := range e.Symbols[i].OrderType {
			r.Equal(orderType, a.Symbols[i].OrderType[j], "OrderType")
		}
		r.Len(a.Symbols[i].TimeInForce, len(e.Symbols[i].TimeInForce), "TimeInForce")
		for j, timeInForce := range e.Symbols[i].TimeInForce {
			r.Equal(timeInForce, a.Symbols[i].TimeInForce[j], "TimeInForce")
		}
	}
}

func (s *exchangeInfoServiceTestSuite) assertLotSizeFilterEqual(e, a *LotSizeFilter) {
	r := s.r()
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertPriceFilterEqual(e, a *PriceFilter) {
	r := s.r()
	r.Equal(e.MaxPrice, a.MaxPrice, "MaxPrice")
	r.Equal(e.MinPrice, a.MinPrice, "MinPrice")
	r.Equal(e.TickSize, a.TickSize, "TickSize")
}

func (s *exchangeInfoServiceTestSuite) assertPercentPriceFilterEqual(e, a *PercentPriceFilter) {
	r := s.r()
	r.Equal(e.MultiplierDecimal, a.MultiplierDecimal, "MultiplierDecimal")
	r.Equal(e.MultiplierUp, a.MultiplierUp, "MultiplierUp")
	r.Equal(e.MultiplierDown, a.MultiplierDown, "MultiplierDown")
}

func (s *exchangeInfoServiceTestSuite) assertMarketLotSizeFilterEqual(e, a *MarketLotSizeFilter) {
	r := s.r()
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertMaxNumOrdersFilterEqual(e, a *MaxNumOrdersFilter) {
	r := s.r()
	r.Equal(e.Limit, a.Limit, "Limit")
}

func (s *exchangeInfoServiceTestSuite) assertMaxNumAlgoOrdersFilterEqual(e, a *MaxNumAlgoOrdersFilter) {
	r := s.r()
	r.Equal(e.Limit, a.Limit, "Limit")
}
