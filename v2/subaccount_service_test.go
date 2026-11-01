package binance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type subAccountServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountService(t *testing.T) {
	suite.Run(t, new(subAccountServiceTestSuite))
}

func (s *subAccountServiceTestSuite) TestSubaccountDepositAddressService() {
	data := []byte(`
	{
		"address":"TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV",
		"coin":"USDT",
		"tag":"a_tag",
		"url":"https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	coin := "a_coin"
	network := "a_network"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"email":   email,
			"coin":    coin,
			"network": network,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSubaccountDepositAddressService().
		Email(email).
		Coin(coin).
		Network(network).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV", res.Address, "Address")
	r.Equal("USDT", res.Coin, "Coin")
	r.Equal("a_tag", res.Tag, "Tag")
	r.Equal("https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV", res.URL, "URL")
}

func (s *subAccountServiceTestSuite) TestSubAccountListService() {
	data := []byte(`
	{
		"subAccounts":[
			{
				"email":"testsub@gmail.com",
				"isFreeze":false,
				"createTime":1544433328000,
				"isManagedSubAccount": false,
				"isAssetManagementSubAccount": false
			}
		]
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	isFreeze := false
	page := 1
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"email":    email,
			"isFreeze": isFreeze,
			"page":     1,
			"limit":    10,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewSubAccountListService().
		Email(email).
		IsFreeze(false).
		Page(page).
		Limit(limit).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	s.assertSubAccountEqual(SubAccount{
		Email:      email,
		CreateTime: 1544433328000,
	}, res.SubAccounts[0])
}

func (s *subAccountServiceTestSuite) assertSubAccountEqual(e, a SubAccount) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.IsFreeze, a.IsFreeze, "IsFreeze")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.IsManagedSubAccount, a.IsManagedSubAccount, "IsManagedSubAccount")
	r.Equal(e.IsAssetManagementSubAccount, a.IsAssetManagementSubAccount, "IsAssetManagementSubAccount")
}

func (s *subAccountServiceTestSuite) TestSubManagedSubAccountDepositService() {
	data := []byte(` { "tranId": 12345678 } `)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	asset := "USDT"
	amount := "1.0"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"toEmail": email,
			"asset":   asset,
			"amount":  amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewManagedSubAccountDepositService().
		ToEmail(email).
		Asset(asset).
		Amount(amount).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	r.Equal(int64(12345678), res.ID)
}

func (s *subAccountServiceTestSuite) TestSubManagedSubAccountWithdrawalService() {
	data := []byte(` { "tranId": 12345678 } `)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	asset := "USDT"
	amount := "1.0"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromEmail": email,
			"asset":     asset,
			"amount":    amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewManagedSubAccountWithdrawalService().
		FromEmail(email).
		Asset(asset).
		Amount(amount).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	r.Equal(int64(12345678), res.ID)
}

func (s *subAccountServiceTestSuite) TestSubManagedSubAccountAssetsService() {
	data := []byte(`
		[
			{
			"coin": "INJ",                
			"name": "Injective Protocol", 
			"totalBalance": "0",          
			"availableBalance": "0",      
			"inOrder": "0",                
			"btcValue": "0"               
			}
		]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{"email": email})
		s.assertRequestEqual(e, r)
	})

	assets, err := s.client.NewManagedSubAccountAssetsService().Email(email).Do(newContext())

	r := s.r()
	r.NoError(err)

	s.assertAssetsEqual(&ManagedSubAccountAsset{
		Coin:             "INJ",
		Name:             "Injective Protocol",
		TotalBalance:     "0",
		AvailableBalance: "0",
		InOrder:          "0",
		BtcValue:         "0",
	}, assets[0])
}

func (s *subAccountServiceTestSuite) assertAssetsEqual(e, a *ManagedSubAccountAsset) {
	r := s.r()
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.Name, a.Name, "Name")
	r.Equal(e.TotalBalance, a.TotalBalance, "TotalBalance")
	r.Equal(e.AvailableBalance, a.AvailableBalance, "AvailableBalance")
	r.Equal(e.InOrder, a.InOrder, "InOrder")
	r.Equal(e.BtcValue, a.BtcValue, "BtcValue")
}

func (s *subAccountServiceTestSuite) TestSubAccountFuturesService() {
	data := []byte(`
		{
		    "email": "abc@test.com",
		    "asset": "USDT",
		    "assets":[
		        {
		            "asset": "USDT",
		            "initialMargin": "0.00000000",
		            "maintenanceMargin": "0.00000000",
		            "marginBalance": "0.88308000",
		            "maxWithdrawAmount": "0.88308000",
		            "openOrderInitialMargin": "0.00000000",
		            "positionInitialMargin": "0.00000000",
		            "unrealizedProfit": "0.00000000",
		            "walletBalance": "0.88308000"
		         }
		    ],
		    "canDeposit": true,
		    "canTrade": true,
		    "canWithdraw": true,
		    "feeTier": 2,
		    "maxWithdrawAmount": "0.88308000",
		    "totalInitialMargin": "0.00000000",
		    "totalMaintenanceMargin": "0.00000000",
		    "totalMarginBalance": "0.88308000",
		    "totalOpenOrderInitialMargin": "0.00000000",
		    "totalPositionInitialMargin": "0.00000000",
		    "totalUnrealizedProfit": "0.00000000",
		    "totalWalletBalance": "0.88308000",
		    "updateTime": 1576756674610
		 }
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "abc@test.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{"email": email})
		s.assertRequestEqual(e, r)
	})

	account, err := s.client.NewSubAccountFuturesAccountService().Email(email).Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("abc@test.com", account.Email, "Email")
	r.Equal("USDT", account.Asset, "Asset")
	r.Equal(true, account.CanDeposit, "CanDeposit")
	r.Equal(true, account.CanTrade, "CanTrade")
	r.Equal(true, account.CanWithdraw, "CanWithdraw")
	r.Equal(2, account.FeeTier, "FeeTier")
	r.Equal("0.88308000", account.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal("0.00000000", account.TotalInitialMargin, "TotalInitialMargin")
	r.Equal("0.00000000", account.TotalMaintenanceMargin, "TotalMaintenanceMargin")
	r.Equal("0.88308000", account.TotalMarginBalance, "TotalMarginBalance")
	r.Equal("0.00000000", account.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal("0.00000000", account.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal("0.00000000", account.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal("0.88308000", account.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(int64(1576756674610), account.UpdateTime, "UpdateTime")

	s.assertAccountFuturesAssetsEqual(&SubAccountFuturesAccountAsset{
		Asset:                  "USDT",
		InitialMargin:          "0.00000000",
		MaintenanceMargin:      "0.00000000",
		MarginBalance:          "0.88308000",
		MaxWithdrawAmount:      "0.88308000",
		OpenOrderInitialMargin: "0.00000000",
		PositionInitialMargin:  "0.00000000",
		UnrealizedProfit:       "0.00000000",
		WalletBalance:          "0.88308000",
	}, &account.Assets[0])
}

func (s *subAccountServiceTestSuite) assertAccountFuturesAssetsEqual(e, a *SubAccountFuturesAccountAsset) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.InitialMargin, a.InitialMargin, "InitialMargin")
	r.Equal(e.MaintenanceMargin, a.MaintenanceMargin, "MaintenanceMargin")
	r.Equal(e.MarginBalance, a.MarginBalance, "MarginBalance")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal(e.OpenOrderInitialMargin, a.OpenOrderInitialMargin, "OpenOrderInitialMargin")
	r.Equal(e.PositionInitialMargin, a.PositionInitialMargin, "PositionInitialMargin")
	r.Equal(e.UnrealizedProfit, a.UnrealizedProfit, "UnrealizedProfit")
	r.Equal(e.WalletBalance, a.WalletBalance, "WalletBalance")
}

func (s *subAccountServiceTestSuite) TestSubAccountFuturesTransferService() {
	data := []byte(`
		{
		    "tranId": 123456789
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "abc@test.com"
	asset := "USDT"
	amount := 1.0
	tType := 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().
			setParams(params{
				"email":  email,
				"asset":  asset,
				"amount": amount,
				"type":   tType,
			})
		s.assertRequestEqual(e, r)
	})

	response, err := s.client.NewSubAccountFuturesTransferV1Service().Email(email).Asset(asset).Amount(amount).TransferType(tType).Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(int64(123456789), response.TranID, "TranID")

}

func (s *subAccountServiceTestSuite) TestSubAccountTransferHistoryService() {
	data := []byte(`
		[
			{
				"counterParty":"master",
				"email":"master@test.com",
				"type":1,
				"asset":"BTC",
				"qty":"1",
				"fromAccountType":"SPOT",
				"toAccountType":"SPOT",
				"status":"SUCCESS",
				"tranId":11798835829,
				"time":1544433325000
			},
			{
				"counterParty": "subAccount",
				"email": "sub2@test.com",
				"type":  2,                                 
				"asset":"ETH",
				"qty":"2",
				"fromAccountType":"SPOT",
				"toAccountType":"COIN_FUTURE",
				"status":"SUCCESS",
				"tranId":11798829519,
				"time":1544433326000
			}
		]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	transferType := SubAccountTransferTypeTransferIn
	startTime := time.Date(2018, 9, 15, 0, 0, 0, 0, time.UTC).UnixMilli()
	endTime := time.Date(2018, 9, 16, 0, 0, 0, 0, time.UTC).UnixMilli()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"type":      int(transferType),
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	response, err := s.client.NewSubAccountTransferHistoryService().
		TransferType(transferType).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	s.assertSubAccountTransferHistoryEqual(&SubAccountTransferHistory{
		CounterParty:    "master",
		Email:           "master@test.com",
		Type:            1,
		Asset:           "BTC",
		Qty:             "1",
		FromAccountType: "SPOT",
		ToAccountType:   "SPOT",
		Status:          "SUCCESS",
		TranID:          11798835829,
		Time:            1544433325000,
	}, response[0])
}

func (s *subAccountServiceTestSuite) assertSubAccountTransferHistoryEqual(e, a *SubAccountTransferHistory) {
	r := s.r()
	r.Equal(e.CounterParty, a.CounterParty, "CounterParty")
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Qty, a.Qty, "Qty")
	r.Equal(e.FromAccountType, a.FromAccountType, "FromAccountType")
	r.Equal(e.ToAccountType, a.ToAccountType, "ToAccountType")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TranID, a.TranID, "TranId")
	r.Equal(e.Time, a.Time, "Time")
}

type createVirtualSubAccountServiceTestSuite struct {
	baseTestSuite
}

func TestCreateVirtualSubAccountService(t *testing.T) {
	suite.Run(t, new(createVirtualSubAccountServiceTestSuite))
}

func (s *createVirtualSubAccountServiceTestSuite) TestCreateVirtualSubAccount() {
	data := []byte(`{"email": "addsdd_virtual@aasaixwqnoemail.com"}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var subAccountString string = "testSubAccount"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("subAccountString", subAccountString)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewCreateVirtualSubAccountService().SubAccountString(subAccountString).Do(newContext())
	s.r().NoError(err)

	e := &CreateVirtualSubAccountResponse{
		Email: "addsdd_virtual@aasaixwqnoemail.com"}
	s.assertCreateVirtualSubAccountResponseEqual(e, res)
}

func (s *createVirtualSubAccountServiceTestSuite) assertCreateVirtualSubAccountResponseEqual(e, a *CreateVirtualSubAccountResponse) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
}

type subAccountSpotTransferHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountSpotTransferHistoryService(t *testing.T) {
	suite.Run(t, new(subAccountSpotTransferHistoryServiceTestSuite))
}

func (s *subAccountSpotTransferHistoryServiceTestSuite) TestSubAccountSpotTransferHistory() {
	data := []byte(`[{
        "from": "aaa@test.com",
        "to": "bbb@test.com",
        "asset": "BTC",
        "qty": "10",
        "status": "SUCCESS",
        "tranId": 6489943656,
        "time": 1544433328000
    },
    {
        "from": "bbb@test.com",
        "to": "ccc@test.com",
        "asset": "ETH",
        "qty": "2",
        "status": "SUCCESS",
        "tranId": 6489938713,
        "time": 1544433328000
    }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var fromEmail string = "xxyyzz@gmail.com"
	var toEmail string = "aabb@gmail.com"
	var limit int32 = 20
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("fromEmail", fromEmail).setParam("toEmail", toEmail).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountSpotTransferHistoryService().FromEmail(fromEmail).ToEmail(toEmail).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := []*SubAccountSpotTransfer{
		{
			From:   "aaa@test.com",
			To:     "bbb@test.com",
			Asset:  "BTC",
			Qty:    "10",
			Status: "SUCCESS",
			TranId: 6489943656,
			Time:   1544433328000},
		{
			From:   "bbb@test.com",
			To:     "ccc@test.com",
			Asset:  "ETH",
			Qty:    "2",
			Status: "SUCCESS",
			TranId: 6489938713,
			Time:   1544433328000}}

	s.assertSubAccountSpotTransfersEqual(e, res)
}

func (s *subAccountSpotTransferHistoryServiceTestSuite) assertSubAccountSpotTransferEqual(e, a *SubAccountSpotTransfer) {
	r := s.r()
	r.Equal(e.From, a.From, "From")
	r.Equal(e.To, a.To, "To")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Qty, a.Qty, "Qty")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TranId, a.TranId, "TranId")
	r.Equal(e.Time, a.Time, "Time")
}

func (s *subAccountSpotTransferHistoryServiceTestSuite) assertSubAccountSpotTransfersEqual(e, a []*SubAccountSpotTransfer) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertSubAccountSpotTransferEqual(e[i], a[i])
	}
}

type subAccountFuturesTransferHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountFuturesTransferHistoryService(t *testing.T) {
	suite.Run(t, new(subAccountFuturesTransferHistoryServiceTestSuite))
}

func (s *subAccountFuturesTransferHistoryServiceTestSuite) TestSubAccountFuturesTransferHistory() {
	data := []byte(`{
    "success": true,
    "futuresType": 2,
    "transfers": [{
            "from": "aaa@test.com",
            "to": "bbb@test.com",
            "asset": "BTC",
            "qty": "1",
            "tranId": 11897001102,
            "time": 1544433328000
        },
        {
            "from": "bbb@test.com",
            "to": "ccc@test.com",
            "asset": "ETH",
            "qty": "2",
            "tranId": 11631474902,
            "time": 1544433328000
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var futuresType int64 = 1
	var limit int32 = 20
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("futuresType", futuresType).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountFuturesTransferHistoryService().Email(email).FuturesType(futuresType).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountFuturesTransferHistoryResponse{
		Success:     true,
		FuturesType: 2,
		Transfers: []*SubAccountFuturesTransfer{{
			From:   "aaa@test.com",
			To:     "bbb@test.com",
			Asset:  "BTC",
			Qty:    "1",
			TranId: 11897001102,
			Time:   1544433328000},
			{
				From:   "bbb@test.com",
				To:     "ccc@test.com",
				Asset:  "ETH",
				Qty:    "2",
				TranId: 11631474902,
				Time:   1544433328000}}}
	s.assertSubAccountFuturesTransferHistoryResponseEqual(e, res)
}

func (s *subAccountFuturesTransferHistoryServiceTestSuite) assertSubAccountFuturesTransferHistoryResponseEqual(e, a *SubAccountFuturesTransferHistoryResponse) {
	r := s.r()
	r.Equal(e.Success, a.Success, "Success")
	r.Equal(e.FuturesType, a.FuturesType, "FuturesType")
	r.Len(e.Transfers, len(a.Transfers))
	for i := range e.Transfers {
		r.Equal(e.Transfers[i].From, a.Transfers[i].From, "Transfers[i].From")
		r.Equal(e.Transfers[i].To, a.Transfers[i].To, "Transfers[i].To")
		r.Equal(e.Transfers[i].Asset, a.Transfers[i].Asset, "Transfers[i].Asset")
		r.Equal(e.Transfers[i].Qty, a.Transfers[i].Qty, "Transfers[i].Qty")
		r.Equal(e.Transfers[i].Status, a.Transfers[i].Status, "Transfers[i].Status")
		r.Equal(e.Transfers[i].TranId, a.Transfers[i].TranId, "Transfers[i].TranId")
		r.Equal(e.Transfers[i].Time, a.Transfers[i].Time, "Transfers[i].Time")
	}

}

type subAccountFuturesInternalTransferServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountFuturesInternalTransferService(t *testing.T) {
	suite.Run(t, new(subAccountFuturesInternalTransferServiceTestSuite))
}

func (s *subAccountFuturesInternalTransferServiceTestSuite) TestSubAccountFuturesInternalTransfer() {
	data := []byte(`{
    "success": true,
    "txnId": "2934662589"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var fromEmail string = "xxyyzz@gmail.com"
	var toEmail string = "jjkk@gmail.com"
	var futuresType int64 = 1
	var asset string = "USDT"
	var amount string = "1.000085"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("fromEmail", fromEmail).setFormParam("toEmail", toEmail).setFormParam("futuresType", futuresType).setFormParam("asset", asset).setFormParam("amount", amount)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountFuturesInternalTransferService().FromEmail(fromEmail).ToEmail(toEmail).FuturesType(futuresType).Asset(asset).Amount(amount).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountFuturesInternalTransferResponse{
		Success: true,
		TxnId:   "2934662589"}
	s.assertSubAccountFuturesInternalTransferResponseEqual(e, res)
}

func (s *subAccountFuturesInternalTransferServiceTestSuite) assertSubAccountFuturesInternalTransferResponseEqual(e, a *SubAccountFuturesInternalTransferResponse) {
	r := s.r()
	r.Equal(e.Success, a.Success, "Success")
	r.Equal(e.TxnId, a.TxnId, "TxnId")
}

type subAccountDepositRecordServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountDepositRecordService(t *testing.T) {
	suite.Run(t, new(subAccountDepositRecordServiceTestSuite))
}

func (s *subAccountDepositRecordServiceTestSuite) TestSubAccountDepositRecord() {
	data := []byte(`[{
        "id": "769800519366885376",
        "amount": "0.001",
        "coin": "BNB",
        "network": "BNB",
        "status": 0,
        "address": "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23",
        "addressTag": "101764890",
        "txId": "98A3EA560C6B3336D348B6C83F0F95ECE4F1F5919E94BD006E5BF3BF264FACFC",
        "insertTime": 1661493146000,
        "transferType": 0,
        "confirmTimes": "1/1",
        "unlockConfirm": 0,
        "walletType": 0
    },
    {
        "id": "769754833590042625",
        "amount": "0.50000000",
        "coin": "IOTA",
        "network": "IOTA",
        "status": 1,
        "address": "SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
        "addressTag": "",
        "txId": "ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
        "insertTime": 1599620082000,
        "transferType": 0,
        "confirmTimes": "1/1",
        "unlockConfirm": 0,
        "walletType": 0
    }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var coin string = "USDT"
	var status int32 = 6
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("coin", coin).setParam("status", status)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountDepositRecordService().Email(email).Coin(coin).Status(status).Do(newContext())
	s.r().NoError(err)

	e := []*SubAccountDepositRecord{
		{
			Id:            "769800519366885376",
			Amount:        "0.001",
			Coin:          "BNB",
			Network:       "BNB",
			Status:        0,
			Address:       "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23",
			AddressTag:    "101764890",
			TxId:          "98A3EA560C6B3336D348B6C83F0F95ECE4F1F5919E94BD006E5BF3BF264FACFC",
			InsertTime:    1661493146000,
			TransferType:  0,
			ConfirmTimes:  "1/1",
			UnlockConfirm: 0,
			WalletType:    0},
		{
			Id:            "769754833590042625",
			Amount:        "0.50000000",
			Coin:          "IOTA",
			Network:       "IOTA",
			Status:        1,
			Address:       "SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW",
			AddressTag:    "",
			TxId:          "ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999",
			InsertTime:    1599620082000,
			TransferType:  0,
			ConfirmTimes:  "1/1",
			UnlockConfirm: 0,
			WalletType:    0}}

	s.assertSubAccountDepositRecordsEqual(e, res)
}

func (s *subAccountDepositRecordServiceTestSuite) assertSubAccountDepositRecordEqual(e, a *SubAccountDepositRecord) {
	r := s.r()
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.Network, a.Network, "Network")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.TxId, a.TxId, "TxId")
	r.Equal(e.AddressTag, a.AddressTag, "AddressTag")
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.TransferType, a.TransferType, "TransferType")
	r.Equal(e.ConfirmTimes, a.ConfirmTimes, "ConfirmTimes")
	r.Equal(e.UnlockConfirm, a.UnlockConfirm, "UnlockConfirm")
	r.Equal(e.WalletType, a.WalletType, "WalletType")
}

func (s *subAccountDepositRecordServiceTestSuite) assertSubAccountDepositRecordsEqual(e, a []*SubAccountDepositRecord) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertSubAccountDepositRecordEqual(e[i], a[i])
	}
}

type subAccountMarginFuturesStatusServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountMarginFuturesStatusService(t *testing.T) {
	suite.Run(t, new(subAccountMarginFuturesStatusServiceTestSuite))
}

func (s *subAccountMarginFuturesStatusServiceTestSuite) TestSubAccountMarginFuturesStatus() {
	data := []byte(`[{
    "email": "123@test.com",
    "isSubUserEnabled": true,
    "isUserActive": true,
    "insertTime": 1570791523523,
    "isMarginEnabled": true,
    "isFutureEnabled": true,
    "mobile": 1570791523523
}]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountMarginFuturesStatusService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := []*SubAccountMarginFuturesStatus{
		{
			Email:            "123@test.com",
			IsSubUserEnabled: true,
			IsUserActive:     true,
			InsertTime:       1570791523523,
			IsMarginEnabled:  true,
			IsFutureEnabled:  true,
			Mobile:           1570791523523}}

	s.assertSubAccountMarginFuturesStatusesEqual(e, res)
}

func (s *subAccountMarginFuturesStatusServiceTestSuite) assertSubAccountMarginFuturesStatusEqual(e, a *SubAccountMarginFuturesStatus) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.IsSubUserEnabled, a.IsSubUserEnabled, "IsSubUserEnabled")
	r.Equal(e.IsUserActive, a.IsUserActive, "IsUserActive")
	r.Equal(e.InsertTime, a.InsertTime, "InsertTime")
	r.Equal(e.IsMarginEnabled, a.IsMarginEnabled, "IsMarginEnabled")
	r.Equal(e.IsFutureEnabled, a.IsFutureEnabled, "IsFutureEnabled")
	r.Equal(e.Mobile, a.Mobile, "Mobile")
}

func (s *subAccountMarginFuturesStatusServiceTestSuite) assertSubAccountMarginFuturesStatusesEqual(e, a []*SubAccountMarginFuturesStatus) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertSubAccountMarginFuturesStatusEqual(e[i], a[i])
	}
}

type subAccountMarginEnableServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountMarginEnableService(t *testing.T) {
	suite.Run(t, new(subAccountMarginEnableServiceTestSuite))
}

func (s *subAccountMarginEnableServiceTestSuite) TestSubAccountMarginEnable() {
	data := []byte(`{

    "email": "123@test.com",
    "isMarginEnabled": true

}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountMarginEnableService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountMarginEnableResponse{
		Email:           "123@test.com",
		IsMarginEnabled: true}
	s.assertSubAccountMarginEnableResponseEqual(e, res)
}

func (s *subAccountMarginEnableServiceTestSuite) assertSubAccountMarginEnableResponseEqual(e, a *SubAccountMarginEnableResponse) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.IsMarginEnabled, a.IsMarginEnabled, "IsMarginEnabled")
}

type subAccountMarginAccountInfoServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountMarginAccountInfoService(t *testing.T) {
	suite.Run(t, new(subAccountMarginAccountInfoServiceTestSuite))
}

func (s *subAccountMarginAccountInfoServiceTestSuite) TestSubAccountMarginAccountInfo() {
	data := []byte(`{
    "email": "123@test.com",
    "marginLevel": "11.64405625",
    "totalAssetOfBtc": "6.82728457",
    "totalLiabilityOfBtc": "0.58633215",
    "totalNetAssetOfBtc": "6.24095242",
    "marginTradeCoeffVo": {
        "forceLiquidationBar": "1.10000000",
        "marginCallBar": "1.50000000",
        "normalBar": "2.00000000"
    },
    "marginUserAssetVoList": [{
            "asset": "BTC",
            "borrowed": "0.00000000",
            "free": "0.00499500",
            "interest": "0.00000000",
            "locked": "0.00000000",
            "netAsset": "0.00499500"
        },
        {
            "asset": "BNB",
            "borrowed": "201.66666672",
            "free": "2346.50000000",
            "interest": "0.00000000",
            "locked": "0.00000000",
            "netAsset": "2144.83333328"
        },
        {
            "asset": "ETH",
            "borrowed": "0.00000000",
            "free": "0.00000000",
            "interest": "0.00000000",
            "locked": "0.00000000",
            "netAsset": "0.00000000"
        },
        {
            "asset": "USDT",
            "borrowed": "0.00000000",
            "free": "0.00000000",
            "interest": "0.00000000",
            "locked": "0.00000000",
            "netAsset": "0.00000000"
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountMarginAccountInfoService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountMarginAccountInfo{
		Email:               "123@test.com",
		MarginLevel:         "11.64405625",
		TotalAssetOfBtc:     "6.82728457",
		TotalLiabilityOfBtc: "0.58633215",
		TotalNetAssetOfBtc:  "6.24095242",
		MarginTradeCoeffVo: &MarginTradeCoeffVo{
			ForceLiquidationBar: "1.10000000",
			MarginCallBar:       "1.50000000",
			NormalBar:           "2.00000000"},
		MarginUserAssetVoList: []*MarginUserAssetVo{{
			Asset:    "BTC",
			Borrowed: "0.00000000",
			Free:     "0.00499500",
			Interest: "0.00000000",
			Locked:   "0.00000000",
			NetAsset: "0.00499500"},
			{
				Asset:    "BNB",
				Borrowed: "201.66666672",
				Free:     "2346.50000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "2144.83333328"},
			{
				Asset:    "ETH",
				Borrowed: "0.00000000",
				Free:     "0.00000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00000000"},
			{
				Asset:    "USDT",
				Borrowed: "0.00000000",
				Free:     "0.00000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00000000"}}}
	s.assertSubAccountMarginAccountInfoEqual(e, res)
}

func (s *subAccountMarginAccountInfoServiceTestSuite) assertSubAccountMarginAccountInfoEqual(e, a *SubAccountMarginAccountInfo) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.MarginLevel, a.MarginLevel, "MarginLevel")
	r.Equal(e.TotalAssetOfBtc, a.TotalAssetOfBtc, "TotalAssetOfBtc")
	r.Equal(e.TotalLiabilityOfBtc, a.TotalLiabilityOfBtc, "TotalLiabilityOfBtc")
	r.Equal(e.TotalNetAssetOfBtc, a.TotalNetAssetOfBtc, "TotalNetAssetOfBtc")
	r.Equal(e.MarginTradeCoeffVo.ForceLiquidationBar, a.MarginTradeCoeffVo.ForceLiquidationBar, "MarginTradeCoeffVo.ForceLiquidationBar")
	r.Equal(e.MarginTradeCoeffVo.MarginCallBar, a.MarginTradeCoeffVo.MarginCallBar, "MarginTradeCoeffVo.MarginCallBar")
	r.Equal(e.MarginTradeCoeffVo.NormalBar, a.MarginTradeCoeffVo.NormalBar, "MarginTradeCoeffVo.NormalBar")
	r.Len(e.MarginUserAssetVoList, len(a.MarginUserAssetVoList))
	for i := range e.MarginUserAssetVoList {
		r.Equal(e.MarginUserAssetVoList[i].Asset, a.MarginUserAssetVoList[i].Asset, "MarginUserAssetVoList[i].Asset")
		r.Equal(e.MarginUserAssetVoList[i].Borrowed, a.MarginUserAssetVoList[i].Borrowed, "MarginUserAssetVoList[i].Borrowed")
		r.Equal(e.MarginUserAssetVoList[i].Free, a.MarginUserAssetVoList[i].Free, "MarginUserAssetVoList[i].Free")
		r.Equal(e.MarginUserAssetVoList[i].Interest, a.MarginUserAssetVoList[i].Interest, "MarginUserAssetVoList[i].Interest")
		r.Equal(e.MarginUserAssetVoList[i].Locked, a.MarginUserAssetVoList[i].Locked, "MarginUserAssetVoList[i].Locked")
		r.Equal(e.MarginUserAssetVoList[i].NetAsset, a.MarginUserAssetVoList[i].NetAsset, "MarginUserAssetVoList[i].NetAsset")
	}

}

type subAccountMarginAccountSummaryServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountMarginAccountSummaryService(t *testing.T) {
	suite.Run(t, new(subAccountMarginAccountSummaryServiceTestSuite))
}

func (s *subAccountMarginAccountSummaryServiceTestSuite) TestSubAccountMarginAccountSummary() {
	data := []byte(`{
    "totalAssetOfBtc": "4.33333333",
    "totalLiabilityOfBtc": "2.11111112",
    "totalNetAssetOfBtc": "2.22222221",
    "subAccountList": [{
            "email": "123@test.com",
            "totalAssetOfBtc": "2.11111111",
            "totalLiabilityOfBtc": "1.11111111",
            "totalNetAssetOfBtc": "1.00000000"
        },
        {
            "email": "345@test.com",
            "totalAssetOfBtc": "2.22222222",
            "totalLiabilityOfBtc": "1.00000001",
            "totalNetAssetOfBtc": "1.22222221"
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountMarginAccountSummaryService().Do(newContext())
	s.r().NoError(err)

	e := &SubAccountMarginAccountSummary{
		TotalAssetOfBtc:     "4.33333333",
		TotalLiabilityOfBtc: "2.11111112",
		TotalNetAssetOfBtc:  "2.22222221",
		SubAccountList: []*MarginSubAccount{{
			Email:               "123@test.com",
			TotalAssetOfBtc:     "2.11111111",
			TotalLiabilityOfBtc: "1.11111111",
			TotalNetAssetOfBtc:  "1.00000000"},
			{
				Email:               "345@test.com",
				TotalAssetOfBtc:     "2.22222222",
				TotalLiabilityOfBtc: "1.00000001",
				TotalNetAssetOfBtc:  "1.22222221"}}}
	s.assertSubAccountMarginAccountSummaryEqual(e, res)
}

func (s *subAccountMarginAccountSummaryServiceTestSuite) assertSubAccountMarginAccountSummaryEqual(e, a *SubAccountMarginAccountSummary) {
	r := s.r()
	r.Equal(e.TotalAssetOfBtc, a.TotalAssetOfBtc, "TotalAssetOfBtc")
	r.Equal(e.TotalLiabilityOfBtc, a.TotalLiabilityOfBtc, "TotalLiabilityOfBtc")
	r.Equal(e.TotalNetAssetOfBtc, a.TotalNetAssetOfBtc, "TotalNetAssetOfBtc")
	r.Len(e.SubAccountList, len(a.SubAccountList))
	for i := range e.SubAccountList {
		r.Equal(e.SubAccountList[i].Email, a.SubAccountList[i].Email, "SubAccountList[i].Email")
		r.Equal(e.SubAccountList[i].TotalAssetOfBtc, a.SubAccountList[i].TotalAssetOfBtc, "SubAccountList[i].TotalAssetOfBtc")
		r.Equal(e.SubAccountList[i].TotalLiabilityOfBtc, a.SubAccountList[i].TotalLiabilityOfBtc, "SubAccountList[i].TotalLiabilityOfBtc")
		r.Equal(e.SubAccountList[i].TotalNetAssetOfBtc, a.SubAccountList[i].TotalNetAssetOfBtc, "SubAccountList[i].TotalNetAssetOfBtc")
	}

}

type subAccountFuturesEnableServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountFuturesEnableService(t *testing.T) {
	suite.Run(t, new(subAccountFuturesEnableServiceTestSuite))
}

func (s *subAccountFuturesEnableServiceTestSuite) TestSubAccountFuturesEnable() {
	data := []byte(`{

    "email": "123@test.com",
    "isFuturesEnabled": true

}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountFuturesEnableService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountFuturesEnableResponse{
		Email:            "123@test.com",
		IsFuturesEnabled: true}
	s.assertSubAccountFuturesEnableResponseEqual(e, res)
}

func (s *subAccountFuturesEnableServiceTestSuite) assertSubAccountFuturesEnableResponseEqual(e, a *SubAccountFuturesEnableResponse) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.IsFuturesEnabled, a.IsFuturesEnabled, "IsFuturesEnabled")
}

type subAccountFuturesAccountV2ServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountFuturesAccountV2Service(t *testing.T) {
	suite.Run(t, new(subAccountFuturesAccountV2ServiceTestSuite))
}

func (s *subAccountFuturesAccountV2ServiceTestSuite) TestSubAccountFuturesAccountV2() {
	data := []byte(`{
    "futureAccountResp": {
        "email": "abc@test.com",
        "assets": [{
            "asset": "USDT",
            "initialMargin": "0.00000000",
            "maintenanceMargin": "0.00000000",
            "marginBalance": "0.88308000",
            "maxWithdrawAmount": "0.88308000",
            "openOrderInitialMargin": "0.00000000",
            "positionInitialMargin": "0.00000000",
            "unrealizedProfit": "0.00000000",
            "walletBalance": "0.88308000"
        }],
        "canDeposit": true,
        "canTrade": true,
        "canWithdraw": true,
        "feeTier": 2,
        "maxWithdrawAmount": "0.88308000",
        "totalInitialMargin": "0.00000000",
        "totalMaintenanceMargin": "0.00000000",
        "totalMarginBalance": "0.88308000",
        "totalOpenOrderInitialMargin": "0.00000000",
        "totalPositionInitialMargin": "0.00000000",
        "totalUnrealizedProfit": "0.00000000",
        "totalWalletBalance": "0.88308000",
        "updateTime": 1576756674610
    },
    "deliveryAccountResp": {
        "email": "abc@test.com",
        "assets": [{
            "asset": "BTC",
            "initialMargin": "0.00000000",
            "maintenanceMargin": "0.00000000",
            "marginBalance": "0.88308000",
            "maxWithdrawAmount": "0.88308000",
            "openOrderInitialMargin": "0.00000000",
            "positionInitialMargin": "0.00000000",
            "unrealizedProfit": "0.00000000",
            "walletBalance": "0.88308000"
        }],
        "canDeposit": true,
        "canTrade": true,
        "canWithdraw": true,
        "feeTier": 2,
        "updateTime": 1598959682001
    }
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var futuresType int32 = 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("futuresType", futuresType)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountFuturesAccountV2Service().Email(email).FuturesType(futuresType).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountFuturesAccountV2ServiceResponse{
		FutureAccountResp: &SubAccountFuturesAccountV2{
			Email: "abc@test.com",
			Assets: []*FuturesAsset{{
				Asset:                  "USDT",
				InitialMargin:          "0.00000000",
				MaintenanceMargin:      "0.00000000",
				MarginBalance:          "0.88308000",
				MaxWithdrawAmount:      "0.88308000",
				OpenOrderInitialMargin: "0.00000000",
				PositionInitialMargin:  "0.00000000",
				UnrealizedProfit:       "0.00000000",
				WalletBalance:          "0.88308000"}},
			CanDeposit:                  true,
			CanTrade:                    true,
			CanWithdraw:                 true,
			FeeTier:                     2,
			MaxWithdrawAmount:           "0.88308000",
			TotalInitialMargin:          "0.00000000",
			TotalMaintenanceMargin:      "0.00000000",
			TotalMarginBalance:          "0.88308000",
			TotalOpenOrderInitialMargin: "0.00000000",
			TotalPositionInitialMargin:  "0.00000000",
			TotalUnrealizedProfit:       "0.00000000",
			TotalWalletBalance:          "0.88308000",
			UpdateTime:                  1576756674610},
		DeliveryAccountResp: &SubAccountDeliveryAccountV2{
			Email: "abc@test.com",
			Assets: []*FuturesAsset{{
				Asset:                  "BTC",
				InitialMargin:          "0.00000000",
				MaintenanceMargin:      "0.00000000",
				MarginBalance:          "0.88308000",
				MaxWithdrawAmount:      "0.88308000",
				OpenOrderInitialMargin: "0.00000000",
				PositionInitialMargin:  "0.00000000",
				UnrealizedProfit:       "0.00000000",
				WalletBalance:          "0.88308000"}},
			CanDeposit:  true,
			CanTrade:    true,
			CanWithdraw: true,
			FeeTier:     2,
			UpdateTime:  1598959682001}}
	s.assertSubAccountFuturesAccountV2ServiceResponseEqual(e, res)
}

func (s *subAccountFuturesAccountV2ServiceTestSuite) assertSubAccountFuturesAccountV2ServiceResponseEqual(e, a *SubAccountFuturesAccountV2ServiceResponse) {
	r := s.r()
	r.Equal(e.FutureAccountResp.Email, a.FutureAccountResp.Email, "FutureAccountResp.Email")
	r.Equal(e.FutureAccountResp.Asset, a.FutureAccountResp.Asset, "FutureAccountResp.Asset")
	r.Len(e.FutureAccountResp.Assets, len(a.FutureAccountResp.Assets))
	for j := range e.FutureAccountResp.Assets {
		r.Equal(e.FutureAccountResp.Assets[j].Asset, a.FutureAccountResp.Assets[j].Asset, "FutureAccountResp.Assets[j].Asset")
		r.Equal(e.FutureAccountResp.Assets[j].InitialMargin, a.FutureAccountResp.Assets[j].InitialMargin, "FutureAccountResp.Assets[j].InitialMargin")
		r.Equal(e.FutureAccountResp.Assets[j].MaintenanceMargin, a.FutureAccountResp.Assets[j].MaintenanceMargin, "FutureAccountResp.Assets[j].MaintenanceMargin")
		r.Equal(e.FutureAccountResp.Assets[j].MarginBalance, a.FutureAccountResp.Assets[j].MarginBalance, "FutureAccountResp.Assets[j].MarginBalance")
		r.Equal(e.FutureAccountResp.Assets[j].MaxWithdrawAmount, a.FutureAccountResp.Assets[j].MaxWithdrawAmount, "FutureAccountResp.Assets[j].MaxWithdrawAmount")
		r.Equal(e.FutureAccountResp.Assets[j].OpenOrderInitialMargin, a.FutureAccountResp.Assets[j].OpenOrderInitialMargin, "FutureAccountResp.Assets[j].OpenOrderInitialMargin")
		r.Equal(e.FutureAccountResp.Assets[j].PositionInitialMargin, a.FutureAccountResp.Assets[j].PositionInitialMargin, "FutureAccountResp.Assets[j].PositionInitialMargin")
		r.Equal(e.FutureAccountResp.Assets[j].UnrealizedProfit, a.FutureAccountResp.Assets[j].UnrealizedProfit, "FutureAccountResp.Assets[j].UnrealizedProfit")
		r.Equal(e.FutureAccountResp.Assets[j].WalletBalance, a.FutureAccountResp.Assets[j].WalletBalance, "FutureAccountResp.Assets[j].WalletBalance")
	}

	r.Equal(e.FutureAccountResp.CanDeposit, a.FutureAccountResp.CanDeposit, "FutureAccountResp.CanDeposit")
	r.Equal(e.FutureAccountResp.CanTrade, a.FutureAccountResp.CanTrade, "FutureAccountResp.CanTrade")
	r.Equal(e.FutureAccountResp.CanWithdraw, a.FutureAccountResp.CanWithdraw, "FutureAccountResp.CanWithdraw")
	r.Equal(e.FutureAccountResp.FeeTier, a.FutureAccountResp.FeeTier, "FutureAccountResp.FeeTier")
	r.Equal(e.FutureAccountResp.MaxWithdrawAmount, a.FutureAccountResp.MaxWithdrawAmount, "FutureAccountResp.MaxWithdrawAmount")
	r.Equal(e.FutureAccountResp.TotalInitialMargin, a.FutureAccountResp.TotalInitialMargin, "FutureAccountResp.TotalInitialMargin")
	r.Equal(e.FutureAccountResp.TotalMaintenanceMargin, a.FutureAccountResp.TotalMaintenanceMargin, "FutureAccountResp.TotalMaintenanceMargin")
	r.Equal(e.FutureAccountResp.TotalMarginBalance, a.FutureAccountResp.TotalMarginBalance, "FutureAccountResp.TotalMarginBalance")
	r.Equal(e.FutureAccountResp.TotalOpenOrderInitialMargin, a.FutureAccountResp.TotalOpenOrderInitialMargin, "FutureAccountResp.TotalOpenOrderInitialMargin")
	r.Equal(e.FutureAccountResp.TotalPositionInitialMargin, a.FutureAccountResp.TotalPositionInitialMargin, "FutureAccountResp.TotalPositionInitialMargin")
	r.Equal(e.FutureAccountResp.TotalUnrealizedProfit, a.FutureAccountResp.TotalUnrealizedProfit, "FutureAccountResp.TotalUnrealizedProfit")
	r.Equal(e.FutureAccountResp.TotalWalletBalance, a.FutureAccountResp.TotalWalletBalance, "FutureAccountResp.TotalWalletBalance")
	r.Equal(e.FutureAccountResp.UpdateTime, a.FutureAccountResp.UpdateTime, "FutureAccountResp.UpdateTime")
	r.Equal(e.DeliveryAccountResp.Email, a.DeliveryAccountResp.Email, "DeliveryAccountResp.Email")
	r.Len(e.DeliveryAccountResp.Assets, len(a.DeliveryAccountResp.Assets))
	for j := range e.DeliveryAccountResp.Assets {
		r.Equal(e.DeliveryAccountResp.Assets[j].Asset, a.DeliveryAccountResp.Assets[j].Asset, "DeliveryAccountResp.Assets[j].Asset")
		r.Equal(e.DeliveryAccountResp.Assets[j].InitialMargin, a.DeliveryAccountResp.Assets[j].InitialMargin, "DeliveryAccountResp.Assets[j].InitialMargin")
		r.Equal(e.DeliveryAccountResp.Assets[j].MaintenanceMargin, a.DeliveryAccountResp.Assets[j].MaintenanceMargin, "DeliveryAccountResp.Assets[j].MaintenanceMargin")
		r.Equal(e.DeliveryAccountResp.Assets[j].MarginBalance, a.DeliveryAccountResp.Assets[j].MarginBalance, "DeliveryAccountResp.Assets[j].MarginBalance")
		r.Equal(e.DeliveryAccountResp.Assets[j].MaxWithdrawAmount, a.DeliveryAccountResp.Assets[j].MaxWithdrawAmount, "DeliveryAccountResp.Assets[j].MaxWithdrawAmount")
		r.Equal(e.DeliveryAccountResp.Assets[j].OpenOrderInitialMargin, a.DeliveryAccountResp.Assets[j].OpenOrderInitialMargin, "DeliveryAccountResp.Assets[j].OpenOrderInitialMargin")
		r.Equal(e.DeliveryAccountResp.Assets[j].PositionInitialMargin, a.DeliveryAccountResp.Assets[j].PositionInitialMargin, "DeliveryAccountResp.Assets[j].PositionInitialMargin")
		r.Equal(e.DeliveryAccountResp.Assets[j].UnrealizedProfit, a.DeliveryAccountResp.Assets[j].UnrealizedProfit, "DeliveryAccountResp.Assets[j].UnrealizedProfit")
		r.Equal(e.DeliveryAccountResp.Assets[j].WalletBalance, a.DeliveryAccountResp.Assets[j].WalletBalance, "DeliveryAccountResp.Assets[j].WalletBalance")
	}

	r.Equal(e.DeliveryAccountResp.CanDeposit, a.DeliveryAccountResp.CanDeposit, "DeliveryAccountResp.CanDeposit")
	r.Equal(e.DeliveryAccountResp.CanTrade, a.DeliveryAccountResp.CanTrade, "DeliveryAccountResp.CanTrade")
	r.Equal(e.DeliveryAccountResp.CanWithdraw, a.DeliveryAccountResp.CanWithdraw, "DeliveryAccountResp.CanWithdraw")
	r.Equal(e.DeliveryAccountResp.FeeTier, a.DeliveryAccountResp.FeeTier, "DeliveryAccountResp.FeeTier")
	r.Equal(e.DeliveryAccountResp.UpdateTime, a.DeliveryAccountResp.UpdateTime, "DeliveryAccountResp.UpdateTime")
}

type subAccountFuturesAccountSummaryServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountFuturesAccountSummaryService(t *testing.T) {
	suite.Run(t, new(subAccountFuturesAccountSummaryServiceTestSuite))
}

func (s *subAccountFuturesAccountSummaryServiceTestSuite) TestSubAccountFuturesAccountSummary() {
	data := []byte(`{
    "futureAccountSummaryResp": {
        "totalInitialMargin": "9.83137400",
        "totalMaintenanceMargin": "0.41568700",
        "totalMarginBalance": "23.03235621",
        "totalOpenOrderInitialMargin": "9.00000000",
        "totalPositionInitialMargin": "0.83137400",
        "totalUnrealizedProfit": "0.03219710",
        "totalWalletBalance": "22.15879444",
        "asset": "USD",
        "subAccountList": [{
                "email": "123@test.com",
                "totalInitialMargin": "9.00000000",
                "totalMaintenanceMargin": "0.00000000",
                "totalMarginBalance": "22.12659734",
                "totalOpenOrderInitialMargin": "9.00000000",
                "totalPositionInitialMargin": "0.00000000",
                "totalUnrealizedProfit": "0.00000000",
                "totalWalletBalance": "22.12659734",
                "asset": "USD"
            },
            {
                "email": "345@test.com",
                "totalInitialMargin": "0.83137400",
                "totalMaintenanceMargin": "0.41568700",
                "totalMarginBalance": "0.90575887",
                "totalOpenOrderInitialMargin": "0.00000000",
                "totalPositionInitialMargin": "0.83137400",
                "totalUnrealizedProfit": "0.03219710",
                "totalWalletBalance": "0.87356177",
                "asset": "USD"
            }
        ]
    },

    "deliveryAccountSummaryResp": {
        "totalMarginBalanceOfBTC": "25.03221121",
        "totalUnrealizedProfitOfBTC": "0.12233410",
        "totalWalletBalanceOfBTC": "22.15879444",
        "asset": "BTC",
        "subAccountList": [{
                "email": "123@test.com",
                "totalMarginBalance": "22.12659734",
                "totalUnrealizedProfit": "0.00000000",
                "totalWalletBalance": "22.12659734",
                "asset": "BTC"
            },
            {
                "email": "345@test.com",
                "totalMarginBalance": "0.90575887",
                "totalUnrealizedProfit": "0.03219710",
                "totalWalletBalance": "0.87356177",
                "asset": "BTC"
            }
        ]
    }
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var futuresType int32 = 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("futuresType", futuresType)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountFuturesAccountSummaryService().FuturesType(futuresType).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountFuturesAccountSummaryServiceResponse{
		FutureAccountSummaryResp: &SubAccountFuturesAccountSummary{
			TotalInitialMargin:          "9.83137400",
			TotalMaintenanceMargin:      "0.41568700",
			TotalMarginBalance:          "23.03235621",
			TotalOpenOrderInitialMargin: "9.00000000",
			TotalPositionInitialMargin:  "0.83137400",
			TotalUnrealizedProfit:       "0.03219710",
			TotalWalletBalance:          "22.15879444",
			Asset:                       "USD",
			SubAccountList: []*FuturesSubAccount{{
				Email:                       "123@test.com",
				TotalInitialMargin:          "9.00000000",
				TotalMaintenanceMargin:      "0.00000000",
				TotalMarginBalance:          "22.12659734",
				TotalOpenOrderInitialMargin: "9.00000000",
				TotalPositionInitialMargin:  "0.00000000",
				TotalUnrealizedProfit:       "0.00000000",
				TotalWalletBalance:          "22.12659734",
				Asset:                       "USD"},
				{
					Email:                       "345@test.com",
					TotalInitialMargin:          "0.83137400",
					TotalMaintenanceMargin:      "0.41568700",
					TotalMarginBalance:          "0.90575887",
					TotalOpenOrderInitialMargin: "0.00000000",
					TotalPositionInitialMargin:  "0.83137400",
					TotalUnrealizedProfit:       "0.03219710",
					TotalWalletBalance:          "0.87356177",
					Asset:                       "USD"}}},
		DeliveryAccountSummaryResp: &SubAccountDeliveryAccountSummary{
			TotalMarginBalanceOfBTC:    "25.03221121",
			TotalUnrealizedProfitOfBTC: "0.12233410",
			TotalWalletBalanceOfBTC:    "22.15879444",
			Asset:                      "BTC",
			SubAccountList: []*DeliverySubAccount{{
				Email:                 "123@test.com",
				TotalMarginBalance:    "22.12659734",
				TotalUnrealizedProfit: "0.00000000",
				TotalWalletBalance:    "22.12659734",
				Asset:                 "BTC"},
				{
					Email:                 "345@test.com",
					TotalMarginBalance:    "0.90575887",
					TotalUnrealizedProfit: "0.03219710",
					TotalWalletBalance:    "0.87356177",
					Asset:                 "BTC"}}}}
	s.assertSubAccountFuturesAccountSummaryServiceResponseEqual(e, res)
}

func (s *subAccountFuturesAccountSummaryServiceTestSuite) assertSubAccountFuturesAccountSummaryServiceResponseEqual(e, a *SubAccountFuturesAccountSummaryServiceResponse) {
	r := s.r()
	r.Equal(e.FutureAccountSummaryResp.TotalInitialMargin, a.FutureAccountSummaryResp.TotalInitialMargin, "FutureAccountSummaryResp.TotalInitialMargin")
	r.Equal(e.FutureAccountSummaryResp.TotalMaintenanceMargin, a.FutureAccountSummaryResp.TotalMaintenanceMargin, "FutureAccountSummaryResp.TotalMaintenanceMargin")
	r.Equal(e.FutureAccountSummaryResp.TotalMarginBalance, a.FutureAccountSummaryResp.TotalMarginBalance, "FutureAccountSummaryResp.TotalMarginBalance")
	r.Equal(e.FutureAccountSummaryResp.TotalOpenOrderInitialMargin, a.FutureAccountSummaryResp.TotalOpenOrderInitialMargin, "FutureAccountSummaryResp.TotalOpenOrderInitialMargin")
	r.Equal(e.FutureAccountSummaryResp.TotalPositionInitialMargin, a.FutureAccountSummaryResp.TotalPositionInitialMargin, "FutureAccountSummaryResp.TotalPositionInitialMargin")
	r.Equal(e.FutureAccountSummaryResp.TotalUnrealizedProfit, a.FutureAccountSummaryResp.TotalUnrealizedProfit, "FutureAccountSummaryResp.TotalUnrealizedProfit")
	r.Equal(e.FutureAccountSummaryResp.TotalWalletBalance, a.FutureAccountSummaryResp.TotalWalletBalance, "FutureAccountSummaryResp.TotalWalletBalance")
	r.Equal(e.FutureAccountSummaryResp.Asset, a.FutureAccountSummaryResp.Asset, "FutureAccountSummaryResp.Asset")
	r.Len(e.FutureAccountSummaryResp.SubAccountList, len(a.FutureAccountSummaryResp.SubAccountList))
	for j := range e.FutureAccountSummaryResp.SubAccountList {
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].Email, a.FutureAccountSummaryResp.SubAccountList[j].Email, "FutureAccountSummaryResp.SubAccountList[j].Email")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalInitialMargin, a.FutureAccountSummaryResp.SubAccountList[j].TotalInitialMargin, "FutureAccountSummaryResp.SubAccountList[j].TotalInitialMargin")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalMaintenanceMargin, a.FutureAccountSummaryResp.SubAccountList[j].TotalMaintenanceMargin, "FutureAccountSummaryResp.SubAccountList[j].TotalMaintenanceMargin")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalMarginBalance, a.FutureAccountSummaryResp.SubAccountList[j].TotalMarginBalance, "FutureAccountSummaryResp.SubAccountList[j].TotalMarginBalance")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalOpenOrderInitialMargin, a.FutureAccountSummaryResp.SubAccountList[j].TotalOpenOrderInitialMargin, "FutureAccountSummaryResp.SubAccountList[j].TotalOpenOrderInitialMargin")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalPositionInitialMargin, a.FutureAccountSummaryResp.SubAccountList[j].TotalPositionInitialMargin, "FutureAccountSummaryResp.SubAccountList[j].TotalPositionInitialMargin")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalUnrealizedProfit, a.FutureAccountSummaryResp.SubAccountList[j].TotalUnrealizedProfit, "FutureAccountSummaryResp.SubAccountList[j].TotalUnrealizedProfit")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].TotalWalletBalance, a.FutureAccountSummaryResp.SubAccountList[j].TotalWalletBalance, "FutureAccountSummaryResp.SubAccountList[j].TotalWalletBalance")
		r.Equal(e.FutureAccountSummaryResp.SubAccountList[j].Asset, a.FutureAccountSummaryResp.SubAccountList[j].Asset, "FutureAccountSummaryResp.SubAccountList[j].Asset")
	}

	r.Equal(e.DeliveryAccountSummaryResp.TotalMarginBalanceOfBTC, a.DeliveryAccountSummaryResp.TotalMarginBalanceOfBTC, "DeliveryAccountSummaryResp.TotalMarginBalanceOfBTC")
	r.Equal(e.DeliveryAccountSummaryResp.TotalUnrealizedProfitOfBTC, a.DeliveryAccountSummaryResp.TotalUnrealizedProfitOfBTC, "DeliveryAccountSummaryResp.TotalUnrealizedProfitOfBTC")
	r.Equal(e.DeliveryAccountSummaryResp.TotalWalletBalanceOfBTC, a.DeliveryAccountSummaryResp.TotalWalletBalanceOfBTC, "DeliveryAccountSummaryResp.TotalWalletBalanceOfBTC")
	r.Equal(e.DeliveryAccountSummaryResp.Asset, a.DeliveryAccountSummaryResp.Asset, "DeliveryAccountSummaryResp.Asset")
	r.Len(e.DeliveryAccountSummaryResp.SubAccountList, len(a.DeliveryAccountSummaryResp.SubAccountList))
	for j := range e.DeliveryAccountSummaryResp.SubAccountList {
		r.Equal(e.DeliveryAccountSummaryResp.SubAccountList[j].Email, a.DeliveryAccountSummaryResp.SubAccountList[j].Email, "DeliveryAccountSummaryResp.SubAccountList[j].Email")
		r.Equal(e.DeliveryAccountSummaryResp.SubAccountList[j].TotalMarginBalance, a.DeliveryAccountSummaryResp.SubAccountList[j].TotalMarginBalance, "DeliveryAccountSummaryResp.SubAccountList[j].TotalMarginBalance")
		r.Equal(e.DeliveryAccountSummaryResp.SubAccountList[j].TotalUnrealizedProfit, a.DeliveryAccountSummaryResp.SubAccountList[j].TotalUnrealizedProfit, "DeliveryAccountSummaryResp.SubAccountList[j].TotalUnrealizedProfit")
		r.Equal(e.DeliveryAccountSummaryResp.SubAccountList[j].TotalWalletBalance, a.DeliveryAccountSummaryResp.SubAccountList[j].TotalWalletBalance, "DeliveryAccountSummaryResp.SubAccountList[j].TotalWalletBalance")
		r.Equal(e.DeliveryAccountSummaryResp.SubAccountList[j].Asset, a.DeliveryAccountSummaryResp.SubAccountList[j].Asset, "DeliveryAccountSummaryResp.SubAccountList[j].Asset")
	}

}

type subAccountFuturesPositionsServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountFuturesPositionsService(t *testing.T) {
	suite.Run(t, new(subAccountFuturesPositionsServiceTestSuite))
}

func (s *subAccountFuturesPositionsServiceTestSuite) TestSubAccountFuturesPositions() {
	data := []byte(`{
    "futurePositionRiskVos": [{
        "entryPrice": "9975.12000",
        "leverage": "50",
        "maxNotional": "1000000",
        "liquidationPrice": "7963.54",
        "markPrice": "9973.50770517",
        "positionAmount": "0.010",
        "symbol": "BTCUSDT",
        "unrealizedProfit": "-0.01612295"
    }],
    "deliveryPositionRiskVos": [{
        "entryPrice": "9975.12000",
        "markPrice": "9973.50770517",
        "leverage": "20",
        "isolated": "false",
        "isolatedWallet": "9973.50770517",
        "isolatedMargin": "0.00000000",
        "isAutoAddMargin": "false",
        "positionSide": "BOTH",
        "positionAmount": "1.230",
        "symbol": "BTCUSD_201225",
        "unrealizedProfit": "-0.01612295"
    }]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var futuresType int32 = 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("futuresType", futuresType)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountFuturesPositionsService().Email(email).FuturesType(futuresType).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountFuturesPositionsServiceResponse{
		FuturePositionRiskVos: []*SubAccountFuturesPosition{{
			EntryPrice:       "9975.12000",
			Leverage:         "50",
			MaxNotional:      "1000000",
			LiquidationPrice: "7963.54",
			MarkPrice:        "9973.50770517",
			PositionAmount:   "0.010",
			Symbol:           "BTCUSDT",
			UnrealizedProfit: "-0.01612295"}},
		DeliveryPositionRiskVos: []*SubAccountDeliveryPosition{{
			EntryPrice:       "9975.12000",
			MarkPrice:        "9973.50770517",
			Leverage:         "20",
			Isolated:         "false",
			IsolatedWallet:   "9973.50770517",
			IsolatedMargin:   "0.00000000",
			IsAutoAddMargin:  "false",
			PositionSide:     "BOTH",
			PositionAmount:   "1.230",
			Symbol:           "BTCUSD_201225",
			UnrealizedProfit: "-0.01612295"}}}
	s.assertSubAccountFuturesPositionsServiceResponseEqual(e, res)
}

func (s *subAccountFuturesPositionsServiceTestSuite) assertSubAccountFuturesPositionsServiceResponseEqual(e, a *SubAccountFuturesPositionsServiceResponse) {
	r := s.r()
	r.Len(e.FuturePositionRiskVos, len(a.FuturePositionRiskVos))
	for i := range e.FuturePositionRiskVos {
		r.Equal(e.FuturePositionRiskVos[i].EntryPrice, a.FuturePositionRiskVos[i].EntryPrice, "FuturePositionRiskVos[i].EntryPrice")
		r.Equal(e.FuturePositionRiskVos[i].Leverage, a.FuturePositionRiskVos[i].Leverage, "FuturePositionRiskVos[i].Leverage")
		r.Equal(e.FuturePositionRiskVos[i].MaxNotional, a.FuturePositionRiskVos[i].MaxNotional, "FuturePositionRiskVos[i].MaxNotional")
		r.Equal(e.FuturePositionRiskVos[i].LiquidationPrice, a.FuturePositionRiskVos[i].LiquidationPrice, "FuturePositionRiskVos[i].LiquidationPrice")
		r.Equal(e.FuturePositionRiskVos[i].MarkPrice, a.FuturePositionRiskVos[i].MarkPrice, "FuturePositionRiskVos[i].MarkPrice")
		r.Equal(e.FuturePositionRiskVos[i].PositionAmount, a.FuturePositionRiskVos[i].PositionAmount, "FuturePositionRiskVos[i].PositionAmount")
		r.Equal(e.FuturePositionRiskVos[i].Symbol, a.FuturePositionRiskVos[i].Symbol, "FuturePositionRiskVos[i].Symbol")
		r.Equal(e.FuturePositionRiskVos[i].UnrealizedProfit, a.FuturePositionRiskVos[i].UnrealizedProfit, "FuturePositionRiskVos[i].UnrealizedProfit")
	}

	r.Len(e.DeliveryPositionRiskVos, len(a.DeliveryPositionRiskVos))
	for i := range e.DeliveryPositionRiskVos {
		r.Equal(e.DeliveryPositionRiskVos[i].EntryPrice, a.DeliveryPositionRiskVos[i].EntryPrice, "DeliveryPositionRiskVos[i].EntryPrice")
		r.Equal(e.DeliveryPositionRiskVos[i].MarkPrice, a.DeliveryPositionRiskVos[i].MarkPrice, "DeliveryPositionRiskVos[i].MarkPrice")
		r.Equal(e.DeliveryPositionRiskVos[i].Leverage, a.DeliveryPositionRiskVos[i].Leverage, "DeliveryPositionRiskVos[i].Leverage")
		r.Equal(e.DeliveryPositionRiskVos[i].Isolated, a.DeliveryPositionRiskVos[i].Isolated, "DeliveryPositionRiskVos[i].Isolated")
		r.Equal(e.DeliveryPositionRiskVos[i].IsolatedWallet, a.DeliveryPositionRiskVos[i].IsolatedWallet, "DeliveryPositionRiskVos[i].IsolatedWallet")
		r.Equal(e.DeliveryPositionRiskVos[i].IsolatedMargin, a.DeliveryPositionRiskVos[i].IsolatedMargin, "DeliveryPositionRiskVos[i].IsolatedMargin")
		r.Equal(e.DeliveryPositionRiskVos[i].IsAutoAddMargin, a.DeliveryPositionRiskVos[i].IsAutoAddMargin, "DeliveryPositionRiskVos[i].IsAutoAddMargin")
		r.Equal(e.DeliveryPositionRiskVos[i].PositionSide, a.DeliveryPositionRiskVos[i].PositionSide, "DeliveryPositionRiskVos[i].PositionSide")
		r.Equal(e.DeliveryPositionRiskVos[i].PositionAmount, a.DeliveryPositionRiskVos[i].PositionAmount, "DeliveryPositionRiskVos[i].PositionAmount")
		r.Equal(e.DeliveryPositionRiskVos[i].Symbol, a.DeliveryPositionRiskVos[i].Symbol, "DeliveryPositionRiskVos[i].Symbol")
		r.Equal(e.DeliveryPositionRiskVos[i].UnrealizedProfit, a.DeliveryPositionRiskVos[i].UnrealizedProfit, "DeliveryPositionRiskVos[i].UnrealizedProfit")
	}

}

type subAccountMarginTransferServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountMarginTransferService(t *testing.T) {
	suite.Run(t, new(subAccountMarginTransferServiceTestSuite))
}

func (s *subAccountMarginTransferServiceTestSuite) TestSubAccountMarginTransfer() {
	data := []byte(`{
    "txnId":"2966662589"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var asset string = "USDT"
	var amount string = "0.0125586"
	var transferType int32 = 1
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email).setFormParam("asset", asset).setFormParam("amount", amount).setFormParam("type", transferType)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountMarginTransferService().Email(email).Asset(asset).Amount(amount).TransferType(transferType).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountMarginTransferResponse{
		TxnId: "2966662589"}
	s.assertSubAccountMarginTransferResponseEqual(e, res)
}

func (s *subAccountMarginTransferServiceTestSuite) assertSubAccountMarginTransferResponseEqual(e, a *SubAccountMarginTransferResponse) {
	r := s.r()
	r.Equal(e.TxnId, a.TxnId, "TxnId")
}

type subAccountTransferSubToMasterServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountTransferSubToMasterService(t *testing.T) {
	suite.Run(t, new(subAccountTransferSubToMasterServiceTestSuite))
}

func (s *subAccountTransferSubToMasterServiceTestSuite) TestSubAccountTransferSubToMaster() {
	data := []byte(`{
    "txnId":"2966662589"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var asset string = "USDT"
	var amount string = "0.0125586"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("asset", asset).setFormParam("amount", amount)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountTransferSubToMasterService().Asset(asset).Amount(amount).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountTransferSubToMasterResponse{
		TxnId: "2966662589"}
	s.assertSubAccountTransferSubToMasterResponseEqual(e, res)
}

func (s *subAccountTransferSubToMasterServiceTestSuite) assertSubAccountTransferSubToMasterResponseEqual(e, a *SubAccountTransferSubToMasterResponse) {
	r := s.r()
	r.Equal(e.TxnId, a.TxnId, "TxnId")
}

type subAccountUniversalTransferServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountUniversalTransferService(t *testing.T) {
	suite.Run(t, new(subAccountUniversalTransferServiceTestSuite))
}

func (s *subAccountUniversalTransferServiceTestSuite) TestSubAccountUniversalTransfer() {
	data := []byte(`{
    "tranId":11945860693,
    "clientTranId":"test"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var fromAccountType string = "SPOT"
	var toAccountType string = "USDT_FUTURE"
	var asset string = "USDT"
	var amount string = "0.0125586"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("fromAccountType", fromAccountType).setFormParam("toAccountType", toAccountType).setFormParam("asset", asset).setFormParam("amount", amount)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountUniversalTransferService().FromAccountType(fromAccountType).ToAccountType(toAccountType).Asset(asset).Amount(amount).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountUniversalTransferResponse{
		TranId:       11945860693,
		ClientTranId: "test"}
	s.assertSubAccountUniversalTransferResponseEqual(e, res)
}

func (s *subAccountUniversalTransferServiceTestSuite) assertSubAccountUniversalTransferResponseEqual(e, a *SubAccountUniversalTransferResponse) {
	r := s.r()
	r.Equal(e.TranId, a.TranId, "TranId")
	r.Equal(e.ClientTranId, a.ClientTranId, "ClientTranId")
}

type subAccUniversalTransferHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccUniversalTransferHistoryService(t *testing.T) {
	suite.Run(t, new(subAccUniversalTransferHistoryServiceTestSuite))
}

func (s *subAccUniversalTransferHistoryServiceTestSuite) TestSubAccUniversalTransferHistory() {
	data := []byte(`{
    "result": [{
        "tranId": 92275823339,
        "fromEmail": "abctest@gmail.com",
        "toEmail": "deftest@gmail.com",
        "asset": "BNB",
        "amount": "0.01",
        "createTimeStamp": 1640317374000,
        "fromAccountType": "USDT_FUTURE",
        "toAccountType": "SPOT",
        "status": "SUCCESS",
        "clientTranId": "test"
    }],
    "totalCount": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var fromEmail string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("fromEmail", fromEmail)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccUniversalTransferHistoryService().FromEmail(fromEmail).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountUniversalTransferHistoryServiceResponse{
		Result: []*SubAccountUniversalTransferRecord{{
			TranId:          92275823339,
			FromEmail:       "abctest@gmail.com",
			ToEmail:         "deftest@gmail.com",
			Asset:           "BNB",
			Amount:          "0.01",
			CreateTimeStamp: 1640317374000,
			FromAccountType: "USDT_FUTURE",
			ToAccountType:   "SPOT",
			Status:          "SUCCESS",
			ClientTranId:    "test"}},
		TotalCount: 1}
	s.assertSubAccountUniversalTransferHistoryServiceResponseEqual(e, res)
}

func (s *subAccUniversalTransferHistoryServiceTestSuite) assertSubAccountUniversalTransferHistoryServiceResponseEqual(e, a *SubAccountUniversalTransferHistoryServiceResponse) {
	r := s.r()
	r.Len(e.Result, len(a.Result))
	for i := range e.Result {
		r.Equal(e.Result[i].TranId, a.Result[i].TranId, "Result[i].TranId")
		r.Equal(e.Result[i].FromEmail, a.Result[i].FromEmail, "Result[i].FromEmail")
		r.Equal(e.Result[i].ToEmail, a.Result[i].ToEmail, "Result[i].ToEmail")
		r.Equal(e.Result[i].Asset, a.Result[i].Asset, "Result[i].Asset")
		r.Equal(e.Result[i].Amount, a.Result[i].Amount, "Result[i].Amount")
		r.Equal(e.Result[i].CreateTimeStamp, a.Result[i].CreateTimeStamp, "Result[i].CreateTimeStamp")
		r.Equal(e.Result[i].FromAccountType, a.Result[i].FromAccountType, "Result[i].FromAccountType")
		r.Equal(e.Result[i].ToAccountType, a.Result[i].ToAccountType, "Result[i].ToAccountType")
		r.Equal(e.Result[i].Status, a.Result[i].Status, "Result[i].Status")
		r.Equal(e.Result[i].ClientTranId, a.Result[i].ClientTranId, "Result[i].ClientTranId")
	}

	r.Equal(e.TotalCount, a.TotalCount, "TotalCount")
}

type subAccountBlvtEnableServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountBlvtEnableService(t *testing.T) {
	suite.Run(t, new(subAccountBlvtEnableServiceTestSuite))
}

func (s *subAccountBlvtEnableServiceTestSuite) TestSubAccountBlvtEnable() {
	data := []byte(`{
    "email":"123@test.com",
    "enableBlvt":true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var enableBlvt bool = true
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email).setFormParam("enableBlvt", enableBlvt)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountBlvtEnableService().Email(email).EnableBlvt(enableBlvt).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountBlvtEnableServiceResponse{
		Email:      "123@test.com",
		EnableBlvt: true}
	s.assertSubAccountBlvtEnableServiceResponseEqual(e, res)
}

func (s *subAccountBlvtEnableServiceTestSuite) assertSubAccountBlvtEnableServiceResponseEqual(e, a *SubAccountBlvtEnableServiceResponse) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.EnableBlvt, a.EnableBlvt, "EnableBlvt")
}

type subAccountApiIpRestrictionServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountApiIpRestrictionService(t *testing.T) {
	suite.Run(t, new(subAccountApiIpRestrictionServiceTestSuite))
}

func (s *subAccountApiIpRestrictionServiceTestSuite) TestSubAccountApiIpRestriction() {
	data := []byte(`{
    "ipRestrict": "true",
    "ipList": [
        "69.210.67.14",
        "8.34.21.10"
    ],
    "updateTime": 1636371437000,
    "apiKey": "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var subAccountApiKey string = "aabb"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("subAccountApiKey", subAccountApiKey)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountApiIpRestrictionService().Email(email).SubAccountApiKey(subAccountApiKey).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountApiIpRestrictServiceResponse{
		IpRestrict: "true",
		IpList: []string{"69.210.67.14",
			"8.34.21.10"},
		UpdateTime: 1636371437000,
		ApiKey:     "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}
	s.assertSubAccountApiIpRestrictServiceResponseEqual(e, res)
}

func (s *subAccountApiIpRestrictionServiceTestSuite) assertSubAccountApiIpRestrictServiceResponseEqual(e, a *SubAccountApiIpRestrictServiceResponse) {
	r := s.r()
	r.Equal(e.IpRestrict, a.IpRestrict, "IpRestrict")
	for i := range e.IpList {
		r.Equal(e.IpList[i], a.IpList[i], "IpList[i]")
	}

	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.ApiKey, a.ApiKey, "ApiKey")
}

type subAccountApiDeleteIpRestrictionServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountApiDeleteIpRestrictionService(t *testing.T) {
	suite.Run(t, new(subAccountApiDeleteIpRestrictionServiceTestSuite))
}

func (s *subAccountApiDeleteIpRestrictionServiceTestSuite) TestSubAccountApiDeleteIpRestriction() {
	data := []byte(`{
  "ipRestrict": "true",
  "ipList": [
    "69.210.67.14",
    "8.34.21.10"
  ],
  "updateTime": 1636371437000,
  "apiKey": "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var subAccountApiKey string = "aabb"
	var ipAddress string = "1.2.3.4"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email).setFormParam("subAccountApiKey", subAccountApiKey).setFormParam("ipAddress", ipAddress)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountApiDeleteIpRestrictionService().Email(email).SubAccountApiKey(subAccountApiKey).IpAddress(ipAddress).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountApiDeleteIpRestrictServiceResponse{
		IpRestrict: "true",
		IpList: []string{"69.210.67.14",
			"8.34.21.10"},
		UpdateTime: 1636371437000,
		ApiKey:     "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}
	s.assertSubAccountApiDeleteIpRestrictServiceResponseEqual(e, res)
}

func (s *subAccountApiDeleteIpRestrictionServiceTestSuite) assertSubAccountApiDeleteIpRestrictServiceResponseEqual(e, a *SubAccountApiDeleteIpRestrictServiceResponse) {
	r := s.r()
	r.Equal(e.IpRestrict, a.IpRestrict, "IpRestrict")
	for i := range e.IpList {
		r.Equal(e.IpList[i], a.IpList[i], "IpList[i]")
	}

	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.ApiKey, a.ApiKey, "ApiKey")
}

type subAccountApiAddIpRestrictionServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountApiAddIpRestrictionService(t *testing.T) {
	suite.Run(t, new(subAccountApiAddIpRestrictionServiceTestSuite))
}

func (s *subAccountApiAddIpRestrictionServiceTestSuite) TestSubAccountApiAddIpRestriction() {
	data := []byte(`{
  "ipRestrict": "true",
  "ipList": [
    "69.210.67.14",
    "8.34.21.10"
  ],
  "updateTime": 1636371437000,
  "apiKey": "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var subAccountApiKey string = "aabb"
	var status string = "1"
	var ipAddress string = "1.2.3.4"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email).setFormParam("subAccountApiKey", subAccountApiKey).setFormParam("status", status).setFormParam("ipAddress", ipAddress)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountApiAddIpRestrictionService().Email(email).SubAccountApiKey(subAccountApiKey).Status(status).IpAddress(ipAddress).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountApiAddIpRestrictServiceResponse{
		IpRestrict: "true",
		IpList: []string{"69.210.67.14",
			"8.34.21.10"},
		UpdateTime: 1636371437000,
		ApiKey:     "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}
	s.assertSubAccountApiAddIpRestrictServiceResponseEqual(e, res)
}

func (s *subAccountApiAddIpRestrictionServiceTestSuite) assertSubAccountApiAddIpRestrictServiceResponseEqual(e, a *SubAccountApiAddIpRestrictServiceResponse) {
	r := s.r()
	r.Equal(e.IpRestrict, a.IpRestrict, "IpRestrict")
	for i := range e.IpList {
		r.Equal(e.IpList[i], a.IpList[i], "IpList[i]")
	}

	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.ApiKey, a.ApiKey, "ApiKey")
}

type managedSubAccountWithdrawServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountWithdrawService(t *testing.T) {
	suite.Run(t, new(managedSubAccountWithdrawServiceTestSuite))
}

func (s *managedSubAccountWithdrawServiceTestSuite) TestManagedSubAccountWithdraw() {
	data := []byte(`{
    "tranId":66157362489
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var fromEmail string = "xxyyzz@gmail.com"
	var asset string = "USDT"
	var amount string = "1.000548036"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("fromEmail", fromEmail).setFormParam("asset", asset).setFormParam("amount", amount)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountWithdrawService().FromEmail(fromEmail).Asset(asset).Amount(amount).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountWithdrawServiceResponse{
		TranId: 66157362489}
	s.assertManagedSubAccountWithdrawServiceResponseEqual(e, res)
}

func (s *managedSubAccountWithdrawServiceTestSuite) assertManagedSubAccountWithdrawServiceResponseEqual(e, a *ManagedSubAccountWithdrawServiceResponse) {
	r := s.r()
	r.Equal(e.TranId, a.TranId, "TranId")
}

type managedSubAccountSnapshotServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountSnapshotService(t *testing.T) {
	suite.Run(t, new(managedSubAccountSnapshotServiceTestSuite))
}

func (s *managedSubAccountSnapshotServiceTestSuite) TestManagedSubAccountSnapshot() {
	data := []byte(`{
    "code": 200,
    "msg": "",
    "snapshotVos": [{
        "data": {
            "balances": [{
                    "asset": "BTC",
                    "free": "0.09905021",
                    "locked": "0.00000000"
                },
                {
                    "asset": "USDT",
                    "free": "1.89109409",
                    "locked": "0.00000000"
                }
            ],
            "totalAssetOfBtc": "0.09942700",

            "marginLevel": "2748.02909813",
            "totalLiabilityOfBtc": "0.00000100",
            "totalNetAssetOfBtc": "0.00274750",
            "userAssets": [{
                "asset": "XRP",
                "borrowed": "0.00000000",
                "free": "1.00000000",
                "interest": "0.00000000",
                "locked": "0.00000000",
                "netAsset": "1.00000000"
            }],

            "assets": [{
                "asset": "USDT",
                "marginBalance": "118.99782335",
                "walletBalance": "120.23811389"
            }],
            "position": [{
                "entryPrice": "7130.41000000",
                "markPrice": "7257.66239673",
                "positionAmt": "0.01000000",
                "symbol": "BTCUSDT",
                "unRealizedProfit": "1.24029054"
            }]
        },
        "type": "spot",
        "updateTime": 1576281599000
    }]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var accType string = "USDT"
	var limit int32 = 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("type", accType).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountSnapshotService().Email(email).AccType(accType).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountSnapshotServiceResponse{
		Code: 200,
		Msg:  "",
		SnapshotVos: []*SnapshotVo{{
			Data: SnapshotVoData{
				Balances: []*SnapShotSpotBalance{{
					Asset:  "BTC",
					Free:   "0.09905021",
					Locked: "0.00000000"},
					{
						Asset:  "USDT",
						Free:   "1.89109409",
						Locked: "0.00000000"}},
				TotalAssetOfBtc:     "0.09942700",
				MarginLevel:         "2748.02909813",
				TotalLiabilityOfBtc: "0.00000100",
				TotalNetAssetOfBtc:  "0.00274750",
				UserAssets: []*MarginUserAsset{{
					Asset:    "XRP",
					Borrowed: "0.00000000",
					Free:     "1.00000000",
					Interest: "0.00000000",
					Locked:   "0.00000000",
					NetAsset: "1.00000000"}},
				Assets: []*FuturesUserAsset{{
					Asset:         "USDT",
					MarginBalance: "118.99782335",
					WalletBalance: "120.23811389"}},
				Position: []*FuturesUserPosition{{
					EntryPrice:       "7130.41000000",
					MarkPrice:        "7257.66239673",
					PositionAmt:      "0.01000000",
					Symbol:           "BTCUSDT",
					UnRealizedProfit: "1.24029054"}}},
			Type:       "spot",
			UpdateTime: 1576281599000}}}
	s.assertManagedSubAccountSnapshotServiceResponseEqual(e, res)
}

func (s *managedSubAccountSnapshotServiceTestSuite) assertManagedSubAccountSnapshotServiceResponseEqual(e, a *ManagedSubAccountSnapshotServiceResponse) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
	r.Len(e.SnapshotVos, len(a.SnapshotVos))
	for i := range e.SnapshotVos {
		r.Len(e.SnapshotVos[i].Data.Balances, len(a.SnapshotVos[i].Data.Balances))
		for k := range e.SnapshotVos[i].Data.Balances {
			r.Equal(e.SnapshotVos[i].Data.Balances[k].Asset, a.SnapshotVos[i].Data.Balances[k].Asset, "SnapshotVos[i].Data.Balances[k].Asset")
			r.Equal(e.SnapshotVos[i].Data.Balances[k].Free, a.SnapshotVos[i].Data.Balances[k].Free, "SnapshotVos[i].Data.Balances[k].Free")
			r.Equal(e.SnapshotVos[i].Data.Balances[k].Locked, a.SnapshotVos[i].Data.Balances[k].Locked, "SnapshotVos[i].Data.Balances[k].Locked")
		}

		r.Equal(e.SnapshotVos[i].Data.TotalAssetOfBtc, a.SnapshotVos[i].Data.TotalAssetOfBtc, "SnapshotVos[i].Data.TotalAssetOfBtc")
		r.Equal(e.SnapshotVos[i].Data.MarginLevel, a.SnapshotVos[i].Data.MarginLevel, "SnapshotVos[i].Data.MarginLevel")
		r.Equal(e.SnapshotVos[i].Data.TotalLiabilityOfBtc, a.SnapshotVos[i].Data.TotalLiabilityOfBtc, "SnapshotVos[i].Data.TotalLiabilityOfBtc")
		r.Equal(e.SnapshotVos[i].Data.TotalNetAssetOfBtc, a.SnapshotVos[i].Data.TotalNetAssetOfBtc, "SnapshotVos[i].Data.TotalNetAssetOfBtc")
		r.Len(e.SnapshotVos[i].Data.UserAssets, len(a.SnapshotVos[i].Data.UserAssets))
		for k := range e.SnapshotVos[i].Data.UserAssets {
			r.Equal(e.SnapshotVos[i].Data.UserAssets[k].Asset, a.SnapshotVos[i].Data.UserAssets[k].Asset, "SnapshotVos[i].Data.UserAssets[k].Asset")
			r.Equal(e.SnapshotVos[i].Data.UserAssets[k].Borrowed, a.SnapshotVos[i].Data.UserAssets[k].Borrowed, "SnapshotVos[i].Data.UserAssets[k].Borrowed")
			r.Equal(e.SnapshotVos[i].Data.UserAssets[k].Free, a.SnapshotVos[i].Data.UserAssets[k].Free, "SnapshotVos[i].Data.UserAssets[k].Free")
			r.Equal(e.SnapshotVos[i].Data.UserAssets[k].Interest, a.SnapshotVos[i].Data.UserAssets[k].Interest, "SnapshotVos[i].Data.UserAssets[k].Interest")
			r.Equal(e.SnapshotVos[i].Data.UserAssets[k].Locked, a.SnapshotVos[i].Data.UserAssets[k].Locked, "SnapshotVos[i].Data.UserAssets[k].Locked")
			r.Equal(e.SnapshotVos[i].Data.UserAssets[k].NetAsset, a.SnapshotVos[i].Data.UserAssets[k].NetAsset, "SnapshotVos[i].Data.UserAssets[k].NetAsset")
		}

		r.Len(e.SnapshotVos[i].Data.Assets, len(a.SnapshotVos[i].Data.Assets))
		for k := range e.SnapshotVos[i].Data.Assets {
			r.Equal(e.SnapshotVos[i].Data.Assets[k].Asset, a.SnapshotVos[i].Data.Assets[k].Asset, "SnapshotVos[i].Data.Assets[k].Asset")
			r.Equal(e.SnapshotVos[i].Data.Assets[k].MarginBalance, a.SnapshotVos[i].Data.Assets[k].MarginBalance, "SnapshotVos[i].Data.Assets[k].MarginBalance")
			r.Equal(e.SnapshotVos[i].Data.Assets[k].WalletBalance, a.SnapshotVos[i].Data.Assets[k].WalletBalance, "SnapshotVos[i].Data.Assets[k].WalletBalance")
		}

		r.Len(e.SnapshotVos[i].Data.Position, len(a.SnapshotVos[i].Data.Position))
		for k := range e.SnapshotVos[i].Data.Position {
			r.Equal(e.SnapshotVos[i].Data.Position[k].EntryPrice, a.SnapshotVos[i].Data.Position[k].EntryPrice, "SnapshotVos[i].Data.Position[k].EntryPrice")
			r.Equal(e.SnapshotVos[i].Data.Position[k].MarkPrice, a.SnapshotVos[i].Data.Position[k].MarkPrice, "SnapshotVos[i].Data.Position[k].MarkPrice")
			r.Equal(e.SnapshotVos[i].Data.Position[k].PositionAmt, a.SnapshotVos[i].Data.Position[k].PositionAmt, "SnapshotVos[i].Data.Position[k].PositionAmt")
			r.Equal(e.SnapshotVos[i].Data.Position[k].Symbol, a.SnapshotVos[i].Data.Position[k].Symbol, "SnapshotVos[i].Data.Position[k].Symbol")
			r.Equal(e.SnapshotVos[i].Data.Position[k].UnRealizedProfit, a.SnapshotVos[i].Data.Position[k].UnRealizedProfit, "SnapshotVos[i].Data.Position[k].UnRealizedProfit")
		}

		r.Equal(e.SnapshotVos[i].Type, a.SnapshotVos[i].Type, "SnapshotVos[i].Type")
		r.Equal(e.SnapshotVos[i].UpdateTime, a.SnapshotVos[i].UpdateTime, "SnapshotVos[i].UpdateTime")
	}

}

type managedSubAccountQueryTransferLogForInvestorServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountQueryTransferLogForInvestorService(t *testing.T) {
	suite.Run(t, new(managedSubAccountQueryTransferLogForInvestorServiceTestSuite))
}

func (s *managedSubAccountQueryTransferLogForInvestorServiceTestSuite) TestManagedSubAccountQueryTransferLogForInvestor() {
	data := []byte(`{
    "managerSubTransferHistoryVos": [{
            "fromEmail": "test_0_virtual@kq3kno9imanagedsub.com",
            "fromAccountType": "SPOT",
            "toEmail": "wdywl0lddakh@test.com",
            "toAccountType": "SPOT",
            "asset": "BNB",
            "amount": "0.01",
            "scheduledData": 1679416673000,
            "createTime": 1679416673000,
            "status": "SUCCESS",
            "tranId": 91077779
        },
        {
            "fromEmail": "wdywl0lddakh@test.com",
            "fromAccountType": "SPOT",
            "toEmail": "test_0_virtual@kq3kno9imanagedsub.com",
            "toAccountType": "SPOT",
            "asset": "BNB",
            "amount": "1",
            "scheduledData": 1679416616000,
            "createTime": 1679416616000,
            "status": "SUCCESS",
            "tranId": 91077676
        }
    ],
    "count": 2
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var startTime int64 = 1111111235466
	var endTime int64 = 2222233355456
	var page int32 = 1
	var limit int32 = 20
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("startTime", startTime).setParam("endTime", endTime).setParam("page", page).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountQueryTransferLogForInvestorService().Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountQueryTransferLogForInvestorServiceResponse{
		ManagerSubTransferHistoryVos: []*ManagedSubTransferHistoryVo{{
			FromEmail:       "test_0_virtual@kq3kno9imanagedsub.com",
			FromAccountType: "SPOT",
			ToEmail:         "wdywl0lddakh@test.com",
			ToAccountType:   "SPOT",
			Asset:           "BNB",
			Amount:          "0.01",
			ScheduledData:   1679416673000,
			CreateTime:      1679416673000,
			Status:          "SUCCESS",
			TranId:          91077779},
			{
				FromEmail:       "wdywl0lddakh@test.com",
				FromAccountType: "SPOT",
				ToEmail:         "test_0_virtual@kq3kno9imanagedsub.com",
				ToAccountType:   "SPOT",
				Asset:           "BNB",
				Amount:          "1",
				ScheduledData:   1679416616000,
				CreateTime:      1679416616000,
				Status:          "SUCCESS",
				TranId:          91077676}},
		Count: 2}
	s.assertManagedSubAccountQueryTransferLogForInvestorServiceResponseEqual(e, res)
}

func (s *managedSubAccountQueryTransferLogForInvestorServiceTestSuite) assertManagedSubAccountQueryTransferLogForInvestorServiceResponseEqual(e, a *ManagedSubAccountQueryTransferLogForInvestorServiceResponse) {
	r := s.r()
	r.Len(e.ManagerSubTransferHistoryVos, len(a.ManagerSubTransferHistoryVos))
	for i := range e.ManagerSubTransferHistoryVos {
		r.Equal(e.ManagerSubTransferHistoryVos[i].FromEmail, a.ManagerSubTransferHistoryVos[i].FromEmail, "ManagerSubTransferHistoryVos[i].FromEmail")
		r.Equal(e.ManagerSubTransferHistoryVos[i].FromAccountType, a.ManagerSubTransferHistoryVos[i].FromAccountType, "ManagerSubTransferHistoryVos[i].FromAccountType")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ToEmail, a.ManagerSubTransferHistoryVos[i].ToEmail, "ManagerSubTransferHistoryVos[i].ToEmail")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ToAccountType, a.ManagerSubTransferHistoryVos[i].ToAccountType, "ManagerSubTransferHistoryVos[i].ToAccountType")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Asset, a.ManagerSubTransferHistoryVos[i].Asset, "ManagerSubTransferHistoryVos[i].Asset")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Amount, a.ManagerSubTransferHistoryVos[i].Amount, "ManagerSubTransferHistoryVos[i].Amount")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ScheduledData, a.ManagerSubTransferHistoryVos[i].ScheduledData, "ManagerSubTransferHistoryVos[i].ScheduledData")
		r.Equal(e.ManagerSubTransferHistoryVos[i].CreateTime, a.ManagerSubTransferHistoryVos[i].CreateTime, "ManagerSubTransferHistoryVos[i].CreateTime")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Status, a.ManagerSubTransferHistoryVos[i].Status, "ManagerSubTransferHistoryVos[i].Status")
		r.Equal(e.ManagerSubTransferHistoryVos[i].TranId, a.ManagerSubTransferHistoryVos[i].TranId, "ManagerSubTransferHistoryVos[i].TranId")
	}

	r.Equal(e.Count, a.Count, "Count")
}

type managedSubAccountQueryTransferLogForTradeParentServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountQueryTransferLogForTradeParentService(t *testing.T) {
	suite.Run(t, new(managedSubAccountQueryTransferLogForTradeParentServiceTestSuite))
}

func (s *managedSubAccountQueryTransferLogForTradeParentServiceTestSuite) TestManagedSubAccountQueryTransferLogForTradeParent() {
	data := []byte(`{
    "managerSubTransferHistoryVos": [{
            "fromEmail": "test_0_virtual@kq3kno9imanagedsub.com",
            "fromAccountType": "SPOT",
            "toEmail": "wdywl0lddakh@test.com",
            "toAccountType": "SPOT",
            "asset": "BNB",
            "amount": "0.01",
            "scheduledData": 1679416673000,
            "createTime": 1679416673000,
            "status": "SUCCESS",
            "tranId": 91077779
        },
        {
            "fromEmail": "wdywl0lddakh@test.com",
            "fromAccountType": "SPOT",
            "toEmail": "test_0_virtual@kq3kno9imanagedsub.com",
            "toAccountType": "SPOT",
            "asset": "BNB",
            "amount": "1",
            "scheduledData": 1679416616000,
            "createTime": 1679416616000,
            "status": "SUCCESS",
            "tranId": 91077676
        }
    ],
    "count": 2
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var startTime int64 = 1111111235466
	var endTime int64 = 2222233355456
	var page int32 = 1
	var limit int32 = 20
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("startTime", startTime).setParam("endTime", endTime).setParam("page", page).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountQueryTransferLogForTradeParentService().Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountQueryTransferLogForTradeParentServiceResponse{
		ManagerSubTransferHistoryVos: []*ManagedSubTransferHistoryVo{{
			FromEmail:       "test_0_virtual@kq3kno9imanagedsub.com",
			FromAccountType: "SPOT",
			ToEmail:         "wdywl0lddakh@test.com",
			ToAccountType:   "SPOT",
			Asset:           "BNB",
			Amount:          "0.01",
			ScheduledData:   1679416673000,
			CreateTime:      1679416673000,
			Status:          "SUCCESS",
			TranId:          91077779},
			{
				FromEmail:       "wdywl0lddakh@test.com",
				FromAccountType: "SPOT",
				ToEmail:         "test_0_virtual@kq3kno9imanagedsub.com",
				ToAccountType:   "SPOT",
				Asset:           "BNB",
				Amount:          "1",
				ScheduledData:   1679416616000,
				CreateTime:      1679416616000,
				Status:          "SUCCESS",
				TranId:          91077676}},
		Count: 2}
	s.assertManagedSubAccountQueryTransferLogForTradeParentServiceResponseEqual(e, res)
}

func (s *managedSubAccountQueryTransferLogForTradeParentServiceTestSuite) assertManagedSubAccountQueryTransferLogForTradeParentServiceResponseEqual(e, a *ManagedSubAccountQueryTransferLogForTradeParentServiceResponse) {
	r := s.r()
	r.Len(e.ManagerSubTransferHistoryVos, len(a.ManagerSubTransferHistoryVos))
	for i := range e.ManagerSubTransferHistoryVos {
		r.Equal(e.ManagerSubTransferHistoryVos[i].FromEmail, a.ManagerSubTransferHistoryVos[i].FromEmail, "ManagerSubTransferHistoryVos[i].FromEmail")
		r.Equal(e.ManagerSubTransferHistoryVos[i].FromAccountType, a.ManagerSubTransferHistoryVos[i].FromAccountType, "ManagerSubTransferHistoryVos[i].FromAccountType")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ToEmail, a.ManagerSubTransferHistoryVos[i].ToEmail, "ManagerSubTransferHistoryVos[i].ToEmail")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ToAccountType, a.ManagerSubTransferHistoryVos[i].ToAccountType, "ManagerSubTransferHistoryVos[i].ToAccountType")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Asset, a.ManagerSubTransferHistoryVos[i].Asset, "ManagerSubTransferHistoryVos[i].Asset")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Amount, a.ManagerSubTransferHistoryVos[i].Amount, "ManagerSubTransferHistoryVos[i].Amount")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ScheduledData, a.ManagerSubTransferHistoryVos[i].ScheduledData, "ManagerSubTransferHistoryVos[i].ScheduledData")
		r.Equal(e.ManagerSubTransferHistoryVos[i].CreateTime, a.ManagerSubTransferHistoryVos[i].CreateTime, "ManagerSubTransferHistoryVos[i].CreateTime")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Status, a.ManagerSubTransferHistoryVos[i].Status, "ManagerSubTransferHistoryVos[i].Status")
		r.Equal(e.ManagerSubTransferHistoryVos[i].TranId, a.ManagerSubTransferHistoryVos[i].TranId, "ManagerSubTransferHistoryVos[i].TranId")
	}

	r.Equal(e.Count, a.Count, "Count")
}

type managedSubAccountQueryFuturesAssetServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountQueryFuturesAssetService(t *testing.T) {
	suite.Run(t, new(managedSubAccountQueryFuturesAssetServiceTestSuite))
}

func (s *managedSubAccountQueryFuturesAssetServiceTestSuite) TestManagedSubAccountQueryFuturesAsset() {
	data := []byte(`{
  "code": 200,
  "message": "OK",
  "snapshotVos": [
    {
      "type": "FUTURES",
      "updateTime": 1672893855394,
      "data": {
        "assets": [
          {
            "asset": "USDT",
            "marginBalance": "100",
            "walletBalance": "120"
          }
        ],
        "position": [
          {
            "symbol": "BTCUSDT",
            "entryPrice": "17000",
            "markPrice": "17000",
            "positionAmt": "0.0001"
          }
        ]
      }
    }
  ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountQueryFuturesAssetService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountQueryFuturesAssetServiceResponse{
		Code:    200,
		Message: "OK",
		SnapshotVos: []*ManagedSubFuturesAccountSnapVo{{
			Type:       "FUTURES",
			UpdateTime: 1672893855394,
			Data: &ManagedSubFuturesAccountSnapVoData{
				Assets: []*ManagedSubFuturesAccountSnapVoDataAsset{{
					Asset:         "USDT",
					MarginBalance: "100",
					WalletBalance: "120"}},
				Position: []*ManagedSubFuturesAccountSnapVoDataPosition{{
					Symbol:      "BTCUSDT",
					EntryPrice:  "17000",
					MarkPrice:   "17000",
					PositionAmt: "0.0001"}}}}}}
	s.assertManagedSubAccountQueryFuturesAssetServiceResponseEqual(e, res)
}

func (s *managedSubAccountQueryFuturesAssetServiceTestSuite) assertManagedSubAccountQueryFuturesAssetServiceResponseEqual(e, a *ManagedSubAccountQueryFuturesAssetServiceResponse) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Message, a.Message, "Message")
	r.Len(e.SnapshotVos, len(a.SnapshotVos))
	for i := range e.SnapshotVos {
		r.Equal(e.SnapshotVos[i].Type, a.SnapshotVos[i].Type, "SnapshotVos[i].Type")
		r.Equal(e.SnapshotVos[i].UpdateTime, a.SnapshotVos[i].UpdateTime, "SnapshotVos[i].UpdateTime")
		r.Len(e.SnapshotVos[i].Data.Assets, len(a.SnapshotVos[i].Data.Assets))
		for k := range e.SnapshotVos[i].Data.Assets {
			r.Equal(e.SnapshotVos[i].Data.Assets[k].Asset, a.SnapshotVos[i].Data.Assets[k].Asset, "SnapshotVos[i].Data.Assets[k].Asset")
			r.Equal(e.SnapshotVos[i].Data.Assets[k].MarginBalance, a.SnapshotVos[i].Data.Assets[k].MarginBalance, "SnapshotVos[i].Data.Assets[k].MarginBalance")
			r.Equal(e.SnapshotVos[i].Data.Assets[k].WalletBalance, a.SnapshotVos[i].Data.Assets[k].WalletBalance, "SnapshotVos[i].Data.Assets[k].WalletBalance")
		}

		r.Len(e.SnapshotVos[i].Data.Position, len(a.SnapshotVos[i].Data.Position))
		for k := range e.SnapshotVos[i].Data.Position {
			r.Equal(e.SnapshotVos[i].Data.Position[k].Symbol, a.SnapshotVos[i].Data.Position[k].Symbol, "SnapshotVos[i].Data.Position[k].Symbol")
			r.Equal(e.SnapshotVos[i].Data.Position[k].EntryPrice, a.SnapshotVos[i].Data.Position[k].EntryPrice, "SnapshotVos[i].Data.Position[k].EntryPrice")
			r.Equal(e.SnapshotVos[i].Data.Position[k].MarkPrice, a.SnapshotVos[i].Data.Position[k].MarkPrice, "SnapshotVos[i].Data.Position[k].MarkPrice")
			r.Equal(e.SnapshotVos[i].Data.Position[k].PositionAmt, a.SnapshotVos[i].Data.Position[k].PositionAmt, "SnapshotVos[i].Data.Position[k].PositionAmt")
		}

	}

}

type managedSubAccountQueryMarginAssetServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountQueryMarginAssetService(t *testing.T) {
	suite.Run(t, new(managedSubAccountQueryMarginAssetServiceTestSuite))
}

func (s *managedSubAccountQueryMarginAssetServiceTestSuite) TestManagedSubAccountQueryMarginAsset() {
	data := []byte(`{
    "marginLevel": "999",
    "totalAssetOfBtc": "0",
    "totalLiabilityOfBtc": "0",
    "totalNetAssetOfBtc": "0",
    "userAssets": [{
        "asset": "MATIC",
        "borrowed": "0",
        "free": "0",
        "interest": "0",
        "locked": "0",
        "netAsset": "0"
    }]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountQueryMarginAssetService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountQueryMarginAssetServiceResponse{
		MarginLevel:         "999",
		TotalAssetOfBtc:     "0",
		TotalLiabilityOfBtc: "0",
		TotalNetAssetOfBtc:  "0",
		UserAssets: []*ManagedSubAccountMarginAsset{{
			Asset:    "MATIC",
			Borrowed: "0",
			Free:     "0",
			Interest: "0",
			Locked:   "0",
			NetAsset: "0"}}}
	s.assertManagedSubAccountQueryMarginAssetServiceResponseEqual(e, res)
}

func (s *managedSubAccountQueryMarginAssetServiceTestSuite) assertManagedSubAccountQueryMarginAssetServiceResponseEqual(e, a *ManagedSubAccountQueryMarginAssetServiceResponse) {
	r := s.r()
	r.Equal(e.MarginLevel, a.MarginLevel, "MarginLevel")
	r.Equal(e.TotalAssetOfBtc, a.TotalAssetOfBtc, "TotalAssetOfBtc")
	r.Equal(e.TotalLiabilityOfBtc, a.TotalLiabilityOfBtc, "TotalLiabilityOfBtc")
	r.Equal(e.TotalNetAssetOfBtc, a.TotalNetAssetOfBtc, "TotalNetAssetOfBtc")
	r.Len(e.UserAssets, len(a.UserAssets))
	for i := range e.UserAssets {
		r.Equal(e.UserAssets[i].Asset, a.UserAssets[i].Asset, "UserAssets[i].Asset")
		r.Equal(e.UserAssets[i].Borrowed, a.UserAssets[i].Borrowed, "UserAssets[i].Borrowed")
		r.Equal(e.UserAssets[i].Free, a.UserAssets[i].Free, "UserAssets[i].Free")
		r.Equal(e.UserAssets[i].Interest, a.UserAssets[i].Interest, "UserAssets[i].Interest")
		r.Equal(e.UserAssets[i].Locked, a.UserAssets[i].Locked, "UserAssets[i].Locked")
		r.Equal(e.UserAssets[i].NetAsset, a.UserAssets[i].NetAsset, "UserAssets[i].NetAsset")
	}

}

type subAccountAssetServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountAssetService(t *testing.T) {
	suite.Run(t, new(subAccountAssetServiceTestSuite))
}

func (s *subAccountAssetServiceTestSuite) TestSubAccountAsset() {
	data := []byte(`{
    "balances":[
        {
            "asset":"ADA",
            "free":"10000",
            "locked":"0"
        },
        {
            "asset":"BNB",
            "free":"10003",
            "locked":"0"
        },
        {
            "asset":"BTC",
            "free":"11467.6399",
            "locked":"0"
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountAssetService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountAssetServiceResponse{
		Balances: []*SubAccountAssetBalance{{
			Asset:  "ADA",
			Free:   "10000",
			Locked: "0"},
			{
				Asset:  "BNB",
				Free:   "10003",
				Locked: "0"},
			{
				Asset:  "BTC",
				Free:   "11467.6399",
				Locked: "0"}}}
	s.assertSubAccountAssetServiceResponseEqual(e, res)
}

func (s *subAccountAssetServiceTestSuite) assertSubAccountAssetServiceResponseEqual(e, a *SubAccountAssetServiceResponse) {
	r := s.r()
	r.Len(e.Balances, len(a.Balances))
	for i := range e.Balances {
		r.Equal(e.Balances[i].Asset, a.Balances[i].Asset, "Balances[i].Asset")
		r.Equal(e.Balances[i].Free, a.Balances[i].Free, "Balances[i].Free")
		r.Equal(e.Balances[i].Locked, a.Balances[i].Locked, "Balances[i].Locked")
	}

}

type managedSubAccountInfoServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountInfoService(t *testing.T) {
	suite.Run(t, new(managedSubAccountInfoServiceTestSuite))
}

func (s *managedSubAccountInfoServiceTestSuite) TestManagedSubAccountInfo() {
	data := []byte(`{
    "total": 3,
    "managerSubUserInfoVoList": [
        {
            "rootUserId": 1000138475670,
            "managersubUserId": 1000137842513,
            "bindParentUserId": 1000138475669,
            "email": "test_0_virtual@kq3kno9imanagedsub.com",
            "insertTimeStamp": 1678435149000,
            "bindParentEmail": "wdyw8xsh8pey@test.com",
            "isSubUserEnabled": true,
            "isUserActive": true,
            "isMarginEnabled": false,
            "isFutureEnabled": false,
            "isSignedLVTRiskAgreement": false
        },
        {
            "rootUserId": 1000138475670,
            "managersubUserId": 1000137842514,
            "bindParentUserId": 1000138475669,
            "email": "test_1_virtual@4qd2u7zxmanagedsub.com",
            "insertTimeStamp": 1678435152000,
            "bindParentEmail": "wdyw8xsh8pey@test.com",
            "isSubUserEnabled": true,
            "isUserActive": true,
            "isMarginEnabled": false,
            "isFutureEnabled": false,
            "isSignedLVTRiskAgreement": false
        },
        {
            "rootUserId": 1000138475670,
            "managersubUserId": 1000137842515,
            "bindParentUserId": 1000138475669,
            "email": "test_2_virtual@akc05o8hmanagedsub.com",
            "insertTimeStamp": 1678435153000,
            "bindParentEmail": "wdyw8xsh8pey@test.com",
            "isSubUserEnabled": true,
            "isUserActive": true,
            "isMarginEnabled": false,
            "isFutureEnabled": false,
            "isSignedLVTRiskAgreement": false
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountInfoService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountInfoServiceResponse{
		Total: 3,
		ManagerSubUserInfoVoList: []*ManagedSubAccountUserInfoVo{{
			RootUserId:               1000138475670,
			ManagersubUserId:         1000137842513,
			BindParentUserId:         1000138475669,
			Email:                    "test_0_virtual@kq3kno9imanagedsub.com",
			InsertTimeStamp:          1678435149000,
			BindParentEmail:          "wdyw8xsh8pey@test.com",
			IsSubUserEnabled:         true,
			IsUserActive:             true,
			IsMarginEnabled:          false,
			IsFutureEnabled:          false,
			IsSignedLVTRiskAgreement: false},
			{
				RootUserId:               1000138475670,
				ManagersubUserId:         1000137842514,
				BindParentUserId:         1000138475669,
				Email:                    "test_1_virtual@4qd2u7zxmanagedsub.com",
				InsertTimeStamp:          1678435152000,
				BindParentEmail:          "wdyw8xsh8pey@test.com",
				IsSubUserEnabled:         true,
				IsUserActive:             true,
				IsMarginEnabled:          false,
				IsFutureEnabled:          false,
				IsSignedLVTRiskAgreement: false},
			{
				RootUserId:               1000138475670,
				ManagersubUserId:         1000137842515,
				BindParentUserId:         1000138475669,
				Email:                    "test_2_virtual@akc05o8hmanagedsub.com",
				InsertTimeStamp:          1678435153000,
				BindParentEmail:          "wdyw8xsh8pey@test.com",
				IsSubUserEnabled:         true,
				IsUserActive:             true,
				IsMarginEnabled:          false,
				IsFutureEnabled:          false,
				IsSignedLVTRiskAgreement: false}}}
	s.assertManagedSubAccountInfoServiceResponseEqual(e, res)
}

func (s *managedSubAccountInfoServiceTestSuite) assertManagedSubAccountInfoServiceResponseEqual(e, a *ManagedSubAccountInfoServiceResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(e.ManagerSubUserInfoVoList, len(a.ManagerSubUserInfoVoList))
	for i := range e.ManagerSubUserInfoVoList {
		r.Equal(e.ManagerSubUserInfoVoList[i].RootUserId, a.ManagerSubUserInfoVoList[i].RootUserId, "ManagerSubUserInfoVoList[i].RootUserId")
		r.Equal(e.ManagerSubUserInfoVoList[i].ManagersubUserId, a.ManagerSubUserInfoVoList[i].ManagersubUserId, "ManagerSubUserInfoVoList[i].ManagersubUserId")
		r.Equal(e.ManagerSubUserInfoVoList[i].BindParentUserId, a.ManagerSubUserInfoVoList[i].BindParentUserId, "ManagerSubUserInfoVoList[i].BindParentUserId")
		r.Equal(e.ManagerSubUserInfoVoList[i].Email, a.ManagerSubUserInfoVoList[i].Email, "ManagerSubUserInfoVoList[i].Email")
		r.Equal(e.ManagerSubUserInfoVoList[i].InsertTimeStamp, a.ManagerSubUserInfoVoList[i].InsertTimeStamp, "ManagerSubUserInfoVoList[i].InsertTimeStamp")
		r.Equal(e.ManagerSubUserInfoVoList[i].BindParentEmail, a.ManagerSubUserInfoVoList[i].BindParentEmail, "ManagerSubUserInfoVoList[i].BindParentEmail")
		r.Equal(e.ManagerSubUserInfoVoList[i].IsSubUserEnabled, a.ManagerSubUserInfoVoList[i].IsSubUserEnabled, "ManagerSubUserInfoVoList[i].IsSubUserEnabled")
		r.Equal(e.ManagerSubUserInfoVoList[i].IsUserActive, a.ManagerSubUserInfoVoList[i].IsUserActive, "ManagerSubUserInfoVoList[i].IsUserActive")
		r.Equal(e.ManagerSubUserInfoVoList[i].IsMarginEnabled, a.ManagerSubUserInfoVoList[i].IsMarginEnabled, "ManagerSubUserInfoVoList[i].IsMarginEnabled")
		r.Equal(e.ManagerSubUserInfoVoList[i].IsFutureEnabled, a.ManagerSubUserInfoVoList[i].IsFutureEnabled, "ManagerSubUserInfoVoList[i].IsFutureEnabled")
		r.Equal(e.ManagerSubUserInfoVoList[i].IsSignedLVTRiskAgreement, a.ManagerSubUserInfoVoList[i].IsSignedLVTRiskAgreement, "ManagerSubUserInfoVoList[i].IsSignedLVTRiskAgreement")
	}

}

type subAccountTransactionStatisticsServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountTransactionStatisticsService(t *testing.T) {
	suite.Run(t, new(subAccountTransactionStatisticsServiceTestSuite))
}

func (s *subAccountTransactionStatisticsServiceTestSuite) TestSubAccountTransactionStatistics() {
	data := []byte(`{
    "recent30BtcTotal": "0",
    "recent30BtcFuturesTotal": "0",
    "recent30BtcMarginTotal": "0",
    "recent30BusdTotal": "0",
    "recent30BusdFuturesTotal": "0",
    "recent30BusdMarginTotal": "0",
    "tradeInfoVos": [
        {
            "userId": 1000138138384,
            "btc": 0,
            "btcFutures": 0,
            "btcMargin": 0,
            "busd": 0,
            "busdFutures": 0,
            "busdMargin": 0,
            "date": 1676851200000
        },
        {
            "userId": 1000138138384,
            "btc": 0,
            "btcFutures": 0,
            "btcMargin": 0,
            "busd": 0,
            "busdFutures": 0,
            "busdMargin": 0,
            "date": 1676937600000
        }]}
        `)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountTransactionStatisticsService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountTransactionStatisticServiceResponse{
		Recent30BtcTotal:         "0",
		Recent30BtcFuturesTotal:  "0",
		Recent30BtcMarginTotal:   "0",
		Recent30BusdTotal:        "0",
		Recent30BusdFuturesTotal: "0",
		Recent30BusdMarginTotal:  "0",
		TradeInfoVos: []*TradeInfoVo{{
			UserId:      1000138138384,
			Btc:         0,
			BtcFutures:  0,
			BtcMargin:   0,
			Busd:        0,
			BusdFutures: 0,
			BusdMargin:  0,
			Date:        1676851200000},
			{
				UserId:      1000138138384,
				Btc:         0,
				BtcFutures:  0,
				BtcMargin:   0,
				Busd:        0,
				BusdFutures: 0,
				BusdMargin:  0,
				Date:        1676937600000}}}
	s.assertSubAccountTransactionStatisticServiceResponseEqual(e, res)
}

func (s *subAccountTransactionStatisticsServiceTestSuite) assertSubAccountTransactionStatisticServiceResponseEqual(e, a *SubAccountTransactionStatisticServiceResponse) {
	r := s.r()
	r.Equal(e.Recent30BtcTotal, a.Recent30BtcTotal, "Recent30BtcTotal")
	r.Equal(e.Recent30BtcFuturesTotal, a.Recent30BtcFuturesTotal, "Recent30BtcFuturesTotal")
	r.Equal(e.Recent30BtcMarginTotal, a.Recent30BtcMarginTotal, "Recent30BtcMarginTotal")
	r.Equal(e.Recent30BusdTotal, a.Recent30BusdTotal, "Recent30BusdTotal")
	r.Equal(e.Recent30BusdFuturesTotal, a.Recent30BusdFuturesTotal, "Recent30BusdFuturesTotal")
	r.Equal(e.Recent30BusdMarginTotal, a.Recent30BusdMarginTotal, "Recent30BusdMarginTotal")
	r.Len(e.TradeInfoVos, len(a.TradeInfoVos))
	for i := range e.TradeInfoVos {
		r.Equal(e.TradeInfoVos[i].UserId, a.TradeInfoVos[i].UserId, "TradeInfoVos[i].UserId")
		r.Equal(e.TradeInfoVos[i].Btc, a.TradeInfoVos[i].Btc, "TradeInfoVos[i].Btc")
		r.Equal(e.TradeInfoVos[i].BtcFutures, a.TradeInfoVos[i].BtcFutures, "TradeInfoVos[i].BtcFutures")
		r.Equal(e.TradeInfoVos[i].BtcMargin, a.TradeInfoVos[i].BtcMargin, "TradeInfoVos[i].BtcMargin")
		r.Equal(e.TradeInfoVos[i].Busd, a.TradeInfoVos[i].Busd, "TradeInfoVos[i].Busd")
		r.Equal(e.TradeInfoVos[i].BusdFutures, a.TradeInfoVos[i].BusdFutures, "TradeInfoVos[i].BusdFutures")
		r.Equal(e.TradeInfoVos[i].BusdMargin, a.TradeInfoVos[i].BusdMargin, "TradeInfoVos[i].BusdMargin")
		r.Equal(e.TradeInfoVos[i].Date, a.TradeInfoVos[i].Date, "TradeInfoVos[i].Date")
	}

}

type managedSubAccountDepositAddressServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountDepositAddressService(t *testing.T) {
	suite.Run(t, new(managedSubAccountDepositAddressServiceTestSuite))
}

func (s *managedSubAccountDepositAddressServiceTestSuite) TestManagedSubAccountDepositAddress() {
	data := []byte(`{
    "coin": "USDT",
    "address": "0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d",
    "tag": "",
    "url": "https://etherscan.io/address/0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	var coin string = "USDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("email", email).setParam("coin", coin)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountDepositAddressService().Email(email).Coin(coin).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountDepositAddressServiceResponse{
		Coin:    "USDT",
		Address: "0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d",
		Tag:     "",
		Url:     "https://etherscan.io/address/0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d"}
	s.assertManagedSubAccountDepositAddressServiceResponseEqual(e, res)
}

func (s *managedSubAccountDepositAddressServiceTestSuite) assertManagedSubAccountDepositAddressServiceResponseEqual(e, a *ManagedSubAccountDepositAddressServiceResponse) {
	r := s.r()
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.Tag, a.Tag, "Tag")
	r.Equal(e.Url, a.Url, "Url")
}

type subAccountOptionsEnableServiceTestSuite struct {
	baseTestSuite
}

func TestSubAccountOptionsEnableService(t *testing.T) {
	suite.Run(t, new(subAccountOptionsEnableServiceTestSuite))
}

func (s *subAccountOptionsEnableServiceTestSuite) TestSubAccountOptionsEnable() {
	data := []byte(`{
    "email": "123@test.com",
    "isEOptionsEnabled": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var email string = "xxyyzz@gmail.com"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("email", email)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewSubAccountOptionsEnableService().Email(email).Do(newContext())
	s.r().NoError(err)

	e := &SubAccountOptionsEnableServiceResponse{
		Email:             "123@test.com",
		IsEOptionsEnabled: true}
	s.assertSubAccountOptionsEnableServiceResponseEqual(e, res)
}

func (s *subAccountOptionsEnableServiceTestSuite) assertSubAccountOptionsEnableServiceResponseEqual(e, a *SubAccountOptionsEnableServiceResponse) {
	r := s.r()
	r.Equal(e.Email, a.Email, "Email")
	r.Equal(e.IsEOptionsEnabled, a.IsEOptionsEnabled, "IsEOptionsEnabled")
}

type managedSubAccountQueryTransferLogServiceTestSuite struct {
	baseTestSuite
}

func TestManagedSubAccountQueryTransferLogService(t *testing.T) {
	suite.Run(t, new(managedSubAccountQueryTransferLogServiceTestSuite))
}

func (s *managedSubAccountQueryTransferLogServiceTestSuite) TestManagedSubAccountQueryTransferLog() {
	data := []byte(`{
    "managerSubTransferHistoryVos": [
        {
            "fromEmail": "test_0_virtual@kq3kno9imanagedsub.com",
            "fromAccountType": "SPOT",
            "toEmail": "wdywl0lddakh@test.com",
            "toAccountType": "SPOT",
            "asset": "BNB",
            "amount": "0.01",
            "scheduledData": 1679416673000,
            "createTime": 1679416673000,
            "status": "SUCCESS",
            "tranId": 91077779
        },
        {
            "fromEmail": "wdywl0lddakh@test.com",
            "fromAccountType": "SPOT",
            "toEmail": "test_0_virtual@kq3kno9imanagedsub.com",
            "toAccountType": "SPOT",
            "asset": "BNB",
            "amount": "1",
            "scheduledData": 1679416616000,
            "createTime": 1679416616000,
            "status": "SUCCESS",
            "tranId": 91077676
        }
    ],
    "count": 2
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var startTime int64 = 1122334455
	var endTime int64 = 223344558899
	var page int32 = 1
	var limit int32 = 20
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("startTime", startTime).setParam("endTime", endTime).setParam("page", page).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewManagedSubAccountQueryTransferLogService().StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := &ManagedSubAccountQueryTransferLogServiceResponse{
		ManagerSubTransferHistoryVos: []*ManagedSubTransferHistoryVo{{
			FromEmail:       "test_0_virtual@kq3kno9imanagedsub.com",
			FromAccountType: "SPOT",
			ToEmail:         "wdywl0lddakh@test.com",
			ToAccountType:   "SPOT",
			Asset:           "BNB",
			Amount:          "0.01",
			ScheduledData:   1679416673000,
			CreateTime:      1679416673000,
			Status:          "SUCCESS",
			TranId:          91077779},
			{
				FromEmail:       "wdywl0lddakh@test.com",
				FromAccountType: "SPOT",
				ToEmail:         "test_0_virtual@kq3kno9imanagedsub.com",
				ToAccountType:   "SPOT",
				Asset:           "BNB",
				Amount:          "1",
				ScheduledData:   1679416616000,
				CreateTime:      1679416616000,
				Status:          "SUCCESS",
				TranId:          91077676}},
		Count: 2}
	s.assertManagedSubAccountQueryTransferLogServiceResponseEqual(e, res)
}

func (s *managedSubAccountQueryTransferLogServiceTestSuite) assertManagedSubAccountQueryTransferLogServiceResponseEqual(e, a *ManagedSubAccountQueryTransferLogServiceResponse) {
	r := s.r()
	r.Len(e.ManagerSubTransferHistoryVos, len(a.ManagerSubTransferHistoryVos))
	for i := range e.ManagerSubTransferHistoryVos {
		r.Equal(e.ManagerSubTransferHistoryVos[i].FromEmail, a.ManagerSubTransferHistoryVos[i].FromEmail, "ManagerSubTransferHistoryVos[i].FromEmail")
		r.Equal(e.ManagerSubTransferHistoryVos[i].FromAccountType, a.ManagerSubTransferHistoryVos[i].FromAccountType, "ManagerSubTransferHistoryVos[i].FromAccountType")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ToEmail, a.ManagerSubTransferHistoryVos[i].ToEmail, "ManagerSubTransferHistoryVos[i].ToEmail")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ToAccountType, a.ManagerSubTransferHistoryVos[i].ToAccountType, "ManagerSubTransferHistoryVos[i].ToAccountType")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Asset, a.ManagerSubTransferHistoryVos[i].Asset, "ManagerSubTransferHistoryVos[i].Asset")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Amount, a.ManagerSubTransferHistoryVos[i].Amount, "ManagerSubTransferHistoryVos[i].Amount")
		r.Equal(e.ManagerSubTransferHistoryVos[i].ScheduledData, a.ManagerSubTransferHistoryVos[i].ScheduledData, "ManagerSubTransferHistoryVos[i].ScheduledData")
		r.Equal(e.ManagerSubTransferHistoryVos[i].CreateTime, a.ManagerSubTransferHistoryVos[i].CreateTime, "ManagerSubTransferHistoryVos[i].CreateTime")
		r.Equal(e.ManagerSubTransferHistoryVos[i].Status, a.ManagerSubTransferHistoryVos[i].Status, "ManagerSubTransferHistoryVos[i].Status")
		r.Equal(e.ManagerSubTransferHistoryVos[i].TranId, a.ManagerSubTransferHistoryVos[i].TranId, "ManagerSubTransferHistoryVos[i].TranId")
	}

	r.Equal(e.Count, a.Count, "Count")
}
