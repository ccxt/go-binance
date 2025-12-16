package options

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
		"serverTime": 1592387337630,
		"optionContracts": [
		  {
			"id": 1,
			"baseAsset": "BTC",
			"quoteAsset": "USDT",
			"underlying": "BTCUSDT",
			"settleAsset": "USDT"
		  }
		],
		"optionAssets": [
		  {
			"id": 1,
			"name": "USDT"
		  }
		],
		"optionSymbols": [
		  {
			  "contractId": 2,
			  "expiryDate": 1660521600000,
			  "filters": [
				  {
					  "filterType": "PRICE_FILTER",
					  "minPrice": "0.02",
					  "maxPrice": "80000.01",
					  "tickSize": "0.01"
				  },
				  {
					  "filterType": "LOT_SIZE",
					  "minQty": "0.01",
					  "maxQty": "100",
					  "stepSize": "0.01"
				  }
			  ],
			  "id": 17,
			  "symbol": "BTC-220815-50000-C",
			  "side": "CALL",
			  "strikePrice": "50000",
			  "underlying": "BTCUSDT",
			  "unit": 1,
			  "makerFeeRate": "0.0002",
			  "takerFeeRate": "0.0002",
			  "minQty": "0.01",
			  "maxQty": "100",
			  "initialMargin": "0.15",
			  "maintenanceMargin": "0.075",
			  "minInitialMargin": "0.1",
			  "minMaintenanceMargin": "0.05",
			  "priceScale": 2,
			  "quantityScale": 2,
			  "quoteAsset": "USDT"
		  }
		],
		"rateLimits": [
		  {
			  "rateLimitType": "REQUEST_WEIGHT",
			  "interval": "MINUTE",
			  "intervalNum": 1,
			  "limit": 2400
		  },
		  {
			  "rateLimitType": "ORDERS",
			  "interval": "MINUTE",
			  "intervalNum": 1,
			  "limit": 1200
		  },
		  {
			  "rateLimitType": "ORDERS",
			  "interval": "SECOND",
			  "intervalNum": 10,
			  "limit": 300
		  }
		]
	  }	  
	`)
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
		ServerTime: 1592387337630,
		OptionContracts: []OptionContract{
			{
				Id:          1,
				BaseAsset:   "BTC",
				QuoteAsset:  "USDT",
				Underlying:  "BTCUSDT",
				SettleAsset: "USDT",
			},
		},
		OptionAssets: []OptionAsset{
			{
				Id:   1,
				Name: "USDT",
			},
		},
		OptionSymbols: []OptionSymbol{
			{
				ContractId: 2,
				ExpiryDate: 1660521600000,
				Filters: []map[string]any{
					{"filterType": "PRICE_FILTER", "minPrice": "0.02", "maxPrice": "80000.01", "tickSize": "0.01"},
					{"filterType": "LOT_SIZE", "minQty": "0.01", "maxQty": "100", "stepSize": "0.01"},
				},
				Id:                   17,
				Symbol:               "BTC-220815-50000-C",
				Side:                 "CALL",
				StrikePrice:          "50000",
				Underlying:           "BTCUSDT",
				Unit:                 1,
				MakerFeeRate:         "0.0002",
				TakerFeeRate:         "0.0002",
				MinQty:               "0.01",
				MaxQty:               "100",
				InitialMargin:        "0.15",
				MaintenanceMargin:    "0.075",
				MinInitialMargin:     "0.1",
				MinMaintenanceMargin: "0.05",
				PriceScale:           2,
				QuantityScale:        2,
				QuoteAsset:           "USDT",
			},
		},
		RateLimits: []RateLimit{
			{RateLimitType: "REQUEST_WEIGHT", Interval: "MINUTE", IntervalNum: 1, Limit: 2400},
			{RateLimitType: "ORDERS", Interval: "MINUTE", IntervalNum: 1, Limit: 1200},
			{RateLimitType: "ORDERS", Interval: "SECOND", IntervalNum: 10, Limit: 300},
		},
	}
	s.assertExchangeInfoEqual(ei, res)
	s.r().Len(ei.OptionSymbols[0].Filters, 2, "Filters")

	ePriceFilter := &PriceFilter{
		MinPrice: "0.02",
		MaxPrice: "80000.01",
		TickSize: "0.01",
	}
	s.assertPriceFilterEqual(ePriceFilter, res.OptionSymbols[0].PriceFilter())

	eLotSizeFilter := &LotSizeFilter{
		MinQuantity: "0.01",
		MaxQuantity: "100",
		StepSize:    "0.01",
	}
	s.assertLotSizeFilterEqual(eLotSizeFilter, res.OptionSymbols[0].LotSizeFilter())
}

func (s *exchangeInfoServiceTestSuite) assertExchangeInfoEqual(e, a *ExchangeInfo) {
	r := s.r()

	r.Equal(e.Timezone, a.Timezone, "Timezone")
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")

	r.Len(a.OptionContracts, len(e.OptionContracts), "OptionContracts")
	for i := range a.OptionContracts {
		r.Equal(e.OptionContracts[i].Id, a.OptionContracts[i].Id, "Id")
		r.Equal(e.OptionContracts[i].BaseAsset, a.OptionContracts[i].BaseAsset, "BaseAsset")
		r.Equal(e.OptionContracts[i].QuoteAsset, a.OptionContracts[i].QuoteAsset, "QuoteAsset")
		r.Equal(e.OptionContracts[i].Underlying, a.OptionContracts[i].Underlying, "Underlying")
		r.Equal(e.OptionContracts[i].SettleAsset, a.OptionContracts[i].SettleAsset, "SettleAsset")
	}

	r.Len(a.OptionAssets, len(e.OptionAssets), "OptionAssets")
	for i := range a.OptionAssets {
		r.Equal(e.OptionAssets[i].Id, a.OptionAssets[i].Id, "Id")
		r.Equal(e.OptionAssets[i].Name, a.OptionAssets[i].Name, "Name")
	}

	r.Len(a.OptionSymbols, len(e.OptionSymbols), "Symbols")
	for i := range a.OptionSymbols {
		r.Equal(e.OptionSymbols[i].ContractId, a.OptionSymbols[i].ContractId, "ContractId")
		r.Equal(e.OptionSymbols[i].ExpiryDate, a.OptionSymbols[i].ExpiryDate, "ExpiryDate")
		r.Equal(e.OptionSymbols[i].Id, a.OptionSymbols[i].Id, "Id")
		r.Equal(e.OptionSymbols[i].Symbol, a.OptionSymbols[i].Symbol, "Symbol")
		r.Equal(e.OptionSymbols[i].Side, a.OptionSymbols[i].Side, "Side")
		r.Equal(e.OptionSymbols[i].StrikePrice, a.OptionSymbols[i].StrikePrice, "StrikePrice")
		r.Equal(e.OptionSymbols[i].Underlying, a.OptionSymbols[i].Underlying, "Underlying")
		r.Equal(e.OptionSymbols[i].Unit, a.OptionSymbols[i].Unit, "Unit")
		r.Equal(e.OptionSymbols[i].MakerFeeRate, a.OptionSymbols[i].MakerFeeRate, "MakerFeeRate")
		r.Equal(e.OptionSymbols[i].TakerFeeRate, a.OptionSymbols[i].TakerFeeRate, "TakerFeeRate")
		r.Equal(e.OptionSymbols[i].MinQty, a.OptionSymbols[i].MinQty, "MinQty")
		r.Equal(e.OptionSymbols[i].MaxQty, a.OptionSymbols[i].MaxQty, "MaxQty")
		r.Equal(e.OptionSymbols[i].InitialMargin, a.OptionSymbols[i].InitialMargin, "InitialMargin")
		r.Equal(e.OptionSymbols[i].MaintenanceMargin, a.OptionSymbols[i].MaintenanceMargin, "MaintenanceMargin")
		r.Equal(e.OptionSymbols[i].MinInitialMargin, a.OptionSymbols[i].MinInitialMargin, "MinInitialMargin")
		r.Equal(e.OptionSymbols[i].MinMaintenanceMargin, a.OptionSymbols[i].MinMaintenanceMargin, "MinMaintenanceMargin")
		r.Equal(e.OptionSymbols[i].PriceScale, a.OptionSymbols[i].PriceScale, "PriceScale")
		r.Equal(e.OptionSymbols[i].QuantityScale, a.OptionSymbols[i].QuantityScale, "QuantityScale")
		r.Equal(e.OptionSymbols[i].QuoteAsset, a.OptionSymbols[i].QuoteAsset, "QuoteAsset")

		for fi, currentFilter := range a.OptionSymbols[i].Filters {
			r.Len(currentFilter, len(e.OptionSymbols[i].Filters[fi]))
			switch currentFilter["filterType"] {
			case "PRICE_FILTER":
				r.Equal(e.OptionSymbols[i].PriceFilter().MinPrice, a.OptionSymbols[i].PriceFilter().MinPrice, "MinPrice")
				r.Equal(e.OptionSymbols[i].PriceFilter().MaxPrice, a.OptionSymbols[i].PriceFilter().MaxPrice, "MaxPrice")
				r.Equal(e.OptionSymbols[i].PriceFilter().TickSize, a.OptionSymbols[i].PriceFilter().TickSize, "TickSize")
			case "LOT_SIZE":
				r.Equal(e.OptionSymbols[i].LotSizeFilter().MinQuantity, a.OptionSymbols[i].LotSizeFilter().MinQuantity, "MinQuantity")
				r.Equal(e.OptionSymbols[i].LotSizeFilter().MaxQuantity, a.OptionSymbols[i].LotSizeFilter().MaxQuantity, "MaxQuantity")
				r.Equal(e.OptionSymbols[i].LotSizeFilter().StepSize, a.OptionSymbols[i].LotSizeFilter().StepSize, "StepSize")
			}
		}
	}

	r.Len(a.RateLimits, len(e.RateLimits), "RateLimits")
	for i := range a.RateLimits {
		r.Equal(e.RateLimits[i].RateLimitType, a.RateLimits[i].RateLimitType, "RateLimitType")
		r.Equal(e.RateLimits[i].Limit, a.RateLimits[i].Limit, "Limit")
		r.Equal(e.RateLimits[i].Interval, a.RateLimits[i].Interval, "Interval")
		r.Equal(e.RateLimits[i].IntervalNum, a.RateLimits[i].IntervalNum, "IntervalNum")
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
