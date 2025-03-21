package portfolio

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

func (s *accountServiceTestSuite) TestGetAccount() {
	data := []byte(`{
		"uniMMR": "5167.92171923",
		"accountEquity": "122607.35137903",
		"actualEquity": "73.47428058",
		"accountInitialMargin": "23.72469206",
		"accountMaintMargin": "23.72469206",
		"accountStatus": "NORMAL",
		"virtualMaxWithdrawAmount": "1627523.32459208",
		"totalAvailableBalance": "100.00",
		"totalMarginOpenLoss": "0.00",
		"updateTime": 1657707212154
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAccountService().Do(newContext())
	s.r().NoError(err)
	s.assertAccountEqual(res, &Account{
		UniMMR:                   "5167.92171923",
		AccountEquity:            "122607.35137903",
		ActualEquity:             "73.47428058",
		AccountInitialMargin:     "23.72469206",
		AccountMaintMargin:       "23.72469206",
		AccountStatus:            "NORMAL",
		VirtualMaxWithdrawAmount: "1627523.32459208",
		TotalAvailableBalance:    "100.00",
		TotalMarginOpenLoss:      "0.00",
		UpdateTime:               1657707212154,
	})
}

func (s *accountServiceTestSuite) assertAccountEqual(a, e *Account) {
	r := s.r()
	r.Equal(e.UniMMR, a.UniMMR, "UniMMR")
	r.Equal(e.AccountEquity, a.AccountEquity, "AccountEquity")
	r.Equal(e.ActualEquity, a.ActualEquity, "ActualEquity")
	r.Equal(e.AccountInitialMargin, a.AccountInitialMargin, "AccountInitialMargin")
	r.Equal(e.AccountMaintMargin, a.AccountMaintMargin, "AccountMaintMargin")
	r.Equal(e.AccountStatus, a.AccountStatus, "AccountStatus")
	r.Equal(e.VirtualMaxWithdrawAmount, a.VirtualMaxWithdrawAmount, "VirtualMaxWithdrawAmount")
	r.Equal(e.TotalAvailableBalance, a.TotalAvailableBalance, "TotalAvailableBalance")
	r.Equal(e.TotalMarginOpenLoss, a.TotalMarginOpenLoss, "TotalMarginOpenLoss")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
