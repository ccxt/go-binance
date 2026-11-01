package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginNextHourlyInterestRateServiceTestSuite struct {
	baseTestSuite
}

func TestMarginNextHourlyInterestRateService(t *testing.T) {
	suite.Run(t, new(marginNextHourlyInterestRateServiceTestSuite))
}

func (s *marginNextHourlyInterestRateServiceTestSuite) TestMarginNextHourlyInterestRate() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"nextHourlyInterestRate": "0.00000571"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	assets := "BTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"assets": assets,
		})
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewMarginNextHourlyInterestRateService().
		Assets(assets).
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	s.Len(*history, 1)

	item := &(*history)[0]
	s.Equal("BTC", item.Asset)
	s.Equal("0.00000571", item.NextHourlyInterestRate)
}
