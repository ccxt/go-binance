package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umOrderServiceTestSuite struct {
	baseTestSuite
}

func TestUMOrderService(t *testing.T) {
	suite.Run(t, new(umOrderServiceTestSuite))
}

func (s *umOrderServiceTestSuite) TestUMOrder() {
	data := []byte(`{
		"clientOrderId": "testOrder",
		"cumQty": "0.0",
		"cumQuote": "0.0",
		"executedQty": "0.0",
		"orderId": 123456,
		"avgPrice": "0.0",
		"origQty": "1.0",
		"price": "2000",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "LONG",
		"status": "NEW",
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 1625097600000,
		"updateTime": 1625097500000,
		"priceMatch": "NONE"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	quantity := "1.0"
	price := "2000"
	priceMatch := PriceMatchTypeNone
	selfTradePreventionMode := SelfTradePreventionModeNone
	goodTillDate := int64(1625097600000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("side", side)
		e.setParam("type", orderType)
		e.setParam("timeInForce", timeInForce)
		e.setParam("quantity", quantity)
		e.setParam("price", price)
		e.setParam("priceMatch", priceMatch)
		e.setParam("selfTradePreventionMode", selfTradePreventionMode)
		e.setParam("goodTillDate", goodTillDate)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMOrderService().Symbol(symbol).
		Side(side).
		Type(orderType).
		TimeInForce(timeInForce).
		Quantity(quantity).
		Price(price).
		PriceMatch(priceMatch).
		SelfTradePreventionMode(selfTradePreventionMode).
		GoodTillDate(goodTillDate).
		Do(newContext())

	s.r().NoError(err)
	e := &UMOrder{
		ClientOrderID:           "testOrder",
		CumQty:                  "0.0",
		CumQuote:                "0.0",
		ExecutedQty:             "0.0",
		OrderID:                 123456,
		AvgPrice:                "0.0",
		OrigQty:                 "1.0",
		Price:                   "2000",
		ReduceOnly:              false,
		Side:                    SideTypeBuy,
		PositionSide:            PositionSideTypeLong,
		Status:                  "NEW",
		Symbol:                  "BTCUSDT",
		TimeInForce:             TimeInForceTypeGTC,
		Type:                    OrderTypeLimit,
		SelfTradePreventionMode: SelfTradePreventionModeNone,
		GoodTillDate:            1625097600000,
		UpdateTime:              1625097500000,
		PriceMatch:              PriceMatchTypeNone,
	}
	s.assertOrderEqual(e, res)
}

func (s *umOrderServiceTestSuite) assertOrderEqual(e, a *UMOrder) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CumQty, a.CumQty, "CumQty")
	r.Equal(e.CumQuote, a.CumQuote, "CumQuote")
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
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.SelfTradePreventionMode, a.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(e.GoodTillDate, a.GoodTillDate, "GoodTillDate")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.PriceMatch, a.PriceMatch, "PriceMatch")
}
