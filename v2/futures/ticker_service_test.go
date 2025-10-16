package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tickerServiceTestSuite struct {
	baseTestSuite
}

func TestTickerService(t *testing.T) {
	suite.Run(t, new(tickerServiceTestSuite))
}

func (s *tickerServiceTestSuite) TestListBookTickers() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "bidPrice": "4.00000000",
            "bidQty": "431.00000000",
            "askPrice": "4.00000200",
            "askQty": "9.00000000",
			"time":1757119264167,
			"lastUpdateId":8524099887582
        },
        {
            "symbol": "ETHBTC",
            "bidPrice": "0.07946700",
            "bidQty": "9.00000000",
            "askPrice": "100000.00000000",
            "askQty": "1000.00000000",
			"time":1757119264167,
			"lastUpdateId":8524099887582
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	tickers, _, err := s.client.NewListBookTickersService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 2)
	e1 := &BookTicker{
		Symbol:       "LTCBTC",
		BidPrice:     "4.00000000",
		BidQuantity:  "431.00000000",
		AskPrice:     "4.00000200",
		AskQuantity:  "9.00000000",
		Time:         1757119264167,
		LastUpdateId: 8524099887582,
	}
	e2 := &BookTicker{
		Symbol:       "ETHBTC",
		BidPrice:     "0.07946700",
		BidQuantity:  "9.00000000",
		AskPrice:     "100000.00000000",
		AskQuantity:  "1000.00000000",
		Time:         1757119264167,
		LastUpdateId: 8524099887582,
	}
	s.assertBookTickerEqual(e1, tickers[0])
	s.assertBookTickerEqual(e2, tickers[1])
}

func (s *tickerServiceTestSuite) TestSingleBookTicker() {
	data := []byte(`{
            "symbol": "LTCBTC",
            "bidPrice": "4.00000000",
            "bidQty": "431.00000000",
            "askPrice": "4.00000200",
            "askQty": "9.00000000",
            "time":1757119264167,
            "lastUpdateId":8524099887582
        }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"

	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	tickers, _, err := s.client.NewListBookTickersService().Symbol("LTCBTC").Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 1)
	e := &BookTicker{
		Symbol:       "LTCBTC",
		BidPrice:     "4.00000000",
		BidQuantity:  "431.00000000",
		AskPrice:     "4.00000200",
		AskQuantity:  "9.00000000",
		Time:         1757119264167,
		LastUpdateId: 8524099887582,
	}
	s.assertBookTickerEqual(e, tickers[0])
}

func (s *tickerServiceTestSuite) assertBookTickerEqual(e, a *BookTicker) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.BidQuantity, a.BidQuantity, "BidQuantity")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.AskQuantity, a.AskQuantity, "AskQuantity")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.LastUpdateId, a.LastUpdateId, "LastUpdateId")
}

func (s *tickerServiceTestSuite) TestListPrices() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "price": "4.00000200"
        },
        {
            "symbol": "ETHBTC",
            "price": "0.07946600"
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	prices, _, err := s.client.NewListPricesService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 2)
	e1 := &SymbolPrice{
		Symbol: "LTCBTC",
		Price:  "4.00000200",
	}
	e2 := &SymbolPrice{
		Symbol: "ETHBTC",
		Price:  "0.07946600",
	}
	s.assertSymbolPriceEqual(e1, prices[0])
	s.assertSymbolPriceEqual(e2, prices[1])
}

func (s *tickerServiceTestSuite) TestListSinglePrice() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"price": "4.00000200"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	prices, _, err := s.client.NewListPricesService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(prices, 1)
	e1 := &SymbolPrice{
		Symbol: "LTCBTC",
		Price:  "4.00000200",
	}
	s.assertSymbolPriceEqual(e1, prices[0])
}

func (s *tickerServiceTestSuite) assertSymbolPriceEqual(e, a *SymbolPrice) {
	r := s.r()
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *tickerServiceTestSuite) TestPriceChangeStats() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"priceChange": "-94.99999800",
		"priceChangePercent": "-95.960",
		"weightedAvgPrice": "0.29628482",
		"prevClosePrice": "0.10002000",
		"lastPrice": "4.00000200",
		"lastQty": "200.00000000",
		"openPrice": "99.00000000",
		"highPrice": "100.00000000",
		"lowPrice": "0.10000000",
		"volume": "8913.30000000",
		"quoteVolume": "15.30000000",
		"openTime": 1499783499040,
		"closeTime": 1499869899040,
		"firstId": 28385,
		"lastId": 28460,
		"count": 76
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	stats, _, err := s.client.NewListPriceChangeStatsService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(stats, 1)
	e := &PriceChangeStats{
		Symbol:             "BTCUSDT",
		PriceChange:        "-94.99999800",
		PriceChangePercent: "-95.960",
		WeightedAvgPrice:   "0.29628482",
		PrevClosePrice:     "0.10002000",
		LastPrice:          "4.00000200",
		LastQuantity:       "200.00000000",
		OpenPrice:          "99.00000000",
		HighPrice:          "100.00000000",
		LowPrice:           "0.10000000",
		Volume:             "8913.30000000",
		OpenTime:           1499783499040,
		CloseTime:          1499869899040,
		FirstID:            28385,
		LastID:             28460,
		Count:              76,
	}
	s.assertPriceChangeStatsEqual(e, stats[0])
}

func (s *tickerServiceTestSuite) assertPriceChangeStatsEqual(e, a *PriceChangeStats) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.PrevClosePrice, a.PrevClosePrice, "PrevClosePrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.LastQuantity, a.LastQuantity, "LastQuantity")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.Count, a.Count, "Count")
}

func (s *tickerServiceTestSuite) TestListPriceChangeStats() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"priceChange": "-94.99999800",
			"priceChangePercent": "-95.960",
			"weightedAvgPrice": "0.29628482",
			"prevClosePrice": "0.10002000",
			"lastPrice": "4.00000200",
			"lastQty": "200.00000000",
			"openPrice": "99.00000000",
			"highPrice": "100.00000000",
			"lowPrice": "0.10000000",
			"volume": "8913.30000000",
			"quoteVolume": "15.30000000",
			"openTime": 1499783499040,
			"closeTime": 1499869899040,
			"firstId": 28385,
			"lastId": 28460,
			"count": 76
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, _, err := s.client.NewListPriceChangeStatsService().Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*PriceChangeStats{
		{
			Symbol:             "BTCUSDT",
			PriceChange:        "-94.99999800",
			PriceChangePercent: "-95.960",
			WeightedAvgPrice:   "0.29628482",
			PrevClosePrice:     "0.10002000",
			LastPrice:          "4.00000200",
			LastQuantity:       "200.00000000",
			OpenPrice:          "99.00000000",
			HighPrice:          "100.00000000",
			LowPrice:           "0.10000000",
			Volume:             "8913.30000000",
			QuoteVolume:        "15.30000000",
			OpenTime:           1499783499040,
			CloseTime:          1499869899040,
			FirstID:            28385,
			LastID:             28460,
			Count:              76,
		},
	}
	s.r().Len(res, len(e))
	for i := range res {
		s.assertPriceChangeStatsEqual(e[i], res[i])
	}
}
