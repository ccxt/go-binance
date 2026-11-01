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
 "time": 1719554596473,
 "constituents": [
  {
   "exchange": "binance",
   "symbol": "BTCUSDT"
  },
  {
   "exchange": "okex",
   "symbol": "BTC-USDT"
  },
  {
   "exchange": "coinbase",
   "symbol": "BTC-USDT"
  },
  {
   "exchange": "gateio",
   "symbol": "BTC_USDT"
  },
  {
   "exchange": "kucoin",
   "symbol": "BTC-USDT"
  },
  {
   "exchange": "mxc",
   "symbol": "BTCUSDT"
  },
  {
   "exchange": "bitget",
   "symbol": "BTCUSDT"
  },
  {
   "exchange": "bybit",
   "symbol": "BTCUSDT"
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
		Time:   1719554596473,
		Constituents: []*Constituents{{
			Exchange: "binance",
			Symbol:   "BTCUSDT"},
			{
				Exchange: "okex",
				Symbol:   "BTC-USDT"},
			{
				Exchange: "coinbase",
				Symbol:   "BTC-USDT"},
			{
				Exchange: "gateio",
				Symbol:   "BTC_USDT"},
			{
				Exchange: "kucoin",
				Symbol:   "BTC-USDT"},
			{
				Exchange: "mxc",
				Symbol:   "BTCUSDT"},
			{
				Exchange: "bitget",
				Symbol:   "BTCUSDT"},
			{
				Exchange: "bybit",
				Symbol:   "BTCUSDT"}}}
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
	}

}
