package futures

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

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

func (s *websocketServiceTestSuite) TestAggTradeServe() {
	data := []byte(`{
		"e": "aggTrade",
		"E": 123456789,
		"s": "BTCUSDT",
		"a": 5933014,
		"p": "0.001",
		"q": "100",
		"f": 100,
		"l": 105,
		"T": 123456785,
		"m": true
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAggTradeServe("BTCUSDT", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:            "aggTrade",
			Time:             123456789,
			Symbol:           "BTCUSDT",
			AggregateTradeID: 5933014,
			Price:            "0.001",
			Quantity:         "100",
			FirstTradeID:     100,
			LastTradeID:      105,
			TradeTime:        123456785,
			Maker:            true,
		}
		s.assertWsAggTradeEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestCombinedAggTradeServe() {
	data := []byte(`{
			"stream":"btcusdt@aggTrade",
			"data":{
				"e":"aggTrade",
				"E":1628843331742,
				"a":105688535,
				"s":"BTCUSDT",
				"p":"46063.00",
				"q":"0.005",
				"f":188354417,
				"l":188354417,
				"T":1628843331590,
				"m":false}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedAggTradeServe([]string{"BTCUSDT"}, func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:            "aggTrade",
			Time:             1628843331742,
			Symbol:           "BTCUSDT",
			AggregateTradeID: 105688535,
			Price:            "46063.00",
			Quantity:         "0.005",
			FirstTradeID:     188354417,
			LastTradeID:      188354417,
			TradeTime:        1628843331590,
			Maker:            false,
		}
		s.assertWsAggTradeEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAggTradeEvent(e, a *WsAggTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AggregateTradeID, a.AggregateTradeID, "AggregateTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstTradeID, a.FirstTradeID, "FirstTradeID")
	r.Equal(e.LastTradeID, a.LastTradeID, "LastTradeID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.Maker, a.Maker, "Maker")
}

func (s *websocketServiceTestSuite) testMarkPriceServe(rate *time.Duration, expectedErr error, expectedServeCnt int) {
	data := []byte(`{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11794.15000000",
		"i": "11784.62659091",
		"r": "0.00038167",
		"T": 1562306400000
	  }`)
	s.mockWsServe(data, expectedErr)
	defer s.assertWsServe(expectedServeCnt)

	handler := func(event *WsMarkPriceEvent) {
		e := &WsMarkPriceEvent{
			Event:           "markPriceUpdate",
			Time:            1562305380000,
			Symbol:          "BTCUSDT",
			MarkPrice:       "11794.15000000",
			IndexPrice:      "11784.62659091",
			FundingRate:     "0.00038167",
			NextFundingTime: 1562306400000,
		}
		s.assertWsMarkPriceEvent(e, event)
	}
	errHandler := func(err error) {
	}

	var doneC, stopC chan struct{}
	var err error
	if rate == nil {
		doneC, stopC, err = WsMarkPriceServe("BTCUSDT", handler, errHandler)
	} else {
		doneC, stopC, err = WsMarkPriceServeWithRate("BTCUSDT", *rate, handler, errHandler)
	}

	if expectedErr == nil {
		s.r().NoError(err)
	} else {
		s.r().EqualError(err, expectedErr.Error())
	}

	if stopC != nil {
		stopC <- struct{}{}
	}
	if doneC != nil {
		<-doneC
	}
}

func (s *websocketServiceTestSuite) TestMarkPriceServe() {
	s.testMarkPriceServe(nil, nil, 1)
}

func (s *websocketServiceTestSuite) TestMarkPriceServeWithValidRate() {
	rate := 3 * time.Second
	s.testMarkPriceServe(&rate, nil, 1)
	rate = time.Second
	s.testMarkPriceServe(&rate, nil, 2)
}

func (s *websocketServiceTestSuite) TestMarkPriceServeWithInvalidRate() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	for {
		rate := time.Duration(rand.Intn(10)) * time.Second
		switch rate {
		case 3 * time.Second:
		case 1 * time.Second:
		default:
			s.testMarkPriceServe(&rate, errors.New("Invalid rate"), 0)
			return
		}
	}
}

func (s *websocketServiceTestSuite) testAllMarkPriceServe(rate *time.Duration, expectedErr error, expectedServeCnt int) {
	data := []byte(`[{
		"e": "markPriceUpdate",
		"E": 1562305380000,
		"s": "BTCUSDT",
		"p": "11794.15000000",
		"i": "11784.62659091",
		"r": "0.00038167",
		"T": 1562306400000
	  }]`)
	s.mockWsServe(data, expectedErr)
	defer s.assertWsServe(expectedServeCnt)

	handler := func(event WsAllMarkPriceEvent) {
		e := WsAllMarkPriceEvent{{
			Event:           "markPriceUpdate",
			Time:            1562305380000,
			Symbol:          "BTCUSDT",
			MarkPrice:       "11794.15000000",
			IndexPrice:      "11784.62659091",
			FundingRate:     "0.00038167",
			NextFundingTime: 1562306400000,
		}}
		s.assertWsMarkPriceEvent(e[0], event[0])
	}
	errHandler := func(err error) {}

	var doneC, stopC chan struct{}
	var err error
	if rate == nil {
		doneC, stopC, err = WsAllMarkPriceServe(handler, errHandler)
	} else {
		doneC, stopC, err = WsAllMarkPriceServeWithRate(*rate, handler, errHandler)
	}

	if expectedErr == nil {
		s.r().NoError(err)
	} else {
		s.r().EqualError(err, expectedErr.Error())
	}

	if stopC != nil {
		stopC <- struct{}{}
	}
	if doneC != nil {
		<-doneC
	}
}

func (s *websocketServiceTestSuite) assertWsMarkPriceEvent(e, a *WsMarkPriceEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.IndexPrice, a.IndexPrice, "IndexPrice")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	r.Equal(e.NextFundingTime, a.NextFundingTime, "NextFundingTime")
}

func (s *websocketServiceTestSuite) TestAllMarkPriceServe() {
	s.testAllMarkPriceServe(nil, nil, 1)
}

func (s *websocketServiceTestSuite) TestAllMarkPriceServeWithValidRate() {
	rate := 3 * time.Second
	s.testAllMarkPriceServe(&rate, nil, 1)
	rate = time.Second
	s.testAllMarkPriceServe(&rate, nil, 2)
}

func (s *websocketServiceTestSuite) TestAllMarkPriceServeWithInvalidRate() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	for {
		rate := time.Duration(rand.Intn(10)) * time.Second
		switch rate {
		case 3 * time.Second:
		case 1 * time.Second:
		default:
			s.testAllMarkPriceServe(&rate, errors.New("Invalid rate"), 0)
			return
		}
	}
}

func (s *websocketServiceTestSuite) testCombinedMarkPriceServe(rate *time.Duration, expectedErr error, expectedServeCnt int) {
	data := []byte(`{
    "stream": "btcusdt@markPrice",
    "data": {
        "e": "markPriceUpdate",
        "E": 1681724175000,
        "s": "BTCUSDT",
        "p": "29892.78738889",
        "P": "29903.84541674",
        "i": "29904.57564103",
        "r": "0.00010000",
        "T": 1681747200000
    }}`)
	s.mockWsServe(data, expectedErr)
	defer s.assertWsServe(expectedServeCnt)

	handler := func(event *WsMarkPriceEvent) {
		e := &WsMarkPriceEvent{
			Event:           "markPriceUpdate",
			Time:            1681724175000,
			Symbol:          "BTCUSDT",
			MarkPrice:       "29892.78738889",
			IndexPrice:      "29904.57564103",
			FundingRate:     "0.00010000",
			NextFundingTime: 1681747200000,
		}
		s.assertWsMarkPriceEvent(e, event)
	}
	errHandler := func(err error) {
	}

	var doneC, stopC chan struct{}
	var err error

	if rate == nil {
		input := []string{"BTCUSDT"}
		doneC, stopC, err = WsCombinedMarkPriceServe(input, handler, errHandler)
	} else {
		input := map[string]time.Duration{"BTCUSDT": *rate}
		doneC, stopC, err = WsCombinedMarkPriceServeWithRate(input, handler, errHandler)
	}

	if expectedErr == nil {
		s.r().NoError(err)
	} else {
		s.r().EqualError(err, expectedErr.Error())
	}

	if stopC != nil {
		stopC <- struct{}{}
	}
	if doneC != nil {
		<-doneC
	}
}

func (s *websocketServiceTestSuite) TestCombinedMarkPriceServe() {
	s.testCombinedMarkPriceServe(nil, nil, 1)
}

func (s *websocketServiceTestSuite) TestCombinedMarkPriceServeWithValidRate() {
	rate := 3 * time.Second
	s.testCombinedMarkPriceServe(&rate, nil, 1)
	rate = time.Second
	s.testCombinedMarkPriceServe(&rate, nil, 2)
}

func (s *websocketServiceTestSuite) TestCombinedMarkPriceServeWithInvalidRate() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	for {
		rate := time.Duration(rand.Intn(10)) * time.Second
		switch rate {
		case 3 * time.Second:
		case 1 * time.Second:
		default:
			s.testCombinedMarkPriceServe(&rate, fmt.Errorf("invalid rate. Symbol BTCUSDT (rate %d)", rate), 0)
			return
		}
	}
}

func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
		"e": "kline",
		"E": 123456789,
		"s": "BTCUSDT",
		"k": {
		  "t": 123400000,
		  "T": 123460000,
		  "s": "BTCUSDT",
		  "i": "1m",
		  "f": 100,
		  "L": 200,
		  "o": "0.0010",
		  "c": "0.0020",
		  "h": "0.0025",
		  "l": "0.0015",
		  "v": "1000",
		  "n": 100,
		  "x": false,
		  "q": "1.0000",
		  "V": "500",
		  "Q": "0.500"
		}
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsKlineServe("ETHBTC", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   123456789,
			Symbol: "BTCUSDT",
			Kline: WsKline{
				StartTime:            123400000,
				EndTime:              123460000,
				Symbol:               "BTCUSDT",
				Interval:             "1m",
				FirstTradeID:         100,
				LastTradeID:          200,
				Open:                 "0.0010",
				Close:                "0.0020",
				High:                 "0.0025",
				Low:                  "0.0015",
				Volume:               "1000",
				TradeNum:             100,
				IsFinal:              false,
				QuoteVolume:          "1.0000",
				ActiveBuyVolume:      "500",
				ActiveBuyQuoteVolume: "0.500",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsKlineEventEqual(e, a *WsKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstTradeID, ak.FirstTradeID, "FirstTradeID")
	r.Equal(ek.LastTradeID, ak.LastTradeID, "LastTradeID")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.Volume, ak.Volume, "Volume")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
	r.Equal(ek.QuoteVolume, ak.QuoteVolume, "QuoteVolume")
	r.Equal(ek.ActiveBuyVolume, ak.ActiveBuyVolume, "ActiveBuyVolume")
	r.Equal(ek.ActiveBuyQuoteVolume, ak.ActiveBuyQuoteVolume, "ActiveBuyQuoteVolume")
}

func (s *websocketServiceTestSuite) TestWsCombinedKlineServe() {
	data := []byte(`{
	"stream":"ethbtc@kline_1m",
	"data": {
        "e": "kline",
        "E": 1499404907056,
        "s": "ETHBTC",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
	}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	input := map[string]string{
		"ETHBTC": "1m",
	}
	doneC, stopC, err := WsCombinedKlineServe(input, func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1499404907056,
			Symbol: "ETHBTC",
			Kline: WsKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Symbol:               "ETHBTC",
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestWsCombinedKlineServeMultiInterval() {
	data := []byte(`{
	"stream":"ethbtc@kline_1m",
	"data": {
        "e": "kline",
        "E": 1499404907056,
        "s": "ETHBTC",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
	}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	input := map[string][]string{
		"ETHBTC": {"1m", "5m"},
	}
	doneC, stopC, err := WsCombinedKlineServeMultiInterval(input, func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1499404907056,
			Symbol: "ETHBTC",
			Kline: WsKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Symbol:               "ETHBTC",
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestContinuousKlineServe() {
	data := []byte(`{
		"e": "continuous_kline",
		"E": 123456789,
		"ps": "BTCUSDT",
		"ct": "PERPETUAL",
		"k": {
		  "t": 123400000,
		  "T": 123460000,
		  "i": "1m",
		  "f": 100,
		  "L": 200,
		  "o": "0.0010",
		  "c": "0.0020",
		  "h": "0.0025",
		  "l": "0.0015",
		  "v": "1000",
		  "n": 100,
		  "x": false,
		  "q": "1.0000",
		  "V": "500",
		  "Q": "0.500"
		}
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsContinuousKlineServe(&WsContinuousKlineSubscribeArgs{
		Pair:         "BTCUSDT",
		ContractType: "PERPETUAL",
		Interval:     "1m",
	},
		func(event *WsContinuousKlineEvent) {
			e := &WsContinuousKlineEvent{
				Event:        "continuous_kline",
				Time:         123456789,
				PairSymbol:   "BTCUSDT",
				ContractType: "PERPETUAL",
				Kline: WsContinuousKline{
					StartTime:            123400000,
					EndTime:              123460000,
					Interval:             "1m",
					FirstTradeID:         100,
					LastTradeID:          200,
					Open:                 "0.0010",
					Close:                "0.0020",
					High:                 "0.0025",
					Low:                  "0.0015",
					Volume:               "1000",
					TradeNum:             100,
					IsFinal:              false,
					QuoteVolume:          "1.0000",
					ActiveBuyVolume:      "500",
					ActiveBuyQuoteVolume: "0.500",
				},
			}
			s.assertWsContinuousKlineEventEqual(e, event)
		}, func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsContinuousKlineEventEqual(e, a *WsContinuousKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.PairSymbol, a.PairSymbol, "PairSymbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstTradeID, ak.FirstTradeID, "FirstTradeID")
	r.Equal(ek.LastTradeID, ak.LastTradeID, "LastTradeID")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.Volume, ak.Volume, "Volume")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
	r.Equal(ek.QuoteVolume, ak.QuoteVolume, "QuoteVolume")
	r.Equal(ek.ActiveBuyVolume, ak.ActiveBuyVolume, "ActiveBuyVolume")
	r.Equal(ek.ActiveBuyQuoteVolume, ak.ActiveBuyQuoteVolume, "ActiveBuyQuoteVolume")
}

func (s *websocketServiceTestSuite) TestWsCombinedContinuousKlineServe() {
	data := []byte(`{
	"stream":"ethbtc_perpetual@continuousKline_1m",
	"data": {
        "e": "continuous_kline",
        "E": 1499404907056,
        "ps": "ETHBTC",
		"ct": "PERPETUAL",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
	}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	input := []*WsContinuousKlineSubscribeArgs{
		{
			Pair:         "ETHBTC",
			ContractType: "PERPETUAL",
			Interval:     "1m",
		},
	}
	doneC, stopC, err := WsCombinedContinuousKlineServe(input, func(event *WsContinuousKlineEvent) {
		e := &WsContinuousKlineEvent{
			Event:        "continuous_kline",
			Time:         1499404907056,
			PairSymbol:   "ETHBTC",
			ContractType: "PERPETUAL",
			Kline: WsContinuousKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsContinuousKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestMiniMarketTickerServe() {
	data := []byte(`{
		"e": "24hrMiniTicker",
		"E": 123456789,
		"s": "BTCUSDT",
		"c": "0.0025",
		"o": "0.0010",
		"h": "0.0025",
		"l": "0.0010",
		"v": "10000",
		"q": "18"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMiniMarketTickerServe("BTCUSDT", func(event *WsMiniMarketTickerEvent) {
		e := &WsMiniMarketTickerEvent{
			Event:       "24hrMiniTicker",
			Time:        123456789,
			Symbol:      "BTCUSDT",
			ClosePrice:  "0.0025",
			OpenPrice:   "0.0010",
			HighPrice:   "0.0025",
			LowPrice:    "0.0010",
			Volume:      "10000",
			QuoteVolume: "18",
		}
		s.assertWsMinMarketTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllMiniMarketTickerServe() {
	data := []byte(`[{
		"e": "24hrMiniTicker",
		"E": 123456789,
		"s": "BTCUSDT",
		"c": "0.0025",
		"o": "0.0010",
		"h": "0.0025",
		"l": "0.0010",
		"v": "10000",
		"q": "18"
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMiniMarketTickerServe(func(event WsAllMiniMarketTickerEvent) {
		e := []*WsMiniMarketTickerEvent{{
			Event:       "24hrMiniTicker",
			Time:        123456789,
			Symbol:      "BTCUSDT",
			ClosePrice:  "0.0025",
			OpenPrice:   "0.0010",
			HighPrice:   "0.0025",
			LowPrice:    "0.0010",
			Volume:      "10000",
			QuoteVolume: "18",
		}}
		s.assertWsMinMarketTickerEvent(e[0], event[0])
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMinMarketTickerEvent(e, a *WsMiniMarketTickerEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ClosePrice, a.ClosePrice, "ClosePrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
}

func (s *websocketServiceTestSuite) TestMarketTickerServe() {
	data := []byte(`{
		"e": "24hrTicker",
		"E": 123456789,
		"s": "BTCUSDT",
		"p": "0.0015",
		"P": "250.00",
		"w": "0.0018",
		"c": "0.0025",
		"Q": "10",
		"o": "0.0010",
		"h": "0.0025",
		"l": "0.0010",
		"v": "10000",
		"q": "18",
		"O": 0,
		"C": 86400000,
		"F": 0,
		"L": 18150,
		"n": 18151
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsMarketTickerServe("BTCUSDT", func(event *WsMarketTickerEvent) {
		e := &WsMarketTickerEvent{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BTCUSDT",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			ClosePrice:         "0.0025",
			CloseQty:           "10",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0025",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			TradeCount:         18151,
		}
		s.assertWsMarketTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllMarketTickerServe() {
	data := []byte(`[{
		"e": "24hrTicker",
		"E": 123456789,
		"s": "BTCUSDT",
		"p": "0.0015",
		"P": "250.00",
		"w": "0.0018",
		"c": "0.0025",
		"Q": "10",
		"o": "0.0010",
		"h": "0.0025",
		"l": "0.0010",
		"v": "10000",
		"q": "18",
		"O": 0,
		"C": 86400000,
		"F": 0,
		"L": 18150,
		"n": 18151
	  }]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMarketTickerServe(func(event WsAllMarketTickerEvent) {
		e := WsAllMarketTickerEvent{{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BTCUSDT",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			ClosePrice:         "0.0025",
			CloseQty:           "10",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0025",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			TradeCount:         18151,
		}}
		s.assertWsMarketTickerEvent(e[0], event[0])
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarketTickerEvent(e, a *WsMarketTickerEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.ClosePrice, a.ClosePrice, "ClosePrice")
	r.Equal(e.CloseQty, a.CloseQty, "CloseQty")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.TradeCount, a.TradeCount, "TradeCount")
}

func (s *websocketServiceTestSuite) TestBookTickerServe() {
	data := []byte(`{
		"e":"bookTicker",
		"u":400900217,
		"E": 1568014460893,
		"T": 1568014460891,
		"s":"BNBUSDT",
		"b":"25.35190000",
		"B":"31.21000000",
		"a":"25.36520000",
		"A":"40.66000000"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBookTickerServe("BNBUSDT", func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			Event:           "bookTicker",
			UpdateID:        400900217,
			Time:            1568014460893,
			TransactionTime: 1568014460891,
			Symbol:          "BNBUSDT",
			BestBidPrice:    "25.35190000",
			BestBidQty:      "31.21000000",
			BestAskPrice:    "25.36520000",
			BestAskQty:      "40.66000000",
		}
		s.assertWsBookTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllBookTickerServe() {
	data := []byte(`{
		"e":"bookTicker",
		"u":400900217,
		"E": 1568014460893,
		"T": 1568014460891,
		"s":"BNBUSDT",
		"b":"25.35190000",
		"B":"31.21000000",
		"a":"25.36520000",
		"A":"40.66000000"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllBookTickerServe(func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			Event:           "bookTicker",
			UpdateID:        400900217,
			Time:            1568014460893,
			TransactionTime: 1568014460891,
			Symbol:          "BNBUSDT",
			BestBidPrice:    "25.35190000",
			BestBidQty:      "31.21000000",
			BestAskPrice:    "25.36520000",
			BestAskQty:      "40.66000000",
		}
		s.assertWsBookTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsBookTickerEvent(e, a *WsBookTickerEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.UpdateID, a.UpdateID, "UpdateID")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BestBidPrice, a.BestBidPrice, "BestBidPrice")
	r.Equal(e.BestBidQty, a.BestBidQty, "BestBidQty")
	r.Equal(e.BestAskPrice, a.BestAskPrice, "BestAskPrice")
	r.Equal(e.BestAskQty, a.BestAskQty, "BestAskQty")
}

func (s *websocketServiceTestSuite) TestLiquidationOrderServe() {
	data := []byte(`{
		"e":"forceOrder",
		"E":1568014460893,
		"o":{
			"s":"BTCUSDT",
			"S":"SELL",
			"o":"LIMIT",
			"f":"IOC",
			"q":"0.014",
			"p":"9910",
			"ap":"9910",
			"X":"FILLED",
			"l":"0.014",
			"z":"0.014",
			"T":1568014460893
		}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsLiquidationOrderServe("BTCUSDT", func(event *WsLiquidationOrderEvent) {
		e := &WsLiquidationOrderEvent{
			Event: "forceOrder",
			Time:  1568014460893,
			LiquidationOrder: WsLiquidationOrder{
				Symbol:               "BTCUSDT",
				Side:                 SideTypeSell,
				OrderType:            OrderTypeLimit,
				TimeInForce:          TimeInForceTypeIOC,
				OrigQuantity:         "0.014",
				Price:                "9910",
				AvgPrice:             "9910",
				OrderStatus:          OrderStatusTypeFilled,
				LastFilledQty:        "0.014",
				AccumulatedFilledQty: "0.014",
				TradeTime:            1568014460893,
			},
		}
		s.assertLiquidationOrderEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestAllLiquidationOrderServe() {
	data := []byte(`{
		"e":"forceOrder",
		"E":1568014460893,
		"o":{
			"s":"BTCUSDT",
			"S":"SELL",
			"o":"LIMIT",
			"f":"IOC",
			"q":"0.014",
			"p":"9910",
			"ap":"9910",
			"X":"FILLED",
			"l":"0.014",
			"z":"0.014",
			"T":1568014460893
		}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllLiquidationOrderServe(func(event *WsLiquidationOrderEvent) {
		e := &WsLiquidationOrderEvent{
			Event: "forceOrder",
			Time:  1568014460893,
			LiquidationOrder: WsLiquidationOrder{
				Symbol:               "BTCUSDT",
				Side:                 SideTypeSell,
				OrderType:            OrderTypeLimit,
				TimeInForce:          TimeInForceTypeIOC,
				OrigQuantity:         "0.014",
				Price:                "9910",
				AvgPrice:             "9910",
				OrderStatus:          OrderStatusTypeFilled,
				LastFilledQty:        "0.014",
				AccumulatedFilledQty: "0.014",
				TradeTime:            1568014460893,
			},
		}
		s.assertLiquidationOrderEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertLiquidationOrderEvent(e, a *WsLiquidationOrderEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	elo, alo := e.LiquidationOrder, a.LiquidationOrder
	r.Equal(elo.Symbol, alo.Symbol, "Symbol")
	r.Equal(elo.Side, alo.Side, "Side")
	r.Equal(elo.OrderType, alo.OrderType, "OrderType")
	r.Equal(elo.TimeInForce, alo.TimeInForce, "TimeInForce")
	r.Equal(elo.OrigQuantity, alo.OrigQuantity, "OrigQuantity")
	r.Equal(elo.Price, alo.Price, "Price")
	r.Equal(elo.AvgPrice, alo.AvgPrice, "AvgPrice")
	r.Equal(elo.OrderStatus, alo.OrderStatus, "OrderStatus")
	r.Equal(elo.LastFilledQty, alo.LastFilledQty, "LastFilledQty")
	r.Equal(elo.AccumulatedFilledQty, alo.AccumulatedFilledQty, "AccumulatedFilledQty")
	r.Equal(elo.TradeTime, alo.TradeTime, "TradeTime")
}

func (s *websocketServiceTestSuite) testPartialDepthServe(rate *time.Duration, expectedErr error, expectedServeCnt int) {
	data := []byte(`{
		"e": "depthUpdate",
		"E": 1571889248277,
		"T": 1571889248276,
		"s": "BTCUSDT",
		"U": 390497796,
		"u": 390497878,
		"pu": 390497794,
		"b": [
		  [
			"7403.89",
			"0.002"
		  ]
		],
		"a": [
		  [
			"7405.96",
			"3.340"
		  ]
		]
	  }`)
	s.mockWsServe(data, expectedErr)
	defer s.assertWsServe(expectedServeCnt)

	handler := func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             1571889248277,
			TransactionTime:  1571889248276,
			Symbol:           "BTCUSDT",
			FirstUpdateID:    390497796,
			LastUpdateID:     390497878,
			PrevLastUpdateID: 390497794,
			Bids:             []Bid{{Price: "7403.89", Quantity: "0.002"}},
			Asks:             []Ask{{Price: "7405.96", Quantity: "3.340"}},
		}
		s.assertDepthEvent(e, event)
	}
	errHandler := func(err error) {
	}

	var doneC, stopC chan struct{}
	var err error
	if rate == nil {
		doneC, stopC, err = WsPartialDepthServe("BTCUSDT", 5, handler, errHandler)
	} else {
		doneC, stopC, err = WsPartialDepthServeWithRate("BTCUSDT", 5, *rate, handler, errHandler)
	}

	if expectedErr == nil {
		s.r().NoError(err)
	} else {
		s.r().EqualError(err, expectedErr.Error())
	}

	if stopC != nil {
		stopC <- struct{}{}
	}
	if doneC != nil {
		<-doneC
	}
}

func (s *websocketServiceTestSuite) TestPartialDepthServe() {
	s.testPartialDepthServe(nil, nil, 1)
}

func (s *websocketServiceTestSuite) TestPartialDepthServeWithValidRate() {
	rate := 250 * time.Millisecond
	s.testPartialDepthServe(&rate, nil, 1)
	rate = 500 * time.Millisecond
	s.testPartialDepthServe(&rate, nil, 2)
	rate = 100 * time.Millisecond
	s.testPartialDepthServe(&rate, nil, 3)
}

func (s *websocketServiceTestSuite) TestPartialDepthServeWithInvalidRate() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	for {
		rate := time.Duration(rand.Intn(100)*10) * time.Millisecond
		switch rate {
		case 250 * time.Millisecond:
		case 500 * time.Millisecond:
		case 100 * time.Millisecond:
		default:
			s.testPartialDepthServe(&rate, errors.New("Invalid rate"), 0)
			return
		}
	}
}

func (s *websocketServiceTestSuite) testDiffDepthServe(rate *time.Duration, expectedErr error, expectedServeCnt int) {
	data := []byte(`{
		"e": "depthUpdate",
		"E": 123456789,
		"T": 123456788,
		"s": "BTCUSDT",
		"U": 157,
		"u": 160,
		"pu": 149,
		"b": [
		  [
			"0.0024",
			"10"
		  ]
		],
		"a": [
		  [
			"0.0026",
			"100"
		  ]
		]
	  }`)
	s.mockWsServe(data, expectedErr)
	defer s.assertWsServe(expectedServeCnt)

	handler := func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             123456789,
			TransactionTime:  123456788,
			Symbol:           "BTCUSDT",
			FirstUpdateID:    157,
			LastUpdateID:     160,
			PrevLastUpdateID: 149,
			Bids:             []Bid{{Price: "0.0024", Quantity: "10"}},
			Asks:             []Ask{{Price: "0.0026", Quantity: "100"}},
		}
		s.assertDepthEvent(e, event)
	}
	errHandler := func(err error) {
	}

	var doneC, stopC chan struct{}
	var err error
	if rate == nil {
		doneC, stopC, err = WsDiffDepthServe("BTCUSDT", handler, errHandler)
	} else {
		doneC, stopC, err = WsDiffDepthServeWithRate("BTCUSDT", *rate, handler, errHandler)
	}

	if expectedErr == nil {
		s.r().NoError(err)
	} else {
		s.r().EqualError(err, expectedErr.Error())
	}

	if stopC != nil {
		stopC <- struct{}{}
	}
	if doneC != nil {
		<-doneC
	}
}

func (s *websocketServiceTestSuite) TestWsCombinedDiffDepthServe() {
	symbols := []string{"BTCUSDT"}
	data := []byte(`{
		"stream":"btcusdt@depth",
		"data":{
			"e":"depthUpdate",
			"E":1628847118038,
			"T":1628847117814,
			"s":"BTCUSDT",
			"U":21925649843,
			"u":21925649849,
			"pu":21925649651,
			"b":[["46248.03","0.000"]],
			"a":[["46249.88","71.870"]]}}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedDiffDepthServe(symbols, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:            "depthUpdate",
			Time:             1628847118038,
			TransactionTime:  1628847117814,
			Symbol:           "BTCUSDT",
			FirstUpdateID:    21925649843,
			LastUpdateID:     21925649849,
			PrevLastUpdateID: 21925649651,
			Bids:             []Bid{{Price: "46248.03", Quantity: "0.000"}},
			Asks:             []Ask{{Price: "46249.88", Quantity: "71.870"}},
		}
		s.assertDepthEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertDepthEvent(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.FirstUpdateID, a.FirstUpdateID, "FirstUpdateID")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	r.Equal(e.PrevLastUpdateID, a.PrevLastUpdateID, "PrevLastUpdateID")
	for i, b := range e.Bids {
		r.Equal(b.Price, a.Bids[i].Price, "Price")
		r.Equal(b.Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i, b := range e.Asks {
		r.Equal(b.Price, a.Asks[i].Price, "Price")
		r.Equal(b.Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestDiffDepthServe() {
	s.testDiffDepthServe(nil, nil, 1)
}

func (s *websocketServiceTestSuite) TestDiffDepthServeWithValidRate() {
	rate := 250 * time.Millisecond
	s.testDiffDepthServe(&rate, nil, 1)
	rate = 500 * time.Millisecond
	s.testDiffDepthServe(&rate, nil, 2)
	rate = 100 * time.Millisecond
	s.testDiffDepthServe(&rate, nil, 3)
}

func (s *websocketServiceTestSuite) TestDiffDepthServeWithInvalidRate() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(randSrc)
	for {
		rate := time.Duration(rand.Intn(100)*10) * time.Millisecond
		switch rate {
		case 250 * time.Millisecond:
		case 500 * time.Millisecond:
		case 100 * time.Millisecond:
		default:
			s.testDiffDepthServe(&rate, errors.New("Invalid rate"), 0)
			return
		}
	}
}

func (s *websocketServiceTestSuite) TestBLVTInfoServe() {
	data := []byte(`{
		"e":"nav",
		"E":1600245286355,
		"s":"TRXDOWN",
		"m":74164.75496502663,
		"b":[
			{
				"s":"TRXUSDT",
				"n":-87988261
			}
		],
		"n":14.78454447,
		"l":2.1786579638117898,
		"t":3,
		"f":-0.0048925
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBLVTInfoServe("TRXDOWN", func(event *WsBLVTInfoEvent) {
		e := &WsBLVTInfoEvent{
			Event:          "nav",
			Time:           1600245286355,
			Symbol:         "TRXDOWN",
			Issued:         74164.75496502663,
			Baskets:        []WsBLVTBasket{{Symbol: "TRXUSDT", Position: -87988261}},
			Nav:            14.78454447,
			Leverage:       2.1786579638117898,
			TargetLeverage: 3,
			FundingRate:    -0.0048925,
		}
		s.assertBLVTInfoEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertBLVTInfoEvent(e, a *WsBLVTInfoEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Issued, a.Issued, "Issued")
	r.Equal(e.Nav, a.Nav, "Nav")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
	r.Equal(e.TargetLeverage, a.TargetLeverage, "TargetLeverage")
	r.Equal(e.FundingRate, a.FundingRate, "FundingRate")
	for i, b := range e.Baskets {
		r.Equal(b.Position, a.Baskets[i].Position, "Position")
		r.Equal(b.Symbol, a.Baskets[i].Symbol, "Symbol")
	}
}

func (s *websocketServiceTestSuite) TestBLVTKlineServe() {
	data := []byte(`{
		"e":"kline",
		"E":1600243159447,
		"s":"TRXDOWN",
		"k":{
			"t":1600243140000,
			"T":1600243199999,
			"s":"TRXDOWN",
			"i":"1m",
			"f":1600243140484,
			"L":1600243159424,
			"o":"14.56800297",
			"c":"14.59766270",
			"h":"14.63325437",
			"l":"14.56207102",
			"v":"2.22524220",
			"n":33
	   }
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBLVTKlineServe("TRXDOWN", "1m", func(event *WsBLVTKlineEvent) {
		e := &WsBLVTKlineEvent{
			Event:  "kline",
			Time:   1600243159447,
			Symbol: "TRXDOWN",
			Kline: WsBLVTKline{
				StartTime:       1600243140000,
				CloseTime:       1600243199999,
				Symbol:          "TRXDOWN",
				Interval:        "1m",
				FirstUpdateTime: 1600243140484,
				LastUpdateTime:  1600243159424,
				OpenPrice:       "14.56800297",
				ClosePrice:      "14.59766270",
				HighPrice:       "14.63325437",
				LowPrice:        "14.56207102",
				Leverage:        "2.22524220",
				Count:           33,
			},
		}
		s.assertBLVTKlineEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertBLVTKlineEvent(e, a *WsBLVTKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.CloseTime, ak.CloseTime, "CloseTime")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstUpdateTime, ak.FirstUpdateTime, "FirstUpdateTime")
	r.Equal(ek.LastUpdateTime, ak.LastUpdateTime, "LastUpdateTime")
	r.Equal(ek.OpenPrice, ak.OpenPrice, "OpenPrice")
	r.Equal(ek.ClosePrice, ak.ClosePrice, "ClosePrice")
	r.Equal(ek.HighPrice, ak.HighPrice, "HighPrice")
	r.Equal(ek.LowPrice, ak.LowPrice, "LowPrice")
	r.Equal(ek.Leverage, ak.Leverage, "Leverage")
	r.Equal(ek.Count, ak.Count, "Count")
}

func (s *websocketServiceTestSuite) TestWsCompositiveIndexServe() {
	data := []byte(`{
		"e":"compositeIndex",
		"E":1602310596000,
		"s":"DEFIUSDT",
		"p":"554.41604065",
		"c":[
		  {
			  "b":"BAL",
			  "w":"1.35038833",
			  "W":"0.03957100"
		  },
		  {
			"b":"BAND",
			"w":"3.53782729",
			"W":"0.03935200"
		  }
		]
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCompositiveIndexServe("TRXDOWN", func(event *WsCompositeIndexEvent) {
		e := &WsCompositeIndexEvent{
			Event:  "compositeIndex",
			Time:   1602310596000,
			Symbol: "DEFIUSDT",
			Price:  "554.41604065",
			Composition: []WsComposition{
				{
					BaseAsset:    "BAL",
					WeightQty:    "1.35038833",
					WeighPercent: "0.03957100",
				},
				{
					BaseAsset:    "BAND",
					WeightQty:    "3.53782729",
					WeighPercent: "0.03935200",
				},
			},
		}
		s.assertCompositeIndexEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertCompositeIndexEvent(e, a *WsCompositeIndexEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
	for i, c := range e.Composition {
		r.Equal(c.BaseAsset, a.Composition[i].BaseAsset, "Position")
		r.Equal(c.WeightQty, a.Composition[i].WeightQty, "WeightQty")
		r.Equal(c.WeighPercent, a.Composition[i].WeighPercent, "WeighPercent")
	}
}

func (s *websocketServiceTestSuite) testWsUserDataServe(data []byte, expectedEvent *WsUserDataEvent) {
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsUserDataServe("fakeListenKey", func(event *WsUserDataEvent) {
		s.assertUserDataEvent(expectedEvent, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestWsUserDataServeStreamExpired() {
	data := []byte(`{
		"e": "listenKeyExpired",
		"E": 1576653824250
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "listenKeyExpired",
		Time:  1576653824250,
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeMarginCall() {
	data := []byte(`{
		"e":"MARGIN_CALL",
		"E":1587727187525,
		"cw":"3.16812045",
		"p":[
			{
				"s":"ETHUSDT",
				"ps":"LONG",
				"pa":"1.327",
				"mt":"CROSSED",
				"iw":"0",
				"mp":"187.17127",
				"up":"-1.166074",
				"mm":"1.614445"
			}
		]
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "MARGIN_CALL",
		Time:  1587727187525,
		WsUserDataMarginCall: WsUserDataMarginCall{CrossWalletBalance: "3.16812045",
			MarginCallPositions: []WsPosition{
				{
					Symbol:                    "ETHUSDT",
					Side:                      "LONG",
					Amount:                    "1.327",
					MarginType:                "CROSSED",
					IsolatedWallet:            "0",
					MarkPrice:                 "187.17127",
					UnrealizedPnL:             "-1.166074",
					MaintenanceMarginRequired: "1.614445",
				},
			}},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeAccountUpdate() {
	data := []byte(`{
		"e": "ACCOUNT_UPDATE",
		"E": 1564745798939,
		"T": 1564745798938,
		"a":
		  {
			"m":"ORDER",
			"B":[
			  {
				"a":"USDT",
				"wb":"122624.12345678",
				"cw":"100.12345678"
			  },
			  {
				"a":"BNB",
				"wb":"1.00000000",
				"cw":"0.00000000"
			  }
			],
			"P":[
			  {
				"s":"BTCUSDT",
				"pa":"0",
				"ep":"0.00000",
				"cr":"200",
				"up":"0",
				"mt":"isolated",
				"iw":"0.00000000",
				"ps":"BOTH"
			  },
			  {
				  "s":"BTCUSDT",
				  "pa":"20",
				  "ep":"6563.66500",
				  "cr":"0",
				  "up":"2850.21200",
				  "mt":"isolated",
				  "iw":"13200.70726908",
				  "ps":"LONG"
			   },
			  {
				  "s":"BTCUSDT",
				  "pa":"-10",
				  "ep":"6563.86000",
				  "cr":"-45.04000000",
				  "up":"-1423.15600",
				  "mt":"isolated",
				  "iw":"6570.42511771",
				  "ps":"SHORT"
			  }
			]
		  }
	}`)
	expectedEvent := &WsUserDataEvent{
		Event:           "ACCOUNT_UPDATE",
		Time:            1564745798939,
		TransactionTime: 1564745798938,
		WsUserDataAccountUpdate: WsUserDataAccountUpdate{
			AccountUpdate: WsAccountUpdate{
				Reason: "ORDER",
				Balances: []WsBalance{
					{
						Asset:              "USDT",
						Balance:            "122624.12345678",
						CrossWalletBalance: "100.12345678",
					},
					{
						Asset:              "BNB",
						Balance:            "1.00000000",
						CrossWalletBalance: "0.00000000",
					},
				},
				Positions: []WsPosition{
					{
						Symbol:              "BTCUSDT",
						Amount:              "0",
						EntryPrice:          "0.00000",
						AccumulatedRealized: "200",
						UnrealizedPnL:       "0",
						MarginType:          "isolated",
						IsolatedWallet:      "0.00000000",
						Side:                "BOTH",
					},
					{
						Symbol:              "BTCUSDT",
						Amount:              "20",
						EntryPrice:          "6563.66500",
						AccumulatedRealized: "0",
						UnrealizedPnL:       "2850.21200",
						MarginType:          "isolated",
						IsolatedWallet:      "13200.70726908",
						Side:                "LONG",
					},
					{
						Symbol:              "BTCUSDT",
						Amount:              "-10",
						EntryPrice:          "6563.86000",
						AccumulatedRealized: "-45.04000000",
						UnrealizedPnL:       "-1423.15600",
						MarginType:          "isolated",
						IsolatedWallet:      "6570.42511771",
						Side:                "SHORT",
					},
				},
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeOrderTradeUpdate() {
	data := []byte(`{
		"e":"ORDER_TRADE_UPDATE",
		"E":1568879465651,
		"T":1568879465650,
		"o":{
		  "s":"BTCUSDT",
		  "c":"TEST",
		  "S":"SELL",
		  "o":"TRAILING_STOP_MARKET",
		  "f":"GTC",
		  "q":"0.001",
		  "p":"0",
		  "ap":"0",
		  "sp":"7103.04",
		  "x":"NEW",
		  "X":"NEW",
		  "i":8886774,
		  "l":"0",
		  "z":"0",
		  "L":"0",
		  "N":"USDT",
		  "n":"0",
		  "T":1568879465651,
		  "t":0,
		  "b":"0",
		  "a":"9.91",
		  "m":false,
		  "R":false,
		  "wt":"CONTRACT_PRICE",
		  "ot":"TRAILING_STOP_MARKET",
		  "ps":"LONG",
		  "cp":false,
		  "AP":"7476.89",
		  "cr":"5.0",
		  "rp":"0"
		}
	}`)
	expectedEvent := &WsUserDataEvent{
		Event:           "ORDER_TRADE_UPDATE",
		Time:            1568879465651,
		TransactionTime: 1568879465650,
		WsUserDataOrderTradeUpdate: WsUserDataOrderTradeUpdate{
			OrderTradeUpdate: WsOrderTradeUpdate{
				Symbol:               "BTCUSDT",
				ClientOrderID:        "TEST",
				Side:                 "SELL",
				Type:                 "TRAILING_STOP_MARKET",
				TimeInForce:          "GTC",
				OriginalQty:          "0.001",
				OriginalPrice:        "0",
				AveragePrice:         "0",
				StopPrice:            "7103.04",
				ExecutionType:        "NEW",
				Status:               "NEW",
				ID:                   8886774,
				LastFilledQty:        "0",
				AccumulatedFilledQty: "0",
				LastFilledPrice:      "0",
				CommissionAsset:      "USDT",
				Commission:           "0",
				TradeTime:            1568879465651,
				TradeID:              0,
				BidsNotional:         "0",
				AsksNotional:         "9.91",
				IsMaker:              false,
				IsReduceOnly:         false,
				WorkingType:          "CONTRACT_PRICE",
				OriginalType:         "TRAILING_STOP_MARKET",
				PositionSide:         "LONG",
				IsClosingPosition:    false,
				ActivationPrice:      "7476.89",
				CallbackRate:         "5.0",
				RealizedPnL:          "0",
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeAccountConfigUpdate() {
	data := []byte(`{
		"e":"ACCOUNT_CONFIG_UPDATE",
		"E":1611646737479,
		"T":1611646737476,
		"ac":{
		"s":"BTCUSDT",
		"l":25
		}
	}`)
	expectedEvent := &WsUserDataEvent{
		Event:           "ACCOUNT_CONFIG_UPDATE",
		Time:            1611646737479,
		TransactionTime: 1611646737476,
		WsUserDataAccountConfigUpdate: WsUserDataAccountConfigUpdate{
			AccountConfigUpdate: WsAccountConfigUpdate{
				Symbol:   "BTCUSDT",
				Leverage: 25,
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeTradeLite() {
	data := []byte(`{
		"e":"TRADE_LITE",
		"E":1721895408092,
		"T":1721895408214,
		"s":"BTCUSDT",
		"q":"0.001",
		"p":"0",
		"m":false,
		"c":"z8hcUoOsqEdKMeKPSABslD",
		"S":"BUY",
		"L":"64089.20",
		"l":"0.040",
		"t":109100866,
		"i":8886774
	}`)

	expectedEvent := &WsUserDataEvent{
		Event:           "TRADE_LITE",
		Time:            1721895408092,
		TransactionTime: 1721895408214,
		WsUserDataTradeLite: WsUserDataTradeLite{
			Symbol:          "BTCUSDT",
			OriginalQty:     "0.001",
			OriginalPrice:   "0",
			IsMaker:         false,
			ClientOrderID:   "z8hcUoOsqEdKMeKPSABslD",
			Side:            "BUY",
			LastFilledPrice: "64089.20",
			LastFilledQty:   "0.040",
			TradeID:         109100866,
			OrderID:         8886774,
		},
	}

	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeAlgoUpdate() {
	data := []byte(`{
  "e":"ALGO_UPDATE",
  "E":1750515742303,
  "T":1750515742297,
  "o":{
    "caid":"Q5xaq5EGKgXXa0fD7fs0Ip",
    "aid":2148719,
    "at":"CONDITIONAL",
    "o":"TAKE_PROFIT",
    "s":"BNBUSDT",
    "S":"SELL",
    "ps":"BOTH",
    "f":"GTC",
    "q":"0.01",
    "X":"CANCELED",
    "ai":"",
    "ap": "0.00000",
    "aq": "0.00000",
    "act": "0",
    "tp":"750",
    "p":"750",
    "V":"EXPIRE_MAKER",
    "wt":"CONTRACT_PRICE",
    "pm":"NONE",
    "cp":false,
    "pP":false,
    "R":false,
    "tt":0,
    "gtd":0,
    "rm": "Reduce Only reject"
  }
}`)

	expectedEvent := &WsUserDataEvent{
		Event:           "ALGO_UPDATE",
		Time:            1750515742303,
		TransactionTime: 1750515742297,
		WsUserDataAlgoUpdate: WsUserDataAlgoUpdate{
			AlgoUpdate: WsAlgoUpdate{
				ClientAlgoID:     "Q5xaq5EGKgXXa0fD7fs0Ip",
				AlgoID:           2148719,
				AlgoType:         "CONDITIONAL",
				OrderType:        "TAKE_PROFIT",
				Symbol:           "BNBUSDT",
				Side:             "SELL",
				PositionSide:     "BOTH",
				TimeInForce:      "GTC",
				Quantity:         "0.01",
				AlgoStatus:       AlgoOrderStatusTypeCanceled,
				OrderID:          "",
				AvgPrice:         "0.00000",
				ExecutedQuantity: "0.00000",
				ActualOrderType:  "0",
				TriggerPrice:     "750",
				OrderPrice:       "750",
				STPMode:          "EXPIRE_MAKER",
				WorkingType:      "CONTRACT_PRICE",
				PriceMatchMode:   "NONE",
				CloseAll:         false,
				PriceProtection:  false,
				ReduceOnly:       false,
				TriggerTime:      0,
				GoodTillTime:     0,
				FailedReason:     "Reduce Only reject",
			},
		},
	}

	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) assertUserDataEvent(e, a *WsUserDataEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	for i, e := range e.MarginCallPositions {
		a := a.MarginCallPositions[i]
		s.assertPosition(e, a)
	}
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	s.assertAccountUpdate(e.AccountUpdate, a.AccountUpdate)
	s.assertOrderTradeUpdate(e.OrderTradeUpdate, a.OrderTradeUpdate)
	s.assertAccountConfigUpdate(e.AccountConfigUpdate, a.AccountConfigUpdate)
	s.assertTradeLite(e.WsUserDataTradeLite, a.WsUserDataTradeLite)
}

func (s *websocketServiceTestSuite) assertTradeLite(e, a WsUserDataTradeLite) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OriginalQty, a.OriginalQty, "OriginalQty")
	r.Equal(e.OriginalPrice, a.OriginalPrice, "OriginalPrice")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.LastFilledPrice, a.LastFilledPrice, "LastFilledPrice")
	r.Equal(e.LastFilledQty, a.LastFilledQty, "LastFilledQty")
	r.Equal(e.TradeID, a.TradeID, "TradeID")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
}

func (s *websocketServiceTestSuite) assertPosition(e, a WsPosition) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.MarginType, a.MarginType, "MarginType")
	r.Equal(e.IsolatedWallet, a.IsolatedWallet, "IsolatedWallet")
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.MarkPrice, a.MarkPrice, "MarkPrice")
	r.Equal(e.UnrealizedPnL, a.UnrealizedPnL, "UnrealizedPnL")
	r.Equal(e.AccumulatedRealized, a.AccumulatedRealized, "AccumulatedRealized")
	r.Equal(e.MaintenanceMarginRequired, a.MaintenanceMarginRequired, "MaintenanceMarginRequired")
}

func (s *websocketServiceTestSuite) assertAccountUpdate(e, a WsAccountUpdate) {
	r := s.r()
	r.Equal(e.Reason, a.Reason, "Reason")
	for i, e := range e.Balances {
		a := a.Balances[i]
		r.Equal(e.Asset, a.Asset, "Asset")
		r.Equal(e.Balance, a.Balance, "Balance")
		r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	}
	for i, e := range e.Positions {
		a := a.Positions[i]
		s.assertPosition(e, a)
	}
}

func (s *websocketServiceTestSuite) assertOrderTradeUpdate(e, a WsOrderTradeUpdate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.OriginalQty, a.OriginalQty, "OriginalQty")
	r.Equal(e.OriginalPrice, a.OriginalPrice, "OriginalPrice")
	r.Equal(e.AveragePrice, a.AveragePrice, "AveragePrice")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.ExecutionType, a.ExecutionType, "ExecutionType")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.LastFilledQty, a.LastFilledQty, "LastFilledQty")
	r.Equal(e.AccumulatedFilledQty, a.AccumulatedFilledQty, "AccumulatedFilledQty")
	r.Equal(e.LastFilledPrice, a.LastFilledPrice, "LastFilledPrice")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.TradeID, a.TradeID, "TradeID")
	r.Equal(e.BidsNotional, a.BidsNotional, "BidsNotional")
	r.Equal(e.AsksNotional, a.AsksNotional, "AsksNotional")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.IsReduceOnly, a.IsReduceOnly, "IsReduceOnly")
	r.Equal(e.WorkingType, a.WorkingType, "WorkingType")
	r.Equal(e.OriginalType, a.OriginalType, "OriginalType")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.IsClosingPosition, a.IsClosingPosition, "IsClosingPosition")
	r.Equal(e.ActivationPrice, a.ActivationPrice, "ActivationPrice")
	r.Equal(e.CallbackRate, a.CallbackRate, "CallbackRate")
	r.Equal(e.RealizedPnL, a.RealizedPnL, "RealizedPnL")
}

func (s *websocketServiceTestSuite) assertAccountConfigUpdate(e, a WsAccountConfigUpdate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
}
