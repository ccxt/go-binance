package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type negativeBalanceInterestHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestNegativeBalanceInterestHistoryService(t *testing.T) {
	suite.Run(t, new(negativeBalanceInterestHistoryServiceTestSuite))
}

func (s *negativeBalanceInterestHistoryServiceTestSuite) TestGetNegativeBalanceInterestHistory() {
	data := []byte(`[
		{
			"asset": "USDT",
			"interest": "24.4440",
			"interestAccuredTime": 1670227200000,
			"interestRate": "0.0001164",
			"principal": "210000"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "USDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		s.assertRequestEqual(e, r)
	})

	interests, err := s.client.NewGetNegativeBalanceInterestHistoryService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	s.r().Len(interests, 1)

	interest := interests[0]
	s.r().Equal("USDT", interest.Asset)
	s.r().Equal("24.4440", interest.Interest)
	s.r().Equal(int64(1670227200000), interest.InterestAccuredTime)
	s.r().Equal("0.0001164", interest.InterestRate)
	s.r().Equal("210000", interest.Principal)
}
