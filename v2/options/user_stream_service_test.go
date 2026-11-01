package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserStreamServiceTestSuite struct {
	baseTestSuite
}

func TestUserStreamService(t *testing.T) {
	suite.Run(t, new(UserStreamServiceTestSuite))
}

func (s *UserStreamServiceTestSuite) TestStartUserStream() {
	data := []byte(`{"listenKey": "listenKeyxxzzyyaabbccdd"}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey, err := s.client.NewStartUserStreamService().Do(newContext())
	targetListenKey := "listenKeyxxzzyyaabbccdd"
	s.r().NoError(err)
	s.r().Equal(targetListenKey, listenKey)
}
