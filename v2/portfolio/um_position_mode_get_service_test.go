package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umPositionModeGetServiceTestSuite struct {
	baseTestSuite
}

func TestUMPositionModeGetService(t *testing.T) {
	suite.Run(t, new(umPositionModeGetServiceTestSuite))
}

func (s *umPositionModeGetServiceTestSuite) TestGetPositionMode() {
	data := []byte(`{
		"dualSidePosition": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetUMPositionModeService().Do(newContext())
	s.r().NoError(err)
	s.assertPositionModeEqual(res, &PositionMode{
		DualSidePosition: true,
	})
}

func (s *umPositionModeGetServiceTestSuite) assertPositionModeEqual(a, e *PositionMode) {
	r := s.r()
	r.Equal(e.DualSidePosition, a.DualSidePosition, "DualSidePosition")
}
