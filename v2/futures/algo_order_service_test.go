package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// algoOrderServiceTestSuite algo order service test suite
type algoOrderServiceTestSuite struct {
	baseTestSuite
}

// TestAlgoOrderService run algo order service test suite
func TestAlgoOrderService(t *testing.T) {
	suite.Run(t, new(algoOrderServiceTestSuite))
}

// TestCreateAlgoOrderService test create algo order service
func (s *algoOrderServiceTestSuite) TestCreateAlgoOrderService() {
	// mock API response
	data := []byte(`{
		"algoId": 2146760,
		"clientAlgoId": "6B2I9XVcJpCjqPAJ4YoFX7",
		"algoType": "CONDITIONAL",
		"orderType": "TAKE_PROFIT",
		"symbol": "BNBUSDT",
		"side": "SELL",
		"positionSide": "BOTH",
		"timeInForce": "GTC",
		"quantity": "0.01",
		"algoStatus": "NEW",
		"triggerPrice": "750.000",
		"price": "750.000",
		"icebergQuantity": null,
		"selfTradePreventionMode": "EXPIRE_MAKER",
		"workingType": "CONTRACT_PRICE",
		"priceMatch": "NONE",
		"closePosition": false,
		"priceProtect": false,
		"reduceOnly": false,
		"activatePrice": "",
		"callbackRate": "",
		"createTime": 1750485492076,
		"updateTime": 1750485492076,
		"triggerTime": 0,
		"goodTillDate": 0
	}`)

	// set mock response
	s.mockDo(data, nil)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewCreateAlgoOrderService().
		AlgoType(OrderAlgoTypeConditional).
		Symbol("BNBUSDT").
		Side(SideTypeSell).
		Type(AlgoOrderTypeTakeProfit).
		Quantity("0.01").
		Price("750.000").
		TriggerPrice("750.000").
		Do(newContext())

	// verify response
	s.r().NoError(err)
	e := &CreateAlgoOrderResp{
		AlgoId:                  2146760,
		ClientAlgoId:            "6B2I9XVcJpCjqPAJ4YoFX7",
		AlgoType:                OrderAlgoTypeConditional,
		OrderType:               AlgoOrderTypeTakeProfit,
		Symbol:                  "BNBUSDT",
		Side:                    SideTypeSell,
		PositionSide:            PositionSideTypeBoth,
		TimeInForce:             TimeInForceTypeGTC,
		Quantity:                "0.01",
		AlgoStatus:              AlgoOrderStatusTypeNew,
		TriggerPrice:            "750.000",
		Price:                   "750.000",
		IcebergQuantity:         nil,
		SelfTradePreventionMode: SelfTradePreventionModeExpireMaker,
		WorkingType:             WorkingTypeContractPrice,
		PriceMatch:              PriceMatchTypeNone,
		ClosePosition:           false,
		PriceProtect:            false,
		ReduceOnly:              false,
		ActivatePrice:           "",
		CallbackRate:            "",
		CreateTime:              1750485492076,
		UpdateTime:              1750485492076,
		TriggerTime:             0,
		GoodTillDate:            0,
	}
	s.assertCreateAlgoOrderRespEqual(e, resp)
}

// TestCreateAlgoOrderServiceError test create algo order service error
func (s *algoOrderServiceTestSuite) TestCreateAlgoOrderServiceError() {
	// mock API error response
	data := []byte(`{
		"code": -1102,
		"msg": "Invalid symbol"
	}`)

	// set mock response
	s.mockDo(data, nil, 400)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewCreateAlgoOrderService().
		AlgoType(OrderAlgoTypeConditional).
		Symbol("INVALID_SYMBOL").
		Side(SideTypeSell).
		Type(AlgoOrderTypeTakeProfit).
		Quantity("0.01").
		Price("750.000").
		TriggerPrice("750.000").
		Do(newContext())

	// verify response
	s.r().Error(err)
	s.r().Nil(resp)
	s.r().Contains(err.Error(), "Invalid symbol")
}

// TestCancelAlgoOrderService test cancel algo order service
func (s *algoOrderServiceTestSuite) TestCancelAlgoOrderService() {
	// mock API response
	data := []byte(`{
		"algoId": 2146760,
		"clientAlgoId": "6B2I9XVcJpCjqPAJ4YoFX7",
		"code": "0",
		"msg": "success"
	}`)

	// set mock response
	s.mockDo(data, nil)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewCancelAlgoOrderService().
		AlgoID(2146760).
		Do(newContext())

	// verify response
	s.r().NoError(err)
	e := &CancelAlgoOrderResp{
		AlgoId:       2146760,
		ClientAlgoId: "6B2I9XVcJpCjqPAJ4YoFX7",
		Code:         "0",
		Message:      "success",
	}
	s.assertCancelAlgoOrderRespEqual(e, resp)
}

// TestCancelAlgoOrderServiceError test cancel algo order service error
func (s *algoOrderServiceTestSuite) TestCancelAlgoOrderServiceError() {
	// mock API error response
	data := []byte(`{
		"code": -1102,
		"msg": "Invalid algoId"
	}`)

	// set mock response
	s.mockDo(data, nil, 400)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewCancelAlgoOrderService().
		AlgoID(9999999).
		Do(newContext())

	// verify response
	s.r().Error(err)
	s.r().Nil(resp)
	s.r().Contains(err.Error(), "Invalid algoId")
}

// TestCancelAllAlgoOpenOrdersService test cancel all algo open orders service
func (s *algoOrderServiceTestSuite) TestCancelAllAlgoOpenOrdersService() {
	// mock API response
	data := []byte(`{
		"code": 200,
		"msg": "success"
	}`)

	// set mock response
	s.mockDo(data, nil)
	defer s.assertDo()

	// execute request
	err := s.client.NewCancelAllAlgoOpenOrdersService().
		Symbol("BNBUSDT").
		Do(newContext())

	// verify response
	s.r().NoError(err)
}

// TestCancelAllAlgoOpenOrdersServiceError test cancel all algo open orders service error
func (s *algoOrderServiceTestSuite) TestCancelAllAlgoOpenOrdersServiceError() {
	// mock API error response
	data := []byte(`{
		"code": -1102,
		"msg": "Invalid symbol"
	}`)

	// set mock response
	s.mockDo(data, nil, 400)
	defer s.assertDo()

	// execute request
	err := s.client.NewCancelAllAlgoOpenOrdersService().
		Symbol("INVALID_SYMBOL").
		Do(newContext())

	// verify response
	s.r().Error(err)
	s.r().Contains(err.Error(), "Invalid symbol")
}

// TestGetAlgoOrderService test get algo order service
func (s *algoOrderServiceTestSuite) TestGetAlgoOrderService() {
	// mock API response
	data := []byte(`{
		"algoId": 2146760,
		"clientAlgoId": "6B2I9XVcJpCjqPAJ4YoFX7",
		"algoType": "CONDITIONAL",
		"orderType": "TAKE_PROFIT",
		"symbol": "BNBUSDT",
		"side": "SELL",
		"positionSide": "BOTH",
		"timeInForce": "GTC",
		"quantity": "0.01",
		"algoStatus": "CANCELED",
		"actualOrderId": "",
		"actualPrice": "0.00000",
		"triggerPrice": "750.000",
		"price": "750.000",
		"icebergQuantity": null,
		"tpTriggerPrice": "0.000",
		"tpPrice": "0.000",
		"slTriggerPrice": "0.000",
		"slPrice": "0.000",
		"tpOrderType": "",
		"selfTradePreventionMode": "EXPIRE_MAKER",
		"workingType": "CONTRACT_PRICE",
		"priceMatch": "NONE",
		"closePosition": false,
		"priceProtect": false,
		"reduceOnly": false,
		"createTime": 1750485492076,
		"updateTime": 1750514545091,
		"triggerTime": 0,
		"goodTillDate": 0
	}`)

	// set mock response
	s.mockDo(data, nil)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewGetAlgoOrderService().
		AlgoID(2146760).
		Do(newContext())

	// verify response
	s.r().NoError(err)
	e := &GetAlgoOrderResp{
		AlgoId:                  2146760,
		ClientAlgoId:            "6B2I9XVcJpCjqPAJ4YoFX7",
		AlgoType:                OrderAlgoTypeConditional,
		OrderType:               AlgoOrderTypeTakeProfit,
		Symbol:                  "BNBUSDT",
		Side:                    SideTypeSell,
		PositionSide:            PositionSideTypeBoth,
		TimeInForce:             TimeInForceTypeGTC,
		Quantity:                "0.01",
		AlgoStatus:              AlgoOrderStatusTypeCanceled,
		ActualOrderId:           "",
		ActualPrice:             "0.00000",
		TriggerPrice:            "750.000",
		Price:                   "750.000",
		IcebergQuantity:         nil,
		TpTriggerPrice:          "0.000",
		TpPrice:                 "0.000",
		SlTriggerPrice:          "0.000",
		SlPrice:                 "0.000",
		TpOrderType:             "",
		SelfTradePreventionMode: SelfTradePreventionModeExpireMaker,
		WorkingType:             WorkingTypeContractPrice,
		PriceMatch:              PriceMatchTypeNone,
		ClosePosition:           false,
		PriceProtect:            false,
		ReduceOnly:              false,
		CreateTime:              1750485492076,
		UpdateTime:              1750514545091,
		TriggerTime:             0,
		GoodTillDate:            0,
	}
	s.assertGetAlgoOrderRespEqual(e, resp)
}

// TestGetAlgoOrderServiceError test get algo order service error
func (s *algoOrderServiceTestSuite) TestGetAlgoOrderServiceError() {
	// mock API error response
	data := []byte(`{
		"code": -1102,
		"msg": "Invalid algoId"
	}`)

	// set mock response
	s.mockDo(data, nil, 400)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewGetAlgoOrderService().
		AlgoID(9999999).
		Do(newContext())

	// verify response
	s.r().Error(err)
	s.r().Nil(resp)
	s.r().Contains(err.Error(), "Invalid algoId")
}

// assertGetAlgoOrderRespEqual assert GetAlgoOrderResp equal
func (s *algoOrderServiceTestSuite) assertGetAlgoOrderRespEqual(e, a *GetAlgoOrderResp) {
	r := s.r()
	r.Equal(e.AlgoId, a.AlgoId, "AlgoId")
	r.Equal(e.ClientAlgoId, a.ClientAlgoId, "ClientAlgoId")
	r.Equal(e.AlgoType, a.AlgoType, "AlgoType")
	r.Equal(e.OrderType, a.OrderType, "OrderType")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.AlgoStatus, a.AlgoStatus, "AlgoStatus")
	r.Equal(e.ActualOrderId, a.ActualOrderId, "ActualOrderId")
	r.Equal(e.ActualPrice, a.ActualPrice, "ActualPrice")
	r.Equal(e.TriggerPrice, a.TriggerPrice, "TriggerPrice")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.IcebergQuantity, a.IcebergQuantity, "IcebergQuantity")
	r.Equal(e.TpTriggerPrice, a.TpTriggerPrice, "TpTriggerPrice")
	r.Equal(e.TpPrice, a.TpPrice, "TpPrice")
	r.Equal(e.SlTriggerPrice, a.SlTriggerPrice, "SlTriggerPrice")
	r.Equal(e.SlPrice, a.SlPrice, "SlPrice")
	r.Equal(e.TpOrderType, a.TpOrderType, "TpOrderType")
	r.Equal(e.SelfTradePreventionMode, a.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.PriceMatch, a.PriceMatch, "PriceMatch")
	r.Equal(e.ClosePosition, a.ClosePosition, "ClosePosition")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.TriggerTime, a.TriggerTime, "TriggerTime")
	r.Equal(e.GoodTillDate, a.GoodTillDate, "GoodTillDate")
}

// TestGetOpenAlgoOrdersService test get open algo orders service
func (s *algoOrderServiceTestSuite) TestGetOpenAlgoOrdersService() {
	// mock API response
	data := []byte(`[
		{
			"algoId": 2146760,
			"clientAlgoId": "6B2I9XVcJpCjqPAJ4YoFX7",
			"algoType": "CONDITIONAL",
			"orderType": "TAKE_PROFIT",
			"symbol": "BNBUSDT",
			"side": "SELL",
			"positionSide": "BOTH",
			"timeInForce": "GTC",
			"quantity": "0.01",
			"algoStatus": "NEW",
			"actualOrderId": "",
			"actualPrice": "0.00000",
			"triggerPrice": "750.000",
			"price": "750.000",
			"icebergQuantity": null,
			"tpTriggerPrice": "0.000",
			"tpPrice": "0.000",
			"slTriggerPrice": "0.000",
			"slPrice": "0.000",
			"tpOrderType": "",
			"selfTradePreventionMode": "EXPIRE_MAKER",
			"workingType": "CONTRACT_PRICE",
			"priceMatch": "NONE",
			"closePosition": false,
			"priceProtect": false,
			"reduceOnly": false,
			"createTime": 1750485492076,
			"updateTime": 1750485492076,
			"triggerTime": 0,
			"goodTillDate": 0
		}
	]`)

	// set mock response
	s.mockDo(data, nil)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewListOpenAlgoOrdersService().
		Symbol("BNBUSDT").
		Do(newContext())

	// verify response
	s.r().NoError(err)
	s.r().Len(resp, 1)
	e := &GetAlgoOrderResp{
		AlgoId:                  2146760,
		ClientAlgoId:            "6B2I9XVcJpCjqPAJ4YoFX7",
		AlgoType:                OrderAlgoTypeConditional,
		OrderType:               AlgoOrderTypeTakeProfit,
		Symbol:                  "BNBUSDT",
		Side:                    SideTypeSell,
		PositionSide:            PositionSideTypeBoth,
		TimeInForce:             TimeInForceTypeGTC,
		Quantity:                "0.01",
		AlgoStatus:              AlgoOrderStatusTypeNew,
		ActualOrderId:           "",
		ActualPrice:             "0.00000",
		TriggerPrice:            "750.000",
		Price:                   "750.000",
		IcebergQuantity:         nil,
		TpTriggerPrice:          "0.000",
		TpPrice:                 "0.000",
		SlTriggerPrice:          "0.000",
		SlPrice:                 "0.000",
		TpOrderType:             "",
		SelfTradePreventionMode: SelfTradePreventionModeExpireMaker,
		WorkingType:             WorkingTypeContractPrice,
		PriceMatch:              PriceMatchTypeNone,
		ClosePosition:           false,
		PriceProtect:            false,
		ReduceOnly:              false,
		CreateTime:              1750485492076,
		UpdateTime:              1750485492076,
		TriggerTime:             0,
		GoodTillDate:            0,
	}
	s.assertGetAlgoOrderRespEqual(e, &resp[0])
}

// TestGetOpenAlgoOrdersServiceError test get open algo orders service error
func (s *algoOrderServiceTestSuite) TestGetOpenAlgoOrdersServiceError() {
	// mock API error response
	data := []byte(`{
		"code": -1102,
		"msg": "Invalid symbol"
	}`)

	// set mock response
	s.mockDo(data, nil, 400)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewListOpenAlgoOrdersService().
		Symbol("INVALID_SYMBOL").
		Do(newContext())

	// verify response
	s.r().Error(err)
	s.r().Nil(resp)
	s.r().Contains(err.Error(), "Invalid symbol")
}

// TestGetAllAlgoOrdersService test get all algo orders service
func (s *algoOrderServiceTestSuite) TestGetAllAlgoOrdersService() {
	// mock API response
	data := []byte(`[
		{
			"algoId": 2146760,
			"clientAlgoId": "6B2I9XVcJpCjqPAJ4YoFX7",
			"algoType": "CONDITIONAL",
			"orderType": "TAKE_PROFIT",
			"symbol": "BNBUSDT",
			"side": "SELL",
			"positionSide": "BOTH",
			"timeInForce": "GTC",
			"quantity": "0.01",
			"algoStatus": "CANCELED",
			"actualOrderId": "",
			"actualPrice": "0.00000",
			"triggerPrice": "750.000",
			"price": "750.000",
			"icebergQuantity": null,
			"tpTriggerPrice": "0.000",
			"tpPrice": "0.000",
			"slTriggerPrice": "0.000",
			"slPrice": "0.000",
			"tpOrderType": "",
			"selfTradePreventionMode": "EXPIRE_MAKER",
			"workingType": "CONTRACT_PRICE",
			"priceMatch": "NONE",
			"closePosition": false,
			"priceProtect": false,
			"reduceOnly": false,
			"createTime": 1750485492076,
			"updateTime": 1750514545091,
			"triggerTime": 0,
			"goodTillDate": 0
		}
	]`)

	// set mock response
	s.mockDo(data, nil)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewListAllAlgoOrdersService().Symbol("BNBUSDT").
		Limit(500).
		Do(newContext())

	// verify response
	s.r().NoError(err)
	s.r().Len(resp, 1)
	e := &GetAlgoOrderResp{
		AlgoId:                  2146760,
		ClientAlgoId:            "6B2I9XVcJpCjqPAJ4YoFX7",
		AlgoType:                OrderAlgoTypeConditional,
		OrderType:               AlgoOrderTypeTakeProfit,
		Symbol:                  "BNBUSDT",
		Side:                    SideTypeSell,
		PositionSide:            PositionSideTypeBoth,
		TimeInForce:             TimeInForceTypeGTC,
		Quantity:                "0.01",
		AlgoStatus:              AlgoOrderStatusTypeCanceled,
		ActualOrderId:           "",
		ActualPrice:             "0.00000",
		TriggerPrice:            "750.000",
		Price:                   "750.000",
		IcebergQuantity:         nil,
		TpTriggerPrice:          "0.000",
		TpPrice:                 "0.000",
		SlTriggerPrice:          "0.000",
		SlPrice:                 "0.000",
		TpOrderType:             "",
		SelfTradePreventionMode: SelfTradePreventionModeExpireMaker,
		WorkingType:             WorkingTypeContractPrice,
		PriceMatch:              PriceMatchTypeNone,
		ClosePosition:           false,
		PriceProtect:            false,
		ReduceOnly:              false,
		CreateTime:              1750485492076,
		UpdateTime:              1750514545091,
		TriggerTime:             0,
		GoodTillDate:            0,
	}
	s.assertGetAlgoOrderRespEqual(e, &resp[0])
}

// TestGetAllAlgoOrdersServiceError test get all algo orders service error
func (s *algoOrderServiceTestSuite) TestGetAllAlgoOrdersServiceError() {
	// mock API error response
	data := []byte(`{
		"code": -1102,
		"msg": "Invalid symbol"
	}`)

	// set mock response
	s.mockDo(data, nil, 400)
	defer s.assertDo()

	// execute request
	resp, err := s.client.NewListAllAlgoOrdersService().Symbol("INVALID_SYMBOL").
		Limit(500).
		Do(newContext())

	// verify response
	s.r().Error(err)
	s.r().Nil(resp)
	s.r().Contains(err.Error(), "Invalid symbol")
}

// assertCreateAlgoOrderRespEqual assert CreateAlgoOrderResp equal
func (s *algoOrderServiceTestSuite) assertCreateAlgoOrderRespEqual(e, a *CreateAlgoOrderResp) {
	r := s.r()
	r.Equal(e.AlgoId, a.AlgoId, "AlgoId")
	r.Equal(e.ClientAlgoId, a.ClientAlgoId, "ClientAlgoId")
	r.Equal(e.AlgoType, a.AlgoType, "AlgoType")
	r.Equal(e.OrderType, a.OrderType, "OrderType")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.AlgoStatus, a.AlgoStatus, "AlgoStatus")
	r.Equal(e.TriggerPrice, a.TriggerPrice, "TriggerPrice")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.IcebergQuantity, a.IcebergQuantity, "IcebergQuantity")
	r.Equal(e.SelfTradePreventionMode, a.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.PriceMatch, a.PriceMatch, "PriceMatch")
	r.Equal(e.ClosePosition, a.ClosePosition, "ClosePosition")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.CallbackRate, a.CallbackRate, "CallbackRate")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.TriggerTime, a.TriggerTime, "TriggerTime")
	r.Equal(e.GoodTillDate, a.GoodTillDate, "GoodTillDate")
}

// assertCancelAlgoOrderRespEqual assert CancelAlgoOrderResp equal
func (s *algoOrderServiceTestSuite) assertCancelAlgoOrderRespEqual(e, a *CancelAlgoOrderResp) {
	r := s.r()
	r.Equal(e.AlgoId, a.AlgoId, "AlgoId")
	r.Equal(e.ClientAlgoId, a.ClientAlgoId, "ClientAlgoId")
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Message, a.Message, "Message")
}
