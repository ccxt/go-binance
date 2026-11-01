package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type repayFuturesNegativeBalanceServiceTestSuite struct {
	baseTestSuite
}

func TestRepayFuturesNegativeBalanceService(t *testing.T) {
	suite.Run(t, new(repayFuturesNegativeBalanceServiceTestSuite))
}

func (s *repayFuturesNegativeBalanceServiceTestSuite) TestRepayFuturesNegativeBalance() {
	data := []byte(`{
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewRepayFuturesNegativeBalanceService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("success", res.Msg)
}
