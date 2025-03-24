package futures

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2/common"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	reduceOnly              *string
	price                   *string
	newClientOrderID        *string
	stopPrice               *string
	workingType             *WorkingType
	activationPrice         *string
	callbackRate            *string
	priceProtect            *string
	newOrderRespType        NewOrderRespType
	closePosition           *string
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

// PositionSide set side
func (s *CreateOrderService) PositionSide(positionSide PositionSideType) *CreateOrderService {
	s.positionSide = &positionSide
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
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	reduceOnlyStr := strconv.FormatBool(reduceOnly)
	s.reduceOnly = &reduceOnlyStr
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

// WorkingType set workingType
func (s *CreateOrderService) WorkingType(workingType WorkingType) *CreateOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateOrderService) ActivationPrice(activationPrice string) *CreateOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateOrderService) CallbackRate(callbackRate string) *CreateOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateOrderService) PriceProtect(priceProtect bool) *CreateOrderService {
	priceProtectStr := strconv.FormatBool(priceProtect)
	s.priceProtect = &priceProtectStr
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *CreateOrderService) ClosePosition(closePosition bool) *CreateOrderService {
	closePositionStr := strconv.FormatBool(closePosition)
	s.closePosition = &closePositionStr
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"newOrderRespType": s.newOrderRespType,
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
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	} else {
		m["newClientOrderId"] = common.GenerateSwapId()
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
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
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/fapi/v1/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	Symbol                  string           `json:"symbol"`                      //
	OrderID                 int64            `json:"orderId"`                     //
	ClientOrderID           string           `json:"clientOrderId"`               //
	Price                   string           `json:"price"`                       //
	OrigQuantity            string           `json:"origQty"`                     //
	ExecutedQuantity        string           `json:"executedQty"`                 //
	CumQuote                string           `json:"cumQuote"`                    //
	ReduceOnly              bool             `json:"reduceOnly"`                  //
	Status                  OrderStatusType  `json:"status"`                      //
	StopPrice               string           `json:"stopPrice"`                   // please ignore when order type is TRAILING_STOP_MARKET
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
	OrigType                OrderType        `json:"origType"`                    //
	RateLimitOrder10s       string           `json:"rateLimitOrder10s,omitempty"` //
	RateLimitOrder1m        string           `json:"rateLimitOrder1m,omitempty"`  //
}

// ModifyOrderService create order
type ModifyOrderService struct {
	c                 *Client
	orderID           *int64
	origClientOrderID *string
	symbol            string
	side              SideType
	quantity          string
	price             *string
	priceMatch        *PriceMatchType
}

// Symbol set symbol
func (s *ModifyOrderService) Symbol(symbol string) *ModifyOrderService {
	s.symbol = symbol
	return s
}

// OrderID will prevail over OrigClientOrderID
func (s *ModifyOrderService) OrderID(orderID int64) *ModifyOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID is not necessary if OrderID is provided
func (s *ModifyOrderService) OrigClientOrderID(origClientOrderID string) *ModifyOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Side set side
func (s *ModifyOrderService) Side(side SideType) *ModifyOrderService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *ModifyOrderService) Quantity(quantity string) *ModifyOrderService {
	s.quantity = quantity
	return s
}

// Price set price
func (s *ModifyOrderService) Price(price string) *ModifyOrderService {
	s.price = &price
	return s
}

// PriceMatch set priceMatch
func (s *ModifyOrderService) PriceMatch(priceMatch PriceMatchType) *ModifyOrderService {
	s.priceMatch = &priceMatch
	return s
}

func (s *ModifyOrderService) modifyOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
	}
	if s.orderID != nil {
		m["orderId"] = *s.orderID
	}
	if s.origClientOrderID != nil {
		m["origClientOrderId"] = *s.origClientOrderID
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.priceMatch != nil {
		m["priceMatch"] = *s.priceMatch
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request:
//   - Either orderId or origClientOrderId must be sent, and the orderId will prevail if both are sent
//   - Either price or priceMatch must be sent. Sending both will fail the request
//   - When the new quantity or price doesn't satisfy PriceFilter / PercentPriceFilter / LotSizeFilter,
//     amendment will be rejected and the order will stay as it is
//   - However the order will be cancelled by the amendment in the following situations:
//     -- when the order is in partially filled status and the new quantity <= executedQty
//     -- when the order is TimeInForceTypeGTX and the new price will cause it to be executed immediately
//   - One order can only be modified for less than 10000 times
//   - Will set ModifyOrderResponse.SelfTradePreventionMode to "NONE"
func (s *ModifyOrderService) Do(ctx context.Context, opts ...RequestOption) (res *ModifyOrderResponse, err error) {
	data, _, err := s.modifyOrder(ctx, "/fapi/v1/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(ModifyOrderResponse)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}
	return res, nil
}

type ModifyOrderResponse struct {
	OrderID                 int64            `json:"orderId"`
	Symbol                  string           `json:"symbol"`
	Pair                    string           `json:"pair"`
	Status                  OrderStatusType  `json:"status"`
	ClientOrderID           string           `json:"clientOrderId"`
	Price                   string           `json:"price"`
	AveragePrice            string           `json:"avgPrice"`
	OriginalQuantity        string           `json:"origQty"`
	ExecutedQuantity        string           `json:"executedQty"`
	CumulativeQuantity      string           `json:"cumQty"`
	CumulativeBase          string           `json:"cumBase"`
	TimeInForce             TimeInForceType  `json:"timeInForce"`
	Type                    OrderType        `json:"type"`
	ReduceOnly              bool             `json:"reduceOnly"`
	ClosePosition           bool             `json:"closePosition"`
	Side                    SideType         `json:"side"`
	PositionSide            PositionSideType `json:"positionSide"`
	StopPrice               string           `json:"stopPrice"`
	WorkingType             WorkingType      `json:"workingType"`
	PriceProtect            bool             `json:"priceProtect"` // if conditional order trigger is protected
	OriginalType            OrderType        `json:"origType"`
	PriceMatch              PriceMatchType   `json:"priceMatch"`
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`
	GoodTillDate            int64            `json:"goodTillDate"` // order pre-set auto cancel time for TIF GTD order
	UpdateTime              int64            `json:"updateTime"`
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
		endpoint: "/fapi/v1/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
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

// GetOpenOrderService query current open order
type GetOpenOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

func (s *GetOpenOrderService) Symbol(symbol string) *GetOpenOrderService {
	s.symbol = symbol
	return s
}

func (s *GetOpenOrderService) OrderID(orderID int64) *GetOpenOrderService {
	s.orderID = &orderID
	return s
}

func (s *GetOpenOrderService) OrigClientOrderID(origClientOrderID string) *GetOpenOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *GetOpenOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID == nil && s.origClientOrderID == nil {
		return nil, errors.New("either orderId or origClientOrderId must be sent")
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
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
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
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

// Order define order info
type Order struct {
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
		endpoint: "/fapi/v1/allOrders",
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

// CancelOrderService cancel an order
type CancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
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

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
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

// CancelOrderResponse define response of canceling order
type CancelOrderResponse struct {
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

// CancelAllOpenOrdersService cancel all open orders
type CancelAllOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllOpenOrdersService) Symbol(symbol string) *CancelAllOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CancelMultiplesOrdersService cancel a list of orders
type CancelMultiplesOrdersService struct {
	c                     *Client
	symbol                string
	orderIDList           []int64
	origClientOrderIDList []string
}

// Symbol set symbol
func (s *CancelMultiplesOrdersService) Symbol(symbol string) *CancelMultiplesOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelMultiplesOrdersService) OrderIDList(orderIDList []int64) *CancelMultiplesOrdersService {
	s.orderIDList = orderIDList
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelMultiplesOrdersService) OrigClientOrderIDList(origClientOrderIDList []string) *CancelMultiplesOrdersService {
	s.origClientOrderIDList = origClientOrderIDList
	return s
}

// Do send request
func (s *CancelMultiplesOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderIDList != nil {
		// convert a slice of integers to a string e.g. [1 2 3] => "[1,2,3]"
		orderIDListString := strings.Join(strings.Fields(fmt.Sprint(s.orderIDList)), ",")
		r.setFormParam("orderIdList", orderIDListString)
	}
	if s.origClientOrderIDList != nil {
		r.setFormParam("origClientOrderIdList", s.origClientOrderIDList)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*CancelOrderResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CancelOrderResponse{}, err
	}
	return res, nil
}

// ListLiquidationOrdersService list liquidation orders
type ListLiquidationOrdersService struct {
	c         *Client
	symbol    *string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListLiquidationOrdersService) Symbol(symbol string) *ListLiquidationOrdersService {
	s.symbol = &symbol
	return s
}

// StartTime set startTime
func (s *ListLiquidationOrdersService) StartTime(startTime int64) *ListLiquidationOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *ListLiquidationOrdersService) EndTime(endTime int64) *ListLiquidationOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListLiquidationOrdersService) Limit(limit int) *ListLiquidationOrdersService {
	s.limit = &limit
	return s
}

// Do send request
// Deprecated: use /fapi/v1/forceOrders instead
func (s *ListLiquidationOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*LiquidationOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allForceOrders",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*LiquidationOrder{}, err
	}
	res = make([]*LiquidationOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LiquidationOrder{}, err
	}
	return res, nil
}

// LiquidationOrder define liquidation order
type LiquidationOrder struct {
	Symbol           string          `json:"symbol"`
	Price            string          `json:"price"`
	OrigQuantity     string          `json:"origQty"`
	ExecutedQuantity string          `json:"executedQty"`
	AveragePrice     string          `json:"avragePrice"`
	Status           OrderStatusType `json:"status"`
	TimeInForce      TimeInForceType `json:"timeInForce"`
	Type             OrderType       `json:"type"`
	Side             SideType        `json:"side"`
	Time             int64           `json:"time"`
}

// ListUserLiquidationOrdersService lists user's liquidation orders
type ListUserLiquidationOrdersService struct {
	c             *Client
	symbol        *string
	autoCloseType ForceOrderCloseType
	startTime     *int64
	endTime       *int64
	limit         *int
}

// Symbol set symbol
func (s *ListUserLiquidationOrdersService) Symbol(symbol string) *ListUserLiquidationOrdersService {
	s.symbol = &symbol
	return s
}

// AutoCloseType set symbol
func (s *ListUserLiquidationOrdersService) AutoCloseType(autoCloseType ForceOrderCloseType) *ListUserLiquidationOrdersService {
	s.autoCloseType = autoCloseType
	return s
}

// StartTime set startTime
func (s *ListUserLiquidationOrdersService) StartTime(startTime int64) *ListUserLiquidationOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListUserLiquidationOrdersService) EndTime(endTime int64) *ListUserLiquidationOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListUserLiquidationOrdersService) Limit(limit int) *ListUserLiquidationOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListUserLiquidationOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*UserLiquidationOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/forceOrders",
		secType:  secTypeSigned,
	}

	r.setParam("autoCloseType", s.autoCloseType)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*UserLiquidationOrder{}, err
	}
	res = make([]*UserLiquidationOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UserLiquidationOrder{}, err
	}
	return res, nil
}

// UserLiquidationOrder defines user's liquidation order
type UserLiquidationOrder struct {
	OrderId          int64            `json:"orderId"`
	Symbol           string           `json:"symbol"`
	Status           OrderStatusType  `json:"status"`
	ClientOrderId    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	AveragePrice     string           `json:"avgPrice"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuote         string           `json:"cumQuote"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ReduceOnly       bool             `json:"reduceOnly"`
	ClosePosition    bool             `json:"closePosition"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	StopPrice        string           `json:"stopPrice"`
	WorkingType      WorkingType      `json:"workingType"`
	OrigType         string           `json:"origType"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
}

type CreateBatchOrdersService struct {
	c      *Client
	orders []*CreateOrderService
}

// CreateBatchOrdersResponse contains the response from CreateBatchOrders operation
type CreateBatchOrdersResponse struct {
	// Total number of messages in the response
	N int
	// List of orders which were placed successfully which can have a length between 0 and N
	Orders []*Order
	// List of errors of length N, where each item corresponds to a nil value if
	// the order from that specific index was placed succeessfully OR an non-nil *APIError if there was an error with
	// the order at that index
	Errors []error
}

func newCreateBatchOrdersResponse(n int) *CreateBatchOrdersResponse {
	return &CreateBatchOrdersResponse{
		N:      n,
		Errors: make([]error, n),
	}
}

func (s *CreateBatchOrdersService) OrderList(orders []*CreateOrderService) *CreateBatchOrdersService {
	s.orders = orders
	return s
}

func (s *CreateBatchOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBatchOrdersResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}

	orders := []params{}
	for _, order := range s.orders {
		m := params{
			"symbol":           order.symbol,
			"side":             order.side,
			"type":             order.orderType,
			"quantity":         order.quantity,
			"newOrderRespType": order.newOrderRespType,
		}

		if order.positionSide != nil {
			m["positionSide"] = *order.positionSide
		}
		if order.timeInForce != nil {
			m["timeInForce"] = *order.timeInForce
		}
		if order.reduceOnly != nil {
			m["reduceOnly"] = *order.reduceOnly
		}
		if order.price != nil {
			m["price"] = *order.price
		}
		if order.newClientOrderID != nil {
			m["newClientOrderId"] = *order.newClientOrderID
		} else {
			m["newClientOrderId"] = common.GenerateSwapId()
		}
		if order.stopPrice != nil {
			m["stopPrice"] = *order.stopPrice
		}
		if order.workingType != nil {
			m["workingType"] = *order.workingType
		}
		if order.priceProtect != nil {
			m["priceProtect"] = *order.priceProtect
		}
		if order.activationPrice != nil {
			m["activationPrice"] = *order.activationPrice
		}
		if order.callbackRate != nil {
			m["callbackRate"] = *order.callbackRate
		}
		if order.closePosition != nil {
			m["closePosition"] = *order.closePosition
		}
		orders = append(orders, m)
	}
	b, err := json.Marshal(orders)
	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}
	m := params{
		"batchOrders": string(b),
	}

	r.setFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}

	rawMessages := make([]*json.RawMessage, 0)

	err = json.Unmarshal(data, &rawMessages)
	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}

	batchCreateOrdersResponse := newCreateBatchOrdersResponse(len(rawMessages))
	for i, j := range rawMessages {
		// check if response is an API error
		e := new(common.APIError)
		if err := json.Unmarshal(*j, e); err != nil {
			return nil, err
		}

		if e.Code > 0 || e.Message != "" {
			batchCreateOrdersResponse.Errors[i] = e
			continue
		}

		o := new(Order)
		if err := json.Unmarshal(*j, o); err != nil {
			return nil, err
		}

		batchCreateOrdersResponse.Orders = append(batchCreateOrdersResponse.Orders, o)
	}

	return batchCreateOrdersResponse, nil
}

// ModifyOrder contains parameters for order modification request
type ModifyOrder struct {
	orderID           *int64
	origClientOrderID *string
	symbol            string
	side              SideType
	quantity          string
	price             *string
	priceMatch        *PriceMatchType
}

// Symbol set symbol
func (s *ModifyOrder) Symbol(symbol string) *ModifyOrder {
	s.symbol = symbol
	return s
}

// OrderID will prevail over OrigClientOrderID
func (s *ModifyOrder) OrderID(orderID int64) *ModifyOrder {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID is not necessary if OrderID is provided
func (s *ModifyOrder) OrigClientOrderID(origClientOrderID string) *ModifyOrder {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Side set side
func (s *ModifyOrder) Side(side SideType) *ModifyOrder {
	s.side = side
	return s
}

// Quantity set quantity
func (s *ModifyOrder) Quantity(quantity string) *ModifyOrder {
	s.quantity = quantity
	return s
}

// Price set price
func (s *ModifyOrder) Price(price string) *ModifyOrder {
	s.price = &price
	return s
}

// PriceMatch set priceMatch
func (s *ModifyOrder) PriceMatch(priceMatch PriceMatchType) *ModifyOrder {
	s.priceMatch = &priceMatch
	return s
}

// ModifyBatchOrdersService handles batch modification of orders
type ModifyBatchOrdersService struct {
	c      *Client
	orders []*ModifyOrder
}

// CreateBatchOrdersResponse contains the response from CreateBatchOrders operation
type ModifyBatchOrdersResponse struct {
	// Total number of messages in the response
	N int
	// List of orders which were modified successfully which can have a length between 0 and N
	Orders []*Order
	// List of errors of length N, where each item corresponds to a nil value if
	// the order from that specific index was placed succeessfully OR an non-nil *APIError if there was an error with
	// the order at that index
	Errors []error
}

func newModifyBatchOrdersResponse(n int) *ModifyBatchOrdersResponse {
	return &ModifyBatchOrdersResponse{
		N:      n,
		Errors: make([]error, n),
	}
}

// OrderList set the list of ModifyOrder to be used in the ModifyBatchOrders operation
func (s *ModifyBatchOrdersService) OrderList(orders []*ModifyOrder) *ModifyBatchOrdersService {
	s.orders = orders
	return s
}

// Do sends a request to modify a batch of orders.
// It constructs the necessary parameters for each order and marshals them into a JSON payload.
// The function returns a ModifyBatchOrdersResponse, which contains the results of the modification attempt.
func (s *ModifyBatchOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *ModifyBatchOrdersResponse, err error) {
	// Create a new request with method PUT and the appropriate endpoint.
	r := &request{
		method:   http.MethodPut,
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}

	orders := []params{}
	// Iterate through the orders to construct parameters for each order.
	for _, order := range s.orders {
		m := params{
			"symbol":   order.symbol,
			"side":     order.side,
			"quantity": order.quantity,
			"price":    order.price,
		}

		// Convert orderID to string to avoid API error with code -1102.
		if order.orderID != nil {
			m["orderId"] = strconv.FormatInt(*order.orderID, 10)
		}
		if order.origClientOrderID != nil {
			m["origClientOrderId"] = *order.origClientOrderID
		} else {
			m["newClientOrderId"] = common.GenerateSwapId()
		}
		if order.priceMatch != nil {
			m["priceMatch"] = *order.priceMatch
		}

		orders = append(orders, m)
	}

	// Marshal the orders into a JSON payload.
	b, err := json.Marshal(orders)
	if err != nil {
		return &ModifyBatchOrdersResponse{}, err
	}

	// Set the marshaled orders as form parameters.
	m := params{
		"batchOrders": string(b),
	}
	r.setFormParams(m)

	// Call the API with the constructed request.
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &ModifyBatchOrdersResponse{}, err
	}

	rawMessages := make([]*json.RawMessage, 0)
	// Unmarshal the response into raw JSON messages.
	err = json.Unmarshal(data, &rawMessages)
	if err != nil {
		return &ModifyBatchOrdersResponse{}, err
	}

	// Create a response object to hold the results.
	batchModifyOrdersResponse := newModifyBatchOrdersResponse(len(rawMessages))
	for i, j := range rawMessages {
		// Check if the response contains an API error.
		e := new(common.APIError)
		if err := json.Unmarshal(*j, e); err != nil {
			return nil, err
		}

		// If there's an error code or message, record it and continue.
		if e.Code > 0 || e.Message != "" {
			batchModifyOrdersResponse.Errors[i] = e
			continue
		}

		// Otherwise, unmarshal the order information.
		o := new(Order)
		if err := json.Unmarshal(*j, o); err != nil {
			return nil, err
		}

		batchModifyOrdersResponse.Orders = append(batchModifyOrdersResponse.Orders, o)
	}

	return batchModifyOrdersResponse, nil
}
