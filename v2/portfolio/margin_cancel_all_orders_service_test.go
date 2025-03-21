package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginCancelAllOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestMarginCancelAllOrdersService(t *testing.T) {
	suite.Run(t, new(marginCancelAllOrdersServiceTestSuite))
}

func (s *marginCancelAllOrdersServiceTestSuite) TestCancelAllOrders() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"origClientOrderId": "E6APeyTJvkMvLMYMqu1KQ4",
			"orderId": 11,
			"orderListId": -1,
			"clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
			"price": "0.089853",
			"origQty": "0.178622",
			"executedQty": "0.000000",
			"cummulativeQuoteQty": "0.000000",
			"status": "CANCELED",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "BUY"
		},
		{
			"orderListId": 1929,
			"contingencyType": "OCO",
			"listStatusType": "ALL_DONE",
			"listOrderStatus": "ALL_DONE",
			"listClientOrderId": "2inzWQdDvZLHbbAmAozX2N",
			"transactionTime": 1585230948299,
			"symbol": "BTCUSDT",
			"orders": [
				{
					"symbol": "BTCUSDT",
					"orderId": 20,
					"clientOrderId": "CwOOIPHSmYywx6jZX77TdL"
				},
				{
					"symbol": "BTCUSDT",
					"orderId": 21,
					"clientOrderId": "461cPg51vQjV3zIMOXNz39"
				}
			],
			"orderReports": [
				{
					"symbol": "BTCUSDT",
					"origClientOrderId": "CwOOIPHSmYywx6jZX77TdL",
					"orderId": 20,
					"orderListId": 1929,
					"clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
					"price": "0.668611",
					"origQty": "0.690354",
					"executedQty": "0.000000",
					"cummulativeQuoteQty": "0.000000",
					"status": "CANCELED",
					"timeInForce": "GTC",
					"type": "STOP_LOSS_LIMIT",
					"side": "BUY",
					"stopPrice": "0.378131",
					"icebergQty": "0.017083"
				},
				{
					"symbol": "BTCUSDT",
					"origClientOrderId": "461cPg51vQjV3zIMOXNz39",
					"orderId": 21,
					"orderListId": 1929,
					"clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
					"price": "0.008791",
					"origQty": "0.690354",
					"executedQty": "0.000000",
					"cummulativeQuoteQty": "0.000000",
					"status": "CANCELED",
					"timeInForce": "GTC",
					"type": "LIMIT_MAKER",
					"side": "BUY",
					"icebergQty": "0.639962"
				}
			]
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginCancelAllOrdersService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 2)
	s.r().Equal("BTCUSDT", res[0].Symbol)
	s.r().Equal(int64(11), res[0].OrderID)
	s.r().Equal("CANCELED", res[0].Status)
	s.r().Equal(int64(1929), res[1].OrderListID)
	s.r().Equal("OCO", res[1].ContingencyType)
	s.r().Len(res[1].Orders, 2)
	s.r().Len(res[1].OrderReports, 2)
}
