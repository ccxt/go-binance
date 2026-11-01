package binance

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adshao/go-binance/v2/common/websocket"
	"github.com/adshao/go-binance/v2/common/websocket/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func (s *orderListCancelServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "c5899911-d3f4-47ae-8835-97da553d27d0"

	s.symbol = "BTCUSDT"
	s.orderListID = int64(1274512)
	s.listClientOrderID = "6023531d7edaad348f5aff"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderListCancel = &OrderListCancelWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderListCancelRequest = NewOrderListCancelWsRequest().
		Symbol(s.symbol).
		OrderListID(s.orderListID).
		ListClientOrderID(s.listClientOrderID)
}

func (s *orderListCancelServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderListCancelServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockClient

	requestID         string
	symbol            string
	orderListID       int64
	listClientOrderID string

	orderListCancel        *OrderListCancelWsService
	orderListCancelRequest *OrderListCancelWsRequest
}

func TestOrderListCancelServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderListCancelServiceWsTestSuite))
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancel() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderListCancel.Do(s.requestID, s.orderListCancelRequest)
	s.NoError(err)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancel_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderListCancel.Do("", s.orderListCancelRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancel_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListCancel.Do(s.requestID, s.orderListCancelRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancel_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListCancel.Do(s.requestID, s.orderListCancelRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancel_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListCancel.Do(s.requestID, s.orderListCancelRequest)
	s.Error(err)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancelSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderListCancelResponse := CancelOrderListWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CancelOrderListResult{
			OrderListId:       s.orderListID,
			ContingencyType:   "OCO",
			ListStatusType:    "ALL_DONE",
			ListOrderStatus:   "ALL_DONE",
			ListClientOrderId: s.listClientOrderID,
			TransactionTime:   1660801720215,
			Symbol:            s.symbol,
			Orders: []struct {
				Symbol        string `json:"symbol"`
				OrderId       int64  `json:"orderId"`
				ClientOrderId string `json:"clientOrderId"`
			}{
				{Symbol: s.symbol, OrderId: 12569138901, ClientOrderId: "BqtFCj5odMoWtSqGk2X9tU"},
				{Symbol: s.symbol, OrderId: 12569138902, ClientOrderId: "jLnZpj5enfMXTuhKB1d0us"},
			},
		},
	}

	rawResponseData, err := json.Marshal(orderListCancelResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderListCancelRequest
	response, err := s.orderListCancel.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(s.orderListID, response.Result.OrderListId)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal("ALL_DONE", response.Result.ListStatusType)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancelSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderListCancelRequest
	response, err := s.orderListCancel.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancelSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListCancel.SyncDo(s.requestID, s.orderListCancelRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancelSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListCancel.SyncDo(s.requestID, s.orderListCancelRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListCancelServiceWsTestSuite) TestOrderListCancelSync_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListCancel.SyncDo(s.requestID, s.orderListCancelRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderListCancelServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderListCancel = &OrderListCancelWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
