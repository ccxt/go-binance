package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAllOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestUMAllOrdersService(t *testing.T) {
	suite.Run(t, new(umAllOrdersServiceTestSuite))
}

func (s *umAllOrdersServiceTestSuite) TestAllOrders() {
	data := []byte(`[
		{
			"avgPrice": "0.00000",
			"clientOrderId": "abc",
			"cumQuote": "0",
			"executedQty": "0",
			"orderId": 1917641,
			"origQty": "0.40",
			"origType": "LIMIT",
			"price": "0",
			"reduceOnly": false,
			"side": "BUY",
			"positionSide": "SHORT",
			"status": "NEW",
			"symbol": "BTCUSDT",
			"time": 1579276756075,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1579276756075,
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"priceMatch": "NONE"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderID := int64(1917641)
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
			"limit":   limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMAllOrdersService().
		Symbol(symbol).
		OrderID(orderID).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1917641), orders[0].OrderID)
	s.r().Equal("BTCUSDT", orders[0].Symbol)
	s.r().Equal("NEW", orders[0].Status)
}

func (s *umAllOrdersServiceTestSuite) TestAllOrdersWithTimeRange() {
	data := []byte(`[
		{
			"avgPrice": "0.00000",
			"clientOrderId": "abc",
			"cumQuote": "0",
			"executedQty": "0",
			"orderId": 1917641,
			"origQty": "0.40",
			"origType": "LIMIT",
			"price": "0",
			"reduceOnly": false,
			"side": "BUY",
			"positionSide": "SHORT",
			"status": "NEW",
			"symbol": "BTCUSDT",
			"time": 1579276756075,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1579276756075,
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"priceMatch": "NONE"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1579276756075)
	endTime := int64(1579276756076)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMAllOrdersService().
		Symbol(symbol).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1579276756075), orders[0].Time)
}
