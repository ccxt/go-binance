package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmOpenConditionalOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMOpenConditionalOrderService(t *testing.T) {
	suite.Run(t, new(cmOpenConditionalOrderServiceTestSuite))
}

func (s *cmOpenConditionalOrderServiceTestSuite) TestOpenConditionalOrder() {
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
		"symbol": "BTCUSD",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"timeInForce": "GTC",
		"activatePrice": "9020",
		"priceRate": "0.3"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	strategyID := int64(123445)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"strategyId": strategyID,
		})
		s.assertRequestEqual(e, r)
	})

	order, err := s.client.NewCMOpenConditionalOrderService().
		Symbol(symbol).
		StrategyID(strategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(123445), order.StrategyID)
	s.r().Equal("BTCUSD", order.Symbol)
	s.r().Equal("NEW", order.StrategyStatus)
	s.r().Equal("TRAILING_STOP_MARKET", order.StrategyType)
	s.r().Equal("9020", order.ActivatePrice)
	s.r().Equal("0.3", order.PriceRate)
}

func (s *cmOpenConditionalOrderServiceTestSuite) TestOpenConditionalOrderWithClientStrategyID() {
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
		"symbol": "BTCUSD",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"timeInForce": "GTC",
		"activatePrice": "9020",
		"priceRate": "0.3"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	newClientStrategyID := "abc"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":              symbol,
			"newClientStrategyId": newClientStrategyID,
		})
		s.assertRequestEqual(e, r)
	})

	order, err := s.client.NewCMOpenConditionalOrderService().
		Symbol(symbol).
		NewClientStrategyID(newClientStrategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", order.NewClientStrategyID)
}
