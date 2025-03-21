package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umTradingStatusServiceTestSuite struct {
	baseTestSuite
}

func TestUMTradingStatusService(t *testing.T) {
	suite.Run(t, new(umTradingStatusServiceTestSuite))
}

func (s *umTradingStatusServiceTestSuite) TestGetTradingStatus() {
	data := []byte(`{
		"indicators": {
			"BTCUSDT": [
				{
					"isLocked": true,
					"plannedRecoverTime": 1545741270000,
					"indicator": "UFR",
					"value": 0.05,
					"triggerValue": 0.995
				}
			]
		},
		"updateTime": 1545741270000
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	status, err := s.client.NewGetUMTradingStatusService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1545741270000), status.UpdateTime)
	s.r().Len(status.Indicators["BTCUSDT"], 1)

	indicator := status.Indicators["BTCUSDT"][0]
	s.r().True(indicator.IsLocked)
	s.r().Equal(int64(1545741270000), indicator.PlannedRecoverTime)
	s.r().Equal("UFR", indicator.Indicator)
	s.r().Equal(0.05, indicator.Value)
	s.r().Equal(0.995, indicator.TriggerValue)
}

func (s *umTradingStatusServiceTestSuite) TestGetAccountViolation() {
	data := []byte(`{
		"indicators": {
			"ACCOUNT": [
				{
					"indicator": "TMV",
					"value": 10,
					"triggerValue": 1,
					"plannedRecoverTime": 1644919865000,
					"isLocked": true
				}
			]
		},
		"updateTime": 1644913304748
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	status, err := s.client.NewGetUMTradingStatusService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1644913304748), status.UpdateTime)
	s.r().Len(status.Indicators["ACCOUNT"], 1)

	indicator := status.Indicators["ACCOUNT"][0]
	s.r().True(indicator.IsLocked)
	s.r().Equal(int64(1644919865000), indicator.PlannedRecoverTime)
	s.r().Equal("TMV", indicator.Indicator)
	s.r().Equal(float64(10), indicator.Value)
	s.r().Equal(float64(1), indicator.TriggerValue)
}
