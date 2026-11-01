package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OpenInterestServiceTestSuite struct {
	baseTestSuite
}

func TestOpenInterestService(t *testing.T) {
	suite.Run(t, new(OpenInterestServiceTestSuite))
}

func (s *OpenInterestServiceTestSuite) TestOpenInterest() {
	data := []byte(`[
		{
		 "symbol": "BTC-240628-36000-P",
		 "sumOpenInterest": "17.250",
		 "sumOpenInterestUsd": "1174670.00553185250",
		 "timestamp": "1717036680000"
		},
		{
		 "symbol": "BTC-240628-110000-C",
		 "sumOpenInterest": "25.180",
		 "sumOpenInterestUsd": "1714677.72401693020",
		 "timestamp": "1717036680000"
		}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	OIs, err := s.client.NewOpenInterestService().Do(newContext())
	targetOIs := []*OpenInterest{
		{
			Symbol:             "BTC-240628-36000-P",
			SumOpenInterest:    "17.250",
			SumOpenInterestUsd: "1174670.00553185250",
			Timestamp:          "1717036680000",
		},
		{
			Symbol:             "BTC-240628-110000-C",
			SumOpenInterest:    "25.180",
			SumOpenInterestUsd: "1714677.72401693020",
			Timestamp:          "1717036680000",
		},
	}

	s.r().Equal(err, nil, "err != nil")
	for i := range OIs {
		r := s.r()
		r.Equal(OIs[i].Symbol, targetOIs[i].Symbol, "Symbol")
		r.Equal(OIs[i].SumOpenInterest, targetOIs[i].SumOpenInterest, "SumOpenInterest")
		r.Equal(OIs[i].SumOpenInterestUsd, targetOIs[i].SumOpenInterestUsd, "SumOpenInterestUsd")
		r.Equal(OIs[i].Timestamp, targetOIs[i].Timestamp, "Timestamp")
	}
}
