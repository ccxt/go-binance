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

func (s *orderListPlaceOtocoServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "1712544408508"

	s.symbol = "LTCBNB"
	s.workingType = OrderTypeLimit
	s.workingSide = SideTypeBuy
	s.workingPrice = "1.5"
	s.workingQuantity = "1"
	s.pendingSide = SideTypeSell
	s.pendingQuantity = "5"
	s.pendingAboveType = OrderTypeStopLoss
	s.pendingAboveStopPrice = "0.5"
	s.pendingBelowType = OrderTypeLimitMaker
	s.pendingBelowPrice = "5"
	s.listClientOrderID = "testOTOCOList"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderListPlaceOtoco = &OrderListPlaceOtocoWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderListPlaceOtocoRequest = NewOrderListPlaceOtocoWsRequest().
		Symbol(s.symbol).
		WorkingType(s.workingType).
		WorkingSide(s.workingSide).
		WorkingPrice(s.workingPrice).
		WorkingQuantity(s.workingQuantity).
		PendingSide(s.pendingSide).
		PendingQuantity(s.pendingQuantity).
		PendingAboveType(s.pendingAboveType).
		PendingAboveStopPrice(s.pendingAboveStopPrice).
		PendingBelowType(s.pendingBelowType).
		PendingBelowPrice(s.pendingBelowPrice).
		ListClientOrderID(s.listClientOrderID).
		NewOrderRespType(NewOrderRespTypeRESULT)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderListPlaceOtocoServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockClient

	requestID             string
	symbol                string
	workingType           OrderType
	workingSide           SideType
	workingPrice          string
	workingQuantity       string
	pendingSide           SideType
	pendingQuantity       string
	pendingAboveType      OrderType
	pendingAboveStopPrice string
	pendingBelowType      OrderType
	pendingBelowPrice     string
	listClientOrderID     string

	orderListPlaceOtoco        *OrderListPlaceOtocoWsService
	orderListPlaceOtocoRequest *OrderListPlaceOtocoWsRequest
}

func TestOrderListPlaceOtocoServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderListPlaceOtocoServiceWsTestSuite))
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtoco() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderListPlaceOtoco.Do(s.requestID, s.orderListPlaceOtocoRequest)
	s.NoError(err)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtoco_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOtoco.Do("", s.orderListPlaceOtocoRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtoco_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOtoco.Do(s.requestID, s.orderListPlaceOtocoRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtoco_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOtoco.Do(s.requestID, s.orderListPlaceOtocoRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtoco_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOtoco.Do(s.requestID, s.orderListPlaceOtocoRequest)
	s.Error(err)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtocoSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderListPlaceOtocoResponse := CreateOrderListWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CreateOrderListResult{
			OrderListId:       629,
			ContingencyType:   "OTO",
			ListStatusType:    "EXEC_STARTED",
			ListOrderStatus:   "EXECUTING",
			ListClientOrderId: s.listClientOrderID,
			TransactionTime:   1712544408537,
			Symbol:            s.symbol,
			Orders: []struct {
				Symbol        string `json:"symbol"`
				OrderId       int64  `json:"orderId"`
				ClientOrderId string `json:"clientOrderId"`
			}{
				{Symbol: s.symbol, OrderId: 23, ClientOrderId: "OVQOpKwfmPCfaBTD0n7e7H"},
				{Symbol: s.symbol, OrderId: 24, ClientOrderId: "YcCPKCDMQIjNvLtNswt82X"},
				{Symbol: s.symbol, OrderId: 25, ClientOrderId: "ilpIoShcFZ1ZGgSASKxMPt"},
			},
		},
	}

	rawResponseData, err := json.Marshal(orderListPlaceOtocoResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderListPlaceOtocoRequest
	response, err := s.orderListPlaceOtoco.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(*req.listClientOrderID, response.Result.ListClientOrderId)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal("OTO", response.Result.ContingencyType)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtocoSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderListPlaceOtocoRequest
	response, err := s.orderListPlaceOtoco.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtocoSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlaceOtoco.SyncDo(s.requestID, s.orderListPlaceOtocoRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtocoSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlaceOtoco.SyncDo(s.requestID, s.orderListPlaceOtocoRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) TestOrderListPlaceOtocoSync_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlaceOtoco.SyncDo(s.requestID, s.orderListPlaceOtocoRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderListPlaceOtocoServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderListPlaceOtoco = &OrderListPlaceOtocoWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
