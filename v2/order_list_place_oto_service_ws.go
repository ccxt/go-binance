package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderListPlaceOtoWsService creates OTO order list
type OrderListPlaceOtoWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderListPlaceOtoWsService init OrderListPlaceOtoWsService
func NewOrderListPlaceOtoWsService(apiKey, secretKey string) (*OrderListPlaceOtoWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderListPlaceOtoWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderListPlaceOtoWsRequest parameters for 'orderList.place.oto' websocket API
type OrderListPlaceOtoWsRequest struct {
	symbol                  string
	listClientOrderID       *string
	newOrderRespType        NewOrderRespType
	selfTradePreventionMode *SelfTradePreventionMode
	workingType             OrderType
	workingSide             SideType
	workingClientOrderID    *string
	workingPrice            string
	workingQuantity         string
	workingIcebergQty       *string
	workingTimeInForce      *TimeInForceType
	workingStrategyId       *int64
	workingStrategyType     *int32
	pendingType             OrderType
	pendingSide             SideType
	pendingClientOrderID    *string
	pendingPrice            *string
	pendingStopPrice        *string
	pendingTrailingDelta    *int64
	pendingQuantity         string
	pendingIcebergQty       *string
	pendingTimeInForce      *TimeInForceType
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	recvWindow              *uint16
}

// NewOrderListPlaceOtoWsRequest init OrderListPlaceOtoWsRequest
func NewOrderListPlaceOtoWsRequest() *OrderListPlaceOtoWsRequest {
	return &OrderListPlaceOtoWsRequest{
		newOrderRespType: NewOrderRespTypeFULL,
	}
}

func (s *OrderListPlaceOtoWsRequest) GetParams() map[string]interface{} {
	return s.buildParams()
}

// buildParams builds params
func (s *OrderListPlaceOtoWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"newOrderRespType": s.newOrderRespType,
		"workingType":      s.workingType,
		"workingSide":      s.workingSide,
		"workingPrice":     s.workingPrice,
		"workingQuantity":  s.workingQuantity,
		"pendingType":      s.pendingType,
		"pendingSide":      s.pendingSide,
		"pendingQuantity":  s.pendingQuantity,
	}
	if s.listClientOrderID != nil {
		m["listClientOrderId"] = *s.listClientOrderID
	} else {
		m["listClientOrderId"] = common.GenerateSpotId()
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.workingClientOrderID != nil {
		m["workingClientOrderId"] = *s.workingClientOrderID
	} else {
		m["workingClientOrderId"] = common.GenerateSpotId()
	}
	if s.workingIcebergQty != nil {
		m["workingIcebergQty"] = *s.workingIcebergQty
	}
	if s.workingTimeInForce != nil {
		m["workingTimeInForce"] = *s.workingTimeInForce
	}
	if s.workingStrategyId != nil {
		m["workingStrategyId"] = *s.workingStrategyId
	}
	if s.workingStrategyType != nil {
		m["workingStrategyType"] = *s.workingStrategyType
	}
	if s.pendingClientOrderID != nil {
		m["pendingClientOrderId"] = *s.pendingClientOrderID
	} else {
		m["pendingClientOrderId"] = common.GenerateSpotId()
	}
	if s.pendingPrice != nil {
		m["pendingPrice"] = *s.pendingPrice
	}
	if s.pendingStopPrice != nil {
		m["pendingStopPrice"] = *s.pendingStopPrice
	}
	if s.pendingTrailingDelta != nil {
		m["pendingTrailingDelta"] = *s.pendingTrailingDelta
	}
	if s.pendingIcebergQty != nil {
		m["pendingIcebergQty"] = *s.pendingIcebergQty
	}
	if s.pendingTimeInForce != nil {
		m["pendingTimeInForce"] = *s.pendingTimeInForce
	}
	if s.pendingStrategyId != nil {
		m["pendingStrategyId"] = *s.pendingStrategyId
	}
	if s.pendingStrategyType != nil {
		m["pendingStrategyType"] = *s.pendingStrategyType
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'orderList.place.oto' request
func (s *OrderListPlaceOtoWsService) Do(requestID string, request *OrderListPlaceOtoWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceOtoSpotWsApiMethod,
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

// SyncDo - sends 'orderList.place.oto' request and receives response
func (s *OrderListPlaceOtoWsService) SyncDo(requestID string, request *OrderListPlaceOtoWsRequest) (*CreateOrderListWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceOtoSpotWsApiMethod,
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
func (s *OrderListPlaceOtoWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderListPlaceOtoWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderListPlaceOtoWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderListPlaceOtoWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *OrderListPlaceOtoWsRequest) Symbol(symbol string) *OrderListPlaceOtoWsRequest {
	s.symbol = symbol
	return s
}

// ListClientOrderID set listClientOrderID
func (s *OrderListPlaceOtoWsRequest) ListClientOrderID(listClientOrderID string) *OrderListPlaceOtoWsRequest {
	s.listClientOrderID = &listClientOrderID
	return s
}

// NewOrderRespType set newOrderRespType
func (s *OrderListPlaceOtoWsRequest) NewOrderRespType(newOrderRespType NewOrderRespType) *OrderListPlaceOtoWsRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *OrderListPlaceOtoWsRequest) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *OrderListPlaceOtoWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// WorkingType set workingType
func (s *OrderListPlaceOtoWsRequest) WorkingType(workingType OrderType) *OrderListPlaceOtoWsRequest {
	s.workingType = workingType
	return s
}

// WorkingSide set workingSide
func (s *OrderListPlaceOtoWsRequest) WorkingSide(workingSide SideType) *OrderListPlaceOtoWsRequest {
	s.workingSide = workingSide
	return s
}

// WorkingClientOrderID set workingClientOrderID
func (s *OrderListPlaceOtoWsRequest) WorkingClientOrderID(workingClientOrderID string) *OrderListPlaceOtoWsRequest {
	s.workingClientOrderID = &workingClientOrderID
	return s
}

// WorkingPrice set workingPrice
func (s *OrderListPlaceOtoWsRequest) WorkingPrice(workingPrice string) *OrderListPlaceOtoWsRequest {
	s.workingPrice = workingPrice
	return s
}

// WorkingQuantity set workingQuantity
func (s *OrderListPlaceOtoWsRequest) WorkingQuantity(workingQuantity string) *OrderListPlaceOtoWsRequest {
	s.workingQuantity = workingQuantity
	return s
}

// WorkingIcebergQty set workingIcebergQty
func (s *OrderListPlaceOtoWsRequest) WorkingIcebergQty(workingIcebergQty string) *OrderListPlaceOtoWsRequest {
	s.workingIcebergQty = &workingIcebergQty
	return s
}

// WorkingTimeInForce set workingTimeInForce
func (s *OrderListPlaceOtoWsRequest) WorkingTimeInForce(workingTimeInForce TimeInForceType) *OrderListPlaceOtoWsRequest {
	s.workingTimeInForce = &workingTimeInForce
	return s
}

// WorkingStrategyId set workingStrategyId
func (s *OrderListPlaceOtoWsRequest) WorkingStrategyId(workingStrategyId int64) *OrderListPlaceOtoWsRequest {
	s.workingStrategyId = &workingStrategyId
	return s
}

// WorkingStrategyType set workingStrategyType
func (s *OrderListPlaceOtoWsRequest) WorkingStrategyType(workingStrategyType int32) *OrderListPlaceOtoWsRequest {
	s.workingStrategyType = &workingStrategyType
	return s
}

// PendingType set pendingType
func (s *OrderListPlaceOtoWsRequest) PendingType(pendingType OrderType) *OrderListPlaceOtoWsRequest {
	s.pendingType = pendingType
	return s
}

// PendingSide set pendingSide
func (s *OrderListPlaceOtoWsRequest) PendingSide(pendingSide SideType) *OrderListPlaceOtoWsRequest {
	s.pendingSide = pendingSide
	return s
}

// PendingClientOrderID set pendingClientOrderID
func (s *OrderListPlaceOtoWsRequest) PendingClientOrderID(pendingClientOrderID string) *OrderListPlaceOtoWsRequest {
	s.pendingClientOrderID = &pendingClientOrderID
	return s
}

// PendingPrice set pendingPrice
func (s *OrderListPlaceOtoWsRequest) PendingPrice(pendingPrice string) *OrderListPlaceOtoWsRequest {
	s.pendingPrice = &pendingPrice
	return s
}

// PendingStopPrice set pendingStopPrice
func (s *OrderListPlaceOtoWsRequest) PendingStopPrice(pendingStopPrice string) *OrderListPlaceOtoWsRequest {
	s.pendingStopPrice = &pendingStopPrice
	return s
}

// PendingTrailingDelta set pendingTrailingDelta
func (s *OrderListPlaceOtoWsRequest) PendingTrailingDelta(pendingTrailingDelta int64) *OrderListPlaceOtoWsRequest {
	s.pendingTrailingDelta = &pendingTrailingDelta
	return s
}

// PendingQuantity set pendingQuantity
func (s *OrderListPlaceOtoWsRequest) PendingQuantity(pendingQuantity string) *OrderListPlaceOtoWsRequest {
	s.pendingQuantity = pendingQuantity
	return s
}

// PendingIcebergQty set pendingIcebergQty
func (s *OrderListPlaceOtoWsRequest) PendingIcebergQty(pendingIcebergQty string) *OrderListPlaceOtoWsRequest {
	s.pendingIcebergQty = &pendingIcebergQty
	return s
}

// PendingTimeInForce set pendingTimeInForce
func (s *OrderListPlaceOtoWsRequest) PendingTimeInForce(pendingTimeInForce TimeInForceType) *OrderListPlaceOtoWsRequest {
	s.pendingTimeInForce = &pendingTimeInForce
	return s
}

// PendingStrategyId set pendingStrategyId
func (s *OrderListPlaceOtoWsRequest) PendingStrategyId(pendingStrategyId int64) *OrderListPlaceOtoWsRequest {
	s.pendingStrategyId = &pendingStrategyId
	return s
}

// PendingStrategyType set pendingStrategyType
func (s *OrderListPlaceOtoWsRequest) PendingStrategyType(pendingStrategyType int32) *OrderListPlaceOtoWsRequest {
	s.pendingStrategyType = &pendingStrategyType
	return s
}

// RecvWindow set recvWindow
func (s *OrderListPlaceOtoWsRequest) RecvWindow(recvWindow uint16) *OrderListPlaceOtoWsRequest {
	s.recvWindow = &recvWindow
	return s
}
