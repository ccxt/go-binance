package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmCancelAllOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMCancelAllOrdersService(t *testing.T) {
	suite.Run(t, new(cmCancelAllOrdersServiceTestSuite))
}

func (s *cmCancelAllOrdersServiceTestSuite) TestCancelAllOrders() {
	data := []byte(`{
		"code": 200,
		"msg": "The operation of cancel all open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelAllOrdersService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(200, res.Code)
	s.r().Equal("The operation of cancel all open order is done.", res.Msg)
}

func (s *cmCancelAllOrdersServiceTestSuite) TestCancelAllOrders_WithRecvWindow() {
	data := []byte(`{
		"code": 200,
		"msg": "The operation of cancel all open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelAllOrdersService().
		Symbol(symbol).
		RecvWindow(recvWindow).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(200, res.Code)
}
