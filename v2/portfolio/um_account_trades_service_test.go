package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAccountTradesServiceTestSuite struct {
	baseTestSuite
}

func TestUMAccountTradesService(t *testing.T) {
	suite.Run(t, new(umAccountTradesServiceTestSuite))
}

func (s *umAccountTradesServiceTestSuite) TestGetTrades() {
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
	fromID := int64(67880589)
	startTime := int64(1680688557875)
	endTime := int64(1680688557876)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"limit":     limit,
			"fromId":    fromID,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewUMAccountTradesService().
		Symbol(symbol).
		Limit(limit).
		FromID(fromID).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
	s.r().Equal("BTCUSDT", trades[0].Symbol)
	s.r().Equal(int64(67880589), trades[0].ID)
	s.r().Equal(int64(270093109), trades[0].OrderID)
	s.r().Equal("SELL", trades[0].Side)
	s.r().Equal("28511.00", trades[0].Price)
	s.r().Equal("BOTH", trades[0].PositionSide)
}
