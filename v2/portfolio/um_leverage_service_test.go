package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umLeverageServiceTestSuite struct {
	baseTestSuite
}

func TestUMLeverageService(t *testing.T) {
	suite.Run(t, new(umLeverageServiceTestSuite))
}

func (s *umLeverageServiceTestSuite) TestChangeLeverage() {
	data := []byte(`{
		"leverage": 21,
		"maxNotionalValue": "1000000",
		"symbol": "BTCUSDT"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	leverage := 21
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("leverage", leverage)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewChangeUMInitialLeverageService().
		Symbol(symbol).
		Leverage(leverage).
		Do(newContext())
	s.r().NoError(err)
	s.assertLeverageEqual(res, &UMLeverage{
		Leverage:         21,
		MaxNotionalValue: "1000000",
		Symbol:           "BTCUSDT",
	})
}

func (s *umLeverageServiceTestSuite) assertLeverageEqual(a, e *UMLeverage) {
	r := s.r()
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.MaxNotionalValue, a.MaxNotionalValue, "MaxNotionalValue")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}
