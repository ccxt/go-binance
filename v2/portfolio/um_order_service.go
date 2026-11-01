package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMOrderService service to place UM orders
type UMOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                *string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	newOrderRespType        *NewOrderRespType
	priceMatch              *PriceMatchType
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            *int64
}

// Symbol set symbol
func (s *UMOrderService) Symbol(symbol string) *UMOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *UMOrderService) Side(side SideType) *UMOrderService {
	s.side = side
	return s
}

// PositionSide set position side
func (s *UMOrderService) PositionSide(positionSide PositionSideType) *UMOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set order type
func (s *UMOrderService) Type(orderType OrderType) *UMOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set time in force
func (s *UMOrderService) TimeInForce(timeInForce TimeInForceType) *UMOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *UMOrderService) Quantity(quantity string) *UMOrderService {
	s.quantity = &quantity
	return s
}

// ReduceOnly set reduce only
func (s *UMOrderService) ReduceOnly(reduceOnly bool) *UMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *UMOrderService) Price(price string) *UMOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set custom order id
func (s *UMOrderService) NewClientOrderID(newClientOrderID string) *UMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderRespType set response type
func (s *UMOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *UMOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// PriceMatch set price match type
func (s *UMOrderService) PriceMatch(priceMatch PriceMatchType) *UMOrderService {
	s.priceMatch = &priceMatch
	return s
}

// SelfTradePreventionMode set self trade prevention mode
func (s *UMOrderService) SelfTradePreventionMode(mode SelfTradePreventionMode) *UMOrderService {
	s.selfTradePreventionMode = &mode
	return s
}

// GoodTillDate set good till date timestamp
func (s *UMOrderService) GoodTillDate(timestamp int64) *UMOrderService {
	s.goodTillDate = &timestamp
	return s
}

// Do send request
func (s *UMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UMOrder, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/order",
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
	if s.priceMatch != nil {
		r.setParam("priceMatch", *s.priceMatch)
	}
	if s.selfTradePreventionMode != nil {
		r.setParam("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.goodTillDate != nil {
		r.setParam("goodTillDate", *s.goodTillDate)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMOrder define UM order info
type UMOrder struct {
	ClientOrderID           string                  `json:"clientOrderId"`
	CumQty                  string                  `json:"cumQty"`
	CumQuote                string                  `json:"cumQuote"`
	ExecutedQty             string                  `json:"executedQty"`
	OrderID                 int64                   `json:"orderId"`
	AvgPrice                string                  `json:"avgPrice"`
	OrigQty                 string                  `json:"origQty"`
	Price                   string                  `json:"price"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	Side                    SideType                `json:"side"`
	PositionSide            PositionSideType        `json:"positionSide"`
	Status                  string                  `json:"status"`
	Symbol                  string                  `json:"symbol"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	GoodTillDate            int64                   `json:"goodTillDate"`
	UpdateTime              int64                   `json:"updateTime"`
	PriceMatch              PriceMatchType          `json:"priceMatch"`
}
