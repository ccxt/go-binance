package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmOrderServiceTestSuite struct {
	baseTestSuite
}

func TestCMOrderService(t *testing.T) {
	suite.Run(t, new(cmOrderServiceTestSuite))
}

func (s *cmOrderServiceTestSuite) TestCMOrder() {
	data := []byte(`{
		"clientOrderId": "testOrder",
		"cumQty": "0",
		"cumBase": "0",
		"executedQty": "0",
		"orderId": 22542179,
		"avgPrice": "0.0",
		"origQty": "10",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"status": "NEW",
		"symbol": "BTCUSD_200925",
		"pair": "BTCUSD",
		"timeInForce": "GTC",
		"type": "MARKET",
		"updateTime": 1566818724722
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	side := SideTypeBuy
	orderType := OrderTypeMarket
	quantity := "10"

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("side", side)
		e.setParam("type", orderType)
		e.setParam("quantity", quantity)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMOrderService().
		Symbol(symbol).
		Side(side).
		Type(orderType).
		Quantity(quantity).
		Do(newContext())

	s.r().NoError(err)
	e := &CMOrder{
		ClientOrderID: "testOrder",
		CumQty:        "0",
		CumBase:       "0",
		ExecutedQty:   "0",
		OrderID:       22542179,
		AvgPrice:      "0.0",
		OrigQty:       "10",
		Price:         "0",
		ReduceOnly:    false,
		Side:          SideTypeBuy,
		PositionSide:  PositionSideTypeShort,
		Status:        "NEW",
		Symbol:        "BTCUSD_200925",
		Pair:          "BTCUSD",
		TimeInForce:   TimeInForceTypeGTC,
		Type:          OrderTypeMarket,
		UpdateTime:    1566818724722,
	}
	s.assertOrderEqual(e, res)
}

func (s *cmOrderServiceTestSuite) assertOrderEqual(e, a *CMOrder) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CumQty, a.CumQty, "CumQty")
	r.Equal(e.CumBase, a.CumBase, "CumBase")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.OrigQty, a.OrigQty, "OrigQty")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
