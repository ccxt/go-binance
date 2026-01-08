package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type vipLoanServiceTestSuite struct {
	baseTestSuite
}

func TestVipLoanService(t *testing.T) {
	suite.Run(t, new(vipLoanServiceTestSuite))
}

func (s *vipLoanServiceTestSuite) TestVipLoanInterestRateService() {
	data := []byte(`[
    {
     "asset": "BUSD",
     "flexibleDailyInterestRate": "0.001503",
     "flexibleYearlyInterestRate": "0.548595",
     "time": "1577233578000"
    },
    {
     "asset": "BTC",
     "flexibleDailyInterestRate": "0.001503",
     "flexibleYearlyInterestRate": "0.548595",
     "time": "1577233562000"
    }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("loanCoin", "BTC,BUSD")
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewVipLoanService().InterestRate().LoanCoin("BTC,BUSD").Do(newContext())
	s.r().NoError(err)
	e := &VipLoanInterestRate{
		{
			Asset:                      "BUSD",
			FlexibleDailyInterestRate:  "0.001503",
			FlexibleYearlyInterestRate: "0.548595",
			Time:                       "1577233578000",
		},
		{
			Asset:                      "BTC",
			FlexibleDailyInterestRate:  "0.001503",
			FlexibleYearlyInterestRate: "0.548595",
			Time:                       "1577233562000",
		},
	}
	s.assertVipLoanInterestRateEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanInterestRateEqual(e, a VipLoanInterestRate) {
	r := s.r()
	for i := 0; i < len(e); i++ {
		r.Equal(e[i].Asset, a[i].Asset, "Asset")
		r.Equal(e[i].FlexibleDailyInterestRate, a[i].FlexibleDailyInterestRate, "FlexibleDailyInterestRate")
		r.Equal(e[i].FlexibleYearlyInterestRate, a[i].FlexibleYearlyInterestRate, "FlexibleYearlyInterestRate")
		r.Equal(e[i].Time, a[i].Time, "Time")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanInterestRateHistory() {
	data := []byte(`{
  "rows": [
    {
      "coin": "USDT",
      "annualizedInterestRate": "0.0647",
      "time": "1575018510000"
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().InterestRateHistory().Coin("USDT").Do(newContext())
	s.r().NoError(err)
	e := &VipLoanInterestRateHistoryResponse{
		Rows: []VipLoanInterestRateHistory{
			{
				Coin:                   "USDT",
				AnnualizedInterestRate: "0.0647",
				Time:                   "1575018510000",
			},
		},
		Total: 1,
	}
	s.assertVipLoanInterestRateHistoryResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanInterestRateHistoryResponseEqual(e, a VipLoanInterestRateHistoryResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].Coin, a.Rows[i].Coin, "Coin")
		r.Equal(e.Rows[i].AnnualizedInterestRate, a.Rows[i].AnnualizedInterestRate, "AnnualizedInterestRate")
		r.Equal(e.Rows[i].Time, a.Rows[i].Time, "Time")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanLoanableAssetData() {
	data := []byte(`{
  "rows": [
    {
    "loanCoin": "BUSD",
    "_flexibleDailyInterestRate": "0.001503",
    "_flexibleYearlyInterestRate": "0.548595",
    "_30dDailyInterestRate": "0.000136",
    "_30dYearlyInterestRate": "0.03450",
    "_60dDailyInterestRate": "0.000145",
    "_60dYearlyInterestRate": "0.04103",
    "minLimit": "100",
    "maxLimit": "1000000",
    "vipLevel": 1
    }
  ],
  "total": 1
}
`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().LoanableAssetData().LoanCoin("BUSD").VipLevel(1).Do(newContext())
	s.r().NoError(err)
	e := &VipLoanLoanableAssetDataResponse{
		Rows: []VipLoanLoanableAssetData{
			{
				LoanCoin:                    "BUSD",
				FlexibleDailyInterestRate:   "0.001503",
				FlexibleYearlyInterestRate:  "0.548595",
				ThirtyDayDailyInterestRate:  "0.000136",
				ThirtyDayYearlyInterestRate: "0.03450",
				SixtyDayDailyInterestRate:   "0.000145",
				SixtyDayYearlyInterestRate:  "0.04103",
			},
		},
		Total: 1,
	}
	s.assertVipLoanLoanableAssetDataResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanLoanableAssetDataResponseEqual(e, a VipLoanLoanableAssetDataResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].LoanCoin, a.Rows[i].LoanCoin, "LoanCoin")
		r.Equal(e.Rows[i].FlexibleYearlyInterestRate, a.Rows[i].FlexibleYearlyInterestRate, "FlexibleYearlyInterestRate")
		r.Equal(e.Rows[i].FlexibleDailyInterestRate, a.Rows[i].FlexibleDailyInterestRate, "FlexibleDailyInterestRate")
		r.Equal(e.Rows[i].ThirtyDayYearlyInterestRate, a.Rows[i].ThirtyDayYearlyInterestRate, "ThirtyDayYearlyInterestRate")
		r.Equal(e.Rows[i].ThirtyDayDailyInterestRate, a.Rows[i].ThirtyDayDailyInterestRate, "ThirtyDayDailyInterestRate")
		r.Equal(e.Rows[i].SixtyDayYearlyInterestRate, a.Rows[i].SixtyDayYearlyInterestRate, "SixtyDayYearlyInterestRate")
		r.Equal(e.Rows[i].SixtyDayDailyInterestRate, a.Rows[i].SixtyDayDailyInterestRate, "SixtyDayDailyInterestRate")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanLoanCollateralData() {
	data := []byte(`{
  "rows": [
    {
      "collateralCoin": "BUSD",
      "_1stCollateralRatio": "100%",
      "_1stCollateralRange": "1-10000000",
      "_2ndCollateralRatio": "80%",
      "_2ndCollateralRange": "10000000-100000000",
      "_3rdCollateralRatio": "60%",
      "_3rdCollateralRange": "100000000-1000000000",
      "_4thCollateralRatio": "0%",
      "_4thCollateralRange": ">10000000000"
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().CollateralAssetData().CollateralCoin("BUSD").Do(newContext())
	s.r().NoError(err)
	e := &VipLoanCollateralAssetDataResponse{
		Rows: []VipLoanCollateralAssetData{
			{
				CollateralCoin:        "BUSD",
				FirstCollateralRatio:  "100%",
				FirstCollateralRange:  "1-10000000",
				SecondCollateralRatio: "80%",
				SecondCollateralRange: "10000000-100000000",
				ThirdCollateralRatio:  "60%",
				ThirdCollateralRange:  "100000000-1000000000",
				FourthCollateralRatio: "0%",
				FourthCollateralRange: ">10000000000",
			},
		},
		Total: 1,
	}

	s.assertVipLoanCollateralAssetDataResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanCollateralAssetDataResponseEqual(e, a VipLoanCollateralAssetDataResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].CollateralCoin, a.Rows[i].CollateralCoin, "CollateralCoin")
		r.Equal(e.Rows[i].FirstCollateralRatio, a.Rows[i].FirstCollateralRatio, "FirstCollateralRatio")
		r.Equal(e.Rows[i].FirstCollateralRange, a.Rows[i].FirstCollateralRange, "FirstCollateralRange")
		r.Equal(e.Rows[i].SecondCollateralRatio, a.Rows[i].SecondCollateralRatio, "SecondCollateralRatio")
		r.Equal(e.Rows[i].SecondCollateralRange, a.Rows[i].SecondCollateralRange, "SecondCollateralRange")
		r.Equal(e.Rows[i].ThirdCollateralRatio, a.Rows[i].ThirdCollateralRatio, "ThirdCollateralRatio")
		r.Equal(e.Rows[i].ThirdCollateralRange, a.Rows[i].ThirdCollateralRange, "ThirdCollateralRange")
		r.Equal(e.Rows[i].FourthCollateralRatio, a.Rows[i].FourthCollateralRatio, "FourthCollateralRatio")
		r.Equal(e.Rows[i].FourthCollateralRange, a.Rows[i].FourthCollateralRange, "FourthCollateralRange")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanOngoingOrders() {
	data := []byte(`{
  "rows": [
    {
      "orderId": 100000001,
      "loanCoin": "BUSD",
      "totalDebt": "10000",
      "loanRate": "0.0123",
      "residualInterest": "10.27687923",
      "collateralAccountId": "12345678,23456789",
      "collateralCoin": "BNB,BTC,ETH",
      "totalCollateralValueAfterHaircut": "25000.27565492",
      "lockedCollateralValue": "25000.27565492",
      "currentLTV": "0.57",
      "expirationTime": 1575018510000,
      "loanDate": "1676851200000",
      "loanTerm": "30days"
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().OngoingOrders().LoanCoin("BUSD").OrderId(100000001).Do(newContext())
	s.r().NoError(err)
	e := &VipLoanOngoingOrderResponse{
		Rows: []VipLoanOngoingOrder{
			{
				OrderId:                          100000001,
				LoanCoin:                         "BUSD",
				TotalDebt:                        "10000",
				LoanRate:                         "0.0123",
				ResidualInterest:                 "10.27687923",
				CollateralAccountId:              "12345678,23456789",
				CollateralCoin:                   "BNB,BTC,ETH",
				TotalCollateralValueAfterHaircut: "25000.27565492",
				LockedCollateralValue:            "25000.27565492",
				CurrentLTV:                       "0.57",
				ExpirationTime:                   1575018510000,
				LoanDate:                         "1676851200000",
				LoanTerm:                         "30days",
			},
		},
		Total: 1,
	}
	s.assertVipLoanOngoingOrderResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanOngoingOrderResponseEqual(e, a VipLoanOngoingOrderResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].OrderId, a.Rows[i].OrderId, "OrderId")
		r.Equal(e.Rows[i].LoanCoin, a.Rows[i].LoanCoin, "LoanCoin")
		r.Equal(e.Rows[i].TotalDebt, a.Rows[i].TotalDebt, "TotalDebt")
		r.Equal(e.Rows[i].LoanRate, a.Rows[i].LoanRate, "LoanRate")
		r.Equal(e.Rows[i].ResidualInterest, a.Rows[i].ResidualInterest, "ResidualInterest")
		r.Equal(e.Rows[i].CollateralAccountId, a.Rows[i].CollateralAccountId, "CollateralAccountId")
		r.Equal(e.Rows[i].CollateralCoin, a.Rows[i].CollateralCoin, "CollateralCoin")
		r.Equal(e.Rows[i].TotalCollateralValueAfterHaircut, a.Rows[i].TotalCollateralValueAfterHaircut, "TotalCollateralValueAfterHaircut")
		r.Equal(e.Rows[i].LockedCollateralValue, a.Rows[i].LockedCollateralValue, "LockedCollateralValue")
		r.Equal(e.Rows[i].CurrentLTV, a.Rows[i].CurrentLTV, "CurrentLTV")
		r.Equal(e.Rows[i].ExpirationTime, a.Rows[i].ExpirationTime, "ExpirationTime")
		r.Equal(e.Rows[i].LoanDate, a.Rows[i].LoanDate, "LoanDate")
		r.Equal(e.Rows[i].LoanTerm, a.Rows[i].LoanTerm, "LoanTerm")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanRepaymentHistory() {
	data := []byte(`{
  "rows": [
    {
      "loanCoin": "BUSD",
      "repayAmount": "10000",
      "collateralCoin": "BNB,BTC,ETH",
      "repayStatus": "Repaid",
      "loanDate": "1676851200000",
      "repayTime": "1575018510000",
      "orderId": "756783308056935434"
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().RepaymentHistory().LoanCoin("BUSD").OrderId(756783308056935434).Do(newContext())
	s.r().NoError(err)
	e := &VipLoanRepaymentHistoryResponse{
		Rows: []VipLoanRepaymentHistory{
			{
				LoanCoin:       "BUSD",
				RepayAmount:    "10000",
				CollateralCoin: "BNB,BTC,ETH",
				RepayStatus:    "Repaid",
				LoanDate:       "1676851200000",
				RepayTime:      "1575018510000",
				OrderId:        "756783308056935434",
			},
		},
		Total: 1,
	}
	s.assertVipLoanRepaymentRecordResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanRepaymentRecordResponseEqual(e, a VipLoanRepaymentHistoryResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].LoanCoin, a.Rows[i].LoanCoin, "LoanCoin")
		r.Equal(e.Rows[i].RepayAmount, a.Rows[i].RepayAmount, "RepayAmount")
		r.Equal(e.Rows[i].CollateralCoin, a.Rows[i].CollateralCoin, "CollateralCoin")
		r.Equal(e.Rows[i].RepayStatus, a.Rows[i].RepayStatus, "RepayStatus")
		r.Equal(e.Rows[i].LoanDate, a.Rows[i].LoanDate, "LoanDate")
		r.Equal(e.Rows[i].RepayTime, a.Rows[i].RepayTime, "RepayTime")
		r.Equal(e.Rows[i].OrderId, a.Rows[i].OrderId, "OrderId")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanAccruedInterest() {
	data := []byte(`{
  "rows": [
    {
      "loanCoin": "USDT",
      "principalAmount": "10000",
      "interestAmount": "1.2",
      "annualInterestRate": "0.001273",
      "accrualTime": 1575018510000,
      "orderId": 756783308056935434
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().AccruedInterest().LoanCoin("USDT").OrderId(756783308056935434).Do(newContext())
	s.r().NoError(err)
	e := &VipLoanAccruedInterestResponse{
		Rows: []VipLoanAccruedInterest{
			{
				LoanCoin:           "USDT",
				PrincipalAmount:    "10000",
				InterestAmount:     "1.2",
				AnnualInterestRate: "0.001273",
				AccrualTime:        1575018510000,
				OrderId:            756783308056935434,
			},
		},
		Total: 1,
	}
	s.assertVipLoanAccruedInterestResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanAccruedInterestResponseEqual(e, a VipLoanAccruedInterestResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].LoanCoin, a.Rows[i].LoanCoin, "LoanCoin")
		r.Equal(e.Rows[i].PrincipalAmount, a.Rows[i].PrincipalAmount, "PrincipalAmount")
		r.Equal(e.Rows[i].InterestAmount, a.Rows[i].InterestAmount, "InterestAmount")
		r.Equal(e.Rows[i].AnnualInterestRate, a.Rows[i].AnnualInterestRate, "AnnualInterestRate")
		r.Equal(e.Rows[i].AccrualTime, a.Rows[i].AccrualTime, "AccrualTime")
		r.Equal(e.Rows[i].OrderId, a.Rows[i].OrderId, "OrderId")
	}
}

func (s *vipLoanServiceTestSuite) TestVipCollateralAccount() {
	data := []byte(`{
  "rows": [
    {
      "collateralAccountId": "12345678",
      "collateralCoin": "BNB,BTC,ETH"
    },
    {
      "collateralAccountId": "23456789",
      "collateralCoin": "BNB,BTC,ETH"
    }
  ],
  "total": 2
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewVipLoanService().CollateralAccount().Do(newContext())
	s.r().NoError(err)
	e := &VipLoanCollateralAccountResponse{
		Rows: []VipLoanCollateralAccount{
			{
				CollateralAccountId: "12345678",
				CollateralCoin:      "BNB,BTC,ETH",
			},
			{
				CollateralAccountId: "23456789",
				CollateralCoin:      "BNB,BTC,ETH",
			},
		},
		Total: 2,
	}
	s.assertVipLoanCollateralAccountResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanCollateralAccountResponseEqual(e, a VipLoanCollateralAccountResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].CollateralAccountId, a.Rows[i].CollateralAccountId, "CollateralAccountId")
		r.Equal(e.Rows[i].CollateralCoin, a.Rows[i].CollateralCoin, "CollateralCoin")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanApplicationStatus() {
	data := []byte(`{
  "rows": [
    {
      "loanAccountId": "12345678",
      "orderId": "12345678",
      "requestId": "12345678",
      "loanCoin": "BTC",
      "loanAmount": "100.55",
      "collateralAccountId": "12345678,12345678,12345678",
      "collateralCoin": "BUSD,USDT,ETH",
      "loanTerm": "30",
      "status": "Repaid",
      "loanDate": "1676851200000"
    }
  ],
  "total": 1
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewVipLoanService().ApplicationStatus().Do(newContext())
	s.r().NoError(err)
	e := &VipLoanApplicationStatusResponse{
		Rows: []VipLoanApplicationStatus{
			{
				LoanAccountId: "12345678",
				OrderId:       "12345678",
				RequestId:     "12345678",
				LoanCoin:      "BTC",
				LoanAmount:    "100.55",
				LoanTerm:      "30",
				Status:        "Repaid",
				LoanDate:      "1676851200000",
			},
		},
		Total: 1,
	}
	s.assertVipLoanApplicationStatusResponseEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanApplicationStatusResponseEqual(e, a VipLoanApplicationStatusResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total)
	for i := 0; i < len(e.Rows); i++ {
		r.Equal(e.Rows[i].LoanAccountId, a.Rows[i].LoanAccountId, "LoanAccountId")
		r.Equal(e.Rows[i].OrderId, a.Rows[i].OrderId, "OrderId")
		r.Equal(e.Rows[i].RequestId, a.Rows[i].RequestId, "RequestId")
		r.Equal(e.Rows[i].LoanCoin, a.Rows[i].LoanCoin, "LoanCoin")
		r.Equal(e.Rows[i].LoanAmount, a.Rows[i].LoanAmount, "LoanAmount")
		r.Equal(e.Rows[i].LoanTerm, a.Rows[i].LoanTerm, "LoanTerm")
		r.Equal(e.Rows[i].Status, a.Rows[i].Status, "Status")
		r.Equal(e.Rows[i].LoanDate, a.Rows[i].LoanDate, "LoanDate")
	}
}

func (s *vipLoanServiceTestSuite) TestVipLoanRenew() {
	data := []byte(`{
  "loanAccountId": "12345678",
  "loanCoin": "BTC",
  "loanAmount": "100.55",
  "collateralAccountId": "12345677,12345678,12345679",
  "collateralCoin": "BUSD,USDT,ETH",
  "loanTerm": "30"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewVipLoanService().Renew().OrderId(12345678).LoanTerm(30).Do(newContext())
	s.r().NoError(err)
	e := &VipLoanRenew{LoanCoin: "BTC", LoanAmount: "100.55", LoanTerm: "30", CollateralAccountId: "12345677,12345678,12345679", CollateralCoin: "BUSD,USDT,ETH"}
	s.assertVipLoanRenewEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertVipLoanRenewEqual(e, a VipLoanRenew) {
	r := s.r()
	r.Equal(e.LoanCoin, a.LoanCoin, "LoanCoin")
	r.Equal(e.LoanAmount, a.LoanAmount, "LoanAmount")
	r.Equal(e.LoanTerm, a.LoanTerm, "LoanTerm")
	r.Equal(e.CollateralAccountId, a.CollateralAccountId, "CollateralAccountId")
	r.Equal(e.CollateralCoin, a.CollateralCoin, "CollateralCoin")
}

func (s *vipLoanServiceTestSuite) TestVipLoanRepay() {
	data := []byte(`{
  "loanCoin": "BUSD",
  "repayAmount": "200.5",
  "remainingPrincipal": "100.5",
  "remainingInterest": "0",
  "collateralCoin": "BNB,BTC,ETH",
  "currentLTV": "0.25",
  "repayStatus": "Repaid"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewVipLoanService().Repay().OrderId(12345678).Amount(200.5).Do(newContext())
	s.r().NoError(err)

	e := &VipLoanRepay{
		LoanCoin:           "BUSD",
		RepayAmount:        "200.5",
		RemainingPrincipal: "100.5",
		RemainingInterest:  "0",
		CollateralCoin:     "BNB,BTC,ETH",
		CurrentLTV:         "0.25",
		RepayStatus:        "Repaid",
	}

	s.assertRepayEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertRepayEqual(e, a VipLoanRepay) {
	r := s.r()
	r.Equal(e.LoanCoin, a.LoanCoin, "LoanCoin")
	r.Equal(e.RepayAmount, a.RepayAmount, "RepayAmount")
	r.Equal(e.RemainingPrincipal, a.RemainingPrincipal, "RemainingPrincipal")
	r.Equal(e.RemainingInterest, a.RemainingInterest, "RemainingInterest")
	r.Equal(e.CollateralCoin, a.CollateralCoin, "CollateralCoin")
	r.Equal(e.CurrentLTV, a.CurrentLTV, "CurrentLTV")
	r.Equal(e.RepayStatus, a.RepayStatus, "RepayStatus")
}

func (s *vipLoanServiceTestSuite) TestVipLoanBorrow() {
	data := []byte(`{
  "loanAccountId": "12345678",
  "requestId": "12345678",
  "loanCoin": "BTC",
  "isFlexibleRate": "No",
  "loanAmount": "100.55",
  "collateralAccountId": "12345678,12345678,12345678",
  "collateralCoin": "BUSD,USDT,ETH",
  "loanTerm": "30"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewVipLoanService().Borrow().LoanAccountId(12345678).LoanCoin("BTC").
		CollateralAccountId("12345678,12345678,12345678").CollateralCoin("BUSD,USDT,ETH").
		IsFlexibleRate(false).LoanTerm(30).Do(newContext())
	s.r().NoError(err)

	e := &VipLoanBorrow{
		LoanAccountId:       "12345678",
		RequestId:           "12345678",
		LoanCoin:            "BTC",
		IsFlexibleRate:      "No",
		LoanAmount:          "100.55",
		CollateralAccountId: "12345678,12345678,12345678",
		CollateralCoin:      "BUSD,USDT,ETH",
		LoanTerm:            "30",
	}

	s.assertBorrowEqual(*e, *res)
}

func (s *vipLoanServiceTestSuite) assertBorrowEqual(e, a VipLoanBorrow) {
	r := s.r()
	r.Equal(e.LoanAccountId, a.LoanAccountId, "LoanAccountId")
	r.Equal(e.RequestId, a.RequestId, "RequestId")
	r.Equal(e.LoanCoin, a.LoanCoin, "LoanCoin")
	r.Equal(e.IsFlexibleRate, a.IsFlexibleRate, "IsFlexibleRate")
	r.Equal(e.LoanAmount, a.LoanAmount, "LoanAmount")
	r.Equal(e.CollateralAccountId, a.CollateralAccountId, "CollateralAccountId")
	r.Equal(e.CollateralCoin, a.CollateralCoin, "CollateralCoin")
	r.Equal(e.LoanTerm, a.LoanTerm, "LoanTerm")
}
