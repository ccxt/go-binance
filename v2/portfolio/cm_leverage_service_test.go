package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmLeverageServiceTestSuite struct {
	baseTestSuite
}

func TestCMLeverageService(t *testing.T) {
	suite.Run(t, new(cmLeverageServiceTestSuite))
}

func (s *cmLeverageServiceTestSuite) TestChangeLeverage() {
	data := []byte(`{
		"leverage": 21,
		"maxQty": "1000",
		"symbol": "BTCUSD_200925"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	leverage := 21
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("leverage", leverage)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewChangeCMInitialLeverageService().
		Symbol(symbol).
		Leverage(leverage).
		Do(newContext())
	s.r().NoError(err)
	s.assertLeverageEqual(res, &CMLeverage{
		Leverage: 21,
		MaxQty:   "1000",
		Symbol:   "BTCUSD_200925",
	})
}

func (s *cmLeverageServiceTestSuite) assertLeverageEqual(a, e *CMLeverage) {
	r := s.r()
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.MaxQty, a.MaxQty, "MaxQty")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}
