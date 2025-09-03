package binance

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
)

// SorOrderTestWsService tests order using SOR
type SorOrderTestWsService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
}

// NewSorOrderTestWsService init SorOrderTestWsService
func NewSorOrderTestWsService(apiKey, secretKey string) (*SorOrderTestWsService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &SorOrderTestWsService{
		c:         client,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		KeyType:   common.KeyTypeHmac,
	}, nil
}

// SorOrderTestWsRequest parameters for 'sor.order.test' websocket API
type SorOrderTestWsRequest struct {
	symbol                  string
	side                    SideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	price                   *string
	quantity                string
	newClientOrderID        *string
	icebergQty              *string
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *SelfTradePreventionMode
	computeCommissionRates  *bool
	recvWindow              *uint16
}

// NewSorOrderTestWsRequest init SorOrderTestWsRequest
func NewSorOrderTestWsRequest() *SorOrderTestWsRequest {
	return &SorOrderTestWsRequest{}
}

func (s *SorOrderTestWsRequest) GetParams() map[string]interface{} {
	return s.buildParams()
}

// buildParams builds params
func (s *SorOrderTestWsRequest) buildParams() params {
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"type":     s.orderType,
		"quantity": s.quantity,
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
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
	if s.computeCommissionRates != nil {
		m["computeCommissionRates"] = *s.computeCommissionRates
	}
	if s.recvWindow != nil {
		m["recvWindow"] = *s.recvWindow
	}
	return m
}

// Do - sends 'sor.order.test' request
func (s *SorOrderTestWsService) Do(requestID string, request *SorOrderTestWsRequest) error {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.SorOrderTestSpotWsApiMethod,
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

// SyncDo - sends 'sor.order.test' request and receives response
func (s *SorOrderTestWsService) SyncDo(requestID string, request *SorOrderTestWsRequest) (*SorOrderTestWsResponse, error) {
	rawData, err := websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		websocket.SorOrderTestSpotWsApiMethod,
		request.buildParams(),
	)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	sorOrderTestWsResponse := &SorOrderTestWsResponse{}
	if err := json.Unmarshal(response, sorOrderTestWsResponse); err != nil {
		return nil, err
	}

	return sorOrderTestWsResponse, nil
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *SorOrderTestWsService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *SorOrderTestWsService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *SorOrderTestWsService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *SorOrderTestWsService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

// Symbol set symbol
func (s *SorOrderTestWsRequest) Symbol(symbol string) *SorOrderTestWsRequest {
	s.symbol = symbol
	return s
}

// Side set side
func (s *SorOrderTestWsRequest) Side(side SideType) *SorOrderTestWsRequest {
	s.side = side
	return s
}

// Type set orderType
func (s *SorOrderTestWsRequest) Type(orderType OrderType) *SorOrderTestWsRequest {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *SorOrderTestWsRequest) TimeInForce(timeInForce TimeInForceType) *SorOrderTestWsRequest {
	s.timeInForce = &timeInForce
	return s
}

// Price set price
func (s *SorOrderTestWsRequest) Price(price string) *SorOrderTestWsRequest {
	s.price = &price
	return s
}

// Quantity set quantity
func (s *SorOrderTestWsRequest) Quantity(quantity string) *SorOrderTestWsRequest {
	s.quantity = quantity
	return s
}

// NewClientOrderID set newClientOrderID
func (s *SorOrderTestWsRequest) NewClientOrderID(newClientOrderID string) *SorOrderTestWsRequest {
	s.newClientOrderID = &newClientOrderID
	return s
}

// IcebergQty set icebergQty
func (s *SorOrderTestWsRequest) IcebergQty(icebergQty string) *SorOrderTestWsRequest {
	s.icebergQty = &icebergQty
	return s
}

// StrategyId set strategyId
func (s *SorOrderTestWsRequest) StrategyId(strategyId int64) *SorOrderTestWsRequest {
	s.strategyId = &strategyId
	return s
}

// StrategyType set strategyType
func (s *SorOrderTestWsRequest) StrategyType(strategyType int32) *SorOrderTestWsRequest {
	s.strategyType = &strategyType
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *SorOrderTestWsRequest) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *SorOrderTestWsRequest {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// ComputeCommissionRates set computeCommissionRates
func (s *SorOrderTestWsRequest) ComputeCommissionRates(computeCommissionRates bool) *SorOrderTestWsRequest {
	s.computeCommissionRates = &computeCommissionRates
	return s
}

// RecvWindow set recvWindow
func (s *SorOrderTestWsRequest) RecvWindow(recvWindow uint16) *SorOrderTestWsRequest {
	s.recvWindow = &recvWindow
	return s
}

// SorOrderTestResult define SOR order test result
type SorOrderTestResult struct {
	StandardCommissionForOrder *struct {
		Maker string `json:"maker"`
		Taker string `json:"taker"`
	} `json:"standardCommissionForOrder,omitempty"`
	TaxCommissionForOrder *struct {
		Maker string `json:"maker"`
		Taker string `json:"taker"`
	} `json:"taxCommissionForOrder,omitempty"`
	Discount *struct {
		EnabledForAccount bool   `json:"enabledForAccount"`
		EnabledForSymbol  bool   `json:"enabledForSymbol"`
		DiscountAsset     string `json:"discountAsset"`
		Discount          string `json:"discount"`
	} `json:"discount,omitempty"`
}

// SorOrderTestWsResponse define 'sor.order.test' websocket API response
type SorOrderTestWsResponse struct {
	Id     string             `json:"id"`
	Status int                `json:"status"`
	Result SorOrderTestResult `json:"result"`

	// error response
	Error *common.APIError `json:"error,omitempty"`
}
