package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type lvtKlinesServiceTestSuite struct {
	baseTestSuite
}

func TestLvtKlinesService(t *testing.T) {
	suite.Run(t, new(lvtKlinesServiceTestSuite))
}

func (s *lvtKlinesServiceTestSuite) TestLvtKlines() {
	data := []byte(`[
 [
  1719108000000,
  "0.00084008",
  "0.00084008",
  "0.00084008",
  "0.00084008",
  "1.64657746",
  1719108899999,
  "0",
  0,
  "0",
  "0",
  "0"
 ],
 [
  1719108900000,
  "0.00084008",
  "0.00084008",
  "0.00084008",
  "0.00084008",
  "1.64657746",
  1719109799999,
  "0",
  0,
  "0",
  "0",
  "0"
 ]]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCDOWN"
	var interval string = "15m"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).setParam("interval", interval)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewLvtKlinesService().Symbol(symbol).Interval(interval).Do(newContext())
	s.r().NoError(err)
	s.Len(res, 2)
	e := []*LvtKline{
		{
			OpenTime:      1719108000000,
			Open:          "0.00084008",
			High:          "0.00084008",
			Low:           "0.00084008",
			Close:         "0.00084008",
			CloseLeverage: "1.64657746",
			CloseTime:     1719108899999},
		{
			OpenTime:      1719108900000,
			Open:          "0.00084008",
			High:          "0.00084008",
			Low:           "0.00084008",
			Close:         "0.00084008",
			CloseLeverage: "1.64657746",
			CloseTime:     1719109799999}}

	s.assertLvtKlinesEqual(e, res)
}

func (s *lvtKlinesServiceTestSuite) assertLvtKlineEqual(e, a *LvtKline) {
	r := s.r()
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.Open, a.Open, "Open")
	r.Equal(e.High, a.High, "High")
	r.Equal(e.Low, a.Low, "Low")
	r.Equal(e.Close, a.Close, "Close")
	r.Equal(e.CloseLeverage, a.CloseLeverage, "CloseLeverage")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
}

func (s *lvtKlinesServiceTestSuite) assertLvtKlinesEqual(e, a []*LvtKline) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertLvtKlineEqual(e[i], a[i])
	}
}
