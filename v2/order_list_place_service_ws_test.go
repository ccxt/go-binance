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

func (s *orderListPlaceDeprecatedServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "e2a85d9f-07a5-4f94-8d5f-789dc3deb098"

	s.symbol = "BTCUSDT"
	s.side = SideTypeSell
	s.price = "23420.00000000"
	s.quantity = "0.00650000"
	s.stopPrice = "23410.00000000"
	s.listClientOrderID = "testOCOList"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderListPlace = &OrderListPlaceWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderListPlaceRequest = NewOrderListPlaceWsRequest().
		Symbol(s.symbol).
		Side(s.side).
		Price(s.price).
		Quantity(s.quantity).
		StopPrice(s.stopPrice).
		ListClientOrderID(s.listClientOrderID).
		NewOrderRespType(NewOrderRespTypeRESULT)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderListPlaceDeprecatedServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockClient

	requestID         string
	symbol            string
	side              SideType
	price             string
	quantity          string
	stopPrice         string
	listClientOrderID string

	orderListPlace        *OrderListPlaceWsService
	orderListPlaceRequest *OrderListPlaceWsRequest
}

func TestOrderListPlaceDeprecatedServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderListPlaceDeprecatedServiceWsTestSuite))
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlace() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.NoError(err)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlace_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do("", s.orderListPlaceRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlace_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlace_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlace_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.Error(err)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlaceSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderListPlaceResponse := CreateOrderListWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CreateOrderListResult{
			OrderListId:       1274512,
			ContingencyType:   "OCO",
			ListStatusType:    "EXEC_STARTED",
			ListOrderStatus:   "EXECUTING",
			ListClientOrderId: s.listClientOrderID,
			TransactionTime:   1660801713793,
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

	rawResponseData, err := json.Marshal(orderListPlaceResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderListPlaceRequest
	response, err := s.orderListPlace.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(*req.listClientOrderID, response.Result.ListClientOrderId)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal("OCO", response.Result.ContingencyType)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlaceSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderListPlaceRequest
	response, err := s.orderListPlace.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlaceSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlace.SyncDo(s.requestID, s.orderListPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlaceSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlace.SyncDo(s.requestID, s.orderListPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) TestOrderListPlaceSync_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlace.SyncDo(s.requestID, s.orderListPlaceRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderListPlaceDeprecatedServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderListPlace = &OrderListPlaceWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
