package futures

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// AlgoOrderStatusType defines algorithmic order status type.
type AlgoOrderStatusType string

const (
	AlgoOrderStatusTypeNew AlgoOrderStatusType = "NEW"
	// AlgoOrderStatusTypePartiallyFilled AlgoOrderStatusType = "PARTIALLY_FILLED"
	// AlgoOrderStatusTypeFilled          AlgoOrderStatusType = "FILLED"
	AlgoOrderStatusTypeCanceled AlgoOrderStatusType = "CANCELED"
	AlgoOrderStatusTypeRejected AlgoOrderStatusType = "REJECTED"
	AlgoOrderStatusTypeExpired  AlgoOrderStatusType = "EXPIRED"
)

// OrderAlgoType defines the algorithmic order type.
type OrderAlgoType string

const (
	OrderAlgoTypeConditional OrderAlgoType = "CONDITIONAL"
)

// AlgoOrderType defines the type of algorithmic order.
type AlgoOrderType string

const (
	AlgoOrderTypeStop               AlgoOrderType = "STOP"
	AlgoOrderTypeStopMarket         AlgoOrderType = "STOP_MARKET"
	AlgoOrderTypeTakeProfitMarket   AlgoOrderType = "TAKE_PROFIT_MARKET"
	AlgoOrderTypeTakeProfit         AlgoOrderType = "TAKE_PROFIT"
	AlgoOrderTypeTrailingStopMarket AlgoOrderType = "TRAILING_STOP_MARKET"
)

// CreateAlgoOrderService creates a new algorithmic order.
type CreateAlgoOrderService struct {
	c                       *Client
	algoType                OrderAlgoType // required
	symbol                  string        // required
	side                    SideType      // required
	_type                   AlgoOrderType // required
	positionSide            *PositionSideType
	timeInForceType         *TimeInForceType
	quantity                *string // Cannot be sent with closePosition=true(Close-All)
	price                   *string
	triggerPrice            *string
	workingType             *WorkingType
	priceMatch              *PriceMatchType
	closePosition           *bool
	priceProtect            *bool
	reduceOnly              *bool
	activationPrice         *string
	callbackRate            *string
	clientAlgoId            *string
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            *int64

	param map[string]any
}

// newCreateAlgoOrderService creates a new CreateAlgoOrderService instance.
func newCreateAlgoOrderService(c *Client) *CreateAlgoOrderService {
	return &CreateAlgoOrderService{
		c:        c,
		algoType: OrderAlgoTypeConditional, // default CONDITIONAL
		param: map[string]any{
			"algoType": OrderAlgoTypeConditional,
		},
	}
}

// AlgoType sets the algorithmic order type.
func (s *CreateAlgoOrderService) AlgoType(algoType OrderAlgoType) *CreateAlgoOrderService {
	s.algoType = algoType
	s.param["algoType"] = algoType
	return s
}

// Symbol sets the trading symbol.
func (s *CreateAlgoOrderService) Symbol(symbol string) *CreateAlgoOrderService {
	s.symbol = symbol
	s.param["symbol"] = symbol
	return s
}

// Side sets the order side.
func (s *CreateAlgoOrderService) Side(side SideType) *CreateAlgoOrderService {
	s.side = side
	s.param["side"] = side
	return s
}

// Type sets the algorithmic order type.
func (s *CreateAlgoOrderService) Type(_type AlgoOrderType) *CreateAlgoOrderService {
	s._type = _type
	s.param["type"] = _type
	return s
}

// PositionSide sets the position side.
func (s *CreateAlgoOrderService) PositionSide(positionSide PositionSideType) *CreateAlgoOrderService {
	s.positionSide = &positionSide
	s.param["positionSide"] = positionSide
	return s
}

// TimeInForce sets the time in force type.
func (s *CreateAlgoOrderService) TimeInForce(timeInForceType TimeInForceType) *CreateAlgoOrderService {
	s.timeInForceType = &timeInForceType
	s.param["timeInForce"] = timeInForceType
	return s
}

// Quantity sets the order quantity.
func (s *CreateAlgoOrderService) Quantity(quantity string) *CreateAlgoOrderService {
	s.quantity = &quantity
	s.param["quantity"] = quantity
	return s
}

// Price sets the order price.
func (s *CreateAlgoOrderService) Price(price string) *CreateAlgoOrderService {
	s.price = &price
	s.param["price"] = price
	return s
}

// TriggerPrice sets the trigger price.
func (s *CreateAlgoOrderService) TriggerPrice(triggerPrice string) *CreateAlgoOrderService {
	s.triggerPrice = &triggerPrice
	s.param["triggerPrice"] = triggerPrice
	return s
}

// WorkingType sets the working type.
func (s *CreateAlgoOrderService) WorkingType(workingType WorkingType) *CreateAlgoOrderService {
	s.workingType = &workingType
	s.param["workingType"] = workingType
	return s
}

// PriceMatch sets the price match type.
func (s *CreateAlgoOrderService) PriceMatch(priceMatch PriceMatchType) *CreateAlgoOrderService {
	s.priceMatch = &priceMatch
	s.param["priceMatch"] = priceMatch
	return s
}

// ClosePosition sets whether to close the position.
func (s *CreateAlgoOrderService) ClosePosition(closePosition bool) *CreateAlgoOrderService {
	s.closePosition = &closePosition
	if closePosition {
		s.param["closePosition"] = "true"
	} else {
		s.param["closePosition"] = "false"
	}
	return s
}

// PriceProtect sets whether to enable price protection.
func (s *CreateAlgoOrderService) PriceProtect(priceProtect bool) *CreateAlgoOrderService {
	s.priceProtect = &priceProtect
	if priceProtect {
		s.param["priceProtect"] = "true"
	} else {
		s.param["priceProtect"] = "false"
	}
	return s
}

// ReduceOnly sets whether the order is reduce-only.
func (s *CreateAlgoOrderService) ReduceOnly(reduceOnly bool) *CreateAlgoOrderService {
	s.reduceOnly = &reduceOnly
	if reduceOnly {
		s.param["reduceOnly"] = "true"
	} else {
		s.param["reduceOnly"] = "false"
	}
	return s
}

// ActivationPrice sets the activation price for trailing stop orders.
func (s *CreateAlgoOrderService) ActivationPrice(activationPrice string) *CreateAlgoOrderService {
	s.activationPrice = &activationPrice
	if activationPrice != "" {
		s.param["activationPrice"] = activationPrice
	}
	return s
}

// CallbackRate sets the callback rate for trailing stop orders.
func (s *CreateAlgoOrderService) CallbackRate(callbackRate string) *CreateAlgoOrderService {
	s.callbackRate = &callbackRate
	if callbackRate != "" {
		s.param["callbackRate"] = callbackRate
	}
	return s
}

// ClientAlgoId sets the client-defined algorithmic order ID.
func (s *CreateAlgoOrderService) ClientAlgoId(clientAlgoId string) *CreateAlgoOrderService {
	s.clientAlgoId = &clientAlgoId
	if clientAlgoId != "" {
		s.param["clientAlgoId"] = clientAlgoId
	}
	return s
}

// SelfTradePreventionMode sets the self-trade prevention mode.
func (s *CreateAlgoOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateAlgoOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	if selfTradePreventionMode != "" {
		s.param["selfTradePreventionMode"] = selfTradePreventionMode
	}
	return s
}

// GoodTillDate sets the good till date for the order.
func (s *CreateAlgoOrderService) GoodTillDate(goodTillDate int64) *CreateAlgoOrderService {
	s.goodTillDate = &goodTillDate
	if goodTillDate != 0 {
		s.param["goodTillDate"] = goodTillDate
	}
	return s
}

// CreateAlgoOrderResp represents the response from creating an algorithmic order.
type CreateAlgoOrderResp struct {
	AlgoId                  int64                   `json:"algoId"`
	ClientAlgoId            string                  `json:"clientAlgoId"`
	AlgoType                OrderAlgoType           `json:"algoType"`
	OrderType               AlgoOrderType           `json:"orderType"`
	Symbol                  string                  `json:"symbol"`
	Side                    SideType                `json:"side"`
	PositionSide            PositionSideType        `json:"positionSide"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Quantity                string                  `json:"quantity"`
	AlgoStatus              AlgoOrderStatusType     `json:"algoStatus"`
	TriggerPrice            string                  `json:"triggerPrice"`
	Price                   string                  `json:"price"`
	IcebergQuantity         *int64                  `json:"icebergQuantity,omitempty"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	WorkingType             WorkingType             `json:"workingType"`
	PriceMatch              PriceMatchType          `json:"priceMatch"`
	ClosePosition           bool                    `json:"closePosition"`
	PriceProtect            bool                    `json:"priceProtect"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	ActivatePrice           string                  `json:"activatePrice"`
	CallbackRate            string                  `json:"callbackRate"`
	CreateTime              int64                   `json:"createTime"`
	UpdateTime              int64                   `json:"updateTime"`
	TriggerTime             int64                   `json:"triggerTime"`
	GoodTillDate            int64                   `json:"goodTillDate"`
}

// Do sends the request to create an algorithmic order.
func (s *CreateAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (*CreateAlgoOrderResp, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/algoOrder",
		secType:  secTypeSigned,
	}
	r.setFormParams(s.param)
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := &CreateAlgoOrderResp{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelAlgoOrderService cancels an algorithmic order.
type CancelAlgoOrderService struct {
	c            *Client
	algoID       int64
	clientAlgoID *string
}

// AlgoID sets the algorithmic order ID.
func (s *CancelAlgoOrderService) AlgoID(algoID int64) *CancelAlgoOrderService {
	s.algoID = algoID
	return s
}

// ClientAlgoID sets the client-defined algorithmic order ID.
func (s *CancelAlgoOrderService) ClientAlgoID(clientAlgoID string) *CancelAlgoOrderService {
	s.clientAlgoID = &clientAlgoID
	return s
}

// CancelAlgoOrderResp represents the response from canceling an algorithmic order.
type CancelAlgoOrderResp struct {
	AlgoId       int    `json:"algoId"`
	ClientAlgoId string `json:"clientAlgoId"`
	Code         string `json:"code"`
	Message      string `json:"msg"`
}

// Do sends the request to cancel an algorithmic order.
func (s *CancelAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (*CancelAlgoOrderResp, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/algoOrder",
		secType:  secTypeSigned,
	}
	param := map[string]any{}
	if s.algoID != 0 {
		param["algoId"] = s.algoID
	}
	if s.clientAlgoID != nil {
		param["clientAlgoId"] = *s.clientAlgoID
	}
	r.setFormParams(param)
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := &CancelAlgoOrderResp{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelAllAlgoOpenOrdersService cancels all open algorithmic orders.
type CancelAllAlgoOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol sets the trading symbol.
func (s *CancelAllAlgoOpenOrdersService) Symbol(symbol string) *CancelAllAlgoOpenOrdersService {
	s.symbol = symbol
	return s
}

// CancelAllAlgoOpenOrdersResp represents the response from canceling all open algorithmic orders.
type CancelAllAlgoOpenOrdersResp struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// Do sends the request to cancel all open algorithmic orders.
func (s *CancelAllAlgoOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/algoOpenOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	res := &CancelAllAlgoOpenOrdersResp{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// GetAlgoOrderService gets an algorithmic order.
type GetAlgoOrderService struct {
	c            *Client
	algoId       *int64
	clientAlgoId *string
}

// AlgoID sets the algorithmic order ID.
func (s *GetAlgoOrderService) AlgoID(algoId int64) *GetAlgoOrderService {
	s.algoId = &algoId
	return s
}

// ClientAlgoID sets the client-defined algorithmic order ID.
func (s *GetAlgoOrderService) ClientAlgoID(clientAlgoId string) *GetAlgoOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// GetAlgoOrderResp represents the response from getting an algorithmic order.
type GetAlgoOrderResp struct {
	AlgoId                  int64                   `json:"algoId"`
	ClientAlgoId            string                  `json:"clientAlgoId"`
	AlgoType                OrderAlgoType           `json:"algoType"`
	OrderType               AlgoOrderType           `json:"orderType"`
	Symbol                  string                  `json:"symbol"`
	Side                    SideType                `json:"side"`
	PositionSide            PositionSideType        `json:"positionSide"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Quantity                string                  `json:"quantity"`
	AlgoStatus              AlgoOrderStatusType     `json:"algoStatus"`
	ActualOrderId           string                  `json:"actualOrderId"`
	ActualPrice             string                  `json:"actualPrice"`
	TriggerPrice            string                  `json:"triggerPrice"`
	Price                   string                  `json:"price"`
	IcebergQuantity         *int64                  `json:"icebergQuantity,omitempty"`
	TpTriggerPrice          string                  `json:"tpTriggerPrice"`
	TpPrice                 string                  `json:"tpPrice"`
	SlTriggerPrice          string                  `json:"slTriggerPrice"`
	SlPrice                 string                  `json:"slPrice"`
	TpOrderType             string                  `json:"tpOrderType"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	WorkingType             WorkingType             `json:"workingType"`
	PriceMatch              PriceMatchType          `json:"priceMatch"`
	ClosePosition           bool                    `json:"closePosition"`
	PriceProtect            bool                    `json:"priceProtect"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	CreateTime              int64                   `json:"createTime"`
	UpdateTime              int64                   `json:"updateTime"`
	TriggerTime             int64                   `json:"triggerTime"`
	GoodTillDate            int64                   `json:"goodTillDate"`
}

// Do sends the request to get an algorithmic order.
func (s *GetAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (*GetAlgoOrderResp, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/algoOrder",
		secType:  secTypeSigned,
	}
	if s.algoId != nil {
		r.setParam("algoId", *s.algoId)
	}
	if s.clientAlgoId != nil {
		r.setParam("clientAlgoId", *s.clientAlgoId)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := &GetAlgoOrderResp{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListOpenAlgoOrdersService lists all open algorithmic orders.
type ListOpenAlgoOrdersService struct {
	c        *Client
	algoType *OrderAlgoType
	symbol   *string
	algoId   *int64
}

// AlgoType sets the algorithmic order type.
func (s *ListOpenAlgoOrdersService) AlgoType(algoType OrderAlgoType) *ListOpenAlgoOrdersService {
	s.algoType = &algoType
	return s
}

// Symbol sets the trading symbol.
func (s *ListOpenAlgoOrdersService) Symbol(symbol string) *ListOpenAlgoOrdersService {
	s.symbol = &symbol
	return s
}

// AlgoID sets the algorithmic order ID.
func (s *ListOpenAlgoOrdersService) AlgoID(algoId int64) *ListOpenAlgoOrdersService {
	s.algoId = &algoId
	return s
}

// Do sends the request to list all open algorithmic orders.
func (s *ListOpenAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]GetAlgoOrderResp, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openAlgoOrders",
		secType:  secTypeSigned,
	}
	if s.algoType != nil {
		r.setParam("algoType", string(*s.algoType))
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.algoId != nil {
		r.setParam("algoId", *s.algoId)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []GetAlgoOrderResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListAllAlgoOrdersService lists all historical algorithmic orders.
type ListAllAlgoOrdersService struct {
	c         *Client
	symbol    string // required
	algoId    *int64
	startTime *int64
	endTime   *int64
	page      *int
	limit     *int
}

// Symbol sets the trading symbol.
func (s *ListAllAlgoOrdersService) Symbol(symbol string) *ListAllAlgoOrdersService {
	s.symbol = symbol
	return s
}

// AlgoID sets the algorithmic order ID.
func (s *ListAllAlgoOrdersService) AlgoID(algoId int64) *ListAllAlgoOrdersService {
	s.algoId = &algoId
	return s
}

// StartTime sets the start time for filtering.
func (s *ListAllAlgoOrdersService) StartTime(startTime int64) *ListAllAlgoOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime sets the end time for filtering.
func (s *ListAllAlgoOrdersService) EndTime(endTime int64) *ListAllAlgoOrdersService {
	s.endTime = &endTime
	return s
}

// Page sets the page number for pagination.
func (s *ListAllAlgoOrdersService) Page(page int) *ListAllAlgoOrdersService {
	s.page = &page
	return s
}

// Limit sets the number of items per page.
func (s *ListAllAlgoOrdersService) Limit(limit int) *ListAllAlgoOrdersService {
	s.limit = &limit
	return s
}

// Do sends the request to list all historical algorithmic orders.
func (s *ListAllAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]GetAlgoOrderResp, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allAlgoOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.algoId != nil {
		r.setParam("algoId", *s.algoId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []GetAlgoOrderResp
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil

}
