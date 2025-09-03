package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderListPlaceWsService creates order list (deprecated OCO)
type OrderListPlaceWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderListPlaceWsService init OrderListPlaceWsService
func NewOrderListPlaceWsService(apiKey, secretKey string) (*OrderListPlaceWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderListPlaceWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderListPlaceWsRequest parameters for 'orderList.place' websocket API (deprecated OCO)
type OrderListPlaceWsRequest struct {
	symbol                  string
	side                    SideType
	price                   string
	quantity                string
	listClientOrderID       *string
	limitClientOrderID      *string
	limitIcebergQty         *string
	limitStrategyId         *int64
	limitStrategyType       *int32
	stopPrice               *string
	trailingDelta           *int64
	stopClientOrderID       *string
	stopLimitPrice          *string
	stopLimitTimeInForce    *TimeInForceType
	stopIcebergQty          *string
	stopStrategyId          *int64
	stopStrategyType        *int32
	newOrderRespType        NewOrderRespType
	selfTradePreventionMode *SelfTradePreventionMode
	recvWindow              *uint16
}

// NewOrderListPlaceWsRequest init OrderListPlaceWsRequest
func NewOrderListPlaceWsRequest() *OrderListPlaceWsRequest {
	return &OrderListPlaceWsRequest{
		newOrderRespType: NewOrderRespTypeFULL,
	}
}

func (s *OrderListPlaceWsRequest) GetParams() map[string]interface{} {
	return s.buildParams()
}

// buildParams builds params
func (s *OrderListPlaceWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"price":            s.price,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.listClientOrderID != nil {
		m["listClientOrderId"] = *s.listClientOrderID
	} else {
		m["listClientOrderId"] = common.GenerateSpotId()
	}
	if s.limitClientOrderID != nil {
		m["limitClientOrderId"] = *s.limitClientOrderID
	} else {
		m["limitClientOrderId"] = common.GenerateSpotId()
	}
	if s.limitIcebergQty != nil {
		m["limitIcebergQty"] = *s.limitIcebergQty
	}
	if s.limitStrategyId != nil {
		m["limitStrategyId"] = *s.limitStrategyId
	}
	if s.limitStrategyType != nil {
		m["limitStrategyType"] = *s.limitStrategyType
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.trailingDelta != nil {
		m["trailingDelta"] = *s.trailingDelta
	}
	if s.stopClientOrderID != nil {
		m["stopClientOrderId"] = *s.stopClientOrderID
	} else {
		m["stopClientOrderId"] = common.GenerateSpotId()
	}
	if s.stopLimitPrice != nil {
		m["stopLimitPrice"] = *s.stopLimitPrice
	}
	if s.stopLimitTimeInForce != nil {
		m["stopLimitTimeInForce"] = *s.stopLimitTimeInForce
	}
	if s.stopIcebergQty != nil {
		m["stopIcebergQty"] = *s.stopIcebergQty
	}
	if s.stopStrategyId != nil {
		m["stopStrategyId"] = *s.stopStrategyId
	}
	if s.stopStrategyType != nil {
		m["stopStrategyType"] = *s.stopStrategyType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'orderList.place' request
func (s *OrderListPlaceWsService) Do(requestID string, request *OrderListPlaceWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceSpotWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

// SyncDo - sends 'orderList.place' request and receives response
func (s *OrderListPlaceWsService) SyncDo(requestID string, request *OrderListPlaceWsRequest) (*CreateOrderListWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceSpotWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	createOrderListWsResponse := &CreateOrderListWsResponse{}
	if err := json.Unmarshal(response, createOrderListWsResponse); err != nil {
		return nil, err
	}

	return createOrderListWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderListPlaceWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderListPlaceWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderListPlaceWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderListPlaceWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *OrderListPlaceWsRequest) Symbol(symbol string) *OrderListPlaceWsRequest {
	s.symbol = symbol
	return s
}

// Side set side
func (s *OrderListPlaceWsRequest) Side(side SideType) *OrderListPlaceWsRequest {
	s.side = side
	return s
}

// Price set price
func (s *OrderListPlaceWsRequest) Price(price string) *OrderListPlaceWsRequest {
	s.price = price
	return s
}

// Quantity set quantity
func (s *OrderListPlaceWsRequest) Quantity(quantity string) *OrderListPlaceWsRequest {
	s.quantity = quantity
	return s
}

// ListClientOrderID set listClientOrderID
func (s *OrderListPlaceWsRequest) ListClientOrderID(listClientOrderID string) *OrderListPlaceWsRequest {
	s.listClientOrderID = &listClientOrderID
	return s
}

// LimitClientOrderID set limitClientOrderID
func (s *OrderListPlaceWsRequest) LimitClientOrderID(limitClientOrderID string) *OrderListPlaceWsRequest {
	s.limitClientOrderID = &limitClientOrderID
	return s
}

// LimitIcebergQty set limitIcebergQty
func (s *OrderListPlaceWsRequest) LimitIcebergQty(limitIcebergQty string) *OrderListPlaceWsRequest {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

// LimitStrategyId set limitStrategyId
func (s *OrderListPlaceWsRequest) LimitStrategyId(limitStrategyId int64) *OrderListPlaceWsRequest {
	s.limitStrategyId = &limitStrategyId
	return s
}

// LimitStrategyType set limitStrategyType
func (s *OrderListPlaceWsRequest) LimitStrategyType(limitStrategyType int32) *OrderListPlaceWsRequest {
	s.limitStrategyType = &limitStrategyType
	return s
}

// StopPrice set stopPrice
func (s *OrderListPlaceWsRequest) StopPrice(stopPrice string) *OrderListPlaceWsRequest {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta set trailingDelta
func (s *OrderListPlaceWsRequest) TrailingDelta(trailingDelta int64) *OrderListPlaceWsRequest {
	s.trailingDelta = &trailingDelta
	return s
}

// StopClientOrderID set stopClientOrderID
func (s *OrderListPlaceWsRequest) StopClientOrderID(stopClientOrderID string) *OrderListPlaceWsRequest {
	s.stopClientOrderID = &stopClientOrderID
	return s
}

// StopLimitPrice set stopLimitPrice
func (s *OrderListPlaceWsRequest) StopLimitPrice(stopLimitPrice string) *OrderListPlaceWsRequest {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

// StopLimitTimeInForce set stopLimitTimeInForce
func (s *OrderListPlaceWsRequest) StopLimitTimeInForce(stopLimitTimeInForce TimeInForceType) *OrderListPlaceWsRequest {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

// StopIcebergQty set stopIcebergQty
func (s *OrderListPlaceWsRequest) StopIcebergQty(stopIcebergQty string) *OrderListPlaceWsRequest {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

// StopStrategyId set stopStrategyId
func (s *OrderListPlaceWsRequest) StopStrategyId(stopStrategyId int64) *OrderListPlaceWsRequest {
	s.stopStrategyId = &stopStrategyId
	return s
}

// StopStrategyType set stopStrategyType
func (s *OrderListPlaceWsRequest) StopStrategyType(stopStrategyType int32) *OrderListPlaceWsRequest {
	s.stopStrategyType = &stopStrategyType
	return s
}

// NewOrderRespType set newOrderRespType
func (s *OrderListPlaceWsRequest) NewOrderRespType(newOrderRespType NewOrderRespType) *OrderListPlaceWsRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *OrderListPlaceWsRequest) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *OrderListPlaceWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// RecvWindow set recvWindow
func (s *OrderListPlaceWsRequest) RecvWindow(recvWindow uint16) *OrderListPlaceWsRequest {
	s.recvWindow = &recvWindow
	return s
}
