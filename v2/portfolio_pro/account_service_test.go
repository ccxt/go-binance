package portfolio_pro

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
		"actualEquity": "142607.35137903",
		"accountMaintMargin": "23.72469206",
		"accountInitialMargin": "47.44938412",
		"totalAvailableBalance": "122559.90199491",
		"accountStatus": "NORMAL",
		"accountType": "PM_1"
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
		UniMMR:                "5167.92171923",
		AccountEquity:         "122607.35137903",
		ActualEquity:          "142607.35137903",
		AccountMaintMargin:    "23.72469206",
		AccountInitialMargin:  "47.44938412",
		TotalAvailableBalance: "122559.90199491",
		AccountStatus:         "NORMAL",
		AccountType:           "PM_1",
	})
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *Account) {
	r := s.r()
	r.Equal(e.UniMMR, a.UniMMR, "UniMMR")
	r.Equal(e.AccountEquity, a.AccountEquity, "AccountEquity")
	r.Equal(e.ActualEquity, a.ActualEquity, "ActualEquity")
	r.Equal(e.AccountMaintMargin, a.AccountMaintMargin, "AccountMaintMargin")
	r.Equal(e.AccountInitialMargin, a.AccountInitialMargin, "AccountInitialMargin")
	r.Equal(e.TotalAvailableBalance, a.TotalAvailableBalance, "TotalAvailableBalance")
	r.Equal(e.AccountStatus, a.AccountStatus, "AccountStatus")
	r.Equal(e.AccountType, a.AccountType, "AccountType")
}
