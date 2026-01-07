package futures

import (
	"encoding/json"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/google/uuid"
)

type WsAccountService struct {
	c          websocket.Client
	ApiKey     string
	SecretKey  string
	KeyType    string
	TimeOffset int64
	RecvWindow int64
}

func NewWsAccountService(apiKey, secretKey string, recvWindow ...int64) (*WsAccountService, error) {
	conn, err := websocket.NewConnection(WsApiInitReadWriteConn, WebsocketKeepalive, WebsocketTimeoutReadWriteConnection)
	if err != nil {
		return nil, err
	}

	client, err := websocket.NewClient(conn)
	if err != nil {
		return nil, err
	}

	window := int64(5000)
	if len(recvWindow) > 0 {
		window = recvWindow[0]
	}

	return &WsAccountService{
		c:          client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    common.KeyTypeHmac,
		RecvWindow: window,
	}, nil
}

type WsAccountV2InfoResponse struct {
	ID        string             `json:"id"`
	Status    int                `json:"status"`
	Result    AccountV3          `json:"result"`
	RateLimit []AccountRateLimit `json:"rateLimits"`
	Error     *common.APIError   `json:"error,omitempty"`
}

type WsAccountV2BalanceResponse struct {
	ID        string             `json:"id"`
	Status    int                `json:"status"`
	Result    []*Balance         `json:"result"`
	RateLimit []AccountRateLimit `json:"rateLimits"`
	Error     *common.APIError   `json:"error,omitempty"`
}

type AccountRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

const (
	AccountV2InfoMethod    websocket.WsApiMethodType = "v2/account.status"
	AccountV2BalanceMethod websocket.WsApiMethodType = "v2/account.balance"
)

func (s *WsAccountService) GetAccountInfo(requestID string) error {
	rawData, err := s.buildRequest(requestID, AccountV2InfoMethod)
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

func (s *WsAccountService) SyncGetAccountInfo(requestID string) (*WsAccountV2InfoResponse, error) {
	rawData, err := s.buildRequest(requestID, AccountV2InfoMethod)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	info := &WsAccountV2InfoResponse{}
	if err := json.Unmarshal(response, info); err != nil {
		return nil, err
	}

	return info, nil
}

func (s *WsAccountService) GetAccountBalance(requestID string) error {
	rawData, err := s.buildRequest(requestID, AccountV2BalanceMethod)
	if err != nil {
		return err
	}

	if err := s.c.Write(requestID, rawData); err != nil {
		return err
	}

	return nil
}

func (s *WsAccountService) SyncGetAccountBalance(requestID string) (*WsAccountV2BalanceResponse, error) {
	rawData, err := s.buildRequest(requestID, AccountV2BalanceMethod)
	if err != nil {
		return nil, err
	}

	response, err := s.c.WriteSync(requestID, rawData, websocket.WriteSyncWsTimeout)
	if err != nil {
		return nil, err
	}

	balance := &WsAccountV2BalanceResponse{}
	if err := json.Unmarshal(response, balance); err != nil {
		return nil, err
	}

	return balance, nil
}

func (s *WsAccountService) buildRequest(requestID string, method websocket.WsApiMethodType) ([]byte, error) {
	return websocket.CreateRequest(
		websocket.NewRequestData(
			requestID,
			s.ApiKey,
			s.SecretKey,
			s.TimeOffset,
			s.KeyType,
		),
		method,
		map[string]any{
			"recvWindow": s.RecvWindow,
		},
	)
}

// ReceiveAllDataBeforeStop waits until all responses will be received from websocket until timeout expired
func (s *WsAccountService) ReceiveAllDataBeforeStop(timeout time.Duration) {
	s.c.Wait(timeout)
}

// GetReadChannel returns channel with API response data (including API errors)
func (s *WsAccountService) GetReadChannel() <-chan []byte {
	return s.c.GetReadChannel()
}

// GetReadErrorChannel returns channel with errors which are occurred while reading websocket connection
func (s *WsAccountService) GetReadErrorChannel() <-chan error {
	return s.c.GetReadErrorChannel()
}

// GetReconnectCount returns count of reconnect attempts by client
func (s *WsAccountService) GetReconnectCount() int64 {
	return s.c.GetReconnectCount()
}

func (c *Client) NewWsAccountService(recvWindow ...int64) (*WsAccountService, error) {
	return NewWsAccountService(c.APIKey, c.SecretKey, recvWindow...)
}

// GetAccountInfoWs Get account info by websocket like RESTful
func (c *Client) GetAccountInfoWs(recvWindow ...int64) (*WsAccountV2InfoResponse, error) {
	service, err := c.NewWsAccountService(recvWindow...)
	if err != nil {
		return nil, err
	}
	defer service.c.Close()

	response, err := service.SyncGetAccountInfo(uuid.New().String())
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetAccountBalanceWs(recvWindow ...int64) (*WsAccountV2BalanceResponse, error) {
	service, err := c.NewWsAccountService(recvWindow...)
	if err != nil {
		return nil, err
	}
	defer service.c.Close()

	response, err := service.SyncGetAccountBalance(uuid.New().String())
	if err != nil {
		return nil, err
	}

	return response, nil
}
