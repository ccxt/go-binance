package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMConditionalOrderService service to place CM conditional orders
type CMConditionalOrderService struct {
	c                   *Client
	symbol              string
	side                SideType
	positionSide        *PositionSideType
	strategyType        string
	timeInForce         *TimeInForceType
	quantity            *string
	reduceOnly          *bool
	price               *string
	workingType         *string
	priceProtect        *bool
	newClientStrategyId *string
	stopPrice           *string
	activationPrice     *string
	callbackRate        *string
}

// Symbol set symbol
func (s *CMConditionalOrderService) Symbol(symbol string) *CMConditionalOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CMConditionalOrderService) Side(side SideType) *CMConditionalOrderService {
	s.side = side
	return s
}

// PositionSide set position side
func (s *CMConditionalOrderService) PositionSide(positionSide PositionSideType) *CMConditionalOrderService {
	s.positionSide = &positionSide
	return s
}

// StrategyType set strategy type
func (s *CMConditionalOrderService) StrategyType(strategyType string) *CMConditionalOrderService {
	s.strategyType = strategyType
	return s
}

// TimeInForce set time in force
func (s *CMConditionalOrderService) TimeInForce(timeInForce TimeInForceType) *CMConditionalOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CMConditionalOrderService) Quantity(quantity string) *CMConditionalOrderService {
	s.quantity = &quantity
	return s
}

// Price set price
func (s *CMConditionalOrderService) Price(price string) *CMConditionalOrderService {
	s.price = &price
	return s
}

// StopPrice set stop price
func (s *CMConditionalOrderService) StopPrice(stopPrice string) *CMConditionalOrderService {
	s.stopPrice = &stopPrice
	return s
}

// ActivationPrice set activation price
func (s *CMConditionalOrderService) ActivationPrice(price string) *CMConditionalOrderService {
	s.activationPrice = &price
	return s
}

// CallbackRate set callback rate
func (s *CMConditionalOrderService) CallbackRate(rate string) *CMConditionalOrderService {
	s.callbackRate = &rate
	return s
}

// WorkingType set working type
func (s *CMConditionalOrderService) WorkingType(workingType string) *CMConditionalOrderService {
	s.workingType = &workingType
	return s
}

// PriceProtect set price protect
func (s *CMConditionalOrderService) PriceProtect(protect bool) *CMConditionalOrderService {
	s.priceProtect = &protect
	return s
}

// NewClientStrategyID set client strategy ID
func (s *CMConditionalOrderService) NewClientStrategyID(id string) *CMConditionalOrderService {
	s.newClientStrategyId = &id
	return s
}

// Do send request
func (s *CMConditionalOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CMConditionalOrder, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/cm/conditional/order",
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
	if s.newClientStrategyId != nil {
		r.setParam("newClientStrategyId", *s.newClientStrategyId)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CMConditionalOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMConditionalOrder define conditional order info
type CMConditionalOrder struct {
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
	Pair                string           `json:"pair"`
	TimeInForce         TimeInForceType  `json:"timeInForce"`
	ActivatePrice       string           `json:"activatePrice"`
	PriceRate           string           `json:"priceRate"`
	BookTime            int64            `json:"bookTime"`
	UpdateTime          int64            `json:"updateTime"`
	WorkingType         string           `json:"workingType"`
	PriceProtect        bool             `json:"priceProtect"`
}
