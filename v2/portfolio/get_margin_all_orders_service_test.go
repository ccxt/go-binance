package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type getMarginAllOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestGetMarginAllOrdersService(t *testing.T) {
	suite.Run(t, new(getMarginAllOrdersServiceTestSuite))
}

func (s *getMarginAllOrdersServiceTestSuite) TestGetAllOrders() {
	data := []byte(`[
		{
			"clientOrderId": "D2KDy4DIeS56PvkM13f8cP",
			"cummulativeQuoteQty": "0.00000000",
			"executedQty": "0.00000000",
			"icebergQty": "0.00000000",
			"isWorking": false,
			"orderId": 41295,
			"origQty": "5.31000000",
			"price": "0.22500000",
			"side": "SELL",
			"status": "CANCELED",
			"stopPrice": "0.18000000",
			"symbol": "BNBBTC",
			"time": 1565769338806,
			"timeInForce": "GTC",
			"type": "TAKE_PROFIT_LIMIT",
			"updateTime": 1565769342148,
			"accountId": 152950866,
			"selfTradePreventionMode": "EXPIRE_TAKER",
			"preventedMatchId": null,
			"preventedQuantity": null
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	orderID := int64(41295)
	limit := 500
	startTime := int64(1565769338806)
	endTime := int64(1565769342148)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"orderId":   orderID,
			"limit":     limit,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewGetMarginAllOrdersService().
		Symbol(symbol).
		OrderID(orderID).
		Limit(limit).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal("BNBBTC", orders[0].Symbol)
	s.r().Equal("D2KDy4DIeS56PvkM13f8cP", orders[0].ClientOrderID)
	s.r().Equal(int64(41295), orders[0].OrderID)
	s.r().Equal(SideTypeSell, orders[0].Side)
	s.r().Equal(OrderTypeTakeProfitLimit, orders[0].Type)
	s.r().Equal("CANCELED", orders[0].Status)
}
