package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmOpenOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMOpenOrdersService(t *testing.T) {
	suite.Run(t, new(cmOpenOrdersServiceTestSuite))
}

func (s *cmOpenOrdersServiceTestSuite) TestOpenOrders() {
	data := []byte(`[
		{
			"avgPrice": "0.0",
			"clientOrderId": "abc",
			"cumBase": "0",
			"executedQty": "0",
			"orderId": 1917641,
			"origQty": "0.40",
			"origType": "LIMIT",
			"price": "0",
			"reduceOnly": false,
			"side": "BUY",
			"positionSide": "SHORT",
			"status": "NEW",
			"symbol": "BTCUSD_200925",
			"pair": "BTCUSD",
			"time": 1579276756075,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1579276756075
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMOpenOrdersService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1917641), orders[0].OrderID)
	s.r().Equal("BTCUSD_200925", orders[0].Symbol)
	s.r().Equal("BTCUSD", orders[0].Pair)
	s.r().Equal("NEW", orders[0].Status)
}

func (s *cmOpenOrdersServiceTestSuite) TestOpenOrdersWithPair() {
	data := []byte(`[
		{
			"avgPrice": "0.0",
			"clientOrderId": "abc",
			"cumBase": "0",
			"executedQty": "0",
			"orderId": 1917641,
			"origQty": "0.40",
			"origType": "LIMIT",
			"price": "0",
			"reduceOnly": false,
			"side": "BUY",
			"positionSide": "SHORT",
			"status": "NEW",
			"symbol": "BTCUSD_200925",
			"pair": "BTCUSD",
			"time": 1579276756075,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1579276756075
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

	orders, err := s.client.NewCMOpenOrdersService().Pair(pair).Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal("BTCUSD", orders[0].Pair)
}
