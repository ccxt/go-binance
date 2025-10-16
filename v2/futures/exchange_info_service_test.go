package futures

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
				"limit": 2400,
				"rateLimitType": "REQUEST_WEIGHT" 
			},
			{
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 1200,
				"rateLimitType": "ORDERS"
			}
		],
		"serverTime": 1565613908500, 
		"symbols": [
			{
				"symbol": "BLZUSDT",
				"pair": "BLZUSDT",
				"contractType": "PERPETUAL",
				"deliveryDate": 4133404800000,
				"onboardDate": 1598252400000,
				"status": "TRADING",
				"maintMarginPercent": "2.5000",   
				"requiredMarginPercent": "5.0000",  
				"baseAsset": "BLZ", 
				"quoteAsset": "USDT",
				"marginAsset": "USDT",
				"pricePrecision": 5,
				"quantityPrecision": 0,
				"baseAssetPrecision": 8,
				"quotePrecision": 8, 
				"underlyingType": "COIN",
				"underlyingSubType": ["STORAGE"],
				"settlePlan": 0,
				"triggerProtect": "0.15",
				"filters": [
					{
						"filterType": "PRICE_FILTER",
						"maxPrice": "300",
						"minPrice": "0.0001", 
						"tickSize": "0.0001"
					},
					{
						"filterType": "LOT_SIZE", 
						"maxQty": "10000000",
						"minQty": "1",
						"stepSize": "1"
					},
					{
						"filterType": "MARKET_LOT_SIZE",
						"maxQty": "590119",
						"minQty": "1",
						"stepSize": "1"
					},
					{
						"filterType": "MAX_NUM_ORDERS",
						"limit": 200
					},
					{
						"filterType": "MAX_NUM_ALGO_ORDERS",
						"limit": 100
					},
					{
						"notional": "5",
						"filterType": "MIN_NOTIONAL"
					},
					{
						"filterType": "PERCENT_PRICE",
						"multiplierUp": "1.1500",
						"multiplierDown": "0.8500",
						"multiplierDecimal": "4"
					}
				],
				"orderType": [
					"LIMIT",
					"MARKET",
					"STOP",
					"STOP_MARKET",
					"TAKE_PROFIT",
					"TAKE_PROFIT_MARKET",
					"TRAILING_STOP_MARKET" 
				],
				"timeInForce": [
					"GTC", 
					"IOC", 
					"FOK", 
					"GTX" 
				],
				"liquidationFee": "0.010000",
				"marketTakeBound": "0.30"
			}
		],
		"timezone": "UTC" 
	}
	
	`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, _, err := s.client.NewExchangeInfoService().Do(newContext())
	s.r().NoError(err)
	ei := &ExchangeInfo{
		Timezone:   "UTC",
		ServerTime: 1565613908500,
		RateLimits: []RateLimit{
			{RateLimitType: "REQUEST_WEIGHT", Interval: "MINUTE", IntervalNum: 1, Limit: 2400},
			{RateLimitType: "ORDERS", Interval: "MINUTE", IntervalNum: 1, Limit: 1200},
		},
		ExchangeFilters: []interface{}{},
		Symbols: []Symbol{
			{
				Symbol:                "BLZUSDT",
				Pair:                  "BLZUSDT",
				ContractType:          ContractTypePerpetual,
				DeliveryDate:          4133404800000,
				OnboardDate:           1598252400000,
				Status:                "TRADING",
				RequiredMarginPercent: "5.0000",
				MaintMarginPercent:    "2.5000",
				BaseAsset:             "BLZ",
				QuoteAsset:            "USDT",
				MarginAsset:           "USDT",
				PricePrecision:        5,
				QuantityPrecision:     0,
				BaseAssetPrecision:    8,
				QuotePrecision:        8,
				UnderlyingType:        "COIN",
				UnderlyingSubType:     []string{"STORAGE"},
				SettlePlan:            0,
				TriggerProtect:        "0.15",
				OrderType: []OrderType{OrderTypeLimit, OrderTypeMarket, OrderTypeStop, OrderTypeStopMarket,
					OrderTypeTakeProfit, OrderTypeTakeProfitMarket, OrderTypeTrailingStopMarket},
				TimeInForce: []TimeInForceType{TimeInForceTypeGTC, TimeInForceTypeIOC, TimeInForceTypeFOK, TimeInForceTypeGTX},
				Filters: []map[string]interface{}{
					{"filterType": "PRICE_FILTER", "minPrice": "0.0001", "maxPrice": "300", "tickSize": "0.0001"},
					{"filterType": "LOT_SIZE", "minQty": "1", "maxQty": "10000000", "stepSize": "1"},
					{"filterType": "MARKET_LOT_SIZE", "maxQty": "590119", "minQty": "1", "stepSize": "1"},
					{"filterType": "MAX_NUM_ORDERS", "limit": 200},
					{"filterType": "MAX_NUM_ALGO_ORDERS", "limit": 100},
					{"filterType": "MIN_NOTIONAL", "notional": "5"},
					{"filterType": "PERCENT_PRICE", "multiplierUp": "1.1500", "multiplierDown": "0.8500", "multiplierDecimal": "4"},
				},
				LiquidationFee:  "0.010000",
				MarketTakeBound: "0.30",
			},
		},
	}
	s.assertExchangeInfoEqual(ei, res)
	s.r().Len(ei.Symbols[0].Filters, 7, "Filters")

	ePriceFilter := &PriceFilter{
		MaxPrice: "300",
		MinPrice: "0.0001",
		TickSize: "0.0001",
	}
	s.assertPriceFilterEqual(ePriceFilter, res.Symbols[0].PriceFilter())

	eLotSizeFilter := &LotSizeFilter{
		MaxQuantity: "10000000",
		MinQuantity: "1",
		StepSize:    "1",
	}
	s.assertLotSizeFilterEqual(eLotSizeFilter, res.Symbols[0].LotSizeFilter())

	eMarketLotSizeFilter := &MarketLotSizeFilter{
		MaxQuantity: "590119",
		MinQuantity: "1",
		StepSize:    "1",
	}
	s.assertMarketLotSizeFilterEqual(eMarketLotSizeFilter, res.Symbols[0].MarketLotSizeFilter())

	eMaxNumOrdersFilter := &MaxNumOrdersFilter{
		Limit: 200,
	}
	s.assertMaxNumOrdersFilterEqual(eMaxNumOrdersFilter, res.Symbols[0].MaxNumOrdersFilter())

	eMaxNumAlgoOrdersFilter := &MaxNumAlgoOrdersFilter{
		Limit: 100,
	}
	s.assertMaxNumAlgoOrdersFilterEqual(eMaxNumAlgoOrdersFilter, res.Symbols[0].MaxNumAlgoOrdersFilter())

	eMinNotional := &MinNotionalFilter{
		Notional: "5",
	}
	s.assertMinNotionalFilterEqual(eMinNotional, res.Symbols[0].MinNotionalFilter())

	ePercentPriceFilter := &PercentPriceFilter{
		MultiplierDecimal: "4",
		MultiplierUp:      "1.1500",
		MultiplierDown:    "0.8500",
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
		r.Equal(e.Symbols[i].Pair, a.Symbols[i].Pair, "Pair")
		r.Equal(e.Symbols[i].ContractType, a.Symbols[i].ContractType, "ContractType")
		r.Equal(e.Symbols[i].DeliveryDate, a.Symbols[i].DeliveryDate, "DeliveryDate")
		r.Equal(e.Symbols[i].OnboardDate, a.Symbols[i].OnboardDate, "OnboardDate")
		r.Equal(e.Symbols[i].BaseAsset, a.Symbols[i].BaseAsset, "BaseAsset")
		r.Equal(e.Symbols[i].QuoteAsset, a.Symbols[i].QuoteAsset, "QuoteAsset")
		r.Equal(e.Symbols[i].MarginAsset, a.Symbols[i].MarginAsset, "MarginAsset")
		r.Equal(e.Symbols[i].Status, a.Symbols[i].Status, "Status")
		r.Equal(e.Symbols[i].MaintMarginPercent, a.Symbols[i].MaintMarginPercent, "MaintMarginPercent")
		r.Equal(e.Symbols[i].RequiredMarginPercent, a.Symbols[i].RequiredMarginPercent, "RequiredMarginPercent")
		r.Equal(e.Symbols[i].PricePrecision, a.Symbols[i].PricePrecision, "PricePrecision")
		r.Equal(e.Symbols[i].QuantityPrecision, a.Symbols[i].QuantityPrecision, "QuantityPrecision")
		r.Equal(e.Symbols[i].BaseAssetPrecision, a.Symbols[i].BaseAssetPrecision, "BaseAssetPrecision")
		r.Equal(e.Symbols[i].QuotePrecision, a.Symbols[i].QuotePrecision, "QuotePrecision")
		r.Equal(e.Symbols[i].UnderlyingType, a.Symbols[i].UnderlyingType, "UnderlyingType")
		r.Equal(e.Symbols[i].SettlePlan, a.Symbols[i].SettlePlan, "SettlePlan")
		r.Equal(e.Symbols[i].TriggerProtect, a.Symbols[i].TriggerProtect, "TriggerProtect")

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
			case "MIN_NOTIONAL":
				r.Equal(e.Symbols[i].MinNotionalFilter().Notional, a.Symbols[i].MinNotionalFilter().Notional, "Notional")
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
		r.Len(a.Symbols[i].UnderlyingSubType, len(e.Symbols[i].UnderlyingSubType), "UnderlyingSubType")
		for j, UnderlyingSubType := range e.Symbols[i].UnderlyingSubType {
			r.Equal(UnderlyingSubType, a.Symbols[i].UnderlyingSubType[j], "UnderlyingSubType")
		}
		r.Equal(e.Symbols[i].LiquidationFee, a.Symbols[i].LiquidationFee, "LiquidationFee")
		r.Equal(e.Symbols[i].MarketTakeBound, a.Symbols[i].MarketTakeBound, "MarketTakeBound")
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

func (s *exchangeInfoServiceTestSuite) assertMinNotionalFilterEqual(e, a *MinNotionalFilter) {
	r := s.r()
	r.Equal(e.Notional, a.Notional, "Notional")
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
