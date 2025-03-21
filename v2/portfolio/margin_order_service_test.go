package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginOrderServiceTestSuite struct {
	baseTestSuite
}

func TestMarginOrderService(t *testing.T) {
	suite.Run(t, new(marginOrderServiceTestSuite))
}

func (s *marginOrderServiceTestSuite) TestMarginOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"orderId": 28,
		"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
		"transactTime": 1507725176595,
		"price": "1.00000000",
		"origQty": "10.00000000",
		"executedQty": "10.00000000",
		"cummulativeQuoteQty": "10.00000000",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "MARKET",
		"side": "SELL",
		"marginBuyBorrowAmount": "5",
		"marginBuyBorrowAsset": "BTC",
		"fills": [
			{
				"price": "4000.00000000",
				"qty": "1.00000000",
				"commission": "4.00000000",
				"commissionAsset": "USDT"
			},
			{
				"price": "3999.00000000",
				"qty": "5.00000000",
				"commission": "19.99500000",
				"commissionAsset": "USDT"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	side := SideTypeSell
	orderType := OrderTypeMarket
	quantity := "10.00000000"

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("side", side)
		e.setParam("type", orderType)
		e.setParam("quantity", quantity)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginOrderService().
		Symbol(symbol).
		Side(side).
		Type(orderType).
		Quantity(quantity).
		Do(newContext())

	s.r().NoError(err)
	e := &MarginOrder{
		Symbol:                "BTCUSDT",
		OrderID:               28,
		ClientOrderID:         "6gCrw2kRUAF9CvJDGP16IP",
		TransactTime:          1507725176595,
		Price:                 "1.00000000",
		OrigQty:               "10.00000000",
		ExecutedQty:           "10.00000000",
		CummulativeQuoteQty:   "10.00000000",
		Status:                "FILLED",
		TimeInForce:           TimeInForceTypeGTC,
		Type:                  OrderTypeMarket,
		Side:                  SideTypeSell,
		MarginBuyBorrowAmount: "5",
		MarginBuyBorrowAsset:  "BTC",
		Fills: []*Fill{
			{
				Price:           "4000.00000000",
				Qty:             "1.00000000",
				Commission:      "4.00000000",
				CommissionAsset: "USDT",
			},
			{
				Price:           "3999.00000000",
				Qty:             "5.00000000",
				Commission:      "19.99500000",
				CommissionAsset: "USDT",
			},
		},
	}
	s.assertOrderEqual(e, res)
}

func (s *marginOrderServiceTestSuite) assertOrderEqual(e, a *MarginOrder) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQty, a.OrigQty, "OrigQty")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.CummulativeQuoteQty, a.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.MarginBuyBorrowAmount, a.MarginBuyBorrowAmount, "MarginBuyBorrowAmount")
	r.Equal(e.MarginBuyBorrowAsset, a.MarginBuyBorrowAsset, "MarginBuyBorrowAsset")
	r.Len(a.Fills, len(e.Fills))
	for idx := range e.Fills {
		r.Equal(e.Fills[idx].Price, a.Fills[idx].Price, "Fill.Price")
		r.Equal(e.Fills[idx].Qty, a.Fills[idx].Qty, "Fill.Qty")
		r.Equal(e.Fills[idx].Commission, a.Fills[idx].Commission, "Fill.Commission")
		r.Equal(e.Fills[idx].CommissionAsset, a.Fills[idx].CommissionAsset, "Fill.CommissionAsset")
	}
}
