package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// OrderListCreateWsService creates OCO order list
type OrderListCreateWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewOrderListCreateWsService init OrderListCreateWsService
func NewOrderListCreateWsService(apiKey, secretKey string) (*OrderListCreateWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &OrderListCreateWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// OrderListCreateWsRequest parameters for 'orderList.place.oco' websocket API
type OrderListCreateWsRequest struct {
	symbol                  string
	listClientOrderID       *string
	side                    SideType
	quantity                string
	aboveType               OrderType
	aboveClientOrderID      *string
	aboveIcebergQty         *string
	abovePrice              *string
	aboveStopPrice          *string
	aboveTrailingDelta      *int64
	aboveTimeInForce        *TimeInForceType
	aboveStrategyId         *int64
	aboveStrategyType       *int64
	belowType               OrderType
	belowClientOrderID      *string
	belowIcebergQty         *string
	belowPrice              *string
	belowStopPrice          *string
	belowTrailingDelta      *int64
	belowTimeInForce        *TimeInForceType
	belowStrategyId         *int64
	belowStrategyType       *int64
	newOrderRespType        NewOrderRespType
	selfTradePreventionMode *SelfTradePreventionMode
	recvWindow              *uint16
}

// NewOrderListCreateWsRequest init OrderListCreateWsRequest
func NewOrderListCreateWsRequest() *OrderListCreateWsRequest {
	return &OrderListCreateWsRequest{}
}

func (s *OrderListCreateWsRequest) GetParams() map[string]any {
	return s.buildParams()
}

// buildParams builds params
func (s *OrderListCreateWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"quantity":         s.quantity,
		"aboveType":        s.aboveType,
		"belowType":        s.belowType,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.listClientOrderID != nil {
		m["listClientOrderId"] = *s.listClientOrderID
	} else {
		m["listClientOrderId"] = common.GenerateSpotId()
	}
	if s.aboveClientOrderID != nil {
		m["aboveClientOrderId"] = *s.aboveClientOrderID
	} else {
		m["aboveClientOrderId"] = common.GenerateSpotId()
	}
	if s.aboveIcebergQty != nil {
		m["aboveIcebergQty"] = *s.aboveIcebergQty
	}
	if s.abovePrice != nil {
		m["abovePrice"] = *s.abovePrice
	}
	if s.aboveStopPrice != nil {
		m["aboveStopPrice"] = *s.aboveStopPrice
	}
	if s.aboveTrailingDelta != nil {
		m["aboveTrailingDelta"] = *s.aboveTrailingDelta
	}
	if s.aboveTimeInForce != nil {
		m["aboveTimeInForce"] = *s.aboveTimeInForce
	}
	if s.aboveStrategyId != nil {
		m["aboveStrategyId"] = *s.aboveStrategyId
	}
	if s.aboveStrategyType != nil {
		m["aboveStrategyType"] = *s.aboveStrategyType
	}
	if s.belowClientOrderID != nil {
		m["belowClientOrderId"] = *s.belowClientOrderID
	} else {
		m["belowClientOrderId"] = common.GenerateSpotId()
	}
	if s.belowIcebergQty != nil {
		m["belowIcebergQty"] = *s.belowIcebergQty
	}
	if s.belowPrice != nil {
		m["belowPrice"] = *s.belowPrice
	}
	if s.belowStopPrice != nil {
		m["belowStopPrice"] = *s.belowStopPrice
	}
	if s.belowTrailingDelta != nil {
		m["belowTrailingDelta"] = *s.belowTrailingDelta
	}
	if s.belowTimeInForce != nil {
		m["belowTimeInForce"] = *s.belowTimeInForce
	}
	if s.belowStrategyId != nil {
		m["belowStrategyId"] = *s.belowStrategyId
	}
	if s.belowStrategyType != nil {
		m["belowStrategyType"] = *s.belowStrategyType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'orderList.place.oco' request
func (s *OrderListCreateWsService) Do(requestID string, request *OrderListCreateWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceOcoSpotWsApiMethod,
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

// SyncDo - sends 'orderList.place.oco' request and receives response
func (s *OrderListCreateWsService) SyncDo(requestID string, request *OrderListCreateWsRequest) (*CreateOrderListWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.OrderListPlaceOcoSpotWsApiMethod,
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
func (s *OrderListCreateWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *OrderListCreateWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *OrderListCreateWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *OrderListCreateWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *OrderListCreateWsRequest) Symbol(symbol string) *OrderListCreateWsRequest {
	s.symbol = symbol
	return s
}

// ListClientOrderID set listClientOrderID
func (s *OrderListCreateWsRequest) ListClientOrderID(listClientOrderID string) *OrderListCreateWsRequest {
	s.listClientOrderID = &listClientOrderID
	return s
}

// Side set side
func (s *OrderListCreateWsRequest) Side(side SideType) *OrderListCreateWsRequest {
	s.side = side
	return s
}

// Quantity set quantity
func (s *OrderListCreateWsRequest) Quantity(quantity string) *OrderListCreateWsRequest {
	s.quantity = quantity
	return s
}

// AboveType set aboveType
func (s *OrderListCreateWsRequest) AboveType(aboveType OrderType) *OrderListCreateWsRequest {
	s.aboveType = aboveType
	return s
}

// AboveClientOrderID set aboveClientOrderID
func (s *OrderListCreateWsRequest) AboveClientOrderID(aboveClientOrderID string) *OrderListCreateWsRequest {
	s.aboveClientOrderID = &aboveClientOrderID
	return s
}

// AboveIcebergQty set aboveIcebergQty
func (s *OrderListCreateWsRequest) AboveIcebergQty(aboveIcebergQty string) *OrderListCreateWsRequest {
	s.aboveIcebergQty = &aboveIcebergQty
	return s
}

// AbovePrice set abovePrice
func (s *OrderListCreateWsRequest) AbovePrice(abovePrice string) *OrderListCreateWsRequest {
	s.abovePrice = &abovePrice
	return s
}

// AboveStopPrice set aboveStopPrice
func (s *OrderListCreateWsRequest) AboveStopPrice(aboveStopPrice string) *OrderListCreateWsRequest {
	s.aboveStopPrice = &aboveStopPrice
	return s
}

// AboveTrailingDelta set aboveTrailingDelta
func (s *OrderListCreateWsRequest) AboveTrailingDelta(aboveTrailingDelta int64) *OrderListCreateWsRequest {
	s.aboveTrailingDelta = &aboveTrailingDelta
	return s
}

// AboveTimeInForce set aboveTimeInForce
func (s *OrderListCreateWsRequest) AboveTimeInForce(aboveTimeInForce TimeInForceType) *OrderListCreateWsRequest {
	s.aboveTimeInForce = &aboveTimeInForce
	return s
}

// AboveStrategyId set aboveStrategyId
func (s *OrderListCreateWsRequest) AboveStrategyId(aboveStrategyId int64) *OrderListCreateWsRequest {
	s.aboveStrategyId = &aboveStrategyId
	return s
}

// AboveStrategyType set aboveStrategyType
func (s *OrderListCreateWsRequest) AboveStrategyType(aboveStrategyType int64) *OrderListCreateWsRequest {
	s.aboveStrategyType = &aboveStrategyType
	return s
}

// BelowType set belowType
func (s *OrderListCreateWsRequest) BelowType(belowType OrderType) *OrderListCreateWsRequest {
	s.belowType = belowType
	return s
}

// BelowClientOrderID set belowClientOrderID
func (s *OrderListCreateWsRequest) BelowClientOrderID(belowClientOrderID string) *OrderListCreateWsRequest {
	s.belowClientOrderID = &belowClientOrderID
	return s
}

// BelowIcebergQty set belowIcebergQty
func (s *OrderListCreateWsRequest) BelowIcebergQty(belowIcebergQty string) *OrderListCreateWsRequest {
	s.belowIcebergQty = &belowIcebergQty
	return s
}

// BelowPrice set belowPrice
func (s *OrderListCreateWsRequest) BelowPrice(belowPrice string) *OrderListCreateWsRequest {
	s.belowPrice = &belowPrice
	return s
}

// BelowStopPrice set belowStopPrice
func (s *OrderListCreateWsRequest) BelowStopPrice(belowStopPrice string) *OrderListCreateWsRequest {
	s.belowStopPrice = &belowStopPrice
	return s
}

// BelowTrailingDelta set belowTrailingDelta
func (s *OrderListCreateWsRequest) BelowTrailingDelta(belowTrailingDelta int64) *OrderListCreateWsRequest {
	s.belowTrailingDelta = &belowTrailingDelta
	return s
}

// BelowTimeInForce set belowTimeInForce
func (s *OrderListCreateWsRequest) BelowTimeInForce(belowTimeInForce TimeInForceType) *OrderListCreateWsRequest {
	s.belowTimeInForce = &belowTimeInForce
	return s
}

// BelowStrategyId set belowStrategyId
func (s *OrderListCreateWsRequest) BelowStrategyId(belowStrategyId int64) *OrderListCreateWsRequest {
	s.belowStrategyId = &belowStrategyId
	return s
}

// BelowStrategyType set belowStrategyType
func (s *OrderListCreateWsRequest) BelowStrategyType(belowStrategyType int64) *OrderListCreateWsRequest {
	s.belowStrategyType = &belowStrategyType
	return s
}

// NewOrderRespType set newOrderRespType
func (s *OrderListCreateWsRequest) NewOrderRespType(newOrderRespType NewOrderRespType) *OrderListCreateWsRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *OrderListCreateWsRequest) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *OrderListCreateWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// RecvWindow set recvWindow
func (s *OrderListCreateWsRequest) RecvWindow(recvWindow uint16) *OrderListCreateWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// CreateOrderListResult define order list creation result
type CreateOrderListResult struct {
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
		WorkingTime             int64           `json:"workingTime"`
		SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	} `json:"orderReports"`
}

// CreateOrderListWsResponse define 'orderList.place.oco' websocket API response
type CreateOrderListWsResponse struct {
	Id     string                `json:"id"`
	Status int                   `json:"status"`
	Result CreateOrderListResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}
