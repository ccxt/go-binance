package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmLeverageBracketServiceTestSuite struct {
	baseTestSuite
}

func TestCMLeverageBracketService(t *testing.T) {
	suite.Run(t, new(cmLeverageBracketServiceTestSuite))
}

func (s *cmLeverageBracketServiceTestSuite) TestGetLeverageBracket() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_PERP",
			"brackets": [
				{
					"bracket": 1,
					"initialLeverage": 125,
					"qtyCap": 50,
					"qtyFloor": 0,
					"maintMarginRatio": 0.004,
					"cum": 0.0
				}
			]
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	brackets, err := s.client.NewGetCMLeverageBracketService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(brackets, 1)
	s.assertLeverageBracketEqual(brackets[0], &CMLeverageBracket{
		Symbol: "BTCUSD_PERP",
		Brackets: []CMBracket{
			{
				Bracket:          1,
				InitialLeverage:  125,
				QtyCap:           50,
				QtyFloor:         0,
				MaintMarginRatio: 0.004,
				Cum:              0.0,
			},
		},
	})
}

func (s *cmLeverageBracketServiceTestSuite) assertLeverageBracketEqual(a, e *CMLeverageBracket) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Len(a.Brackets, len(e.Brackets))
	for i := range e.Brackets {
		r.Equal(e.Brackets[i].Bracket, a.Brackets[i].Bracket, "Bracket")
		r.Equal(e.Brackets[i].InitialLeverage, a.Brackets[i].InitialLeverage, "InitialLeverage")
		r.Equal(e.Brackets[i].QtyCap, a.Brackets[i].QtyCap, "QtyCap")
		r.Equal(e.Brackets[i].QtyFloor, a.Brackets[i].QtyFloor, "QtyFloor")
		r.Equal(e.Brackets[i].MaintMarginRatio, a.Brackets[i].MaintMarginRatio, "MaintMarginRatio")
		r.Equal(e.Brackets[i].Cum, a.Brackets[i].Cum, "Cum")
	}
}
