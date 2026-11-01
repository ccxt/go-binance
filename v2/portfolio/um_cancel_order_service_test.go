package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umCancelOrderServiceTestSuite struct {
	baseTestSuite
}

func TestUMCancelOrderService(t *testing.T) {
	suite.Run(t, new(umCancelOrderServiceTestSuite))
}

func (s *umCancelOrderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"avgPrice": "0.00000",
		"clientOrderId": "myOrder1",
		"cumQty": "0",
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 4611875134427365377,
		"origQty": "0.40",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"status": "CANCELED",
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1571110484038,
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0,
		"priceMatch": "NONE"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderID := int64(4611875134427365377)
	origClientOrderID := "myOrder1"

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("orderId", orderID)
		e.setParam("origClientOrderId", origClientOrderID)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMCancelOrderService().
		Symbol(symbol).
		OrderID(orderID).
		OrigClientOrderID(origClientOrderID).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(symbol, res.Symbol)
	s.r().Equal(orderID, res.OrderID)
	s.r().Equal(origClientOrderID, res.ClientOrderID)
	s.r().Equal("CANCELED", res.Status)
	s.r().Equal(SideTypeBuy, res.Side)
	s.r().Equal(PositionSideTypeShort, res.PositionSide)
	s.r().Equal(OrderTypeLimit, res.Type)
	s.r().Equal(TimeInForceTypeGTC, res.TimeInForce)
}
