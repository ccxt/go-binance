package options

import (
	"encoding/json"
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

func (s *baseOrderTestSuite) assertOrderEqual(e, a *Order) {
	r := s.r()
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.PostOnly, a.PostOnly, "PostOnly")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.Status, a.Status, "Status")
	if e.Reason != nil {
		r.Equal(e.Reason, a.Reason, "Reason")
	}
	r.Equal(e.AvgPrice, a.AvgPrice, "AvgPrice")
	r.Equal(e.Source, a.Source, "Source")
	r.Equal(e.ClientOrderId, a.ClientOrderId, "ClientOrderId")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Mmp, a.Mmp, "Mmp")
	if e.LastTrade != nil {
		r.Equal(e.LastTrade.Id, a.LastTrade.Id, "LastTrade.Id")
		r.Equal(e.LastTrade.TradeId, a.LastTrade.TradeId, "LastTrade.TradeId")
		r.Equal(e.LastTrade.Time, a.LastTrade.Time, "LastTrade.Time")
		r.Equal(e.LastTrade.Price, a.LastTrade.Price, "LastTrade.Price")
		r.Equal(e.LastTrade.Qty, a.LastTrade.Qty, "LastTrade.Qty")
	}
}

func (s *baseOrderTestSuite) assertOrdersEqual(e, a []*Order) {
	for i := range e {
		s.assertOrderEqual(e[i], a[i])
	}
}

func (s *baseOrderTestSuite) assertOrderAndAPIErrorListEqual(e, a []any) {
	for i := range e {
		switch ee := e[i].(type) {
		case Order:
			aa, ok := a[i].(Order)
			s.r().Equal(true, ok, "convert Order failed")
			s.assertOrderEqual(&ee, &aa)
		case common.APIError:
			aa, ok := a[i].(common.APIError)
			s.r().Equal(true, ok, "convert APIError failed")
			s.r().Equal(ee.Code, aa.Code, "Code")
			s.r().Equal(ee.Message, aa.Message, "Message")
		}
	}
}

func (s *orderServiceTestSuite) TestCreateOrder() {
	data := []byte(`{
		"orderId": 4729003411963445248,
		"symbol": "DOGE-240607-0.158-C",
		"price": "4.2000",
		"quantity": "0.01",
		"executedQty": "0.00",
		"fee": "0",
		"side": "BUY",
		"type": "LIMIT",
		"timeInForce": "GTC",
		"reduceOnly": false,
		"postOnly": false,
		"createTime": 1717666659088,
		"updateTime": 1717666659088,
		"status": "ACCEPTED",
		"avgPrice": "0",
		"source": "API",
		"clientOrderId": "053023",
		"priceScale": 4,
		"quantityScale": 2,
		"optionSide": "CALL",
		"quoteAsset": "USDT",
		"mmp": false
	   }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"
	side := SideTypeBuy
	orderType := OrderTypeLimit
	quantity := "0.01"
	price := "4.2000"
	timeInForce := TimeInForceTypeGTC
	reduceOnly := false
	postOnly := false
	newOrderResponseType := NewOrderRespTypeRESULT
	clientOrderId := "053023"
	isMmp := false

	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"quantity":         quantity,
			"price":            price,
			"timeInForce":      timeInForce,
			"reduceOnly":       reduceOnly,
			"postOnly":         postOnly,
			"newOrderRespType": newOrderResponseType,
			"clientOrderId":    clientOrderId,
			"isMmp":            isMmp,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).Price(price).TimeInForce(timeInForce).
		ReduceOnly(reduceOnly).PostOnly(postOnly).NewOrderResponseType(newOrderResponseType).
		ClientOrderId(clientOrderId).IsMmp(isMmp).
		Do(newContext())
	s.r().NoError(err)

	e := &Order{
		OrderId:       4729003411963445248,
		Symbol:        "DOGE-240607-0.158-C",
		Price:         "4.2000",
		Quantity:      "0.01",
		ExecutedQty:   "0.00",
		Fee:           "0",
		Side:          "BUY",
		Type:          "LIMIT",
		TimeInForce:   "GTC",
		ReduceOnly:    false,
		PostOnly:      false,
		CreateTime:    1717666659088,
		UpdateTime:    1717666659088,
		Status:        "ACCEPTED",
		AvgPrice:      "0",
		Source:        "API",
		ClientOrderId: "053023",
		PriceScale:    4,
		QuantityScale: 2,
		OptionSide:    "CALL",
		QuoteAsset:    "USDT",
		Mmp:           false,
	}
	s.assertOrderEqual(e, res)
}

func (s *orderServiceTestSuite) TestCreateBatchOrders() {
	data := []byte(`[
		{
		 "orderId": 4710989013445263360,
		 "symbol": "DOGE-240607-0.158-C",
		 "price": "4.2000",
		 "quantity": "0.01",
		 "executedQty": "0.00",
		 "fee": "0",
		 "side": "BUY",
		 "type": "LIMIT",
		 "timeInForce": "GTC",
		 "reduceOnly": false,
		 "postOnly": false,
		 "createTime": 1717666656964,
		 "updateTime": 1717666656964,
		 "status": "ACCEPTED",
		 "avgPrice": "0",
		 "source": "API",
		 "clientOrderId": "053020",
		 "priceScale": 4,
		 "quantityScale": 2,
		 "optionSide": "CALL",
		 "quoteAsset": "USDT",
		 "mmp": false
		},
		{
		 "orderId": 4710989013445263361,
		 "symbol": "DOGE-240607-0.158-C",
		 "price": "4.2000",
		 "quantity": "0.01",
		 "executedQty": "0.00",
		 "fee": "0",
		 "side": "BUY",
		 "type": "LIMIT",
		 "timeInForce": "GTC",
		 "reduceOnly": false,
		 "postOnly": false,
		 "createTime": 1717666656964,
		 "updateTime": 1717666656964,
		 "status": "ACCEPTED",
		 "avgPrice": "0",
		 "source": "API",
		 "clientOrderId": "053021",
		 "priceScale": 4,
		 "quantityScale": 2,
		 "optionSide": "CALL",
		 "quoteAsset": "USDT",
		 "mmp": false
		},
		{
		 "code": 1002,
		 "msg": "test 1002"
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	strHelper := func(s string) *string {
		return &s
	}
	tifHelper := func(s TimeInForceType) *TimeInForceType {
		return &s
	}
	boolHelper := func(i bool) *bool {
		return &i
	}

	orderLists := []*CreateOrderService{
		{
			symbol:           "DOGE-240607-0.158-C",
			side:             SideTypeBuy,
			orderType:        OrderTypeLimit,
			quantity:         "0.01",
			price:            strHelper("4.2000"),
			timeInForce:      tifHelper(TimeInForceTypeGTC),
			reduceOnly:       boolHelper(false),
			postOnly:         boolHelper(false),
			newOrderRespType: NewOrderRespTypeRESULT,
			clientOrderId:    strHelper("053020"),
			isMmp:            boolHelper(false),
		},
		{
			symbol:           "DOGE-240607-0.158-C",
			side:             SideTypeBuy,
			orderType:        OrderTypeLimit,
			quantity:         "0.01",
			price:            strHelper("4.2000"),
			timeInForce:      tifHelper(TimeInForceTypeGTC),
			reduceOnly:       boolHelper(false),
			postOnly:         boolHelper(false),
			newOrderRespType: NewOrderRespTypeRESULT,
			clientOrderId:    strHelper("053021"),
			isMmp:            boolHelper(false),
		},
	}
	s.assertReq(func(r *request) {
		packup := [](map[string]any){}
		for _, cos := range orderLists {
			t := map[string]any{
				"symbol":           cos.symbol,
				"side":             string(cos.side),
				"type":             string(cos.orderType),
				"quantity":         cos.quantity,
				"newOrderRespType": string(cos.newOrderRespType),
			}
			if cos.price != nil {
				t["price"] = *cos.price
			}
			if cos.timeInForce != nil {
				t["timeInForce"] = string(*cos.timeInForce)
			}
			if cos.reduceOnly != nil {
				t["reduceOnly"] = *cos.reduceOnly
			}
			if cos.postOnly != nil {
				t["postOnly"] = *cos.postOnly
			}
			if cos.clientOrderId != nil {
				t["clientOrderId"] = *cos.clientOrderId
			}
			if cos.isMmp != nil {
				t["isMmp"] = *cos.isMmp
			}
			packup = append(packup, t)
		}
		tmp, err := json.Marshal(packup)
		s.r().NoError(err)
		e := newSignedRequest().setFormParams(params{
			"orders": string(tmp),
		})
		s.assertRequestEqual(e, r)
	})

	returnOrders, err := s.client.NewCreateBatchOrdersService().OrderList(orderLists).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(returnOrders, 3)

	orders := []any{
		Order{
			OrderId:       4710989013445263360,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.2000",
			Quantity:      "0.01",
			ExecutedQty:   "0.00",
			Fee:           "0",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666656964,
			Status:        "ACCEPTED",
			AvgPrice:      "0",
			Source:        "API",
			ClientOrderId: "053020",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
		Order{
			OrderId:       4710989013445263361,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.2000",
			Quantity:      "0.01",
			ExecutedQty:   "0.00",
			Fee:           "0",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666656964,
			Status:        "ACCEPTED",
			AvgPrice:      "0",
			Source:        "API",
			ClientOrderId: "053021",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
		common.APIError{
			Code:    1002,
			Message: "test 1002",
		},
	}
	s.assertOrderAndAPIErrorListEqual(orders, returnOrders)
}

func (s *orderServiceTestSuite) TestGetOrder() {
	data := []byte(`{
		"orderId": 4710989013445263360,
		"symbol": "DOGE-240607-0.158-C",
		"price": "4.2000",
		"quantity": "0.01",
		"executedQty": "0.00",
		"fee": "0",
		"side": "BUY",
		"type": "LIMIT",
		"timeInForce": "GTC",
		"reduceOnly": false,
		"postOnly": false,
		"createTime": 1717666656964,
		"updateTime": 1717666656964,
		"status": "ACCEPTED",
		"avgPrice": "0",
		"source": "API",
		"clientOrderId": "053020",
		"priceScale": 4,
		"quantityScale": 2,
		"optionSide": "CALL",
		"quoteAsset": "USDT",
		"mmp": false
	   }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"
	orderID := int64(4710989013445263360)
	clientOrderId := "053020"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":        symbol,
			"clientOrderId": clientOrderId,
			"orderId":       orderID,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetOrderService().Symbol(symbol).
		OrderId(orderID).ClientOrderId(clientOrderId).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Order{
		OrderId:       4710989013445263360,
		Symbol:        "DOGE-240607-0.158-C",
		Price:         "4.2000",
		Quantity:      "0.01",
		ExecutedQty:   "0.00",
		Fee:           "0",
		Side:          "BUY",
		Type:          "LIMIT",
		TimeInForce:   "GTC",
		ReduceOnly:    false,
		PostOnly:      false,
		CreateTime:    1717666656964,
		UpdateTime:    1717666656964,
		Status:        "ACCEPTED",
		AvgPrice:      "0",
		Source:        "API",
		ClientOrderId: "053020",
		PriceScale:    4,
		QuantityScale: 2,
		OptionSide:    "CALL",
		QuoteAsset:    "USDT",
		Mmp:           false,
	}
	s.assertOrderEqual(e, order)
}

func (s *orderServiceTestSuite) TestCancelOrder() {
	data := []byte(`{
		"orderId": 4701981814192529408,
		"symbol": "DOGE-240607-0.158-C",
		"price": "4.2000",
		"quantity": "0.01",
		"executedQty": "0.00",
		"fee": "0",
		"side": "BUY",
		"type": "LIMIT",
		"timeInForce": "GTC",
		"reduceOnly": false,
		"postOnly": false,
		"createTime": 1717666657454,
		"updateTime": 1717666658463,
		"status": "CANCELLED",
		"avgPrice": "0",
		"source": "API",
		"clientOrderId": "053022",
		"priceScale": 4,
		"quantityScale": 2,
		"optionSide": "CALL",
		"quoteAsset": "USDT",
		"mmp": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"
	orderID := int64(4701981814192529408)
	clientOrderId := "053022"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":        symbol,
			"orderId":       orderID,
			"clientOrderId": clientOrderId,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelOrderService().Symbol(symbol).
		OrderId(orderID).ClientOrderId(clientOrderId).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Order{
		OrderId:       4701981814192529408,
		Symbol:        "DOGE-240607-0.158-C",
		Price:         "4.2000",
		Quantity:      "0.01",
		ExecutedQty:   "0.00",
		Fee:           "0",
		Side:          "BUY",
		Type:          "LIMIT",
		TimeInForce:   "GTC",
		ReduceOnly:    false,
		PostOnly:      false,
		CreateTime:    1717666657454,
		UpdateTime:    1717666658463,
		Status:        "CANCELLED",
		AvgPrice:      "0",
		Source:        "API",
		ClientOrderId: "053022",
		PriceScale:    4,
		QuantityScale: 2,
		OptionSide:    "CALL",
		QuoteAsset:    "USDT",
		Mmp:           false,
	}
	s.assertOrderEqual(e, res)
}

func (s *orderServiceTestSuite) TestCancelBatchOrders() {
	data := []byte(`[
		{
		 "orderId": 4710989013445263361,
		 "symbol": "DOGE-240607-0.158-C",
		 "price": "4.2000",
		 "quantity": "0.01",
		 "executedQty": "0.00",
		 "fee": "0",
		 "side": "BUY",
		 "type": "LIMIT",
		 "timeInForce": "GTC",
		 "reduceOnly": false,
		 "postOnly": false,
		 "createTime": 1717666656964,
		 "updateTime": 1717666658789,
		 "status": "CANCELLED",
		 "avgPrice": "0",
		 "source": "API",
		 "clientOrderId": "053021",
		 "priceScale": 4,
		 "quantityScale": 2,
		 "optionSide": "CALL",
		 "quoteAsset": "USDT",
		 "mmp": false
		},
		{
		 "orderId": 4710989013445263360,
		 "symbol": "DOGE-240607-0.158-C",
		 "price": "4.2000",
		 "quantity": "0.01",
		 "executedQty": "0.00",
		 "fee": "0",
		 "side": "BUY",
		 "type": "LIMIT",
		 "timeInForce": "GTC",
		 "reduceOnly": false,
		 "postOnly": false,
		 "createTime": 1717666656964,
		 "updateTime": 1717666658789,
		 "status": "CANCELLED",
		 "avgPrice": "0",
		 "source": "API",
		 "clientOrderId": "053020",
		 "priceScale": 4,
		 "quantityScale": 2,
		 "optionSide": "CALL",
		 "quoteAsset": "USDT",
		 "mmp": false
		},
		{
		 "code": 1002,
		 "msg": "test 1002"
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"
	orderIds := []int64{4710989013445263360, 4710989013445263361}
	clientOrderIds := []string{"053021", "053020"}

	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol":         symbol,
			"orderIds":       "[4710989013445263360,4710989013445263361]",
			"clientOrderIds": "[\"053021\",\"053020\"]",
		})
		//expectClientOrderIds := "[\"053021\",\"053020\"]"
		//s.r().Equal(expectClientOrderIds, r.form["clientOrderIds"])
		//r.form.Del("clientOrderIds")
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCancelBatchOrdersService().Symbol(symbol).OrderIds(
		orderIds).ClientOrderIds(clientOrderIds).Do(newContext())
	r := s.r()
	r.NoError(err)

	e := []any{
		Order{
			OrderId:       4710989013445263361,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.2000",
			Quantity:      "0.01",
			ExecutedQty:   "0.00",
			Fee:           "0",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666658789,
			Status:        "CANCELLED",
			AvgPrice:      "0",
			Source:        "API",
			ClientOrderId: "053021",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
		Order{
			OrderId:       4710989013445263360,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.2000",
			Quantity:      "0.01",
			ExecutedQty:   "0.00",
			Fee:           "0",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666658789,
			Status:        "CANCELLED",
			AvgPrice:      "0",
			Source:        "API",
			ClientOrderId: "053020",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
		common.APIError{
			Code:    1002,
			Message: "test 1002",
		},
	}
	s.assertOrderAndAPIErrorListEqual(e, res)
}

func (s *orderServiceTestSuite) TestCancelAllOpenOrders() {
	data := []byte(`{
		"code": "0",
		"msg": "success"
	   }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	rsp, err := s.client.NewCancelAllOpenOrdersService().Symbol(symbol).
		Do(newContext())
	targetRsp := &CancelAllOpenOrdersRsp{
		Code: "0",
		Msg:  "success",
	}

	s.r().NoError(err)
	s.r().Equal(rsp.Code, targetRsp.Code, "Code")
	s.r().Equal(rsp.Msg, targetRsp.Msg, "Msg")
}

func (s *orderServiceTestSuite) TestCancelAllOpenOrdersByUnderlying() {
	data := []byte(`{
		"code": "0",
		"msg": "success"
	   }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	underlying := "DOGEUSDT"

	res, err := s.client.NewCancelAllOpenOrdersByUnderlyingService().Underlying(
		underlying).Do(newContext())
	r := s.r()
	r.NoError(err)

	e := &CancelAllOpenOrdersByUnderlyingRsp{
		Code: "0",
		Msg:  "success",
	}
	s.r().Equal(e.Code, res.Code, "Code")
	s.r().Equal(e.Msg, res.Msg, "Msg")
	if res.Data != nil {
		s.r().Equal(*e.Data, *res.Data, "Data")
	}
}

func (s *orderServiceTestSuite) TestListOpenOrders() {
	data := []byte(`[
		{
			"orderId": 4710989013445263361,
			"symbol": "DOGE-240607-0.158-C",
			"price": "4.2000",
			"quantity": "0.01",
			"executedQty": "0.00",
			"fee": "0",
			"side": "BUY",
			"type": "LIMIT",
			"timeInForce": "GTC",
			"reduceOnly": false,
			"postOnly": false,
			"createTime": 1717666656964,
			"updateTime": 1717666656964,
			"status": "ACCEPTED",
			"avgPrice": "0",
			"source": "API",
			"clientOrderId": "053021",
			"priceScale": 4,
			"quantityScale": 2,
			"optionSide": "CALL",
			"quoteAsset": "USDT",
			"mmp": false
		   },
		   {
			"orderId": 4710989013445263360,
			"symbol": "DOGE-240607-0.158-C",
			"price": "4.2000",
			"quantity": "0.01",
			"executedQty": "0.00",
			"fee": "0",
			"side": "BUY",
			"type": "LIMIT",
			"timeInForce": "GTC",
			"reduceOnly": false,
			"postOnly": false,
			"createTime": 1717666656964,
			"updateTime": 1717666656964,
			"status": "ACCEPTED",
			"avgPrice": "0",
			"source": "API",
			"clientOrderId": "053020",
			"priceScale": 4,
			"quantityScale": 2,
			"optionSide": "CALL",
			"quoteAsset": "USDT",
			"mmp": false
		   }]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"
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
	r.Len(orders, 2)
	e := []*Order{
		{
			OrderId:       4710989013445263361,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.2000",
			Quantity:      "0.01",
			ExecutedQty:   "0.00",
			Fee:           "0",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666656964,
			Status:        "ACCEPTED",
			AvgPrice:      "0",
			Source:        "API",
			ClientOrderId: "053021",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
		{
			OrderId:       4710989013445263360,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.2000",
			Quantity:      "0.01",
			ExecutedQty:   "0.00",
			Fee:           "0",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666656964,
			Status:        "ACCEPTED",
			AvgPrice:      "0",
			Source:        "API",
			ClientOrderId: "053020",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
	}
	s.assertOrdersEqual(e, orders)
}

func (s *orderServiceTestSuite) TestHistoryOrders() {
	data := []byte(`[	
			{
			"orderId": 4710989013445263360,
			"symbol": "DOGE-240607-0.158-C",
			"price": "4.20000000",
			"quantity": "0.01000000",
			"executedQty": "0.00000000",
			"fee": "0.00000000",
			"side": "BUY",
			"type": "LIMIT",
			"timeInForce": "GTC",
			"reduceOnly": false,
			"postOnly": false,
			"createTime": 1717666656964,
			"updateTime": 1717666658789,
			"status": "CANCELLED",
			"reason": "1",
			"avgPrice": "0.00000000",
			"source": "API",
			"clientOrderId": "053020",
			"priceScale": 4,
			"quantityScale": 2,
			"optionSide": "CALL",
			"quoteAsset": "USDT",
			"mmp": false
			},
			{
			"orderId": 4710989013445263361,
			"symbol": "DOGE-240607-0.158-C",
			"price": "4.20000000",
			"quantity": "0.01000000",
			"executedQty": "0.00000000",
			"fee": "0.00000000",
			"side": "BUY",
			"type": "LIMIT",
			"timeInForce": "GTC",
			"reduceOnly": false,
			"postOnly": false,
			"createTime": 1717666656964,
			"updateTime": 1717666658789,
			"status": "CANCELLED",
			"reason": "1",
			"avgPrice": "0.00000000",
			"source": "API",
			"clientOrderId": "053021",
			"priceScale": 4,
			"quantityScale": 2,
			"optionSide": "CALL",
			"quoteAsset": "USDT",
			"mmp": false
			}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240607-0.158-C"

	orders, err := s.client.NewHistoryOrdersService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(orders, 2)

	strHelper := func(s string) *string {
		return &s
	}

	e := []*Order{
		{
			OrderId:       4710989013445263360,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.20000000",
			Quantity:      "0.01000000",
			ExecutedQty:   "0.00000000",
			Fee:           "0.00000000",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666658789,
			Status:        "CANCELLED",
			Reason:        strHelper("1"),
			AvgPrice:      "0.00000000",
			Source:        "API",
			ClientOrderId: "053020",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
		{
			OrderId:       4710989013445263361,
			Symbol:        "DOGE-240607-0.158-C",
			Price:         "4.20000000",
			Quantity:      "0.01000000",
			ExecutedQty:   "0.00000000",
			Fee:           "0.00000000",
			Side:          "BUY",
			Type:          "LIMIT",
			TimeInForce:   "GTC",
			ReduceOnly:    false,
			PostOnly:      false,
			CreateTime:    1717666656964,
			UpdateTime:    1717666658789,
			Status:        "CANCELLED",
			Reason:        strHelper("1"),
			AvgPrice:      "0.00000000",
			Source:        "API",
			ClientOrderId: "053021",
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Mmp:           false,
		},
	}
	s.assertOrdersEqual(e, orders)
}

func (s *orderServiceTestSuite) assertPositionEqual(e, a *Position) {
	r := s.r()
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.ReducibleQty, a.ReducibleQty, "ReducibleQty")
	r.Equal(e.MarkValue, a.MarkValue, "MarkValue")
	r.Equal(e.Ror, a.Ror, "Ror")
	r.Equal(e.UnrealizedPNL, a.UnrealizedPNL, "UnrealizedPNL")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.StrikePrice, a.StrikePrice, "StrikePrice")
	r.Equal(e.ExpiryDate, a.ExpiryDate, "ExpiryDate")
	r.Equal(e.PositionCost, a.PositionCost, "PositionCost")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Time, a.Time, "Time")
}

func (s *orderServiceTestSuite) assertPositionsEqual(e, a []*Position) {
	for idx := range e {
		s.assertPositionEqual(e[idx], a[idx])
	}
}

func (s *orderServiceTestSuite) TestPosition() {
	data := []byte(`[
		{
		 "entryPrice": "24.3",
		 "symbol": "DOGE-240531-0.14-C",
		 "side": "LONG",
		 "quantity": "0.01",
		 "reducibleQty": "0.01",
		 "markValue": "0.216621",
		 "ror": "-0.1085",
		 "unrealizedPNL": "-0.026379",
		 "markPrice": "21.6621",
		 "strikePrice": "0.14000000",
		 "positionCost": "0.243",
		 "expiryDate": 1717142400000,
		 "priceScale": 4,
		 "quantityScale": 2,
		 "optionSide": "CALL",
		 "quoteAsset": "USDT",
		 "time": 1717047449679
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240531-0.14-C"

	positions, err := s.client.NewPositionService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(positions, 1)

	e := []*Position{
		{
			EntryPrice:    "24.3",
			Symbol:        "DOGE-240531-0.14-C",
			Side:          "LONG",
			Quantity:      "0.01",
			ReducibleQty:  "0.01",
			MarkValue:     "0.216621",
			Ror:           "-0.1085",
			UnrealizedPNL: "-0.026379",
			MarkPrice:     "21.6621",
			StrikePrice:   "0.14000000",
			PositionCost:  "0.243",
			ExpiryDate:    1717142400000,
			PriceScale:    4,
			QuantityScale: 2,
			OptionSide:    "CALL",
			QuoteAsset:    "USDT",
			Time:          1717047449679,
		},
	}
	s.assertPositionsEqual(e, positions)
}

func (s *orderServiceTestSuite) assertUserTrade(e, a *UserTrade) {
	r := s.r()
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.TradeId, a.TradeId, "TradeId")
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.RealizedProfit, a.RealizedProfit, "RealizedProfit")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Volatility, a.Volatility, "Volatility")
	r.Equal(e.Liquidity, a.Liquidity, "Liquidity")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
}

func (s *orderServiceTestSuite) assertUserTrades(e, a []*UserTrade) {
	for idx := range e {
		s.assertUserTrade(e[idx], a[idx])
	}
}

func (s *orderServiceTestSuite) TestUserTrades() {
	data := []byte(`[
		{
		 "id": 1125899906892495057,
		 "tradeId": 62,
		 "orderId": 4729000875681705984,
		 "symbol": "DOGE-240531-0.14-C",
		 "price": "24.30000000",
		 "quantity": "0.01000000",
		 "fee": "0.00049172",
		 "realizedProfit": "0.00000000",
		 "side": "BUY",
		 "type": "LIMIT",
		 "volatility": "1.10148926",
		 "liquidity": "TAKER",
		 "time": 1717047449679,
		 "priceScale": 4,
		 "quantityScale": 2,
		 "optionSide": "CALL",
		 "quoteAsset": "USDT"
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240531-0.14-C"

	positions, err := s.client.NewUserTradesService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(positions, 1)

	e := []*UserTrade{
		{
			Id:             1125899906892495057,
			TradeId:        62,
			OrderId:        4729000875681705984,
			Symbol:         "DOGE-240531-0.14-C",
			Price:          "24.30000000",
			Quantity:       "0.01000000",
			Fee:            "0.00049172",
			RealizedProfit: "0.00000000",
			Side:           "BUY",
			Type:           "LIMIT",
			Volatility:     "1.10148926",
			Liquidity:      "TAKER",
			Time:           1717047449679,
			PriceScale:     4,
			QuantityScale:  2,
			OptionSide:     "CALL",
			QuoteAsset:     "USDT",
		},
	}
	s.assertUserTrades(e, positions)
}

func (s *orderServiceTestSuite) assertExerciseRecordEqual(e, a *ExerciseRecord) {
	r := s.r()
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.Currency, a.Currency, "Currency")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ExercisePrice, a.ExercisePrice, "ExercisePrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Fee, a.Fee, "Fee")
	r.Equal(e.CreateDate, a.CreateDate, "CreateDate")
	r.Equal(e.PriceScale, a.PriceScale, "PriceScale")
	r.Equal(e.QuantityScale, a.QuantityScale, "QuantityScale")
	r.Equal(e.OptionSide, a.OptionSide, "OptionSide")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
}

func (s *orderServiceTestSuite) assertExerciseRecordsEqual(e, a []*ExerciseRecord) {
	for idx := range e {
		s.assertExerciseRecordEqual(e[idx], a[idx])
	}
}

func (s *orderServiceTestSuite) TestExerciseRecord() {
	data := []byte(`[
		{
			"id": "1125899906842624042",        
			"currency": "USDT",                  
			"symbol": "DOGE-240531-0.14-C",      
			"exercisePrice": "25000.00000000",   
			"markPrice": "25000.00000000",       
			"quantity": "1.00000000",            
			"amount": "0.00000000",              
			"fee": "0.00000000",                 
			"createDate": 1658361600000,        
			"priceScale": 2,                     
			"quantityScale": 2,                  
			"optionSide": "CALL",               
			"positionSide": "LONG",              
			"quoteAsset": "USDT"                
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "DOGE-240531-0.14-C"

	exercises, err := s.client.NewExercistRecordService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(exercises, 1)

	e := []*ExerciseRecord{
		{
			Id:            "1125899906842624042",
			Currency:      "USDT",
			Symbol:        "DOGE-240531-0.14-C",
			ExercisePrice: "25000.00000000",
			MarkPrice:     "25000.00000000",
			Quantity:      "1.00000000",
			Amount:        "0.00000000",
			Fee:           "0.00000000",
			CreateDate:    1658361600000,
			PriceScale:    2,
			QuantityScale: 2,
			OptionSide:    "CALL",
			PositionSide:  "LONG",
			QuoteAsset:    "USDT",
		},
	}
	s.assertExerciseRecordsEqual(e, exercises)
}

func (s *orderServiceTestSuite) assertBillEqual(e, a *Bill) {
	r := s.r()
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.CreateDate, a.CreateDate, "CreateDate")
}

func (s *orderServiceTestSuite) TestBill() {
	data := []byte(`[
		{
		 "id": "1125899906873951554",
		 "asset": "USDT",
		 "amount": "-0.00049172",
		 "type": "FEE",
		 "createDate": 1717047449679
		},
		{
		 "id": "1125899906873951553",
		 "asset": "USDT",
		 "amount": "-0.24300000",
		 "type": "CONTRACT",
		 "createDate": 1717047449679
		}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	bills, err := s.client.NewBillService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(bills, 2)

	e := []*Bill{
		{
			Id:         "1125899906873951554",
			Asset:      "USDT",
			Amount:     "-0.00049172",
			Type:       "FEE",
			CreateDate: 1717047449679,
		},
		{
			Id:         "1125899906873951553",
			Asset:      "USDT",
			Amount:     "-0.24300000",
			Type:       "CONTRACT",
			CreateDate: 1717047449679,
		},
	}
	for idx := range e {
		s.assertBillEqual(e[idx], bills[idx])
	}
}

func (s *orderServiceTestSuite) TestIncomeDownloadIdService() {
	data := []byte(`{
		"avgCostTimestampOfLast30d":7241837, 
		"downloadId":"546975389218332672"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	downloadId, err := s.client.NewIncomeDownloadIdService().Do(newContext())
	r := s.r()
	r.NoError(err)

	e := &IncomeDownloadId{
		AvgCostTimestampOfLast30d: 7241837,
		DownloadId:                "546975389218332672",
	}

	s.r().Equal(e.AvgCostTimestampOfLast30d, downloadId.AvgCostTimestampOfLast30d, "AvgCostTimestampOfLast30d")
	s.r().Equal(e.DownloadId, downloadId.DownloadId, "DownloadId")
	if e.Error != nil {
		s.r().Equal(e.Error, downloadId.Error, "Error")
	}
}

func (s *orderServiceTestSuite) TestIncomeDownloadLink() {
	data := []byte(`{
		"downloadId":"545923594199212032",
		"status":"completed",     
		"url":"www.binance.com",  
		"notified":true,          
		"expirationTimestamp":1645009771000, 
		"isExpired":null
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	link, err := s.client.NewIncomeDownloadLinkService().DownloadId("546975389218332672").Do(newContext())
	r := s.r()
	r.NoError(err)

	e := &IncomeDownloadLink{
		DownloadId:          "545923594199212032",
		Status:              "completed",
		Url:                 "www.binance.com",
		Notified:            true,
		ExpirationTimestamp: 1645009771000,
		IsExpired:           nil,
	}

	s.r().Equal(e.DownloadId, link.DownloadId, "DownloadId")
	s.r().Equal(e.Status, link.Status, "Status")
	s.r().Equal(e.Url, link.Url, "Url")
	s.r().Equal(e.Notified, link.Notified, "Notified")
	s.r().Equal(e.ExpirationTimestamp, link.ExpirationTimestamp, "ExpirationTimestamp")
	s.r().Equal(e.IsExpired, link.IsExpired, "IsExpired")
}
