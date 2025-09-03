package futures

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderStatusWsService query order
type OrderStatusWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderStatusWsService init OrderStatusWsService
func NewOrderStatusWsService(apiKey, secretKey string) (*OrderStatusWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderStatusWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderStatusWsRequest parameters for 'order.status' websocket API
type OrderStatusWsRequest struct {
	symbol            string
	orderID           int64
	origClientOrderID string
	timestamp         int64
}

// NewOrderStatusWsRequest init OrderStatusWsRequest
func NewOrderStatusWsRequest() *OrderStatusWsRequest {
	return &OrderStatusWsRequest{}
}

// Symbol set symbol
func (s *OrderStatusWsRequest) Symbol(symbol string) *OrderStatusWsRequest {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *OrderStatusWsRequest) OrderID(orderID int64) *OrderStatusWsRequest {
	s.orderID = orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *OrderStatusWsRequest) OrigClientOrderID(origClientOrderID string) *OrderStatusWsRequest {
	s.origClientOrderID = origClientOrderID
	return s
}

// QueryOrderResponse define query order response
type QueryOrderResponse struct {
	AvgPrice      string `json:"avgPrice"`
	ClientOrderID string `json:"clientOrderId"`
	CumQuote      string `json:"cumQuote"`
	ExecutedQty   string `json:"executedQty"`
	OrderID       int64  `json:"orderId"`
	OrigQty       string `json:"origQty"`
	OrigType      string `json:"origType"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Time          int64  `json:"time"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
	ClosePosition bool   `json:"closePosition"`
	PriceProtect  bool   `json:"priceProtect"`
	StopPrice     string `json:"stopPrice"`
	ActivatePrice string `json:"activatePrice"`
	PriceRate     string `json:"priceRate"`
	WorkingType   string `json:"workingType"`
}

// QueryOrderResult define order creation result
type QueryOrderResult struct {
	QueryOrderResponse
}

// QueryOrderWsResponse define 'order.status' websocket API response
type QueryOrderWsResponse struct {
	Id     string           `json:"id"`
	Status int              `json:"status"`
	Result QueryOrderResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}

func (r *OrderStatusWsRequest) GetParams() map[string]interface{} {
	return r.buildParams()
}

// buildParams builds params
func (s *OrderStatusWsRequest) buildParams() params {
	m := params{
		"symbol":  s.symbol,
		"orderId": s.orderID,
	}
	if s.origClientOrderID != "" {
		m["origClientOrderId"] = s.origClientOrderID
	}

	return m
}

// Do - sends 'order.status' request
func (s *OrderStatusWsService) Do(requestID string, request *OrderStatusWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderStatusFuturesWsApiMethod,
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

// SyncDo - sends 'order.status' request and receives response
func (s *OrderStatusWsService) SyncDo(requestID string, request *OrderStatusWsRequest) (*QueryOrderWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderStatusFuturesWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	queryOrderWsResponse := &QueryOrderWsResponse{}
	if err := json.Unmarshal(response, queryOrderWsResponse); err != nil {
		return nil, err
	}

	return queryOrderWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderStatusWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderStatusWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderStatusWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderStatusWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}
