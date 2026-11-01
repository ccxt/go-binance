package futures

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type fundingRateInfoServiceTestSuite struct {
	baseTestSuite
}

func TestFundingRateInfoService(t *testing.T) {
	suite.Run(t, new(fundingRateInfoServiceTestSuite))
}

func (s *fundingRateInfoServiceTestSuite) TestGetFundingRateInfo() {
	data := []byte(`[{
		"symbol": "BTCUSDT",
		"adjustedFundingRateCap": "0.02500000",
		"adjustedFundingRateFloor": "-0.02500000",
		"fundingIntervalHours": 8
	}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFundingRateInfoService().Do(newContext())
	s.r().NoError(err)
	e := []*FundingRateInfo{
		{
			Symbol:                   "BTCUSDT",
			AdjustedFundingRateCap:   "0.02500000",
			AdjustedFundingRateFloor: "-0.02500000",
			FundingIntervalHours:     8,
		},
	}
	s.assertFundingRateInfoEqual(e, res)
}

func (s *fundingRateInfoServiceTestSuite) assertFundingRateInfoEqual(e, a []*FundingRateInfo) {
	r := s.r()
	r.Equal(e[0].Symbol, a[0].Symbol, "Symbol")
	r.Equal(e[0].AdjustedFundingRateCap, a[0].AdjustedFundingRateCap, "AdjustedFundingRateCap")
	r.Equal(e[0].AdjustedFundingRateFloor, a[0].AdjustedFundingRateFloor, "AdjustedFundingRateFloor")
	r.Equal(e[0].FundingIntervalHours, a[0].FundingIntervalHours, "FundingIntervalHours")
}
