package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServerServiceTestSuite struct {
	baseTestSuite
}

func TestPingService(t *testing.T) {
	suite.Run(t, new(ServerServiceTestSuite))
}

func (s *ServerServiceTestSuite) TestPing() {

	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	err := s.client.NewPingService().Do(newContext())
	s.r().Equal(err, nil, "err != nil")
}

func (s *ServerServiceTestSuite) TestServerTime() {

	data := []byte(`{
		"serverTime": 1592387156596
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	st, err := s.client.NewServerTimeService().Do(newContext())
	var targetServerTime int64 = 1592387156596
	s.r().Equal(st, targetServerTime, "serverTime")
	s.r().Equal(err, nil, "err != nil")
}
