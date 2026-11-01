package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmAccountTradesServiceTestSuite struct {
	baseTestSuite
}

func TestCMAccountTradesService(t *testing.T) {
	suite.Run(t, new(cmAccountTradesServiceTestSuite))
}

func (s *cmAccountTradesServiceTestSuite) TestGetTradesBySymbol() {
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

	trades, err := s.client.NewCMAccountTradesService().
		Symbol(symbol).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
	s.r().Equal("BTCUSD_200626", trades[0].Symbol)
	s.r().Equal("BTCUSD", trades[0].Pair)
	s.r().Equal(int64(6), trades[0].ID)
	s.r().Equal(int64(28), trades[0].OrderID)
	s.r().Equal("SELL", trades[0].Side)
	s.r().Equal("BTC", trades[0].MarginAsset)
}

func (s *cmAccountTradesServiceTestSuite) TestGetTradesByPair() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_200626",
			"pair": "BTCUSD",
			"id": 6,
			"orderId": 28
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	pair := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"pair": pair,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewCMAccountTradesService().
		Pair(pair).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
	s.r().Equal("BTCUSD", trades[0].Pair)
}

func (s *cmAccountTradesServiceTestSuite) TestGetTradesValidation() {
	_, err := s.client.NewCMAccountTradesService().Do(newContext())
	s.r().Error(err)
	s.r().Equal("either symbol or pair must be sent", err.Error())

	_, err = s.client.NewCMAccountTradesService().
		Symbol("BTCUSD_200626").
		Pair("BTCUSD").
		Do(newContext())
	s.r().Error(err)
	s.r().Equal("symbol and pair cannot be sent together", err.Error())

	_, err = s.client.NewCMAccountTradesService().
		Pair("BTCUSD").
		FromID(6).
		Do(newContext())
	s.r().Error(err)
	s.r().Equal("pair and fromId cannot be sent together", err.Error())
}
