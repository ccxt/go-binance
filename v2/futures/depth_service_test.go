package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type depthServiceTestSuite struct {
	baseTestSuite
}

func TestDepthService(t *testing.T) {
	suite.Run(t, new(depthServiceTestSuite))
}

func (s *depthServiceTestSuite) TestDepth() {
	data := []byte(`{
        "lastUpdateId": 1027024,
        "bids": [
            [
                "4.00000000",
                "431.00000000"
            ]
        ],
        "asks": [
            [
                "4.00000200",
                "12.00000000"
            ]
        ]
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "LTCBTC"
	limit := 3
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).
			setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, _, err := s.client.NewDepthService().Symbol(symbol).Limit(limit).Do(newContext())
	s.r().NoError(err)
	e := &DepthResponse{
		LastUpdateID: 1027024,
		Bids: []Bid{
			{
				Price:    "4.00000000",
				Quantity: "431.00000000",
			},
		},
		Asks: []Ask{
			{
				Price:    "4.00000200",
				Quantity: "12.00000000",
			},
		},
	}
	s.assertDepthResponseEqual(e, res)
}

func (s *depthServiceTestSuite) assertDepthResponseEqual(e, a *DepthResponse) {
	r := s.r()
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	r.Len(a.Bids, len(e.Bids))
	for i := 0; i < len(a.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	r.Len(a.Asks, len(e.Asks))
	for i := 0; i < len(a.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}
