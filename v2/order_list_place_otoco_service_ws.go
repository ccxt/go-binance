package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderListPlaceOtocoWsService creates OTOCO order list
type OrderListPlaceOtocoWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderListPlaceOtocoWsService init OrderListPlaceOtocoWsService
func NewOrderListPlaceOtocoWsService(apiKey, secretKey string) (*OrderListPlaceOtocoWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderListPlaceOtocoWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderListPlaceOtocoWsRequest parameters for 'orderList.place.otoco' websocket API
type OrderListPlaceOtocoWsRequest struct {
	symbol                    string
	listClientOrderID         *string
	newOrderRespType          NewOrderRespType
	selfTradePreventionMode   *SelfTradePreventionMode
	workingType               OrderType
	workingSide               SideType
	workingClientOrderID      *string
	workingPrice              string
	workingQuantity           string
	workingIcebergQty         *string
	workingTimeInForce        *TimeInForceType
	workingStrategyId         *int64
	workingStrategyType       *int32
	pendingSide               SideType
	pendingQuantity           string
	pendingAboveType          OrderType
	pendingAboveClientOrderID *string
	pendingAbovePrice         *string
	pendingAboveStopPrice     *string
	pendingAboveTrailingDelta *int64
	pendingAboveIcebergQty    *string
	pendingAboveTimeInForce   *TimeInForceType
	pendingAboveStrategyId    *int64
	pendingAboveStrategyType  *int32
	pendingBelowType          *OrderType
	pendingBelowClientOrderID *string
	pendingBelowPrice         *string
	pendingBelowStopPrice     *string
	pendingBelowTrailingDelta *int64
	pendingBelowIcebergQty    *string
	pendingBelowTimeInForce   *TimeInForceType
	pendingBelowStrategyId    *int64
	pendingBelowStrategyType  *int32
	recvWindow                *uint16
}

// NewOrderListPlaceOtocoWsRequest init OrderListPlaceOtocoWsRequest
func NewOrderListPlaceOtocoWsRequest() *OrderListPlaceOtocoWsRequest {
	return &OrderListPlaceOtocoWsRequest{
		newOrderRespType: NewOrderRespTypeFULL,
	}
}

func (s *OrderListPlaceOtocoWsRequest) GetParams() map[string]any {
	return s.buildParams()
}

// buildParams builds params
func (s *OrderListPlaceOtocoWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"newOrderRespType": s.newOrderRespType,
		"workingType":      s.workingType,
		"workingSide":      s.workingSide,
		"workingPrice":     s.workingPrice,
		"workingQuantity":  s.workingQuantity,
		"pendingSide":      s.pendingSide,
		"pendingQuantity":  s.pendingQuantity,
		"pendingAboveType": s.pendingAboveType,
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
	if s.pendingAboveClientOrderID != nil {
		m["pendingAboveClientOrderId"] = *s.pendingAboveClientOrderID
	} else {
		m["pendingAboveClientOrderId"] = common.GenerateSpotId()
	}
	if s.pendingAbovePrice != nil {
		m["pendingAbovePrice"] = *s.pendingAbovePrice
	}
	if s.pendingAboveStopPrice != nil {
		m["pendingAboveStopPrice"] = *s.pendingAboveStopPrice
	}
	if s.pendingAboveTrailingDelta != nil {
		m["pendingAboveTrailingDelta"] = *s.pendingAboveTrailingDelta
	}
	if s.pendingAboveIcebergQty != nil {
		m["pendingAboveIcebergQty"] = *s.pendingAboveIcebergQty
	}
	if s.pendingAboveTimeInForce != nil {
		m["pendingAboveTimeInForce"] = *s.pendingAboveTimeInForce
	}
	if s.pendingAboveStrategyId != nil {
		m["pendingAboveStrategyId"] = *s.pendingAboveStrategyId
	}
	if s.pendingAboveStrategyType != nil {
		m["pendingAboveStrategyType"] = *s.pendingAboveStrategyType
	}
	if s.pendingBelowType != nil {
		m["pendingBelowType"] = *s.pendingBelowType
	}
	if s.pendingBelowClientOrderID != nil {
		m["pendingBelowClientOrderId"] = *s.pendingBelowClientOrderID
	} else if s.pendingBelowType != nil {
		m["pendingBelowClientOrderId"] = common.GenerateSpotId()
	}
	if s.pendingBelowPrice != nil {
		m["pendingBelowPrice"] = *s.pendingBelowPrice
	}
	if s.pendingBelowStopPrice != nil {
		m["pendingBelowStopPrice"] = *s.pendingBelowStopPrice
	}
	if s.pendingBelowTrailingDelta != nil {
		m["pendingBelowTrailingDelta"] = *s.pendingBelowTrailingDelta
	}
	if s.pendingBelowIcebergQty != nil {
		m["pendingBelowIcebergQty"] = *s.pendingBelowIcebergQty
	}
	if s.pendingBelowTimeInForce != nil {
		m["pendingBelowTimeInForce"] = *s.pendingBelowTimeInForce
	}
	if s.pendingBelowStrategyId != nil {
		m["pendingBelowStrategyId"] = *s.pendingBelowStrategyId
	}
	if s.pendingBelowStrategyType != nil {
		m["pendingBelowStrategyType"] = *s.pendingBelowStrategyType
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'orderList.place.otoco' request
func (s *OrderListPlaceOtocoWsService) Do(requestID string, request *OrderListPlaceOtocoWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceOtocoSpotWsApiMethod,
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

// SyncDo - sends 'orderList.place.otoco' request and receives response
func (s *OrderListPlaceOtocoWsService) SyncDo(requestID string, request *OrderListPlaceOtocoWsRequest) (*CreateOrderListWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceOtocoSpotWsApiMethod,
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
func (s *OrderListPlaceOtocoWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderListPlaceOtocoWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderListPlaceOtocoWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderListPlaceOtocoWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *OrderListPlaceOtocoWsRequest) Symbol(symbol string) *OrderListPlaceOtocoWsRequest {
	s.symbol = symbol
	return s
}

// ListClientOrderID set listClientOrderID
func (s *OrderListPlaceOtocoWsRequest) ListClientOrderID(listClientOrderID string) *OrderListPlaceOtocoWsRequest {
	s.listClientOrderID = &listClientOrderID
	return s
}

// NewOrderRespType set newOrderRespType
func (s *OrderListPlaceOtocoWsRequest) NewOrderRespType(newOrderRespType NewOrderRespType) *OrderListPlaceOtocoWsRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *OrderListPlaceOtocoWsRequest) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *OrderListPlaceOtocoWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// WorkingType set workingType
func (s *OrderListPlaceOtocoWsRequest) WorkingType(workingType OrderType) *OrderListPlaceOtocoWsRequest {
	s.workingType = workingType
	return s
}

// WorkingSide set workingSide
func (s *OrderListPlaceOtocoWsRequest) WorkingSide(workingSide SideType) *OrderListPlaceOtocoWsRequest {
	s.workingSide = workingSide
	return s
}

// WorkingClientOrderID set workingClientOrderID
func (s *OrderListPlaceOtocoWsRequest) WorkingClientOrderID(workingClientOrderID string) *OrderListPlaceOtocoWsRequest {
	s.workingClientOrderID = &workingClientOrderID
	return s
}

// WorkingPrice set workingPrice
func (s *OrderListPlaceOtocoWsRequest) WorkingPrice(workingPrice string) *OrderListPlaceOtocoWsRequest {
	s.workingPrice = workingPrice
	return s
}

// WorkingQuantity set workingQuantity
func (s *OrderListPlaceOtocoWsRequest) WorkingQuantity(workingQuantity string) *OrderListPlaceOtocoWsRequest {
	s.workingQuantity = workingQuantity
	return s
}

// WorkingIcebergQty set workingIcebergQty
func (s *OrderListPlaceOtocoWsRequest) WorkingIcebergQty(workingIcebergQty string) *OrderListPlaceOtocoWsRequest {
	s.workingIcebergQty = &workingIcebergQty
	return s
}

// WorkingTimeInForce set workingTimeInForce
func (s *OrderListPlaceOtocoWsRequest) WorkingTimeInForce(workingTimeInForce TimeInForceType) *OrderListPlaceOtocoWsRequest {
	s.workingTimeInForce = &workingTimeInForce
	return s
}

// WorkingStrategyId set workingStrategyId
func (s *OrderListPlaceOtocoWsRequest) WorkingStrategyId(workingStrategyId int64) *OrderListPlaceOtocoWsRequest {
	s.workingStrategyId = &workingStrategyId
	return s
}

// WorkingStrategyType set workingStrategyType
func (s *OrderListPlaceOtocoWsRequest) WorkingStrategyType(workingStrategyType int32) *OrderListPlaceOtocoWsRequest {
	s.workingStrategyType = &workingStrategyType
	return s
}

// PendingSide set pendingSide
func (s *OrderListPlaceOtocoWsRequest) PendingSide(pendingSide SideType) *OrderListPlaceOtocoWsRequest {
	s.pendingSide = pendingSide
	return s
}

// PendingQuantity set pendingQuantity
func (s *OrderListPlaceOtocoWsRequest) PendingQuantity(pendingQuantity string) *OrderListPlaceOtocoWsRequest {
	s.pendingQuantity = pendingQuantity
	return s
}

// PendingAboveType set pendingAboveType
func (s *OrderListPlaceOtocoWsRequest) PendingAboveType(pendingAboveType OrderType) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveType = pendingAboveType
	return s
}

// PendingAboveClientOrderID set pendingAboveClientOrderID
func (s *OrderListPlaceOtocoWsRequest) PendingAboveClientOrderID(pendingAboveClientOrderID string) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveClientOrderID = &pendingAboveClientOrderID
	return s
}

// PendingAbovePrice set pendingAbovePrice
func (s *OrderListPlaceOtocoWsRequest) PendingAbovePrice(pendingAbovePrice string) *OrderListPlaceOtocoWsRequest {
	s.pendingAbovePrice = &pendingAbovePrice
	return s
}

// PendingAboveStopPrice set pendingAboveStopPrice
func (s *OrderListPlaceOtocoWsRequest) PendingAboveStopPrice(pendingAboveStopPrice string) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveStopPrice = &pendingAboveStopPrice
	return s
}

// PendingAboveTrailingDelta set pendingAboveTrailingDelta
func (s *OrderListPlaceOtocoWsRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta int64) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return s
}

// PendingAboveIcebergQty set pendingAboveIcebergQty
func (s *OrderListPlaceOtocoWsRequest) PendingAboveIcebergQty(pendingAboveIcebergQty string) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return s
}

// PendingAboveTimeInForce set pendingAboveTimeInForce
func (s *OrderListPlaceOtocoWsRequest) PendingAboveTimeInForce(pendingAboveTimeInForce TimeInForceType) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return s
}

// PendingAboveStrategyId set pendingAboveStrategyId
func (s *OrderListPlaceOtocoWsRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveStrategyId = &pendingAboveStrategyId
	return s
}

// PendingAboveStrategyType set pendingAboveStrategyType
func (s *OrderListPlaceOtocoWsRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) *OrderListPlaceOtocoWsRequest {
	s.pendingAboveStrategyType = &pendingAboveStrategyType
	return s
}

// PendingBelowType set pendingBelowType
func (s *OrderListPlaceOtocoWsRequest) PendingBelowType(pendingBelowType OrderType) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowType = &pendingBelowType
	return s
}

// PendingBelowClientOrderID set pendingBelowClientOrderID
func (s *OrderListPlaceOtocoWsRequest) PendingBelowClientOrderID(pendingBelowClientOrderID string) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowClientOrderID = &pendingBelowClientOrderID
	return s
}

// PendingBelowPrice set pendingBelowPrice
func (s *OrderListPlaceOtocoWsRequest) PendingBelowPrice(pendingBelowPrice string) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowPrice = &pendingBelowPrice
	return s
}

// PendingBelowStopPrice set pendingBelowStopPrice
func (s *OrderListPlaceOtocoWsRequest) PendingBelowStopPrice(pendingBelowStopPrice string) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowStopPrice = &pendingBelowStopPrice
	return s
}

// PendingBelowTrailingDelta set pendingBelowTrailingDelta
func (s *OrderListPlaceOtocoWsRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta int64) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return s
}

// PendingBelowIcebergQty set pendingBelowIcebergQty
func (s *OrderListPlaceOtocoWsRequest) PendingBelowIcebergQty(pendingBelowIcebergQty string) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return s
}

// PendingBelowTimeInForce set pendingBelowTimeInForce
func (s *OrderListPlaceOtocoWsRequest) PendingBelowTimeInForce(pendingBelowTimeInForce TimeInForceType) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return s
}

// PendingBelowStrategyId set pendingBelowStrategyId
func (s *OrderListPlaceOtocoWsRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowStrategyId = &pendingBelowStrategyId
	return s
}

// PendingBelowStrategyType set pendingBelowStrategyType
func (s *OrderListPlaceOtocoWsRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) *OrderListPlaceOtocoWsRequest {
	s.pendingBelowStrategyType = &pendingBelowStrategyType
	return s
}

// RecvWindow set recvWindow
func (s *OrderListPlaceOtocoWsRequest) RecvWindow(recvWindow uint16) *OrderListPlaceOtocoWsRequest {
	s.recvWindow = &recvWindow
	return s
}
