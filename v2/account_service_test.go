package binance

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
			"makerCommission": 15,
			"takerCommission": 15,
			"buyerCommission": 0,
			"sellerCommission": 0,
			"commissionRates": {
				"maker": "0.00150000",
				"taker": "0.00150000",
				"buyer": "0.00000000",
				"seller": "0.00000000"
			},
			"canTrade": true,
			"canWithdraw": true,
			"canDeposit": true,
			"updateTime": 123456789,
			"accountType": "SPOT",
			"balances": [
					{
							"asset": "BTC",
							"free": "4723846.89208129",
							"locked": "0.00000000"
					},
					{
							"asset": "LTC",
							"free": "4763368.68006011",
							"locked": "0.00000000"
					}
			],
			"permissions": [
				"SPOT"
			],
            "uid": 354937868
  }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	omitZeroBalances := true

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"omitZeroBalances": omitZeroBalances,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAccountService().OmitZeroBalances(omitZeroBalances).Do(newContext())
	s.r().NoError(err)
	e := &Account{
		MakerCommission:  15,
		TakerCommission:  15,
		BuyerCommission:  0,
		SellerCommission: 0,
		CommissionRates: CommissionRates{
			Maker:  "0.00150000",
			Taker:  "0.00150000",
			Buyer:  "0.00000000",
			Seller: "0.00000000",
		},
		CanTrade:    true,
		CanWithdraw: true,
		CanDeposit:  true,
		UpdateTime:  123456789,
		AccountType: "SPOT",
		Balances: []Balance{
			{
				Asset:  "BTC",
				Free:   "4723846.89208129",
				Locked: "0.00000000",
			},
			{
				Asset:  "LTC",
				Free:   "4763368.68006011",
				Locked: "0.00000000",
			},
		},
		Permissions: []string{"SPOT"},
		UID:         354937868,
	}
	s.assertAccountEqual(e, res)
}

func (s *accountServiceTestSuite) assertAccountEqual(e, a *Account) {
	r := s.r()
	r.Equal(e.MakerCommission, a.MakerCommission, "MakerCommission")
	r.Equal(e.TakerCommission, a.TakerCommission, "TakerCommission")
	r.Equal(e.BuyerCommission, a.BuyerCommission, "BuyerCommission")
	r.Equal(e.SellerCommission, a.SellerCommission, "SellerCommission")
	r.Equal(e.CommissionRates.Maker, a.CommissionRates.Maker, "CommissionRates.Maker")
	r.Equal(e.CommissionRates.Taker, a.CommissionRates.Taker, "CommissionRates.Taker")
	r.Equal(e.CommissionRates.Buyer, a.CommissionRates.Buyer, "CommissionRates.Buyer")
	r.Equal(e.CommissionRates.Seller, a.CommissionRates.Seller, "CommissionRates.Seller")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Len(a.Balances, len(e.Balances))
	for i := 0; i < len(a.Balances); i++ {
		r.Equal(e.Balances[i].Asset, a.Balances[i].Asset, "Asset")
		r.Equal(e.Balances[i].Free, a.Balances[i].Free, "Free")
		r.Equal(e.Balances[i].Locked, a.Balances[i].Locked, "Locked")
	}
	r.Equal(e.UID, a.UID, "UID")
}

func (s *accountServiceTestSuite) TestGetAccountSnapshot() {
	data := []byte(`{
		"code":200,
		"msg":"",
		"snapshotVos":[
		   {
			  "data":{
				 "balances":[
					{
					   "asset":"BTC",
					   "free":"0.09905021",
					   "locked":"0.00000000"
					}
				 ],
				 "totalAssetOfBtc":"0.09942700"
			  },
			  "type":"spot",
			  "updateTime":1576281599000
		   }
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	accountType := "SPOT"
	startTime := int64(1498793709153)
	endTime := int64(1498793709156)
	limit := 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"type":      accountType,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	accSnapshot, err := s.client.NewGetAccountSnapshotService().Type(accountType).StartTime(startTime).EndTime(endTime).Limit(limit).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Snapshot{
		Code: 200,
		Msg:  "",
		Snapshot: []*SnapshotVos{
			&SnapshotVos{
				Type:       "spot",
				UpdateTime: 1576281599000,
				Data: &SnapshotData{
					TotalAssetOfBtc: "0.09942700",
					Balances: []*SnapshotBalances{
						&SnapshotBalances{
							Asset:  "BTC",
							Free:   "0.09905021",
							Locked: "0.00000000",
						},
					},
				},
			},
		},
	}
	s.assertSnapshotAccountEqual(e, accSnapshot)
}

func (s *accountServiceTestSuite) assertSnapshotAccountEqual(e, a *Snapshot) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
	for i := 0; i < len(a.Snapshot); i++ {
		r.Equal(e.Snapshot[i].Type, a.Snapshot[i].Type, "Type")
		r.Equal(e.Snapshot[i].UpdateTime, a.Snapshot[i].UpdateTime, "UpdateTime")
		r.Equal(e.Snapshot[i].Data.TotalAssetOfBtc, a.Snapshot[i].Data.TotalAssetOfBtc, "TotalAssetOfBtc")
		for j := 0; j < len(a.Snapshot[i].Data.Balances); j++ {
			r.Equal(e.Snapshot[i].Data.Balances[j].Asset, a.Snapshot[i].Data.Balances[j].Asset, "Asset")
			r.Equal(e.Snapshot[i].Data.Balances[j].Free, a.Snapshot[i].Data.Balances[j].Free, "Free")
			r.Equal(e.Snapshot[i].Data.Balances[j].Locked, a.Snapshot[i].Data.Balances[j].Locked, "Locked")
		}
	}
}

func (s *accountServiceTestSuite) TestGetAPIKeyPermission() {
	data := []byte(`{
   			"ipRestrict": false,
   			"createTime": 1623840271000,   
   			"enableWithdrawals": false,
   			"enableInternalTransfer": true,
   			"permitsUniversalTransfer": true,
   			"enableVanillaOptions": false,
   			"enableReading": true,
   			"enableFutures": false,
   			"enableMargin": false,
   			"enableSpotAndMarginTrading": false,
   			"tradingAuthorityExpirationTime": 1628985600000
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAPIKeyPermission().Do(newContext())
	s.r().NoError(err)
	e := &APIKeyPermission{
		IPRestrict:                     false,
		CreateTime:                     1623840271000,
		EnableWithdrawals:              false,
		EnableInternalTransfer:         true,
		PermitsUniversalTransfer:       true,
		EnableVanillaOptions:           false,
		EnableReading:                  true,
		EnableFutures:                  false,
		EnableMargin:                   false,
		EnableSpotAndMarginTrading:     false,
		TradingAuthorityExpirationTime: 1628985600000,
	}
	s.assertAPIKeyPermissionEqual(e, res)
}

func (s *accountServiceTestSuite) assertAPIKeyPermissionEqual(e, a *APIKeyPermission) {
	r := s.r()
	r.Equal(e.IPRestrict, a.IPRestrict, "IPRestrict")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.EnableWithdrawals, a.EnableWithdrawals, "EnableWithdrawals")
	r.Equal(e.EnableInternalTransfer, a.EnableInternalTransfer, "EnableInternalTransfer")
	r.Equal(e.PermitsUniversalTransfer, a.PermitsUniversalTransfer, "PermitsUniversalTransfer")
	r.Equal(e.EnableVanillaOptions, a.EnableVanillaOptions, "EnableVanillaOptions")
	r.Equal(e.EnableReading, a.EnableReading, "EnableReading")
	r.Equal(e.EnableFutures, a.EnableFutures, "EnableFutures")
	r.Equal(e.EnableMargin, a.EnableMargin, "EnableMargin")
	r.Equal(e.EnableSpotAndMarginTrading, a.EnableSpotAndMarginTrading, "EnableSpotAndMarginTrading")
	r.Equal(e.TradingAuthorityExpirationTime, a.TradingAuthorityExpirationTime, "TradingAuthorityExpirationTime")
}
