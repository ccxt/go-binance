package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginWithdrawServiceTestSuite struct {
	baseTestSuite
}

func TestMarginWithdrawService(t *testing.T) {
	suite.Run(t, new(marginWithdrawServiceTestSuite))
}

func (s *marginWithdrawServiceTestSuite) TestGetMaxWithdraw() {
	data := []byte(`{
		"amount": "60"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "USDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetMarginMaxWithdrawService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	s.assertMaxWithdrawEqual(res, &MaxWithdraw{
		Amount: "60",
	})
}

func (s *marginWithdrawServiceTestSuite) assertMaxWithdrawEqual(a, e *MaxWithdraw) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, "Amount")
}
