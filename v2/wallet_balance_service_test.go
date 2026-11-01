package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type walletBalanceServiceTestSuite struct {
	baseTestSuite
}

func TestWalletBalanceService(t *testing.T) {
	suite.Run(t, new(walletBalanceServiceTestSuite))
}

func (s *walletBalanceServiceTestSuite) TestWalletBalance() {
	data := []byte(`[
	{
		"activate": true,
		"balance": "0.00000551",
		"walletName": "Spot"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "Funding"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "Cross Margin"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "Isolated Margin"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "USDⓈ-M Futures"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "COIN-M Futures"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "Earn"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "Options"
	},
	{
		"activate": true,
		"balance": "0",
		"walletName": "Trading Bots"
	},
	{
		"activate": true,
		"balance": "0.01630125",
		"walletName": "Copy Trading"
	}
]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	quoteAsset := "USDT"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("quoteAsset", quoteAsset)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewWalletBalanceService().QuoteAsset(quoteAsset).Do(newContext())
	s.r().NoError(err)

	e := []*WalletBalance{
		{
			Activate:   true,
			Balance:    "0.00000551",
			WalletName: "Spot",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "Funding",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "Cross Margin",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "Isolated Margin",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "USDⓈ-M Futures",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "COIN-M Futures",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "Earn",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "Options",
		},
		{
			Activate:   true,
			Balance:    "0",
			WalletName: "Trading Bots",
		},
		{
			Activate:   true,
			Balance:    "0.01630125",
			WalletName: "Copy Trading",
		},
	}
	s.assertWalletBalancesEqual(e, res)
}

func (s *walletBalanceServiceTestSuite) assertWalletBalanceEqual(e, a *WalletBalance) {
	r := s.r()
	r.Equal(e.Activate, a.Activate, "Activate")
	r.Equal(e.Balance, a.Balance, "Balance")
	r.Equal(e.WalletName, a.WalletName, "WalletName")
}

func (s *walletBalanceServiceTestSuite) assertWalletBalancesEqual(e, a []*WalletBalance) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertWalletBalanceEqual(e[i], a[i])
	}
}
