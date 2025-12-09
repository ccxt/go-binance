package futures

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/adshao/go-binance/v2/common"
)

// CreateAlgoOrderService create order
type CreateAlgoOrderService struct {
	c                       *Client
	algoType *string
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	price                   *string
	triggerPrice               *string
	workingType             *WorkingType
	priceMatch              *PriceMatchType
	closePosition           *string
	priceProtect            *string
	reduceOnly              *string
	activationPrice         *string
	callbackRate            *string
	clientAlgoID        *string
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            int64
}

// Symbol set symbol
func (s *CreateAlgoOrderService) Symbol(symbol string) *CreateAlgoOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateAlgoOrderService) Side(side SideType) *CreateAlgoOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateAlgoOrderService) PositionSide(positionSide PositionSideType) *CreateAlgoOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateAlgoOrderService) Type(orderType OrderType) *CreateAlgoOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateAlgoOrderService) TimeInForce(timeInForce TimeInForceType) *CreateAlgoOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateAlgoOrderService) Quantity(quantity string) *CreateAlgoOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateAlgoOrderService) ReduceOnly(reduceOnly bool) *CreateAlgoOrderService {
	reduceOnlyStr := strconv.FormatBool(reduceOnly)
	s.reduceOnly = &reduceOnlyStr
	return s
}

// Price set price
func (s *CreateAlgoOrderService) Price(price string) *CreateAlgoOrderService {
	s.price = &price
	return s
}

// PriceMatch set priceMatch
func (s *CreateAlgoOrderService) PriceMatch(priceMatch PriceMatchType) *CreateAlgoOrderService {
	s.priceMatch = &priceMatch
	return s
}

// ClientAlgoID set clientAlgoID rule: ^[\.A-Z\:/a-z0-9_-]{1,36}$
func (s *CreateAlgoOrderService) ClientAlgoID(clientAlgoID string) *CreateAlgoOrderService {
	s.clientAlgoID = &clientAlgoID
	return s
}

// StopPrice set triggerPrice
func (s *CreateAlgoOrderService) StopPrice(stopPrice string) *CreateAlgoOrderService {
	s.triggerPrice = &stopPrice
	return s
}

// TriggerPrice set triggerPrice
func (s *CreateAlgoOrderService) TriggerPrice(triggerPrice string) *CreateAlgoOrderService {
	s.triggerPrice = &triggerPrice
	return s
}

// WorkingType set workingType
func (s *CreateAlgoOrderService) WorkingType(workingType WorkingType) *CreateAlgoOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateAlgoOrderService) ActivationPrice(activationPrice string) *CreateAlgoOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateAlgoOrderService) CallbackRate(callbackRate string) *CreateAlgoOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateAlgoOrderService) PriceProtect(priceProtect bool) *CreateAlgoOrderService {
	priceProtectStr := strconv.FormatBool(priceProtect)
	s.priceProtect = &priceProtectStr
	return s
}

// ClosePosition set closePosition
func (s *CreateAlgoOrderService) ClosePosition(closePosition bool) *CreateAlgoOrderService {
	closePositionStr := strconv.FormatBool(closePosition)
	s.closePosition = &closePositionStr
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateAlgoOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateAlgoOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// GoodTillDate set goodTillDate
func (s *CreateAlgoOrderService) GoodTillDate(goodTillDate int64) *CreateAlgoOrderService {
	s.goodTillDate = goodTillDate
	return s
}

func (s *CreateAlgoOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
	}
	if s.quantity != "" {
		m["quantity"] = s.quantity
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.priceMatch != nil {
		m["priceMatch"] = *s.priceMatch
	}
	if s.clientAlgoID != nil {
		m["clientAlgoId"] = *s.clientAlgoID
	} else {
		m["clientAlgoId"] = common.GenerateSwapId()
	}
	if s.triggerPrice != nil {
		m["triggerPrice"] = *s.triggerPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.goodTillDate > 0 && *s.timeInForce == TimeInForceTypeGTD {
		m["goodTillDate"] = s.goodTillDate
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateAlgoOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/fapi/v1/algoOrder", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateAlgoOrderResponse define create algo order response
type CreateAlgoOrderResponse struct {
	Symbol                  string           `json:"symbol"`                      //
	OrderID                 int64            `json:"orderId"`                     //
	ClientAlgoID           string           `json:"clientAlgoId"`               //
	Price                   string           `json:"price"`                       //
	OrigQuantity            string           `json:"origQty"`                     //
	ExecutedQuantity        string           `json:"executedQty"`                 //
	CumQuote                string           `json:"cumQuote"`                    //
	ReduceOnly              bool             `json:"reduceOnly"`                  //
	Status                  OrderStatusType  `json:"status"`                      //
	TriggerPrice               string           `json:"triggerPrice"`                   // please ignore when order type is TRAILING_STOP_MARKET
	TimeInForce             TimeInForceType  `json:"timeInForce"`                 //
	Type                    OrderType        `json:"type"`                        //
	Side                    SideType         `json:"side"`                        //
	UpdateTime              int64            `json:"updateTime"`                  // update time
	WorkingType             WorkingType      `json:"workingType"`                 //
	ActivatePrice           string           `json:"activatePrice"`               // activation price, only return with TRAILING_STOP_MARKET order
	PriceRate               string           `json:"priceRate"`                   // callback rate, only return with TRAILING_STOP_MARKET order
	AvgPrice                string           `json:"avgPrice"`                    //
	PositionSide            PositionSideType `json:"positionSide"`                //
	ClosePosition           bool             `json:"closePosition"`               // if Close-All
	PriceProtect            bool             `json:"priceProtect"`                // if conditional order trigger is protected
	PriceMatch              string           `json:"priceMatch"`                  // price match mode
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`     // self trading prevention mode
	GoodTillDate            int64            `json:"goodTillDate"`                // order pre-set auto cancel time for TIF GTD order
	CumQty                  string           `json:"cumQty"`                      //
	RateLimitOrder10s       string           `json:"rateLimitOrder10s,omitempty"` //
	RateLimitOrder1m        string           `json:"rateLimitOrder1m,omitempty"`  //
}

// ListOpenAlgoOrdersService list opened orders
type ListOpenAlgoOrdersService struct {
	c      *Client
	symbol string
	algoType string
	algoID int64 // todo: test it
}

// Symbol set symbol
func (s *ListOpenAlgoOrdersService) Symbol(symbol string) *ListOpenAlgoOrdersService {
	s.symbol = symbol
	return s
}

// AlgoType set algoType
func (s *ListOpenAlgoOrdersService) AlgoType(algoType string) *ListOpenAlgoOrdersService {
	s.algoType = algoType
	return s
}

// AlgoID set algoID
func (s *ListOpenAlgoOrdersService) AlgoID(algoID int64) *ListOpenAlgoOrdersService {
	s.algoID = algoID
	return s
}

// Do send request
func (s *ListOpenAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openAlgoOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.algoType != "" {
		r.setParam("algoType", s.algoType)
	}
	if s.algoID != 0 {
		r.setParam("algoId", s.algoID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
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

// GetAlgoOrderService query algo order
type GetAlgoOrderService struct {
	c                 *Client
	algoID           *int64
	clientAlgoID *string
}

func (s *GetAlgoOrderService) AlgoID(algoID int64) *GetAlgoOrderService {
	s.algoID = &algoID
	return s
}

func (s *GetAlgoOrderService) ClientAlgoID(clientAlgoID string) *GetAlgoOrderService {
	s.clientAlgoID = &clientAlgoID
	return s
}

func (s *GetAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/algoOrder",
		secType:  secTypeSigned,
	}
	if s.algoID == nil && s.clientAlgoID == nil {
		return nil, errors.New("either algoID or clientAlgoID must be sent")
	}
	if s.algoID != nil {
		r.setParam("algoId", *s.algoID)
	}
	if s.clientAlgoID != nil {
		r.setParam("clientAlgoId", *s.clientAlgoID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
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

// AlgoOrder define algo order info
// todo: check this struct
type AlgoOrder struct {
	Symbol                  string           `json:"symbol"`
	OrderID                 int64            `json:"orderId"`
	ClientOrderID           string           `json:"clientOrderId"`
	Price                   string           `json:"price"`
	ReduceOnly              bool             `json:"reduceOnly"`
	OrigQuantity            string           `json:"origQty"`
	ExecutedQuantity        string           `json:"executedQty"`
	CumQuantity             string           `json:"cumQty"` // deprecated: use ExecutedQuantity instead
	CumQuote                string           `json:"cumQuote"`
	Status                  OrderStatusType  `json:"status"`
	TimeInForce             TimeInForceType  `json:"timeInForce"`
	Type                    OrderType        `json:"type"`
	Side                    SideType         `json:"side"`
	StopPrice               string           `json:"stopPrice"`
	Time                    int64            `json:"time"`
	UpdateTime              int64            `json:"updateTime"`
	WorkingType             WorkingType      `json:"workingType"`
	ActivatePrice           string           `json:"activatePrice"`
	PriceRate               string           `json:"priceRate"`
	AvgPrice                string           `json:"avgPrice"`
	OrigType                OrderType        `json:"origType"`
	PositionSide            PositionSideType `json:"positionSide"`
	PriceProtect            bool             `json:"priceProtect"`
	ClosePosition           bool             `json:"closePosition"`
	PriceMatch              string           `json:"priceMatch"`
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`
	GoodTillDate            int64            `json:"goodTillDate"`
}

// ListAlgoOrdersService all account algo orders; active, canceled, triggered, or finished
type ListAlgoOrdersService struct {
	c         *Client
	symbol    string
	algoID   *int64
	startTime *int64
	endTime   *int64
	page     *int
	limit     *int
}

// Symbol set symbol
func (s *ListAlgoOrdersService) Symbol(symbol string) *ListAlgoOrdersService {
	s.symbol = symbol
	return s
}

// AlgoID set algoID
func (s *ListAlgoOrdersService) AlgoID(algoID int64) *ListAlgoOrdersService {
	s.algoID = &algoID
	return s
}

// StartTime set starttime
func (s *ListAlgoOrdersService) StartTime(startTime int64) *ListAlgoOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListAlgoOrdersService) EndTime(endTime int64) *ListAlgoOrdersService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *ListAlgoOrdersService) Page(page int) *ListAlgoOrdersService {
	s.page = &page
	return s
}

// Limit set limit
func (s *ListAlgoOrdersService) Limit(limit int) *ListAlgoOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allAlgoOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.algoID != nil {
		r.setParam("algoId", *s.algoID)
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
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// CancelAlgoOrderService cancel an order
type CancelAlgoOrderService struct {
	c                 *Client
	symbol            string
	algoID           *int64
	clientAlgoID *string
}

// Symbol set symbol
func (s *CancelAlgoOrderService) Symbol(symbol string) *CancelAlgoOrderService {
	s.symbol = symbol
	return s
}

// AlgoID set algoID
func (s *CancelAlgoOrderService) AlgoID(algoID int64) *CancelAlgoOrderService {
	s.algoID = &algoID
	return s
}

// ClientAlgoID set clientAlgoID
func (s *CancelAlgoOrderService) ClientAlgoID(clientAlgoID string) *CancelAlgoOrderService {
	s.clientAlgoID = &clientAlgoID
	return s
}

// Do send request
func (s *CancelAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelAlgoOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/algoOrder",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.algoID != nil {
		r.setFormParam("algoId", *s.algoID)
	}
	if s.clientAlgoID != nil {
		r.setFormParam("clientalgoid", *s.clientAlgoID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelAlgoOrderResponse define response of canceling algo order
// todo: check this
type CancelAlgoOrderResponse struct {
	ClientOrderID           string                  `json:"clientOrderId"`
	CumQuantity             string                  `json:"cumQty"` // deprecated: use ExecutedQuantity instead
	CumQuote                string                  `json:"cumQuote"`
	ExecutedQuantity        string                  `json:"executedQty"`
	OrderID                 int64                   `json:"orderId"`
	OrigQuantity            string                  `json:"origQty"`
	Price                   string                  `json:"price"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	Side                    SideType                `json:"side"`
	Status                  OrderStatusType         `json:"status"`
	StopPrice               string                  `json:"stopPrice"`
	Symbol                  string                  `json:"symbol"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	UpdateTime              int64                   `json:"updateTime"`
	WorkingType             WorkingType             `json:"workingType"`
	ActivatePrice           string                  `json:"activatePrice"`
	PriceRate               string                  `json:"priceRate"`
	OrigType                string                  `json:"origType"`
	PositionSide            PositionSideType        `json:"positionSide"`
	PriceProtect            bool                    `json:"priceProtect"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
}

// CancelAllOpenAlgoOrdersService cancel all open algo orders
type CancelAllOpenAlgoOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllOpenAlgoOrdersService) Symbol(symbol string) *CancelAllOpenAlgoOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllOpenAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/algoOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
