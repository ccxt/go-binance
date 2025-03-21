package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type autoRepayFuturesStatusServiceTestSuite struct {
	baseTestSuite
}

func TestAutoRepayFuturesStatusService(t *testing.T) {
	suite.Run(t, new(autoRepayFuturesStatusServiceTestSuite))
}

func (s *autoRepayFuturesStatusServiceTestSuite) TestGetAutoRepayFuturesStatus() {
	data := []byte(`{
		"autoRepay": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	status, err := s.client.NewGetAutoRepayFuturesStatusService().Do(newContext())
	s.r().NoError(err)
	s.r().True(status.AutoRepay)
}
