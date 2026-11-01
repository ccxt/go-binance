package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginRepayDebtServiceTestSuite struct {
	baseTestSuite
}

func TestMarginRepayDebtService(t *testing.T) {
	suite.Run(t, new(marginRepayDebtServiceTestSuite))
}

func (s *marginRepayDebtServiceTestSuite) TestMarginRepayDebt() {
	data := []byte(`{
		"amount": "0.10000000",
		"asset": "BNB",
		"specifyRepayAssets": ["USDT", "BTC"],
		"updateTime": 1636371437000,
		"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BNB"
	amount := "0.10000000"
	specifyRepayAssets := "USDT,BTC"
	recvWindow := int64(1000)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":              asset,
			"amount":             amount,
			"specifyRepayAssets": specifyRepayAssets,
			"recvWindow":         recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginRepayDebtService().
		Asset(asset).
		Amount(amount).
		SpecifyRepayAssets(specifyRepayAssets).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal("0.10000000", res.Amount)
	s.r().Equal("BNB", res.Asset)
	s.r().Equal([]string{"USDT", "BTC"}, res.SpecifyRepayAssets)
	s.r().Equal(int64(1636371437000), res.UpdateTime)
	s.r().True(res.Success)
}

func (s *marginRepayDebtServiceTestSuite) TestMarginRepayDebtWithoutAmount() {
	data := []byte(`{
		"amount": "0.10000000",
		"asset": "BNB",
		"specifyRepayAssets": ["USDT", "BTC"],
		"updateTime": 1636371437000,
		"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BNB"
	specifyRepayAssets := "USDT,BTC"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":              asset,
			"specifyRepayAssets": specifyRepayAssets,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginRepayDebtService().
		Asset(asset).
		SpecifyRepayAssets(specifyRepayAssets).
		Do(newContext())

	s.r().NoError(err)
	s.r().Equal("0.10000000", res.Amount)
	s.r().Equal("BNB", res.Asset)
	s.r().True(res.Success)
}
