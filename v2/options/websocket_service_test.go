package options

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	origWsServe func(*WsConfig, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketServiceTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketServiceTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, innerErr error) {
		s.serveCount++
		doneC = make(chan struct{})
		stopC = make(chan struct{})
		go func() {
			<-stopC
			close(doneC)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneC, stopC, nil
	}
}

func (s *websocketServiceTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketServiceTestSuite) assertWsTradeEvent(e, a *WsTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TradeId, a.TradeId, "TradeId")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.BuyId, a.BuyId, "BuyId")
	r.Equal(e.SellId, a.SellId, "SellId")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.Side, a.Side, "Side")
}

func (s *websocketServiceTestSuite) TestTradeServe() {
	data := []byte(`{
		"e": "trade",
		"E": 1716883280754,
		"s": "ETH-240531-3600-P",
		"t": "381",
		"p": "15.5",
		"q": "1.25",
		"b": "4674957007717330944",
		"a": "4692971406119415808",
		"T": 1716883280751,
		"S": "1"
	   }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsTradeServe("ETH", func(event *WsTradeEvent) {
		e := &WsTradeEvent{
			Event:     "trade",
			Time:      1716883280754,
			Symbol:    "ETH-240531-3600-P",
			TradeId:   "381",
			Price:     "15.5",
			Quantity:  "1.25",
			BuyId:     "4674957007717330944",
			SellId:    "4692971406119415808",
			TradeTime: 1716883280751,
			Side:      "1",
		}
		s.assertWsTradeEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsIndexEvent(e, a *WsIndexEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
}

func (s *websocketServiceTestSuite) TestIndexServe() {
	data := []byte(`{
		"e": "index",
		"E": 1716883243048,
		"s": "ETHUSDT",
		"p": "3846.63204545"
	   }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsIndexServe("ETHUSDT", func(event *WsIndexEvent) {
		e := &WsIndexEvent{
			Event:  "index",
			Time:   1716883243048,
			Symbol: "ETHUSDT",
			Price:  "3846.63204545",
		}
		s.assertWsIndexEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarkPriceEvent(ee, a []*WsMarkPriceEvent) {
	r := s.r()
	for i, e := range ee {
		r.Equal(e.Event, a[i].Event, "Event")
		r.Equal(e.Time, a[i].Time, "Time")
		r.Equal(e.Symbol, a[i].Symbol, "Symbol")
		r.Equal(e.MarkPrice, a[i].MarkPrice, "MarkPrice")
	}
}

func (s *websocketServiceTestSuite) TestMarkPriceServe() {
	data := []byte(`[
		{
		 "e": "markPrice",
		 "E": 1716884520102,
		 "s": "ETH-240628-800-C",
		 "mp": "3066.5"
		},
		{
		 "e": "markPrice",
		 "E": 1716884520102,
		 "s": "ETH-240628-800-P",
		 "mp": "0.1"
		},
		{
		 "e": "markPrice",
		 "E": 1716884520102,
		 "s": "ETH-240628-1200-C",
		 "mp": "2669.9"
		}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarkPriceServe("ETH", func(event []*WsMarkPriceEvent) {
		e := []*WsMarkPriceEvent{
			{
				Event:     "markPrice",
				Time:      1716884520102,
				Symbol:    "ETH-240628-800-C",
				MarkPrice: "3066.5",
			},
			{
				Event:     "markPrice",
				Time:      1716884520102,
				Symbol:    "ETH-240628-800-P",
				MarkPrice: "0.1",
			},
			{
				Event:     "markPrice",
				Time:      1716884520102,
				Symbol:    "ETH-240628-1200-C",
				MarkPrice: "2669.9",
			},
		}
		s.assertWsMarkPriceEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsKlineEvent(e, a *WsKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Kline.StartTime, a.Kline.StartTime, "Kline.StartTime")
	r.Equal(e.Kline.EndTime, a.Kline.EndTime, "Kline.EndTime")
	r.Equal(e.Kline.Symbol, a.Kline.Symbol, "Kline.Symbol")
	r.Equal(e.Kline.Interval, a.Kline.Interval, "Kline.Interval")
	r.Equal(e.Kline.FirstTradeID, a.Kline.FirstTradeID, "Kline.FirstTradeID")
	r.Equal(e.Kline.LastTradeID, a.Kline.LastTradeID, "Kline.LastTradeID")
	r.Equal(e.Kline.Open, a.Kline.Open, "Kline.Open")
	r.Equal(e.Kline.High, a.Kline.High, "Kline.High")
	r.Equal(e.Kline.Low, a.Kline.Low, "Kline.Low")
	r.Equal(e.Kline.Close, a.Kline.Close, "Kline.Close")
	r.Equal(e.Kline.Volume, a.Kline.Volume, "Kline.Volume")
	r.Equal(e.Kline.TradeNum, a.Kline.TradeNum, "Kline.TradeNum")
	r.Equal(e.Kline.IsFinal, a.Kline.IsFinal, "Kline.IsFinal")
	r.Equal(e.Kline.QuoteVolume, a.Kline.QuoteVolume, "Kline.QuoteVolume")
	r.Equal(e.Kline.TakerQuoteVolume, a.Kline.TakerQuoteVolume, "Kline.TakerQuoteVolume")
	r.Equal(e.Kline.TakerVolume, a.Kline.TakerVolume, "Kline.TakerVolume")
}

func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
		"e": "kline",
		"E": 1716885720086,
		"s": "ETH-240628-800-C",
		"k": {
		 "t": 1716885720000,
		 "T": 1716885780000,
		 "s": "ETH-240628-800-C",
		 "i": "1m",
		 "F": "0",
		 "L": "-1",
		 "o": "3010.6",
		 "h": "3010.6",
		 "l": "3010.6",
		 "c": "3010.6",
		 "v": "0",
		 "n": 0,
		 "q": "0",
		 "x": false,
		 "V": "0",
		 "Q": "0"
		}
	   }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsKlineServe("ETH-240628-800-C", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1716885720086,
			Symbol: "ETH-240628-800-C",
			Kline: &WsKline{
				StartTime:        1716885720000,
				EndTime:          1716885780000,
				Symbol:           "ETH-240628-800-C",
				Interval:         "1m",
				FirstTradeID:     "0",
				LastTradeID:      "-1",
				Open:             "3010.6",
				High:             "3010.6",
				Low:              "3010.6",
				Close:            "3010.6",
				Volume:           "0",
				TradeNum:         0,
				QuoteVolume:      "0",
				IsFinal:          false,
				TakerQuoteVolume: "0",
				TakerVolume:      "0",
			},
		}
		s.assertWsKlineEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsTickerEvent(ee, aa []*WsTickerEvent) {
	r := s.r()
	for i, e := range ee {
		a := aa[i]
		r.Equal(e.Event, a.Event, "Event")
		r.Equal(e.Time, a.Time, "Time")
		r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
		r.Equal(e.Symbol, a.Symbol, "Symbol")
		r.Equal(e.Open, a.Open, "Open")
		r.Equal(e.High, a.High, "High")
		r.Equal(e.Low, a.Low, "Low")
		r.Equal(e.Close, a.Close, "Close")
		r.Equal(e.Volume, a.Volume, "Volume")
		r.Equal(e.Amount, a.Amount, "Amount")
		r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
		r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
		r.Equal(e.LastQty, a.LastQty, "LastQty")
		r.Equal(e.FirstID, a.FirstID, "FirstID")
		r.Equal(e.LastID, a.LastID, "LastID")
		r.Equal(e.TradeNum, a.TradeNum, "TradeNum")
		r.Equal(e.AskOpenPrice, a.AskOpenPrice, "AskOpenPrice")
		r.Equal(e.BidOpenPrice, a.BidOpenPrice, "BidOpenPrice")
		r.Equal(e.AskQty, a.AskQty, "AskQty")
		r.Equal(e.BidQty, a.BidQty, "BidQty")
		r.Equal(e.AskIV, a.AskIV, "AskIV")
		r.Equal(e.BidIV, a.BidIV, "BidIV")
		r.Equal(e.Delta, a.Delta, "Delta")
		r.Equal(e.Theta, a.Theta, "Theta")
		r.Equal(e.Gamma, a.Gamma, "Gamma")
		r.Equal(e.Vega, a.Vega, "Vega")
		r.Equal(e.Volatility, a.Volatility, "Volatility")
		r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
		r.Equal(e.HighPriceLimit, a.HighPriceLimit, "HighPriceLimit")
		r.Equal(e.LowPriceLimit, a.LowPriceLimit, "LowPriceLimit")
		r.Equal(e.ExercisPrice, a.ExercisPrice, "ExercisPrice")
	}
}

func (s *websocketServiceTestSuite) TestTickerServe() {
	data := []byte(`{
		"e": "24hrTicker",
		"E": 1716887445078,
		"T": 1716887445000,
		"s": "ETH-240628-800-C",
		"o": "3010.6",
		"h": "3010.6",
		"l": "3010.6",
		"c": "3010.6",
		"V": "0",
		"A": "0",
		"P": "0",
		"p": "0",
		"Q": "0",
		"F": "0",
		"L": "0",
		"n": 0,
		"bo": "60.9",
		"ao": "0",
		"bq": "0.01",
		"aq": "0",
		"b": "0.0000006",
		"a": "-0.00000001",
		"d": "1",
		"t": "0",
		"g": "0",
		"v": "0",
		"vo": "0.575",
		"mp": "3083.7",
		"hl": "3374.4",
		"ll": "2793",
		"eep": "0",
		"r": "0.1"
	   }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsTickerServe("ETH-240628-800-C", func(event []*WsTickerEvent) {
		e := []*WsTickerEvent{
			{
				Event:              "24hrTicker",
				Time:               1716887445078,
				TradeTime:          1716887445000,
				Symbol:             "ETH-240628-800-C",
				Open:               "3010.6",
				High:               "3010.6",
				Low:                "3010.6",
				Close:              "3010.6",
				Volume:             "0",
				Amount:             "0",
				PriceChangePercent: "0",
				PriceChange:        "0",
				LastQty:            "0",
				FirstID:            "0",
				LastID:             "0",
				TradeNum:           0,
				BidOpenPrice:       "60.9",
				AskOpenPrice:       "0",
				BidQty:             "0.01",
				AskQty:             "0",
				BidIV:              "0.0000006",
				AskIV:              "-0.00000001",
				Delta:              "1",
				Theta:              "0",
				Gamma:              "0",
				Vega:               "0",
				Volatility:         "0.575",
				MarkPrice:          "3083.7",
				HighPriceLimit:     "3374.4",
				LowPriceLimit:      "2793",
				ExercisPrice:       "0",
				RiskFreeInterest:   "0.1",
			},
		}
		s.assertWsTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestTickerWithExpireServe() {
	data := []byte(`[
		{
		 "e": "24hrTicker",
		 "E": 1716889413050,
		 "T": 1716889413000,
		 "s": "ETH-240628-800-C",
		 "o": "3010.6",
		 "h": "3010.6",
		 "l": "3010.6",
		 "c": "3010.6",
		 "V": "0",
		 "A": "0",
		 "P": "0",
		 "p": "0",
		 "Q": "0",
		 "F": "0",
		 "L": "0",
		 "n": 0,
		 "bo": "50.9",
		 "ao": "0",
		 "bq": "0.01",
		 "aq": "0",
		 "b": "0.0000006",
		 "a": "-0.00000001",
		 "d": "1",
		 "t": "0",
		 "g": "0",
		 "v": "0",
		 "vo": "0.575",
		 "mp": "3086.9",
		 "hl": "3377.9",
		 "ll": "2795.9",
		 "eep": "0",
		 "r": "0.1"
		},
		{
			"e": "24hrTicker",
			"E": 1716889413050,
			"T": 1716889413000,
			"s": "ETH-240628-800-C",
			"o": "3010.6",
			"h": "3010.6",
			"l": "3010.6",
			"c": "3010.6",
			"V": "0",
			"A": "0",
			"P": "0",
			"p": "0",
			"Q": "0",
			"F": "0",
			"L": "0",
			"n": 0,
			"bo": "50.9",
			"ao": "0",
			"bq": "0.01",
			"aq": "0",
			"b": "0.0000006",
			"a": "-0.00000001",
			"d": "1",
			"t": "0",
			"g": "0",
			"v": "0",
			"vo": "0.575",
			"mp": "3086.9",
			"hl": "3377.9",
			"ll": "2795.9",
			"eep": "0",
			"r": "0.1"
		   }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsTickerWithExpireServe("ETH", "240628", func(event []*WsTickerEvent) {
		e := []*WsTickerEvent{
			{
				Event:              "24hrTicker",
				Time:               1716889413050,
				TradeTime:          1716889413000,
				Symbol:             "ETH-240628-800-C",
				Open:               "3010.6",
				High:               "3010.6",
				Low:                "3010.6",
				Close:              "3010.6",
				Volume:             "0",
				Amount:             "0",
				PriceChangePercent: "0",
				PriceChange:        "0",
				LastQty:            "0",
				FirstID:            "0",
				LastID:             "0",
				TradeNum:           0,
				BidOpenPrice:       "50.9",
				AskOpenPrice:       "0",
				BidQty:             "0.01",
				AskQty:             "0",
				BidIV:              "0.0000006",
				AskIV:              "-0.00000001",
				Delta:              "1",
				Theta:              "0",
				Gamma:              "0",
				Vega:               "0",
				Volatility:         "0.575",
				MarkPrice:          "3086.9",
				HighPriceLimit:     "3377.9",
				LowPriceLimit:      "2795.9",
				ExercisPrice:       "0",
				RiskFreeInterest:   "0.1",
			},
			{
				Event:              "24hrTicker",
				Time:               1716889413050,
				TradeTime:          1716889413000,
				Symbol:             "ETH-240628-800-C",
				Open:               "3010.6",
				High:               "3010.6",
				Low:                "3010.6",
				Close:              "3010.6",
				Volume:             "0",
				Amount:             "0",
				PriceChangePercent: "0",
				PriceChange:        "0",
				LastQty:            "0",
				FirstID:            "0",
				LastID:             "0",
				TradeNum:           0,
				BidOpenPrice:       "50.9",
				AskOpenPrice:       "0",
				BidQty:             "0.01",
				AskQty:             "0",
				BidIV:              "0.0000006",
				AskIV:              "-0.00000001",
				Delta:              "1",
				Theta:              "0",
				Gamma:              "0",
				Vega:               "0",
				Volatility:         "0.575",
				MarkPrice:          "3086.9",
				HighPriceLimit:     "3377.9",
				LowPriceLimit:      "2795.9",
				ExercisPrice:       "0",
				RiskFreeInterest:   "0.1",
			},
		}
		s.assertWsTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsOpenInterestEvent(ee, aa []*WsOpenInterestEvent) {
	r := s.r()
	for i, e := range ee {
		a := aa[i]
		r.Equal(e.Event, a.Event, "Event")
		r.Equal(e.Time, a.Time, "Time")
		r.Equal(e.Symbol, a.Symbol, "Symbol")
		r.Equal(e.OpenInterest, a.OpenInterest, "OpenInterest")
		r.Equal(e.Hold, a.Hold, "Hold")
	}
}

func (s *websocketServiceTestSuite) TestOpenInterestServe() {
	data := []byte(`[
		{
		 "e": "openInterest",
		 "E": 1716889802320,
		 "s": "ETH-240628-7500-C",
		 "o": "346.04",
		 "h": "1346822.0480630072"
		},
		{
		 "e": "openInterest",
		 "E": 1716889802320,
		 "s": "ETH-240628-3800-C",
		 "o": "390.05",
		 "h": "1518113.339056109"
		}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsOpenInterestServe("ETH", "20240628", func(event []*WsOpenInterestEvent) {
		e := []*WsOpenInterestEvent{
			{
				Event:        "openInterest",
				Time:         1716889802320,
				Symbol:       "ETH-240628-7500-C",
				OpenInterest: "346.04",
				Hold:         "1346822.0480630072",
			},
			{
				Event:        "openInterest",
				Time:         1716889802320,
				Symbol:       "ETH-240628-3800-C",
				OpenInterest: "390.05",
				Hold:         "1518113.339056109",
			},
		}
		s.assertWsOpenInterestEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsOptionPairEvent(e, a *WsOptionPairEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.CId, a.CId, "CId")
	r.Equal(e.Underlying, a.Underlying, "Underlying")
	r.Equal(e.QuoteAsset, a.QuoteAsset, "QuoteAsset")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Unit, a.Unit, "Unit")
	r.Equal(e.MinQuantity, a.MinQuantity, "MinQuantity")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.StrikePrice, a.StrikePrice, "StrikePrice")
	r.Equal(e.ExerciseDate, a.ExerciseDate, "ExerciseDate")
}

func (s *websocketServiceTestSuite) TestOptionPairServe() {
	data := []byte(`{
		"e":"OPTION_PAIR",        
		"E":1668573571842,       
		"id":652,                 
		"cid":2,                  
		"u":"BTCUSDT",            
		"qa":"USDT",              
		"s":"BTC-221116-21000-C", 
		"unit":1,                 
		"mq":"0.01",             
		"d":"CALL",               
		"sp":"21000",             
		"ed":1668585600000        
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsOptionPairServe(func(event *WsOptionPairEvent) {
		e := &WsOptionPairEvent{
			Event:        "OPTION_PAIR",
			Time:         1668573571842,
			Id:           652,
			CId:          2,
			Underlying:   "BTCUSDT",
			QuoteAsset:   "USDT",
			Symbol:       "BTC-221116-21000-C",
			Unit:         1,
			MinQuantity:  "0.01",
			Type:         "CALL",
			StrikePrice:  "21000",
			ExerciseDate: 1668585600000,
		}
		s.assertWsOptionPairEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsDepthEvent(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	r.Equal(e.PrevLastUpdateID, a.PrevLastUpdateID, "PrevLastUpdateID")
	for i, _ := range e.Bids {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Bids.Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Bids.Quantity")
	}
	for i, _ := range e.Asks {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Asks.Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Asks.Quantity")
	}
}

func (s *websocketServiceTestSuite) TestDepthServe() {
	data := []byte(`{
		"e": "depth",
		"E": 1716890957050,
		"T": 1716890955587,
		"s": "ETH-240927-5500-P",
		"u": 91644,
		"pu": 91644,
		"b": [
		 [
		  "1676.9",
		  "19.95"
		 ],
		 [
		  "1578.5",
		  "11.12"
		 ],
		 [
		  "1369.7",
		  "3.74"
		 ],
		 [
		  "0.1",
		  "0.3"
		 ]
		],
		"a": [
		 [
		  "1778.5",
		  "19.79"
		 ],
		 [
		  "1796.6",
		  "11.12"
		 ],
		 [
		  "2054.1",
		  "6.87"
		 ]
		]
	   }
	   `)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsDepthServe("ETH-240927-5500-P", "10", nil, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depth",
			Time:             1716890957050,
			TransactionTime:  1716890955587,
			Symbol:           "ETH-240927-5500-P",
			LastUpdateID:     91644,
			PrevLastUpdateID: 91644,
			Bids: []Bid{
				{
					Price:    "1676.9",
					Quantity: "19.95",
				},
				{
					Price:    "1578.5",
					Quantity: "11.12",
				},
				{
					Price:    "1369.7",
					Quantity: "3.74",
				},
				{
					Price:    "0.1",
					Quantity: "0.3",
				},
			},
			Asks: []Ask{
				{
					Price:    "1778.5",
					Quantity: "19.79",
				},
				{
					Price:    "1796.6",
					Quantity: "11.12",
				},
				{
					Price:    "2054.1",
					Quantity: "6.87",
				},
			},
		}

		s.assertWsDepthEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestCombinedServe1() {
	data := []byte(`{"stream": "ETH-240927-5500-P@depth10", "data": {
		"e": "depth",
		"E": 1716890957050,
		"T": 1716890955587,
		"s": "ETH-240927-5500-P",
		"u": 91644,
		"pu": 91644,
		"b": [
		 [
		  "1676.9",
		  "19.95"
		 ],
		 [
		  "1578.5",
		  "11.12"
		 ],
		 [
		  "1369.7",
		  "3.74"
		 ],
		 [
		  "0.1",
		  "0.3"
		 ]
		],
		"a": [
		 [
		  "1778.5",
		  "19.79"
		 ],
		 [
		  "1796.6",
		  "11.12"
		 ],
		 [
		  "2054.1",
		  "6.87"
		 ]
		]
	   }}
	   `)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedServe([]string{"ETH-240927-5500-P@depth10"}, map[string]any{"depth": (func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depth",
			Time:             1716890957050,
			TransactionTime:  1716890955587,
			Symbol:           "ETH-240927-5500-P",
			LastUpdateID:     91644,
			PrevLastUpdateID: 91644,
			Bids: []Bid{
				{
					Price:    "1676.9",
					Quantity: "19.95",
				},
				{
					Price:    "1578.5",
					Quantity: "11.12",
				},
				{
					Price:    "1369.7",
					Quantity: "3.74",
				},
				{
					Price:    "0.1",
					Quantity: "0.3",
				},
			},
			Asks: []Ask{
				{
					Price:    "1778.5",
					Quantity: "19.79",
				},
				{
					Price:    "1796.6",
					Quantity: "11.12",
				},
				{
					Price:    "2054.1",
					Quantity: "6.87",
				},
			},
		}

		s.assertWsDepthEvent(e, event)
	})},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestCombinedServe2() {
	data := []byte(`{"stream": "ETH-240927-5500-P@ticker@240927", "data": [
		{
		 "e": "24hrTicker",
		 "E": 1716889413050,
		 "T": 1716889413000,
		 "s": "ETH-240628-800-C",
		 "o": "3010.6",
		 "h": "3010.6",
		 "l": "3010.6",
		 "c": "3010.6",
		 "V": "0",
		 "A": "0",
		 "P": "0",
		 "p": "0",
		 "Q": "0",
		 "F": "0",
		 "L": "0",
		 "n": 0,
		 "bo": "50.9",
		 "ao": "0",
		 "bq": "0.01",
		 "aq": "0",
		 "b": "0.0000006",
		 "a": "-0.00000001",
		 "d": "1",
		 "t": "0",
		 "g": "0",
		 "v": "0",
		 "vo": "0.575",
		 "mp": "3086.9",
		 "hl": "3377.9",
		 "ll": "2795.9",
		 "eep": "0",
		 "r": "0.1"
		},
		{
			"e": "24hrTicker",
			"E": 1716889413050,
			"T": 1716889413000,
			"s": "ETH-240628-800-C",
			"o": "3010.6",
			"h": "3010.6",
			"l": "3010.6",
			"c": "3010.6",
			"V": "0",
			"A": "0",
			"P": "0",
			"p": "0",
			"Q": "0",
			"F": "0",
			"L": "0",
			"n": 0,
			"bo": "50.9",
			"ao": "0",
			"bq": "0.01",
			"aq": "0",
			"b": "0.0000006",
			"a": "-0.00000001",
			"d": "1",
			"t": "0",
			"g": "0",
			"v": "0",
			"vo": "0.575",
			"mp": "3086.9",
			"hl": "3377.9",
			"ll": "2795.9",
			"eep": "0",
			"r": "0.1"
		   }]}
	   `)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedServe([]string{"ETH-240927-5500-P@ticker"}, map[string]any{"ticker": (func(event []*WsTickerEvent) {
		e := []*WsTickerEvent{
			{
				Event:              "24hrTicker",
				Time:               1716889413050,
				TradeTime:          1716889413000,
				Symbol:             "ETH-240628-800-C",
				Open:               "3010.6",
				High:               "3010.6",
				Low:                "3010.6",
				Close:              "3010.6",
				Volume:             "0",
				Amount:             "0",
				PriceChangePercent: "0",
				PriceChange:        "0",
				LastQty:            "0",
				FirstID:            "0",
				LastID:             "0",
				TradeNum:           0,
				BidOpenPrice:       "50.9",
				AskOpenPrice:       "0",
				BidQty:             "0.01",
				AskQty:             "0",
				BidIV:              "0.0000006",
				AskIV:              "-0.00000001",
				Delta:              "1",
				Theta:              "0",
				Gamma:              "0",
				Vega:               "0",
				Volatility:         "0.575",
				MarkPrice:          "3086.9",
				HighPriceLimit:     "3377.9",
				LowPriceLimit:      "2795.9",
				ExercisPrice:       "0",
				RiskFreeInterest:   "0.1",
			},
			{
				Event:              "24hrTicker",
				Time:               1716889413050,
				TradeTime:          1716889413000,
				Symbol:             "ETH-240628-800-C",
				Open:               "3010.6",
				High:               "3010.6",
				Low:                "3010.6",
				Close:              "3010.6",
				Volume:             "0",
				Amount:             "0",
				PriceChangePercent: "0",
				PriceChange:        "0",
				LastQty:            "0",
				FirstID:            "0",
				LastID:             "0",
				TradeNum:           0,
				BidOpenPrice:       "50.9",
				AskOpenPrice:       "0",
				BidQty:             "0.01",
				AskQty:             "0",
				BidIV:              "0.0000006",
				AskIV:              "-0.00000001",
				Delta:              "1",
				Theta:              "0",
				Gamma:              "0",
				Vega:               "0",
				Volatility:         "0.575",
				MarkPrice:          "3086.9",
				HighPriceLimit:     "3377.9",
				LowPriceLimit:      "2795.9",
				ExercisPrice:       "0",
				RiskFreeInterest:   "0.1",
			},
		}

		s.assertWsTickerEvent(e, event)
	})},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsUserDataEvent(e, a *WsUserDataEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event)
	r.Equal(e.Time, a.Time)
	if e.RLCStatus != nil {
		r.Equal(e.RLCStatus, a.RLCStatus)
		r.Equal(e.RLCMaintenanceMargin, a.RLCMaintenanceMargin)
		r.Equal(e.RLCMarginBalance, a.RLCMarginBalance)
	}
	if e.AUBalance != nil || e.AUGreek != nil {
		for i, _ := range e.AUBalance {
			r.Equal(e.AUBalance[i].Asset, a.AUBalance[i].Asset)
			r.Equal(e.AUBalance[i].Balance, a.AUBalance[i].Balance)
			r.Equal(e.AUBalance[i].Merit, a.AUBalance[i].Merit)
			r.Equal(e.AUBalance[i].UnrealizedPnL, a.AUBalance[i].UnrealizedPnL)
			r.Equal(e.AUBalance[i].MaintenanceMargin, a.AUBalance[i].MaintenanceMargin)
			r.Equal(e.AUBalance[i].InitialMargin, a.AUBalance[i].InitialMargin)
		}

		for i, _ := range e.AUGreek {
			r.Equal(e.AUGreek[i].UnderlyingId, a.AUGreek[i].UnderlyingId)
			r.Equal(e.AUGreek[i].Delta, a.AUGreek[i].Delta)
			r.Equal(e.AUGreek[i].Theta, a.AUGreek[i].Theta)
			r.Equal(e.AUGreek[i].Gamma, a.AUGreek[i].Gamma)
			r.Equal(e.AUGreek[i].Vega, a.AUGreek[i].Vega)
		}

		for i, _ := range e.AUPosition {
			r.Equal(e.AUPosition[i].Symbol, a.AUPosition[i].Symbol)
			r.Equal(e.AUPosition[i].CountQty, a.AUPosition[i].CountQty)
			r.Equal(e.AUPosition[i].ReduleQty, a.AUPosition[i].ReduleQty)
			r.Equal(e.AUPosition[i].PositionVal, a.AUPosition[i].PositionVal)
			r.Equal(e.AUPosition[i].AvgPrice, a.AUPosition[i].AvgPrice)
		}
		r.Equal(e.AUUid, a.AUUid)
	}

	if e.OTU != nil {
		for j, _ := range e.OTU {
			r.Equal(e.OTU[j].CreateTime, a.OTU[j].CreateTime)
			r.Equal(e.OTU[j].UpdateTime, a.OTU[j].UpdateTime)
			r.Equal(e.OTU[j].Symbol, a.OTU[j].Symbol)
			r.Equal(e.OTU[j].ClientOrderID, a.OTU[j].ClientOrderID)
			r.Equal(e.OTU[j].OrderId, a.OTU[j].OrderId)
			r.Equal(e.OTU[j].Price, a.OTU[j].Price)
			r.Equal(e.OTU[j].Quantity, a.OTU[j].Quantity)
			r.Equal(e.OTU[j].Stp, a.OTU[j].Stp)
			r.Equal(e.OTU[j].ReduleOnly, a.OTU[j].ReduleOnly)
			r.Equal(e.OTU[j].PostOnly, a.OTU[j].PostOnly)
			r.Equal(e.OTU[j].Status, a.OTU[j].Status)
			r.Equal(e.OTU[j].ExecutedQty, a.OTU[j].ExecutedQty)
			r.Equal(e.OTU[j].ExecutedCost, a.OTU[j].ExecutedCost)
			r.Equal(e.OTU[j].Fee, a.OTU[j].Fee)
			r.Equal(e.OTU[j].TimeInForce, a.OTU[j].TimeInForce)
			r.Equal(e.OTU[j].OrderType, a.OTU[j].OrderType)
			for i, _ := range e.OTU[j].Filled {
				r.Equal(e.OTU[j].Filled[i].TradeId, a.OTU[j].Filled[i].TradeId)
				r.Equal(e.OTU[j].Filled[i].Price, a.OTU[j].Filled[i].Price)
				r.Equal(e.OTU[j].Filled[i].Quantity, a.OTU[j].Filled[i].Quantity)
				r.Equal(e.OTU[j].Filled[i].TradedTime, a.OTU[j].Filled[i].TradedTime)
				r.Equal(e.OTU[j].Filled[i].Marker, a.OTU[j].Filled[i].Marker)
				r.Equal(e.OTU[j].Filled[i].Fee, a.OTU[j].Filled[i].Fee)
			}
		}

	}
}

func (s *websocketServiceTestSuite) TestUserDataServe1() {
	data := []byte(`{
		"e":"RISK_LEVEL_CHANGE", 
		"E":1587727187525,       
		"s":"REDUCE_ONLY",      
		"mb":"1534.11708371",   
		"mm":"254789.11708371"  
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	strHelper := func(s string) *string {
		return &s
	}

	doneC, stopC, err := WsUserDataServe("xxyyzz", func(event *WsUserDataEvent) {
		e := &WsUserDataEvent{
			Event:                "RISK_LEVEL_CHANGE",
			Time:                 1587727187525,
			RLCStatus:            strHelper("REDUCE_ONLY"),
			RLCMarginBalance:     strHelper("1534.11708371"),
			RLCMaintenanceMargin: strHelper("254789.11708371"),
		}
		s.assertWsUserDataEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestUserDataServe2() {
	data := []byte(`{
		"e":"ORDER_TRADE_UPDATE",           
		"E":1657613775883,                  
		"o":[
		  {
			"T":1657613342918,              
			"t":1657613342918,              
			"s":"BTC-220930-18000-C",       
			"c":"",                         
			"oid":"4611869636869226548",    
			"p":"1993",                     
			"q":"1",                        
			"stp":0,                        
			"r":false,                      
			"po":true,                      
			"S":"PARTIALLY_FILLED",         
			"e":"0.1",                      
			"ec":"199.3",                   
			"f":"2",                       
			"tif": "GTC",                   
			"oty":"LIMIT",                  
			"fi":[
			  {
				"t":"20",                   
				"p":"1993",                 
				"q":"0.1",                  
				"T":1657613774336,          
				"m":"MAKER",                
				"f":"0.0002"                
			  }
			]
		  }
		]
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsUserDataServe("xxyyzz", func(event *WsUserDataEvent) {
		e := &WsUserDataEvent{
			Event: "ORDER_TRADE_UPDATE",
			Time:  1657613775883,
			OTU: []*WsOrderTradeUpdate{
				{
					CreateTime:    1657613342918,
					UpdateTime:    1657613342918,
					Symbol:        "BTC-220930-18000-C",
					ClientOrderID: "",
					OrderId:       "4611869636869226548",
					Price:         "1993",
					Quantity:      "1",
					Stp:           0,
					ReduleOnly:    false,
					PostOnly:      true,
					Status:        "PARTIALLY_FILLED",
					ExecutedQty:   "0.1",
					ExecutedCost:  "199.3",
					Fee:           "2",
					TimeInForce:   "GTC",
					OrderType:     "LIMIT",
					Filled: []WsFilled{
						{
							TradeId:    "20",
							Price:      "1993",
							Quantity:   "0.1",
							TradedTime: 1657613774336,
							Marker:     "MAKER",
							Fee:        "0.0002",
						},
					},
				},
			}}
		s.assertWsUserDataEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestUserDataServe3() {
	data := []byte(`{
		"e":"ORDER_TRADE_UPDATE",           
		"E":1657613775883,                  
		"o":[
		  {
			"T":1657613342918,              
			"t":1657613342918,              
			"s":"BTC-220930-18000-C",       
			"c":"",                         
			"oid":"4611869636869226548",    
			"p":"1993",                     
			"q":"1",                        
			"stp":0,                        
			"r":false,                      
			"po":true,                      
			"S":"PARTIALLY_FILLED",         
			"e":"0.1",                      
			"ec":"199.3",                   
			"f":"2",                       
			"tif": "GTC",                   
			"oty":"LIMIT"                 
		  }
		]
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsUserDataServe("xxyyzz", func(event *WsUserDataEvent) {
		e := &WsUserDataEvent{
			Event: "ORDER_TRADE_UPDATE",
			Time:  1657613775883,
			OTU: []*WsOrderTradeUpdate{
				{
					CreateTime:    1657613342918,
					UpdateTime:    1657613342918,
					Symbol:        "BTC-220930-18000-C",
					ClientOrderID: "",
					OrderId:       "4611869636869226548",
					Price:         "1993",
					Quantity:      "1",
					Stp:           0,
					ReduleOnly:    false,
					PostOnly:      true,
					Status:        "PARTIALLY_FILLED",
					ExecutedQty:   "0.1",
					ExecutedCost:  "199.3",
					Fee:           "2",
					TimeInForce:   "GTC",
					OrderType:     "LIMIT",
				},
			}}
		s.assertWsUserDataEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestUserDataServe4() {
	data := []byte(`{
		"e":"ACCOUNT_UPDATE",                
		"E":1591696384141,                    
		"B":[
			{
			  "b":"100007992.26053177",       
			  "m":"0",                        
			  "u":"458.782655111111",         
			  "M":"-15452.328456",            
			  "i":"-18852.328456",           
			  "a":"USDT"                    
			}
		],
		"G":[
			{
			 "ui":"SOLUSDT",                  
			 "d": -33.2933905,                
			 "t": 35.5926375,                  
			 "g": -14.3023855,                 
			 "v": -0.1929375                     
			}
		],
		"P":[
		  {
		   "s":"SOL-220912-35-C",              
		   "c":"-50",                          
		   "r":"-50",                         
		   "p":"-100",                        
		   "a":"2.2"                        
		  }
		],
		"uid":1000006559949}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	int64Helper := func(i int64) *int64 {
		return &i
	}

	doneC, stopC, err := WsUserDataServe("xxyyzz", func(event *WsUserDataEvent) {
		e := &WsUserDataEvent{

			Event: "ACCOUNT_UPDATE",
			Time:  1591696384141,
			AUBalance: []*WsBalance{
				{
					Balance:           "100007992.26053177",
					Merit:             "0",
					UnrealizedPnL:     "458.782655111111",
					MaintenanceMargin: "-15452.328456",
					InitialMargin:     "-18852.328456",
					Asset:             "USDT",
				},
			},
			AUGreek: []*WsGreek{
				{
					UnderlyingId: "SOLUSDT",
					Delta:        -33.2933905,
					Theta:        35.5926375,
					Gamma:        -14.3023855,
					Vega:         -0.1929375,
				},
			},
			AUPosition: []*WsPosition{
				{
					Symbol:      "SOL-220912-35-C",
					CountQty:    "-50",
					ReduleQty:   "-50",
					PositionVal: "-100",
					AvgPrice:    "2.2",
				},
			},
			AUUid: int64Helper(1000006559949),
		}
		s.assertWsUserDataEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}
