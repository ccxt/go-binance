package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmConditionalOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMConditionalOrdersService(t *testing.T) {
	suite.Run(t, new(cmConditionalOrdersServiceTestSuite))
}

func (s *cmConditionalOrdersServiceTestSuite) TestConditionalOrders() {
	data := []byte(`[
		{
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
			"priceRate": "0.3"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMConditionalOrdersService().
		Symbol(symbol).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(123445), orders[0].StrategyID)
	s.r().Equal("TRIGGERED", orders[0].StrategyStatus)
	s.r().Equal(int64(12123343534), orders[0].OrderID)
	s.r().Equal("NEW", orders[0].Status)
}

func (s *cmConditionalOrdersServiceTestSuite) TestConditionalOrdersWithAllParams() {
	data := []byte(`[
		{
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
			"priceRate": "0.3"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	strategyID := int64(123445)
	startTime := int64(1566818724000)
	endTime := int64(1566818724999)
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"strategyId": strategyID,
			"startTime":  startTime,
			"endTime":    endTime,
			"limit":      limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMConditionalOrdersService().
		Symbol(symbol).
		StrategyID(strategyID).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
}
