package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAccountDetailServiceTestSuite struct {
	baseTestSuite
}

func TestUMAccountDetailService(t *testing.T) {
	suite.Run(t, new(umAccountDetailServiceTestSuite))
}

func (s *umAccountDetailServiceTestSuite) TestGetUMAccountDetail() {
	data := []byte(`{
		"assets": [
			{
				"asset": "USDT",
				"crossWalletBalance": "23.72469206",
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
				"symbol": "BTCUSDT",
				"initialMargin": "0",
				"maintMargin": "0",
				"unrealizedProfit": "0.00000000",
				"positionInitialMargin": "0",
				"openOrderInitialMargin": "0",
				"leverage": "100",
				"entryPrice": "0.00000",
				"maxNotional": "250000",
				"bidNotional": "0",
				"askNotional": "0",
				"positionSide": "BOTH",
				"positionAmt": "0",
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

	res, err := s.client.NewGetUMAccountDetailService().Do(newContext())
	s.r().NoError(err)

	s.r().Len(res.Assets, 1)
	s.r().Equal("USDT", res.Assets[0].Asset)
	s.r().Equal("23.72469206", res.Assets[0].CrossWalletBalance)
	s.r().Equal("0.00000000", res.Assets[0].CrossUnPnl)
	s.r().Equal("0.00000000", res.Assets[0].MaintMargin)
	s.r().Equal("0.00000000", res.Assets[0].InitialMargin)
	s.r().Equal("0.00000000", res.Assets[0].PositionInitialMargin)
	s.r().Equal("0.00000000", res.Assets[0].OpenOrderInitialMargin)
	s.r().Equal(int64(1625474304765), res.Assets[0].UpdateTime)

	s.r().Len(res.Positions, 1)
	s.r().Equal("BTCUSDT", res.Positions[0].Symbol)
	s.r().Equal("0", res.Positions[0].InitialMargin)
	s.r().Equal("0", res.Positions[0].MaintMargin)
	s.r().Equal("0.00000000", res.Positions[0].UnrealizedProfit)
	s.r().Equal("0", res.Positions[0].PositionInitialMargin)
	s.r().Equal("0", res.Positions[0].OpenOrderInitialMargin)
	s.r().Equal("100", res.Positions[0].Leverage)
	s.r().Equal("0.00000", res.Positions[0].EntryPrice)
	s.r().Equal("250000", res.Positions[0].MaxNotional)
	s.r().Equal("0", res.Positions[0].BidNotional)
	s.r().Equal("0", res.Positions[0].AskNotional)
	s.r().Equal("BOTH", res.Positions[0].PositionSide)
	s.r().Equal("0", res.Positions[0].PositionAmt)
	s.r().Equal(int64(0), res.Positions[0].UpdateTime)
}
