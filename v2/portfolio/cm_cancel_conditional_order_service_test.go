package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmCancelConditionalOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMCancelConditionalOrderService(t *testing.T) {
	suite.Run(t, new(cmCancelConditionalOrderServiceTestSuite))
}

func (s *cmCancelConditionalOrderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"newClientStrategyId": "myOrder1",
		"strategyId": 123445,
		"strategyStatus": "CANCELED",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "11",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSD",
		"timeInForce": "GTC",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false
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

	res, err := s.client.NewCMCancelConditionalOrderService().
		Symbol(symbol).
		StrategyID(strategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(123445), res.StrategyID)
	s.r().Equal("CANCELED", res.StrategyStatus)
	s.r().Equal("TRAILING_STOP_MARKET", res.StrategyType)
	s.r().Equal("BTCUSD", res.Symbol)
}

func (s *cmCancelConditionalOrderServiceTestSuite) TestCancelOrderWithClientStrategyID() {
	data := []byte(`{
		"newClientStrategyId": "myOrder1",
		"strategyId": 123445,
		"strategyStatus": "CANCELED",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "11",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSD",
		"timeInForce": "GTC",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	newClientStrategyID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":              symbol,
			"newClientStrategyId": newClientStrategyID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelConditionalOrderService().
		Symbol(symbol).
		NewClientStrategyID(newClientStrategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("myOrder1", res.NewClientStrategyID)
	s.r().Equal("CANCELED", res.StrategyStatus)
}
