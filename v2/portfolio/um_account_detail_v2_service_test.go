package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAccountDetailV2ServiceTestSuite struct {
	baseTestSuite
}

func TestUMAccountDetailV2Service(t *testing.T) {
	suite.Run(t, new(umAccountDetailV2ServiceTestSuite))
}

func (s *umAccountDetailV2ServiceTestSuite) TestGetUMAccountDetailV2() {
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
				"positionSide": "BOTH",
				"positionAmt": "0",
				"updateTime": 0,
				"notional": "86.98650000"
			}
		]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetUMAccountDetailV2Service().Do(newContext())
	s.r().NoError(err)

	// Validate assets
	s.r().Len(res.Assets, 1)
	s.r().Equal("USDT", res.Assets[0].Asset)
	s.r().Equal("23.72469206", res.Assets[0].CrossWalletBalance)
	s.r().Equal(int64(1625474304765), res.Assets[0].UpdateTime)

	// Validate positions
	s.r().Len(res.Positions, 1)
	s.r().Equal("BTCUSDT", res.Positions[0].Symbol)
	s.r().Equal("BOTH", res.Positions[0].PositionSide)
	s.r().Equal("86.98650000", res.Positions[0].Notional)
}
