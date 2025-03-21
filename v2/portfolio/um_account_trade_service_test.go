package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAccountTradeServiceTestSuite struct {
	baseTestSuite
}

func TestUMAccountTradeService(t *testing.T) {
	suite.Run(t, new(umAccountTradeServiceTestSuite))
}

func (s *umAccountTradeServiceTestSuite) TestAccountTrades() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"id": 67880589,
			"orderId": 270093109,
			"side": "SELL",
			"price": "28511.00",
			"qty": "0.010",
			"realizedPnl": "2.58500000",
			"quoteQty": "285.11000",
			"commission": "-0.11404400",
			"commissionAsset": "USDT",
			"time": 1680688557875,
			"buyer": false,
			"maker": false,
			"positionSide": "BOTH"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewUMAccountTradeService().
		Symbol(symbol).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
	s.r().Equal(int64(67880589), trades[0].ID)
	s.r().Equal("BTCUSDT", trades[0].Symbol)
	s.r().Equal("SELL", trades[0].Side)
	s.r().Equal("28511.00", trades[0].Price)
	s.r().Equal("2.58500000", trades[0].RealizedPnl)
}

func (s *umAccountTradeServiceTestSuite) TestAccountTradesWithAllParams() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"id": 67880589,
			"orderId": 270093109,
			"side": "SELL",
			"price": "28511.00",
			"qty": "0.010",
			"realizedPnl": "2.58500000",
			"quoteQty": "285.11000",
			"commission": "-0.11404400",
			"commissionAsset": "USDT",
			"time": 1680688557875,
			"buyer": false,
			"maker": false,
			"positionSide": "BOTH"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1680688557000)
	endTime := int64(1680688557999)
	fromID := int64(67880500)
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"fromId":    fromID,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewUMAccountTradeService().
		Symbol(symbol).
		StartTime(startTime).
		EndTime(endTime).
		FromID(fromID).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
}
