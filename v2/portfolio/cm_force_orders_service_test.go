package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmForceOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMForceOrdersService(t *testing.T) {
	suite.Run(t, new(cmForceOrdersServiceTestSuite))
}

func (s *cmForceOrdersServiceTestSuite) TestForceOrders() {
	data := []byte(`[
		{
			"orderId": 165123080,
			"symbol": "BTCUSD_200925",
			"pair": "BTCUSD",
			"status": "FILLED",
			"clientOrderId": "autoclose-1596542005017000006",
			"price": "11326.9",
			"avgPrice": "11326.9",
			"origQty": "1",
			"executedQty": "1",
			"cumBase": "0.00882854",
			"timeInForce": "IOC",
			"type": "LIMIT",
			"reduceOnly": false,
			"side": "SELL",
			"positionSide": "BOTH",
			"origType": "LIMIT",
			"time": 1596542005019,
			"updateTime": 1596542005050
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
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

	orders, err := s.client.NewCMForceOrdersService().
		Symbol(symbol).
		AutoCloseType(autoCloseType).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(165123080), orders[0].OrderID)
	s.r().Equal("BTCUSD_200925", orders[0].Symbol)
	s.r().Equal("BTCUSD", orders[0].Pair)
	s.r().Equal("FILLED", orders[0].Status)
}

func (s *cmForceOrdersServiceTestSuite) TestForceOrdersWithAllParams() {
	data := []byte(`[
		{
			"orderId": 165123080,
			"symbol": "BTCUSD_200925",
			"pair": "BTCUSD",
			"status": "FILLED",
			"clientOrderId": "autoclose-1596542005017000006",
			"price": "11326.9",
			"avgPrice": "11326.9",
			"origQty": "1",
			"executedQty": "1",
			"cumBase": "0.00882854",
			"timeInForce": "IOC",
			"type": "LIMIT",
			"reduceOnly": false,
			"side": "SELL",
			"positionSide": "BOTH",
			"origType": "LIMIT",
			"time": 1596542005019,
			"updateTime": 1596542005050
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	autoCloseType := "ADL"
	startTime := int64(1596542005000)
	endTime := int64(1596542005999)
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

	orders, err := s.client.NewCMForceOrdersService().
		Symbol(symbol).
		AutoCloseType(autoCloseType).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
}
