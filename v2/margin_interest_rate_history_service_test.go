package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginInterestRateHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestMarginInterestRateHistoryService(t *testing.T) {
	suite.Run(t, new(marginInterestRateHistoryServiceTestSuite))
}

func (s *marginInterestRateHistoryServiceTestSuite) TestMarginInterestRateHistory() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"dailyInterestRate": "0.00025000",
			"timestamp": 1611544731000,
			"vipLevel": 1    
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": asset,
		})
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewMarginInterestRateHistoryService().
		Asset(asset).
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	s.Len(*history, 1)

	item := &(*history)[0]
	s.Equal(int64(1611544731000), item.Timestamp)
	s.Equal(int64(1), item.VipLevel)
	s.Equal("0.00025000", item.DailyInterestRate)
	s.Equal("BTC", item.Asset)
}
