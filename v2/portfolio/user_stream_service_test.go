package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type userStreamServiceTestSuite struct {
	baseTestSuite
}

func TestUserStreamService(t *testing.T) {
	suite.Run(t, new(userStreamServiceTestSuite))
}

func (s *userStreamServiceTestSuite) TestStartUserStream() {
	data := []byte(`{
        "listenKey": "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		s.assertRequestEqual(newSignedRequest(), r)
	})

	listenKey, err := s.client.NewStartUserStreamService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1", listenKey)
}

func (s *userStreamServiceTestSuite) TestKeepaliveUserStream() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey := "dummykey"
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newSignedRequest().setFormParam("listenKey", listenKey), r)
	})

	err := s.client.NewKeepaliveUserStreamService().ListenKey(listenKey).Do(newContext())
	s.r().NoError(err)
}

func (s *userStreamServiceTestSuite) TestCloseUserStream() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey := "dummykey"
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newSignedRequest().setFormParam("listenKey", listenKey), r)
	})

	err := s.client.NewCloseUserStreamService().ListenKey(listenKey).Do(newContext())
	s.r().NoError(err)
}
