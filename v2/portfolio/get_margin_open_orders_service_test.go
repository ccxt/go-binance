package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type getMarginOpenOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestGetMarginOpenOrdersService(t *testing.T) {
	suite.Run(t, new(getMarginOpenOrdersServiceTestSuite))
}

func (s *getMarginOpenOrdersServiceTestSuite) TestGetOpenOrders() {
	data := []byte(`[
		{
			"clientOrderId": "qhcZw71gAkCCTv0t0k8LUK",
			"cummulativeQuoteQty": "0.00000000",
			"executedQty": "0.00000000",
			"icebergQty": "0.00000000",
			"isWorking": true,
			"orderId": 211842552,
			"origQty": "0.30000000",
			"price": "0.00475010",
			"side": "SELL",
			"status": "NEW",
			"stopPrice": "0.00000000",
			"symbol": "BNBBTC",
			"time": 1562040170089,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1562040170089,
			"accountId": 152950866,
			"selfTradePreventionMode": "EXPIRE_TAKER",
			"preventedMatchId": null,
			"preventedQuantity": null
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewGetMarginOpenOrdersService().Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal("BNBBTC", orders[0].Symbol)
	s.r().Equal("qhcZw71gAkCCTv0t0k8LUK", orders[0].ClientOrderID)
	s.r().Equal(int64(211842552), orders[0].OrderID)
	s.r().Equal(SideTypeSell, orders[0].Side)
	s.r().Equal(OrderTypeLimit, orders[0].Type)
	s.r().Equal("NEW", orders[0].Status)
}

func (s *getMarginOpenOrdersServiceTestSuite) TestGetAllOpenOrders() {
	data := []byte(`[
		{
			"clientOrderId": "order1",
			"symbol": "BNBBTC",
			"orderId": 1,
			"status": "NEW"
		},
		{
			"clientOrderId": "order2",
			"symbol": "ETHBTC",
			"orderId": 2,
			"status": "NEW"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewGetMarginOpenOrdersService().
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 2)
	s.r().Equal("BNBBTC", orders[0].Symbol)
	s.r().Equal("ETHBTC", orders[1].Symbol)
}
