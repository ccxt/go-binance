package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umCancelAllOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestUMCancelAllOrdersService(t *testing.T) {
	suite.Run(t, new(umCancelAllOrdersServiceTestSuite))
}

func (s *umCancelAllOrdersServiceTestSuite) TestCancelAllOrders() {
	data := []byte(`{
		"code": 200,
		"msg": "The operation of cancel all open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	recvWindow := int64(5000)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("recvWindow", recvWindow)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMCancelAllOrdersService().
		Symbol(symbol).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(200, res.Code)
	s.r().Equal("The operation of cancel all open order is done.", res.Msg)
}
