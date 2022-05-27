package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type stakingServiceTestSuite struct {
	baseTestSuite
}

func TestStakingService(t *testing.T) {
	suite.Run(t, new(stakingServiceTestSuite))
}

func (s *stakingServiceTestSuite) TestStakingProductPosition() {
	data := []byte(`[
	  {
		"positionId": 123123,
		"productId": "Axs*90",
		"asset": "AXS",
		"amount": "122.09202928",
		"purchaseTime": 1646182276000,
		"duration": 60,
		"accrualDays": 4,
		"rewardAsset": "AXS",
		"APY": "0.2032",
		"rewardAmt": "5.17181528",
		"extraRewardAsset": "BNB",
		"extraRewardAPY": "0.0203",
		"estExtraRewardAmt": "5.17181528",
		"nextInterestPay": "1.29295383",
		"nextInterestPayDate": "1646697600000",
		"payInterestPeriod": 1,
		"redeemAmountEarly": "2802.24068892",
		"interestEndDate": 1651449600000,
		"deliverDate": 1651536000000,
		"redeemPeriod": 1,
		"redeemingAmt": "232.2323",
		"partialAmtDeliverDate": 1651536000000,
		"canRedeemEarly": true,
		"renewable": true,
		"type": "AUTO",
		"status": "HOLDING"
	  }
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"product": StakingProductLockedStaking,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewStakingProductPositionService().
		Product(StakingProductLockedStaking).
		Do(newContext())
	s.r().NoError(err)
	e := &StakingProductPositions{
		{
			PositionId:                 123123,
			ProductId:                  "Axs*90",
			Asset:                      "AXS",
			Amount:                     "122.09202928",
			PurchaseTime:               1646182276000,
			Duration:                   60,
			AccrualDays:                4,
			RewardAsset:                "AXS",
			APY:                        "0.2032",
			RewardAmount:               "5.17181528",
			ExtraRewardAsset:           "BNB",
			ExtraRewardAPY:             "0.0203",
			EstimatedExtraRewardAmount: "5.17181528",
			NextInterestPay:            "1.29295383",
			NextInterestPayDate:        "1646697600000",
			PayInterestPeriod:          1,
			RedeemAmountEarly:          "2802.24068892",
			InterestEndDate:            1651449600000,
			DeliverDate:                1651536000000,
			RedeemPeriod:               1,
			RedeemingAmount:            "232.2323",
			PartialAmountDeliverDate:   1651536000000,
			CanRedeemEarly:             true,
			Renewable:                  true,
			Type:                       "AUTO",
			Status:                     "HOLDING",
		},
	}
	s.assertStakingProductPositionsEqual(e, res)
}

func (s *stakingServiceTestSuite) assertStakingProductPositionsEqual(e, a *StakingProductPositions) {
	r := s.r()
	r.Len(*a, len(*e))
	for i := 0; i < len(*a); i++ {
		s.assertStakingProductPositionEqual(&(*e)[i], &(*a)[i])
	}
}

func (s *stakingServiceTestSuite) assertStakingProductPositionEqual(e, a *StakingProductPosition) {
	r := s.r()
	r.Equal(e.PositionId, a.PositionId, "PositionId")
	r.Equal(e.ProductId, a.ProductId, "ProductId")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.PurchaseTime, a.PurchaseTime, "PurchaseTime")
	r.Equal(e.Duration, a.Duration, "Duration")
	r.Equal(e.AccrualDays, a.AccrualDays, "AccrualDays")
	r.Equal(e.RewardAsset, a.RewardAsset, "RewardAsset")
	r.Equal(e.APY, a.APY, "APY")
	r.Equal(e.RewardAmount, a.RewardAmount, "RewardAmount")
	r.Equal(e.ExtraRewardAsset, a.ExtraRewardAsset, "ExtraRewardAsset")
	r.Equal(e.ExtraRewardAPY, a.ExtraRewardAPY, "ExtraRewardAPY")
	r.Equal(e.EstimatedExtraRewardAmount, a.EstimatedExtraRewardAmount, "EstimatedExtraRewardAmount")
	r.Equal(e.NextInterestPay, a.NextInterestPay, "NextInterestPay")
	r.Equal(e.NextInterestPayDate, a.NextInterestPayDate, "NextInterestPayDate")
	r.Equal(e.PayInterestPeriod, a.PayInterestPeriod, "PayInterestPeriod")
	r.Equal(e.RedeemAmountEarly, a.RedeemAmountEarly, "RedeemAmountEarly")
	r.Equal(e.InterestEndDate, a.InterestEndDate, "InterestEndDate")
	r.Equal(e.DeliverDate, a.DeliverDate, "DeliverDate")
	r.Equal(e.RedeemPeriod, a.RedeemPeriod, "RedeemPeriod")
	r.Equal(e.RedeemingAmount, a.RedeemingAmount, "RedeemingAmount")
	r.Equal(e.PartialAmountDeliverDate, a.PartialAmountDeliverDate, "PartialAmountDeliverDate")
	r.Equal(e.CanRedeemEarly, a.CanRedeemEarly, "CanRedeemEarly")
	r.Equal(e.Renewable, a.Renewable, "Renewable")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Status, a.Status, "Status")
}
