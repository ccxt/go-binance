package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginOCOServiceTestSuite struct {
	baseTestSuite
}

func TestMarginOCOService(t *testing.T) {
	suite.Run(t, new(marginOCOServiceTestSuite))
}

func (s *marginOCOServiceTestSuite) TestMarginOCO() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "JYVpp3F0f5CAG15DhtrqLp",
		"transactionTime": 1563417480525,
		"symbol": "LTCBTC",
		"marginBuyBorrowAmount": "5",
		"marginBuyBorrowAsset": "BTC",
		"orders": [
			{
				"symbol": "LTCBTC",
				"orderId": 2,
				"clientOrderId": "Kk7sqHb9J6mJWTMDVW7Vos"
			},
			{
				"symbol": "LTCBTC",
				"orderId": 3,
				"clientOrderId": "xTXKaGYd4bluPVp78IVRvl"
			}
		],
		"orderReports": [
			{
				"symbol": "LTCBTC",
				"orderId": 2,
				"orderListId": 0,
				"clientOrderId": "Kk7sqHb9J6mJWTMDVW7Vos",
				"transactTime": 1563417480525,
				"price": "0.000000",
				"origQty": "0.624363",
				"executedQty": "0.000000",
				"cummulativeQuoteQty": "0.000000",
				"status": "NEW",
				"timeInForce": "GTC",
				"type": "STOP_LOSS",
				"side": "BUY",
				"stopPrice": "0.960664"
			},
			{
				"symbol": "LTCBTC",
				"orderId": 3,
				"orderListId": 0,
				"clientOrderId": "xTXKaGYd4bluPVp78IVRvl",
				"transactTime": 1563417480525,
				"price": "0.036435",
				"origQty": "0.624363",
				"executedQty": "0.000000",
				"cummulativeQuoteQty": "0.000000",
				"status": "NEW",
				"timeInForce": "GTC",
				"type": "LIMIT_MAKER",
				"side": "BUY"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	side := SideTypeBuy
	quantity := "0.624363"
	price := "0.036435"
	stopPrice := "0.960664"
	listClientOrderID := "JYVpp3F0f5CAG15DhtrqLp"

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("side", side)
		e.setParam("quantity", quantity)
		e.setParam("price", price)
		e.setParam("stopPrice", stopPrice)
		e.setParam("listClientOrderId", listClientOrderID)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginOCOService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		StopPrice(stopPrice).
		ListClientOrderID(listClientOrderID).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(int64(0), res.OrderListID)
	s.r().Equal("OCO", res.ContingencyType)
	s.r().Equal("EXEC_STARTED", res.ListStatusType)
	s.r().Equal("EXECUTING", res.ListOrderStatus)
	s.r().Equal(listClientOrderID, res.ListClientOrderID)
	s.r().Equal(int64(1563417480525), res.TransactionTime)
	s.r().Equal(symbol, res.Symbol)
	s.r().Equal("5", res.MarginBuyBorrowAmount)
	s.r().Equal("BTC", res.MarginBuyBorrowAsset)
	s.r().Len(res.Orders, 2)
	s.r().Len(res.OrderReports, 2)
}
