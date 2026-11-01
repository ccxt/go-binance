package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmCancelOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMCancelOrderService(t *testing.T) {
	suite.Run(t, new(cmCancelOrderServiceTestSuite))
}

func (s *cmCancelOrderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"avgPrice": "0.0",
		"clientOrderId": "myOrder1",
		"cumQty": "0",
		"cumBase": "0",
		"executedQty": "0",
		"orderId": 283194212,
		"origQty": "2",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"status": "CANCELED",
		"symbol": "BTCUSD_200925",
		"pair": "BTCUSD",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1571110484038
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	orderID := int64(283194212)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelOrderService().
		Symbol(symbol).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(283194212), res.OrderID)
	s.r().Equal("CANCELED", res.Status)
	s.r().Equal("BTCUSD_200925", res.Symbol)
	s.r().Equal("BTCUSD", res.Pair)
}

func (s *cmCancelOrderServiceTestSuite) TestCancelOrderWithClientOrderID() {
	data := []byte(`{
		"avgPrice": "0.0",
		"clientOrderId": "myOrder1",
		"cumQty": "0",
		"cumBase": "0",
		"executedQty": "0",
		"orderId": 283194212,
		"origQty": "2",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"status": "CANCELED",
		"symbol": "BTCUSD_200925",
		"pair": "BTCUSD",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1571110484038
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	origClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelOrderService().
		Symbol(symbol).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("myOrder1", res.ClientOrderID)
	s.r().Equal("CANCELED", res.Status)
}
