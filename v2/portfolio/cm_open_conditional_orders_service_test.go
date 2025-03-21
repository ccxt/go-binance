package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmOpenConditionalOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMOpenConditionalOrdersService(t *testing.T) {
	suite.Run(t, new(cmOpenConditionalOrdersServiceTestSuite))
}

func (s *cmOpenConditionalOrdersServiceTestSuite) TestOpenConditionalOrders() {
	data := []byte(`[
		{
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
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMOpenConditionalOrdersService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(123445), orders[0].StrategyID)
	s.r().Equal("BTCUSD", orders[0].Symbol)
	s.r().Equal("NEW", orders[0].StrategyStatus)
	s.r().Equal("TRAILING_STOP_MARKET", orders[0].StrategyType)
	s.r().Equal("9020", orders[0].ActivatePrice)
	s.r().Equal("0.3", orders[0].PriceRate)
}

func (s *cmOpenConditionalOrdersServiceTestSuite) TestOpenConditionalOrdersForAllSymbols() {
	data := []byte(`[
		{
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
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMOpenConditionalOrdersService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal("abc", orders[0].NewClientStrategyID)
}
