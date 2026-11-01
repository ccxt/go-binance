package portfolio_pro

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mintBFUSDServiceTestSuite struct {
	baseTestSuite
}

func TestMintBFUSDService(t *testing.T) {
	suite.Run(t, new(mintBFUSDServiceTestSuite))
}

func (s *mintBFUSDServiceTestSuite) TestMintBFUSD() {
	data := []byte(`{
		"fromAsset": "USDC",
		"targetAsset": "BFUSD",
		"fromAssetQty": "50.00000000",
		"targetAssetQty": "49.95000000",
		"mintRate": "0.99900000"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromAsset := "USDC"
	amount := "50.00000000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromAsset":   fromAsset,
			"targetAsset": "BFUSD",
			"amount":      amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMintBFUSDService().
		FromAsset(fromAsset).
		TargetAsset("BFUSD").
		Amount(amount).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(fromAsset, res.FromAsset)
	s.r().Equal("BFUSD", res.TargetAsset)
	s.r().Equal("50.00000000", res.FromAssetQty)
	s.r().Equal("49.95000000", res.TargetAssetQty)
	s.r().Equal("0.99900000", res.MintRate)
}
