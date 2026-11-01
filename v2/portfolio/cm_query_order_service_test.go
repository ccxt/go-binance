package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmQueryOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMQueryOrderService(t *testing.T) {
	suite.Run(t, new(cmQueryOrderServiceTestSuite))
}

func (s *cmQueryOrderServiceTestSuite) TestQueryOrder() {
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
		"status": "NEW",
		"symbol": "BTCUSD_200925",
		"pair": "BTCUSD",
		"positionSide": "SHORT",
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

	res, err := s.client.NewCMQueryOrderService().
		Symbol(symbol).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1917641), res.OrderID)
	s.r().Equal("BTCUSD_200925", res.Symbol)
	s.r().Equal("BTCUSD", res.Pair)
	s.r().Equal("NEW", res.Status)
}

func (s *cmQueryOrderServiceTestSuite) TestQueryOrderWithClientOrderID() {
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
		"status": "NEW",
		"symbol": "BTCUSD_200925",
		"pair": "BTCUSD",
		"positionSide": "SHORT",
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

	res, err := s.client.NewCMQueryOrderService().
		Symbol(symbol).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("abc", res.ClientOrderID)
	s.r().Equal(int64(1579276756075), res.Time)
}
