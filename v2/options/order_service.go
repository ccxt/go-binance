package options

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/adshao/go-binance/v2/common"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	orderType        OrderType
	quantity         string
	price            *string
	timeInForce      *TimeInForceType
	reduceOnly       *bool
	postOnly         *bool
	newOrderRespType NewOrderRespType
	clientOrderId    *string
	isMmp            *bool
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
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// PostOnly set postOnly
func (s *CreateOrderService) PostOnly(postOnly bool) *CreateOrderService {
	s.postOnly = &postOnly
	return s
}

// Price set price
func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

// ClientOrderId set clientOrderId
func (s *CreateOrderService) ClientOrderId(ClientOrderId string) *CreateOrderService {
	s.clientOrderId = &ClientOrderId
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// IsMmp set isMmp
func (s *CreateOrderService) IsMmp(isMmp bool) *CreateOrderService {
	s.isMmp = &isMmp
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	if s.newOrderRespType != "" && s.newOrderRespType != NewOrderRespTypeACK &&
		s.newOrderRespType != NewOrderRespTypeRESULT {
		return []byte{}, &http.Header{}, fmt.Errorf("no expected newOrderRespType value=%v", s.newOrderRespType)
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"type":     s.orderType,
		"quantity": s.quantity,
	}
	if s.newOrderRespType != "" {
		m["newOrderRespType"] = s.newOrderRespType
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.postOnly != nil {
		m["postOnly"] = *s.postOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.clientOrderId != nil {
		m["clientOrderId"] = *s.clientOrderId
	}
	if s.isMmp != nil {
		m["isMmp"] = *s.isMmp
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	data, header, err := s.createOrder(ctx, "/eapi/v1/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

type LastTrade struct {
	Id      int64  `json:"id"`
	TradeId int64  `json:"tradeId"`
	Time    int64  `json:"time"`
	Price   string `json:"price"`
	Qty     string `json:"qty"`
}

// Unified order structure, it would be used in many ways, for example, create order, cancel order,
// query open orders, query historical order and so on
type Order struct {
	OrderId       int64           `json:"orderId"`
	Symbol        string          `json:"symbol"`
	Price         string          `json:"price"`
	Quantity      string          `json:"quantity"`
	ExecutedQty   string          `json:"executedQty"`
	Fee           string          `json:"fee"`
	Side          SideType        `json:"side"`
	Type          OrderType       `json:"type"`
	TimeInForce   TimeInForceType `json:"timeInForce"`
	ReduceOnly    bool            `json:"reduceOnly"`
	PostOnly      bool            `json:"postOnly"`
	CreateTime    int64           `json:"createTime"`
	UpdateTime    int64           `json:"updateTime"`
	Status        OrderStatusType `json:"status"`
	Reason        *string         `json:"reason"` // set while query history orders.
	AvgPrice      string          `json:"avgPrice"`
	Source        string          `json:"source"`
	ClientOrderId string          `json:"clientOrderId"`
	PriceScale    int             `json:"priceScale"`
	QuantityScale int             `json:"quantityScale"`
	OptionSide    OptionSideType  `json:"optionSide"`
	QuoteAsset    string          `json:"quoteAsset"`
	Mmp           bool            `json:"mmp"`
	LastTrade     *LastTrade      `json:"lastTrade"` // order is immediately filled while calling create order, it will be set.

	RateLimitOrder10s string `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string `json:"rateLimitOrder1m,omitempty"`
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	c         *Client
	symbol    string
	orderId   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *ListOpenOrdersService) OrderId(orderId int64) *ListOpenOrdersService {
	s.orderId = &orderId
	return s
}

// StartTime set startTime
func (s *ListOpenOrdersService) StartTime(startTime int64) *ListOpenOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListOpenOrdersService) EndTime(endTime int64) *ListOpenOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListOpenOrdersService) Limit(limit int) *ListOpenOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
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

// GetOrderService get an order
type GetOrderService struct {
	c             *Client
	symbol        string
	orderId       *int64
	clientOrderId *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderId set orderId
func (s *GetOrderService) OrderId(orderId int64) *GetOrderService {
	s.orderId = &orderId
	return s
}

// ClientOrderId set clientOrderId
func (s *GetOrderService) ClientOrderId(clientOrderId string) *GetOrderService {
	s.clientOrderId = &clientOrderId
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.clientOrderId != nil {
		r.setParam("clientOrderId", *s.clientOrderId)
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

// CancelOrderService cancel an order
type CancelOrderService struct {
	c             *Client
	symbol        string
	orderId       *int64
	clientOrderId *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderId set OrderId
func (s *CancelOrderService) OrderId(orderId int64) *CancelOrderService {
	s.orderId = &orderId
	return s
}

// ClientOrderId set clientOrderId
func (s *CancelOrderService) ClientOrderId(clientOrderId string) *CancelOrderService {
	s.clientOrderId = &clientOrderId
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderId != nil {
		r.setFormParam("orderId", *s.orderId)
	}
	if s.clientOrderId != nil {
		r.setFormParam("clientOrderId", *s.clientOrderId)
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

// CancelAllOpenOrdersService cancel all open orders
type CancelAllOpenOrdersService struct {
	c      *Client
	symbol string
}

type CancelAllOpenOrdersRsp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Symbol set symbol
func (s *CancelAllOpenOrdersService) Symbol(symbol string) *CancelAllOpenOrdersService {
	s.symbol = symbol
	return s
}

func (s *CancelAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (*CancelAllOpenOrdersRsp, error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	rsp := new(CancelAllOpenOrdersRsp)
	err = json.Unmarshal(data, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// CancelBatchOrdersService cancel a list of orders
type CancelBatchOrdersService struct {
	c              *Client
	symbol         string
	orderIds       []int64
	clientOrderIds []string
}

// Symbol set symbol
func (s *CancelBatchOrdersService) Symbol(symbol string) *CancelBatchOrdersService {
	s.symbol = symbol
	return s
}

// OrderId set OrderId
func (s *CancelBatchOrdersService) OrderIds(orderIds []int64) *CancelBatchOrdersService {
	s.orderIds = orderIds
	return s
}

// ClientOrderIds set clientOrderIds
func (s *CancelBatchOrdersService) ClientOrderIds(clientOrderIds []string) *CancelBatchOrdersService {
	s.clientOrderIds = clientOrderIds
	return s
}

// return: [Order, APIError]
func (s *CancelBatchOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []any, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/batchOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderIds != nil {
		// convert a slice of integers to a string e.g. [1 2 3] => "[1,2,3]"
		orderIDListString := strings.Join(strings.Fields(fmt.Sprint(s.orderIds)), ",")
		r.setFormParam("orderIds", orderIDListString)
	}
	if s.clientOrderIds != nil {
		cids := []string{}
		for _, cid := range s.clientOrderIds {
			cids = append(cids, fmt.Sprintf("\"%v\"", cid))
		}
		r.setFormParam("clientOrderIds", strings.Join(strings.Fields(fmt.Sprint(cids)), ","))
	}
	data, header, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	rlos := header.Get("X-Mbx-Order-Count-10s")
	rlom := header.Get("X-Mbx-Order-Count-1m")

	rawMessages := make([]*json.RawMessage, 0)
	err = json.Unmarshal(data, &rawMessages)
	if err != nil {
		return []any{}, err
	}

	res = make([]any, 0)
	for _, j := range rawMessages {
		e := new(common.APIError)
		if err := json.Unmarshal(*j, e); err != nil {
			return []any{}, err
		}
		if e.IsValid() {
			res = append(res, *e)
			continue
		}
		o := new(Order)
		if err := json.Unmarshal(*j, o); err != nil {
			return []any{}, err
		}
		o.RateLimitOrder10s = rlos
		o.RateLimitOrder1m = rlom
		res = append(res, *o)
	}
	return res, nil
}

type CreateBatchOrdersService struct {
	c      *Client
	orders []*CreateOrderService
}

func (s *CreateBatchOrdersService) OrderList(orders []*CreateOrderService) *CreateBatchOrdersService {
	s.orders = orders
	return s
}

// return: [Order, APIError]
func (s *CreateBatchOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []any, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/eapi/v1/batchOrders",
		secType:  secTypeSigned,
	}

	orders := []params{}
	for _, order := range s.orders {
		if order.newOrderRespType != "" && order.newOrderRespType != NewOrderRespTypeACK &&
			order.newOrderRespType != NewOrderRespTypeRESULT {
			return []any{}, fmt.Errorf("no expected newOrderRespType value=%v", order.newOrderRespType)
		}
		m := params{
			"symbol":   order.symbol,
			"side":     order.side,
			"type":     order.orderType,
			"quantity": order.quantity,
		}
		if order.newOrderRespType != "" {
			m["newOrderRespType"] = order.newOrderRespType
		}
		if order.timeInForce != nil {
			m["timeInForce"] = *order.timeInForce
		}
		if order.reduceOnly != nil {
			m["reduceOnly"] = *order.reduceOnly
		}
		if order.postOnly != nil {
			m["postOnly"] = *order.postOnly
		}
		if order.price != nil {
			m["price"] = *order.price
		}
		if order.clientOrderId != nil {
			m["clientOrderId"] = *order.clientOrderId
		}
		if order.isMmp != nil {
			m["isMmp"] = *order.isMmp
		}
		orders = append(orders, m)
	}

	b, err := json.Marshal(orders)
	if err != nil {
		return []any{}, err
	}
	m := params{
		"orders": string(b),
	}

	r.setFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []any{}, err
	}

	rawMessages := make([]*json.RawMessage, 0)
	err = json.Unmarshal(data, &rawMessages)
	if err != nil {
		return []any{}, err
	}

	res = make([]any, 0)
	for _, j := range rawMessages {
		e := new(common.APIError)
		if err := json.Unmarshal(*j, e); err != nil {
			return []any{}, err
		}
		if e.IsValid() {
			res = append(res, *e)
			continue
		}
		o := new(Order)
		if err := json.Unmarshal(*j, o); err != nil {
			return []any{}, err
		}
		res = append(res, *o)
	}
	return res, nil
}

type CancelAllOpenOrdersByUnderlyingService struct {
	c          *Client
	underlying string
}

type CancelAllOpenOrdersByUnderlyingRsp struct {
	Code string  `json:"code"`
	Msg  string  `json:"msg"`
	Data *string `json:"data"`
}

func (s *CancelAllOpenOrdersByUnderlyingService) Underlying(underlying string) *CancelAllOpenOrdersByUnderlyingService {
	s.underlying = underlying
	return s
}

func (s *CancelAllOpenOrdersByUnderlyingService) Do(ctx context.Context, opts ...RequestOption) (res *CancelAllOpenOrdersByUnderlyingRsp, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/eapi/v1/allOpenOrdersByUnderlying",
		secType:  secTypeSigned,
	}

	m := params{
		"underlying": s.underlying,
	}
	r.setFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CancelAllOpenOrdersByUnderlyingRsp{}, err
	}

	res = new(CancelAllOpenOrdersByUnderlyingRsp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &CancelAllOpenOrdersByUnderlyingRsp{}, err
	}

	return res, nil
}

type HistoryOrdersService struct {
	c         *Client
	symbol    string
	orderId   *uint64
	startTime *uint64
	endTime   *uint64
	limit     *int
}

func (s *HistoryOrdersService) Symbol(symbol string) *HistoryOrdersService {
	s.symbol = symbol
	return s
}

func (s *HistoryOrdersService) OrderId(orderId uint64) *HistoryOrdersService {
	s.orderId = &orderId
	return s
}

func (s *HistoryOrdersService) StartTime(st uint64) *HistoryOrdersService {
	s.startTime = &st
	return s
}

func (s *HistoryOrdersService) EndTime(et uint64) *HistoryOrdersService {
	s.endTime = &et
	return s
}

func (s *HistoryOrdersService) Limit(limit int) *HistoryOrdersService {
	s.limit = &limit
	return s
}

func (s *HistoryOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/historyOrders",
		secType:  secTypeSigned,
	}

	m := params{
		"symbol": s.symbol,
	}
	if s.orderId != nil {
		m["orderId"] = *s.orderId
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)

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

type PositionService struct {
	c      *Client
	symbol *string
}

type Position struct {
	EntryPrice    string `json:"entryPrice"`
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	Quantity      string `json:"quantity"`
	ReducibleQty  string `json:"reducibleQty"`
	MarkValue     string `json:"markValue"`
	Ror           string `json:"ror"`
	UnrealizedPNL string `json:"unrealizedPNL"`
	MarkPrice     string `json:"markPrice"`
	StrikePrice   string `json:"strikePrice"`
	ExpiryDate    uint64 `json:"expiryDate"`
	PositionCost  string `json:"positionCost"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	QuoteAsset    string `json:"quoteAsset"`
	Time          int64  `json:"time"`
}

func (s *PositionService) Symbol(symbol string) *PositionService {
	s.symbol = &symbol
	return s
}

func (s *PositionService) Do(ctx context.Context, opts ...RequestOption) (res []*Position, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/position",
		secType:  secTypeSigned,
	}

	m := params{}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	r.setParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Position{}, err
	}

	res = make([]*Position, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Position{}, err
	}

	return res, nil
}

type UserTradesService struct {
	c         *Client
	symbol    *string
	fromId    *uint64
	startTime *uint64
	endTime   *uint64
	limit     *int
}

type UserTrade struct {
	Id             uint64 `json:"id"`
	TradeId        uint32 `json:"tradeId"`
	OrderId        uint64 `json:"orderId"`
	Symbol         string `json:"symbol"`
	Price          string `json:"price"`
	Quantity       string `json:"quantity"`
	Fee            string `json:"fee"`
	RealizedProfit string `json:"realizedProfit"`
	Side           string `json:"side"`
	Type           string `json:"type"`
	Volatility     string `json:"volatility"`
	Liquidity      string `json:"liquidity"`
	Time           uint64 `json:"time"`
	PriceScale     int    `json:"priceScale"`
	QuantityScale  int    `json:"quantityScale"`
	OptionSide     string `json:"optionSide"`
	QuoteAsset     string `json:"quoteAsset"`
}

func (s *UserTradesService) Symbol(symbol string) *UserTradesService {
	s.symbol = &symbol
	return s
}

func (s *UserTradesService) FromId(fromId uint64) *UserTradesService {
	s.fromId = &fromId
	return s
}

func (s *UserTradesService) StartTime(st uint64) *UserTradesService {
	s.startTime = &st
	return s
}

func (s *UserTradesService) EndTime(et uint64) *UserTradesService {
	s.endTime = &et
	return s
}

func (s *UserTradesService) Limit(limit int) *UserTradesService {
	s.limit = &limit
	return s
}

func (s *UserTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*UserTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/userTrades",
		secType:  secTypeSigned,
	}

	m := params{}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	if s.fromId != nil {
		m["fromId"] = *s.fromId
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*UserTrade{}, err
	}

	res = make([]*UserTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UserTrade{}, err
	}

	return res, nil
}

type ExerciseRecordService struct {
	c         *Client
	symbol    *string
	startTime *uint64
	endTime   *uint64
	limit     *int // default 1000, max 1000
}

type ExerciseRecord struct {
	Id            string `json:"id"`
	Currency      string `json:"currency"`
	Symbol        string `json:"symbol"`
	ExercisePrice string `json:"exercisePrice"`
	MarkPrice     string `json:"markPrice"`
	Quantity      string `json:"quantity"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	CreateDate    uint64 `json:"createDate"`
	PriceScale    int    `json:"priceScale"`
	QuantityScale int    `json:"quantityScale"`
	OptionSide    string `json:"optionSide"`
	PositionSide  string `json:"positionSide"`
	QuoteAsset    string `json:"quoteAsset"`
}

func (s *ExerciseRecordService) Symbol(symbol string) *ExerciseRecordService {
	s.symbol = &symbol
	return s
}

func (s *ExerciseRecordService) StartTime(st uint64) *ExerciseRecordService {
	s.startTime = &st
	return s
}

func (s *ExerciseRecordService) EndTime(et uint64) *ExerciseRecordService {
	s.endTime = &et
	return s
}

func (s *ExerciseRecordService) Limit(limit int) *ExerciseRecordService {
	s.limit = &limit
	return s
}

func (s *ExerciseRecordService) Do(ctx context.Context, opts ...RequestOption) (res []*ExerciseRecord, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/exerciseRecord",
		secType:  secTypeSigned,
	}

	m := params{}
	if s.symbol != nil {
		m["symbol"] = *s.symbol
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*ExerciseRecord{}, err
	}

	res = make([]*ExerciseRecord, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*ExerciseRecord{}, err
	}

	return res, nil
}

type BillService struct {
	c         *Client
	currency  string
	recordId  *uint64
	startTime *uint64 // start timestamp, for example 1593511200000
	endTime   *uint64
	limit     *int // default 100, max 1000
}

type Bill struct {
	Id         string `json:"id"`
	Asset      string `json:"asset"`
	Amount     string `json:"amount"`
	Type       string `json:"type"`
	CreateDate int64  `json:"createDate"`
}

func (s *BillService) Currency(c string) *BillService {
	s.currency = c
	return s
}

func (s *BillService) RecordId(ri uint64) *BillService {
	s.recordId = &ri
	return s
}

func (s *BillService) StartTime(st uint64) *BillService {
	s.startTime = &st
	return s
}

func (s *BillService) EndTime(et uint64) *BillService {
	s.endTime = &et
	return s
}

func (s *BillService) Limit(limit int) *BillService {
	s.limit = &limit
	return s
}

func (s *BillService) Do(ctx context.Context, opts ...RequestOption) (res []*Bill, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/bill",
		secType:  secTypeSigned,
	}

	m := params{
		"currency": s.currency,
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	r.setParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Bill{}, err
	}

	res = make([]*Bill, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Bill{}, err
	}

	return res, nil
}

type IncomeDownloadIdService struct {
	c         *Client
	startTime uint64 // timestamp ms
	endTime   uint64
}

type IncomeDownloadId struct {
	AvgCostTimestampOfLast30d uint64  `json:"avgCostTimestampOfLast30d"`
	DownloadId                string  `json:"downloadId"`
	Error                     *string `json:"error"`
}

func (s *IncomeDownloadIdService) StartTime(st uint64) *IncomeDownloadIdService {
	s.startTime = st
	return s
}

func (s *IncomeDownloadIdService) EndTime(et uint64) *IncomeDownloadIdService {
	s.endTime = et
	return s
}

func (s *IncomeDownloadIdService) Do(ctx context.Context, opts ...RequestOption) (res *IncomeDownloadId, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/income/asyn",
		secType:  secTypeSigned,
	}

	m := params{
		"startTime": s.startTime,
		"endTime":   s.endTime,
	}

	r.setParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IncomeDownloadId)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type IncomeDownloadLinkService struct {
	c          *Client
	downloadId string
}

type IncomeDownloadLink struct {
	DownloadId          string `json:"downloadId"`
	Status              string `json:"status"`
	Url                 string `json:"url"`
	Notified            bool   `json:"notified"`
	ExpirationTimestamp int64  `json:"expirationTimestamp"`
	IsExpired           *bool  `json:"isExpired"`
}

func (s *IncomeDownloadLinkService) DownloadId(di string) *IncomeDownloadLinkService {
	s.downloadId = di
	return s
}

func (s *IncomeDownloadLinkService) Do(ctx context.Context, opts ...RequestOption) (res *IncomeDownloadLink, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/income/id",
		secType:  secTypeSigned,
	}

	m := params{
		"downloadId": s.downloadId,
	}

	r.setParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IncomeDownloadLink)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
