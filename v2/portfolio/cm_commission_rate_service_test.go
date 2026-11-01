package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmCommissionRateServiceTestSuite struct {
	baseTestSuite
}

func TestCMCommissionRateService(t *testing.T) {
	suite.Run(t, new(cmCommissionRateServiceTestSuite))
}

func (s *cmCommissionRateServiceTestSuite) TestGetCommissionRate() {
	data := []byte(`{
		"symbol": "BTCUSD_PERP",
		"makerCommissionRate": "0.00015",
		"takerCommissionRate": "0.00040"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	rates, err := s.client.NewGetCMCommissionRateService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("BTCUSD_PERP", rates.Symbol)
	s.r().Equal("0.00015", rates.MakerCommissionRate)
	s.r().Equal("0.00040", rates.TakerCommissionRate)
}
