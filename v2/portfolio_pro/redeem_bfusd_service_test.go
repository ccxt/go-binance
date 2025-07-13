package portfolio_pro

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type redeemBFUSDServiceTestSuite struct {
	baseTestSuite
}

func TestRedeemBFUSDService(t *testing.T) {
	suite.Run(t, new(redeemBFUSDServiceTestSuite))
}

func (s *redeemBFUSDServiceTestSuite) TestRedeemBFUSD() {
	data := []byte(`{
		"fromAsset": "BFUSD",
		"targetAsset": "USDC",
		"fromAssetQty": "50.00000000",
		"targetAssetQty": "49.95000000",
		"redeemRate": "0.99950000"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	targetAsset := "USDC"
	amount := "50.00000000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromAsset":   "BFUSD",
			"targetAsset": targetAsset,
			"amount":      amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewRedeemBFUSDService().
		FromAsset("BFUSD").
		TargetAsset(targetAsset).
		Amount(amount).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("BFUSD", res.FromAsset)
	s.r().Equal(targetAsset, res.TargetAsset)
	s.r().Equal("49.95000000", res.TargetAssetQty)
	s.r().Equal("0.99950000", res.RedeemRate)
}
