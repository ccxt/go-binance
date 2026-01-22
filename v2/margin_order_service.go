package binance

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// CreateMarginOrderService create order
type CreateMarginOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	orderType               OrderType
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	stopPrice               *string
	newClientOrderID        *string
	icebergQuantity         *string
	newOrderRespType        *NewOrderRespType
	sideEffectType          *SideEffectType
	timeInForce             *TimeInForceType
	isIsolated              *bool
	selfTradePreventionMode *SelfTradePreventionMode
	autoRepayAtCancel       *bool
}

// Symbol set symbol
func (s *CreateMarginOrderService) Symbol(symbol string) *CreateMarginOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated sets the order to isolated margin
func (s *CreateMarginOrderService) IsIsolated(isIsolated bool) *CreateMarginOrderService {
	s.isIsolated = &isIsolated
	return s
}

// Side set side
func (s *CreateMarginOrderService) Side(side SideType) *CreateMarginOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateMarginOrderService) Type(orderType OrderType) *CreateMarginOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateMarginOrderService) TimeInForce(timeInForce TimeInForceType) *CreateMarginOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateMarginOrderService) Quantity(quantity string) *CreateMarginOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *CreateMarginOrderService) QuoteOrderQty(quoteOrderQty string) *CreateMarginOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *CreateMarginOrderService) Price(price string) *CreateMarginOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateMarginOrderService) NewClientOrderID(newClientOrderID string) *CreateMarginOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateMarginOrderService) StopPrice(stopPrice string) *CreateMarginOrderService {
	s.stopPrice = &stopPrice
	return s
}

// IcebergQuantity set icebergQuantity
func (s *CreateMarginOrderService) IcebergQuantity(icebergQuantity string) *CreateMarginOrderService {
	s.icebergQuantity = &icebergQuantity
	return s
}

// NewOrderRespType set icebergQuantity
func (s *CreateMarginOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateMarginOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *CreateMarginOrderService) SideEffectType(sideEffectType SideEffectType) *CreateMarginOrderService {
	s.sideEffectType = &sideEffectType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateMarginOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateMarginOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// AutoRepayAtCancel set autoRepayAtCancel
func (s *CreateMarginOrderService) AutoRepayAtCancel(autoRepayAtCancel bool) *CreateMarginOrderService {
	s.autoRepayAtCancel = &autoRepayAtCancel
	return s
}

// Do send request
func (s *CreateMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/order",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = *s.quoteOrderQty
	}
	if s.isIsolated != nil {
		if *s.isIsolated {
			m["isIsolated"] = "TRUE"
		} else {
			m["isIsolated"] = "FALSE"
		}
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	} else {
		m["newClientOrderId"] = common.GenerateSpotId()
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.icebergQuantity != nil {
		m["icebergQty"] = *s.icebergQuantity
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.autoRepayAtCancel != nil {
		if *s.autoRepayAtCancel {
			m["autoRepayAtCancel"] = "TRUE"
		} else {
			m["autoRepayAtCancel"] = "FALSE"
		}
	}
	r.setFormParams(m)
	res = new(CreateOrderResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelMarginOrderService cancel an order
type CancelMarginOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	newClientOrderID  *string
	isIsolated        *bool
}

// Symbol set symbol
func (s *CancelMarginOrderService) Symbol(symbol string) *CancelMarginOrderService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *CancelMarginOrderService) IsIsolated(isIsolated bool) *CancelMarginOrderService {
	s.isIsolated = &isIsolated
	return s
}

// OrderID set orderID
func (s *CancelMarginOrderService) OrderID(orderID int64) *CancelMarginOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelMarginOrderService) OrigClientOrderID(origClientOrderID string) *CancelMarginOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CancelMarginOrderService) NewClientOrderID(newClientOrderID string) *CancelMarginOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// Do send request
func (s *CancelMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelMarginOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setParam("newClientOrderId", *s.newClientOrderID)
	}
	if s.isIsolated != nil {
		if *s.isIsolated {
			r.setParam("isIsolated", "TRUE")
		} else {
			r.setParam("isIsolated", "FALSE")
		}
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelMarginOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelAllMarginOrdersService cancel an order
type CancelAllMarginOrdersService struct {
	c          *Client
	symbol     string
	isIsolated *bool
}

// Symbol set symbol
func (s *CancelAllMarginOrdersService) Symbol(symbol string) *CancelAllMarginOrdersService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *CancelAllMarginOrdersService) IsIsolated(isIsolated bool) *CancelAllMarginOrdersService {
	s.isIsolated = &isIsolated
	return s
}

// Do send request
func (s *CancelAllMarginOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*CancelAllMarginOrdersResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/margin/openOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.isIsolated != nil {
		if *s.isIsolated {
			r.setParam("isIsolated", "TRUE")
		} else {
			r.setParam("isIsolated", "FALSE")
		}
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CancelAllMarginOrdersResponse{}, err
	}
	res = make([]*CancelAllMarginOrdersResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CancelAllMarginOrdersResponse{}, err
	}
	return res, nil
}

// GetMarginOrderService get an order
type GetMarginOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	isIsolated        bool
}

// IsIsolated set isIsolated
func (s *GetMarginOrderService) IsIsolated(isIsolated bool) *GetMarginOrderService {
	s.isIsolated = isIsolated
	return s
}

// Symbol set symbol
func (s *GetMarginOrderService) Symbol(symbol string) *GetMarginOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetMarginOrderService) OrderID(orderID int64) *GetMarginOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetMarginOrderService) OrigClientOrderID(origClientOrderID string) *GetMarginOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginOpenOrdersService list margin open orders
type ListMarginOpenOrdersService struct {
	c          *Client
	symbol     string
	isIsolated bool
}

// Symbol set symbol
func (s *ListMarginOpenOrdersService) Symbol(symbol string) *ListMarginOpenOrdersService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *ListMarginOpenOrdersService) IsIsolated(isIsolated bool) *ListMarginOpenOrdersService {
	s.isIsolated = isIsolated
	return s
}

// Do send request
func (s *ListMarginOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// ListMarginOrdersService all account orders; active, canceled, or filled
type ListMarginOrdersService struct {
	c          *Client
	symbol     string
	orderID    *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	isIsolated bool
}

// Symbol set symbol
func (s *ListMarginOrdersService) Symbol(symbol string) *ListMarginOrdersService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *ListMarginOrdersService) IsIsolated(isIsolated bool) *ListMarginOrdersService {
	s.isIsolated = isIsolated
	return s
}

// OrderID set orderID
func (s *ListMarginOrdersService) OrderID(orderID int64) *ListMarginOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListMarginOrdersService) StartTime(startTime int64) *ListMarginOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginOrdersService) EndTime(endTime int64) *ListMarginOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginOrdersService) Limit(limit int) *ListMarginOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListMarginOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/allOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// CancelMarginOrderResponse define response of canceling order
type CancelMarginOrderResponse struct {
	Symbol                   string          `json:"symbol"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
	OrderID                  string          `json:"orderId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactTime             int64           `json:"transactTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
}

// CancelAllMarginOrdersResponse define response of canceling order
type CancelAllMarginOrdersResponse struct {
	Symbol                   string                         `json:"symbol"`
	OrigClientOrderID        string                         `json:"origClientOrderId"`
	OrderID                  int64                         `json:"orderId"`
	ClientOrderID            string                         `json:"clientOrderId"`
	TransactTime             int64                          `json:"transactTime"`
	Price                    string                         `json:"price"`
	OrigQuantity             string                         `json:"origQty"`
	ExecutedQuantity         string                         `json:"executedQty"`
	CummulativeQuoteQuantity string                         `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType                `json:"status"`
	TimeInForce              TimeInForceType                `json:"timeInForce"`
	Type                     OrderType                      `json:"type"`
	Side                     SideType                       `json:"side"`
	SelfTradePreventionMode  string                         `json:"selfTradePreventionMode"`
	OrderListID              int64                          `json:"orderListId"`
	ContingencyType          string                         `json:"contingencyType"`
	ListStatusType           string                         `json:"listStatusType"`
	ListOrderStatus          string                         `json:"listOrderStatus"`
	ListClientOrderID        string                         `json:"listClientOrderId"`
	TransactionTime          int64                          `json:"transactionTime"`
	IsIsolated               bool                           `json:"isIsolated"`
	Orders                   []*MarginOCOOrder              `json:"orders"`
	OrderReports             []*CancelAllMarginOrdersReport `json:"orderReports"`
}

// CancelAllMarginOrdersReport may be returned in an array of MarginOCOOrderReport in a CreateMarginOCOResponse
type CancelAllMarginOrdersReport struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
	StopPrice                string          `json:"stopPrice"`
	IcebergQty               string          `json:"icebergQty"`
}

// CreateMarginOCOService create a new OCO for a margin account
type CreateMarginOCOService struct {
	c                    *Client
	symbol               string
	isIsolated           *bool
	listClientOrderID    *string
	side                 SideType
	quantity             *string
	limitClientOrderID   *string
	price                *string
	limitIcebergQty      *string
	stopClientOrderID    *string
	stopPrice            *string
	stopLimitPrice       *string
	stopIcebergQty       *string
	stopLimitTimeInForce *TimeInForceType
	newOrderRespType     *NewOrderRespType
	sideEffectType       *SideEffectType
}

// Symbol set symbol
func (s *CreateMarginOCOService) Symbol(symbol string) *CreateMarginOCOService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *CreateMarginOCOService) IsIsolated(isIsolated bool) *CreateMarginOCOService {
	s.isIsolated = &isIsolated
	return s
}

// Side set side
func (s *CreateMarginOCOService) Side(side SideType) *CreateMarginOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *CreateMarginOCOService) Quantity(quantity string) *CreateMarginOCOService {
	s.quantity = &quantity
	return s
}

// ListClientOrderID set listClientOrderID
func (s *CreateMarginOCOService) ListClientOrderID(listClientOrderID string) *CreateMarginOCOService {
	s.listClientOrderID = &listClientOrderID
	return s
}

// LimitClientOrderID set limitClientOrderID
func (s *CreateMarginOCOService) LimitClientOrderID(limitClientOrderID string) *CreateMarginOCOService {
	s.limitClientOrderID = &limitClientOrderID
	return s
}

// Price set price
func (s *CreateMarginOCOService) Price(price string) *CreateMarginOCOService {
	s.price = &price
	return s
}

// LimitIcebergQuantity set limitIcebergQuantity
func (s *CreateMarginOCOService) LimitIcebergQuantity(limitIcebergQty string) *CreateMarginOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// StopClientOrderID set stopClientOrderID
func (s *CreateMarginOCOService) StopClientOrderID(stopClientOrderID string) *CreateMarginOCOService {
	s.stopClientOrderID = &stopClientOrderID
	return s
}

// StopPrice set stop price
func (s *CreateMarginOCOService) StopPrice(stopPrice string) *CreateMarginOCOService {
	s.stopPrice = &stopPrice
	return s
}

// StopLimitPrice set stop limit price
func (s *CreateMarginOCOService) StopLimitPrice(stopLimitPrice string) *CreateMarginOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stop limit price
func (s *CreateMarginOCOService) StopIcebergQty(stopIcebergQty string) *CreateMarginOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *CreateMarginOCOService) StopLimitTimeInForce(stopLimitTimeInForce TimeInForceType) *CreateMarginOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set icebergQuantity
func (s *CreateMarginOCOService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateMarginOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SideEffectType set sideEffectType
func (s *CreateMarginOCOService) SideEffectType(sideEffectType SideEffectType) *CreateMarginOCOService {
	s.sideEffectType = &sideEffectType
	return s
}

func (s *CreateMarginOCOService) createOrder(ctx context.Context, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/order/oco",
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":    s.symbol,
		"side":      s.side,
		"quantity":  *s.quantity,
		"price":     *s.price,
		"stopPrice": *s.stopPrice,
	}

	if s.isIsolated != nil {
		if *s.isIsolated {
			m["isIsolated"] = "TRUE"
		} else {
			m["isIsolated"] = "FALSE"
		}
	}
	if s.listClientOrderID != nil {
		m["listClientOrderId"] = *s.listClientOrderID
	}
	if s.limitClientOrderID != nil {
		m["limitClientOrderId"] = *s.limitClientOrderID
	}
	if s.limitIcebergQty != nil {
		m["limitIcebergQty"] = *s.limitIcebergQty
	}
	if s.stopClientOrderID != nil {
		m["stopClientOrderId"] = *s.stopClientOrderID
	}
	if s.stopLimitPrice != nil {
		m["stopLimitPrice"] = *s.stopLimitPrice
	}
	if s.stopIcebergQty != nil {
		m["stopIcebergQty"] = *s.stopIcebergQty
	}
	if s.stopLimitTimeInForce != nil {
		m["stopLimitTimeInForce"] = *s.stopLimitTimeInForce
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateMarginOCOService) Do(ctx context.Context, opts ...RequestOption) (res *CreateMarginOCOResponse, err error) {
	data, err := s.createOrder(ctx, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateMarginOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateMarginOCOResponse define create order response
type CreateMarginOCOResponse struct {
	OrderListID           int64                   `json:"orderListId"`
	ContingencyType       string                  `json:"contingencyType"`
	ListStatusType        string                  `json:"listStatusType"`
	ListOrderStatus       string                  `json:"listOrderStatus"`
	ListClientOrderID     string                  `json:"listClientOrderId"`
	TransactionTime       int64                   `json:"transactionTime"`
	Symbol                string                  `json:"symbol"`
	MarginBuyBorrowAmount string                  `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string                  `json:"marginBuyBorrowAsset"`
	IsIsolated            bool                    `json:"isIsolated"`
	Orders                []*MarginOCOOrder       `json:"orders"`
	OrderReports          []*MarginOCOOrderReport `json:"orderReports"`
}

// MarginOCOOrder may be returned in an array of MarginOCOOrder in a CreateMarginOCOResponse
type MarginOCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// MarginOCOOrderReport may be returned in an array of MarginOCOOrderReport in a CreateMarginOCOResponse
type MarginOCOOrderReport struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactionTime          int64           `json:"transactionTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
	StopPrice                string          `json:"stopPrice"`
}

// CancelMarginOCOService cancel an entire Order List for a margin account
type CancelMarginOCOService struct {
	c                 *Client
	symbol            string
	isIsolated        *bool
	listClientOrderID string
	orderListID       int64
	newClientOrderID  string
}

// Symbol set symbol
func (s *CancelMarginOCOService) Symbol(symbol string) *CancelMarginOCOService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *CancelMarginOCOService) IsIsolated(isIsolated bool) *CancelMarginOCOService {
	s.isIsolated = &isIsolated
	return s
}

// ListClientOrderID sets listClientOrderId
func (s *CancelMarginOCOService) ListClientOrderID(listClientOrderID string) *CancelMarginOCOService {
	s.listClientOrderID = listClientOrderID
	return s
}

// OrderListID sets orderListId
func (s *CancelMarginOCOService) OrderListID(orderListID int64) *CancelMarginOCOService {
	s.orderListID = orderListID
	return s
}

// NewClientOrderID sets newClientOrderId
func (s *CancelMarginOCOService) NewClientOrderID(newClientOrderID string) *CancelMarginOCOService {
	s.newClientOrderID = newClientOrderID
	return s
}

// Do send request
func (s *CancelMarginOCOService) Do(ctx context.Context, opts ...RequestOption) (res *CancelMarginOCOResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/margin/orderList",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.listClientOrderID != "" {
		r.setParam("listClientOrderId", s.listClientOrderID)
	}
	if s.isIsolated != nil {
		r.setParam("isIsolated", *s.isIsolated)
	}
	if s.orderListID != 0 {
		r.setParam("orderListId", s.orderListID)
	}
	if s.newClientOrderID != "" {
		r.setParam("newClientOrderId", s.newClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelMarginOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelMarginOCOResponse define create cancelled oco response.
type CancelMarginOCOResponse struct {
	OrderListID       int64                   `json:"orderListId"`
	ContingencyType   string                  `json:"contingencyType"`
	ListStatusType    string                  `json:"listStatusType"`
	ListOrderStatus   string                  `json:"listOrderStatus"`
	ListClientOrderID string                  `json:"listClientOrderId"`
	TransactionTime   int64                   `json:"transactionTime"`
	Symbol            string                  `json:"symbol"`
	IsIsolated        bool                    `json:"isIsolated"`
	Orders            []*MarginOCOOrder       `json:"orders"`
	OrderReports      []*MarginOCOOrderReport `json:"orderReports"`
}
