package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type getMarginLoanServiceTestSuite struct {
	baseTestSuite
}

func TestGetMarginLoanService(t *testing.T) {
	suite.Run(t, new(getMarginLoanServiceTestSuite))
}

func (s *getMarginLoanServiceTestSuite) TestGetMarginLoan() {
	data := []byte(`{
		"rows": [
			{
				"txId": 12807067523,
				"asset": "BNB",
				"principal": "0.84624403",
				"timestamp": 1555056425000,
				"status": "CONFIRMED"
			}
		],
		"total": 1
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BNB"
	txID := int64(12807067523)
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		e.setParam("txId", txID)
		s.assertRequestEqual(e, r)
	})

	loans, err := s.client.NewGetMarginLoanService().Asset(asset).TxID(txID).Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1), loans.Total)
	s.r().Len(loans.Rows, 1)

	loan := loans.Rows[0]
	s.r().Equal(int64(12807067523), loan.TxID)
	s.r().Equal("BNB", loan.Asset)
	s.r().Equal("0.84624403", loan.Principal)
	s.r().Equal(int64(1555056425000), loan.Timestamp)
	s.r().Equal("CONFIRMED", loan.Status)
}
