package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginOrderServiceTestSuite struct {
	baseOrderTestSuite
}

func TestMarginOrderService(t *testing.T) {
	suite.Run(t, new(marginOrderServiceTestSuite))
}

func (s *marginOrderServiceTestSuite) TestCreateOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
		"executedQty": "10.00",
		"cummulativeQuoteQty": "10.00",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"selfTradePreventionMode": "EXPIRE_TAKER",
		"autoRepayAtCancel": false
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
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":                  symbol,
			"side":                    side,
			"type":                    orderType,
			"timeInForce":             timeInForce,
			"quantity":                quantity,
			"quoteOrderQty":           quoteOrderQty,
			"price":                   price,
			"newClientOrderId":        newClientOrderID,
			"sideEffectType":          SideEffectTypeNoSideEffect,
			"selfTradePreventionMode": SelfTradePreventionModeExpireTaker,
			"autoRepayAtCancel":       "FALSE",
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateMarginOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).
		Price(price).NewClientOrderID(newClientOrderID).SideEffectType(SideEffectTypeNoSideEffect).
		SelfTradePreventionMode(SelfTradePreventionModeExpireTaker).
		AutoRepayAtCancel(false).
		Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  1,
		ClientOrderID:            "myOrder1",
		TransactTime:             1499827319559,
		Price:                    "0.0001",
		OrigQuantity:             "12.00",
		ExecutedQuantity:         "10.00",
		CummulativeQuoteQuantity: "10.00",
		Status:                   OrderStatusTypeFilled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeBuy,
		SelfTradePreventionMode:  SelfTradePreventionModeExpireTaker,
	}
	s.assertCreateOrderResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) TestCreateOrderFull() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"transactTime": 1499827319559,
		"price": "0.0001",
		"origQty": "12.00",
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
	res, err := s.client.NewCreateMarginOrderService().Symbol(symbol).Side(side).
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
}

func (s *marginOrderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"symbol": "LTCBTC",
		"orderId": "28",
		"origClientOrderId": "myOrder1",
		"clientOrderId": "cancelMyOrder1",
		"transactTime": 1507725176595,
		"price": "1.00000000",
		"origQty": "10.00000000",
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

	res, err := s.client.NewCancelMarginOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).
		NewClientOrderID(newClientOrderID).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelMarginOrderResponse{
		Symbol:                   "LTCBTC",
		OrderID:                  "28",
		OrigClientOrderID:        "myOrder1",
		ClientOrderID:            "cancelMyOrder1",
		TransactTime:             1507725176595,
		Price:                    "1.00000000",
		OrigQuantity:             "10.00000000",
		ExecutedQuantity:         "8.00000000",
		CummulativeQuoteQuantity: "8.00000000",
		Status:                   OrderStatusTypeCanceled,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeSell,
	}
	s.assertCancelMarginOrderResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) TestCancelAllOrder() {
	data := []byte(`[{
		"symbol": "LTCBTC",
		"orderId": "28",
		"origClientOrderId": "myOrder1",
		"clientOrderId": "cancelMyOrder1",
		"transactTime": 1507725176595,
		"price": "1.00000000",
		"origQty": "10.00000000",
		"executedQty": "8.00000000",
		"cummulativeQuoteQty": "8.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "SELL"
	}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelAllMarginOrdersService().Symbol(symbol).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*CancelAllMarginOrdersResponse{
		&CancelAllMarginOrdersResponse{
			Symbol:                   "LTCBTC",
			OrderID:                  "28",
			OrigClientOrderID:        "myOrder1",
			ClientOrderID:            "cancelMyOrder1",
			TransactTime:             1507725176595,
			Price:                    "1.00000000",
			OrigQuantity:             "10.00000000",
			ExecutedQuantity:         "8.00000000",
			CummulativeQuoteQuantity: "8.00000000",
			Status:                   OrderStatusTypeCanceled,
			TimeInForce:              TimeInForceTypeGTC,
			Type:                     OrderTypeLimit,
			Side:                     SideTypeSell,
		},
	}
	s.assertCancelAllMarginOrdersResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) TestGetOrder() {
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
	order, err := s.client.NewGetMarginOrderService().Symbol(symbol).
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

func (s *marginOrderServiceTestSuite) TestListMarginOpenOrders() {
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
			"updateTime": 1562040170089
		   }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	orders, err := s.client.NewListMarginOpenOrdersService().Symbol(symbol).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	r.Len(orders, 1)
	e := &Order{
		Symbol:                   "BNBBTC",
		OrderID:                  211842552,
		ClientOrderID:            "qhcZw71gAkCCTv0t0k8LUK",
		CummulativeQuoteQuantity: "0.00000000",
		Price:                    "0.00475010",
		OrigQuantity:             "0.30000000",
		ExecutedQuantity:         "0.00000000",
		Status:                   OrderStatusTypeNew,
		TimeInForce:              TimeInForceTypeGTC,
		Type:                     OrderTypeLimit,
		Side:                     SideTypeSell,
		StopPrice:                "0.00000000",
		IcebergQuantity:          "0.00000000",
		Time:                     1562040170089,
		UpdateTime:               1562040170089,
		IsWorking:                true,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *marginOrderServiceTestSuite) TestListMarginOrders() {
	data := []byte(`[
		{
			"orderId": 43123876,
			"price": "0.00395740",
			"origQty": "4.06000000",
			"cummulativeQuoteQty": "0.01606704",
			"symbol": "BNBBTC",
			"time": 1556089977693
		},
		{
			"orderId": 43123877,
			"price": "0.00395740",
			"origQty": "0.77000000",
			"cummulativeQuoteQty": "0.00304719",
			"symbol": "BNBBTC",
			"time": 1556089977693
		},
		{
			"orderId": 43253549,
			"price": "0.00428930",
			"origQty": "23.30000000",
			"cummulativeQuoteQty": "0.09994069",
			"symbol": "BNBBTC",
			"time": 1556163963504
		}
  	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BNBBTC"
	limit := 3
	startTime := int64(1556089977693)
	endTime := int64(1556163963504)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewListMarginOrdersService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).
		Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(orders, 3)
	e := []*Order{
		{
			OrderID:                  43123876,
			Price:                    "0.00395740",
			OrigQuantity:             "4.06000000",
			CummulativeQuoteQuantity: "0.01606704",
			Symbol:                   "BNBBTC",
			Time:                     1556089977693,
		},
		{
			OrderID:                  43123877,
			Price:                    "0.00395740",
			OrigQuantity:             "0.77000000",
			CummulativeQuoteQuantity: "0.00304719",
			Symbol:                   "BNBBTC",
			Time:                     1556089977693,
		},
		{
			OrderID:                  43253549,
			Price:                    "0.00428930",
			OrigQuantity:             "23.30000000",
			CummulativeQuoteQuantity: "0.09994069",
			Symbol:                   "BNBBTC",
			Time:                     1556163963504,
		},
	}
	s.r().Len(orders, len(e))
	for i := 0; i < len(orders); i++ {
		s.assertMarginAllOrderEqual(e[i], orders[i])
	}
}

func (s *marginOrderServiceTestSuite) assertMarginAllOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
}

func (s *marginOrderServiceTestSuite) assertCancelMarginOrderResponseEqual(e, a *CancelMarginOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigClientOrderID, a.OrigClientOrderID, "OrigClientOrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CummulativeQuoteQuantity, a.CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
}

func (s *marginOrderServiceTestSuite) assertCancelAllMarginOrdersResponseEqual(e, a []*CancelAllMarginOrdersResponse) {
	r := s.r()
	r.Equal(len(e), len(a), "Length")
	r.Equal(e[0].Symbol, a[0].Symbol, "Symbol")
	r.Equal(e[0].OrderID, a[0].OrderID, "OrderID")
	r.Equal(e[0].OrigClientOrderID, a[0].OrigClientOrderID, "OrigClientOrderID")
	r.Equal(e[0].ClientOrderID, a[0].ClientOrderID, "ClientOrderID")
	r.Equal(e[0].TransactTime, a[0].TransactTime, "TransactTime")
	r.Equal(e[0].Price, a[0].Price, "Price")
	r.Equal(e[0].OrigQuantity, a[0].OrigQuantity, "OrigQuantity")
	r.Equal(e[0].ExecutedQuantity, a[0].ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e[0].CummulativeQuoteQuantity, a[0].CummulativeQuoteQuantity, "CummulativeQuoteQuantity")
	r.Equal(e[0].Status, a[0].Status, "Status")
	r.Equal(e[0].TimeInForce, a[0].TimeInForce, "TimeInForce")
	r.Equal(e[0].Type, a[0].Type, "Type")
	r.Equal(e[0].Side, a[0].Side, "Side")
}

func (s *marginOrderServiceTestSuite) TestCreateOCO() {
	data := []byte(`{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
		"transactionTime": 1574040868128,
		"symbol": "LTCBTC",
		"marginBuyBorrowAmount": "5",
		"marginBuyBorrowAsset": "BTC",
		"isIsolated": true,
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
			"orderId": 2,
			"orderListId": 0,
			"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
			"transactTime": 1563417480525,
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
			"orderId": 3,
			"orderListId": 0,
			"clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
			"transactTime": 1563417480525,
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
	isIsolated := true
	side := SideTypeBuy
	timeInForce := TimeInForceTypeGTC
	quantity := "10"
	price := "3"
	stopPrice := "3.1"
	stopLimitPrice := "3.2"
	limitClientOrderID := "myOrder1"
	newOrderRespType := NewOrderRespTypeFULL
	sideEffectType := SideEffectTypeMarginBuy
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":               symbol,
			"isIsolated":           "TRUE",
			"side":                 side,
			"quantity":             quantity,
			"price":                price,
			"stopPrice":            stopPrice,
			"stopLimitPrice":       stopLimitPrice,
			"stopLimitTimeInForce": timeInForce,
			"limitClientOrderId":   limitClientOrderID,
			"newOrderRespType":     newOrderRespType,
			"sideEffectType":       sideEffectType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateMarginOCOService().
		Symbol(symbol).
		IsIsolated(isIsolated).
		Side(side).
		Quantity(quantity).
		Price(price).
		StopPrice(stopPrice).
		StopLimitPrice(stopLimitPrice).
		StopLimitTimeInForce(timeInForce).
		LimitClientOrderID(limitClientOrderID).
		NewOrderRespType(newOrderRespType).
		SideEffectType(sideEffectType).
		Do(newContext())

	s.r().NoError(err)
	e := &CreateMarginOCOResponse{
		OrderListID:           0,
		ContingencyType:       "OCO",
		ListStatusType:        "EXEC_STARTED",
		ListOrderStatus:       "EXECUTING",
		ListClientOrderID:     "C3wyj4WVEktd7u9aVBRXcN",
		TransactionTime:       1574040868128,
		Symbol:                "LTCBTC",
		MarginBuyBorrowAmount: "5",
		MarginBuyBorrowAsset:  "BTC",
		IsIsolated:            true,
		Orders: []*MarginOCOOrder{
			&MarginOCOOrder{
				Symbol:        "LTCBTC",
				OrderID:       2,
				ClientOrderID: "pO9ufTiFGg3nw2fOdgeOXa",
			},
			&MarginOCOOrder{
				Symbol:        "LTCBTC",
				OrderID:       3,
				ClientOrderID: "TXOvglzXuaubXAaENpaRCB",
			},
		},
		OrderReports: []*MarginOCOOrderReport{
			&MarginOCOOrderReport{
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
			&MarginOCOOrderReport{
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
	s.assertCreateMarginOCOResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) assertCreateMarginOCOResponseEqual(e, a *CreateMarginOCOResponse) {
	r := s.r()
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.MarginBuyBorrowAmount, a.MarginBuyBorrowAmount, "MarginBuyBorrowAmount")
	r.Equal(e.MarginBuyBorrowAsset, a.MarginBuyBorrowAsset, "MarginBuyBorrowAsset")
	r.Equal(e.IsIsolated, a.IsIsolated, "IsIsolated")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertMarginOCOOrderReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertMarginOCOOrderEqual(order, a.Orders[idx])
	}
}

func (s *marginOrderServiceTestSuite) assertMarginOCOOrderReportEqual(e, a *MarginOCOOrderReport) {
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

func (s *marginOrderServiceTestSuite) assertMarginOCOOrderEqual(e, a *MarginOCOOrder) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *marginOrderServiceTestSuite) TestCancelOCO() {
	data := []byte(`{
		"orderListId":1000,
		"contingencyType":"OCO",
		"listStatusType":"ALL_DONE",
		"listOrderStatus":"ALL_DONE",
		"listClientOrderId":"C3wyj4WVEktd7u9aVBRXcN",
		"transactionTime":1614272133000,
		"symbol":"LTCBTC",
		"isIsolated":true,
		"orders":[
			{
				"symbol":"LTCBTC",
				"orderId":1100,
				"clientOrderId":"pO9ufTiFGg3nw2fOdgeOXa"
			},
			{
				"symbol":"LTCBTC",
				"orderId":1010,
				"clientOrderId":"TXOvglzXuaubXAaENpaRCB"
			}
		],
		"orderReports":[
			{
				"symbol":"LTCBTC",
				"origClientOrderId":"pO9ufTiFGg3nw2fOdgeOXa",
				"orderId":1100,
				"orderListId":1000,
				"clientOrderId":"unfWT8ig8i0uj6lPuYLez6",
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
				"symbol":"LTCBTC",
				"origClientOrderId":"TXOvglzXuaubXAaENpaRCB",
				"orderId":1010,
				"orderListId":1000,
				"clientOrderId":"unfWT8ig8i0uj6lPuYLez6",
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

	symbol := "LTCBTC"
	listClientOrderID := "C3wyj4WVEktd7u9aVBRXcN"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"listClientOrderId": listClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.
		NewCancelMarginOCOService().
		Symbol(symbol).
		ListClientOrderID(listClientOrderID).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelMarginOCOResponse{
		OrderListID:       1000,
		ContingencyType:   "OCO",
		ListStatusType:    "ALL_DONE",
		ListOrderStatus:   "ALL_DONE",
		ListClientOrderID: "C3wyj4WVEktd7u9aVBRXcN",
		TransactionTime:   1614272133000,
		Symbol:            "LTCBTC",
		IsIsolated:        true,
		Orders: []*MarginOCOOrder{
			{Symbol: "LTCBTC", OrderID: 1100, ClientOrderID: "pO9ufTiFGg3nw2fOdgeOXa"},
			{Symbol: "LTCBTC", OrderID: 1010, ClientOrderID: "TXOvglzXuaubXAaENpaRCB"},
		},
		OrderReports: []*MarginOCOOrderReport{
			{
				Symbol:                   "LTCBTC",
				OrderID:                  1100,
				OrderListID:              1000,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
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
				Symbol:                   "LTCBTC",
				OrderID:                  1010,
				OrderListID:              1000,
				ClientOrderID:            "unfWT8ig8i0uj6lPuYLez6",
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
	s.assertCancelMarginOCOResponseEqual(e, res)
}

func (s *marginOrderServiceTestSuite) assertCancelMarginOCOResponseEqual(e, a *CancelMarginOCOResponse) {
	r := s.r()
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.IsIsolated, a.IsIsolated, "IsIsolated")

	r.Len(a.OrderReports, len(e.OrderReports))
	for idx, orderReport := range e.OrderReports {
		s.assertMarginOCOOrderReportEqual(orderReport, a.OrderReports[idx])
	}

	r.Len(a.Orders, len(e.Orders))
	for idx, order := range e.Orders {
		s.assertMarginOCOOrderEqual(order, a.Orders[idx])
	}
}
