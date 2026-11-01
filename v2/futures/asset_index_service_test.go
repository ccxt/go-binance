package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetIndexServiceTestSuite struct {
	baseTestSuite
}

func TestAssetIndexService(t *testing.T) {
	suite.Run(t, new(assetIndexServiceTestSuite))
}

func (s *assetIndexServiceTestSuite) TestAssetIndex() {
	data := []byte(`[
 {
  "symbol": "BTCUSD",
  "time": 1719554624000,
  "index": "61475.02720330",
  "bidBuffer": "0.05000000",
  "askBuffer": "0.05000000",
  "bidRate": "58401.27584314",
  "askRate": "64548.77856347",
  "autoExchangeBidBuffer": "0.02500000",
  "autoExchangeAskBuffer": "0.02500000",
  "autoExchangeBidRate": "59938.15152322",
  "autoExchangeAskRate": "63011.90288338"
 },
 {
  "symbol": "USDTUSD",
  "time": 1719554624000,
  "index": "0.99884101",
  "bidBuffer": "0.00010000",
  "askBuffer": "0.00010000",
  "bidRate": "0.99874113",
  "askRate": "0.99894089",
  "autoExchangeBidBuffer": "0.00010000",
  "autoExchangeAskBuffer": "0.00010000",
  "autoExchangeBidRate": "0.99874113",
  "autoExchangeAskRate": "0.99894089"
 }]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSD"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewAssetIndexService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)

	e := []*AssetIndex{
		{
			Symbol:                "BTCUSD",
			Time:                  1719554624000,
			Index:                 "61475.02720330",
			BidBuffer:             "0.05000000",
			AskBuffer:             "0.05000000",
			BidRate:               "58401.27584314",
			AskRate:               "64548.77856347",
			AutoExchangeBidBuffer: "0.02500000",
			AutoExchangeAskBuffer: "0.02500000",
			AutoExchangeBidRate:   "59938.15152322",
			AutoExchangeAskRate:   "63011.90288338"},
		{
			Symbol:                "USDTUSD",
			Time:                  1719554624000,
			Index:                 "0.99884101",
			BidBuffer:             "0.00010000",
			AskBuffer:             "0.00010000",
			BidRate:               "0.99874113",
			AskRate:               "0.99894089",
			AutoExchangeBidBuffer: "0.00010000",
			AutoExchangeAskBuffer: "0.00010000",
			AutoExchangeBidRate:   "0.99874113",
			AutoExchangeAskRate:   "0.99894089"}}

	s.assertAssetIndexesEqual(e, res)
}

func (s *assetIndexServiceTestSuite) assertAssetIndexEqual(e, a *AssetIndex) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Index, a.Index, "Index")
	r.Equal(e.BidBuffer, a.BidBuffer, "BidBuffer")
	r.Equal(e.AskBuffer, a.AskBuffer, "AskBuffer")
	r.Equal(e.BidRate, a.BidRate, "BidRate")
	r.Equal(e.AskRate, a.AskRate, "AskRate")
	r.Equal(e.AutoExchangeBidBuffer, a.AutoExchangeBidBuffer, "AutoExchangeBidBuffer")
	r.Equal(e.AutoExchangeAskBuffer, a.AutoExchangeAskBuffer, "AutoExchangeAskBuffer")
	r.Equal(e.AutoExchangeBidRate, a.AutoExchangeBidRate, "AutoExchangeBidRate")
	r.Equal(e.AutoExchangeAskRate, a.AutoExchangeAskRate, "AutoExchangeAskRate")
}

func (s *assetIndexServiceTestSuite) assertAssetIndexesEqual(e, a []*AssetIndex) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertAssetIndexEqual(e[i], a[i])
	}
}
