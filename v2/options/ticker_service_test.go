package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TickerServiceTestSuite struct {
	baseTestSuite
}

func TestTickerService(t *testing.T) {
	suite.Run(t, new(TickerServiceTestSuite))
}

func (s *TickerServiceTestSuite) TestTicker() {
	data := []byte(`[
		{
		 "symbol": "ETH-240628-800-C",
		 "priceChange": "0",
		 "priceChangePercent": "0",
		 "lastPrice": "3010.6",
		 "lastQty": "0",
		 "open": "3010.6",
		 "high": "3010.6",
		 "low": "3010.6",
		 "volume": "0",
		 "amount": "0",
		 "bidPrice": "2620",
		 "askPrice": "0",
		 "openTime": 0,
		 "closeTime": 0,
		 "firstTradeId": 0,
		 "tradeCount": 0,
		 "strikePrice": "800",
		 "exercisePrice": "3784.30113636"
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	tickers, err := s.client.NewTickerService().Do(newContext())
	targetTickers := []*Ticker{
		{
			Symbol:             "ETH-240628-800-C",
			PriceChange:        "0",
			PriceChangePercent: "0",
			LastPrice:          "3010.6",
			LastQty:            "0",
			Open:               "3010.6",
			High:               "3010.6",
			Low:                "3010.6",
			Volume:             "0",
			Amount:             "0",
			BidPrice:           "2620",
			AskPrice:           "0",
			OpenTime:           0,
			CloseTime:          0,
			FirstTradeId:       0,
			TradeCount:         0,
			StrikePrice:        "800",
			ExercisePrice:      "3784.30113636",
		},
	}

	s.r().Equal(err, nil, "err != nil")
	for i := range tickers {
		r := s.r()
		r.Equal(tickers[i].Symbol, targetTickers[i].Symbol, "Symbol")
		r.Equal(tickers[i].PriceChange, targetTickers[i].PriceChange, "PriceChange")
		r.Equal(tickers[i].PriceChangePercent, targetTickers[i].PriceChangePercent, "PriceChangePercent")
		r.Equal(tickers[i].LastPrice, targetTickers[i].LastPrice, "LastPrice")
		r.Equal(tickers[i].LastQty, targetTickers[i].LastQty, "LastQty")
		r.Equal(tickers[i].Open, targetTickers[i].Open, "Open")
		r.Equal(tickers[i].High, targetTickers[i].High, "High")
		r.Equal(tickers[i].Low, targetTickers[i].Low, "Low")
		r.Equal(tickers[i].Volume, targetTickers[i].Volume, "Volume")
		r.Equal(tickers[i].Amount, targetTickers[i].Amount, "Amount")
		r.Equal(tickers[i].BidPrice, targetTickers[i].BidPrice, "BidPrice")
		r.Equal(tickers[i].AskPrice, targetTickers[i].AskPrice, "AskPrice")
		r.Equal(tickers[i].OpenTime, targetTickers[i].OpenTime, "OpenTime")
		r.Equal(tickers[i].CloseTime, targetTickers[i].CloseTime, "CloseTime")
		r.Equal(tickers[i].FirstTradeId, targetTickers[i].FirstTradeId, "FirstTradeId")
		r.Equal(tickers[i].TradeCount, targetTickers[i].TradeCount, "TradeCount")
		r.Equal(tickers[i].StrikePrice, targetTickers[i].StrikePrice, "StrikePrice")
		r.Equal(tickers[i].ExercisePrice, targetTickers[i].ExercisePrice, "ExercisePrice")
	}
}
