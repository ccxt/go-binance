package binance

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseOrderTestSuite struct {
	baseTestSuite
}

type orderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestOrderService(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (s *orderServiceTestSuite) TestCreateOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
		"origQuoteOrderQty": "10.00",
		"executedQty": "10.00",
		"cummulativeQuoteQty": "10.00",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	quantity := "12.00"
	quoteOrderQty := "10.00"
	price := "0.0001"
	newClientOrderID := "myOrder1"
	trailingDelta := "1000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"quantity":         quantity,
			"quoteOrderQty":    quoteOrderQty,
			"price":            price,
			"newClientOrderId": newClientOrderID,
			"trailingDelta":    trailingDelta,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).TrailingDelta(trailingDelta).Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		TransactTime:             1499827319559,
		Price:                    "0.0001",
		OrigQuantity:             "12.00",
		OrigQuoteOrderQuantity:   "10.00",
		ExecutedQuantity:         "10.00",
		CummulativeQuoteQuantity: "10.00",
		Status:                   OrderStatusTypeFilled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
	}
	s.assertCreateOrderResponseEqual(e, res)

	err = s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).TrailingDelta(trailingDelta).Test(newContext())
	s.r().NoError(err)
}

func (s *orderServiceTestSuite) TestCreateOrderId() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
		"origQuoteOrderQty": "10.00",
		"executedQty": "10.00",
		"cummulativeQuoteQty": "10.00",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	quantity := "12.00"
	quoteOrderQty := "10.00"
	price := "0.0001"
	// trailingDelta := "1000"

	s.assertReq(func(r *request) {
		s.Assertions.True(strings.HasPrefix(r.form.Get("newClientOrderId"), "x-B3AUXNYV"))
	})

	_, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).Do(newContext())
	s.r().NoError(err)
}

func (s *orderServiceTestSuite) TestCreateOrderFull() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
		"origQuoteOrderQty": "10.00",
		"executedQty": "10.00",
		"cummulativeQuoteQty": "10.00",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"fills": [
			{
				"price":"0.00002991",
				"qty":"344.00000000",
				"commission":"0.00332384",
				"commissionAsset":"BNB",
				"tradeId":1566397
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	quantity := "12.00"
	quoteOrderQty := "10.00"
	price := "0.0001"
	newClientOrderID := "myOrder1"
	newOrderRespType := NewOrderRespTypeFULL
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"quantity":         quantity,
			"quoteOrderQty":    quoteOrderQty,
			"price":            price,
			"newClientOrderId": newClientOrderID,
			"newOrderRespType": newOrderRespType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).
		NewOrderRespType(newOrderRespType).Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		TransactTime:             1499827319559,
		Price:                    "0.0001",
		OrigQuantity:             "12.00",
		OrigQuoteOrderQuantity:   "10.00",
		ExecutedQuantity:         "10.00",
		CummulativeQuoteQuantity: "10.00",
		Status:                   OrderStatusTypeFilled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
		Fills: []*Fill{
			{
				Price:           "0.00002991",
				Quantity:        "344.00000000",
				Commission:      "0.00332384",
				CommissionAsset: "BNB",
			},
		},
	}
	s.assertCreateOrderResponseEqual(e, res)

	err = s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).
		NewOrderRespType(newOrderRespType).Test(newContext())
	s.r().NoError(err)
}

func (s *baseOrderTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.OrigQuoteOrderQuantity, a.OrigQuoteOrderQuantity, "OrigQuoteOrderQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")

	r.Len(a.Fills, len(e.Fills))
	for idx, fill := range e.Fills {
		s.assertFillEqual(fill, a.Fills[idx])
	}
}

func (s *baseOrderTestSuite) assertFillEqual(e, a *Fill) {
	r := s.r()
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
}

func (s *orderServiceTestSuite) TestCreateOCO() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
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
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "STOP_LOSS",
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
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT_MAKER",
			"side": "SELL"
		  }
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	timeInForce := TimeInForceTypeGTC
	quantity := "10"
	price := "3"
	stopPrice := "3.1"
	stopLimitPrice := "3.2"
	limitClientOrderID := "myOrder1"
	newOrderRespType := NewOrderRespTypeFULL
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":               symbol,
			"side":                 side,
			"quantity":             quantity,
			"price":                price,
			"stopPrice":            stopPrice,
			"stopLimitPrice":       stopLimitPrice,
			"stopLimitTimeInForce": timeInForce,
			"limitClientOrderId":   limitClientOrderID,
			"newOrderRespType":     newOrderRespType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOCOService().
		Symbol(symbol).
		Side(side).
		Quantity(quantity).
		Price(price).
		StopPrice(stopPrice).
		StopLimitPrice(stopLimitPrice).
		StopLimitTimeInForce(timeInForce).
		LimitClientOrderID(limitClientOrderID).
		NewOrderRespType(newOrderRespType).
		Do(newContext())

	s.r().NoError(err)
	e := &CreateOCOResponse{
		OrderListID:       0,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "C3wyj4WVEktd7u9aVBRXcN",
		TransactionTime:   1574040868128,
		Symbol:            "LTCBTC",
		Orders: []*OCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       2,
				ClientOrderID: "pO9ufTiFGg3nw2fOdgeOXa",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       3,
				ClientOrderID: "TXOvglzXuaubXAaENpaRCB",
			},
		},
		OrderReports: []*OCOOrderReport{
			{
				Symbol:                   "LTCBTC",
				OrderID:                  2,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "1.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeNew,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeStopLoss,
				Side:                     SideTypeSell,
				StopPrice:                "1.00000000",
			},
			{
				Symbol:                   "LTCBTC",
				OrderID:                  3,
				OrderListID:              0,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
				Price:                    "3.00000000",
				OrigQuantity:             "10.00000000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeNew,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeLimitMaker,
				Side:                     SideTypeSell,
			},
		},
	}
	s.assertCreateOCOResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCreateOCOResponseEqual(e, a *CreateOCOResponse) {
	r := s.r()
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertOCOOrderReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertOCOOrderEqual(order, a.Orders[idx])
	}
}

func (s *baseOrderTestSuite) assertOCOOrderReportEqual(e, a *OCOOrderReport) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	// r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
}

func (s *baseOrderTestSuite) assertOCOOrderEqual(e, a *OCOOrder) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *orderServiceTestSuite) TestListOpenOco() {
	data := []byte(`[
		{
			"orderListId": 31,
			"contingencyType": "OCO",
			"listStatusType": "EXEC_STARTED",
			"listOrderStatus": "EXECUTING",
			"listClientOrderId": "wuB13fmulKj3YjdqWEcsnp",
			"transactionTime": 1565246080644,
			"symbol": "LTCBTC",
			"orders": [
			  {
				"symbol": "LTCBTC",
				"orderId": 4,
				"clientOrderId": "r3EH2N76dHfLoSZWIUw1bT"
			  },
			  {
				"symbol": "LTCBTC",
				"orderId": 5,
				"clientOrderId": "Cv1SnyPD3qhqpbjpYEHbd2"
			  }
			]
		  }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	ocos, err := s.client.NewListOpenOcoService().
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(ocos, 1)
	e := &Oco{
		Symbol:            "LTCBTC",
		OrderListId:       31,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "wuB13fmulKj3YjdqWEcsnp",
		TransactionTime:   1565246080644,
		Orders: []*Order{
			{
				Symbol:        "LTCBTC",
				OrderID:       4,
				ClientOrderID: "r3EH2N76dHfLoSZWIUw1bT",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       5,
				ClientOrderID: "Cv1SnyPD3qhqpbjpYEHbd2",
			},
		},
	}
	s.assertOcoEqual(e, ocos[0])
}

func (s *baseOrderTestSuite) assertOcoEqual(e, a *Oco) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListId, a.OrderListId, "OrderListId")
	r.Equal(e.Orders[0].Symbol, a.Orders[0].Symbol, "Orders[0].Symbol")
	r.Equal(e.Orders[0].OrderID, a.Orders[0].OrderID, "Orders[0].OrderID")
	r.Equal(e.Orders[0].ClientOrderID, a.Orders[0].ClientOrderID, "Orders[0].ClientOrderID")
	r.Equal(e.Orders[1].Symbol, a.Orders[1].Symbol, "Orders[1].Symbol")
	r.Equal(e.Orders[1].OrderID, a.Orders[1].OrderID, "Orders[1].OrderID")
	r.Equal(e.Orders[1].ClientOrderID, a.Orders[1].ClientOrderID, "Orders[1].ClientOrderID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
}

func (s *orderServiceTestSuite) TestListOpenOrders() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "orderId": 1,
            "clientOrderId": "myOrder1",
            "price": "0.1",
            "origQty": "1.0",
            "executedQty": "0.0",
            "status": "NEW",
            "timeInForce": "GTC",
            "type": "LIMIT",
            "side": "BUY",
            "stopPrice": "0.0",
            "icebergQty": "0.0",
			"time": 1499827319559,
			"updateTime": 1499827319559,
    		"isWorking": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	orders, err := s.client.NewListOpenOrdersService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &Order{
		Symbol:           "LTCBTC",
		OrderID:          1,
		ClientOrderID:    "myOrder1",
		Price:            "0.1",
		OrigQuantity:     "1.0",
		ExecutedQuantity: "0.0",
		Status:           OrderStatusTypeNew,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		Side:             SideTypeBuy,
		StopPrice:        "0.0",
		IcebergQuantity:  "0.0",
		Time:             1499827319559,
		UpdateTime:       1499827319559,
		IsWorking:        true,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *baseOrderTestSuite) assertOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.IcebergQuantity, a.IcebergQuantity, "IcebergQuantity")
	r.Equal(e.Time, e.Time, "Time")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.IsWorking, a.IsWorking, "IsWorking")
	r.Equal(e.OrigQuoteOrderQuantity, a.OrigQuoteOrderQuantity, "OrigQuoteOrderQuantity")
}

func (s *orderServiceTestSuite) TestGetOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"price": "0.1",
		"origQty": "1.0",
		"executedQty": "0.0",
		"cummulativeQuoteQty": "0.0",
		"status": "NEW",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"stopPrice": "0.0",
		"icebergQty": "0.0",
		"time": 1499827319559,
		"updateTime": 1499827319559,
		"isWorking": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(1)
	origClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Order{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		Price:                    "0.1",
		OrigQuantity:             "1.0",
		ExecutedQuantity:         "0.0",
		CummulativeQuoteQuantity: "0.0",
		Status:                   OrderStatusTypeNew,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
		StopPrice:                "0.0",
		IcebergQuantity:          "0.0",
		Time:                     1499827319559,
		UpdateTime:               1499827319559,
		IsWorking:                true,
	}
	s.assertOrderEqual(e, order)
}

func (s *orderServiceTestSuite) TestListOrders() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "orderId": 1,
            "clientOrderId": "myOrder1",
            "price": "0.1",
            "origQty": "1.0",
            "executedQty": "0.0",
            "status": "NEW",
            "timeInForce": "GTC",
            "type": "LIMIT",
            "side": "BUY",
            "stopPrice": "0.0",
            "icebergQty": "0.0",
			"time": 1499827319559,
			"updateTime": 1499827319559,
			"isWorking": true,
    		"origQuoteOrderQty": "0.000000"
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	orderID := int64(1)
	limit := 3
	startTime := int64(1499827319559)
	endTime := int64(1499827319560)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"orderId":   orderID,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewListOrdersService().Symbol(symbol).
		OrderID(orderID).StartTime(startTime).EndTime(endTime).
		Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &Order{
		Symbol:                 "LTCBTC",
		OrderID:                1,
		ClientOrderID:          "myOrder1",
		Price:                  "0.1",
		OrigQuantity:           "1.0",
		ExecutedQuantity:       "0.0",
		Status:                 OrderStatusTypeNew,
		TimeInForce:            TimeInForceTypeGTC,
		Type:                   OrderTypeLimit,
		Side:                   SideTypeBuy,
		StopPrice:              "0.0",
		IcebergQuantity:        "0.0",
		Time:                   1499827319559,
		UpdateTime:             1499827319559,
		IsWorking:              true,
		OrigQuoteOrderQuantity: "0.000000",
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *orderServiceTestSuite) TestCancelOCO() {
	data := []byte(`{
		"orderListId":1000,
		"contingencyType":"OCO",
		"listStatusType":"ALL_DONE",
		"listOrderStatus":"ALL_DONE",
		"listClientOrderId":"my-list-order-id",
		"transactionTime":1614272133000,
		"symbol":"BTCUSDT",
		"orders":[
			{"symbol":"BTCUSDT","orderId":1100,"clientOrderId":"stop-loss-order-id"},
			{"symbol":"BTCUSDT","orderId":1010,"clientOrderId":"limit-maker-order-id"}
		],
		"orderReports":[
			{
				"symbol":"BTCUSDT",
				"origClientOrderId":"stop-loss-order-id",
				"orderId":1100,
				"orderListId":1000,
				"clientOrderId":"cancel-request-id",
				"price":"50000.00000000",
				"origQty":"0.00030000",
				"executedQty":"0.00000000",
				"cummulativeQuoteQty":"0.00000000",
				"status":"CANCELED",
				"timeInForce":"GTC",
				"type":"STOP_LOSS_LIMIT",
				"side":"SELL",
				"stopPrice":"50000.00000000"
			},
			{
				"symbol":"BTCUSDT",
				"origClientOrderId":"limit-maker-order-id",
				"orderId":1010,
				"orderListId":1000,
				"clientOrderId":"cancel-request-id",
				"price":"52000.00000000",
				"origQty":"0.00030000",
				"executedQty":"0.00000000",
				"cummulativeQuoteQty":"0.00000000",
				"status":"CANCELED",
				"timeInForce":"GTC",
				"type":"LIMIT_MAKER",
				"side":"SELL"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	listClientOrderID := "my-list-order-id"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"listClientOrderId": listClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.
		NewCancelOCOService().
		Symbol(symbol).
		ListClientOrderID(listClientOrderID).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelOCOResponse{
		OrderListID:       1000,
		ContingencyType:   "OCO",
		ListStatusType:    "ALL_DONE",
		ListOrderStatus:   "ALL_DONE",
		ListClientOrderID: "my-list-order-id",
		TransactionTime:   1614272133000,
		Symbol:            "BTCUSDT",
		Orders: []*OCOOrder{
			{Symbol: "BTCUSDT", OrderID: 1100, ClientOrderID: "stop-loss-order-id"},
			{Symbol: "BTCUSDT", OrderID: 1010, ClientOrderID: "limit-maker-order-id"},
		},
		OrderReports: []*OCOOrderReport{
			{
				Symbol:                   "BTCUSDT",
				OrigClientOrderID:        "stop-loss-order-id",
				OrderID:                  1100,
				OrderListID:              1000,
				ClientOrderID:            "cancel-request-id",
				Price:                    "50000.00000000",
				OrigQuantity:             "0.00030000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeCanceled,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeStopLossLimit,
				Side:                     SideTypeSell,
				StopPrice:                "50000.00000000",
			},
			{
				Symbol:                   "BTCUSDT",
				OrigClientOrderID:        "limit-maker-order-id",
				OrderID:                  1010,
				OrderListID:              1000,
				ClientOrderID:            "cancel-request-id",
				Price:                    "52000.00000000",
				OrigQuantity:             "0.00030000",
				ExecutedQuantity:         "0.00000000",
				CummulativeQuoteQuantity: "0.00000000",
				Status:                   OrderStatusTypeCanceled,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeLimitMaker,
				Side:                     SideTypeSell,
			},
		},
	}
	s.assertCancelOCOResponseEqual(e, res)
}

func (s *orderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 28,
		"origClientOrderId": "myOrder1",
		"clientOrderId": "cancelMyOrder1",
		"transactTime": 1507725176595,
		"price": "1.00000000",
		"origQty": "10.00000000",
		"origQuoteOrderQty": "11.00000000",
		"executedQty": "8.00000000",
		"cummulativeQuoteQty": "8.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "SELL"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	orderID := int64(28)
	origClientOrderID := "myOrder1"
	newClientOrderID := "cancelMyOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
			"newClientOrderId":  newClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).
		NewClientOrderID(newClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  28,
		OrigClientOrderID:        "myOrder1",
		ClientOrderID:            "cancelMyOrder1",
		TransactTime:             1507725176595,
		Price:                    "1.00000000",
		OrigQuantity:             "10.00000000",
		OrigQuoteOrderQuantity:   "11.00000000",
		ExecutedQuantity:         "8.00000000",
		CummulativeQuoteQuantity: "8.00000000",
		Status:                   OrderStatusTypeCanceled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeSell,
	}
	s.assertCancelOrderResponseEqual(e, res)
}

func (s *orderServiceTestSuite) TestCancelOpenOrders() {
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

	res, err := s.client.NewCancelOpenOrdersService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(res.Orders, 1)
	eo := &CancelOrderResponse{
		Symbol:                   "BTCUSDT",
		OrigClientOrderID:        "E6APeyTJvkMvLMYMqu1KQ4",
		OrderID:                  11,
		ClientOrderID:            "pXLV6Hz6mprAcVYpVMTGgx",
		Price:                    "0.089853",
		OrigQuantity:             "0.178622",
		ExecutedQuantity:         "0.000000",
		CummulativeQuoteQuantity: "0.000000",
		Status:                   OrderStatusTypeCanceled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
		TransactTime:             0,
	}
	s.assertCancelOrderResponseEqual(eo, res.Orders[0])

	r.Len(res.OCOOrders, 1)
	eoco := &CancelOCOResponse{
		OrderListID:       1929,
		ContingencyType:   "OCO",
		ListStatusType:    "ALL_DONE",
		ListOrderStatus:   "ALL_DONE",
		ListClientOrderID: "2inzWQdDvZLHbbAmAozX2N",
		TransactionTime:   1585230948299,
		Symbol:            "BTCUSDT",
		Orders: []*OCOOrder{
			{
				Symbol:        "BTCUSDT",
				OrderID:       20,
				ClientOrderID: "CwOOIPHSmYywx6jZX77TdL",
			},
			{
				Symbol:        "BTCUSDT",
				OrderID:       21,
				ClientOrderID: "461cPg51vQjV3zIMOXNz39",
			},
		},
		OrderReports: []*OCOOrderReport{
			{
				Symbol:                   "BTCUSDT",
				OrigClientOrderID:        "CwOOIPHSmYywx6jZX77TdL",
				OrderID:                  20,
				OrderListID:              1929,
				ClientOrderID:            "pXLV6Hz6mprAcVYpVMTGgx",
				Price:                    "0.668611",
				OrigQuantity:             "0.690354",
				ExecutedQuantity:         "0.000000",
				CummulativeQuoteQuantity: "0.000000",
				Status:                   OrderStatusTypeCanceled,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeStopLossLimit,
				Side:                     SideTypeBuy,
				StopPrice:                "0.378131",
				IcebergQuantity:          "0.017083",
			},
			{
				Symbol:                   "BTCUSDT",
				OrigClientOrderID:        "461cPg51vQjV3zIMOXNz39",
				OrderID:                  21,
				OrderListID:              1929,
				ClientOrderID:            "pXLV6Hz6mprAcVYpVMTGgx",
				Price:                    "0.008791",
				OrigQuantity:             "0.690354",
				ExecutedQuantity:         "0.000000",
				CummulativeQuoteQuantity: "0.000000",
				Status:                   OrderStatusTypeCanceled,
				TimeInForce:              TimeInForceTypeGTC,
				Type:                     OrderTypeLimitMaker,
				Side:                     SideTypeBuy,
				IcebergQuantity:          "0.639962",
			},
		},
	}
	s.assertCancelOCOResponseEqual(eoco, res.OCOOrders[0])
}

func (s *baseOrderTestSuite) assertCancelOrderResponseEqual(e, a *CancelOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigClientOrderID, a.OrigClientOrderID, "OrigClientOrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.OrigQuoteOrderQuantity, a.OrigQuoteOrderQuantity, "OrigQuoteOrderQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
}

func (s *baseOrderTestSuite) assertCancelOCOResponseEqual(e, a *CancelOCOResponse) {
	r := s.r()
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertOCOOrderReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertOCOOrderEqual(order, a.Orders[idx])
	}
}

func (s *orderServiceTestSuite) TestCancelReplaceOrder() {
	data := []byte(`{
		"cancelResult": "SUCCESS",
		"newOrderResult": "SUCCESS",
		"cancelResponse": {
			"symbol": "LTCBTC",
			"orderId": 4,
			"orderListId": -1,
			"clientOrderId": "myOrder1",
			"origClientOrderId": "JYdYzLVhuxtLm10bYM6FOx",
			"transactTime": 1499827319559,
			"price": "0.00000000",
			"origQty": "12.00000000",
			"origQuoteOrderQty": "0.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "CANCELED",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "BUY",
			"selfTradePreventionMode": "NONE"
		},
		"newOrderResponse": {
			"symbol": "LTCBTC",
			"orderId": 5,
			"orderListId": -1,
			"clientOrderId": "myOrder2",
			"transactTime": 1499827319559,
			"price": "0.0002",
			"origQty": "12.00000000",
			"origQuoteOrderQty": "0.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "BUY",
			"fills": [],
			"selfTradePreventionMode": "NONE"
		}
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	cancelReplaceMode := CancelReplaceModeStopOnFailure
	timeInForce := TimeInForceTypeGTC
	quantity := "12.00"
	price := "0.0002"
	newClientOrderID := "myOrder2"
	cancelOrderID := int64(4)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"side":              side,
			"type":              orderType,
			"cancelReplaceMode": cancelReplaceMode,
			"timeInForce":       timeInForce,
			"quantity":          quantity,
			"price":             price,
			"newClientOrderId":  newClientOrderID,
			"cancelOrderId":     cancelOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelReplaceOrderService().Symbol(symbol).Side(side).
		Type(orderType).CancelReplaceMode(cancelReplaceMode).TimeInForce(timeInForce).
		Quantity(quantity).Price(price).NewClientOrderID(newClientOrderID).
		CancelOrderID(cancelOrderID).Do(newContext())

	s.r().NoError(err)

	e := &CancelReplaceOrderResponse{
		CancelResult:   "SUCCESS",
		NewOrderResult: "SUCCESS",
		CancelResponse: &CancelOrderResponse{
			Symbol:                   "LTCBTC",
			OrderID:                  4,
			OrderListID:              -1,
			ClientOrderID:            "myOrder1",
			OrigClientOrderID:        "JYdYzLVhuxtLm10bYM6FOx",
			TransactTime:             1499827319559,
			Price:                    "0.00000000",
			OrigQuantity:             "12.00000000",
			OrigQuoteOrderQuantity:   "0.00000000",
			ExecutedQuantity:         "0.00000000",
			CummulativeQuoteQuantity: "0.00000000",
			Status:                   OrderStatusTypeCanceled,
			TimeInForce:              TimeInForceTypeGTC,
			Type:                     OrderTypeLimit,
			Side:                     SideTypeBuy,
			SelfTradePreventionMode:  SelfTradePreventionModeNone,
		},
		NewOrderResponse: &CreateOrderResponse{
			Symbol:                   "LTCBTC",
			OrderID:                  5,
			ClientOrderID:            "myOrder2",
			TransactTime:             1499827319559,
			Price:                    "0.0002",
			OrigQuantity:             "12.00000000",
			OrigQuoteOrderQuantity:   "0.00000000",
			ExecutedQuantity:         "0.00000000",
			CummulativeQuoteQuantity: "0.00000000",
			Status:                   OrderStatusTypeNew,
			TimeInForce:              TimeInForceTypeGTC,
			Type:                     OrderTypeLimit,
			Side:                     SideTypeBuy,
			Fills:                    []*Fill{},
			SelfTradePreventionMode:  SelfTradePreventionModeNone,
		},
	}
	s.assertCancelReplaceOrderResponseEqual(e, res)
}

func (s *orderServiceTestSuite) TestCancelReplaceOrderAllowFailure() {
	data := []byte(`{
		"cancelResult": "FAILURE",
		"newOrderResult": "SUCCESS",
		"newOrderResponse": {
			"symbol": "LTCBTC",
			"orderId": 6,
			"orderListId": -1,
			"clientOrderId": "myOrder3",
			"transactTime": 1499827319559,
			"price": "0.0003",
			"origQty": "10.00000000",
			"origQuoteOrderQty": "0.00000000",
			"executedQty": "0.00000000",
			"cummulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "SELL",
			"fills": [],
			"selfTradePreventionMode": "NONE"
		}
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	side := SideTypeSell
	orderType := OrderTypeLimit
	cancelReplaceMode := CancelReplaceModeAllowFailure
	timeInForce := TimeInForceTypeGTC
	quantity := "10.00"
	price := "0.0003"
	newClientOrderID := "myOrder3"
	cancelOrigClientOrderID := "originalOrder123"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":                  symbol,
			"side":                    side,
			"type":                    orderType,
			"cancelReplaceMode":       cancelReplaceMode,
			"timeInForce":             timeInForce,
			"quantity":                quantity,
			"price":                   price,
			"newClientOrderId":        newClientOrderID,
			"cancelOrigClientOrderId": cancelOrigClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelReplaceOrderService().Symbol(symbol).Side(side).
		Type(orderType).CancelReplaceMode(cancelReplaceMode).TimeInForce(timeInForce).
		Quantity(quantity).Price(price).NewClientOrderID(newClientOrderID).
		CancelOrigClientOrderID(cancelOrigClientOrderID).Do(newContext())

	s.r().NoError(err)

	e := &CancelReplaceOrderResponse{
		CancelResult:   "FAILURE",
		NewOrderResult: "SUCCESS",
		CancelResponse: nil,
		NewOrderResponse: &CreateOrderResponse{
			Symbol:                   "LTCBTC",
			OrderID:                  6,
			ClientOrderID:            "myOrder3",
			TransactTime:             1499827319559,
			Price:                    "0.0003",
			OrigQuantity:             "10.00000000",
			OrigQuoteOrderQuantity:   "0.00000000",
			ExecutedQuantity:         "0.00000000",
			CummulativeQuoteQuantity: "0.00000000",
			Status:                   OrderStatusTypeNew,
			TimeInForce:              TimeInForceTypeGTC,
			Type:                     OrderTypeLimit,
			Side:                     SideTypeSell,
			Fills:                    []*Fill{},
			SelfTradePreventionMode:  SelfTradePreventionModeNone,
		},
	}
	s.assertCancelReplaceOrderResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertCancelReplaceOrderResponseEqual(e, a *CancelReplaceOrderResponse) {
	r := s.r()
	r.Equal(e.CancelResult, a.CancelResult, "CancelResult")
	r.Equal(e.NewOrderResult, a.NewOrderResult, "NewOrderResult")

	if e.CancelResponse != nil {
		r.NotNil(a.CancelResponse, "CancelResponse should not be nil")
		s.assertCancelOrderResponseEqual(e.CancelResponse, a.CancelResponse)
	} else {
		r.Nil(a.CancelResponse, "CancelResponse should be nil")
	}

	if e.NewOrderResponse != nil {
		r.NotNil(a.NewOrderResponse, "NewOrderResponse should not be nil")
		s.assertCreateOrderResponseEqual(e.NewOrderResponse, a.NewOrderResponse)
	} else {
		r.Nil(a.NewOrderResponse, "NewOrderResponse should be nil")
	}
}
