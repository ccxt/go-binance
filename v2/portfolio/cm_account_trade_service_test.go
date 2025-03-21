package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmAccountTradeServiceTestSuite struct {
	baseTestSuite
}

func TestCMAccountTradeService(t *testing.T) {
	suite.Run(t, new(cmAccountTradeServiceTestSuite))
}

func (s *cmAccountTradeServiceTestSuite) TestAccountTradesWithSymbol() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_200626",
			"id": 6,
			"orderId": 28,
			"pair": "BTCUSD",
			"side": "SELL",
			"price": "8800",
			"qty": "1",
			"realizedPnl": "0",
			"marginAsset": "BTC",
			"baseQty": "0.01136364",
			"commission": "0.00000454",
			"commissionAsset": "BTC",
			"time": 1590743483586,
			"positionSide": "BOTH",
			"buyer": false,
			"maker": false
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200626"
	limit := 50
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewCMAccountTradeService().
		Symbol(symbol).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
	s.r().Equal(int64(6), trades[0].ID)
	s.r().Equal("BTCUSD_200626", trades[0].Symbol)
	s.r().Equal("BTCUSD", trades[0].Pair)
	s.r().Equal("SELL", trades[0].Side)
	s.r().Equal("BTC", trades[0].MarginAsset)
}

func (s *cmAccountTradeServiceTestSuite) TestAccountTradesWithPair() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_200626",
			"id": 6,
			"orderId": 28,
			"pair": "BTCUSD",
			"side": "SELL",
			"price": "8800",
			"qty": "1",
			"realizedPnl": "0",
			"marginAsset": "BTC",
			"baseQty": "0.01136364",
			"commission": "0.00000454",
			"commissionAsset": "BTC",
			"time": 1590743483586,
			"positionSide": "BOTH",
			"buyer": false,
			"maker": false
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	pair := "BTCUSD"
	startTime := int64(1590743483000)
	endTime := int64(1590743483999)
	limit := 50
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"pair":      pair,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewCMAccountTradeService().
		Pair(pair).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
}
