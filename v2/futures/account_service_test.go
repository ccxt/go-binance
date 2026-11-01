package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type accountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (s *accountServiceTestSuite) TestGetBalance() {
	data := []byte(`[
		{
			"accountAlias": "SgsR",
			"asset": "USDT",
			"balance": "122607.35137903",
			"crossWalletBalance": "23.72469206",
			"crossUnPnl": "0.00000000",
			"availableBalance": "23.72469206",
			"maxWithdrawAmount": "23.72469206",
		    "marginAvailable": true,
            "updateTime": 1617939110373
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetBalanceService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	e := &Balance{
		AccountAlias:       "SgsR",
		Asset:              "USDT",
		Balance:            "122607.35137903",
		CrossWalletBalance: "23.72469206",
		CrossUnPnl:         "0.00000000",
		AvailableBalance:   "23.72469206",
		MaxWithdrawAmount:  "23.72469206",
		MarginAvailable:    true,
		UpdateTime:         1617939110373,
	}
	s.assertBalanceEqual(e, res[0])
}

func (s *accountServiceTestSuite) assertBalanceEqual(e, a *Balance) {
	r := s.r()
	r.Equal(e.AccountAlias, a.AccountAlias, "AccountAlias")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Balance, a.Balance, "Balance")
	r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	r.Equal(e.CrossUnPnl, a.CrossUnPnl, "CrossUnPnl")
	r.Equal(e.AvailableBalance, a.AvailableBalance, "AvailableBalance")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
}

func (s *accountServiceTestSuite) TestGetAccount() {
	data := []byte(`{
		"assets": [
			{
				"asset": "USDT",
				"initialMargin": "0.33683000",
				"maintMargin": "0.02695000",
				"marginBalance": "8.74947592",
				"maxWithdrawAmount": "8.41264592",
				"openOrderInitialMargin": "0.00000000",
				"positionInitialMargin": "0.33683000",
				"unrealizedProfit": "-0.44537584",
				"walletBalance": "9.19485176",
				"crossWalletBalance": "23.72469206",
				"crossUnPnl": "0.00000000",
				"availableBalance": "126.72469206",
				"marginAvailable": true,
				"updateTime": 1625474304765
			}
		 ],
		 "canDeposit": true,
		 "canTrade": true,
		 "canWithdraw": true,
		 "feeTier": 2,
		 "maxWithdrawAmount": "8.41264592",
		 "multiAssetsMargin": false,
		 "positions": [
			 {
				"isolated": false, 
				"leverage": "20",
				"initialMargin": "0.33683",
				"maintMargin": "0.02695",
				"openOrderInitialMargin": "0.00000",
				"positionInitialMargin": "0.33683",
				"symbol": "BTCUSDT",
				"unrealizedProfit": "-0.44537584",
				"entryPrice": "8950.5",
				"maxNotional": "250000",
				"positionSide": "BOTH",
				"positionAmt": "0.436",
				"bidNotional": "0",
				"askNotional": "0",
				"updateTime":1618646402359
			 }
		 ],
		 "totalInitialMargin": "0.33683000",
		 "totalMaintMargin": "0.02695000",
		 "totalMarginBalance": "8.74947592",
		 "totalOpenOrderInitialMargin": "0.00000000",
		 "totalPositionInitialMargin": "0.33683000",
		 "totalUnrealizedProfit": "-0.44537584",
		 "totalWalletBalance": "9.19485176",
		 "updateTime": 0
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAccountService().Do(newContext())
	s.r().NoError(err)
	e := &Account{
		Assets: []*AccountAsset{
			{
				Asset:                  "USDT",
				InitialMargin:          "0.33683000",
				MaintMargin:            "0.02695000",
				MarginBalance:          "8.74947592",
				MaxWithdrawAmount:      "8.41264592",
				OpenOrderInitialMargin: "0.00000000",
				PositionInitialMargin:  "0.33683000",
				UnrealizedProfit:       "-0.44537584",
				WalletBalance:          "9.19485176",
				CrossWalletBalance:     "23.72469206",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "126.72469206",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
		},
		CanTrade:          true,
		CanWithdraw:       true,
		CanDeposit:        true,
		FeeTier:           2,
		MaxWithdrawAmount: "8.41264592",
		MultiAssetsMargin: false,
		Positions: []*AccountPosition{
			{
				Isolated:               false,
				Leverage:               "20",
				InitialMargin:          "0.33683",
				MaintMargin:            "0.02695",
				OpenOrderInitialMargin: "0.00000",
				PositionInitialMargin:  "0.33683",
				Symbol:                 "BTCUSDT",
				UnrealizedProfit:       "-0.44537584",
				EntryPrice:             "8950.5",
				MaxNotional:            "250000",
				PositionSide:           "BOTH",
				PositionAmt:            "0.436",
				BidNotional:            "0",
				AskNotional:            "0",
				UpdateTime:             1618646402359,
			},
		},
		TotalInitialMargin:          "0.33683000",
		TotalMaintMargin:            "0.02695000",
		TotalMarginBalance:          "8.74947592",
		TotalOpenOrderInitialMargin: "0.00000000",
		TotalPositionInitialMargin:  "0.33683000",
		TotalUnrealizedProfit:       "-0.44537584",
		TotalWalletBalance:          "9.19485176",
		UpdateTime:                  0,
	}
	s.assertAccountEqual(e, res)
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *Account) {
	r := s.r()
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.FeeTier, a.FeeTier, "FeeTier")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal(e.TotalInitialMargin, a.TotalInitialMargin, "TotalInitialMargin")
	r.Equal(e.TotalMaintMargin, a.TotalMaintMargin, "TotalMaintMargin")
	r.Equal(e.TotalMarginBalance, a.TotalMarginBalance, "TotalMarginBalance")
	r.Equal(e.TotalOpenOrderInitialMargin, a.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal(e.TotalPositionInitialMargin, a.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal(e.TotalUnrealizedProfit, a.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal(e.TotalWalletBalance, a.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.MultiAssetsMargin, a.MultiAssetsMargin, "MultiAssetsMargin")

	r.Len(a.Assets, len(e.Assets))
	for i := 0; i < len(a.Assets); i++ {
		r.Equal(e.Assets[i].Asset, a.Assets[i].Asset, "Asset")
		r.Equal(e.Assets[i].InitialMargin, a.Assets[i].InitialMargin, "InitialMargin")
		r.Equal(e.Assets[i].MaintMargin, a.Assets[i].MaintMargin, "MaintMargin")
		r.Equal(e.Assets[i].MarginBalance, a.Assets[i].MarginBalance, "MarginBalance")
		r.Equal(e.Assets[i].MaxWithdrawAmount, a.Assets[i].MaxWithdrawAmount, "MaxWithdrawAmount")
		r.Equal(e.Assets[i].OpenOrderInitialMargin, a.Assets[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(e.Assets[i].PositionInitialMargin, a.Assets[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(e.Assets[i].UnrealizedProfit, a.Assets[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(e.Assets[i].WalletBalance, a.Assets[i].WalletBalance, "WalletBalance")
		r.Equal(e.Assets[i].CrossWalletBalance, a.Assets[i].CrossWalletBalance, "CrossWalletBalance")
		r.Equal(e.Assets[i].CrossUnPnl, a.Assets[i].CrossUnPnl, "CrossUnPnl")
		r.Equal(e.Assets[i].AvailableBalance, a.Assets[i].AvailableBalance, "AvailableBalance")
		r.Equal(e.Assets[i].MarginAvailable, a.Assets[i].MarginAvailable, "MarginAvailable")
		r.Equal(e.Assets[i].UpdateTime, a.Assets[i].UpdateTime, "UpdateTime")
	}

	r.Len(a.Positions, len(e.Positions))
	for i := 0; i < len(a.Positions); i++ {
		r.Equal(e.Positions[i].Isolated, a.Positions[i].Isolated, "Isolated")
		r.Equal(e.Positions[i].Leverage, a.Positions[i].Leverage, "Leverage")
		r.Equal(e.Positions[i].InitialMargin, a.Positions[i].InitialMargin, "InitialMargin")
		r.Equal(e.Positions[i].MaintMargin, a.Positions[i].MaintMargin, "MaintMargin")
		r.Equal(e.Positions[i].OpenOrderInitialMargin, a.Positions[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(e.Positions[i].PositionInitialMargin, a.Positions[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(e.Positions[i].Symbol, a.Positions[i].Symbol, "Symbol")
		r.Equal(e.Positions[i].UnrealizedProfit, a.Positions[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(e.Positions[i].EntryPrice, a.Positions[i].EntryPrice, "EntryPrice")
		r.Equal(e.Positions[i].MaxNotional, a.Positions[i].MaxNotional, "MaxNotional")
		r.Equal(e.Positions[i].PositionSide, a.Positions[i].PositionSide, "PositionSide")
		r.Equal(e.Positions[i].PositionAmt, a.Positions[i].PositionAmt, "PositionAmt")
		r.Equal(e.Positions[i].BidNotional, a.Positions[i].BidNotional, "BidNotional")
		r.Equal(e.Positions[i].AskNotional, a.Positions[i].AskNotional, "AskNotional")
		r.Equal(e.Positions[i].UpdateTime, a.Positions[i].UpdateTime, "UpdateTime")
	}
}

func (s *accountServiceTestSuite) TestGetAccountV3() {
	data := []byte(`{
		"totalInitialMargin": "0.00000000",
		"totalMaintMargin": "0.00000000",
		"totalWalletBalance": "126.72469206",
		"totalUnrealizedProfit": "0.00000000",
		"totalMarginBalance": "126.72469206",
		"totalPositionInitialMargin": "0.00000000",
		"totalOpenOrderInitialMargin": "0.00000000",
		"totalCrossWalletBalance": "126.72469206",
		"totalCrossUnPnl": "0.00000000",
		"availableBalance": "126.72469206",
		"maxWithdrawAmount": "126.72469206",
		"assets": [
			{
				"asset": "USDT",
				"walletBalance": "23.72469206",
				"unrealizedProfit": "0.00000000",
				"marginBalance": "23.72469206",
				"maintMargin": "0.00000000",
				"initialMargin": "0.00000000",
				"positionInitialMargin": "0.00000000",
				"openOrderInitialMargin": "0.00000000",
				"crossWalletBalance": "23.72469206",
				"crossUnPnl": "0.00000000",
				"availableBalance": "126.72469206",
				"maxWithdrawAmount": "23.72469206",
				"marginAvailable": true,
				"updateTime": 1625474304765
			},
			{
				"asset": "BUSD",
				"walletBalance": "103.12345678",
				"unrealizedProfit": "0.00000000",
				"marginBalance": "103.12345678",
				"maintMargin": "0.00000000",
				"initialMargin": "0.00000000",
				"positionInitialMargin": "0.00000000",
				"openOrderInitialMargin": "0.00000000",
				"crossWalletBalance": "103.12345678",
				"crossUnPnl": "0.00000000",
				"availableBalance": "126.72469206",
				"maxWithdrawAmount": "103.12345678",
				"marginAvailable": true,
				"updateTime": 1625474304765
			}
		],
		"positions": [
			{
				"symbol": "BTCUSDT",
				"positionSide": "BOTH",
				"positionAmt": "1.000",
				"unrealizedProfit": "0.00000000",
				"isolatedMargin": "0.00000000",
				"notional": "0",
				"isolatedWallet": "0",
				"initialMargin": "0",
				"maintMargin": "0",
				"updateTime": 0
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		expected := newSignedRequest()
		s.assertRequestEqual(expected, r)
	})

	res, err := s.client.NewGetAccountV3Service().Do(newContext())
	s.r().NoError(err)

	expected := &AccountV3{
		TotalInitialMargin:          "0.00000000",
		TotalMaintMargin:            "0.00000000",
		TotalWalletBalance:          "126.72469206",
		TotalUnrealizedProfit:       "0.00000000",
		TotalMarginBalance:          "126.72469206",
		TotalPositionInitialMargin:  "0.00000000",
		TotalOpenOrderInitialMargin: "0.00000000",
		TotalCrossWalletBalance:     "126.72469206",
		TotalCrossUnPnl:             "0.00000000",
		AvailableBalance:            "126.72469206",
		MaxWithdrawAmount:           "126.72469206",
		Assets: []*AccountAssetV3{
			{
				Asset:                  "USDT",
				WalletBalance:          "23.72469206",
				UnrealizedProfit:       "0.00000000",
				MarginBalance:          "23.72469206",
				MaintMargin:            "0.00000000",
				InitialMargin:          "0.00000000",
				PositionInitialMargin:  "0.00000000",
				OpenOrderInitialMargin: "0.00000000",
				CrossWalletBalance:     "23.72469206",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "126.72469206",
				MaxWithdrawAmount:      "23.72469206",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
			{
				Asset:                  "BUSD",
				WalletBalance:          "103.12345678",
				UnrealizedProfit:       "0.00000000",
				MarginBalance:          "103.12345678",
				MaintMargin:            "0.00000000",
				InitialMargin:          "0.00000000",
				PositionInitialMargin:  "0.00000000",
				OpenOrderInitialMargin: "0.00000000",
				CrossWalletBalance:     "103.12345678",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "126.72469206",
				MaxWithdrawAmount:      "103.12345678",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
		},
		Positions: []*AccountPositionV3{
			{
				Symbol:           "BTCUSDT",
				PositionSide:     "BOTH",
				PositionAmt:      "1.000",
				UnrealizedProfit: "0.00000000",
				IsolatedMargin:   "0.00000000",
				Notional:         "0",
				IsolatedWallet:   "0",
				InitialMargin:    "0",
				MaintMargin:      "0",
				UpdateTime:       0,
			},
		},
	}

	s.assertAccountV3Equal(expected, res)
}

func (s *accountServiceTestSuite) assertAccountV3Equal(expected, actual *AccountV3) {
	r := s.r()
	r.Equal(expected.TotalInitialMargin, actual.TotalInitialMargin, "TotalInitialMargin")
	r.Equal(expected.TotalMaintMargin, actual.TotalMaintMargin, "TotalMaintMargin")
	r.Equal(expected.TotalWalletBalance, actual.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(expected.TotalUnrealizedProfit, actual.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal(expected.TotalMarginBalance, actual.TotalMarginBalance, "TotalMarginBalance")
	r.Equal(expected.TotalPositionInitialMargin, actual.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal(expected.TotalOpenOrderInitialMargin, actual.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal(expected.TotalCrossWalletBalance, actual.TotalCrossWalletBalance, "TotalCrossWalletBalance")
	r.Equal(expected.TotalCrossUnPnl, actual.TotalCrossUnPnl, "TotalCrossUnPnl")
	r.Equal(expected.AvailableBalance, actual.AvailableBalance, "AvailableBalance")
	r.Equal(expected.MaxWithdrawAmount, actual.MaxWithdrawAmount, "MaxWithdrawAmount")

	r.Len(actual.Assets, len(expected.Assets))
	for i := 0; i < len(expected.Assets); i++ {
		r.Equal(expected.Assets[i].Asset, actual.Assets[i].Asset, "Asset")
		r.Equal(expected.Assets[i].WalletBalance, actual.Assets[i].WalletBalance, "WalletBalance")
		r.Equal(expected.Assets[i].UnrealizedProfit, actual.Assets[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(expected.Assets[i].MarginBalance, actual.Assets[i].MarginBalance, "MarginBalance")
		r.Equal(expected.Assets[i].MaintMargin, actual.Assets[i].MaintMargin, "MaintMargin")
		r.Equal(expected.Assets[i].InitialMargin, actual.Assets[i].InitialMargin, "InitialMargin")
		r.Equal(expected.Assets[i].PositionInitialMargin, actual.Assets[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(expected.Assets[i].OpenOrderInitialMargin, actual.Assets[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(expected.Assets[i].CrossWalletBalance, actual.Assets[i].CrossWalletBalance, "CrossWalletBalance")
		r.Equal(expected.Assets[i].CrossUnPnl, actual.Assets[i].CrossUnPnl, "CrossUnPnl")
		r.Equal(expected.Assets[i].AvailableBalance, actual.Assets[i].AvailableBalance, "AvailableBalance")
		r.Equal(expected.Assets[i].MaxWithdrawAmount, actual.Assets[i].MaxWithdrawAmount, "MaxWithdrawAmount")
		r.Equal(expected.Assets[i].MarginAvailable, actual.Assets[i].MarginAvailable, "MarginAvailable")
		r.Equal(expected.Assets[i].UpdateTime, actual.Assets[i].UpdateTime, "UpdateTime")
	}

	r.Len(actual.Positions, len(expected.Positions))
	for i := 0; i < len(expected.Positions); i++ {
		r.Equal(expected.Positions[i].Symbol, actual.Positions[i].Symbol, "Symbol")
		r.Equal(expected.Positions[i].PositionSide, actual.Positions[i].PositionSide, "PositionSide")
		r.Equal(expected.Positions[i].PositionAmt, actual.Positions[i].PositionAmt, "PositionAmt")
		r.Equal(expected.Positions[i].UnrealizedProfit, actual.Positions[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(expected.Positions[i].IsolatedMargin, actual.Positions[i].IsolatedMargin, "IsolatedMargin")
		r.Equal(expected.Positions[i].Notional, actual.Positions[i].Notional, "Notional")
		r.Equal(expected.Positions[i].IsolatedWallet, actual.Positions[i].IsolatedWallet, "IsolatedWallet")
		r.Equal(expected.Positions[i].InitialMargin, actual.Positions[i].InitialMargin, "InitialMargin")
		r.Equal(expected.Positions[i].MaintMargin, actual.Positions[i].MaintMargin, "MaintMargin")
		r.Equal(expected.Positions[i].UpdateTime, actual.Positions[i].UpdateTime, "UpdateTime")
	}
}
