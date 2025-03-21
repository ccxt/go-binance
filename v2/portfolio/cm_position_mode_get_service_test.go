package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmPositionModeGetServiceTestSuite struct {
	baseTestSuite
}

func TestCMPositionModeGetService(t *testing.T) {
	suite.Run(t, new(cmPositionModeGetServiceTestSuite))
}

func (s *cmPositionModeGetServiceTestSuite) TestGetPositionMode() {
	data := []byte(`{
		"dualSidePosition": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetCMPositionModeService().Do(newContext())
	s.r().NoError(err)
	s.assertPositionModeEqual(res, &PositionMode{
		DualSidePosition: true,
	})
}

func (s *cmPositionModeGetServiceTestSuite) assertPositionModeEqual(a, e *PositionMode) {
	r := s.r()
	r.Equal(e.DualSidePosition, a.DualSidePosition, "DualSidePosition")
}
