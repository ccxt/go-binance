package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umModifyOrderServiceTestSuite struct {
	baseTestSuite
}

func TestUMModifyOrderService(t *testing.T) {
	suite.Run(t, new(umModifyOrderServiceTestSuite))
}

func (s *umModifyOrderServiceTestSuite) TestModifyOrder() {
	data := []byte(`{
		"orderId": 20072994037,
		"symbol": "BTCUSDT",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumQuote": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "LONG",
		"origType": "LIMIT",
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0,
		"updateTime": 1629182711600,
		"priceMatch": "NONE"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
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

	res, err := s.client.NewUMModifyOrderService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(20072994037), res.OrderID)
	s.r().Equal("BTCUSDT", res.Symbol)
	s.r().Equal("NEW", res.Status)
	s.r().Equal("BUY", res.Side)
}

func (s *umModifyOrderServiceTestSuite) TestModifyOrderWithPriceMatch() {
	data := []byte(`{
		"orderId": 20072994037,
		"symbol": "BTCUSDT",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumQuote": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "LONG",
		"origType": "LIMIT",
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0,
		"updateTime": 1629182711600,
		"priceMatch": "OPPONENT"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	side := SideTypeBuy
	quantity := "1"
	priceMatch := "OPPONENT"
	origClientOrderID := "LJ9R4QZDihCaS8UAOOLpgW"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"side":              side,
			"quantity":          quantity,
			"priceMatch":        priceMatch,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMModifyOrderService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		PriceMatch(priceMatch).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("OPPONENT", res.PriceMatch)
	s.r().Equal("LJ9R4QZDihCaS8UAOOLpgW", res.ClientOrderID)
}
