package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umConditionalOrderServiceTestSuite struct {
	baseTestSuite
}

func TestUMConditionalOrderService(t *testing.T) {
	suite.Run(t, new(umConditionalOrderServiceTestSuite))
}

func (s *umConditionalOrderServiceTestSuite) TestUMConditionalOrder() {
	data := []byte(`{
		"newClientStrategyId": "testOrder",
		"strategyId": 123445,
		"strategyStatus": "NEW",
		"strategyType": "TRAILING_STOP_MARKET",
		"origQty": "10",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"stopPrice": "9300",
		"symbol": "BTCUSDT",
		"timeInForce": "GTD",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"bookTime": 1566818724710,
		"updateTime": 1566818724722,
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	side := SideTypeBuy
	strategyType := "TRAILING_STOP_MARKET"
	callbackRate := "0.3"

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("side", side)
		e.setParam("strategyType", strategyType)
		e.setParam("callbackRate", callbackRate)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMConditionalOrderService().
		Symbol(symbol).
		Side(side).
		StrategyType(strategyType).
		CallbackRate(callbackRate).
		Do(newContext())

	s.r().NoError(err)
	e := &UMConditionalOrder{
		NewClientStrategyId: "testOrder",
		StrategyId:          123445,
		StrategyStatus:      "NEW",
		StrategyType:        "TRAILING_STOP_MARKET",
		OrigQty:             "10",
		Price:               "0",
		ReduceOnly:          false,
		Side:                SideTypeBuy,
		PositionSide:        PositionSideTypeShort,
		StopPrice:           "9300",
		Symbol:              "BTCUSDT",
		TimeInForce:         TimeInForceTypeGTD,
		ActivatePrice:       "9020",
		PriceRate:           "0.3",
		BookTime:            1566818724710,
		UpdateTime:          1566818724722,
		WorkingType:         "CONTRACT_PRICE",
		PriceProtect:        false,
	}
	s.assertConditionalOrderEqual(e, res)
}

func (s *umConditionalOrderServiceTestSuite) assertConditionalOrderEqual(e, a *UMConditionalOrder) {
	r := s.r()
	r.Equal(e.NewClientStrategyId, a.NewClientStrategyId, "NewClientStrategyId")
	r.Equal(e.StrategyId, a.StrategyId, "StrategyId")
	r.Equal(e.StrategyStatus, a.StrategyStatus, "StrategyStatus")
	r.Equal(e.StrategyType, a.StrategyType, "StrategyType")
	r.Equal(e.OrigQty, a.OrigQty, "OrigQty")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.ReduceOnly, a.ReduceOnly, "ReduceOnly")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.ActivatePrice, a.ActivatePrice, "ActivatePrice")
	r.Equal(e.PriceRate, a.PriceRate, "PriceRate")
	r.Equal(e.BookTime, a.BookTime, "BookTime")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.PriceProtect, a.PriceProtect, "PriceProtect")
}
