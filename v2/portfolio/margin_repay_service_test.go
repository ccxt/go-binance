package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginRepayServiceTestSuite struct {
	baseTestSuite
}

func TestMarginRepayService(t *testing.T) {
	suite.Run(t, new(marginRepayServiceTestSuite))
}

func (s *marginRepayServiceTestSuite) TestMarginRepay() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BTC"
	amount := "1.0"
	recvWindow := int64(5000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		e.setParam("amount", amount)
		e.setParam("recvWindow", recvWindow)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginRepayService().
		Asset(asset).
		Amount(amount).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(int64(100000001), res.TranID)
}
