package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umPositionServiceTestSuite struct {
	baseTestSuite
}

func TestUMPositionService(t *testing.T) {
	suite.Run(t, new(umPositionServiceTestSuite))
}

func (s *umPositionServiceTestSuite) TestGetPositionRisk() {
	data := []byte(`[
		{
			"entryPrice": "0.00000",
			"leverage": "10",
			"markPrice": "6679.50671178",
			"maxNotionalValue": "20000000",
			"positionAmt": "0.000",
			"notional": "0",
			"symbol": "BTCUSDT",
			"unRealizedProfit": "0.00000000",
			"liquidationPrice": "6170.20509059",
			"positionSide": "BOTH",
			"updateTime": 1625474304765
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	positions, err := s.client.NewGetUMPositionRiskService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(positions, 1)
	s.assertPositionEqual(positions[0], &UMPositionRisk{
		EntryPrice:       "0.00000",
		Leverage:         "10",
		MarkPrice:        "6679.50671178",
		MaxNotionalValue: "20000000",
		PositionAmt:      "0.000",
		Notional:         "0",
		Symbol:           "BTCUSDT",
		UnRealizedProfit: "0.00000000",
		LiquidationPrice: "6170.20509059",
		PositionSide:     "BOTH",
		UpdateTime:       1625474304765,
	})
}

func (s *umPositionServiceTestSuite) assertPositionEqual(a, e *UMPositionRisk) {
	r := s.r()
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.MaxNotionalValue, a.MaxNotionalValue, "MaxNotionalValue")
	r.Equal(e.PositionAmt, a.PositionAmt, "PositionAmt")
	r.Equal(e.Notional, a.Notional, "Notional")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.UnRealizedProfit, a.UnRealizedProfit, "UnRealizedProfit")
	r.Equal(e.LiquidationPrice, a.LiquidationPrice, "LiquidationPrice")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
