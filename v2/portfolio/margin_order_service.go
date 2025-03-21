package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginOrderService service to place margin orders
type MarginOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	orderType               OrderType
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	stopPrice               *string
	newClientOrderID        *string
	newOrderRespType        *NewOrderRespType
	icebergQty              *string
	sideEffectType          *string
	timeInForce             *TimeInForceType
	selfTradePreventionMode *SelfTradePreventionMode
	autoRepayAtCancel       *bool
}

// Symbol set symbol
func (s *MarginOrderService) Symbol(symbol string) *MarginOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *MarginOrderService) Side(side SideType) *MarginOrderService {
	s.side = side
	return s
}

// Type set order type
func (s *MarginOrderService) Type(orderType OrderType) *MarginOrderService {
	s.orderType = orderType
	return s
}

// Quantity set quantity
func (s *MarginOrderService) Quantity(quantity string) *MarginOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quote order quantity
func (s *MarginOrderService) QuoteOrderQty(quoteOrderQty string) *MarginOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *MarginOrderService) Price(price string) *MarginOrderService {
	s.price = &price
	return s
}

// StopPrice set stop price
func (s *MarginOrderService) StopPrice(stopPrice string) *MarginOrderService {
	s.stopPrice = &stopPrice
	return s
}

// NewClientOrderID set custom order id
func (s *MarginOrderService) NewClientOrderID(newClientOrderID string) *MarginOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderRespType set response type
func (s *MarginOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *MarginOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// IcebergQty set iceberg quantity
func (s *MarginOrderService) IcebergQty(icebergQty string) *MarginOrderService {
	s.icebergQty = &icebergQty
	return s
}

// SideEffectType set side effect type
func (s *MarginOrderService) SideEffectType(sideEffectType string) *MarginOrderService {
	s.sideEffectType = &sideEffectType
	return s
}

// TimeInForce set time in force
func (s *MarginOrderService) TimeInForce(timeInForce TimeInForceType) *MarginOrderService {
	s.timeInForce = &timeInForce
	return s
}

// SelfTradePreventionMode set self trade prevention mode
func (s *MarginOrderService) SelfTradePreventionMode(mode SelfTradePreventionMode) *MarginOrderService {
	s.selfTradePreventionMode = &mode
	return s
}

// AutoRepayAtCancel set auto repay at cancel
func (s *MarginOrderService) AutoRepayAtCancel(autoRepay bool) *MarginOrderService {
	s.autoRepayAtCancel = &autoRepay
	return s
}

// Do send request
func (s *MarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *MarginOrder, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/margin/order",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)
	r.setParam("side", s.side)
	r.setParam("type", s.orderType)

	if s.quantity != nil {
		r.setParam("quantity", *s.quantity)
	}
	if s.quoteOrderQty != nil {
		r.setParam("quoteOrderQty", *s.quoteOrderQty)
	}
	if s.price != nil {
		r.setParam("price", *s.price)
	}
	if s.stopPrice != nil {
		r.setParam("stopPrice", *s.stopPrice)
	}
	if s.newClientOrderID != nil {
		r.setParam("newClientOrderId", *s.newClientOrderID)
	}
	if s.newOrderRespType != nil {
		r.setParam("newOrderRespType", *s.newOrderRespType)
	}
	if s.icebergQty != nil {
		r.setParam("icebergQty", *s.icebergQty)
	}
	if s.sideEffectType != nil {
		r.setParam("sideEffectType", *s.sideEffectType)
	}
	if s.timeInForce != nil {
		r.setParam("timeInForce", *s.timeInForce)
	}
	if s.selfTradePreventionMode != nil {
		r.setParam("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.autoRepayAtCancel != nil {
		r.setParam("autoRepayAtCancel", *s.autoRepayAtCancel)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Fill define fill info
type Fill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

// MarginOrder define margin order info
type MarginOrder struct {
	Symbol                  string          `json:"symbol"`
	OrderID                 int64           `json:"orderId"`
	ClientOrderID           string          `json:"clientOrderId"`
	TransactTime            int64           `json:"time"`
	UpdateTime              int64           `json:"updateTime"`
	Price                   string          `json:"price"`
	OrigQty                 string          `json:"origQty"`
	ExecutedQty             string          `json:"executedQty"`
	CummulativeQuoteQty     string          `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             TimeInForceType `json:"timeInForce"`
	Type                    OrderType       `json:"type"`
	Side                    SideType        `json:"side"`
	MarginBuyBorrowAmount   string          `json:"marginBuyBorrowAmount,omitempty"`
	MarginBuyBorrowAsset    string          `json:"marginBuyBorrowAsset,omitempty"`
	IcebergQty              string          `json:"icebergQty"`
	IsWorking               bool            `json:"isWorking"`
	StopPrice               string          `json:"stopPrice"`
	AccountID               int64           `json:"accountId"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	PreventedMatchID        *string         `json:"preventedMatchId"`
	PreventedQuantity       *string         `json:"preventedQuantity"`
	Fills                   []*Fill         `json:"fills"`
}
