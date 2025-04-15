package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adshao/go-binance/v2/futures"
)

// CreateFuturesAlgoTwapOrderService create future algo order
type CreateFuturesAlgoVpOrderService struct {
	c            *Client
	symbol       string
	side         SideType
	positionSide *futures.PositionSideType
	quantity     float64
	urgency      FuturesAlgoUrgencyType
	clientAlgoId *string
	reduceOnly   *bool
	limitPrice   *float64
}

// Symbol set symbol
func (s *CreateFuturesAlgoVpOrderService) Symbol(symbol string) *CreateFuturesAlgoVpOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateFuturesAlgoVpOrderService) Side(side SideType) *CreateFuturesAlgoVpOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateFuturesAlgoVpOrderService) PositionSide(positionSide futures.PositionSideType) *CreateFuturesAlgoVpOrderService {
	s.positionSide = &positionSide
	return s
}

// Quantity set quantity
func (s *CreateFuturesAlgoVpOrderService) Quantity(quantity float64) *CreateFuturesAlgoVpOrderService {
	s.quantity = quantity
	return s
}

// Urgency set urgency
func (s *CreateFuturesAlgoVpOrderService) Urgency(urgency FuturesAlgoUrgencyType) *CreateFuturesAlgoVpOrderService {
	s.urgency = urgency
	return s
}

// ClientAlgoId set clientAlgoId
func (s *CreateFuturesAlgoVpOrderService) ClientAlgoId(clientAlgoId string) *CreateFuturesAlgoVpOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateFuturesAlgoVpOrderService) ReduceOnly(reduceOnly bool) *CreateFuturesAlgoVpOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// LimitPrice set limitPrice
func (s *CreateFuturesAlgoVpOrderService) LimitPrice(limitPrice float64) *CreateFuturesAlgoVpOrderService {
	s.limitPrice = &limitPrice
	return s
}

func (s *CreateFuturesAlgoVpOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
		"urgency":  s.urgency,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.clientAlgoId != nil {
		m["clientAlgoId"] = *s.clientAlgoId
	}
	if s.limitPrice != nil {
		m["limitPrice"] = *s.limitPrice
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateFuturesAlgoVpOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateFuturesAlgoOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/sapi/v1/algo/futures/newOrderVp", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateFuturesAlgoOrderResponse)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateFutureAlgoOrderResponse define create future algo order response
type CreateFuturesAlgoOrderResponse struct {
	ClientAlgoId string `json:"clientAlgoId"`
	Success      bool   `json:"success"`
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
}

// CreateFuturesAlgoTwapOrderService create future algo order
type CreateFuturesAlgoTwapOrderService struct {
	c            *Client
	symbol       string
	side         SideType
	positionSide *futures.PositionSideType
	quantity     float64
	duration     int64
	clientAlgoId *string
	reduceOnly   *bool
	limitPrice   *float64
}

// Symbol set symbol
func (s *CreateFuturesAlgoTwapOrderService) Symbol(symbol string) *CreateFuturesAlgoTwapOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateFuturesAlgoTwapOrderService) Side(side SideType) *CreateFuturesAlgoTwapOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateFuturesAlgoTwapOrderService) PositionSide(positionSide futures.PositionSideType) *CreateFuturesAlgoTwapOrderService {
	s.positionSide = &positionSide
	return s
}

// Quantity set quantity
func (s *CreateFuturesAlgoTwapOrderService) Quantity(quantity float64) *CreateFuturesAlgoTwapOrderService {
	s.quantity = quantity
	return s
}

// Duration set duration
func (s *CreateFuturesAlgoTwapOrderService) Duration(duration int64) *CreateFuturesAlgoTwapOrderService {
	s.duration = duration
	return s
}

// ClientAlgoId set clientAlgoId
func (s *CreateFuturesAlgoTwapOrderService) ClientAlgoId(clientAlgoId string) *CreateFuturesAlgoTwapOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateFuturesAlgoTwapOrderService) ReduceOnly(reduceOnly bool) *CreateFuturesAlgoTwapOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// LimitPrice set limitPrice
func (s *CreateFuturesAlgoTwapOrderService) LimitPrice(limitPrice float64) *CreateFuturesAlgoTwapOrderService {
	s.limitPrice = &limitPrice
	return s
}

func (s *CreateFuturesAlgoTwapOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
		"duration": s.duration,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.clientAlgoId != nil {
		m["clientAlgoId"] = *s.clientAlgoId
	}
	if s.limitPrice != nil {
		m["limitPrice"] = *s.limitPrice
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateFuturesAlgoTwapOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateFuturesAlgoOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/sapi/v1/algo/futures/newOrderTwap", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateFuturesAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	// why do we print response here
	fmt.Printf("response is %v", data)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListOpenFuturesAlgoOrdersService list current open futures algo orders
type ListOpenFuturesAlgoOrdersService struct {
	c *Client
}

// Do send request
func (s *ListOpenFuturesAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *ListOpenFuturesAlgoOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/openOrders",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ListOpenFuturesAlgoOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FuturesAlgoOrder struct {
	//策略订单ID
	AlgoId           int64                      `json:"algoId"`
	Symbol           string                     `json:"symbol"`
	Side             SideType                   `json:"side"`
	PositionSide     futures.PositionSideType   `json:"positionSide"`
	TotalQuantity    string                     `json:"totalQty"`
	ExecutedQuantity string                     `json:"executedQty"`
	ExecutedAmount   string                     `json:"executedAmt"`
	AvgPrice         string                     `json:"avgPrice"`
	ClientAlgoId     string                     `json:"clientAlgoId"`
	BookTime         int64                      `json:"bookTime"`
	EndTime          int64                      `json:"endTime"`
	AlgoStatus       FuturesAlgoOrderStatusType `json:"algoStatus"`
	AlgoType         FuturesAlgoType            `json:"algoType"`
	Urgency          FuturesAlgoUrgencyType     `json:"urgency"`
}

// ListOpenFuturesAlgoOrdersResponse define response of list open future algo orders
type ListOpenFuturesAlgoOrdersResponse struct {
	Total  int64               `json:"total"`
	Orders []*FuturesAlgoOrder `json:"orders"`
}

// ListHistoryFuturesAlgoOrdersService list future algo historical orders
type ListHistoryFuturesAlgoOrdersService struct {
	c         *Client
	symbol    *string
	side      *SideType
	startTime *int64
	endTime   *int64
	page      *int
	pageSize  *int
}

// Symbol set symbol
func (s *ListHistoryFuturesAlgoOrdersService) Symbol(symbol string) *ListHistoryFuturesAlgoOrdersService {
	s.symbol = &symbol
	return s
}

// Side set side
func (s *ListHistoryFuturesAlgoOrdersService) Side(side SideType) *ListHistoryFuturesAlgoOrdersService {
	s.side = &side
	return s
}

// StartTime set startTime
func (s *ListHistoryFuturesAlgoOrdersService) StartTime(startTime int64) *ListHistoryFuturesAlgoOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListHistoryFuturesAlgoOrdersService) EndTime(endTime int64) *ListHistoryFuturesAlgoOrdersService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *ListHistoryFuturesAlgoOrdersService) Page(page int) *ListHistoryFuturesAlgoOrdersService {
	s.page = &page
	return s
}

// PageSize set pageSize
func (s *ListHistoryFuturesAlgoOrdersService) PageSize(pageSize int) *ListHistoryFuturesAlgoOrdersService {
	s.pageSize = &pageSize
	return s
}

// Do send request
func (s *ListHistoryFuturesAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *ListHistoryFuturesAlgoOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/historicalOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.side != nil {
		r.setParam("side", *s.side)
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
	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ListHistoryFuturesAlgoOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListFutureAlgoOrderHistoryResponse defines response of list future algo historical orders
type ListHistoryFuturesAlgoOrdersResponse struct {
	Total  int64               `json:"total"`
	Orders []*FuturesAlgoOrder `json:"orders"`
}

// CancelFuturesAlgoOrderService cancel future algo twap order
type CancelFuturesAlgoOrderService struct {
	c      *Client
	algoId int64
}

func (s *CancelFuturesAlgoOrderService) AlgoId(algoId int64) *CancelFuturesAlgoOrderService {
	s.algoId = algoId
	return s
}

// Do send request
func (s *CancelFuturesAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelFuturesAlgoOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/algo/futures/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("algoId", s.algoId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelFuturesAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelFuturesAlgoOrderResponse define response of cancel future algo twap order
type CancelFuturesAlgoOrderResponse struct {
	AlgoId  int64  `json:"algoId"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

// GetFuturesAlgoSubOrdersService get future algo sub orders
type GetFuturesAlgoSubOrdersService struct {
	c        *Client
	algoId   int64
	page     *int
	pageSize *int
}

// AlgoId set algoId
func (s *GetFuturesAlgoSubOrdersService) AlgoId(algoId int64) *GetFuturesAlgoSubOrdersService {
	s.algoId = algoId
	return s
}

// Page set page
func (s *GetFuturesAlgoSubOrdersService) Page(page int) *GetFuturesAlgoSubOrdersService {
	s.page = &page
	return s
}

// PageSize set pageSize
func (s *GetFuturesAlgoSubOrdersService) PageSize(pageSize int) *GetFuturesAlgoSubOrdersService {
	s.pageSize = &pageSize
	return s
}

// Do send request
func (s *GetFuturesAlgoSubOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *GetFuturesAlgoSubOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/subOrders",
		secType:  secTypeSigned,
	}
	r.setParam("algoId", s.algoId)
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetFuturesAlgoSubOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FutureAlgoSubOrder definen sub order of future algo order
type FuturesAlgoSubOrder struct {
	AlgoId           int64           `json:"algoId"`
	OrderId          int64           `json:"orderId"`
	Symbol           string          `json:"symbol"`
	Side             SideType        `json:"side"`
	OrderStatus      OrderStatusType `json:"orderStatus"`
	ExecutedQuantity string          `json:"executedQty"`
	ExecutedAmount   string          `json:"executedAmt"`
	FeeAmount        string          `json:"feeAmt"`
	FeeAsset         string          `json:"feeAsset"`
	AvgPrice         string          `json:"avgPrice"`
	BookTime         int64           `json:"bookTime"`
	SubId            int64           `json:"subId"`
	TimeInForce      TimeInForceType `json:"timeInForce"`
	OriginQuantity   string          `json:"origQty"`
}

type GetFuturesAlgoSubOrdersResponse struct {
	Total            int64                  `json:"total"`
	ExecutedQuantity string                 `json:"executedQty"`
	ExecutedAmount   string                 `json:"executedAmt"`
	SubOrders        []*FuturesAlgoSubOrder `json:"subOrders"`
}
