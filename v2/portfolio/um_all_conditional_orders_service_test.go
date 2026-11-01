package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAllConditionalOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestUMAllConditionalOrdersService(t *testing.T) {
	suite.Run(t, new(umAllConditionalOrdersServiceTestSuite))
}

func (s *umAllConditionalOrdersServiceTestSuite) TestAllConditionalOrders() {
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
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"priceMatch": "NONE"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	strategyID := int64(123445)
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"strategyId": strategyID,
			"limit":      limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMAllConditionalOrdersService().
		Symbol(symbol).
		StrategyID(strategyID).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(123445), orders[0].StrategyID)
	s.r().Equal("BTCUSDT", orders[0].Symbol)
	s.r().Equal("TRIGGERED", orders[0].StrategyStatus)
}

func (s *umAllConditionalOrdersServiceTestSuite) TestAllConditionalOrdersWithTimeRange() {
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
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"priceMatch": "NONE"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1566818724710)
	endTime := int64(1566818724750)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMAllConditionalOrdersService().
		Symbol(symbol).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1566818724710), orders[0].BookTime)
}
