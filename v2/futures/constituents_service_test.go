package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type constituentsServiceTestSuite struct {
	baseTestSuite
}

func TestConstituentsService(t *testing.T) {
	suite.Run(t, new(constituentsServiceTestSuite))
}

func (s *constituentsServiceTestSuite) TestConstituents() {
	data := []byte(`{
 "symbol": "BTCUSDT",
 "time": 1745401553408,
 "constituents": [
  {
   "exchange": "binance",
   "symbol": "BTCUSDT",
   "price": "94057.03000000",
   "weight": "0.51282051"
  },
  {
   "exchange": "coinbase",
   "symbol": "BTC-USDT",
   "price": "94140.58000000",
   "weight": "0.15384615"
  }
 ]
}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewConstituentsService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)

	e := &ConstituentsServiceRsp{
		Symbol: "BTCUSDT",
		Time:   1745401553408,
		Constituents: []*Constituents{
			{
				Exchange: "binance",
				Symbol:   "BTCUSDT",
				Price:    "94057.03000000",
				Weight:   "0.51282051"},
			{
				Exchange: "coinbase",
				Symbol:   "BTC-USDT",
				Price:    "94140.58000000",
				Weight:   "0.15384615"}}}
	s.assertConstituentsServiceRspEqual(e, res)
}

func (s *constituentsServiceTestSuite) assertConstituentsServiceRspEqual(e, a *ConstituentsServiceRsp) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
	r.Len(e.Constituents, len(a.Constituents))
	for i := range e.Constituents {
		r.Equal(e.Constituents[i].Exchange, a.Constituents[i].Exchange, "Constituents[i].Exchange")
		r.Equal(e.Constituents[i].Symbol, a.Constituents[i].Symbol, "Constituents[i].Symbol")
		r.Equal(e.Constituents[i].Price, a.Constituents[i].Price, "Constituents[i].Price")
		r.Equal(e.Constituents[i].Weight, a.Constituents[i].Weight, "Constituents[i].Weight")
	}

}
