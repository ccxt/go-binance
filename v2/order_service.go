package binance

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	newOrderRespType        *NewOrderRespType
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	newClientOrderID        *string
	stopPrice               *string
	trailingDelta           *string
	icebergQuantity         *string
	selfTradePreventionMode *SelfTradePreventionMode
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce TimeInForceType) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity string) *CreateOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *CreateOrderService) QuoteOrderQty(quoteOrderQty string) *CreateOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// Price set price
func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateOrderService) NewClientOrderID(newClientOrderID string) *CreateOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateOrderService) StopPrice(stopPrice string) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *CreateOrderService) TrailingDelta(trailingDelta string) *CreateOrderService {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQuantity set icebergQuantity
func (s *CreateOrderService) IcebergQuantity(icebergQuantity string) *CreateOrderService {
	s.icebergQuantity = &icebergQuantity
	return s
}

// NewOrderRespType set icebergQuantity
func (s *CreateOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
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
	if s.trailingDelta != nil {
		m["trailingDelta"] = *s.trailingDelta
	}
	if s.icebergQuantity != nil {
		m["icebergQty"] = *s.icebergQuantity
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/api/v3/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Test send test api to check if the request is valid
func (s *CreateOrderService) Test(ctx context.Context, opts ...RequestOption) (err error) {
	_, err = s.createOrder(ctx, "/api/v3/order/test", opts...)
	return err
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	Symbol                   string `json:"symbol"`
	OrderID                  int64  `json:"orderId"`
	ClientOrderID            string `json:"clientOrderId"`
	TransactTime             int64  `json:"transactTime"`
	Price                    string `json:"price"`
	OrigQuantity             string `json:"origQty"`
	OrigQuoteOrderQuantity   string `json:"origQuoteOrderQty"`
	ExecutedQuantity         string `json:"executedQty"`
	CummulativeQuoteQuantity string `json:"cummulativeQuoteQty"`
	IsIsolated               bool   `json:"isIsolated"` // for isolated margin

	Status      OrderStatusType `json:"status"`
	TimeInForce TimeInForceType `json:"timeInForce"`
	Type        OrderType       `json:"type"`
	Side        SideType        `json:"side"`

	// for order response is set to FULL
	Fills                 []*Fill `json:"fills"`
	MarginBuyBorrowAmount string  `json:"marginBuyBorrowAmount"` // for margin
	MarginBuyBorrowAsset  string  `json:"marginBuyBorrowAsset"`

	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
}

// Fill may be returned in an array of fills in a CreateOrderResponse.
type Fill struct {
	TradeID         int64  `json:"tradeId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

// CreateOCOService create order
type CreateOCOService struct {
	c                    *Client
	symbol               string
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
}

// Symbol set symbol
func (s *CreateOCOService) Symbol(symbol string) *CreateOCOService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOCOService) Side(side SideType) *CreateOCOService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *CreateOCOService) Quantity(quantity string) *CreateOCOService {
	s.quantity = &quantity
	return s
}

// ListClientOrderID set listClientOrderID
func (s *CreateOCOService) ListClientOrderID(listClientOrderID string) *CreateOCOService {
	s.listClientOrderID = &listClientOrderID
	return s
}

// LimitClientOrderID set limitClientOrderID
func (s *CreateOCOService) LimitClientOrderID(limitClientOrderID string) *CreateOCOService {
	s.limitClientOrderID = &limitClientOrderID
	return s
}

// Price set price
func (s *CreateOCOService) Price(price string) *CreateOCOService {
	s.price = &price
	return s
}

// LimitIcebergQuantity set limitIcebergQuantity
func (s *CreateOCOService) LimitIcebergQuantity(limitIcebergQty string) *CreateOCOService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// StopClientOrderID set stopClientOrderID
func (s *CreateOCOService) StopClientOrderID(stopClientOrderID string) *CreateOCOService {
	s.stopClientOrderID = &stopClientOrderID
	return s
}

// StopPrice set stop price
func (s *CreateOCOService) StopPrice(stopPrice string) *CreateOCOService {
	s.stopPrice = &stopPrice
	return s
}

// StopLimitPrice set stop limit price
func (s *CreateOCOService) StopLimitPrice(stopLimitPrice string) *CreateOCOService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopIcebergQty set stop limit price
func (s *CreateOCOService) StopIcebergQty(stopIcebergQty string) *CreateOCOService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *CreateOCOService) StopLimitTimeInForce(stopLimitTimeInForce TimeInForceType) *CreateOCOService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// NewOrderRespType set icebergQuantity
func (s *CreateOCOService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateOCOService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *CreateOCOService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":    s.symbol,
		"side":      s.side,
		"quantity":  *s.quantity,
		"price":     *s.price,
		"stopPrice": *s.stopPrice,
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
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateOCOService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOCOResponse, err error) {
	data, err := s.createOrder(ctx, "/api/v3/order/oco", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOCOResponse define create order response
type CreateOCOResponse struct {
	OrderListID       int64             `json:"orderListId"`
	ContingencyType   string            `json:"contingencyType"`
	ListStatusType    string            `json:"listStatusType"`
	ListOrderStatus   string            `json:"listOrderStatus"`
	ListClientOrderID string            `json:"listClientOrderId"`
	TransactionTime   int64             `json:"transactionTime"`
	Symbol            string            `json:"symbol"`
	Orders            []*OCOOrder       `json:"orders"`
	OrderReports      []*OCOOrderReport `json:"orderReports"`
}

// OCOOrder may be returned in an array of OCOOrder in a CreateOCOResponse.
type OCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// OCOOrderReport may be returned in an array of OCOOrderReport in a CreateOCOResponse.
type OCOOrderReport struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
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
	IcebergQuantity          string          `json:"icebergQty"`
}

// ListOpenOcoService list opened oco
type ListOpenOcoService struct {
	c *Client
}

// oco define oco info
type Oco struct {
	Symbol            string   `json:"symbol"`
	OrderListId       int64    `json:"orderListId"`
	ContingencyType   string   `json:"contingencyType"`
	ListStatusType    string   `json:"listStatusType"`
	ListOrderStatus   string   `json:"listOrderStatus"`
	ListClientOrderID string   `json:"listClientOrderId"`
	TransactionTime   int64    `json:"transactionTime"`
	Orders            []*Order `json:"orders"`
}

// Do send request
func (s *ListOpenOcoService) Do(ctx context.Context, opts ...RequestOption) (res []*Oco, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/openOrderList",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Oco{}, err
	}
	res = make([]*Oco, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Oco{}, err
	}
	return res, nil
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
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

// GetOrderService get an order
type GetOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetOrderService) OrderID(orderID int64) *GetOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetOrderService) OrigClientOrderID(origClientOrderID string) *GetOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
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

// Order define order info
type Order struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListId              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
	StopPrice                string          `json:"stopPrice"`
	IcebergQuantity          string          `json:"icebergQty"`
	Time                     int64           `json:"time"`
	UpdateTime               int64           `json:"updateTime"`
	IsWorking                bool            `json:"isWorking"`
	IsIsolated               bool            `json:"isIsolated"`
	OrigQuoteOrderQuantity   string          `json:"origQuoteOrderQty"`
}

// ListOrdersService all account orders; active, canceled, or filled
type ListOrdersService struct {
	c         *Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListOrdersService) Symbol(symbol string) *ListOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *ListOrdersService) OrderID(orderID int64) *ListOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListOrdersService) StartTime(startTime int64) *ListOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListOrdersService) EndTime(endTime int64) *ListOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListOrdersService) Limit(limit int) *ListOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v3/allOrders",
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

// CancelOrderService cancel an order
type CancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
	newClientOrderID  *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelOrderService) OrigClientOrderID(origClientOrderID string) *CancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CancelOrderService) NewClientOrderID(newClientOrderID string) *CancelOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	if s.newClientOrderID != nil {
		r.setFormParam("newClientOrderId", *s.newClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOCOService cancel all active orders on the list order.
type CancelOCOService struct {
	c                 *Client
	symbol            string
	listClientOrderID string
	orderListID       int64
	newClientOrderID  string
}

// Symbol set symbol
func (s *CancelOCOService) Symbol(symbol string) *CancelOCOService {
	s.symbol = symbol
	return s
}

// ListClientOrderID sets listClientOrderId
func (s *CancelOCOService) ListClientOrderID(listClientOrderID string) *CancelOCOService {
	s.listClientOrderID = listClientOrderID
	return s
}

// OrderListID sets orderListId
func (s *CancelOCOService) OrderListID(orderListID int64) *CancelOCOService {
	s.orderListID = orderListID
	return s
}

// NewClientOrderID sets newClientOrderId
func (s *CancelOCOService) NewClientOrderID(newClientOrderID string) *CancelOCOService {
	s.newClientOrderID = newClientOrderID
	return s
}

// Do send request
func (s *CancelOCOService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOCOResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/orderList",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.listClientOrderID != "" {
		r.setFormParam("listClientOrderId", s.listClientOrderID)
	}
	if s.orderListID != 0 {
		r.setFormParam("orderListId", s.orderListID)
	}
	if s.newClientOrderID != "" {
		r.setFormParam("newClientOrderId", s.newClientOrderID)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOpenOrdersService cancel all active orders on a symbol.
type CancelOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelOpenOrdersService) Symbol(symbol string) *CancelOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOpenOrdersResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/api/v3/openOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CancelOpenOrdersResponse{}, err
	}
	rawMessages := make([]*json.RawMessage, 0)
	err = json.Unmarshal(data, &rawMessages)
	if err != nil {
		return &CancelOpenOrdersResponse{}, err
	}
	cancelOpenOrdersResponse := new(CancelOpenOrdersResponse)
	for _, j := range rawMessages {
		o := new(CancelOrderResponse)
		if err := json.Unmarshal(*j, o); err != nil {
			return &CancelOpenOrdersResponse{}, err
		}
		// Non-OCO orders guaranteed to have order list ID of -1
		if o.OrderListID == -1 {
			cancelOpenOrdersResponse.Orders = append(cancelOpenOrdersResponse.Orders, o)
			continue
		}
		oco := new(CancelOCOResponse)
		if err := json.Unmarshal(*j, oco); err != nil {
			return &CancelOpenOrdersResponse{}, err
		}
		cancelOpenOrdersResponse.OCOOrders = append(cancelOpenOrdersResponse.OCOOrders, oco)
	}
	return cancelOpenOrdersResponse, nil
}

// CancelOpenOrdersResponse defines cancel open orders response.
type CancelOpenOrdersResponse struct {
	Orders    []*CancelOrderResponse
	OCOOrders []*CancelOCOResponse
}

// CancelOrderResponse may be returned included in a CancelOpenOrdersResponse.
type CancelOrderResponse struct {
	Symbol                   string                  `json:"symbol"`
	OrigClientOrderID        string                  `json:"origClientOrderId"`
	OrderID                  int64                   `json:"orderId"`
	OrderListID              int64                   `json:"orderListId"`
	ClientOrderID            string                  `json:"clientOrderId"`
	TransactTime             int64                   `json:"transactTime"`
	Price                    string                  `json:"price"`
	OrigQuantity             string                  `json:"origQty"`
	OrigQuoteOrderQuantity   string                  `json:"origQuoteOrderQty"`
	ExecutedQuantity         string                  `json:"executedQty"`
	CummulativeQuoteQuantity string                  `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType         `json:"status"`
	TimeInForce              TimeInForceType         `json:"timeInForce"`
	Type                     OrderType               `json:"type"`
	Side                     SideType                `json:"side"`
	SelfTradePreventionMode  SelfTradePreventionMode `json:"selfTradePreventionMode"`
}

// CancelOCOResponse may be returned included in a CancelOpenOrdersResponse.
type CancelOCOResponse struct {
	OrderListID       int64             `json:"orderListId"`
	ContingencyType   string            `json:"contingencyType"`
	ListStatusType    string            `json:"listStatusType"`
	ListOrderStatus   string            `json:"listOrderStatus"`
	ListClientOrderID string            `json:"listClientOrderId"`
	TransactionTime   int64             `json:"transactionTime"`
	Symbol            string            `json:"symbol"`
	Orders            []*OCOOrder       `json:"orders"`
	OrderReports      []*OCOOrderReport `json:"orderReports"`
}
