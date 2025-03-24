package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type tradingDayTickerServiceTestSuite struct {
	baseTestSuite
}

func TestTradingDayTickerService(t *testing.T) {
	suite.Run(t, new(tradingDayTickerServiceTestSuite))
}

func (s *tradingDayTickerServiceTestSuite) TestTradingDayTicker() {
	data := []byte(`{
 "symbol": "BTCUSDT",
 "priceChange": "-25.63000000",
 "priceChangePercent": "-0.042",
 "weightedAvgPrice": "60989.95178227",
 "openPrice": "60986.68000000",
 "highPrice": "61078.01000000",
 "lowPrice": "60922.00000000",
 "lastPrice": "60961.05000000",
 "volume": "707.02697000",
 "quoteVolume": "43121540.80906740",
 "openTime": 1719705600000,
 "closeTime": 1719791999999,
 "firstId": 3655222498,
 "lastId": 3655289744,
 "count": 67247
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbols = []string{"BTCUSDT", "ETHUSDT"}
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbols", symbols)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewTradingDayTickerService().Symbols(symbols).Do(newContext())
	s.r().NoError(err)

	e := []*TradingDayTicker{
		{
			Symbol:             "BTCUSDT",
			PriceChange:        "-25.63000000",
			PriceChangePercent: "-0.042",
			WeightedAvgPrice:   "60989.95178227",
			OpenPrice:          "60986.68000000",
			HighPrice:          "61078.01000000",
			LowPrice:           "60922.00000000",
			LastPrice:          "60961.05000000",
			Volume:             "707.02697000",
			QuoteVolume:        "43121540.80906740",
			OpenTime:           1719705600000,
			CloseTime:          1719791999999,
			FirstId:            3655222498,
			LastId:             3655289744,
			Count:              67247}}

	s.assertTradingDayTickersEqual(e, res)
}

func (s *tradingDayTickerServiceTestSuite) assertTradingDayTickerEqual(e, a *TradingDayTicker) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstId, a.FirstId, "FirstId")
	r.Equal(e.LastId, a.LastId, "LastId")
	r.Equal(e.Count, a.Count, "Count")
}

func (s *tradingDayTickerServiceTestSuite) assertTradingDayTickersEqual(e, a []*TradingDayTicker) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertTradingDayTickerEqual(e[i], a[i])
	}
}
