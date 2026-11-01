package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fundCollectionByAssetServiceTestSuite struct {
	baseTestSuite
}

func TestFundCollectionByAssetService(t *testing.T) {
	suite.Run(t, new(fundCollectionByAssetServiceTestSuite))
}

func (s *fundCollectionByAssetServiceTestSuite) TestFundCollectionByAsset() {
	data := []byte(`{
		"msg": "success"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "USDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewFundCollectionByAssetService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("success", res.Msg)
}
