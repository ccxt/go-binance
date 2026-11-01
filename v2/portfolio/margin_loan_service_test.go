package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginLoanServiceTestSuite struct {
	baseTestSuite
}

func TestMarginLoanService(t *testing.T) {
	suite.Run(t, new(marginLoanServiceTestSuite))
}

func (s *marginLoanServiceTestSuite) TestMarginLoan() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BTC"
	amount := "1.00000000"

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		e.setParam("amount", amount)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginLoanService().
		Asset(asset).
		Amount(amount).
		Do(newContext())

	s.r().NoError(err)
	e := &MarginLoanResponse{
		TranID: 100000001,
	}
	s.assertLoanResponseEqual(e, res)
}

func (s *marginLoanServiceTestSuite) assertLoanResponseEqual(e, a *MarginLoanResponse) {
	r := s.r()
	r.Equal(e.TranID, a.TranID, "TranID")
}
