package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmConditionalOrderHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestCMConditionalOrderHistoryService(t *testing.T) {
	suite.Run(t, new(cmConditionalOrderHistoryServiceTestSuite))
}

func (s *cmConditionalOrderHistoryServiceTestSuite) TestConditionalOrderHistory() {
	data := []byte(`{
		"newClientStrategyId": "abc",
		"strategyId": 123445,
		"strategyStatus": "TRIGGERED",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "0.40",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSD",
		"orderId": 12123343534,
		"status": "NEW",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"triggerTime": 1566818724750,
		"timeInForce": "GTC",
		"type": "MARKET",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false,
		"priceMatch": "NONE"
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

	order, err := s.client.NewCMConditionalOrderHistoryService().
		Symbol(symbol).
		StrategyID(strategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(123445), order.StrategyID)
	s.r().Equal("TRIGGERED", order.StrategyStatus)
	s.r().Equal(int64(12123343534), order.OrderID)
	s.r().Equal("CONTRACT_PRICE", order.WorkingType)
	s.r().Equal(false, order.PriceProtect)
	s.r().Equal("NONE", order.PriceMatch)
}

func (s *cmConditionalOrderHistoryServiceTestSuite) TestConditionalOrderHistoryWithClientStrategyID() {
	data := []byte(`{
		"newClientStrategyId": "abc",
		"strategyId": 123445,
		"strategyStatus": "TRIGGERED",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "0.40",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSD",
		"orderId": 12123343534,
		"status": "NEW",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"triggerTime": 1566818724750,
		"timeInForce": "GTC",
		"type": "MARKET",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false,
		"priceMatch": "NONE"
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

	order, err := s.client.NewCMConditionalOrderHistoryService().
		Symbol(symbol).
		NewClientStrategyID(newClientStrategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", order.NewClientStrategyID)
}
