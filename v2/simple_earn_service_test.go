package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type simpleEarnServiceTestSuite struct {
	baseTestSuite
}

func TestSimpleEarnService(t *testing.T) {
	suite.Run(t, new(simpleEarnServiceTestSuite))
}

func (s *simpleEarnServiceTestSuite) TestGetAccount() {
	data := []byte(`{
  "totalAmountInBTC": "0.01067982",
  "totalAmountInUSDT": "77.13289230",
  "totalFlexibleAmountInBTC": "0.00000000",
  "totalFlexibleAmountInUSDT": "0.00000000",
  "totalLockedInBTC": "0.01067982",
  "totalLockedInUSDT": "77.13289230"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	account, err := s.client.NewSimpleEarnService().GetAccount().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("0.01067982", account.TotalAmountInBTC)
	s.r().Equal("77.13289230", account.TotalAmountInUSDT)
	s.r().Equal("0.00000000", account.TotalFlexibleAmountInBTC)
	s.r().Equal("0.00000000", account.TotalFlexibleAmountInUSDT)
	s.r().Equal("0.01067982", account.TotalLockedInBTC)
	s.r().Equal("77.13289230", account.TotalLockedInUSDT)
}

func (s *simpleEarnServiceTestSuite) TestListFlexibleProduct() {
	data := []byte(`{
  "rows": [
    {
      "asset": "BTC",
      "latestAnnualPercentageRate": "0.05000000",  
      "tierAnnualPercentageRate": {               
        "0-5BTC": "0.05",
        "5-10BTC": "0.03"
      },
      "airDropPercentageRate": "0.05000000",      
      "canPurchase": true,
      "canRedeem": true,
      "isSoldOut": true,
      "hot": true,                                
      "minPurchaseAmount": "0.01000000",
      "productId": "BTC001",
      "subscriptionStartTime": 1646182276000,
      "status": "PURCHASING"                      
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "BTC",
		})
		s.assertRequestEqual(e, r)
	})

	product, err := s.client.NewSimpleEarnService().FlexibleService().ListProduct().Asset("BTC").Do(newContext())
	s.r().NoError(err)
	s.r().Equal("BTC001", product.Rows[0].ProductId)
	s.r().Equal("BTC", product.Rows[0].Asset)
	s.r().Equal("0.05000000", product.Rows[0].LatestAnnualPercentageRate)
	s.r().Equal(map[string]string{"0-5BTC": "0.05", "5-10BTC": "0.03"}, product.Rows[0].TierAnnualPercentageRate)
	s.r().Equal("0.05000000", product.Rows[0].AirDropPercentageRate)
	s.r().Equal(true, product.Rows[0].CanPurchase)
	s.r().Equal(true, product.Rows[0].CanRedeem)
	s.r().Equal(true, product.Rows[0].IsSoldOut)
	s.r().Equal(true, product.Rows[0].Hot)
	s.r().Equal("0.01000000", product.Rows[0].MinPurchaseAmount)
	s.r().EqualValues(1646182276000, product.Rows[0].SubscriptionStartTime)
	s.r().Equal("PURCHASING", product.Rows[0].Status)
	s.r().Equal(1, product.Total)
}

func (s *simpleEarnServiceTestSuite) TestListLockedProduct() {
	data := []byte(`{
  "rows": [
    {
      "projectId": "Axs*90",
      "detail": {
        "asset": "AXS",                
        "rewardAsset": "AXS",          
        "duration": 90,                
        "renewable": true,             
        "isSoldOut": true,
        "apr": "1.2069",
        "status": "CREATED",           
        "subscriptionStartTime": 1646182276000,
        "extraRewardAsset": "BNB",
        "extraRewardAPR": "0.23"
      },
      "quota": {
        "totalPersonalQuota": "2",     
        "minimum": "0.001"              
      }
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "AXS",
		})
		s.assertRequestEqual(e, r)
	})

	product, err := s.client.NewSimpleEarnService().LockedService().ListProduct().Asset("AXS").Do(newContext())
	s.r().NoError(err)
	s.r().Equal("Axs*90", product.Rows[0].ProjectId)
	s.r().Equal("AXS", product.Rows[0].Detail.Asset)
	s.r().Equal("AXS", product.Rows[0].Detail.RewardAsset)
	s.r().Equal(90, product.Rows[0].Detail.Duration)
	s.r().Equal(true, product.Rows[0].Detail.Renewable)
	s.r().Equal(true, product.Rows[0].Detail.IsSoldOut)
	s.r().Equal("1.2069", product.Rows[0].Detail.Apr)
	s.r().Equal("CREATED", product.Rows[0].Detail.Status)
	s.r().EqualValues(1646182276000, product.Rows[0].Detail.SubscriptionStartTime)
	s.r().Equal("BNB", product.Rows[0].Detail.ExtraRewardAsset)
	s.r().Equal("0.23", product.Rows[0].Detail.ExtraRewardAPR)
	s.r().Equal("2", product.Rows[0].Quota.TotalPersonalQuota)
	s.r().Equal("0.001", product.Rows[0].Quota.Minimum)
	s.r().Equal(1, product.Total)
}

func (s *simpleEarnServiceTestSuite) TestGetSimpleEarnFlexiblePositionService() {
	data := []byte(`{
  "rows": [
    {
      "totalAmount": "75.46000000",
      "tierAnnualPercentageRate": {
        "0-5BTC": "0.05",
        "5-10BTC": "0.03"
      },
      "latestAnnualPercentageRate": "0.02599895",
      "yesterdayAirdropPercentageRate": "0.02599895",
      "asset": "USDT",
      "airDropAsset": "BETH",
      "canRedeem": true,
      "collateralAmount": "232.23123213",
      "productId": "USDT001",
      "yesterdayRealTimeRewards": "0.10293829",
      "cumulativeBonusRewards": "0.22759183",
      "cumulativeRealTimeRewards": "0.22759183",
      "cumulativeTotalRewards": "0.45459183",
      "autoSubscribe": true
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "USDT",
		})
		s.assertRequestEqual(e, r)
	})

	position, err := s.client.NewSimpleEarnService().FlexibleService().GetPosition().Asset("USDT").Do(newContext())
	s.r().NoError(err)
	s.r().Equal("75.46000000", position.Rows[0].TotalAmount)
	s.r().Equal(map[string]string{"0-5BTC": "0.05", "5-10BTC": "0.03"}, position.Rows[0].TierAnnualPercentageRate)
	s.r().Equal("0.02599895", position.Rows[0].LatestAnnualPercentageRate)
	s.r().Equal("0.02599895", position.Rows[0].YesterdayAirdropPercentageRate)
	s.r().Equal("USDT", position.Rows[0].Asset)
	s.r().Equal("BETH", position.Rows[0].AirDropAsset)
	s.r().Equal(true, position.Rows[0].CanRedeem)
	s.r().Equal("232.23123213", position.Rows[0].CollateralAmount)
	s.r().Equal("USDT001", position.Rows[0].ProductId)
	s.r().Equal("0.10293829", position.Rows[0].YesterdayRealTimeRewards)
	s.r().Equal("0.22759183", position.Rows[0].CumulativeBonusRewards)
	s.r().Equal("0.22759183", position.Rows[0].CumulativeRealTimeRewards)
	s.r().Equal("0.45459183", position.Rows[0].CumulativeTotalRewards)
	s.r().Equal(true, position.Rows[0].AutoSubscribe)
	s.r().Equal(1, position.Total)
}

func (s *simpleEarnServiceTestSuite) TestGetSimpleEarnLockedPositionService() {
	data := []byte(`{
  "rows": [
    {
      "positionId": 123123,
      "parentPositionId": 123122,
      "projectId": "Axs*90",
      "asset": "AXS",
      "amount": "122.09202928",
      "purchaseTime": 1646182276000,
      "duration": 60,
      "accrualDays": 4,
      "rewardAsset": "AXS",
      "APY": "0.2032",
      "rewardAmt": "5.17181528",
      "extraRewardAsset": "BNB",
      "extraRewardAPR": "0.0203",
      "estExtraRewardAmt": "5.17181528",
      "nextPay": "1.29295383",
      "nextPayDate": 1646697600000,
      "payPeriod": 1,
      "redeemAmountEarly": "2802.24068892",
      "rewardsEndDate": 1651449600000,
      "deliverDate": 1651536000000,
      "redeemPeriod": 1,
      "redeemingAmt": "232.2323",
      "redeemTo": "FLEXIBLE",
      "partialAmtDeliverDate": 1651536000000,
      "canRedeemEarly": true,
      "canFastRedemption": true,
      "autoSubscribe": true,
      "type": "AUTO",
      "status": "HOLDING",
      "canReStake": true
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":      "AXS",
			"positionId": 123123,
			"projectId":  "Axs*90",
			"current":    1,
			"size":       10,
		})
		s.assertRequestEqual(e, r)
	})

	position, err := s.client.NewSimpleEarnService().LockedService().GetPosition().
		Asset("AXS").
		PositionId(123123).
		ProjectId("Axs*90").
		Current(1).
		Size(10).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(123123, position.Rows[0].PositionId)
	s.r().Equal(123122, position.Rows[0].ParentPositionId)
	s.r().Equal("Axs*90", position.Rows[0].ProjectId)
	s.r().Equal("AXS", position.Rows[0].Asset)
	s.r().Equal("122.09202928", position.Rows[0].Amount)
	s.r().EqualValues(1646182276000, position.Rows[0].PurchaseTime)
	s.r().EqualValues(60, position.Rows[0].Duration)
	s.r().EqualValues(4, position.Rows[0].AccrualDays)
	s.r().Equal("AXS", position.Rows[0].RewardAsset)
	s.r().Equal("0.2032", position.Rows[0].APY)
	s.r().Equal("5.17181528", position.Rows[0].RewardAmt)
	s.r().Equal("BNB", position.Rows[0].ExtraRewardAsset)
	s.r().Equal("0.0203", position.Rows[0].ExtraRewardAPR)
	s.r().Equal("5.17181528", position.Rows[0].EstExtraRewardAmt)
	s.r().Equal("1.29295383", position.Rows[0].NextPay)
	s.r().EqualValues(1646697600000, position.Rows[0].NextPayDate)
	s.r().EqualValues(1, position.Rows[0].PayPeriod)
	s.r().Equal("2802.24068892", position.Rows[0].RedeemAmountEarly)
	s.r().EqualValues(1651449600000, position.Rows[0].RewardsEndDate)
	s.r().EqualValues(1651536000000, position.Rows[0].DeliverDate)
	s.r().EqualValues(1, position.Rows[0].RedeemPeriod)
	s.r().Equal("232.2323", position.Rows[0].RedeemingAmt)
	s.r().Equal("FLEXIBLE", position.Rows[0].RedeemTo)
	s.r().EqualValues(1651536000000, position.Rows[0].PartialAmtDeliverDate)
	s.r().Equal(true, position.Rows[0].CanRedeemEarly)
	s.r().Equal(true, position.Rows[0].CanFastRedemption)
	s.r().Equal(true, position.Rows[0].AutoSubscribe)
	s.r().Equal("AUTO", position.Rows[0].Type)
	s.r().Equal("HOLDING", position.Rows[0].Status)
	s.r().Equal(true, position.Rows[0].CanReStake)
	s.r().Equal(1, position.Total)
}

func (s *simpleEarnServiceTestSuite) TestGetSimpleEarnFlexibleQuotaService() {
	data := []byte(`{
  "leftPersonalQuota": "1000"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId": "BTC001",
		})
		s.assertRequestEqual(e, r)
	})

	quota, err := s.client.NewSimpleEarnService().FlexibleService().GetLeftQuote().
		ProductId("BTC001").
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("1000", quota.LeftPersonalQuota)
}

func (s *simpleEarnServiceTestSuite) TestGetSimpleEarnLockedQuotaService() {
	data := []byte(`{
  "leftPersonalQuota": "1000"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"projectId": "AXS001",
		})
		s.assertRequestEqual(e, r)
	})

	quota, err := s.client.NewSimpleEarnService().LockedService().GetLeftQuote().
		ProjectId("AXS001").
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("1000", quota.LeftPersonalQuota)
}

func (s *simpleEarnServiceTestSuite) TestSubscribeFlexibleProduct() {
	data := []byte(`{
  "purchaseId": 40607,
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId":     "BTC001",
			"amount":        "0.1",
			"autoSubscribe": true,
			"sourceAccount": "SPOT",
		})
		s.assertRequestEqual(e, r)
	})

	subscribeResp, err := s.client.NewSimpleEarnService().FlexibleService().Subscribe().
		ProductId("BTC001").
		Amount("0.1").
		AutoSubscribe(true).
		SourceAccount(SourceAccountSpot).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(40607, subscribeResp.PurchaseId)
	s.r().Equal(true, subscribeResp.Success)
}

func (s *simpleEarnServiceTestSuite) TestSubscribeLockedProduct() {
	data := []byte(`{
  "purchaseId": 40607,
  "positionId": 12345,
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"projectId":     "AXS001",
			"amount":        "0.1",
			"autoSubscribe": true,
			"sourceAccount": "SPOT",
			"redeemTo":      "FLEXIBLE",
		})
		s.assertRequestEqual(e, r)
	})

	subscribeResp, err := s.client.NewSimpleEarnService().LockedService().Subscribe().
		ProjectId("AXS001").
		Amount("0.1").
		AutoSubscribe(true).
		SourceAccount(SourceAccountSpot).
		RedeemTo("FLEXIBLE").
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(40607, subscribeResp.PurchaseId)
	s.r().EqualValues(12345, subscribeResp.PositionId)
	s.r().Equal(true, subscribeResp.Success)
}

func (s *simpleEarnServiceTestSuite) TestRedeemFlexibleProduct() {
	data := []byte(`{
  "redeemId": 40607,
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId":   "BTC001",
			"redeemAll":   true,
			"amount":      "0.1",
			"destAccount": "SPOT",
		})
		s.assertRequestEqual(e, r)
	})

	redeemResp, err := s.client.NewSimpleEarnService().FlexibleService().Redeem().
		ProductId("BTC001").
		RedeemAll(true).
		Amount("0.1").
		DestAccount("SPOT").
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(40607, redeemResp.RedeemId)
	s.r().Equal(true, redeemResp.Success)
}

func (s *simpleEarnServiceTestSuite) TestRedeemLockedProduct() {
	data := []byte(`{
  "redeemId": 40607,
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"positionId": 12345,
		})
		s.assertRequestEqual(e, r)
	})

	redeemResp, err := s.client.NewSimpleEarnService().LockedService().Redeem().
		PositionId(12345).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(40607, redeemResp.RedeemId)
	s.r().Equal(true, redeemResp.Success)
}

func (s *simpleEarnServiceTestSuite) TestSetAutoSubscribeFlexibleProduct() {
	data := []byte(`{
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId":     "BTC001",
			"autoSubscribe": true,
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewSimpleEarnService().FlexibleService().SetAutoSubscribe().
		ProductId("BTC001").
		AutoSubscribe(true).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(true, resp.Success)
}

func (s *simpleEarnServiceTestSuite) TestSetAutoSubscribeLockedProduct() {
	data := []byte(`{
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"positionId":    12345,
			"autoSubscribe": true,
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewSimpleEarnService().LockedService().SetAutoSubscribe().
		PositionId(12345).
		AutoSubscribe(true).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(true, resp.Success)
}

func (s *simpleEarnServiceTestSuite) TestFlexibleSubscriptionPreview() {
	data := []byte(`{
  "totalAmount": "1232.32230982",
  "rewardAsset": "BUSD",
  "airDropAsset": "BETH",
  "estDailyBonusRewards": "0.22759183",
  "estDailyRealTimeRewards": "0.22759183",
  "estDailyAirdropRewards": "0.22759183"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"productId": "BTC001",
			"amount":    "0.1",
		})
		s.assertRequestEqual(e, r)
	})

	preview, err := s.client.NewSimpleEarnService().FlexibleService().PreviewSubscribe().
		ProductId("BTC001").
		Amount("0.1").
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("1232.32230982", preview.TotalAmount)
	s.r().Equal("BUSD", preview.RewardAsset)
	s.r().Equal("BETH", preview.AirDropAsset)
	s.r().Equal("0.22759183", preview.EstDailyBonusRewards)
	s.r().Equal("0.22759183", preview.EstDailyRealTimeRewards)
	s.r().Equal("0.22759183", preview.EstDailyAirdropRewards)
}

func (s *simpleEarnServiceTestSuite) TestLockedSubscriptionPreview() {
	data := []byte(`
  {
    "rewardAsset": "AXS",
    "totalRewardAmt": "5.17181528",
    "extraRewardAsset": "BNB",
    "estTotalExtraRewardAmt": "5.17181528",
    "nextPay": "1.29295383",
    "nextPayDate": 1646697600000,
    "valueDate": 1646697600000,
    "rewardsEndDate": 1651449600000,
    "deliverDate": 1651536000000,
    "nextSubscriptionDate": 1651536000000
  }
`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"projectId":     "AXS001",
			"amount":        "0.1",
			"autoSubscribe": true,
		})
		s.assertRequestEqual(e, r)
	})

	preview, err := s.client.NewSimpleEarnService().LockedService().PreviewSubscribe().
		ProjectId("AXS001").
		Amount("0.1").
		AutoSubscribe(true).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("AXS", preview.RewardAsset)
	s.r().Equal("5.17181528", preview.TotalRewardAmt)
	s.r().Equal("BNB", preview.ExtraRewardAsset)
	s.r().Equal("5.17181528", preview.EstTotalExtraRewardAmt)
	s.r().Equal("1.29295383", preview.NextPay)
	s.r().EqualValues(1646697600000, preview.NextPayDate)
	s.r().EqualValues(1646697600000, preview.ValueDate)
	s.r().EqualValues(1651449600000, preview.RewardsEndDate)
	s.r().EqualValues(1651536000000, preview.DeliverDate)
	s.r().EqualValues(1651536000000, preview.NextSubscriptionDate)
}

func (s *simpleEarnServiceTestSuite) TestLockedSetRedeemOption() {
	data := []byte(`{
  "success": true
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"positionId": "12345",
			"redeemTo":   "SPOT",
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewSimpleEarnService().LockedService().SetRedeemOption().
		PositionId("12345").
		RedeemTo(RedeemToSpot).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(true, resp.Success)
}
