package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/adshao/go-binance/v2/futures"
)

// AlgoOrderPlaceWsService creates algo order using WebSocket API
type AlgoOrderPlaceWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewAlgoOrderPlaceWsService init AlgoOrderPlaceWsService
func NewAlgoOrderPlaceWsService(apiKey, secretKey string) (*AlgoOrderPlaceWsService, error) {
	conn, err := websocket.NewConnection(futures.WsApiInitReadWriteConn, futures.WebsocketKeepalive, futures.WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &AlgoOrderPlaceWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// AlgoOrderPlaceWsRequest parameters for 'algoOrder.place' websocket API
type AlgoOrderPlaceWsRequest struct {
	algoType         futures.OrderAlgoType
	symbol           string
	side             futures.SideType
	_type            futures.AlgoOrderType
	positionSide     *futures.PositionSideType
	timeInForce      *futures.TimeInForceType
	quantity         *string
	price            *string
	triggerPrice     string
	workingType      *futures.WorkingType
	closePosition    *bool
	reduceOnly       *bool
	newClientOrderID *string
	newOrderRespType futures.NewOrderRespType
	recvWindow       *int64
}

// NewAlgoOrderPlaceWsRequest init AlgoOrderPlaceWsRequest
func NewAlgoOrderPlaceWsRequest() *AlgoOrderPlaceWsRequest {
	return &AlgoOrderPlaceWsRequest{
		algoType:         futures.OrderAlgoTypeConditional,
		newOrderRespType: futures.NewOrderRespTypeRESULT,
	}
}

// Symbol set symbol
func (s *AlgoOrderPlaceWsRequest) Symbol(symbol string) *AlgoOrderPlaceWsRequest {
	s.symbol = symbol
	return s
}

// Side set side
func (s *AlgoOrderPlaceWsRequest) Side(side futures.SideType) *AlgoOrderPlaceWsRequest {
	s.side = side
	return s
}

// Type set type
func (s *AlgoOrderPlaceWsRequest) Type(_type futures.AlgoOrderType) *AlgoOrderPlaceWsRequest {
	s._type = _type
	return s
}

// PositionSide set positionSide
func (s *AlgoOrderPlaceWsRequest) PositionSide(positionSide futures.PositionSideType) *AlgoOrderPlaceWsRequest {
	s.positionSide = &positionSide
	return s
}

// TimeInForce set timeInForce
func (s *AlgoOrderPlaceWsRequest) TimeInForce(timeInForce futures.TimeInForceType) *AlgoOrderPlaceWsRequest {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *AlgoOrderPlaceWsRequest) Quantity(quantity string) *AlgoOrderPlaceWsRequest {
	s.quantity = &quantity
	return s
}

// Price set price
func (s *AlgoOrderPlaceWsRequest) Price(price string) *AlgoOrderPlaceWsRequest {
	s.price = &price
	return s
}

// TriggerPrice set triggerPrice
func (s *AlgoOrderPlaceWsRequest) TriggerPrice(triggerPrice string) *AlgoOrderPlaceWsRequest {
	s.triggerPrice = triggerPrice
	return s
}

// WorkingType set workingType
func (s *AlgoOrderPlaceWsRequest) WorkingType(workingType futures.WorkingType) *AlgoOrderPlaceWsRequest {
	s.workingType = &workingType
	return s
}

// ClosePosition set closePosition
func (s *AlgoOrderPlaceWsRequest) ClosePosition(closePosition bool) *AlgoOrderPlaceWsRequest {
	s.closePosition = &closePosition
	return s
}

// ReduceOnly set reduceOnly
func (s *AlgoOrderPlaceWsRequest) ReduceOnly(reduceOnly bool) *AlgoOrderPlaceWsRequest {
	s.reduceOnly = &reduceOnly
	return s
}

// NewClientOrderID set newClientOrderID
func (s *AlgoOrderPlaceWsRequest) NewClientOrderID(newClientOrderID string) *AlgoOrderPlaceWsRequest {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *AlgoOrderPlaceWsRequest) NewOrderResponseType(newOrderResponseType futures.NewOrderRespType) *AlgoOrderPlaceWsRequest {
	s.newOrderRespType = newOrderResponseType
	return s
}

// RecvWindow set recvWindow
func (s *AlgoOrderPlaceWsRequest) RecvWindow(recvWindow int64) *AlgoOrderPlaceWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// buildParams builds params
func (s *AlgoOrderPlaceWsRequest) buildParams() map[string]interface{} {
	m := map[string]interface{}{
		"algoType":         s.algoType,
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s._type,
		"triggerPrice":     s.triggerPrice,
		"newOrderRespType": s.newOrderRespType,
	}

	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}

	return m
}

// CreateAlgoOrderResult define algo order creation result
type CreateAlgoOrderResult struct {
	AlgoId       int64  `json:"algoId"`
	ClientAlgoId string `json:"clientAlgoId"`
}

// CreateAlgoOrderWsResponse define 'algoOrder.place' websocket API response
type CreateAlgoOrderWsResponse struct {
	Id     string                `json:"id"`
	Status int                   `json:"status"`
	Result CreateAlgoOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}

// SyncDo - sends 'algoOrder.place' request and receives response
func (s *AlgoOrderPlaceWsService) SyncDo(requestID string, request *AlgoOrderPlaceWsRequest) (*CreateAlgoOrderWsResponse, error) {
	// Use custom method "algoOrder.place"
	method := websocket.WsApiMethodType("algoOrder.place")

	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		method,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	createAlgoOrderWsResponse := &CreateAlgoOrderWsResponse{}
	if err := json.Unmarshal(response, createAlgoOrderWsResponse); err != nil {
		return nil, err
	}

	return createAlgoOrderWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *AlgoOrderPlaceWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *AlgoOrderPlaceWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *AlgoOrderPlaceWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *AlgoOrderPlaceWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}
