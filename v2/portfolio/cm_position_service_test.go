package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmPositionServiceTestSuite struct {
	baseTestSuite
}

func TestCMPositionService(t *testing.T) {
	suite.Run(t, new(cmPositionServiceTestSuite))
}

func (s *cmPositionServiceTestSuite) TestGetPositionRisk() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_201225",
			"positionAmt": "1",
			"entryPrice": "11707.70000003",
			"markPrice": "11788.66626667",
			"unrealizedProfit": "0.00005866",
			"liquidationPrice": "6170.20509059",
			"leverage": "125",
			"positionSide": "LONG",
			"updateTime": 1627026881327,
			"maxQty": "50",
			"notionalValue": "0.00084827"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	marginAsset := "BTC"
	pair := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("marginAsset", marginAsset)
		e.setParam("pair", pair)
		s.assertRequestEqual(e, r)
	})

	positions, err := s.client.NewGetCMPositionRiskService().
		MarginAsset(marginAsset).
		Pair(pair).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(positions, 1)
	s.assertPositionEqual(positions[0], &CMPosition{
		Symbol:           "BTCUSD_201225",
		PositionAmt:      "1",
		EntryPrice:       "11707.70000003",
		MarkPrice:        "11788.66626667",
		UnrealizedProfit: "0.00005866",
		LiquidationPrice: "6170.20509059",
		Leverage:         "125",
		PositionSide:     "LONG",
		UpdateTime:       1627026881327,
		MaxQty:           "50",
		NotionalValue:    "0.00084827",
	})
}

func (s *cmPositionServiceTestSuite) assertPositionEqual(a, e *CMPosition) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PositionAmt, a.PositionAmt, "PositionAmt")
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.UnrealizedProfit, a.UnrealizedProfit, "UnrealizedProfit")
	r.Equal(e.LiquidationPrice, a.LiquidationPrice, "LiquidationPrice")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.MaxQty, a.MaxQty, "MaxQty")
	r.Equal(e.NotionalValue, a.NotionalValue, "NotionalValue")
}
