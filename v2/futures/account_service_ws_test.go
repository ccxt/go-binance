package futures

import (
	"testing"

	"github.com/adshao/go-binance/v2/common"
	"github.com/adshao/go-binance/v2/common/websocket/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type accountWsServiceTestSuite struct {
	suite.Suite
	apiKey     string
	secretKey  string
	mockClient *mock.MockClient
	mockCtrl   *gomock.Controller
}

func TestAccountWsService(t *testing.T) {
	suite.Run(t, new(accountWsServiceTestSuite))
}

func (s *accountWsServiceTestSuite) SetupTest() {
	s.apiKey = "dummyAPIKey"
	s.secretKey = "dummySecretKey"
	s.mockCtrl = gomock.NewController(s.T())
	s.mockClient = mock.NewMockClient(s.mockCtrl)
}

func (s *accountWsServiceTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

func (s *accountWsServiceTestSuite) TestGetAccountInfo() {
	data := []byte(`{
  "id": "873a969f-2d36-472a-ace5-0a61c5fc7d39",
  "status": 200,
  "result": {
    "totalInitialMargin": "13.63585026",
    "totalMaintMargin": "0.68179251",
    "totalWalletBalance": "192.32494207",
    "totalUnrealizedProfit": "6.95279735",
    "totalMarginBalance": "199.27773942",
    "totalPositionInitialMargin": "13.63585026",
    "totalOpenOrderInitialMargin": "0.00000000",
    "totalCrossWalletBalance": "192.32494207",
    "totalCrossUnPnl": "6.95279735",
    "availableBalance": "185.64188916",
    "maxWithdrawAmount": "185.64188916",
    "assets": [
      {
        "asset": "USDT",
        "walletBalance": "192.32494207",
        "unrealizedProfit": "6.95279735",
        "marginBalance": "199.27773942",
        "maintMargin": "0.68179251",
        "initialMargin": "13.63585026",
        "positionInitialMargin": "13.63585026",
        "openOrderInitialMargin": "0.00000000",
        "crossWalletBalance": "192.32494207",
        "crossUnPnl": "6.95279735",
        "availableBalance": "185.64188916",
        "maxWithdrawAmount": "185.64188916",
        "updateTime": 1747995976003
      },
      {
        "asset": "USDC",
        "walletBalance": "0.00000000",
        "unrealizedProfit": "0.00000000",
        "marginBalance": "0.00000000",
        "maintMargin": "0.00000000",
        "initialMargin": "0.00000000",
        "positionInitialMargin": "0.00000000",
        "openOrderInitialMargin": "0.00000000",
        "crossWalletBalance": "0.00000000",
        "crossUnPnl": "0.00000000",
        "availableBalance": "0.00000000",
        "maxWithdrawAmount": "0.00000000",
        "updateTime": 0
      }
    ],
    "positions": [
      {
        "symbol": "SOLUSDT",
        "positionSide": "SHORT",
        "positionAmt": "-0.77",
        "unrealizedProfit": "6.95279735",
        "isolatedMargin": "0",
        "notional": "-136.35850264",
        "isolatedWallet": "0",
        "initialMargin": "13.63585026",
        "maintMargin": "0.68179251",
        "updateTime": 1747995976003
      }
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 2400,
      "count": 50
    }
  ]
}`)

	requestID := "873a969f-2d36-472a-ace5-0a61c5fc7d39"
	s.mockClient.EXPECT().WriteSync(requestID, gomock.Any(), gomock.Any()).Return(data, nil)

	wsAccountV2Service := &WsAccountService{
		c:          s.mockClient,
		ApiKey:     s.apiKey,
		SecretKey:  s.secretKey,
		KeyType:    common.KeyTypeHmac,
		RecvWindow: 5000,
	}

	response, err := wsAccountV2Service.SyncGetAccountInfo(requestID)
	s.NoError(err)

	// verify
	s.Equal(200, response.Status)
	s.Equal(requestID, response.ID)

	// verify account info
	s.Equal("13.63585026", response.Result.TotalInitialMargin)
	s.Equal("192.32494207", response.Result.TotalWalletBalance)
	s.Equal("185.64188916", response.Result.AvailableBalance)

	// verify assets info
	s.Len(response.Result.Assets, 2)
	s.Equal("USDT", response.Result.Assets[0].Asset)
	s.Equal("192.32494207", response.Result.Assets[0].WalletBalance)

	// verify positions info
	s.Len(response.Result.Positions, 1)
	s.Equal("SOLUSDT", response.Result.Positions[0].Symbol)
	s.Equal("SHORT", response.Result.Positions[0].PositionSide)
	s.Equal("-0.77", response.Result.Positions[0].PositionAmt)
}

func (s *accountWsServiceTestSuite) TestGetAccountBalance() {
	data := []byte(`{
  "id": "7fe4c481-9784-4c02-8121-aacae6d2d38f",
  "status": 200,
  "result": [
    {
      "accountAlias": "SgsR",
      "asset": "USDT",
      "balance": "192.32494207",
      "crossWalletBalance": "192.32494207",
      "crossUnPnl": "6.86729999",
      "availableBalance": "185.54784206",
      "maxWithdrawAmount": "185.54784206",
      "marginAvailable": true,
      "updateTime": 1747995976003
    },
    {
      "accountAlias": "SgsR",
      "asset": "USDC",
      "balance": "0.00000000",
      "crossWalletBalance": "0.00000000",
      "crossUnPnl": "0.00000000",
      "availableBalance": "0.00000000",
      "maxWithdrawAmount": "0.00000000",
      "marginAvailable": true,
      "updateTime": 0
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 2400,
      "count": 10
    }
  ]
}`)

	requestID := "7fe4c481-9784-4c02-8121-aacae6d2d38f"
	s.mockClient.EXPECT().WriteSync(requestID, gomock.Any(), gomock.Any()).Return(data, nil)

	wsAccountV2Service := &WsAccountService{
		c:          s.mockClient,
		ApiKey:     s.apiKey,
		SecretKey:  s.secretKey,
		KeyType:    common.KeyTypeHmac,
		RecvWindow: 5000,
	}

	response, err := wsAccountV2Service.SyncGetAccountBalance(requestID)
	s.NoError(err)

	// verify
	s.Equal(200, response.Status)
	s.Equal(requestID, response.ID)

	// verify balance info
	s.Len(response.Result, 2)
	s.Equal("USDT", response.Result[0].Asset)
	s.Equal("192.32494207", response.Result[0].Balance)
	s.Equal("185.54784206", response.Result[0].AvailableBalance)
}
