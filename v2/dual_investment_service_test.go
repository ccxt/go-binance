package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type dualInvestmentServiceTestSuite struct {
	baseTestSuite
}

func TestDualInvestmentService(t *testing.T) {
	suite.Run(t, new(dualInvestmentServiceTestSuite))
}

func (s *dualInvestmentServiceTestSuite) TestListProducts() {
	data := []byte(`{
    "total": 1,
    "list": [
        {
            "id": "741590",
            "investCoin": "USDT",
            "exercisedCoin": "BNB",
            "strikePrice": "380",
            "duration": 4,
            "settleDate": 1709020800000,
            "purchaseDecimal": 8,
            "purchaseEndTime": 1708934400000,
            "canPurchase": true, 
            "apr": "0.6076",
            "orderId": 8257205859,
            "minAmount": "0.1",
            "maxAmount": "25265.7",
            "createTimestamp": 1708560084000,
            "optionType": "PUT",
            "isAutoCompoundEnable": true, 
            "autoCompoundPlanList": [
                "STANDARD",
                "ADVANCE"
            ]
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"optionType":    "PUT",
			"exercisedCoin": "BNB",
			"investCoin":    "USDT",
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewDualInvestmentService().
		ListProductService().
		InvestCoin("USDT").
		ExercisedCoin("BNB").
		OptionType(DualInvestmentOptionTypePut).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Equal(1, len(resp.List))
	s.assertProduct(&DualInvestmentProduct{
		ID:                   "741590",
		InvestCoin:           "USDT",
		ExercisedCoin:        "BNB",
		StrikePrice:          "380",
		Duration:             4,
		SettleDate:           1709020800000,
		PurchaseDecimal:      8,
		PurchaseEndTime:      1708934400000,
		CanPurchase:          true,
		APR:                  "0.6076",
		OrderID:              8257205859,
		MinAmount:            "0.1",
		MaxAmount:            "25265.7",
		CreateTimestamp:      1708560084000,
		OptionType:           DualInvestmentOptionTypePut,
		IsAutoCompoundEnable: true,
		AutoCompoundPlanList: []string{"STANDARD", "ADVANCE"},
	}, &resp.List[0])
}

func (s *dualInvestmentServiceTestSuite) TestSubscribe() {
	data := []byte(`{
		"positionId": 10208824,
		"investCoin": "BNB",
		"exercisedCoin": "USDT",
		"subscriptionAmount": "0.002",
		"duration": 4,
		"autoCompoundPlan": "STANDARD",
		"strikePrice": "380",
		"settleDate": 1709020800000,
		"purchaseStatus": "PURCHASE_SUCCESS",
		"apr": "0.7397",
		"orderId": 8259117597,
		"purchaseTime": 1708677583874,
		"optionType": "CALL"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"id":               "10208824",
			"orderId":          8259117597,
			"depositAmount":    "0.002",
			"autoCompoundPlan": "STANDARD",
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewDualInvestmentService().
		SubscribeService().
		ID("10208824").
		OrderID(8259117597).
		DepositAmount("0.002").
		AutoCompoundPlan(DualInvestmentCompoundPlanStandard).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Equal(int64(10208824), resp.PositionID)
	r.Equal("BNB", resp.InvestCoin)
	r.Equal("USDT", resp.ExercisedCoin)
	r.Equal("0.002", resp.SubscriptionAmount)
	r.Equal(4, resp.Duration)
	r.Equal(DualInvestmentCompoundPlanStandard, resp.AutoCompoundPlan)
	r.Equal("380", resp.StrikePrice)
	r.Equal(int64(1709020800000), resp.SettleDate)
	r.Equal(ListDualInvestmentPositionStatusPurchaseSuccess, resp.PurchaseStatus)
	r.Equal("0.7397", resp.APR)
	r.Equal(int64(8259117597), resp.OrderID)
	r.Equal(int64(1708677583874), resp.PurchaseTime)
	r.Equal(DualInvestmentOptionTypeCall, resp.OptionType)
}

func (s *dualInvestmentServiceTestSuite) TestListPositions() {
	data := []byte(`{
    "total": 1,
    "list": [
        {
            "id": "10160533",
            "investCoin": "USDT",
            "exercisedCoin": "BNB",
            "subscriptionAmount": "0.5",
            "strikePrice": "330",
            "duration": 4,
            "settleDate": 1708416000000,
            "purchaseStatus": "PURCHASE_SUCCESS",
            "apr": "0.0365",
            "orderId": 7973677530,
            "purchaseEndTime": 1708329600000,
            "optionType": "PUT",
            "autoCompoundPlan": "STANDARD"
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"status": "PURCHASE_SUCCESS",
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewDualInvestmentService().
		ListPositionService().
		Status(ListDualInvestmentPositionStatusPurchaseSuccess).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Equal(1, len(resp.List))
	r.Equal("10160533", resp.List[0].ID)
	r.Equal("USDT", resp.List[0].InvestCoin)
	r.Equal("BNB", resp.List[0].ExercisedCoin)
	r.Equal("0.5", resp.List[0].SubscriptionAmount)
	r.Equal("330", resp.List[0].StrikePrice)
	r.Equal(4, resp.List[0].Duration)
	r.Equal(int64(1708416000000), resp.List[0].SettleDate)
	r.Equal(ListDualInvestmentPositionStatusPurchaseSuccess, resp.List[0].PurchaseStatus)
	r.Equal("0.0365", resp.List[0].APR)
	r.Equal(int64(7973677530), resp.List[0].OrderID)
	r.Equal(int64(1708329600000), resp.List[0].PurchaseEndTime)
	r.Equal(DualInvestmentOptionTypePut, resp.List[0].OptionType)
	r.Equal(DualInvestmentCompoundPlanStandard, resp.List[0].AutoCompoundPlan)
}

func (s *dualInvestmentServiceTestSuite) TestListPositionsSettled() {
	data := []byte(`{
    "total": 1,
    "list": [
        {
            "id": "15653189",
            "investCoin": "BTC",
            "exercisedCoin": "FDUSD",
            "subscriptionAmount": "0.00498",
            "strikePrice": "120500",
            "duration": 7,
            "settleDate": 1758873600000,
            "purchaseStatus": "SETTLED",
            "apr": "0.1123",
            "orderId": 37623188787,
            "purchaseEndTime": 1758816000000,
            "optionType": "CALL",
            "autoCompoundPlan": "NONE",
            "settlePrice": "109308.64",
            "isExercised": false,
            "settleAsset": "BTC",
            "settleAmount": "0.00499045"
        }
    ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"status": "SETTLED",
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewDualInvestmentService().
		ListPositionService().
		Status(ListDualInvestmentPositionStatusSettled).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Equal(1, len(resp.List))

	pos := resp.List[0]
	r.Equal("15653189", pos.ID)
	r.Equal("BTC", pos.InvestCoin)
	r.Equal("FDUSD", pos.ExercisedCoin)
	r.Equal("0.00498", pos.SubscriptionAmount)
	r.Equal("120500", pos.StrikePrice)
	r.Equal(7, pos.Duration)
	r.Equal(int64(1758873600000), pos.SettleDate)
	r.Equal(ListDualInvestmentPositionStatusSettled, pos.PurchaseStatus)
	r.Equal("0.1123", pos.APR)
	r.Equal(int64(37623188787), pos.OrderID)
	r.Equal(int64(1758816000000), pos.PurchaseEndTime)
	r.Equal(DualInvestmentOptionTypeCall, pos.OptionType)
	r.Equal(DualInvestmentCompoundPlanNone, pos.AutoCompoundPlan)

	// Assert settlement fields
	r.NotNil(pos.SettlePrice)
	r.Equal("109308.64", *pos.SettlePrice)
	r.NotNil(pos.SettleAmount)
	r.Equal("0.00499045", *pos.SettleAmount)
	r.NotNil(pos.SettleAsset)
	r.Equal("BTC", *pos.SettleAsset)
	r.NotNil(pos.IsExercised)
	r.Equal(false, *pos.IsExercised)
}

func (s *dualInvestmentServiceTestSuite) TestGetAccounts() {
	data := []byte(`{
   "totalAmountInBTC": "0.01067982",   
   "totalAmountInUSDT": "77.13289230" 
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewDualInvestmentService().
		GetAccountService().
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Equal("0.01067982", resp.TotalAmountInBTC)
	r.Equal("77.13289230", resp.TotalAmountInUSDT)
}

func (s *dualInvestmentServiceTestSuite) TestEditAutoCompoundStatus() {
	data := []byte(`{
    "positionId": 123456789,
    "autoCompoundPlan":"ADVANCED"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"positionId":       "123456789",
			"AutoCompoundPlan": "ADVANCED",
		})
		s.assertRequestEqual(e, r)
	})

	resp, err := s.client.NewDualInvestmentService().
		EditAutoCompoundStatusService().
		PositionID("123456789").
		AutoCompoundPlan(DualInvestmentCompoundPlanAdvanced).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Equal(int64(123456789), resp.PositionID)
	r.Equal(DualInvestmentCompoundPlanAdvanced, resp.AutoCompoundPlan)
}

func (s *dualInvestmentServiceTestSuite) assertProduct(e, a *DualInvestmentProduct) {
	r := s.r()
	r.Equal(e.ID, a.ID)
	r.Equal(e.InvestCoin, a.InvestCoin)
	r.Equal(e.ExercisedCoin, a.ExercisedCoin)
	r.Equal(e.StrikePrice, a.StrikePrice)
	r.Equal(e.Duration, a.Duration)
	r.Equal(e.SettleDate, a.SettleDate)
	r.Equal(e.PurchaseDecimal, a.PurchaseDecimal)
	r.Equal(e.PurchaseEndTime, a.PurchaseEndTime)
	r.Equal(e.CanPurchase, a.CanPurchase)
	r.Equal(e.APR, a.APR)
	r.Equal(e.OrderID, a.OrderID)
	r.Equal(e.MinAmount, a.MinAmount)
	r.Equal(e.MaxAmount, a.MaxAmount)
	r.Equal(e.CreateTimestamp, a.CreateTimestamp)
	r.Equal(e.OptionType, a.OptionType)
	r.Equal(e.IsAutoCompoundEnable, a.IsAutoCompoundEnable)
	r.Equal(e.AutoCompoundPlanList, a.AutoCompoundPlanList)
}
