package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MarkServiceTestSuite struct {
	baseTestSuite
}

func TestMarkService(t *testing.T) {
	suite.Run(t, new(MarkServiceTestSuite))
}

func (s *MarkServiceTestSuite) TestMark() {
	data := []byte(`[
		{
		 "symbol": "ETH-240628-800-C",
		 "markPrice": "2986.3",
		 "bidIV": "0.0000006",
		 "askIV": "-0.00000001",
		 "markIV": "0.575",
		 "delta": "1",
		 "theta": "0",
		 "gamma": "0",
		 "vega": "0",
		 "highPriceLimit": "3269.7",
		 "lowPriceLimit": "2702.9",
		 "riskFreeInterest": "0.1"
		}
	   ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	marks, err := s.client.NewMarkService().Do(newContext())
	targetMarks := []*Mark{
		{
			Symbol:           "ETH-240628-800-C",
			MarkPrice:        "2986.3",
			BidIV:            "0.0000006",
			AskIV:            "-0.00000001",
			MarkIV:           "0.575",
			Delta:            "1",
			Theta:            "0",
			Gamma:            "0",
			Vega:             "0",
			HighPriceLimit:   "3269.7",
			LowPriceLimit:    "2702.9",
			RiskFreeInterest: "0.1",
		},
	}

	s.r().Equal(err, nil, "err != nil")
	for i := range marks {
		r := s.r()
		r.Equal(marks[i].Symbol, targetMarks[i].Symbol, "Symbol")
		r.Equal(marks[i].MarkPrice, targetMarks[i].MarkPrice, "MarkPrice")
		r.Equal(marks[i].BidIV, targetMarks[i].BidIV, "BidIV")
		r.Equal(marks[i].AskIV, targetMarks[i].AskIV, "AskIV")
		r.Equal(marks[i].MarkIV, targetMarks[i].MarkIV, "MarkIV")
		r.Equal(marks[i].Delta, targetMarks[i].Delta, "Delta")
		r.Equal(marks[i].Theta, targetMarks[i].Theta, "Theta")
		r.Equal(marks[i].Gamma, targetMarks[i].Gamma, "Gamma")
		r.Equal(marks[i].Vega, targetMarks[i].Vega, "Vega")
		r.Equal(marks[i].HighPriceLimit, targetMarks[i].HighPriceLimit, "HighPriceLimit")
		r.Equal(marks[i].LowPriceLimit, targetMarks[i].LowPriceLimit, "LowPriceLimit")
		r.Equal(marks[i].RiskFreeInterest, targetMarks[i].RiskFreeInterest, "RiskFreeInterest")
	}
}
