package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmAllOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMAllOrdersService(t *testing.T) {
	suite.Run(t, new(cmAllOrdersServiceTestSuite))
}

func (s *cmAllOrdersServiceTestSuite) TestAllOrders() {
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
	orderID := int64(1917641)
	limit := 50
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
			"limit":   limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMAllOrdersService().
		Symbol(symbol).
		OrderID(orderID).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1917641), orders[0].OrderID)
	s.r().Equal("BTCUSD_200925", orders[0].Symbol)
	s.r().Equal("BTCUSD", orders[0].Pair)
}

func (s *cmAllOrdersServiceTestSuite) TestAllOrdersWithPair() {
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
	startTime := int64(1579276756075)
	endTime := int64(1579276756076)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"pair":      pair,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewCMAllOrdersService().
		Pair(pair).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal("BTCUSD", orders[0].Pair)
}
