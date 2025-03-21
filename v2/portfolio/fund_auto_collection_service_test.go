package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fundAutoCollectionServiceTestSuite struct {
	baseTestSuite
}

func TestFundAutoCollectionService(t *testing.T) {
	suite.Run(t, new(fundAutoCollectionServiceTestSuite))
}

func (s *fundAutoCollectionServiceTestSuite) TestFundAutoCollection() {
	data := []byte(`{
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFundAutoCollectionService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("success", res.Msg)
}
