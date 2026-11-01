package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

func (s *AccountServiceTestSuite) TestAccount() {
	data := []byte(`{
		"asset": [
		 {
		  "asset": "USDT",
		  "marginBalance": "1.73304384",
		  "equity": "1.73304384",
		  "available": "1.73304384",
		  "locked": "0",
		  "unrealizedPNL": "0"
		 }
		],
		"greek":[{
			"underlying": "BTC",
			"delta": "1",
			"gamma": "0.01",
			"theta": "0.001",
			"vega": "0.0001"
		}],
		"time": 1717037493533,
		"riskLevel": "NORMAL"
	   }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	account, err := s.client.NewAccountService().Do(newContext())
	targetAccount := &Account{
		Asset: []*Asset{
			{
				Asset:         "USDT",
				MarginBalance: "1.73304384",
				Equity:        "1.73304384",
				Available:     "1.73304384",
				Locked:        "0",
				UnrealizedPNL: "0",
			},
		},
		Greek: []*Greek{
			{
				Underlying: "BTC",
				Delta:      "1",
				Gamma:      "0.01",
				Theta:      "0.001",
				Vega:       "0.0001",
			},
		},
		Time:      1717037493533,
		RiskLevel: "NORMAL",
	}

	s.r().Equal(err, nil, "err != nil")

	r := s.r()
	for i := range account.Asset {
		r.Equal(account.Asset[i].Asset, targetAccount.Asset[i].Asset, "Asset.Asset")
		r.Equal(account.Asset[i].MarginBalance, targetAccount.Asset[i].MarginBalance, "Asset.MarginBalance")
		r.Equal(account.Asset[i].Equity, targetAccount.Asset[i].Equity, "Asset.Equity")
		r.Equal(account.Asset[i].Available, targetAccount.Asset[i].Available, "Asset.Available")
		r.Equal(account.Asset[i].Locked, targetAccount.Asset[i].Locked, "Asset.Locked")
		r.Equal(account.Asset[i].UnrealizedPNL, targetAccount.Asset[i].UnrealizedPNL, "Asset.UnrealizedPNL")
	}
	for i := range account.Greek {
		r.Equal(account.Greek[i].Underlying, targetAccount.Greek[i].Underlying, "Asset.Underlying")
		r.Equal(account.Greek[i].Delta, targetAccount.Greek[i].Delta, "Asset.Delta")
		r.Equal(account.Greek[i].Gamma, targetAccount.Greek[i].Gamma, "Asset.Gamma")
		r.Equal(account.Greek[i].Theta, targetAccount.Greek[i].Theta, "Asset.Theta")
		r.Equal(account.Greek[i].Vega, targetAccount.Greek[i].Vega, "Asset.Vega")
	}
	r.Equal(account.Time, targetAccount.Time, "Time")
	r.Equal(account.RiskLevel, targetAccount.RiskLevel, "RiskLevel")

}
