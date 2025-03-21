package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginBorrowServiceTestSuite struct {
	baseTestSuite
}

func TestMarginBorrowService(t *testing.T) {
	suite.Run(t, new(marginBorrowServiceTestSuite))
}

func (s *marginBorrowServiceTestSuite) TestGetMaxBorrow() {
	data := []byte(`{
		"amount": "125",
		"borrowLimit": "60"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "USDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetMarginMaxBorrowService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	s.assertMaxBorrowEqual(res, &MaxBorrow{
		Amount:      "125",
		BorrowLimit: "60",
	})
}

func (s *marginBorrowServiceTestSuite) assertMaxBorrowEqual(a, e *MaxBorrow) {
	r := s.r()
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.BorrowLimit, a.BorrowLimit, "BorrowLimit")
}
