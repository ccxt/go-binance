package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type rateLimitServiceTestSuite struct {
	baseTestSuite
}

func TestRateLimitService(t *testing.T) {
	suite.Run(t, new(rateLimitServiceTestSuite))
}

func (s *rateLimitServiceTestSuite) TestGetRateLimit() {
	data := []byte(`[
		{
			"rateLimitType": "ORDERS",
			"interval": "MINUTE",
			"intervalNum": 1,
			"limit": 1200
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	recvWindow := int64(5000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("recvWindow", recvWindow)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetRateLimitService().
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Len(res, 1)
	s.r().Equal("ORDERS", res[0].RateLimitType)
	s.r().Equal("MINUTE", res[0].Interval)
	s.r().Equal(int64(1), res[0].IntervalNum)
	s.r().Equal(int64(1200), res[0].Limit)
}
