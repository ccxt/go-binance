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

func (s *orderListPlaceServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "e2a85d9f-07a5-4f94-8d5f-789dc3deb098"

	s.symbol = "BTCUSDT"
	s.side = SideTypeSell
	s.quantity = "0.1001"
	s.aboveType = OrderTypeStopLossLimit
	s.belowType = OrderTypeLimitMaker
	s.abovePrice = "50000"
	s.aboveStopPrice = "49000"
	s.belowPrice = "48000"
	s.listClientOrderID = "testOCOList"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderListPlace = &OrderListCreateWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderListPlaceRequest = NewOrderListCreateWsRequest().
		Symbol(s.symbol).
		Side(s.side).
		Quantity(s.quantity).
		AboveType(s.aboveType).
		BelowType(s.belowType).
		AbovePrice(s.abovePrice).
		AboveStopPrice(s.aboveStopPrice).
		BelowPrice(s.belowPrice).
		ListClientOrderID(s.listClientOrderID).
		NewOrderRespType(NewOrderRespTypeRESULT)
}

func (s *orderListPlaceServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderListPlaceServiceWsTestSuite struct {
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
	quantity          string
	aboveType         OrderType
	belowType         OrderType
	abovePrice        string
	aboveStopPrice    string
	belowPrice        string
	listClientOrderID string

	orderListPlace        *OrderListCreateWsService
	orderListPlaceRequest *OrderListCreateWsRequest
}

func TestOrderListPlaceServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderListPlaceServiceWsTestSuite))
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlace() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.NoError(err)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlace_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do("", s.orderListPlaceRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlace_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlace_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlace_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlace.Do(s.requestID, s.orderListPlaceRequest)
	s.Error(err)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlaceSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderListPlaceResponse := CreateOrderListWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CreateOrderListResult{
			OrderListId:       12345,
			ContingencyType:   "OCO",
			ListStatusType:    "EXEC_STARTED",
			ListOrderStatus:   "EXECUTING",
			ListClientOrderId: s.listClientOrderID,
			TransactionTime:   1660801715639,
			Symbol:            s.symbol,
			Orders: []struct {
				Symbol        string `json:"symbol"`
				OrderId       int64  `json:"orderId"`
				ClientOrderId string `json:"clientOrderId"`
			}{
				{Symbol: s.symbol, OrderId: 12345, ClientOrderId: "limit_order_id"},
				{Symbol: s.symbol, OrderId: 12346, ClientOrderId: "stop_order_id"},
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

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlaceSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderListPlaceRequest
	response, err := s.orderListPlace.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlaceSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlace.SyncDo(s.requestID, s.orderListPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceServiceWsTestSuite) TestOrderListPlaceSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlace.SyncDo(s.requestID, s.orderListPlaceRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderListPlace = &OrderListCreateWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
