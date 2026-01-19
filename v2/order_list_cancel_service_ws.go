package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderListCancelWsService cancels order list
type OrderListCancelWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderListCancelWsService init OrderListCancelWsService
func NewOrderListCancelWsService(apiKey, secretKey string) (*OrderListCancelWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderListCancelWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderListCancelWsRequest parameters for 'orderList.cancel' websocket API
type OrderListCancelWsRequest struct {
	symbol            string
	orderListID       *int64
	listClientOrderID *string
	newClientOrderID  *string
	recvWindow        *uint16
}

// NewOrderListCancelWsRequest init OrderListCancelWsRequest
func NewOrderListCancelWsRequest() *OrderListCancelWsRequest {
	return &OrderListCancelWsRequest{}
}

func (s *OrderListCancelWsRequest) GetParams() map[string]any {
	return s.buildParams()
}

// buildParams builds params
func (s *OrderListCancelWsRequest) buildParams() params {
	m := params{
		"symbol": s.symbol,
	}
	if s.orderListID != nil {
		m["orderListId"] = *s.orderListID
	}
	if s.listClientOrderID != nil {
		m["listClientOrderId"] = *s.listClientOrderID
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'orderList.cancel' request
func (s *OrderListCancelWsService) Do(requestID string, request *OrderListCancelWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListCancelSpotWsApiMethod,
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

// SyncDo - sends 'orderList.cancel' request and receives response
func (s *OrderListCancelWsService) SyncDo(requestID string, request *OrderListCancelWsRequest) (*CancelOrderListWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListCancelSpotWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	cancelOrderListWsResponse := &CancelOrderListWsResponse{}
	if err := json.Unmarshal(response, cancelOrderListWsResponse); err != nil {
		return nil, err
	}

	return cancelOrderListWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *OrderListCancelWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderListCancelWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderListCancelWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderListCancelWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *OrderListCancelWsRequest) Symbol(symbol string) *OrderListCancelWsRequest {
	s.symbol = symbol
	return s
}

// OrderListID set orderListID
func (s *OrderListCancelWsRequest) OrderListID(orderListID int64) *OrderListCancelWsRequest {
	s.orderListID = &orderListID
	return s
}

// ListClientOrderID set listClientOrderID
func (s *OrderListCancelWsRequest) ListClientOrderID(listClientOrderID string) *OrderListCancelWsRequest {
	s.listClientOrderID = &listClientOrderID
	return s
}

// NewClientOrderID set newClientOrderID
func (s *OrderListCancelWsRequest) NewClientOrderID(newClientOrderID string) *OrderListCancelWsRequest {
	s.newClientOrderID = &newClientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *OrderListCancelWsRequest) RecvWindow(recvWindow uint16) *OrderListCancelWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// CancelOrderListResult define order list cancellation result
type CancelOrderListResult struct {
	OrderListId       int64  `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol                  string          `json:"symbol"`
		OrderId                 int64           `json:"orderId"`
		OrderListId             int64           `json:"orderListId"`
		ClientOrderId           string          `json:"clientOrderId"`
		TransactTime            int64           `json:"transactTime"`
		Price                   string          `json:"price"`
		OrigQty                 string          `json:"origQty"`
		ExecutedQty             string          `json:"executedQty"`
		CummulativeQuoteQty     string          `json:"cummulativeQuoteQty"`
		Status                  OrderStatusType `json:"status"`
		TimeInForce             TimeInForceType `json:"timeInForce"`
		Type                    OrderType       `json:"type"`
		Side                    SideType        `json:"side"`
		StopPrice               string          `json:"stopPrice"`
		SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	} `json:"orderReports"`
}

// CancelOrderListWsResponse define 'orderList.cancel' websocket API response
type CancelOrderListWsResponse struct {
	Id     string                `json:"id"`
	Status int                   `json:"status"`
	Result CancelOrderListResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}
