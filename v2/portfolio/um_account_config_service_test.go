package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umAccountConfigServiceTestSuite struct {
	baseTestSuite
}

func TestUMAccountConfigService(t *testing.T) {
	suite.Run(t, new(umAccountConfigServiceTestSuite))
}

func (s *umAccountConfigServiceTestSuite) TestGetUMAccountConfig() {
	data := []byte(`{
		"feeTier": 0,
		"canTrade": true,
		"canDeposit": true,
		"canWithdraw": true,
		"dualSidePosition": true,
		"updateTime": 1724416653850,
		"multiAssetsMargin": false,
		"tradeGroupId": -1
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	expected := &UMAccountConfig{
		FeeTier:           0,
		CanTrade:          true,
		CanDeposit:        true,
		CanWithdraw:       true,
		DualSidePosition:  true,
		UpdateTime:        1724416653850,
		MultiAssetsMargin: false,
		TradeGroupId:      -1,
	}

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	config, err := s.client.NewGetUMAccountConfigService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal(expected, config)
}
