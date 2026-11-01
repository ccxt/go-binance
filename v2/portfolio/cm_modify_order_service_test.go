package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmModifyOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMModifyOrderService(t *testing.T) {
	suite.Run(t, new(cmModifyOrderServiceTestSuite))
}

func (s *cmModifyOrderServiceTestSuite) TestModifyOrder() {
	data := []byte(`{
		"orderId": 20072994037,
		"symbol": "BTCUSD_PERP",
		"pair": "BTCUSD",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumBase": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "LONG",
		"origType": "LIMIT",
		"updateTime": 1629182711600
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	side := SideTypeBuy
	quantity := "1"
	price := "30005"
	orderID := int64(20072994037)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":   symbol,
			"side":     side,
			"quantity": quantity,
			"price":    price,
			"orderId":  orderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMModifyOrderService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(20072994037), res.OrderID)
	s.r().Equal("BTCUSD_PERP", res.Symbol)
	s.r().Equal("BTCUSD", res.Pair)
	s.r().Equal("NEW", res.Status)
	s.r().Equal("BUY", res.Side)
}

func (s *cmModifyOrderServiceTestSuite) TestModifyOrderWithClientOrderID() {
	data := []byte(`{
		"orderId": 20072994037,
		"symbol": "BTCUSD_PERP",
		"pair": "BTCUSD",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumBase": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "LONG",
		"origType": "LIMIT",
		"updateTime": 1629182711600
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	side := SideTypeBuy
	quantity := "1"
	price := "30005"
	origClientOrderID := "LJ9R4QZDihCaS8UAOOLpgW"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"side":              side,
			"quantity":          quantity,
			"price":             price,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMModifyOrderService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("LJ9R4QZDihCaS8UAOOLpgW", res.ClientOrderID)
	s.r().Equal("BTCUSD", res.Pair)
}
