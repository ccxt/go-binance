package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type apiTradingStatusServiceTestSuite struct {
	baseTestSuite
}

func TestApiTradingStatusService(t *testing.T) {
	suite.Run(t, new(apiTradingStatusServiceTestSuite))
}

func (s *apiTradingStatusServiceTestSuite) TestGetApiTradingStatus() {
	data := []byte(`{
		"indicators": {
			"BTCUSDT": [
				{
					"isLocked": true,
					"plannedRecoverTime": 1545741270000,
					"indicator": "UFR",
					"value": 0.05,
					"triggerValue": 0.995
				},
				{
					"isLocked": true,
					"plannedRecoverTime": 1545741270000,
					"indicator": "IFER",
					"value": 0.99,
					"triggerValue": 0.99
				},
				{
					"isLocked": true,
					"plannedRecoverTime": 1545741270000,
					"indicator": "GCR",
					"value": 0.99,
					"triggerValue": 0.99
				},
				{
					"isLocked": true,
					"plannedRecoverTime": 1545741270000,
					"indicator": "DR",
					"value": 0.99,
					"triggerValue": 0.99
				}
			]
		},
		"updateTime": 1545741270000
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewApiTradingStatusService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)

	e := &TradingStatusIndicators{
		Indicators: map[string][]*IndicatorInfo{
			"BTCUSDT": {
				{
					IsLocked:           true,
					PlannedRecoverTime: 1545741270000,
					Indicator:          "UFR",
					Value:              0.05,
					TriggerValue:       0.995,
				},
				{
					IsLocked:           true,
					PlannedRecoverTime: 1545741270000,
					Indicator:          "IFER",
					Value:              0.99,
					TriggerValue:       0.99,
				},
				{
					IsLocked:           true,
					PlannedRecoverTime: 1545741270000,
					Indicator:          "GCR",
					Value:              0.99,
					TriggerValue:       0.99,
				},
				{
					IsLocked:           true,
					PlannedRecoverTime: 1545741270000,
					Indicator:          "DR",
					Value:              0.99,
					TriggerValue:       0.99,
				},
			},
		},
		UpdateTime: 1545741270000,
	}

	s.assertApiTradingStatusEqual(e, res)
}

func (s *apiTradingStatusServiceTestSuite) assertApiTradingStatusEqual(e, a *TradingStatusIndicators) {
	r := s.r()
	s.r().Len(e.Indicators, len(a.Indicators))
	for k := range e.Indicators {
		s.r().Len(e.Indicators[k], len(a.Indicators[k]))

		for i := range e.Indicators[k] {
			r.Equal(e.Indicators[k][i].IsLocked, a.Indicators[k][i].IsLocked, "IsLocked")
			r.Equal(e.Indicators[k][i].PlannedRecoverTime, a.Indicators[k][i].PlannedRecoverTime, "PlannedRecoverTime")
			r.Equal(e.Indicators[k][i].Indicator, a.Indicators[k][i].Indicator, "Indicator")
			r.Equal(e.Indicators[k][i].Value, a.Indicators[k][i].Value, "Value")
			r.Equal(e.Indicators[k][i].TriggerValue, a.Indicators[k][i].TriggerValue, "TriggerValue")
		}
	}
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
}
