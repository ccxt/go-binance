package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fundingRateServiceTestSuite struct {
	baseTestSuite
}

func TestFundingRateService(t *testing.T) {
	suite.Run(t, new(fundingRateServiceTestSuite))
}

func (s *fundingRateServiceTestSuite) TestGetFundingRate() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"fundingRate": "-0.03750000",
			"fundingTime": 1570608000000,
			"markPrice": "1576566020000"
		},
		{
			"symbol": "BTCUSDT",
			"fundingRate": "0.00010000",
			"fundingTime": 1570636800000,
			"markPrice": "1576566020000"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	startTime := int64(1576566020000)
	endTime := int64(1676566020000)
	limit := 10
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFundingRateService().Symbol(symbol).StartTime(startTime).
		EndTime(endTime).Limit(limit).Do(newContext())
	s.r().NoError(err)
	e := []*FundingRate{
		{
			Symbol:      symbol,
			FundingRate: "-0.03750000",
			FundingTime: int64(1570608000000),
			MarkPrice:   "1576566020000",
		},
		{
			Symbol:      symbol,
			FundingRate: "0.00010000",
			FundingTime: int64(1570636800000),
			MarkPrice:   "1576566020000",
		},
	}
	s.r().Len(res, len(e))
	for i := 0; i < len(res); i++ {
		s.assertFundingRateEqual(e[i], res[i])
	}
}

func (s *fundingRateServiceTestSuite) assertFundingRateEqual(e, a *FundingRate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	r.Equal(e.FundingTime, a.FundingTime, "FundingTime")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
}
