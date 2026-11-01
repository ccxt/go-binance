package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umFeeBurnStatusServiceTestSuite struct {
	baseTestSuite
}

func TestUMFeeBurnStatusService(t *testing.T) {
	suite.Run(t, new(umFeeBurnStatusServiceTestSuite))
}

func (s *umFeeBurnStatusServiceTestSuite) TestGetFeeBurnStatus() {
	data := []byte(`{
		"feeBurn": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMFeeBurnStatusService().
		Do(newContext())

	s.r().NoError(err)
	s.r().True(res.FeeBurn)
}

func (s *umFeeBurnStatusServiceTestSuite) TestGetFeeBurnStatusWithRecvWindow() {
	data := []byte(`{
		"feeBurn": false
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMFeeBurnStatusService().
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().False(res.FeeBurn)
}
