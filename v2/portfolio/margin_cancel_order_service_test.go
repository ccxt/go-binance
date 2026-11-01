package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginCancelOrderServiceTestSuite struct {
	baseTestSuite
}

func TestMarginCancelOrderService(t *testing.T) {
	suite.Run(t, new(marginCancelOrderServiceTestSuite))
}

func (s *marginCancelOrderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 28,
		"origClientOrderId": "myOrder1",
		"clientOrderId": "cancelMyOrder1",
		"price": "1.00000000",
		"origQty": "10.00000000",
		"executedQty": "8.00000000",
		"cummulativeQuoteQty": "8.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "SELL"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(28)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginCancelOrderService().
		Symbol(symbol).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(28), res.OrderID)
	s.r().Equal("CANCELED", res.Status)
	s.r().Equal("LTCBTC", res.Symbol)
	s.r().Equal("SELL", res.Side)
}

func (s *marginCancelOrderServiceTestSuite) TestCancelOrderWithClientOrderID() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 28,
		"origClientOrderId": "myOrder1",
		"clientOrderId": "cancelMyOrder1",
		"price": "1.00000000",
		"origQty": "10.00000000",
		"executedQty": "8.00000000",
		"cummulativeQuoteQty": "8.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "SELL"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	origClientOrderID := "myOrder1"
	newClientOrderID := "cancelMyOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"origClientOrderId": origClientOrderID,
			"newClientOrderId":  newClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginCancelOrderService().
		Symbol(symbol).
		OrigClientOrderID(origClientOrderID).
		NewClientOrderID(newClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("myOrder1", res.OrigClientOrderID)
	s.r().Equal("cancelMyOrder1", res.ClientOrderID)
	s.r().Equal("CANCELED", res.Status)
}
