package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMConditionalOrderService service to place UM conditional orders
type UMConditionalOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	strategyType            string
	timeInForce             *TimeInForceType
	quantity                *string
	reduceOnly              *bool
	price                   *string
	workingType             *string
	priceProtect            *bool
	newClientStrategyId     *string
	stopPrice               *string
	activationPrice         *string
	callbackRate            *string
	priceMatch              *PriceMatchType
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            *int64
}

// Symbol set symbol
func (s *UMConditionalOrderService) Symbol(symbol string) *UMConditionalOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *UMConditionalOrderService) Side(side SideType) *UMConditionalOrderService {
	s.side = side
	return s
}

// PositionSide set position side
func (s *UMConditionalOrderService) PositionSide(positionSide PositionSideType) *UMConditionalOrderService {
	s.positionSide = &positionSide
	return s
}

// StrategyType set strategy type
func (s *UMConditionalOrderService) StrategyType(strategyType string) *UMConditionalOrderService {
	s.strategyType = strategyType
	return s
}

// TimeInForce set time in force
func (s *UMConditionalOrderService) TimeInForce(timeInForce TimeInForceType) *UMConditionalOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *UMConditionalOrderService) Quantity(quantity string) *UMConditionalOrderService {
	s.quantity = &quantity
	return s
}

// Price set price
func (s *UMConditionalOrderService) Price(price string) *UMConditionalOrderService {
	s.price = &price
	return s
}

// StopPrice set stop price
func (s *UMConditionalOrderService) StopPrice(stopPrice string) *UMConditionalOrderService {
	s.stopPrice = &stopPrice
	return s
}

// ActivationPrice set activation price
func (s *UMConditionalOrderService) ActivationPrice(price string) *UMConditionalOrderService {
	s.activationPrice = &price
	return s
}

// CallbackRate set callback rate
func (s *UMConditionalOrderService) CallbackRate(rate string) *UMConditionalOrderService {
	s.callbackRate = &rate
	return s
}

// WorkingType set working type
func (s *UMConditionalOrderService) WorkingType(workingType string) *UMConditionalOrderService {
	s.workingType = &workingType
	return s
}

// PriceProtect set price protect
func (s *UMConditionalOrderService) PriceProtect(protect bool) *UMConditionalOrderService {
	s.priceProtect = &protect
	return s
}

// Do send request
func (s *UMConditionalOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UMConditionalOrder, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/conditional/order",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("strategyType", s.strategyType)

	if s.positionSide != nil {
		r.setParam("positionSide", *s.positionSide)
	}
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setParam("quantity", *s.quantity)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.stopPrice != nil {
		r.setParam("stopPrice", *s.stopPrice)
	}
	if s.activationPrice != nil {
		r.setParam("activationPrice", *s.activationPrice)
	}
	if s.callbackRate != nil {
		r.setParam("callbackRate", *s.callbackRate)
	}
	if s.workingType != nil {
		r.setParam("workingType", *s.workingType)
	}
	if s.priceProtect != nil {
		r.setParam("priceProtect", *s.priceProtect)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMConditionalOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMConditionalOrder define conditional order info
type UMConditionalOrder struct {
	NewClientStrategyId string           `json:"newClientStrategyId"`
	StrategyId          int64            `json:"strategyId"`
	StrategyStatus      string           `json:"strategyStatus"`
	StrategyType        string           `json:"strategyType"`
	OrigQty             string           `json:"origQty"`
	Price               string           `json:"price"`
	ReduceOnly          bool             `json:"reduceOnly"`
	Side                SideType         `json:"side"`
	PositionSide        PositionSideType `json:"positionSide"`
	StopPrice           string           `json:"stopPrice"`
	Symbol              string           `json:"symbol"`
	TimeInForce         TimeInForceType  `json:"timeInForce"`
	ActivatePrice       string           `json:"activatePrice"`
	PriceRate           string           `json:"priceRate"`
	BookTime            int64            `json:"bookTime"`
	UpdateTime          int64            `json:"updateTime"`
	WorkingType         string           `json:"workingType"`
	PriceProtect        bool             `json:"priceProtect"`
	PriceMatch          string           `json:"priceMatch"`
}
