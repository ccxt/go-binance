package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type autoRepayFuturesSwitchServiceTestSuite struct {
	baseTestSuite
}

func TestAutoRepayFuturesSwitchService(t *testing.T) {
	suite.Run(t, new(autoRepayFuturesSwitchServiceTestSuite))
}

func (s *autoRepayFuturesSwitchServiceTestSuite) TestChangeAutoRepayFuturesStatus() {
	data := []byte(`{
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	autoRepay := true
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("autoRepay", autoRepay)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewChangeAutoRepayFuturesStatusService().AutoRepay(autoRepay).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("success", res.Msg)
}
