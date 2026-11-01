package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type balanceServiceTestSuite struct {
	baseTestSuite
}

func TestBalanceService(t *testing.T) {
	suite.Run(t, new(balanceServiceTestSuite))
}

func (s *balanceServiceTestSuite) TestGetBalance() {
	data := []byte(`[
		{
			"asset": "ETH",
			"totalWalletBalance": "0.00057786",
			"crossMarginAsset": "0.0",
			"crossMarginBorrowed": "0.0",
			"crossMarginFree": "0.0",
			"crossMarginInterest": "0.0",
			"crossMarginLocked": "0.0",
			"umWalletBalance": "0.0",
			"umUnrealizedPNL": "0.0",
			"cmWalletBalance": "0.00057786",
			"cmUnrealizedPNL": "0.0",
			"updateTime": 1708067615305,
			"negativeBalance": "0.0"
		},
		{
			"asset": "USDT",
			"totalWalletBalance": "136.03451668",
			"crossMarginAsset": "30.6003261",
			"crossMarginBorrowed": "0.00011146",
			"crossMarginFree": "30.6003261",
			"crossMarginInterest": "0.00009623",
			"crossMarginLocked": "0.0",
			"umWalletBalance": "105.43419058",
			"umUnrealizedPNL": "0.0",
			"cmWalletBalance": "0.0",
			"cmUnrealizedPNL": "0.0",
			"updateTime": 1708064880778,
			"negativeBalance": "0.0"
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
	s.r().Len(res, 2)

	expectedBalances := []*Balance{
		{
			Asset:               "ETH",
			TotalWalletBalance:  "0.00057786",
			CrossMarginAsset:    "0.0",
			CrossMarginBorrowed: "0.0",
			CrossMarginFree:     "0.0",
			CrossMarginInterest: "0.0",
			CrossMarginLocked:   "0.0",
			UMWalletBalance:     "0.0",
			UMUnrealizedPNL:     "0.0",
			CMWalletBalance:     "0.00057786",
			CMUnrealizedPNL:     "0.0",
			UpdateTime:          1708067615305,
			NegativeBalance:     "0.0",
		},
		{
			Asset:               "USDT",
			TotalWalletBalance:  "136.03451668",
			CrossMarginAsset:    "30.6003261",
			CrossMarginBorrowed: "0.00011146",
			CrossMarginFree:     "30.6003261",
			CrossMarginInterest: "0.00009623",
			CrossMarginLocked:   "0.0",
			UMWalletBalance:     "105.43419058",
			UMUnrealizedPNL:     "0.0",
			CMWalletBalance:     "0.0",
			CMUnrealizedPNL:     "0.0",
			UpdateTime:          1708064880778,
			NegativeBalance:     "0.0",
		},
	}

	for i := range expectedBalances {
		s.assertBalanceEqual(expectedBalances[i], res[i])
	}
}

func (s *balanceServiceTestSuite) assertBalanceEqual(e, a *Balance) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.TotalWalletBalance, a.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(e.CrossMarginAsset, a.CrossMarginAsset, "CrossMarginAsset")
	r.Equal(e.CrossMarginBorrowed, a.CrossMarginBorrowed, "CrossMarginBorrowed")
	r.Equal(e.CrossMarginFree, a.CrossMarginFree, "CrossMarginFree")
	r.Equal(e.CrossMarginInterest, a.CrossMarginInterest, "CrossMarginInterest")
	r.Equal(e.CrossMarginLocked, a.CrossMarginLocked, "CrossMarginLocked")
	r.Equal(e.UMWalletBalance, a.UMWalletBalance, "UMWalletBalance")
	r.Equal(e.UMUnrealizedPNL, a.UMUnrealizedPNL, "UMUnrealizedPNL")
	r.Equal(e.CMWalletBalance, a.CMWalletBalance, "CMWalletBalance")
	r.Equal(e.CMUnrealizedPNL, a.CMUnrealizedPNL, "CMUnrealizedPNL")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.NegativeBalance, a.NegativeBalance, "NegativeBalance")
}
