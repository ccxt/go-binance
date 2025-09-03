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

func (s *orderListPlaceOtoServiceWsTestSuite) SetupTest() {
	s.apiKey = "dummyApiKey"
	s.secretKey = "dummySecretKey"
	s.signedKey = "HMAC"
	s.timeOffset = 0

	s.requestID = "1712544395950"

	s.symbol = "LTCBNB"
	s.workingType = OrderTypeLimit
	s.workingSide = SideTypeSell
	s.workingPrice = "1.0"
	s.workingQuantity = "1"
	s.pendingType = OrderTypeMarket
	s.pendingSide = SideTypeBuy
	s.pendingQuantity = "1"
	s.listClientOrderID = "testOTOList"

	s.ctrl = gomock.NewController(s.T())
	s.client = mock.NewMockClient(s.ctrl)

	s.orderListPlaceOto = &OrderListPlaceOtoWsService{
		c:         s.client,
		ApiKey:    s.apiKey,
		SecretKey: s.secretKey,
		KeyType:   s.signedKey,
	}

	s.orderListPlaceOtoRequest = NewOrderListPlaceOtoWsRequest().
		Symbol(s.symbol).
		WorkingType(s.workingType).
		WorkingSide(s.workingSide).
		WorkingPrice(s.workingPrice).
		WorkingQuantity(s.workingQuantity).
		PendingType(s.pendingType).
		PendingSide(s.pendingSide).
		PendingQuantity(s.pendingQuantity).
		ListClientOrderID(s.listClientOrderID).
		NewOrderRespType(NewOrderRespTypeRESULT)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

type orderListPlaceOtoServiceWsTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	signedKey  string
	timeOffset int64

	ctrl   *gomock.Controller
	client *mock.MockClient

	requestID         string
	symbol            string
	workingType       OrderType
	workingSide       SideType
	workingPrice      string
	workingQuantity   string
	pendingType       OrderType
	pendingSide       SideType
	pendingQuantity   string
	listClientOrderID string

	orderListPlaceOto        *OrderListPlaceOtoWsService
	orderListPlaceOtoRequest *OrderListPlaceOtoWsRequest
}

func TestOrderListPlaceOtoServiceWsPlace(t *testing.T) {
	suite.Run(t, new(orderListPlaceOtoServiceWsTestSuite))
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOto() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).AnyTimes()

	err := s.orderListPlaceOto.Do(s.requestID, s.orderListPlaceOtoRequest)
	s.NoError(err)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOto_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOto.Do("", s.orderListPlaceOtoRequest)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOto_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOto.Do(s.requestID, s.orderListPlaceOtoRequest)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOto_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOto.Do(s.requestID, s.orderListPlaceOtoRequest)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOto_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().Write(s.requestID, gomock.Any()).Return(nil).Times(0)

	err := s.orderListPlaceOto.Do(s.requestID, s.orderListPlaceOtoRequest)
	s.Error(err)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOtoSync() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	orderListPlaceOtoResponse := CreateOrderListWsResponse{
		Id:     s.requestID,
		Status: 200,
		Result: CreateOrderListResult{
			OrderListId:       626,
			ContingencyType:   "OTO",
			ListStatusType:    "EXEC_STARTED",
			ListOrderStatus:   "EXECUTING",
			ListClientOrderId: s.listClientOrderID,
			TransactionTime:   1712544395981,
			Symbol:            s.symbol,
			Orders: []struct {
				Symbol        string `json:"symbol"`
				OrderId       int64  `json:"orderId"`
				ClientOrderId string `json:"clientOrderId"`
			}{
				{Symbol: s.symbol, OrderId: 13, ClientOrderId: "YiAUtM9yJjl1a2jXHSp9Ny"},
				{Symbol: s.symbol, OrderId: 14, ClientOrderId: "9MxJSE1TYkmyx5lbGLve7R"},
			},
		},
	}

	rawResponseData, err := json.Marshal(orderListPlaceOtoResponse)
	s.NoError(err)

	s.client.EXPECT().WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(rawResponseData, nil).Times(1)

	req := s.orderListPlaceOtoRequest
	response, err := s.orderListPlaceOto.SyncDo(s.requestID, req)
	s.Require().NoError(err)
	s.Equal(*req.listClientOrderID, response.Result.ListClientOrderId)
	s.Equal(req.symbol, response.Result.Symbol)
	s.Equal("OTO", response.Result.ContingencyType)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOtoSync_EmptyRequestID() {
	s.reset(s.apiKey, s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	req := s.orderListPlaceOtoRequest
	response, err := s.orderListPlaceOto.SyncDo("", req)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorRequestIDNotSet)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOtoSync_EmptyApiKey() {
	s.reset("", s.secretKey, s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlaceOto.SyncDo(s.requestID, s.orderListPlaceOtoRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorApiKeyIsNotSet)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOtoSync_EmptySecretKey() {
	s.reset(s.apiKey, "", s.signedKey, s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlaceOto.SyncDo(s.requestID, s.orderListPlaceOtoRequest)
	s.Nil(response)
	s.ErrorIs(err, websocket.ErrorSecretKeyIsNotSet)
}

func (s *orderListPlaceOtoServiceWsTestSuite) TestOrderListPlaceOtoSync_EmptySignKeyType() {
	s.reset(s.apiKey, s.secretKey, "", s.timeOffset)

	s.client.EXPECT().
		WriteSync(s.requestID, gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("write sync: error")).Times(0)

	response, err := s.orderListPlaceOto.SyncDo(s.requestID, s.orderListPlaceOtoRequest)
	s.Nil(response)
	s.Error(err)
}

func (s *orderListPlaceOtoServiceWsTestSuite) reset(apiKey, secretKey, signKeyType string, timeOffset int64) {
	s.orderListPlaceOto = &OrderListPlaceOtoWsService{
		c:          s.client,
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		KeyType:    signKeyType,
		TimeOffset: timeOffset,
	}
}
