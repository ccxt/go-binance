package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmOpenOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMOpenOrderService(t *testing.T) {
	suite.Run(t, new(cmOpenOrderServiceTestSuite))
}

func (s *cmOpenOrderServiceTestSuite) TestOpenOrder() {
	data := []byte(`{
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
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	orderID := int64(1917641)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMOpenOrderService().
		Symbol(symbol).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1917641), res.OrderID)
	s.r().Equal("BTCUSD_200925", res.Symbol)
	s.r().Equal("BTCUSD", res.Pair)
	s.r().Equal("NEW", res.Status)
}

func (s *cmOpenOrderServiceTestSuite) TestOpenOrderWithClientOrderID() {
	data := []byte(`{
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
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	origClientOrderID := "abc"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMOpenOrderService().
		Symbol(symbol).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", res.ClientOrderID)
	s.r().Equal(int64(1579276756075), res.Time)
}
