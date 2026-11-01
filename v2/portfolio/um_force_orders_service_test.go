package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umForceOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestUMForceOrdersService(t *testing.T) {
	suite.Run(t, new(umForceOrdersServiceTestSuite))
}

func (s *umForceOrdersServiceTestSuite) TestForceOrders() {
	data := []byte(`[
		{
			"orderId": 6071832819,
			"symbol": "BTCUSDT",
			"status": "FILLED",
			"clientOrderId": "autoclose-1596107620040000020",
			"price": "10871.09",
			"avgPrice": "10913.21000",
			"origQty": "0.001",
			"executedQty": "0.001",
			"cumQuote": "10.91321",
			"timeInForce": "IOC",
			"type": "LIMIT",
			"reduceOnly": false,
			"side": "SELL",
			"positionSide": "BOTH",
			"origType": "LIMIT",
			"time": 1596107620044,
			"updateTime": 1596107620087
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	autoCloseType := "LIQUIDATION"
	limit := 50
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":        symbol,
			"autoCloseType": autoCloseType,
			"limit":         limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMForceOrdersService().
		Symbol(symbol).
		AutoCloseType(autoCloseType).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(6071832819), orders[0].OrderID)
	s.r().Equal("BTCUSDT", orders[0].Symbol)
	s.r().Equal("FILLED", orders[0].Status)
	s.r().Equal("autoclose-1596107620040000020", orders[0].ClientOrderID)
}

func (s *umForceOrdersServiceTestSuite) TestForceOrdersWithAllParams() {
	data := []byte(`[
		{
			"orderId": 6071832819,
			"symbol": "BTCUSDT",
			"status": "FILLED",
			"clientOrderId": "autoclose-1596107620040000020",
			"price": "10871.09",
			"avgPrice": "10913.21000",
			"origQty": "0.001",
			"executedQty": "0.001",
			"cumQuote": "10.91321",
			"timeInForce": "IOC",
			"type": "LIMIT",
			"reduceOnly": false,
			"side": "SELL",
			"positionSide": "BOTH",
			"origType": "LIMIT",
			"time": 1596107620044,
			"updateTime": 1596107620087
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	autoCloseType := "ADL"
	startTime := int64(1596107620000)
	endTime := int64(1596107620999)
	limit := 50
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":        symbol,
			"autoCloseType": autoCloseType,
			"startTime":     startTime,
			"endTime":       endTime,
			"limit":         limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewUMForceOrdersService().
		Symbol(symbol).
		AutoCloseType(autoCloseType).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
}
