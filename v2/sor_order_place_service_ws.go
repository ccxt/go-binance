package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// SorOrderPlaceWsService places order using SOR
type SorOrderPlaceWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewSorOrderPlaceWsService init SorOrderPlaceWsService
func NewSorOrderPlaceWsService(apiKey, secretKey string) (*SorOrderPlaceWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &SorOrderPlaceWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// SorOrderPlaceWsRequest parameters for 'sor.order.place' websocket API
type SorOrderPlaceWsRequest struct {
	symbol                  string
	side                    SideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	price                   *string
	quantity                string
	newClientOrderID        *string
	newOrderRespType        NewOrderRespType
	icebergQty              *string
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *SelfTradePreventionMode
	recvWindow              *uint16
}

// NewSorOrderPlaceWsRequest init SorOrderPlaceWsRequest
func NewSorOrderPlaceWsRequest() *SorOrderPlaceWsRequest {
	return &SorOrderPlaceWsRequest{
		newOrderRespType: NewOrderRespTypeFULL,
	}
}

func (s *SorOrderPlaceWsRequest) GetParams() map[string]any {
	return s.buildParams()
}

// buildParams builds params
func (s *SorOrderPlaceWsRequest) buildParams() params {
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
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
	if s.icebergQty != nil {
		m["icebergQty"] = *s.icebergQty
	}
	if s.strategyId != nil {
		m["strategyId"] = *s.strategyId
	}
	if s.strategyType != nil {
		m["strategyType"] = *s.strategyType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'sor.order.place' request
func (s *SorOrderPlaceWsService) Do(requestID string, request *SorOrderPlaceWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.SorOrderPlaceSpotWsApiMethod,
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

// SyncDo - sends 'sor.order.place' request and receives response
func (s *SorOrderPlaceWsService) SyncDo(requestID string, request *SorOrderPlaceWsRequest) (*SorOrderPlaceWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.SorOrderPlaceSpotWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	sorOrderPlaceWsResponse := &SorOrderPlaceWsResponse{}
	if err := json.Unmarshal(response, sorOrderPlaceWsResponse); err != nil {
		return nil, err
	}

	return sorOrderPlaceWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *SorOrderPlaceWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *SorOrderPlaceWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *SorOrderPlaceWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *SorOrderPlaceWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *SorOrderPlaceWsRequest) Symbol(symbol string) *SorOrderPlaceWsRequest {
	s.symbol = symbol
	return s
}

// Side set side
func (s *SorOrderPlaceWsRequest) Side(side SideType) *SorOrderPlaceWsRequest {
	s.side = side
	return s
}

// Type set orderType
func (s *SorOrderPlaceWsRequest) Type(orderType OrderType) *SorOrderPlaceWsRequest {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *SorOrderPlaceWsRequest) TimeInForce(timeInForce TimeInForceType) *SorOrderPlaceWsRequest {
	s.timeInForce = &timeInForce
	return s
}

// Price set price
func (s *SorOrderPlaceWsRequest) Price(price string) *SorOrderPlaceWsRequest {
	s.price = &price
	return s
}

// Quantity set quantity
func (s *SorOrderPlaceWsRequest) Quantity(quantity string) *SorOrderPlaceWsRequest {
	s.quantity = quantity
	return s
}

// NewClientOrderID set newClientOrderID
func (s *SorOrderPlaceWsRequest) NewClientOrderID(newClientOrderID string) *SorOrderPlaceWsRequest {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderRespType set newOrderRespType
func (s *SorOrderPlaceWsRequest) NewOrderRespType(newOrderRespType NewOrderRespType) *SorOrderPlaceWsRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

// IcebergQty set icebergQty
func (s *SorOrderPlaceWsRequest) IcebergQty(icebergQty string) *SorOrderPlaceWsRequest {
	s.icebergQty = &icebergQty
	return s
}

// StrategyId set strategyId
func (s *SorOrderPlaceWsRequest) StrategyId(strategyId int64) *SorOrderPlaceWsRequest {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *SorOrderPlaceWsRequest) StrategyType(strategyType int32) *SorOrderPlaceWsRequest {
	s.strategyType = &strategyType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *SorOrderPlaceWsRequest) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *SorOrderPlaceWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// RecvWindow set recvWindow
func (s *SorOrderPlaceWsRequest) RecvWindow(recvWindow uint16) *SorOrderPlaceWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// SorOrderPlaceResult define SOR order placement result
type SorOrderPlaceResult struct {
	Symbol              string          `json:"symbol"`
	OrderId             int64           `json:"orderId"`
	OrderListId         int64           `json:"orderListId"`
	ClientOrderId       string          `json:"clientOrderId"`
	TransactTime        int64           `json:"transactTime"`
	Price               string          `json:"price"`
	OrigQty             string          `json:"origQty"`
	ExecutedQty         string          `json:"executedQty"`
	CummulativeQuoteQty string          `json:"cummulativeQuoteQty"`
	Status              OrderStatusType `json:"status"`
	TimeInForce         TimeInForceType `json:"timeInForce"`
	Type                OrderType       `json:"type"`
	Side                SideType        `json:"side"`
	WorkingTime         int64           `json:"workingTime"`
	Fills               []struct {
		MatchType       string `json:"matchType"`
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int64  `json:"tradeId"`
		AllocId         int64  `json:"allocId"`
	} `json:"fills"`
	WorkingFloor            string `json:"workingFloor"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	UsedSor                 bool   `json:"usedSor"`
}

// SorOrderPlaceWsResponse define 'sor.order.place' websocket API response
type SorOrderPlaceWsResponse struct {
	Id     string                `json:"id"`
	Status int                   `json:"status"`
	Result []SorOrderPlaceResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}
