package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umOpenConditionalOrderServiceTestSuite struct {
	baseTestSuite
}

func TestUMOpenConditionalOrderService(t *testing.T) {
	suite.Run(t, new(umOpenConditionalOrderServiceTestSuite))
}

func (s *umOpenConditionalOrderServiceTestSuite) TestOpenConditionalOrder() {
	data := []byte(`{
		"newClientStrategyId": "abc",
		"strategyId": 123445,
		"strategyStatus": "NEW",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "0.40",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSDT",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"timeInForce": "GTC",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0,
		"priceMatch": "NONE"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	strategyID := int64(123445)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"strategyId": strategyID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMOpenConditionalOrderService().
		Symbol(symbol).
		StrategyID(strategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(123445), res.StrategyID)
	s.r().Equal("BTCUSDT", res.Symbol)
	s.r().Equal("NEW", res.StrategyStatus)
	s.r().Equal("TRAILING_STOP_MARKET", res.StrategyType)
}

func (s *umOpenConditionalOrderServiceTestSuite) TestOpenConditionalOrderWithClientStrategyID() {
	data := []byte(`{
		"newClientStrategyId": "abc",
		"strategyId": 123445,
		"strategyStatus": "NEW",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "0.40",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSDT",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"timeInForce": "GTC",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0,
		"priceMatch": "NONE"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	newClientStrategyID := "abc"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":              symbol,
			"newClientStrategyId": newClientStrategyID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMOpenConditionalOrderService().
		Symbol(symbol).
		NewClientStrategyID(newClientStrategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", res.NewClientStrategyID)
	s.r().Equal("0.3", res.PriceRate)
}
