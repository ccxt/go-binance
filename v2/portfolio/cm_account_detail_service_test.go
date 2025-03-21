package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmAccountDetailServiceTestSuite struct {
	baseTestSuite
}

func TestCMAccountDetailService(t *testing.T) {
	suite.Run(t, new(cmAccountDetailServiceTestSuite))
}

func (s *cmAccountDetailServiceTestSuite) TestGetCMAccountDetail() {
	data := []byte(`{
		"assets": [
			{
				"asset": "BTC",
				"crossWalletBalance": "0.00241969",
				"crossUnPnl": "0.00000000",
				"maintMargin": "0.00000000",
				"initialMargin": "0.00000000",
				"positionInitialMargin": "0.00000000",
				"openOrderInitialMargin": "0.00000000",
				"updateTime": 1625474304765
			}
		],
		"positions": [
			{
				"symbol": "BTCUSD_201225",
				"positionAmt": "0",
				"initialMargin": "0",
				"maintMargin": "0",
				"unrealizedProfit": "0.00000000",
				"positionInitialMargin": "0",
				"openOrderInitialMargin": "0",
				"leverage": "125",
				"positionSide": "BOTH",
				"entryPrice": "0.0",
				"maxQty": "50",
				"updateTime": 0
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetCMAccountDetailService().Do(newContext())
	s.r().NoError(err)

	s.r().Len(res.Assets, 1)
	s.r().Equal("BTC", res.Assets[0].Asset)
	s.r().Equal("0.00241969", res.Assets[0].CrossWalletBalance)
	s.r().Equal("0.00000000", res.Assets[0].CrossUnPnl)
	s.r().Equal("0.00000000", res.Assets[0].MaintMargin)
	s.r().Equal("0.00000000", res.Assets[0].InitialMargin)
	s.r().Equal("0.00000000", res.Assets[0].PositionInitialMargin)
	s.r().Equal("0.00000000", res.Assets[0].OpenOrderInitialMargin)
	s.r().Equal(int64(1625474304765), res.Assets[0].UpdateTime)

	s.r().Len(res.Positions, 1)
	s.r().Equal("BTCUSD_201225", res.Positions[0].Symbol)
	s.r().Equal("0", res.Positions[0].PositionAmt)
	s.r().Equal("0", res.Positions[0].InitialMargin)
	s.r().Equal("0", res.Positions[0].MaintMargin)
	s.r().Equal("0.00000000", res.Positions[0].UnrealizedProfit)
	s.r().Equal("0", res.Positions[0].PositionInitialMargin)
	s.r().Equal("0", res.Positions[0].OpenOrderInitialMargin)
	s.r().Equal("125", res.Positions[0].Leverage)
	s.r().Equal("BOTH", res.Positions[0].PositionSide)
	s.r().Equal("0.0", res.Positions[0].EntryPrice)
	s.r().Equal("50", res.Positions[0].MaxQty)
	s.r().Equal(int64(0), res.Positions[0].UpdateTime)
}
