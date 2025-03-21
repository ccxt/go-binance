package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginCancelOCOServiceTestSuite struct {
	baseTestSuite
}

func TestMarginCancelOCOService(t *testing.T) {
	suite.Run(t, new(marginCancelOCOServiceTestSuite))
}

func (s *marginCancelOCOServiceTestSuite) TestCancelOCO() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "ALL_DONE",
		"listOrderStatus": "ALL_DONE",
		"listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
		"transactionTime": 1574040868128,
		"symbol": "LTCBTC",
		"orders": [
			{
				"symbol": "LTCBTC",
				"orderId": 2,
				"clientOrderId": "pO9ufTiFGg3nw2fOdgeOXa"
			},
			{
				"symbol": "LTCBTC",
				"orderId": 3,
				"clientOrderId": "TXOvglzXuaubXAaENpaRCB"
			}
		],
		"orderReports": [
			{
				"symbol": "LTCBTC",
				"origClientOrderId": "pO9ufTiFGg3nw2fOdgeOXa",
				"orderId": 2,
				"orderListId": 0,
				"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
				"price": "1.00000000",
				"origQty": "10.00000000",
				"executedQty": "0.00000000",
				"cummulativeQuoteQty": "0.00000000",
				"status": "CANCELED",
				"timeInForce": "GTC",
				"type": "STOP_LOSS_LIMIT",
				"side": "SELL",
				"stopPrice": "1.00000000"
			},
			{
				"symbol": "LTCBTC",
				"origClientOrderId": "TXOvglzXuaubXAaENpaRCB",
				"orderId": 3,
				"orderListId": 0,
				"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
				"price": "3.00000000",
				"origQty": "10.00000000",
				"executedQty": "0.00000000",
				"cummulativeQuoteQty": "0.00000000",
				"status": "CANCELED",
				"timeInForce": "GTC",
				"type": "LIMIT_MAKER",
				"side": "SELL"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderListID := int64(0)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":      symbol,
			"orderListId": orderListID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginCancelOCOService().
		Symbol(symbol).
		OrderListID(orderListID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(0), res.OrderListID)
	s.r().Equal("OCO", res.ContingencyType)
	s.r().Equal("ALL_DONE", res.ListStatusType)
	s.r().Equal("LTCBTC", res.Symbol)
	s.r().Len(res.Orders, 2)
	s.r().Len(res.OrderReports, 2)
}

func (s *marginCancelOCOServiceTestSuite) TestCancelOCOWithListClientOrderID() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "ALL_DONE",
		"listOrderStatus": "ALL_DONE",
		"listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
		"transactionTime": 1574040868128,
		"symbol": "LTCBTC",
		"orders": [
			{
				"symbol": "LTCBTC",
				"orderId": 2,
				"clientOrderId": "pO9ufTiFGg3nw2fOdgeOXa"
			}
		],
		"orderReports": [
			{
				"symbol": "LTCBTC",
				"origClientOrderId": "pO9ufTiFGg3nw2fOdgeOXa",
				"orderId": 2,
				"orderListId": 0,
				"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
				"price": "1.00000000",
				"origQty": "10.00000000",
				"executedQty": "0.00000000",
				"cummulativeQuoteQty": "0.00000000",
				"status": "CANCELED",
				"timeInForce": "GTC",
				"type": "STOP_LOSS_LIMIT",
				"side": "SELL",
				"stopPrice": "1.00000000"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	listClientOrderID := "C3wyj4WVEktd7u9aVBRXcN"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"listClientOrderId": listClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginCancelOCOService().
		Symbol(symbol).
		ListClientOrderID(listClientOrderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("C3wyj4WVEktd7u9aVBRXcN", res.ListClientOrderID)
	s.r().Equal("ALL_DONE", res.ListStatusType)
}
