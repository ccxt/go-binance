package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umPositionModeServiceTestSuite struct {
	baseTestSuite
}

func TestUMPositionModeService(t *testing.T) {
	suite.Run(t, new(umPositionModeServiceTestSuite))
}

func (s *umPositionModeServiceTestSuite) TestChangePositionMode() {
	data := []byte(`{
		"code": 200,
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	dualSidePosition := true
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("dualSidePosition", dualSidePosition)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewChangeUMPositionModeService().
		DualSidePosition(dualSidePosition).
		Do(newContext())
	s.r().NoError(err)
	s.assertPositionModeResponseEqual(res, &APIResponse{
		Code: 200,
		Msg:  "success",
	})
}

func (s *umPositionModeServiceTestSuite) assertPositionModeResponseEqual(a, e *APIResponse) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Msg, a.Msg, "Msg")
}
