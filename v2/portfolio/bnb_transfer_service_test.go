package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type bnbTransferServiceTestSuite struct {
	baseTestSuite
}

func TestBNBTransferService(t *testing.T) {
	suite.Run(t, new(bnbTransferServiceTestSuite))
}

func (s *bnbTransferServiceTestSuite) TestBNBTransfer() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	amount := "1.0"
	transferSide := TransferSideToUM
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("amount", amount)
		e.setParam("transferSide", transferSide)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewBNBTransferService().
		Amount(amount).
		TransferSide(transferSide).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(int64(100000001), res.TranID)
}
