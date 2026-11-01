package futures

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/adshao/go-binance/v2/common/websocket/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func (s *orderStatusServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "e2a85d9f-07a5-4f94-8d5f-789dc3deb098"

	s.symbol = "BTCUSDT"
	s.orderID = int64(123456)
	s.origClientOrderID = "testOrder"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderStatus = &OrderStatusWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderStatusRequest = NewOrderStatusWsRequest().
		Symbol(s.symbol).
		OrderID(s.orderID).
		OrigClientOrderID(s.origClientOrderID)
}

func (s *orderStatusServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderStatusServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockClient

	requestID         string
	symbol            string
	orderID           int64
	origClientOrderID string

	orderStatus        *OrderStatusWsService
	orderStatusRequest *OrderStatusWsRequest
}

func TestOrderStatusServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderStatusServiceWsTestSuite))
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatus() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderStatus.Do(s.requestID, s.orderStatusRequest)
	s.NoError(err)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatus_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderStatus.Do("", s.orderStatusRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatus_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderStatus.Do(s.requestID, s.orderStatusRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatus_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderStatus.Do(s.requestID, s.orderStatusRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatus_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderStatus.Do(s.requestID, s.orderStatusRequest)
	s.Error(err)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatusSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderStatusResponse := QueryOrderWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: QueryOrderResult{
			QueryOrderResponse{
				Symbol:        s.symbol,
				OrderID:       s.orderID,
				ClientOrderID: s.origClientOrderID,
			},
		},
	}

	rawResponseData, err := json.Marshal(orderStatusResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderStatusRequest
	response, err := s.orderStatus.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal(req.orderID, response.Result.OrderID)
	s.Equal(req.origClientOrderID, response.Result.ClientOrderID)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatusSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderStatusRequest
	response, err := s.orderStatus.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatusSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderStatus.SyncDo(s.requestID, s.orderStatusRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatusSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderStatus.SyncDo(s.requestID, s.orderStatusRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderStatusServiceWsTestSuite) TestOrderStatusSync_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderStatus.SyncDo(s.requestID, s.orderStatusRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderStatusServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderStatus = &OrderStatusWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
