package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umOpenOrderServiceTestSuite struct {
	baseTestSuite
}

func TestUMOpenOrderService(t *testing.T) {
	suite.Run(t, new(umOpenOrderServiceTestSuite))
}

func (s *umOpenOrderServiceTestSuite) TestOpenOrder() {
	data := []byte(`{
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
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderID := int64(1917641)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMOpenOrderService().
		Symbol(symbol).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1917641), res.OrderID)
	s.r().Equal("BTCUSDT", res.Symbol)
	s.r().Equal("NEW", res.Status)
	s.r().Equal("BUY", res.Side)
	s.r().Equal("SHORT", res.PositionSide)
}

func (s *umOpenOrderServiceTestSuite) TestOpenOrderWithClientOrderID() {
	data := []byte(`{
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
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	origClientOrderID := "abc"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMOpenOrderService().
		Symbol(symbol).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", res.ClientOrderID)
	s.r().Equal(int64(1579276756075), res.Time)
}
