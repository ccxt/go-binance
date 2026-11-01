package binance

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
  		"timezone": "UTC",
  		"serverTime": 1539281238296,
  		"rateLimits": [
    		{
      			"rateLimitType": "REQUEST_WEIGHT",
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 1200
    		},
    		{
      			"rateLimitType": "ORDERS",
      			"interval": "SECOND",
				"intervalNum": 10,
      			"limit": 10
    		},
    		{
      			"rateLimitType": "ORDERS",
      			"interval": "DAY",
				"intervalNum": 1,
      			"limit": 100000
    		}
  		],
		"exchangeFilters": [],
		"symbols": [
			{
				"symbol": "ETHBTC",
				"status": "TRADING",
				"baseAsset": "ETH",
				"baseAssetPrecision": 8,
				"quoteAsset": "BTC",
				"quotePrecision": 8,
				"orderTypes":["LIMIT","LIMIT_MAKER","MARKET","STOP_LOSS_LIMIT","TAKE_PROFIT_LIMIT"],
				"icebergAllowed": true,
				"ocoAllowed": true,
				"isSpotTradingAllowed": true,
				"isMarginTradingAllowed": false,
				"filters":[{
					"filterType": "PRICE_FILTER",
					"minPrice": "0.00001000",
					"maxPrice": "922327.00000000",
					"tickSize": "0.00001000"
				}, {
					"filterType": "LOT_SIZE",
					"minQty": "0.00010000",
					"maxQty": "100000.00000000",
					"stepSize": "0.00010000"
				}, {
					"filterType": "ICEBERG_PARTS",
					"limit": 10
				}, {
					"filterType": "MARKET_LOT_SIZE",
					"minQty": "0.00000000",
					"maxQty": "3373.58569748",
					"stepSize": "0.00000000"
				}, {
					"filterType": "TRAILING_DELTA",
					"minTrailingAboveDelta": 10,
					"maxTrailingAboveDelta": 2000,
					"minTrailingBelowDelta": 10,
					"maxTrailingBelowDelta": 2000
				}, {
					"filterType": "PERCENT_PRICE_BY_SIDE",
					"bidMultiplierUp": "5",
					"bidMultiplierDown": "0.2",
					"askMultiplierUp": "5",
					"askMultiplierDown": "0.2",
					"avgPriceMins": 5
				}, {
					"filterType": "NOTIONAL",
					"minNotional": "0.00010000",
					"applyMinToMarket": true,
					"maxNotional": "9000000.00000000",
					"applyMaxToMarket": false,
					"avgPriceMins": 5
				}, {
					"filterType": "MAX_NUM_ORDERS",
					"maxNumOrders": 200
				}, {
					"filterType": "MAX_NUM_ALGO_ORDERS",
					"maxNumAlgoOrders": 5
				}],
				"permissions": ["SPOT","MARGIN"]
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "ETHBTC"
	symbols := []string{"ETHBTC", "LTCBTC"}
	permissions := []string{"SPOT", "MARGIN"}
	s.assertReq(func(r *request) {
		e := newRequest().setParams(map[string]any{
			"symbol":      symbol,
			"symbols":     `["ETHBTC","LTCBTC"]`,
			"permissions": `["SPOT","MARGIN"]`,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewExchangeInfoService().Symbol(symbol).Symbols(symbols...).Permissions(permissions...).Do(newContext())
	s.r().NoError(err)
	ei := &ExchangeInfo{
		Timezone:   "UTC",
		ServerTime: 1539281238296,
		RateLimits: []RateLimit{
			{RateLimitType: "REQUEST_WEIGHT", Interval: "MINUTE", IntervalNum: 1, Limit: 1200},
			{RateLimitType: "ORDERS", Interval: "SECOND", IntervalNum: 10, Limit: 10},
			{RateLimitType: "ORDERS", Interval: "DAY", IntervalNum: 1, Limit: 100000},
		},
		ExchangeFilters: []any{},
		Symbols: []Symbol{
			{
				Symbol:                 "ETHBTC",
				Status:                 "TRADING",
				BaseAsset:              "ETH",
				BaseAssetPrecision:     8,
				QuoteAsset:             "BTC",
				QuotePrecision:         8,
				OrderTypes:             []string{"LIMIT", "LIMIT_MAKER", "MARKET", "STOP_LOSS_LIMIT", "TAKE_PROFIT_LIMIT"},
				IcebergAllowed:         true,
				OcoAllowed:             true,
				IsSpotTradingAllowed:   true,
				IsMarginTradingAllowed: false,
				Filters: []map[string]any{
					{"filterType": "PRICE_FILTER", "minPrice": "0.00001000", "maxPrice": "922327.00000000", "tickSize": "0.00001000"},
					{"filterType": "LOT_SIZE", "minQty": "0.00010000", "maxQty": "100000.00000000", "stepSize": "0.00010000"},
					{"filterType": "ICEBERG_PARTS", "limit": 10},
					{"filterType": "MARKET_LOT_SIZE", "minQty": "0.00000000", "maxQty": "3373.58569748", "stepSize": "0.00000000"},
					{"filterType": "TRAILING_DELTA", "minTrailingAboveDelta": 10, "maxTrailingAboveDelta": 2000, "minTrailingBelowDelta": 10, "maxTrailingBelowDelta": 2000},
					{"filterType": "PERCENT_PRICE_BY_SIDE", "bidMultiplierUp": "5", "bidMultiplierDown": "0.2", "askMultiplierUp": "5", "askMultiplierDown": "0.2", "avgPriceMins": 5},
					{"filterType": "NOTIONAL", "minNotional": "0.00010000", "applyMinToMarket": true, "maxNotional": "9000000.00000000", "applyMaxToMarket": false, "avgPriceMins": 5},
					{"filterType": "MAX_NUM_ORDERS", "maxNumOrders": 200},
					{"filterType": "MAX_NUM_ALGO_ORDERS", "maxNumAlgoOrders": 5},
				},
				Permissions: []string{"SPOT", "MARGIN"},
			},
		},
	}
	s.assertExchangeInfoEqual(ei, res)

	ePriceFilter := &PriceFilter{
		MaxPrice: "922327.00000000",
		MinPrice: "0.00001000",
		TickSize: "0.00001000",
	}
	s.assertPriceFilterEqual(ePriceFilter, res.Symbols[0].PriceFilter())

	eLotSizeFilter := &LotSizeFilter{
		MaxQuantity: "100000.00000000",
		MinQuantity: "0.00010000",
		StepSize:    "0.00010000",
	}
	s.assertLotSizeFilterEqual(eLotSizeFilter, res.Symbols[0].LotSizeFilter())

	eIcebergPartsFilter := &IcebergPartsFilter{
		Limit: 10,
	}
	s.assertIcebergPartsFilterEqual(eIcebergPartsFilter, res.Symbols[0].IcebergPartsFilter())

	eMarketLotSizeFilter := &MarketLotSizeFilter{
		MinQuantity: "0.00000000",
		MaxQuantity: "3373.58569748",
		StepSize:    "0.00000000",
	}
	s.assertMarketLotSizeFilterEqual(eMarketLotSizeFilter, res.Symbols[0].MarketLotSizeFilter())

	eTrailingDeltaFilter := &TrailingDeltaFilter{
		MinTrailingAboveDelta: 10,
		MaxTrailingAboveDelta: 2000,
		MinTrailingBelowDelta: 10,
		MaxTrailingBelowDelta: 2000,
	}
	s.assertTrailingDeltaFilterEqual(eTrailingDeltaFilter, res.Symbols[0].TrailingDeltaFilter())

	ePercentPriceBySideFilter := &PercentPriceBySideFilter{
		BidMultiplierUp:   "5",
		BidMultiplierDown: "0.2",
		AskMultiplierUp:   "5",
		AskMultiplierDown: "0.2",
		AveragePriceMins:  5,
	}
	s.assertPercentPriceBySideFilterEqual(ePercentPriceBySideFilter, res.Symbols[0].PercentPriceBySideFilter())

	eMinNotionalFilter := &NotionalFilter{
		MinNotional:      "0.00010000",
		ApplyMinToMarket: true,
		MaxNotional:      "9000000.00000000",
		ApplyMaxToMarket: false,
		AvgPriceMins:     5,
	}
	s.assertMinNotionalFilterEqual(eMinNotionalFilter, res.Symbols[0].NotionalFilter())

	eMaxNumOrdersFilter := &MaxNumOrdersFilter{
		MaxNumOrders: 200,
	}
	s.assertMaxNumOrdersFilterEqual(eMaxNumOrdersFilter, res.Symbols[0].MaxNumOrdersFilter())

	eMaxNumAlgoOrdersFilter := &MaxNumAlgoOrdersFilter{
		MaxNumAlgoOrders: 5,
	}
	s.assertMaxNumAlgoOrdersFilterEqual(eMaxNumAlgoOrdersFilter, res.Symbols[0].MaxNumAlgoOrdersFilter())
}

func (s *exchangeInfoServiceTestSuite) TestExchangeInfoWithPermissionSets() {
	data := []byte(`{
		"timezone":"UTC",
		"serverTime":1733252001653,
		"rateLimits":[
			{
				"rateLimitType":"REQUEST_WEIGHT",
				"interval":"MINUTE",
				"intervalNum":1,
				"limit":6000
			}
		],
		"exchangeFilters":[],
		"symbols":[
			{
				"symbol":"BTCUSDT",
				"status":"TRADING",
				"baseAsset":"BTC",
				"baseAssetPrecision":8,
				"quoteAsset":"USDT",
				"quotePrecision":8,
				"quoteAssetPrecision":8,
				"baseCommissionPrecision":8,
				"quoteCommissionPrecision":8,
				"orderTypes":["LIMIT","LIMIT_MAKER","MARKET","STOP_LOSS","STOP_LOSS_LIMIT","TAKE_PROFIT","TAKE_PROFIT_LIMIT"],
				"icebergAllowed":true,
				"ocoAllowed":true,
				"otoAllowed":true,
				"quoteOrderQtyMarketAllowed":true,
				"allowTrailingStop":true,
				"cancelReplaceAllowed":true,
				"isSpotTradingAllowed":true,
				"isMarginTradingAllowed":true,
				"filters":[
					{
						"filterType":"PRICE_FILTER",
						"minPrice":"0.01000000",
						"maxPrice":"1000000.00000000",
						"tickSize":"0.01000000"
					}
				],
				"permissions":[],
				"permissionSets":[["SPOT","MARGIN"]]
			}
		]
	}`)
	s.mockDo(data, nil)
	res, err := s.client.NewExchangeInfoService().Symbol("BTCUSDT").Do(newContext())
	s.r().NoError(err)
	s.r().Len(res.Symbols[0].PermissionSets, 1)
	s.r().Equal([]string{"SPOT", "MARGIN"}, res.Symbols[0].PermissionSets[0])
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()

	r.Equal(e.Timezone, a.Timezone, "Timezone")
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")

	for i := range a.RateLimits {
		r.Equal(e.RateLimits[i].RateLimitType, a.RateLimits[i].RateLimitType, "RateLimitType")
		r.Equal(e.RateLimits[i].Limit, a.RateLimits[i].Limit, "Limit")
		r.Equal(e.RateLimits[i].Interval, a.RateLimits[i].Interval, "Interval")
		r.Equal(e.RateLimits[i].IntervalNum, a.RateLimits[i].IntervalNum, "IntervalNum")
	}

	r.Equal(e.ExchangeFilters, a.ExchangeFilters, "ExchangeFilters")

	for i, currentSymbol := range a.Symbols {
		if a.Symbols[i].Symbol == e.Symbols[0].Symbol {
			r.Equal(e.Symbols[i].Status, currentSymbol.Status, "Status")
			r.Equal(e.Symbols[i].BaseAsset, currentSymbol.BaseAsset, "BaseAsset")
			r.Equal(e.Symbols[i].BaseAssetPrecision, currentSymbol.BaseAssetPrecision, "BaseAssetPrecision")
			r.Equal(e.Symbols[i].QuoteAsset, currentSymbol.QuoteAsset, "QuoteAsset")
			r.Equal(e.Symbols[i].QuotePrecision, currentSymbol.QuotePrecision, "QuotePrecision")
			r.Len(currentSymbol.OrderTypes, len(e.Symbols[i].OrderTypes))
			r.Equal(e.Symbols[i].OrderTypes, currentSymbol.OrderTypes, "OrderTypes")
			r.Equal(e.Symbols[i].IcebergAllowed, currentSymbol.IcebergAllowed, "IcebergAllowed")
			r.Len(currentSymbol.Filters, len(e.Symbols[i].Filters))

			for fi, currentFilter := range currentSymbol.Filters {
				r.Len(currentFilter, len(e.Symbols[i].Filters[fi]))
				switch currentFilter["filterType"] {
				case "PRICE_FILTER":
					r.Equal(e.Symbols[i].PriceFilter().MinPrice, currentSymbol.PriceFilter().MinPrice, "MinPrice")
					r.Equal(e.Symbols[i].PriceFilter().MaxPrice, currentSymbol.PriceFilter().MaxPrice, "MaxPrice")
					r.Equal(e.Symbols[i].PriceFilter().TickSize, currentSymbol.PriceFilter().TickSize, "TickSize")
				case "LOT_SIZE":
					r.Equal(e.Symbols[i].LotSizeFilter().MinQuantity, currentSymbol.LotSizeFilter().MinQuantity, "MinQuantity")
					r.Equal(e.Symbols[i].LotSizeFilter().MaxQuantity, currentSymbol.LotSizeFilter().MaxQuantity, "MaxQuantity")
					r.Equal(e.Symbols[i].LotSizeFilter().StepSize, currentSymbol.LotSizeFilter().StepSize, "StepSize")
				case "ICEBERG_PARTS":
					r.Equal(e.Symbols[i].IcebergPartsFilter().Limit, currentSymbol.IcebergPartsFilter().Limit, "Limit")
				case "MARKET_LOT_SIZE":
					r.Equal(e.Symbols[i].MarketLotSizeFilter().MinQuantity, currentSymbol.MarketLotSizeFilter().MinQuantity, "MinQuantity")
					r.Equal(e.Symbols[i].MarketLotSizeFilter().MaxQuantity, currentSymbol.MarketLotSizeFilter().MaxQuantity, "MaxQuantity")
					r.Equal(e.Symbols[i].MarketLotSizeFilter().StepSize, currentSymbol.MarketLotSizeFilter().StepSize, "StepSize")
				case "TRAILING_DELTA":
					r.Equal(e.Symbols[i].TrailingDeltaFilter().MinTrailingAboveDelta, currentSymbol.TrailingDeltaFilter().MinTrailingAboveDelta, "MinTrailingAboveDelta")
					r.Equal(e.Symbols[i].TrailingDeltaFilter().MaxTrailingAboveDelta, currentSymbol.TrailingDeltaFilter().MaxTrailingAboveDelta, "MaxTrailingAboveDelta")
					r.Equal(e.Symbols[i].TrailingDeltaFilter().MinTrailingBelowDelta, currentSymbol.TrailingDeltaFilter().MinTrailingBelowDelta, "MinTrailingBelowDelta")
					r.Equal(e.Symbols[i].TrailingDeltaFilter().MaxTrailingBelowDelta, currentSymbol.TrailingDeltaFilter().MaxTrailingBelowDelta, "MaxTrailingBelowDelta")
				case "PERCENT_PRICE_BY_SIDE":
					r.Equal(e.Symbols[i].PercentPriceBySideFilter().BidMultiplierUp, currentSymbol.PercentPriceBySideFilter().BidMultiplierUp, "BidMultiplierUp")
					r.Equal(e.Symbols[i].PercentPriceBySideFilter().BidMultiplierDown, currentSymbol.PercentPriceBySideFilter().BidMultiplierDown, "BidMultiplierDown")
					r.Equal(e.Symbols[i].PercentPriceBySideFilter().AskMultiplierUp, currentSymbol.PercentPriceBySideFilter().AskMultiplierUp, "AskMultiplierUp")
					r.Equal(e.Symbols[i].PercentPriceBySideFilter().AskMultiplierDown, currentSymbol.PercentPriceBySideFilter().AskMultiplierDown, "AskMultiplierDown")
					r.Equal(e.Symbols[i].PercentPriceBySideFilter().AveragePriceMins, currentSymbol.PercentPriceBySideFilter().AveragePriceMins, "AveragePriceMins")
				case "NOTIONAL":
					r.Equal(e.Symbols[i].NotionalFilter().MinNotional, currentSymbol.NotionalFilter().MinNotional, "MinNotional")
					r.Equal(e.Symbols[i].NotionalFilter().ApplyMinToMarket, currentSymbol.NotionalFilter().ApplyMinToMarket, "ApplyMinToMarket")
					r.Equal(e.Symbols[i].NotionalFilter().MaxNotional, currentSymbol.NotionalFilter().MaxNotional, "MaxNotional")
					r.Equal(e.Symbols[i].NotionalFilter().ApplyMaxToMarket, currentSymbol.NotionalFilter().ApplyMaxToMarket, "ApplyMaxToMarket")
					r.Equal(e.Symbols[i].NotionalFilter().AvgPriceMins, currentSymbol.NotionalFilter().AvgPriceMins, "AvgPriceMins")
				case "MAX_NUM_ORDERS":
					r.Equal(e.Symbols[i].MaxNumOrdersFilter().MaxNumOrders, currentSymbol.MaxNumOrdersFilter().MaxNumOrders, "MaxNumOrders")
				case "MAX_NUM_ALGO_ORDERS":
					r.Equal(e.Symbols[i].MaxNumAlgoOrdersFilter().MaxNumAlgoOrders, currentSymbol.MaxNumAlgoOrdersFilter().MaxNumAlgoOrders, "MaxNumAlgoOrders")
				}
			}

			r.Len(currentSymbol.Permissions, len(e.Symbols[i].Permissions))
			r.Equal(e.Symbols[i].Permissions, currentSymbol.Permissions, "Permissions")

			return
		}

	}
	r.Fail("Symbol ETHBTC not found")
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

func (s *exchangeInfoServiceTestSuite) assertPercentPriceBySideFilterEqual(e, a *PercentPriceBySideFilter) {
	r := s.r()
	r.Equal(e.AveragePriceMins, a.AveragePriceMins, "AveragePriceMins")
	r.Equal(e.BidMultiplierUp, a.BidMultiplierUp, "BidMultiplierUp")
	r.Equal(e.BidMultiplierDown, a.BidMultiplierDown, "BidMultiplierDown")
	r.Equal(e.AskMultiplierUp, a.AskMultiplierUp, "AskMultiplierUp")
	r.Equal(e.AskMultiplierDown, a.AskMultiplierDown, "AskMultiplierDown")
}

func (s *exchangeInfoServiceTestSuite) assertMinNotionalFilterEqual(e, a *NotionalFilter) {
	r := s.r()
	r.Equal(e.MinNotional, a.MinNotional, "MinNotional")
	r.Equal(e.ApplyMinToMarket, a.ApplyMinToMarket, "ApplyMinToMarket")
	r.Equal(e.MaxNotional, a.MaxNotional, "MaxNotional")
	r.Equal(e.ApplyMaxToMarket, a.ApplyMaxToMarket, "ApplyMaxToMarket")
	r.Equal(e.AvgPriceMins, a.AvgPriceMins, "AvgPriceMins")
}

func (s *exchangeInfoServiceTestSuite) assertIcebergPartsFilterEqual(e, a *IcebergPartsFilter) {
	r := s.r()
	r.Equal(e.Limit, a.Limit, "Limit")
}

func (s *exchangeInfoServiceTestSuite) assertMarketLotSizeFilterEqual(e, a *MarketLotSizeFilter) {
	r := s.r()
	r.Equal(e.MaxQuantity, a.MaxQuantity, "MaxQuantity")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.StepSize, a.StepSize, "StepSize")
}

func (s *exchangeInfoServiceTestSuite) assertTrailingDeltaFilterEqual(e, a *TrailingDeltaFilter) {
	r := s.r()
	r.Equal(e.MinTrailingAboveDelta, a.MinTrailingAboveDelta, "MinTrailingAboveDelta")
	r.Equal(e.MaxTrailingAboveDelta, a.MaxTrailingAboveDelta, "MaxTrailingAboveDelta")
	r.Equal(e.MinTrailingBelowDelta, a.MinTrailingBelowDelta, "MinTrailingBelowDelta")
	r.Equal(e.MaxTrailingBelowDelta, a.MaxTrailingBelowDelta, "MaxTrailingBelowDelta")
}

func (s *exchangeInfoServiceTestSuite) assertMaxNumOrdersFilterEqual(e, a *MaxNumOrdersFilter) {
	r := s.r()
	r.Equal(e.MaxNumOrders, a.MaxNumOrders, "MaxNumOrders")
}

func (s *exchangeInfoServiceTestSuite) assertMaxNumAlgoOrdersFilterEqual(e, a *MaxNumAlgoOrdersFilter) {
	r := s.r()
	r.Equal(e.MaxNumAlgoOrders, a.MaxNumAlgoOrders, "MaxNumAlgoOrders")
}
