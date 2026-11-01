package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umFeeBurnServiceTestSuite struct {
	baseTestSuite
}

func TestUMFeeBurnService(t *testing.T) {
	suite.Run(t, new(umFeeBurnServiceTestSuite))
}

func (s *umFeeBurnServiceTestSuite) TestToggleFeeBurn() {
	data := []byte(`{
		"code": 200,
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	feeBurn := true
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"feeBurn": feeBurn,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMFeeBurnService().
		FeeBurn(feeBurn).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(200, res.Code)
	s.r().Equal("success", res.Msg)
}

func (s *umFeeBurnServiceTestSuite) TestToggleFeeBurnWithRecvWindow() {
	data := []byte(`{
		"code": 200,
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	feeBurn := false
	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"feeBurn":    feeBurn,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMFeeBurnService().
		FeeBurn(feeBurn).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal(200, res.Code)
	s.r().Equal("success", res.Msg)
}
