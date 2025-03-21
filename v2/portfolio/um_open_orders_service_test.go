package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umOpenOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestUMOpenOrdersService(t *testing.T) {
	suite.Run(t, new(umOpenOrdersServiceTestSuite))
}

func (s *umOpenOrdersServiceTestSuite) TestOpenOrders() {
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
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMOpenOrdersService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1917641), orders[0].OrderID)
	s.r().Equal("BTCUSDT", orders[0].Symbol)
	s.r().Equal("NEW", orders[0].Status)
}

func (s *umOpenOrdersServiceTestSuite) TestOpenOrdersForAllSymbols() {
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

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMOpenOrdersService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(1579276756075), orders[0].Time)
}
