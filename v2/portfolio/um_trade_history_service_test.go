package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umTradeHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestUMTradeHistoryService(t *testing.T) {
	suite.Run(t, new(umTradeHistoryServiceTestSuite))
}

func (s *umTradeHistoryServiceTestSuite) TestGetUMTradeHistoryDownloadID() {
	data := []byte(`{
		"avgCostTimestampOfLast30d": 7241837,
		"downloadId": "546975389218332672"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	startTime := int64(1622555222000)
	endTime := int64(1622555522000)
	recvWindow := int64(5000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("startTime", startTime)
		e.setParam("endTime", endTime)
		e.setParam("recvWindow", recvWindow)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetUMTradeHistoryDownloadIDService().
		StartTime(startTime).
		EndTime(endTime).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(int64(7241837), res.AvgCostTimestampOfLast30d)
	s.r().Equal("546975389218332672", res.DownloadID)
}
