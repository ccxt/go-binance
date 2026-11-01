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

func (s *orderPlaceServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "e2a85d9f-07a5-4f94-8d5f-789dc3deb098"

	s.symbol = "BTCUSDT"
	s.side = SideTypeSell
	s.orderType = OrderTypeLimit
	s.timeInForce = TimeInForceTypeGTC
	s.quantity = "0.1001"
	s.price = "50000"
	s.newClientOrderID = "testOrder"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderPlace = &OrderCreateWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderPlaceRequest = NewOrderCreateWsRequest().
		Symbol(s.symbol).
		Side(s.side).
		Type(s.orderType).
		TimeInForce(s.timeInForce).
		Quantity(s.quantity).
		Price(s.price).
		NewClientOrderID(s.newClientOrderID)
}

func (s *orderPlaceServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderPlaceServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockClient

	requestID        string
	symbol           string
	side             SideType
	orderType        OrderType
	timeInForce      TimeInForceType
	quantity         string
	price            string
	newClientOrderID string

	orderPlace        *OrderCreateWsService
	orderPlaceRequest *OrderCreateWsRequest
}

func TestOrderPlaceServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderPlaceServiceWsTestSuite))
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.NoError(err)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do("", s.orderPlaceRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlace_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderPlace.Do(s.requestID, s.orderPlaceRequest)
	s.Error(err)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderPlaceResponse := CreateOrderWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CreateOrderResult{
			CreateOrderResponse{
				Symbol:        s.symbol,
				OrderID:       0,
				ClientOrderID: s.newClientOrderID,
				Price:         s.price,
				TimeInForce:   s.timeInForce,
				Type:          s.orderType,
				Side:          s.side,
			},
		},
	}

	rawResponseData, err := json.Marshal(orderPlaceResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderPlaceRequest
	response, err := s.orderPlace.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(*req.newClientOrderID, response.Result.ClientOrderID)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal(req.orderType, response.Result.Type)
	s.Equal(*req.price, response.Result.Price)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderPlaceRequest
	response, err := s.orderPlace.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderPlace.SyncDo(s.requestID, s.orderPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderPlace.SyncDo(s.requestID, s.orderPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderPlaceServiceWsTestSuite) TestOrderPlaceSync_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderPlace.SyncDo(s.requestID, s.orderPlaceRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderPlaceServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderPlace = &OrderCreateWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
