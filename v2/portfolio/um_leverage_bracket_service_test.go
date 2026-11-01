package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umLeverageBracketServiceTestSuite struct {
	baseTestSuite
}

func TestUMLeverageBracketService(t *testing.T) {
	suite.Run(t, new(umLeverageBracketServiceTestSuite))
}

func (s *umLeverageBracketServiceTestSuite) TestGetLeverageBracket() {
	data := []byte(`[
		{
			"symbol": "ETHUSDT",
			"notionalCoef": "4.0",
			"brackets": [
				{
					"bracket": 1,
					"initialLeverage": 75,
					"notionalCap": 10000,
					"notionalFloor": 0,
					"maintMarginRatio": 0.0065,
					"cum": 0
				}
			]
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "ETHUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	brackets, err := s.client.NewGetUMLeverageBracketService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(brackets, 1)
	s.assertLeverageBracketEqual(brackets[0], &LeverageBracket{
		Symbol:       "ETHUSDT",
		NotionalCoef: "4.0",
		Brackets: []Bracket{
			{
				Bracket:          1,
				InitialLeverage:  75,
				NotionalCap:      10000,
				NotionalFloor:    0,
				MaintMarginRatio: 0.0065,
				Cum:              0,
			},
		},
	})
}

func (s *umLeverageBracketServiceTestSuite) assertLeverageBracketEqual(a, e *LeverageBracket) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.NotionalCoef, a.NotionalCoef, "NotionalCoef")
	r.Len(a.Brackets, len(e.Brackets))
	for i := range e.Brackets {
		r.Equal(e.Brackets[i].Bracket, a.Brackets[i].Bracket, "Bracket")
		r.Equal(e.Brackets[i].InitialLeverage, a.Brackets[i].InitialLeverage, "InitialLeverage")
		r.Equal(e.Brackets[i].NotionalCap, a.Brackets[i].NotionalCap, "NotionalCap")
		r.Equal(e.Brackets[i].NotionalFloor, a.Brackets[i].NotionalFloor, "NotionalFloor")
		r.Equal(e.Brackets[i].MaintMarginRatio, a.Brackets[i].MaintMarginRatio, "MaintMarginRatio")
		r.Equal(e.Brackets[i].Cum, a.Brackets[i].Cum, "Cum")
	}
}
