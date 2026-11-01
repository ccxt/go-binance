package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type uiKlinesServiceTestSuite struct {
	baseTestSuite
}

func TestUiKlinesService(t *testing.T) {
	suite.Run(t, new(uiKlinesServiceTestSuite))
}

func (s *uiKlinesServiceTestSuite) TestUiKlines() {
	data := []byte(`[
    [

        1499040000000,
        "0.01634790",
        "0.80000000",
        "0.01575800",
        "0.01577100",
        "148976.11427815",
        1499644799999,
        "2434.19055334",
        308,
        "1756.87402397",
        "28.46694368",
        "0"
    ]
]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSDT"
	var interval string = "15m"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).setParam("interval", interval)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewUiKlinesService().Symbol(symbol).Interval(interval).Do(newContext())
	s.r().NoError(err)
	s.Len(res, 1)
	e := []*UiKline{
		{
			OpenTime:                 1499040000000,
			Open:                     "0.01634790",
			High:                     "0.80000000",
			Low:                      "0.01575800",
			Close:                    "0.01577100",
			Volume:                   "148976.11427815",
			CloseTime:                1499644799999,
			QuoteVolume:              "2434.19055334",
			TradeNum:                 308,
			TakerBuyBaseAssetVolume:  "1756.87402397",
			TakerBuyQuoteAssetVolume: "28.46694368"}}

	s.assertUiKlinesEqual(e, res)
}

func (s *uiKlinesServiceTestSuite) assertUiKlineEqual(e, a *UiKline) {
	r := s.r()
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.Open, a.Open, "Open")
	r.Equal(e.High, a.High, "High")
	r.Equal(e.Low, a.Low, "Low")
	r.Equal(e.Close, a.Close, "Close")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.TradeNum, a.TradeNum, "TradeNum")
	r.Equal(e.TakerBuyBaseAssetVolume, a.TakerBuyBaseAssetVolume, "TakerBuyBaseAssetVolume")
	r.Equal(e.TakerBuyQuoteAssetVolume, a.TakerBuyQuoteAssetVolume, "TakerBuyQuoteAssetVolume")
}

func (s *uiKlinesServiceTestSuite) assertUiKlinesEqual(e, a []*UiKline) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertUiKlineEqual(e[i], a[i])
	}
}
