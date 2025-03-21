package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umCommissionRateServiceTestSuite struct {
	baseTestSuite
}

func TestUMCommissionRateService(t *testing.T) {
	suite.Run(t, new(umCommissionRateServiceTestSuite))
}

func (s *umCommissionRateServiceTestSuite) TestGetCommissionRate() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"makerCommissionRate": "0.0002",
		"takerCommissionRate": "0.0004"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	rates, err := s.client.NewGetUMCommissionRateService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("BTCUSDT", rates.Symbol)
	s.r().Equal("0.0002", rates.MakerCommissionRate)
	s.r().Equal("0.0004", rates.TakerCommissionRate)
}
