package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmCancelAllConditionalOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestCMCancelAllConditionalOrdersService(t *testing.T) {
	suite.Run(t, new(cmCancelAllConditionalOrdersServiceTestSuite))
}

func (s *cmCancelAllConditionalOrdersServiceTestSuite) TestCancelAllOrders() {
	data := []byte(`{
		"code": "200",
		"msg": "The operation of cancel all conditional open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelAllConditionalOrdersService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("200", res.Code)
	s.r().Equal("The operation of cancel all conditional open order is done.", res.Msg)
}

func (s *cmCancelAllConditionalOrdersServiceTestSuite) TestCancelAllOrders_WithRecvWindow() {
	data := []byte(`{
		"code": "200",
		"msg": "The operation of cancel all conditional open order is done."
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD"
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":     symbol,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMCancelAllConditionalOrdersService().
		Symbol(symbol).
		RecvWindow(recvWindow).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("200", res.Code)
}
