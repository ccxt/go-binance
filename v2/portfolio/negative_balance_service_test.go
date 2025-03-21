package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type negativeBalanceServiceTestSuite struct {
	baseTestSuite
}

func TestNegativeBalanceService(t *testing.T) {
	suite.Run(t, new(negativeBalanceServiceTestSuite))
}

func (s *negativeBalanceServiceTestSuite) TestGetNegativeBalanceExchangeRecord() {
	data := []byte(`{
		"total": 2,
		"rows": [
			{
				"startTime": 1736263046841,
				"endTime": 1736263248179,
				"details": [
					{
						"asset": "ETH",
						"negativeBalance": 18,
						"negativeMaxThreshold": 5
					}
				]
			},
			{
				"startTime": 1736184913252,
				"endTime": 1736184965474,
				"details": [
					{
						"asset": "BNB",
						"negativeBalance": 1.10264488,
						"negativeMaxThreshold": 0
					}
				]
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	startTime := int64(1736184913252)
	endTime := int64(1736263248179)
	recvWindow := int64(5000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("startTime", startTime)
		e.setParam("endTime", endTime)
		e.setParam("recvWindow", recvWindow)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetNegativeBalanceExchangeRecordService().
		StartTime(startTime).
		EndTime(endTime).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(int64(2), res.Total)
	s.r().Len(res.Rows, 2)
	s.r().Equal(int64(1736263046841), res.Rows[0].StartTime)
	s.r().Equal("ETH", res.Rows[0].Details[0].Asset)
	s.r().Equal(18.0, res.Rows[0].Details[0].NegativeBalance)
	s.r().Equal(5.0, res.Rows[0].Details[0].NegativeMaxThreshold)
}
