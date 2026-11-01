package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/adshao/go-binance/v2/futures"
)

// AlgoOrderCancelWsService cancels algo order using WebSocket API
type AlgoOrderCancelWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewAlgoOrderCancelWsService init AlgoOrderCancelWsService
func NewAlgoOrderCancelWsService(apiKey, secretKey string) (*AlgoOrderCancelWsService, error) {
	conn, err := websocket.NewConnection(futures.WsApiInitReadWriteConn, futures.WebsocketKeepalive, futures.WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &AlgoOrderCancelWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// AlgoOrderCancelWsRequest parameters for 'algoOrder.cancel' websocket API
type AlgoOrderCancelWsRequest struct {
	algoId       *int64
	clientAlgoId *string
	recvWindow   *int64
}

// NewAlgoOrderCancelWsRequest init AlgoOrderCancelWsRequest
func NewAlgoOrderCancelWsRequest() *AlgoOrderCancelWsRequest {
	return &AlgoOrderCancelWsRequest{}
}

// AlgoID set algoID
func (s *AlgoOrderCancelWsRequest) AlgoID(algoID int64) *AlgoOrderCancelWsRequest {
	s.algoId = &algoID
	return s
}

// ClientAlgoID set clientAlgoID
func (s *AlgoOrderCancelWsRequest) ClientAlgoID(clientAlgoID string) *AlgoOrderCancelWsRequest {
	s.clientAlgoId = &clientAlgoID
	return s
}

// RecvWindow set recvWindow
func (s *AlgoOrderCancelWsRequest) RecvWindow(recvWindow int64) *AlgoOrderCancelWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// buildParams builds params
func (s *AlgoOrderCancelWsRequest) buildParams() map[string]interface{} {
	m := map[string]interface{}{}

	if s.algoId != nil {
		m["algoid"] = *s.algoId
	}

	if s.clientAlgoId != nil {
		m["clientalgoid"] = *s.clientAlgoId
	}

	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}

	return m
}

// CancelAlgoOrderResult define algo order cancel result
type CancelAlgoOrderResult struct {
	AlgoId       int64  `json:"algoId"`
	ClientAlgoId string `json:"clientAlgoId"`
	Code         string `json:"code"`
	Message      string `json:"msg"`
}

// CancelAlgoOrderWsResponse define 'algoOrder.cancel' websocket API response
type CancelAlgoOrderWsResponse struct {
	Id     string                `json:"id"`
	Status int                   `json:"status"`
	Result CancelAlgoOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}

// Do - sends 'algoOrder.cancel' request
func (s *AlgoOrderCancelWsService) Do(requestID string, request *AlgoOrderCancelWsRequest) error {
	// Use custom method "algoOrder.cancel"
	method := websocket.WsApiMethodType("algoOrder.cancel")

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
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

// SyncDo - sends 'algoOrder.cancel' request and receives response
func (s *AlgoOrderCancelWsService) SyncDo(requestID string, request *AlgoOrderCancelWsRequest) (*CancelAlgoOrderWsResponse, error) {
	// Use custom method "algoOrder.cancel"
	method := websocket.WsApiMethodType("algoOrder.cancel")

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

	cancelAlgoOrderWsResponse := &CancelAlgoOrderWsResponse{}
	if err := json.Unmarshal(response, cancelAlgoOrderWsResponse); err != nil {
		return nil, err
	}

	return cancelAlgoOrderWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *AlgoOrderCancelWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *AlgoOrderCancelWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *AlgoOrderCancelWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *AlgoOrderCancelWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}
