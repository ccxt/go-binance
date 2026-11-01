package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type indexInfoServiceTestSuite struct {
	baseTestSuite
}

func TestIndexInfoService(t *testing.T) {
	suite.Run(t, new(indexInfoServiceTestSuite))
}

func (s *indexInfoServiceTestSuite) TestIndexInfo() {
	data := []byte(`[{
    "symbol": "DEFIUSDT",
    "time": 1719545343000,
    "component": "baseAsset",
    "baseAssetList": [{
            "baseAsset": "1INCH",
            "quoteAsset": "USDT",
            "weightInQuantity": "68.50746929",
            "weightInPercentage": "0.03365900"
        },
        {
            "baseAsset": "AAVE",
            "quoteAsset": "USDT",
            "weightInQuantity": "0.46687798",
            "weightInPercentage": "0.05457000"
        }
    ]
}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewIndexInfoService().Do(newContext())
	s.r().NoError(err)

	e := []*IndexInfo{
		{
			Symbol:    "DEFIUSDT",
			Time:      1719545343000,
			Component: "baseAsset",
			BaseAssetList: []*BaseAssetList{{
				BaseAsset:          "1INCH",
				QuoteAsset:         "USDT",
				WeightInQuantity:   "68.50746929",
				WeightInPercentage: "0.03365900"},
				{
					BaseAsset:          "AAVE",
					QuoteAsset:         "USDT",
					WeightInQuantity:   "0.46687798",
					WeightInPercentage: "0.05457000"}}}}

	s.assertIndexInfosEqual(e, res)
}

func (s *indexInfoServiceTestSuite) assertIndexInfoEqual(e, a *IndexInfo) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Component, a.Component, "Component")
	r.Len(e.BaseAssetList, len(a.BaseAssetList))
	for i := range e.BaseAssetList {
		r.Equal(e.BaseAssetList[i].BaseAsset, a.BaseAssetList[i].BaseAsset, "BaseAssetList[i].BaseAsset")
		r.Equal(e.BaseAssetList[i].QuoteAsset, a.BaseAssetList[i].QuoteAsset, "BaseAssetList[i].QuoteAsset")
		r.Equal(e.BaseAssetList[i].WeightInQuantity, a.BaseAssetList[i].WeightInQuantity, "BaseAssetList[i].WeightInQuantity")
		r.Equal(e.BaseAssetList[i].WeightInPercentage, a.BaseAssetList[i].WeightInPercentage, "BaseAssetList[i].WeightInPercentage")
	}

}

func (s *indexInfoServiceTestSuite) assertIndexInfosEqual(e, a []*IndexInfo) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertIndexInfoEqual(e[i], a[i])
	}
}
