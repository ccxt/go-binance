package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umConditionalOrderHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestUMConditionalOrderHistoryService(t *testing.T) {
	suite.Run(t, new(umConditionalOrderHistoryServiceTestSuite))
}

func (s *umConditionalOrderHistoryServiceTestSuite) TestConditionalOrderHistory() {
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
		"symbol": "BTCUSDT",
		"orderId": 12132343435,
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
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0
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

	res, err := s.client.NewUMConditionalOrderHistoryService().
		Symbol(symbol).
		StrategyID(strategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(123445), res.StrategyID)
	s.r().Equal("BTCUSDT", res.Symbol)
	s.r().Equal("TRIGGERED", res.StrategyStatus)
	s.r().Equal("CONTRACT_PRICE", res.WorkingType)
}

func (s *umConditionalOrderHistoryServiceTestSuite) TestConditionalOrderHistoryWithClientStrategyID() {
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
		"symbol": "BTCUSDT",
		"orderId": 12132343435,
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
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0
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

	res, err := s.client.NewUMConditionalOrderHistoryService().
		Symbol(symbol).
		NewClientStrategyID(newClientStrategyID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", res.NewClientStrategyID)
	s.r().Equal(int64(12132343435), res.OrderID)
}
