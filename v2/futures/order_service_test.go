package futures

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/adshao/go-binance/v2/common"

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
		"clientOrderId": "testOrder",
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 22542179,
		"origQty": "10",
		"price": "10000",
		"reduceOnly": false,
		"side": "SELL",
		"status": "NEW",
		"stopPrice": "0",
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1566818724722,
		"workingType": "CONTRACT_PRICE",
		"activatePrice": "1000",
		"priceRate": "0.1",
		"positionSide": "BOTH",
		"closePosition": false,
		"priceProtect": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	side := SideTypeSell
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	positionSide := PositionSideTypeBoth
	quantity := "10"
	price := "10000"
	priceMatch := PriceMatchTypeNone
	newClientOrderID := "testOrder"
	reduceOnly := false
	stopPrice := "0"
	activationPrice := "1000"
	callbackRate := "0.1"
	workingType := WorkingTypeContractPrice
	priceProtect := true
	newOrderResponseType := NewOrderRespTypeRESULT
	closePosition := false
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"timeInForce":      timeInForce,
			"positionSide":     positionSide,
			"quantity":         quantity,
			"reduceOnly":       strconv.FormatBool(reduceOnly),
			"price":            price,
			"priceMatch":       priceMatch,
			"newClientOrderId": newClientOrderID,
			"stopPrice":        stopPrice,
			"workingType":      workingType,
			"activationPrice":  activationPrice,
			"callbackRate":     callbackRate,
			"priceProtect":     strconv.FormatBool(priceProtect),
			"newOrderRespType": newOrderResponseType,
			"closePosition":    strconv.FormatBool(closePosition),
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).ClosePosition(closePosition).
		ReduceOnly(reduceOnly).Price(price).PriceMatch(priceMatch).NewClientOrderID(newClientOrderID).
		StopPrice(stopPrice).WorkingType(workingType).ActivationPrice(activationPrice).
		CallbackRate(callbackRate).PositionSide(positionSide).
		PriceProtect(priceProtect).NewOrderResponseType(newOrderResponseType).
		Do(newContext())
	s.r().NoError(err)
	e := &CreateOrderResponse{
		ClientOrderID:    newClientOrderID,
		CumQuote:         "0",
		ExecutedQuantity: "0",
		OrderID:          22542179,
		OrigQuantity:     "10",
		PositionSide:     positionSide,
		Price:            "10000",
		ReduceOnly:       false,
		Side:             SideTypeSell,
		Status:           OrderStatusTypeNew,
		StopPrice:        "0",
		Symbol:           symbol,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		UpdateTime:       1566818724722,
		WorkingType:      WorkingTypeContractPrice,
		ActivatePrice:    activationPrice,
		PriceRate:        callbackRate,
		ClosePosition:    false,
		PriceProtect:     priceProtect,
	}
	s.assertCreateOrderResponseEqual(e, res)
}

func (s *orderServiceTestSuite) TestCreateOrderId() {
	data := []byte(`{
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 22542179,
		"origQty": "10",
		"price": "10000",
		"reduceOnly": false,
		"side": "SELL",
		"status": "NEW",
		"stopPrice": "0",
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1566818724722,
		"workingType": "CONTRACT_PRICE",
		"activatePrice": "1000",
		"priceRate": "0.1",
		"positionSide": "BOTH",
		"closePosition": false,
		"priceProtect": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	side := SideTypeSell
	orderType := OrderTypeLimit
	timeInForce := TimeInForceTypeGTC
	positionSide := PositionSideTypeBoth
	quantity := "10"
	price := "10000"
	reduceOnly := false
	stopPrice := "0"
	activationPrice := "1000"
	callbackRate := "0.1"
	workingType := WorkingTypeContractPrice
	priceProtect := true
	newOrderResponseType := NewOrderRespTypeRESULT
	closePosition := false
	s.assertReq(func(r *request) {
		s.Assertions.True(strings.HasPrefix(r.form.Get("newClientOrderId"), "x-ftGmvgAN"))
	})
	_, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).TimeInForce(timeInForce).Quantity(quantity).ClosePosition(closePosition).
		ReduceOnly(reduceOnly).Price(price).
		StopPrice(stopPrice).WorkingType(workingType).ActivationPrice(activationPrice).
		CallbackRate(callbackRate).PositionSide(positionSide).
		PriceProtect(priceProtect).NewOrderResponseType(newOrderResponseType).
		Do(newContext())
	s.r().NoError(err)
}

func (s *baseOrderTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponse) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CumQuote, a.CumQuote, "CumQuote")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.PriceRate, a.PriceRate, "PriceRate")
	r.Equal(e.ClosePosition, a.ClosePosition, "ClosePosition")
}

func (s *orderServiceTestSuite) TestListOpenOrders() {
	data := []byte(`[
		{
		  "symbol": "BTCUSDT",
		  "orderId": 1,
		  "clientOrderId": "myOrder1",
		  "price": "0.1",
		  "reduceOnly": false,
		  "origQty": "1.0",
		  "cumQty": "1.0",
		  "cumQuote": "1.0",
		  "status": "NEW",
		  "timeInForce": "GTC",
		  "type": "LIMIT",
		  "side": "BUY",
		  "stopPrice": "0.0",
		  "time": 1499827319559,
		  "updateTime": 1499827319559,
		  "workingType": "CONTRACT_PRICE",
		  "activatePrice": "10000",
		  "priceRate":"0.1",
		  "positionSide":"BOTH",
		  "priceProtect": false
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
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
		Symbol:        symbol,
		OrderID:       1,
		ClientOrderID: "myOrder1",
		Price:         "0.1",
		ReduceOnly:    false,
		OrigQuantity:  "1.0",
		CumQuantity:   "1.0",
		CumQuote:      "1.0",
		Status:        OrderStatusTypeNew,
		TimeInForce:   TimeInForceTypeGTC,
		Type:          OrderTypeLimit,
		Side:          SideTypeBuy,
		StopPrice:     "0.0",
		Time:          1499827319559,
		UpdateTime:    1499827319559,
		WorkingType:   WorkingTypeContractPrice,
		ActivatePrice: "10000",
		PriceRate:     "0.1",
		PositionSide:  "BOTH",
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *baseOrderTestSuite) assertOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CumQuantity, a.CumQuantity, "CumQuantity")
	r.Equal(e.CumQuote, a.CumQuote, "CumQuote")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.Time, e.Time, "Time")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.PriceRate, a.PriceRate, "PriceRate")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
}

func (s *orderServiceTestSuite) TestGetOpenOrder() {
	data := []byte(`{
		  "symbol": "BTCUSDT",
		  "orderId": 1,
		  "clientOrderId": "myOrder1",
		  "price": "0.1",
		  "reduceOnly": false,
		  "origQty": "1.0",
		  "cumQty": "1.0",
		  "cumQuote": "1.0",
		  "status": "NEW",
		  "timeInForce": "GTC",
		  "type": "LIMIT",
		  "side": "BUY",
		  "stopPrice": "0.0",
		  "time": 1499827319559,
		  "updateTime": 1499827319559,
		  "workingType": "CONTRACT_PRICE",
		  "activatePrice": "10000",
		  "priceRate":"0.1",
		  "positionSide":"BOTH",
		  "priceProtect": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderId := int64(1)
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"orderId":    1,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetOpenOrderService().Symbol(symbol).OrderID(orderId).
		Do(newContext(), WithRecvWindow(recvWindow))
	r := s.r()
	r.NoError(err)
	e := &Order{
		Symbol:        symbol,
		OrderID:       orderId,
		ClientOrderID: "myOrder1",
		Price:         "0.1",
		ReduceOnly:    false,
		OrigQuantity:  "1.0",
		CumQuantity:   "1.0",
		CumQuote:      "1.0",
		Status:        OrderStatusTypeNew,
		TimeInForce:   TimeInForceTypeGTC,
		Type:          OrderTypeLimit,
		Side:          SideTypeBuy,
		StopPrice:     "0.0",
		Time:          1499827319559,
		UpdateTime:    1499827319559,
		WorkingType:   WorkingTypeContractPrice,
		ActivatePrice: "10000",
		PriceRate:     "0.1",
		PositionSide:  "BOTH",
	}
	s.assertOrderEqual(e, order)
}

func (s *orderServiceTestSuite) TestGetOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"orderId": 1,
		"clientOrderId": "myOrder1",
		"price": "0.1",
		"reduceOnly": false,
		"origQty": "1.0",
		"executedQty": "0.0",
		"cumQuote": "0.0",
		"status": "NEW",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"stopPrice": "0.0",
		"time": 1499827319559,
		"updateTime": 1499827319559,
		"workingType": "CONTRACT_PRICE",
		"activatePrice": "10000",
		"priceRate":"0.1",
		"positionSide": "BOTH",
		"priceProtect": false,
		"closePosition": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
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
		Symbol:           symbol,
		OrderID:          1,
		ClientOrderID:    origClientOrderID,
		Price:            "0.1",
		ReduceOnly:       false,
		OrigQuantity:     "1.0",
		ExecutedQuantity: "0.0",
		CumQuote:         "0.0",
		Status:           OrderStatusTypeNew,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		Side:             SideTypeBuy,
		StopPrice:        "0.0",
		Time:             1499827319559,
		UpdateTime:       1499827319559,
		WorkingType:      WorkingTypeContractPrice,
		ActivatePrice:    "10000",
		PriceRate:        "0.1",
		PositionSide:     "BOTH",
		PriceProtect:     false,
		ClosePosition:    true,
	}
	s.assertOrderEqual(e, order)
}

func (s *orderServiceTestSuite) TestListOrders() {
	data := []byte(`[
		{
		  "symbol": "BTCUSDT",
		  "orderId": 1,
		  "clientOrderId": "myOrder1",
		  "price": "0.1",
		  "reduceOnly": false,
		  "origQty": "1.0",
		  "executedQty": "1.0",
		  "cumQuote": "10.0",
		  "status": "NEW",
		  "timeInForce": "GTC",
		  "type": "LIMIT",
		  "side": "BUY",
		  "stopPrice": "0.0",
		  "time": 1499827319559,
		  "updateTime": 1499827319559,
		  "workingType": "CONTRACT_PRICE",
		  "activatePrice": "10000",
		  "priceRate":"0.1",
		  "priceProtect": false
		}
	  ]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
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
		Symbol:           symbol,
		OrderID:          1,
		ClientOrderID:    "myOrder1",
		Price:            "0.1",
		ReduceOnly:       false,
		OrigQuantity:     "1.0",
		ExecutedQuantity: "1.0",
		CumQuote:         "10.0",
		Status:           OrderStatusTypeNew,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		Side:             SideTypeBuy,
		StopPrice:        "0.0",
		Time:             1499827319559,
		UpdateTime:       1499827319559,
		WorkingType:      WorkingTypeContractPrice,
		ActivatePrice:    "10000",
		PriceRate:        "0.1",
		PriceProtect:     false,
	}
	s.assertOrderEqual(e, orders[0])
}

func (s *orderServiceTestSuite) TestModifyOrder() {
	data := []byte(`{
		"orderId": 20072994037,
		"symbol": "BTCUSDT",
		"pair": "BTCUSDT",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumBase": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"closePosition": false,
		"side": "BUY",
		"positionSide": "LONG",
		"stopPrice": "0",
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false,
		"origType": "LIMIT",
		"priceMatch": "NONE",
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0,
		"updateTime": 1629182711600
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orderID := int64(20072994037)
	origClientOrderID := "LJ9R4QZDihCaS8UAOOLpgW"
	symbol := "BTCUSDT"
	side := SideTypeBuy
	quantity := "1"
	price := "30005"
	priceMatch := PriceMatchTypeNone
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
			"symbol":            symbol,
			"side":              side,
			"quantity":          quantity,
			"price":             price,
			"priceMatch":        priceMatch,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewModifyOrderService().OrderID(orderID).OrigClientOrderID(origClientOrderID).
		Symbol(symbol).Side(side).Quantity(quantity).Price(price).PriceMatch(priceMatch).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &ModifyOrderResponse{
		OrderID:                 orderID,
		Symbol:                  symbol,
		Pair:                    "BTCUSDT",
		Status:                  OrderStatusTypeNew,
		ClientOrderID:           origClientOrderID,
		Price:                   price,
		AveragePrice:            "0.0",
		OriginalQuantity:        quantity,
		ExecutedQuantity:        "0",
		CumulativeQuantity:      "0",
		CumulativeBase:          "0",
		TimeInForce:             TimeInForceTypeGTC,
		Type:                    OrderTypeLimit,
		ReduceOnly:              false,
		ClosePosition:           false,
		Side:                    side,
		PositionSide:            PositionSideTypeLong,
		StopPrice:               "0",
		WorkingType:             WorkingTypeContractPrice,
		PriceProtect:            false,
		OriginalType:            OrderTypeLimit,
		PriceMatch:              priceMatch,
		SelfTradePreventionMode: "NONE",
		GoodTillDate:            0,
		UpdateTime:              1629182711600,
	}
	s.assertModifyOrderResponseEqual(e, res)
}

func (s *baseOrderTestSuite) assertModifyOrderResponseEqual(e, a *ModifyOrderResponse) {
	r := s.r()
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.AveragePrice, a.AveragePrice, "AveragePrice")
	r.Equal(e.OriginalQuantity, a.OriginalQuantity, "OriginalQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.CumulativeQuantity, a.CumulativeQuantity, "CumulativeQuantity")
	r.Equal(e.CumulativeBase, a.CumulativeBase, "CumulativeBase")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.ClosePosition, a.ClosePosition, "ClosePosition")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
	r.Equal(e.OriginalType, a.OriginalType, "OriginalType")
	r.Equal(e.PriceMatch, a.PriceMatch, "PriceMatch")
	r.Equal(e.SelfTradePreventionMode, a.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(e.GoodTillDate, a.GoodTillDate, "GoodTillDate")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}

func (s *orderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"clientOrderId": "myOrder1",
		"cumQty": "0",
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 283194212,
		"origQty": "11",
		"price": "8301",
		"reduceOnly": false,
		"side": "BUY",
		"status": "CANCELED",
		"stopPrice": "8300",
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "TAKE_PROFIT",
		"updateTime": 1571110484038,
		"workingType": "CONTRACT_PRICE",
		"activatePrice": "10000",
		"priceRate":"0.1",
		"positionSide":"BOTH",
		"priceProtect": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderID := int64(28)
	origClientOrderID := "myOrder1"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":            symbol,
			"orderId":           orderID,
			"origClientOrderId": origClientOrderID,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOrderService().Symbol(symbol).
		OrderID(orderID).OrigClientOrderID(origClientOrderID).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CancelOrderResponse{
		ClientOrderID:    origClientOrderID,
		CumQuantity:      "0",
		CumQuote:         "0",
		ExecutedQuantity: "0",
		OrderID:          283194212,
		OrigQuantity:     "11",
		Price:            "8301",
		ReduceOnly:       false,
		Side:             SideTypeBuy,
		Status:           OrderStatusTypeCanceled,
		StopPrice:        "8300",
		Symbol:           symbol,
		TimeInForce:      TimeInForceTypeGTC,
		Type:             OrderTypeLimit,
		UpdateTime:       1571110484038,
		WorkingType:      WorkingTypeContractPrice,
		ActivatePrice:    "10000",
		PriceRate:        "0.1",
		PositionSide:     "BOTH",
		PriceProtect:     false,
	}
	s.assertCancelOrderResponseEqual(e, res)
}

func (s *orderServiceTestSuite) assertCancelOrderResponseEqual(e, a *CancelOrderResponse) {
	r := s.r()
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.CumQuantity, a.CumQuantity, "CumQuantity")
	r.Equal(e.CumQuote, a.CumQuote, "CumQuote")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.PriceRate, a.PriceRate, "PriceRate")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
}

func (s *orderServiceTestSuite) TestCancelAllOpenOrders() {
	data := []byte(`{
		"code": "200",
		"msg": "The operation of cancel all open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewCancelAllOpenOrdersService().Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
}

func (s *orderServiceTestSuite) TestListLiquidationOrders() {
	data := []byte(`[
		{
			  "symbol": "BTCUSDT",
			  "price": "7918.33",
			  "origQty": "0.014",
			  "executedQty": "0.014",
			  "avragePrice": "7918.33",
			  "status": "FILLED",
			  "timeInForce": "IOC",
			  "type": "LIMIT",
			  "side": "SELL",
			  "time": 1568014460893
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1568014460893)
	endTime := int64(1568014460894)
	limit := 1
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListLiquidationOrdersService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := []*LiquidationOrder{
		{
			Symbol:           symbol,
			Price:            "7918.33",
			OrigQuantity:     "0.014",
			ExecutedQuantity: "0.014",
			AveragePrice:     "7918.33",
			Status:           OrderStatusTypeFilled,
			TimeInForce:      TimeInForceTypeIOC,
			Type:             OrderTypeLimit,
			Side:             SideTypeSell,
		},
	}
	s.r().Len(res, len(e))
	for i := range res {
		s.assertLiquidationEqual(e[i], res[i])
	}
}

func (s *orderServiceTestSuite) assertLiquidationEqual(e, a *LiquidationOrder) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQuantity, a.OrigQuantity, "OrigQuantity")
	r.Equal(e.ExecutedQuantity, a.ExecutedQuantity, "ExecutedQuantity")
	r.Equal(e.AveragePrice, a.AveragePrice, "AveragePrice")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
}

func (s *orderServiceTestSuite) TestCreateBatchOrders() {
	data := []byte(`[
		{
			"code": -2014,
			"msg": "API-key format invalid."
		},
		{
			"clientOrderId": "testOrder",
			"cumQty": "0",
			"cumQuote": "0",
			"executedQty": "0",
			"orderId": 22542179,
			"avgPrice": "0.00000",
			"origQty": "10",
			"price": "100",
			"reduceOnly": false,
			"side": "BUY",
			"positionSide": "SHORT",
			"status": "NEW",
			"stopPrice": "9300",
			"symbol": "BTCUSDT",
			"timeInForce": "GTC",
			"type": "TRAILING_STOP_MARKET",
			"origType": "TRAILING_STOP_MARKET",
			"activatePrice": "9020",
			"priceRate": "0.3",
			"updateTime": 1566818724722,
			"workingType": "CONTRACT_PRICE",
			"priceProtect": false,
			"priceMatch": "NONE",
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0
		},
		{
			"code": -2022,
			"msg": "ReduceOnly Order is rejected."
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewCreateBatchOrdersService().OrderList([]*CreateOrderService{{}, {}, {}}).Do(context.Background())
	r := s.r()
	r.NoError(err)

	e := &CreateBatchOrdersResponse{
		N: 3,
		Orders: []*Order{
			{
				Symbol:                  "BTCUSDT",
				OrderID:                 22542179,
				ClientOrderID:           "testOrder",
				Price:                   "100",
				ReduceOnly:              false,
				OrigQuantity:            "10",
				ExecutedQuantity:        "0",
				CumQuantity:             "0",
				CumQuote:                "0",
				Status:                  "NEW",
				TimeInForce:             "GTC",
				Type:                    "TRAILING_STOP_MARKET",
				Side:                    "BUY",
				StopPrice:               "9300",
				Time:                    0,
				UpdateTime:              1566818724722,
				WorkingType:             "CONTRACT_PRICE",
				ActivatePrice:           "9020",
				PriceRate:               "0.3",
				AvgPrice:                "0.00000",
				OrigType:                "TRAILING_STOP_MARKET",
				PositionSide:            "SHORT",
				PriceProtect:            false,
				ClosePosition:           false,
				PriceMatch:              "NONE",
				SelfTradePreventionMode: "NONE",
				GoodTillDate:            0,
			},
		},
		Errors: []error{
			&common.APIError{
				Code:    -2014,
				Message: "API-key format invalid.",
			},
			nil,
			&common.APIError{
				Code:    -2022,
				Message: "ReduceOnly Order is rejected.",
			},
		},
	}
	r.EqualValues(e, res)
}

func (s *orderServiceTestSuite) TestModifyBatchOrders() {
	data := []byte(`[
	{
			"orderId": 42042723,
			"symbol": "BTCUSDT",
			"status": "NEW",
			"clientOrderId": "Ne7DEEvLvv8b8egTqrZceu",
			"price": "99995.00",
			"avgPrice": "0.00",
			"origQty": "1",
			"executedQty": "0",
			"cumQty": "0",
			"cumQuote": "0.00",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"reduceOnly": false,
			"closePosition": false,
			"side": "BUY",
			"positionSide": "BOTH",
			"stopPrice": "0.00",
			"workingType": "CONTRACT_PRICE",
			"priceProtect": false,
			"origType": "LIMIT",
			"priceMatch": "NONE",
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"updateTime": 1733500988978
		},
		{
			"code": -1102,
			"msg": "Mandatory parameter 'price' was not sent, was empty/null, or malformed."
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orders := []*ModifyOrder{
		new(ModifyOrder).
			Symbol("BTCUSDT").
			OrigClientOrderID("Ne7DEEvLvv8b8egTqrZceu").
			Quantity("1").
			Side("BUY").
			Price("99995.00"),
		new(ModifyOrder).
			Symbol("BTCUSDT").
			OrigClientOrderID("GYby67jLvv8b8egTr8Adrf").
			Quantity("1").
			Side("SELL").
			Price("-100005.00"),
	}

	res, err := s.client.NewModifyBatchOrdersService().OrderList(orders).Do(context.Background())

	r := s.r()
	r.NoError(err)

	r.Equal(1, len(res.Orders))

	e := &Order{
		Symbol:                  "BTCUSDT",
		OrderID:                 42042723,
		ClientOrderID:           "Ne7DEEvLvv8b8egTqrZceu",
		Price:                   "99995.00",
		ReduceOnly:              false,
		OrigQuantity:            "1",
		ExecutedQuantity:        "0",
		CumQuantity:             "0",
		CumQuote:                "0.00",
		Status:                  "NEW",
		TimeInForce:             "GTC",
		Type:                    "LIMIT",
		Side:                    "BUY",
		StopPrice:               "0.00",
		Time:                    0,
		UpdateTime:              1733500988978,
		WorkingType:             "CONTRACT_PRICE",
		ActivatePrice:           "",
		PriceRate:               "",
		AvgPrice:                "0.00",
		OrigType:                "LIMIT",
		PositionSide:            "BOTH",
		PriceProtect:            false,
		ClosePosition:           false,
		PriceMatch:              "NONE",
		SelfTradePreventionMode: "NONE",
		GoodTillDate:            0,
	}
	s.assertOrderEqual(e, res.Orders[0])

	r.Equal(
		[]error{
			nil,
			&common.APIError{
				Code:     -1102,
				Message:  "Mandatory parameter 'price' was not sent, was empty/null, or malformed.",
				Response: nil,
			},
		},
		res.Errors)

}
