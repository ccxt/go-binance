package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type getMarginOrderServiceTestSuite struct {
	baseTestSuite
}

func TestGetMarginOrderService(t *testing.T) {
	suite.Run(t, new(getMarginOrderServiceTestSuite))
}

func (s *getMarginOrderServiceTestSuite) TestGetOpenOrders() {
	data := []byte(`
		{
		  "clientOrderId": "ZwfQzuDIGpceVhKW5DvCmO",
		  "cummulativeQuoteQty": "0.00000000",
		  "executedQty": "0.00000000",
		  "icebergQty": "0.00000000",
		  "isWorking": true,
		  "orderId": 213205622,
		  "origQty": "0.30000000",
		  "price": "0.00493630",
		  "side": "SELL",
		  "status": "NEW",
		  "stopPrice": "0.00000000",
		  "symbol": "BNBBTC",
		  "time": 1562133008725,
		  "timeInForce": "GTC",
		  "type": "LIMIT",
		  "updateTime": 1562133008725,
		  "accountId": 152950866,
		  "selfTradePreventionMode": "EXPIRE_TAKER",
		  "preventedMatchId": null,
		  "preventedQuantity": null
		}`,
	)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	order, err := s.client.NewGetMarginOrderService().Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("BNBBTC", order.Symbol)
	s.r().Equal("ZwfQzuDIGpceVhKW5DvCmO", order.ClientOrderID)
	s.r().Equal(int64(213205622), order.OrderID)
	s.r().Equal(SideTypeSell, order.Side)
	s.r().Equal(OrderTypeLimit, order.Type)
	s.r().Equal("NEW", order.Status)
}
