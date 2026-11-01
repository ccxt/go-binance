package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TradeServiceTestSuite struct {
	baseTestSuite
}

func TestTradeService(t *testing.T) {
	suite.Run(t, new(TradeServiceTestSuite))
}

func (s *TradeServiceTestSuite) TestTrades() {
	data := []byte(`[
		{
		 "id": 1125899906889469013,
		 "tradeId": 599,
		 "symbol": "BTC-240405-67000-C",
		 "price": "70.0",
		 "qty": "0.01",
		 "quoteQty": "0.7",
		 "side": -1,
		 "time": 1712303719297
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	trades, err := s.client.NewTradesService().Do(newContext())
	targetTrades := []*Trade{
		{
			Id:       1125899906889469013,
			TradeId:  599,
			Symbol:   "BTC-240405-67000-C",
			Price:    "70.0",
			Qty:      "0.01",
			QuoteQty: "0.7",
			Side:     -1,
			Time:     1712303719297,
		},
	}

	s.r().Equal(err, nil, "err != nil")
	for i := range trades {
		r := s.r()
		r.Equal(trades[i].Id, targetTrades[i].Id, "Id")
		r.Equal(trades[i].TradeId, targetTrades[i].TradeId, "TradeId")
		r.Equal(trades[i].Symbol, targetTrades[i].Symbol, "Symbol")
		r.Equal(trades[i].Price, targetTrades[i].Price, "Price")
		r.Equal(trades[i].Qty, targetTrades[i].Qty, "Qty")
		r.Equal(trades[i].QuoteQty, targetTrades[i].QuoteQty, "QuoteQty")
		r.Equal(trades[i].Side, targetTrades[i].Side, "Side")
		r.Equal(trades[i].Time, targetTrades[i].Time, "Time")
	}
}

func (s *TradeServiceTestSuite) TestHistoricalTrades() {

	data := []byte(` [
		{
		 "id": 1125899906889469013,
		 "tradeId": 599,
		 "price": "70.00000000",
		 "qty": "-0.01000000",
		 "quoteQty": "-0.7000000000000000",
		 "side": -1,
		 "time": 1712303719297
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	ht, err := s.client.NewHistoricalTradesService().Do(newContext())
	s.r().Equal(err, nil, "err != nil")

	targetHistoricalTrades := []*HistoricalTrade{
		{
			Id:       1125899906889469013,
			TradeId:  599,
			Price:    "70.00000000",
			Qty:      "-0.01000000",
			QuoteQty: "-0.7000000000000000",
			Side:     -1,
			Time:     1712303719297,
		},
	}
	for i := range ht {
		r := s.r()
		r.Equal(ht[i].Id, targetHistoricalTrades[i].Id, "Id")
		r.Equal(ht[i].TradeId, targetHistoricalTrades[i].TradeId, "TradeId")
		r.Equal(ht[i].Price, targetHistoricalTrades[i].Price, "Price")
		r.Equal(ht[i].Qty, targetHistoricalTrades[i].Qty, "Qty")
		r.Equal(ht[i].QuoteQty, targetHistoricalTrades[i].QuoteQty, "QuoteQty")
		r.Equal(ht[i].Side, targetHistoricalTrades[i].Side, "Side")
		r.Equal(ht[i].Time, targetHistoricalTrades[i].Time, "Time")
	}
}
