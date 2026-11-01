package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type feeburnServiceTestSuite struct {
	baseTestSuite
}

func TestFeeBurnService(t *testing.T) {
	suite.Run(t, new(feeburnServiceTestSuite))
}

func (s *feeburnServiceTestSuite) TestGetFeeBurn() {
	data := []byte(`{"feeBurn": true}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetFeeBurnService().Do(newContext())
	s.r().NoError(err)
	s.r().True(res.FeeBurn)
}

func (s *feeburnServiceTestSuite) TestFeeBurnEnable() {
	data := []byte(`{"msg": "success"}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("feeBurn", "true")
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewFeeBurnService().Enable().Do(newContext())
	s.r().NoError(err)
}

func (s *feeburnServiceTestSuite) TestFeeBurnDisable() {
	data := []byte(`{"msg": "success"}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParam("feeBurn", "false")
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewFeeBurnService().Disable().Do(newContext())
	s.r().NoError(err)
}
