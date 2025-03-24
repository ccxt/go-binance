package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderCreateWsService creates order
type OrderCreateWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderCreateWsService init OrderCreateWsService
func NewOrderCreateWsService(apiKey, secretKey string) (*OrderCreateWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderCreateWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderCreateWsRequest parameters for 'order.place' websocket API
type OrderCreateWsRequest struct {
	symbol           string
	side             SideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	price            *string
	newClientOrderID *string
	stopPrice        *string
	newOrderRespType NewOrderRespType
	quoteOrderQty    *string
	trailingDelta    *int64
	icebergQty       *string
	strategyId       *uint64
	strategyType     *uint32
	recvWindow       *uint16
}

// NewOrderCreateWsRequest init OrderCreateWsRequest
func NewOrderCreateWsRequest() *OrderCreateWsRequest {
	return &OrderCreateWsRequest{}
}

func (s *OrderCreateWsRequest) GetParams() map[string]interface{} {
	return s.buildParams()
}

// buildParams builds params
func (s *OrderCreateWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.quantity != "" {
		m["quantity"] = s.quantity
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
	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = *s.quoteOrderQty
	}
	if s.trailingDelta != nil {
		m["trailingDelta"] = *s.trailingDelta
	}
	if s.icebergQty != nil {
		m["icebergQty"] = *s.icebergQty
	}
	if s.strategyId != nil {
		m["strategyId"] = *s.strategyId
	}
	if s.strategyType != nil {
		m["strategyType"] = *s.strategyType
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'order.place' request
func (s *OrderCreateWsService) Do(requestID string, request *OrderCreateWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderPlaceSpotWsApiMethod,
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

// SyncDo - sends 'order.place' request and receives response
func (s *OrderCreateWsService) SyncDo(requestID string, request *OrderCreateWsRequest) (*CreateOrderWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderPlaceSpotWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	createOrderWsResponse := &CreateOrderWsResponse{}
	if err := json.Unmarshal(response, createOrderWsResponse); err != nil {
		return nil, err
	}

	return createOrderWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderCreateWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderCreateWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderCreateWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderCreateWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *OrderCreateWsRequest) Symbol(symbol string) *OrderCreateWsRequest {
	s.symbol = symbol
	return s
}

// Side set side
func (s *OrderCreateWsRequest) Side(side SideType) *OrderCreateWsRequest {
	s.side = side
	return s
}

// Type set type
func (s *OrderCreateWsRequest) Type(orderType OrderType) *OrderCreateWsRequest {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *OrderCreateWsRequest) TimeInForce(timeInForce TimeInForceType) *OrderCreateWsRequest {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *OrderCreateWsRequest) Quantity(quantity string) *OrderCreateWsRequest {
	s.quantity = quantity
	return s
}

// Price set price
func (s *OrderCreateWsRequest) Price(price string) *OrderCreateWsRequest {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *OrderCreateWsRequest) NewClientOrderID(newClientOrderID string) *OrderCreateWsRequest {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *OrderCreateWsRequest) StopPrice(stopPrice string) *OrderCreateWsRequest {
	s.stopPrice = &stopPrice
	return s
}

// RecvWindow set recvWindow
func (s *OrderCreateWsRequest) RecvWindow(recvWindow uint16) *OrderCreateWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// StrategyType set strategyType
func (s *OrderCreateWsRequest) StrategyType(strategyType uint32) *OrderCreateWsRequest {
	s.strategyType = &strategyType
	return s
}

// StrategyId set strategyId
func (s *OrderCreateWsRequest) StrategyId(strategyId uint64) *OrderCreateWsRequest {
	s.strategyId = &strategyId
	return s
}

// IcebergQty set icebergQty
func (s *OrderCreateWsRequest) IcebergQty(icebergQty string) *OrderCreateWsRequest {
	s.icebergQty = &icebergQty
	return s
}

// TrailingDelta set trailingDelta
func (s *OrderCreateWsRequest) TrailingDelta(trailingDelta int64) *OrderCreateWsRequest {
	s.trailingDelta = &trailingDelta
	return s
}

// QuoteOrderQty set quoteOrderQty
func (s *OrderCreateWsRequest) QuoteOrderQty(quoteOrderQty string) *OrderCreateWsRequest {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

// NewOrderRespType set newOrderRespType
func (s *OrderCreateWsRequest) NewOrderRespType(newOrderRespType NewOrderRespType) *OrderCreateWsRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

// CreateOrderResult define order creation result
type CreateOrderResult struct {
	CreateOrderResponse
}

// CreateOrderWsResponse define 'order.place' websocket API response
type CreateOrderWsResponse struct {
	Id     string            `json:"id"`
	Status int               `json:"status"`
	Result CreateOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}
