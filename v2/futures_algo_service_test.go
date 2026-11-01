package binance

import (
	"testing"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/stretchr/testify/suite"
)

type baseFuturesAlgoOrderTestSuite struct {
	baseTestSuite
}

type futuresAlgoOrderTestSuite struct {
	baseFuturesAlgoOrderTestSuite
}

func TestFuturesAlgoOrderService(t *testing.T) {
	suite.Run(t, new(futuresAlgoOrderTestSuite))
}

func (s *baseFuturesAlgoOrderTestSuite) assertFuturesAlgoOrderEqual(e, a *FuturesAlgoOrder) {
	r := s.r()
	r.Equal(e.AlgoId, a.AlgoId, "AlgoId")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.TotalQuantity, a.TotalQuantity, "TotalQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.ExecutedAmount, a.ExecutedAmount, "ExecutedAmount")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.ClientAlgoId, a.ClientAlgoId, "ClientAlgoId")
	r.Equal(e.BookTime, a.BookTime, "BookTime")
	r.Equal(e.EndTime, a.EndTime, "EndTime")
	r.Equal(e.AlgoStatus, a.AlgoStatus, "AlgoStatus")
	r.Equal(e.AlgoType, a.AlgoType, "AlgoType")
	r.Equal(e.Urgency, a.Urgency, "Urgency")
}

func (s *baseFuturesAlgoOrderTestSuite) assertFuturesAlgoSubOrderEqual(a, e *FuturesAlgoSubOrder) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.AlgoId, a.AlgoId, "AlgoId")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.OrderStatus, a.OrderStatus, "OrderStatus")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.ExecutedAmount, a.ExecutedAmount, "ExecutedAmount")
	r.Equal(e.FeeAmount, a.FeeAmount, "FeeAmount")
	r.Equal(e.FeeAsset, a.FeeAsset, "FeeAsset")
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.BookTime, a.BookTime, "BookTime")
	r.Equal(e.SubId, a.SubId, "SubId")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.OriginQuantity, a.OriginQuantity, "OriginQuantity")
}

func (s *futuresAlgoOrderTestSuite) TestCreateVpOrder() {
	data := []byte(`{
		"clientAlgoId": "12345678",
		"success": true, 
		"code": 0,
		"msg": "OK"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	side := SideTypeBuy
	positionSide := futures.PositionSideTypeLong
	quantity := 1.1
	urgency := FuturesAlgoUrgencyTypeMedium
	clientAlgoId := "12345678"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":       symbol,
			"side":         string(side),
			"positionSide": string(positionSide),
			"quantity":     quantity,
			"urgency":      string(urgency),
			"clientAlgoId": clientAlgoId,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateFuturesAlgoVpOrderService().Symbol(symbol).Side(side).PositionSide(positionSide).Quantity(quantity).
		Urgency(urgency).ClientAlgoId(clientAlgoId).Do(newContext())
	s.r().NoError(err)
	e := &CreateFuturesAlgoOrderResponse{
		ClientAlgoId: "12345678",
		Success:      true,
		Code:         0,
		Msg:          "OK",
	}
	s.assertCreateAlgoOrderResponseEqual(e, res)
}

func (s *futuresAlgoOrderTestSuite) TestCreateTwapOrder() {
	data := []byte(`{
		"clientAlgoId": "12345678",
		"success": true, 
		"code": 0,
		"msg": "OK"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	side := SideTypeBuy
	positionSide := futures.PositionSideTypeLong
	quantity := 1.1
	duration := int64(60)
	clientAlgoId := "12345678"
	limitPrice := 60000.0
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":       symbol,
			"side":         string(side),
			"positionSide": string(positionSide),
			"quantity":     quantity,
			"duration":     duration,
			"clientAlgoId": clientAlgoId,
			"limitPrice":   limitPrice,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateFuturesAlgoTwapOrderService().Symbol(symbol).Side(side).PositionSide(positionSide).Quantity(quantity).
		Duration(duration).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Do(newContext())
	s.r().NoError(err)
	e := &CreateFuturesAlgoOrderResponse{
		ClientAlgoId: "12345678",
		Success:      true,
		Code:         0,
		Msg:          "OK",
	}
	s.assertCreateAlgoOrderResponseEqual(e, res)
}

func (s *baseFuturesAlgoOrderTestSuite) assertCreateAlgoOrderResponseEqual(e, a *CreateFuturesAlgoOrderResponse) {
	r := s.r()
	r.Equal(e.ClientAlgoId, a.ClientAlgoId, "ClientAlgoId")
	r.Equal(e.Success, a.Success, "Success")
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
}

func (s *futuresAlgoOrderTestSuite) TestListOpenOrders() {
	data := []byte(`{
		"total": 1,
		"orders": [
			{ 
				"algoId": 14517,
				"symbol": "ETHUSDT",
				"side": "SELL",
				"positionSide": "SHORT",
				"totalQty": "5.000",
				"executedQty": "0.000", 
				"executedAmt": "0.00000000",
				"avgPrice": "0.00",
				"clientAlgoId": "d7096549481642f8a0bb69e9e2e31f2e",
				"bookTime": 1649756817004,
				"endTime": 0,
				"algoStatus": "WORKING",
				"algoType": "VP",
				"urgency": "LOW"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListOpenFuturesAlgoOrdersService().Do(newContext())
	s.r().NoError(err)
	e := &ListOpenFuturesAlgoOrdersResponse{
		Total: 1,
		Orders: []*FuturesAlgoOrder{
			{
				AlgoId:           14517,
				Symbol:           "ETHUSDT",
				Side:             SideTypeSell,
				PositionSide:     futures.PositionSideTypeShort,
				TotalQuantity:    "5.000",
				ExecutedQuantity: "0.000",
				ExecutedAmount:   "0.00000000",
				AvgPrice:         "0.00",
				ClientAlgoId:     "d7096549481642f8a0bb69e9e2e31f2e",
				BookTime:         1649756817004,
				EndTime:          0,
				AlgoStatus:       FuturesAlgoOrderStatusTypeWorking,
				AlgoType:         FuturesAlgoTypeVp,
				Urgency:          FuturesAlgoUrgencyTypeLow,
			},
		},
	}
	s.assertListAlgoOrdersResponseEqual(e, res)
}

func (s *baseFuturesAlgoOrderTestSuite) assertListAlgoOrdersResponseEqual(e, a *ListOpenFuturesAlgoOrdersResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	for i := range e.Orders {
		s.assertFuturesAlgoOrderEqual(e.Orders[i], a.Orders[i])
	}
}

func (s *futuresAlgoOrderTestSuite) TestGetHistoryOrders() {
	data := []byte(`{
		"total": 1,
		"orders": [
			{
				"algoId": 14518,
				"symbol": "BNBUSDT",
				"side": "BUY",
				"positionSide": "BOTH", 
				"totalQty": "100.00",   
				"executedQty": "0.00",  
				"executedAmt": "0.00000000", 
				"avgPrice": "3000.000",
				"clientAlgoId": "acacab56b3c44bef9f6a8f8ebd2a8408",
				"bookTime": 1649757019503,   
				"endTime": 1649757088101,    
				"algoStatus": "CANCELLED",   
				"algoType": "VP",            
				"urgency": "LOW"             
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BNBUSDT"
	side := SideTypeBuy
	startTime := int64(1649757019503)
	endTime := int64(1649757088101)
	page := 1
	pageSize := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"side":      string(side),
			"startTime": startTime,
			"endTime":   endTime,
			"page":      page,
			"pageSize":  pageSize,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListHistoryFuturesAlgoOrdersService().Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).
		Page(page).PageSize(pageSize).Do(newContext())
	s.r().NoError(err)
	e := &ListHistoryFuturesAlgoOrdersResponse{
		Total: 1,
		Orders: []*FuturesAlgoOrder{
			{
				AlgoId:           14518,
				Symbol:           "BNBUSDT",
				Side:             SideTypeBuy,
				PositionSide:     futures.PositionSideTypeBoth,
				TotalQuantity:    "100.00",
				ExecutedQuantity: "0.00",
				ExecutedAmount:   "0.00000000",
				AvgPrice:         "3000.000",
				ClientAlgoId:     "acacab56b3c44bef9f6a8f8ebd2a8408",
				BookTime:         1649757019503,
				EndTime:          1649757088101,
				AlgoStatus:       FuturesAlgoOrderStatusTypeCancelled,
				AlgoType:         FuturesAlgoTypeVp,
				Urgency:          FuturesAlgoUrgencyTypeLow,
			},
		},
	}
	s.assertListHistoryAlgoOrdersResponseEqual(e, res)
}

func (s *baseFuturesAlgoOrderTestSuite) assertListHistoryAlgoOrdersResponseEqual(e, a *ListHistoryFuturesAlgoOrdersResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	for i := range e.Orders {
		s.assertFuturesAlgoOrderEqual(e.Orders[i], a.Orders[i])
	}
}

func (s *futuresAlgoOrderTestSuite) TestCancelOrder() {
	data := []byte(`{
		"algoId": 14511,
		"success": true,
		"code": 0,
		"msg": "OK"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	algoId := int64(14511)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"algoId": algoId,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCancelFuturesAlgoOrderService().AlgoId(algoId).Do(newContext())
	s.r().NoError(err)
	e := &CancelFuturesAlgoOrderResponse{
		AlgoId:  14511,
		Success: true,
		Code:    0,
		Msg:     "OK",
	}
	s.assertCancelAlgoOrdersResponseEqual(e, res)
}

func (s *baseFuturesAlgoOrderTestSuite) assertCancelAlgoOrdersResponseEqual(e, a *CancelFuturesAlgoOrderResponse) {
	r := s.r()
	r.Equal(e.AlgoId, a.AlgoId, "AlgoId")
	r.Equal(e.Success, a.Success, "Success")
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
}

func (s *futuresAlgoOrderTestSuite) TestGetSubOrders() {
	data := []byte(`{
		"total": 1,
		"executedQty": "1.000",
		"executedAmt": "3229.44000000",
		"subOrders": [
			{
				"algoId": 13723,
				"orderId": 8389765519993908929, 
				"orderStatus": "FILLED", 
				"executedQty": "1.000",
				"executedAmt": "3229.44000000", 
				"feeAmt": "-1.61471999", 
				"feeAsset": "USDT",
				"bookTime": 1649319001964,
				"avgPrice": "3229.44",
				"side": "SELL", 
				"symbol": "ETHUSDT",
				"subId": 1, 
				"timeInForce": "IOC", 
				"origQty": "1.000"           
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	algoId := int64(14511)
	page := 1
	pageSize := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"algoId":   algoId,
			"page":     page,
			"pageSize": pageSize,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetFuturesAlgoSubOrdersService().AlgoId(algoId).Page(page).PageSize(pageSize).Do(newContext())
	s.r().NoError(err)
	e := &GetFuturesAlgoSubOrdersResponse{
		Total:            1,
		ExecutedQuantity: "1.000",
		ExecutedAmount:   "3229.44000000",
		SubOrders: []*FuturesAlgoSubOrder{
			{
				AlgoId:           13723,
				OrderId:          8389765519993908929,
				OrderStatus:      OrderStatusTypeFilled,
				ExecutedQuantity: "1.000",
				ExecutedAmount:   "3229.44000000",
				FeeAmount:        "-1.61471999",
				FeeAsset:         "USDT",
				BookTime:         1649319001964,
				AvgPrice:         "3229.44",
				Side:             SideTypeSell,
				Symbol:           "ETHUSDT",
				SubId:            1,
				TimeInForce:      TimeInForceTypeIOC,
				OriginQuantity:   "1.000",
			},
		},
	}
	s.assertGetSubOrdersResponseEqual(e, res)
}

func (s *baseFuturesAlgoOrderTestSuite) assertGetSubOrdersResponseEqual(e, a *GetFuturesAlgoSubOrdersResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.ExecutedAmount, a.ExecutedAmount, "ExecutedAmount")
	for i := range e.SubOrders {
		s.assertFuturesAlgoSubOrderEqual(e.SubOrders[i], a.SubOrders[i])
	}
}
