package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetDetailServiceTestSuite struct {
	baseTestSuite
}

func TestAssetDetailService(t *testing.T) {
	suite.Run(t, new(assetDetailServiceTestSuite))
}

func (s *withdrawServiceTestSuite) TestGetAssetDetail() {
	data := []byte(`
	{
		"CTR": {
			"minWithdrawAmount": "70.00000000",
			"depositStatus": false,
			"withdrawFee": "35",
			"withdrawStatus": true,
			"depositTip": "Delisted, Deposit Suspended"
		},
		"SKY": {
			"minWithdrawAmount": "0.02000000",
			"depositStatus": true,
			"withdrawFee": "0.01",
			"withdrawStatus": true
		}
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAssetDetailService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal(res["CTR"].DepositStatus, false, "depositStatus")
	s.r().Equal(res["SKY"].DepositStatus, true, "depositStatus")
}
func (s *assetDetailServiceTestSuite) TestGetAllCoinsInfo() {
	data := []byte(`
	[
		{
			"coin": "BTC",
			"depositAllEnable": true,
			"free": "0.08074558",
			"freeze": "0.00000000",
			"ipoable": "0.00000000",
			"ipoing": "0.00000000",
			"isLegalMoney": false,
			"locked": "0.00000000",
			"name": "Bitcoin",
			"networkList": [
				{
					"addressRegex": "^(bnb1)[0-9a-z]{38}$",
					"coin": "BTC",
					"depositDesc": "Wallet Maintenance, Deposit Suspended", 
					"depositEnable": false,
					"isDefault": false,        
					"memoRegex": "^[0-9A-Za-z\\-_]{1,120}$",
					"minConfirm": 1,  
					"name": "BEP2",
					"network": "BNB",            
					"resetAddressStatus": false,
					"specialTips": "Both a MEMO and an Address are required to successfully deposit your BEP2-BTCB tokens to Binance.",
					"unLockConfirm": 0,   
					"withdrawDesc": "Wallet Maintenance, Withdrawal Suspended", 
					"withdrawEnable": false,
					"withdrawFee": "0.00000220",
					"withdrawIntegerMultiple": "0.00000001",
					"withdrawMax": "9999999999.99999999",
					"withdrawMin": "0.00000440",
					"sameAddress": true,
					"estimatedArrivalTime": 25,
					"busy": false,
					"contractAddressUrl": "https://bscscan.com/token/",
					"contractAddress": "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c"
				},
				{
					"addressRegex": "^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$|^(bc1)[0-9A-Za-z]{39,59}$",
					"coin": "BTC",
					"depositEnable": true,
					"isDefault": true,
					"memoRegex": "",
					"minConfirm": 1,  
					"name": "BTC",
					"network": "BTC",
					"resetAddressStatus": false,
					"specialTips": "",
					"unLockConfirm": 0,  
					"withdrawEnable": true,
					"withdrawFee": "0.00050000",
					"withdrawIntegerMultiple": "0.00000001",
					"withdrawMax": "750",
					"withdrawMin": "0.00100000",
					"sameAddress": false,
					"estimatedArrivalTime": 25,
					"busy": false,
					"contractAddressUrl": "",
					"contractAddress": ""
				}
			],
			"storage": "0.00000000",
			"trading": true,
			"withdrawAllEnable": true,
			"withdrawing": "0.00000000"
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAllCoinsInfoService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal(res[0].DepositAllEnable, true, "depositAllEnable")
	s.r().Equal(res[0].NetworkList[0].WithdrawEnable, false, "withdrawEnable")
	s.r().Equal(res[0].NetworkList[0].EstimatedArrivalTime, 25, "estimatedArrivalTime")
	s.r().Equal(res[0].NetworkList[0].Busy, false, "busy")
	s.r().Equal(res[0].NetworkList[0].ContractAddressUrl, "https://bscscan.com/token/", "contractAddressUrl")
	s.r().Equal(res[0].NetworkList[0].ContractAddress, "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c", "contractAddress")
	s.r().Equal(res[0].NetworkList[1].MinConfirm, 1, "minConfirm")
}

func (s *assetDetailServiceTestSuite) TestGetUserAsset() {
	data := []byte(`
	[
		{
			"asset": "AVAX",
			"free": "1",
			"locked": "0",
			"freeze": "0",
			"withdrawing": "0",
			"ipoable": "0",
			"btcValuation": "0"
		},
		{
			"asset": "BCH",
			"free": "0.9",
			"locked": "0",
			"freeze": "0",
			"withdrawing": "0",
			"ipoable": "0",
			"btcValuation": "0"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetUserAsset().Do(newContext())
	s.r().NoError(err)
	s.assertUserAssetEqual(UserAssetRecord{
		Asset:        "AVAX",
		Free:         "1",
		Locked:       "0",
		Freeze:       "0",
		Withdrawing:  "0",
		Ipoable:      "0",
		BtcValuation: "0",
	}, res[0])
	s.assertUserAssetEqual(UserAssetRecord{
		Asset:        "BCH",
		Free:         "0.9",
		Locked:       "0",
		Freeze:       "0",
		Withdrawing:  "0",
		Ipoable:      "0",
		BtcValuation: "0",
	}, res[1])
}

func (s *assetDetailServiceTestSuite) assertUserAssetEqual(e, a UserAssetRecord) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Free, a.Free, "Free")
	r.Equal(e.Locked, a.Locked, "Locked")
	r.Equal(e.Freeze, a.Freeze, "Freeze")
	r.Equal(e.Withdrawing, a.Withdrawing, "Withdrawing")
	r.Equal(e.Ipoable, a.Ipoable, "Ipoable")
	r.Equal(e.BtcValuation, a.BtcValuation, "BtcValuation")
}

func (s *assetDetailServiceTestSuite) TestGetFundingAsset() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"free": "1",
			"locked": "0",
			"freeze": "0",
			"withdrawing": "0",
			"btcValuation": "0"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetFundingAssetService().Do(newContext())
	s.r().NoError(err)
	s.assertFundingAssetEqual(FundingAsset{
		Asset: "BTC",
		Free:  "1",
		Locked: "0",
		Freeze: "0",
		Withdrawing: "0",
		BtcValuation: "0",
	}, res[0])
}

func (s *assetDetailServiceTestSuite) assertFundingAssetEqual(e, a FundingAsset) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Free, a.Free, "Free")
	r.Equal(e.Locked, a.Locked, "Locked")
	r.Equal(e.Freeze, a.Freeze, "Freeze")
	r.Equal(e.Withdrawing, a.Withdrawing, "Withdrawing")
	r.Equal(e.BtcValuation, a.BtcValuation, "BtcValuation")
}