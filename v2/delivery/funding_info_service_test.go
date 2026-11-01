package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fundingInfoServiceTestSuite struct {
	baseTestSuite
}

func TestFundingInfoService(t *testing.T) {
	suite.Run(t, new(fundingInfoServiceTestSuite))
}

func (s *fundingInfoServiceTestSuite) TestGetFundingInfo() {
	data := []byte(`[
		{
			"symbol": "BLZUSDT",
			"adjustedFundingRateCap": "0.3",
			"adjustedFundingRateFloor": "-0.3",
			"fundingIntervalHours": 8,
			"disclaimer": true
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	fundingInfo, err := s.client.NewGetFundingInfoService().Do(newContext())
	s.r().NoError(err)
	s.Len(fundingInfo, 1)
	e := &FundingInfo{
		Symbol:                   "BLZUSDT",
		AdjustedFundingRateCap:   "0.3",
		AdjustedFundingRateFloor: "-0.3",
		FundingIntervalHours:     8,
		Disclaimer:               true,
	}
	s.assertFundingInfoEqual(e, fundingInfo[0])
}

func (s *fundingInfoServiceTestSuite) assertFundingInfoEqual(e, a *FundingInfo) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AdjustedFundingRateCap, a.AdjustedFundingRateCap, "AdjustedFundingRateCap")
	r.Equal(e.AdjustedFundingRateFloor, a.AdjustedFundingRateFloor, "AdjustedFundingRateFloor")
	r.Equal(e.FundingIntervalHours, a.FundingIntervalHours, "FundingIntervalHours")
	r.Equal(e.Disclaimer, a.Disclaimer, "Disclaimer")
}
