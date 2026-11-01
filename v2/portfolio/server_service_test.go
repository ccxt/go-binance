package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type serverServiceTestSuite struct {
	baseTestSuite
}

func TestServerService(t *testing.T) {
	suite.Run(t, new(serverServiceTestSuite))
}

func (s *serverServiceTestSuite) TestPing() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewPingService().Do(newContext())
	s.r().NoError(err)
}
