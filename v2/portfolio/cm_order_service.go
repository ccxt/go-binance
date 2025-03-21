package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMOrderService service to place CM orders
type CMOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         *string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	newOrderRespType *NewOrderRespType
}

// Symbol set symbol
func (s *CMOrderService) Symbol(symbol string) *CMOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CMOrderService) Side(side SideType) *CMOrderService {
	s.side = side
	return s
}

// PositionSide set position side
func (s *CMOrderService) PositionSide(positionSide PositionSideType) *CMOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set order type
func (s *CMOrderService) Type(orderType OrderType) *CMOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set time in force
func (s *CMOrderService) TimeInForce(timeInForce TimeInForceType) *CMOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CMOrderService) Quantity(quantity string) *CMOrderService {
	s.quantity = &quantity
	return s
}

// ReduceOnly set reduce only
func (s *CMOrderService) ReduceOnly(reduceOnly bool) *CMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CMOrderService) Price(price string) *CMOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set custom order id
func (s *CMOrderService) NewClientOrderID(newClientOrderID string) *CMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderRespType set response type
func (s *CMOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CMOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// Do send request
func (s *CMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CMOrder, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/cm/order",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("type", s.orderType)

	if s.positionSide != nil {
		r.setParam("positionSide", *s.positionSide)
	}
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		r.setParam("quantity", *s.quantity)
	}
	if s.reduceOnly != nil {
		r.setParam("reduceOnly", *s.reduceOnly)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.newClientOrderID != nil {
		r.setParam("newClientOrderId", *s.newClientOrderID)
	}
	if s.newOrderRespType != nil {
		r.setParam("newOrderRespType", *s.newOrderRespType)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CMOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMOrder define CM order info
type CMOrder struct {
	ClientOrderID string           `json:"clientOrderId"`
	CumQty        string           `json:"cumQty"`
	CumBase       string           `json:"cumBase"`
	ExecutedQty   string           `json:"executedQty"`
	OrderID       int64            `json:"orderId"`
	AvgPrice      string           `json:"avgPrice"`
	OrigQty       string           `json:"origQty"`
	Price         string           `json:"price"`
	ReduceOnly    bool             `json:"reduceOnly"`
	Side          SideType         `json:"side"`
	PositionSide  PositionSideType `json:"positionSide"`
	Status        string           `json:"status"`
	Symbol        string           `json:"symbol"`
	Pair          string           `json:"pair"`
	TimeInForce   TimeInForceType  `json:"timeInForce"`
	Type          OrderType        `json:"type"`
	UpdateTime    int64            `json:"updateTime"`
}
