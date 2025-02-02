package binance

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

func (s *websocketServiceTestSuite) TestPartialDepthServe() {
	data := []byte(`{
	  "lastUpdateId": 160,
	  "bids": [
	    [
	      "0.0024",
	      "10",
	      []
	    ]
	  ],
	  "asks": [
	    [
	      "0.0026",
	      "100",
	      []
	    ]
	  ]
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsPartialDepthServe("ETHBTC", "5", func(event *WsPartialDepthEvent) {
		e := &WsPartialDepthEvent{
			Symbol:       "ETHBTC",
			LastUpdateID: 160,
			Bids: []Bid{
				{
					Price:    "0.0024",
					Quantity: "10",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.0026",
					Quantity: "100",
				},
			},
		}
		s.assertWsPartialDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestPartialDepthServe100Ms() {
	data := []byte(`{
	  "lastUpdateId": 160,
	  "bids": [
	    [
	      "0.0024",
	      "10",
	      []
	    ]
	  ],
	  "asks": [
	    [
	      "0.0026",
	      "100",
	      []
	    ]
	  ]
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsPartialDepthServe100Ms("ETHBTC", "5", func(event *WsPartialDepthEvent) {
		e := &WsPartialDepthEvent{
			Symbol:       "ETHBTC",
			LastUpdateID: 160,
			Bids: []Bid{
				{
					Price:    "0.0024",
					Quantity: "10",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.0026",
					Quantity: "100",
				},
			},
		}
		s.assertWsPartialDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestCombinedPartialDepthServe() {
	data := []byte(`{
      "stream":"ethusdt@depth5",
      "data": {
	    "lastUpdateId": 160,
	    "bids": [
	      [
	        "0.0024",
	        "10",
	        []
	      ]
	    ],
	    "asks": [
	      [
	        "0.0026",
	        "100",
	        []
	      ]
	    ]
      }
	}`)
	symbolLevels := map[string]string{
		"BTCUSDT": "5",
		"ETHUSDT": "5",
	}
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()
	doneC, stopC, err := WsCombinedPartialDepthServe(symbolLevels, func(event *WsPartialDepthEvent) {
		e := &WsPartialDepthEvent{
			Symbol:       "ETHUSDT",
			LastUpdateID: 160,
			Bids: []Bid{
				{
					Price:    "0.0024",
					Quantity: "10",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.0026",
					Quantity: "100",
				},
			},
		}
		s.assertWsPartialDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsPartialDepthEventEqual(e, a *WsPartialDepthEvent) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "LastUpdateID")
	for i := 0; i < len(e.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i := 0; i < len(e.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestDepthServe() {
	data := []byte(`{
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "U": 7913452,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsDepthServe("ETHBTC", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:         "depthUpdate",
			Time:          1499404630606,
			Symbol:        "ETHBTC",
			LastUpdateID:  7913455,
			FirstUpdateID: 7913452,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestDepthServe100Ms() {
	data := []byte(`{
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "U": 7913452,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsDepthServe100Ms("ETHBTC", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:         "depthUpdate",
			Time:          1499404630606,
			Symbol:        "ETHBTC",
			LastUpdateID:  7913455,
			FirstUpdateID: 7913452,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsDepthEventEqual(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "UpdateID")
	r.Equal(e.FirstUpdateID, a.FirstUpdateID, "FirstUpdateID")
	for i := 0; i < len(e.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i := 0; i < len(e.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketServiceTestSuite) TestCombinedDepthServe() {
	data := []byte(`{
		"stream":"btcusdt@depth",
		"data":{
			"e":"depthUpdate",
			"E":1629769560797,
			"s":"BTCUSDT",
			"U":13544035,
			"u":13544037,
			"b":[["49095.23000000","0.01018500"],["49081.00000000","0.00000000"]],
			"a":[["49095.65000000","0.01018500"]]}}
	`)
	symbols := []string{
		"BTCUSDT",
		"ETHUSDT",
	}
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()
	doneC, stopC, err := WsCombinedDepthServe(symbols, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Symbol:        "BTCUSDT",
			Time:          1629769560797,
			LastUpdateID:  13544037,
			FirstUpdateID: 13544035,
			Bids: []Bid{
				{
					Price:    "49095.23000000",
					Quantity: "0.01018500",
				},
				{
					Price:    "49081.00000000",
					Quantity: "0.00000000",
				},
			},
			Asks: []Ask{
				{
					Price:    "49095.65000000",
					Quantity: "0.01018500",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestCombinedDepthServe100Ms() {
	data := []byte(`{
		"stream":"btcusdt@depth",
		"data":{
			"e":"depthUpdate",
			"E":1629769560797,
			"s":"BTCUSDT",
			"U":13544035,
			"u":13544037,
			"b":[["49095.23000000","0.01018500"],["49081.00000000","0.00000000"]],
			"a":[["49095.65000000","0.01018500"]]}}
	`)
	symbols := []string{
		"BTCUSDT",
		"ETHUSDT",
	}
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()
	doneC, stopC, err := WsCombinedDepthServe100Ms(symbols, func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Symbol:        "BTCUSDT",
			Time:          1629769560797,
			LastUpdateID:  13544037,
			FirstUpdateID: 13544035,
			Bids: []Bid{
				{
					Price:    "49095.23000000",
					Quantity: "0.01018500",
				},
				{
					Price:    "49081.00000000",
					Quantity: "0.00000000",
				},
			},
			Asks: []Ask{
				{
					Price:    "49095.65000000",
					Quantity: "0.01018500",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestKlineServe() {
	data := []byte(`{
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
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsKlineServe("ETHBTC", "1m", func(event *WsKlineEvent) {
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

func (s *websocketServiceTestSuite) TestWsAggTradeServe() {
	data := []byte(`{
        "e": "aggTrade",
        "E": 1499405254326,
        "s": "ETHBTC",
        "a": 70232,
        "p": "0.10281118",
        "q": "8.15632997",
        "f": 77489,
        "l": 77489,
        "T": 1499405254324,
        "m": false,
        "M": true
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAggTradeServe("ETHBTC", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:                 "aggTrade",
			Time:                  1499405254326,
			Symbol:                "ETHBTC",
			AggTradeID:            70232,
			Price:                 "0.10281118",
			Quantity:              "8.15632997",
			FirstBreakdownTradeID: 77489,
			LastBreakdownTradeID:  77489,
			TradeTime:             1499405254324,
			IsBuyerMaker:          false,
		}
		s.assertWsAggTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
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

func (s *websocketServiceTestSuite) TestWsCombinedAggTradeServe() {
	data := []byte(`{
	"stream":"ethbtc@aggTrade",
	"data": {
		"e": "aggTrade",
		"E": 1499405254326,
		"s": "ETHBTC",
		"a": 70232,
		"p": "0.10281118",
		"q": "8.15632997",
		"f": 77489,
		"l": 77489,
		"T": 1499405254324,
		"m": false,
		"M": true
		}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedAggTradeServe([]string{"ETHBTC"}, func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:                 "aggTrade",
			Time:                  1499405254326,
			Symbol:                "ETHBTC",
			AggTradeID:            70232,
			Price:                 "0.10281118",
			Quantity:              "8.15632997",
			FirstBreakdownTradeID: 77489,
			LastBreakdownTradeID:  77489,
			TradeTime:             1499405254324,
			IsBuyerMaker:          false,
		}
		s.assertWsAggTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAggTradeEventEqual(e, a *WsAggTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstBreakdownTradeID, a.FirstBreakdownTradeID, "FirstBreakdownTradeID")
	r.Equal(e.LastBreakdownTradeID, a.LastBreakdownTradeID, "LastBreakdownTradeID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *websocketServiceTestSuite) assertAccountUpdate(e, a *WsAccountUpdate) {
	r := s.r()
	r.Equal(e.Asset, a.Asset)
	r.Equal(e.Free, a.Free)
	r.Equal(e.Locked, a.Locked)
}

func (s *websocketServiceTestSuite) assertOrderUpdate(e, a *WsOrderUpdate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ClientOrderId, a.ClientOrderId, "ClientOrderId")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.IceBergVolume, a.IceBergVolume, "IceBergVolume")
	r.Equal(e.OrderListId, a.OrderListId, "OrderListId")
	r.Equal(e.OrigCustomOrderId, a.OrigCustomOrderId, "OrigCustomOrderId")
	r.Equal(e.ExecutionType, a.ExecutionType, "ExecutionType")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.RejectReason, a.RejectReason, "RejectReason")
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.LatestVolume, a.LatestVolume, "LatestVolume")
	r.Equal(e.FilledVolume, a.FilledVolume, "FilledVolume")
	r.Equal(e.LatestPrice, a.LatestPrice, "LatestPrice")
	r.Equal(e.FeeAsset, a.FeeAsset, "FeeAsset")
	r.Equal(e.FeeCost, a.FeeCost, "FeeCost")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.TradeId, a.TradeId, "TradeId")
	r.Equal(e.IgnoreI, a.IgnoreI, "IgnoreI")
	r.Equal(e.IsInOrderBook, a.IsInOrderBook, "IsInOrderBook")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.IgnoreM, a.IgnoreM, "IgnoreM")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.FilledQuoteVolume, a.FilledQuoteVolume, "FilledQuoteVolume")
	r.Equal(e.LatestQuoteVolume, a.LatestQuoteVolume, "LatestQuoteVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.SelfTradePreventionMode, a.SelfTradePreventionMode, "SelfTradePreventionMode")

	r.Equal(e.TrailingDelta, a.TrailingDelta, "TrailingDelta")
	r.Equal(e.TrailingTime, a.TrailingTime, "TrailingTime")
	r.Equal(e.StrategyId, a.StrategyId, "StrategyId")
	r.Equal(e.StrategyType, a.StrategyType, "StrategyType")
	r.Equal(e.PreventedMatchId, a.PreventedMatchId, "PreventedMatchId")
	r.Equal(e.PreventedQuantity, a.PreventedQuantity, "PreventedQuantity")
	r.Equal(e.LastPreventedQuantity, a.LastPreventedQuantity, "LastPreventedQuantity")
	r.Equal(e.TradeGroupId, a.TradeGroupId, "TradeGroupId")
	r.Equal(e.CounterOrderId, a.CounterOrderId, "CounterOrderId")
	r.Equal(e.CounterSymbol, a.CounterSymbol, "CounterSymbol")
	r.Equal(e.PreventedExecutionQuantity, a.PreventedExecutionQuantity, "PreventedExecutionQuantity")
	r.Equal(e.PreventedExecutionPrice, a.PreventedExecutionPrice, "PreventedExecutionPrice")
	r.Equal(e.PreventedExecutionQuoteQty, a.PreventedExecutionQuoteQty, "PreventedExecutionQuoteQty")
	r.Equal(e.WorkingTime, a.WorkingTime, "WorkingTime")
	r.Equal(e.MatchType, a.MatchType, "MatchType")
	r.Equal(e.AllocationId, a.AllocationId, "AllocationId")
	r.Equal(e.WorkingFloor, a.WorkingFloor, "WorkingFloor")
	r.Equal(e.UsedSor, a.UsedSor, "UsedSor")
}

func (s *websocketServiceTestSuite) assertBalanceUpdate(e, a *WsBalanceUpdate) {
	r := s.r()
	r.Equal(e.Asset, a.Asset)
	r.Equal(e.Change, a.Change)
	r.Equal(e.TransactionTime, a.TransactionTime)
}

func (s *websocketServiceTestSuite) assertUserDataEvent(e, a *WsUserDataEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	for i, e := range e.AccountUpdate.WsAccountUpdates {
		a := a.AccountUpdate.WsAccountUpdates[i]
		s.assertAccountUpdate(&e, &a)
	}
	s.assertOrderUpdate(&e.OrderUpdate, &a.OrderUpdate)
	s.assertBalanceUpdate(&e.BalanceUpdate, &a.BalanceUpdate)
}

func (s *websocketServiceTestSuite) testWsUserDataServe(data []byte, expectedEvent *WsUserDataEvent) {
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsUserDataServe("fakeListenKey", func(event *WsUserDataEvent) {
		s.assertUserDataEvent(expectedEvent, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestWsUserDataServeAccountUpdate() {
	data := []byte(`{
	   "e":"outboundAccountPosition",
	   "E":1629771130464,
	   "u":1629771130463,
	   "B":[
	      {
	         "a":"LTC",
	         "f":"503.70000000",
	         "l":"0.00000000"
	      }
	   ]
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "outboundAccountPosition",
		Time:  1629771130464,
		AccountUpdate: WsAccountUpdateList{
			1629771130463,
			[]WsAccountUpdate{
				{
					"LTC",
					"503.70000000",
					"0.00000000",
				},
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeOrderUpdateWithExample() {
	//Using example data from the API documentation
	data := []byte(`{
          "e":  "executionReport",
          "E":  1499405658658,
          "s":  "ETHBTC",
          "c":  "mUvoqJxFIILMdfAW5iGSOW",
          "S":  "BUY",
          "o":  "LIMIT",
          "f":  "GTC",
          "q":  "1.00000000",
          "p":  "0.10264410",
          "P":  "0.00000000",
          "F":  "0.00000000",
          "g":  -1,
          "C":  "",
          "x":  "NEW",
          "X":  "NEW",
          "r":  "NONE",
          "i":  4293153,
          "l":  "0.00000000",
          "z":  "0.00000000",
          "L":  "0.00000000",
          "n":  "0",
          "N":  null,
          "T":  1499405658657,
          "t":  -1,
          "I":  8641984,
          "w":  true,
          "m":  false,
          "M":  false,
          "O":  1499405658657,
          "Z":  "0.00000000",
          "Y":  "0.00000000",
          "Q":  "0.00000000",
          "V":  "NONE",
          "d":  4,
          "D":  1668680518494,
          "j":  1,
          "J":  1000000,
          "v":  3,
          "A":  "3.000000",
          "B":  "3.000000",
          "u":  1,
          "U":  37,
          "Cs":  "BTCUSDT",
          "pl":  "2.123456",
          "pL":  "0.10000001",
          "pY":  "0.21234562",
          "W":  1668683798379,
          "b":  "ONE_PARTY_TRADE_REPORT",
          "a":  1234,
          "k":  "SOR",
          "uS":  true
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "executionReport",
		Time:  1499405658658,
		OrderUpdate: WsOrderUpdate{
			Symbol:                     "ETHBTC",
			ClientOrderId:              "mUvoqJxFIILMdfAW5iGSOW",
			Side:                       "BUY",
			Type:                       "LIMIT",
			TimeInForce:                "GTC",
			Volume:                     "1.00000000",
			Price:                      "0.10264410",
			StopPrice:                  "0.00000000",
			IceBergVolume:              "0.00000000",
			OrderListId:                -1,
			OrigCustomOrderId:          "",
			ExecutionType:              "NEW",
			Status:                     "NEW",
			RejectReason:               "NONE",
			Id:                         4293153,
			LatestVolume:               "0.00000000",
			FilledVolume:               "0.00000000",
			LatestPrice:                "0.00000000",
			FeeAsset:                   "",
			FeeCost:                    "0",
			TransactionTime:            1499405658657,
			TradeId:                    -1,
			IgnoreI:                    8641984,
			IsInOrderBook:              true,
			IsMaker:                    false,
			IgnoreM:                    false,
			CreateTime:                 1499405658657,
			FilledQuoteVolume:          "0.00000000",
			LatestQuoteVolume:          "0.00000000",
			QuoteVolume:                "0.00000000",
			SelfTradePreventionMode:    "NONE",
			TrailingDelta:              4,
			TrailingTime:               1668680518494,
			StrategyId:                 1,
			StrategyType:               1000000,
			PreventedMatchId:           3,
			PreventedQuantity:          "3.000000",
			LastPreventedQuantity:      "3.000000",
			TradeGroupId:               1,
			CounterOrderId:             37,
			CounterSymbol:              "BTCUSDT",
			PreventedExecutionQuantity: "2.123456",
			PreventedExecutionPrice:    "0.10000001",
			PreventedExecutionQuoteQty: "0.21234562",
			WorkingTime:                1668683798379,
			MatchType:                  "ONE_PARTY_TRADE_REPORT",
			AllocationId:               1234,
			WorkingFloor:               "SOR",
			UsedSor:                    true,
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}
func (s *websocketServiceTestSuite) TestWsUserDataServeOrderUpdateWithPendingOrder() {
	//Use pending order data
	data := []byte(`{
          "e":  "executionReport",
          "E":  1709393629467,
          "s":  "FDUSDUSDT",
          "c":  "ios_2249e2aa2d394cada23b8e2e5b8ecdb9",
          "S":  "BUY",
          "o":  "LIMIT",
          "f":  "GTC",
          "q":  "1000.00000000",
          "p":  "0.99700000",
          "P":  "0.00000000",
          "F":  "0.00000000",
          "g":  -1,
          "C":  "",
          "x":  "NEW",
          "X":  "NEW",
          "r":  "NONE",
          "i":  98534826,
          "l":  "0.00000000",
          "z":  "0.00000000",
          "L":  "0.00000000",
          "n":  "0",
          "N":  null,
          "T":  1709393629467,
          "t":  -1,
          "I":  270913746,
          "w":  true,
          "m":  false,
          "M":  false,
          "O":  1709393629467,
          "Z":  "0.00000000",
          "Y":  "0.00000000",
          "Q":  "0.00000000",
          "W":  1709393629467,
          "V":  "EXPIRE_MAKER"
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "executionReport",
		Time:  1709393629467,
		OrderUpdate: WsOrderUpdate{
			Symbol:                  "FDUSDUSDT",
			ClientOrderId:           "ios_2249e2aa2d394cada23b8e2e5b8ecdb9",
			Side:                    "BUY",
			Type:                    "LIMIT",
			TimeInForce:             "GTC",
			Volume:                  "1000.00000000",
			Price:                   "0.99700000",
			StopPrice:               "0.00000000",
			IceBergVolume:           "0.00000000",
			OrderListId:             -1,
			OrigCustomOrderId:       "",
			ExecutionType:           "NEW",
			Status:                  "NEW",
			RejectReason:            "NONE",
			Id:                      98534826,
			LatestVolume:            "0.00000000",
			FilledVolume:            "0.00000000",
			LatestPrice:             "0.00000000",
			FeeAsset:                "",
			FeeCost:                 "0",
			TransactionTime:         1709393629467,
			TradeId:                 -1,
			IgnoreI:                 270913746,
			IsInOrderBook:           true,
			IsMaker:                 false,
			IgnoreM:                 false,
			CreateTime:              1709393629467,
			FilledQuoteVolume:       "0.00000000",
			LatestQuoteVolume:       "0.00000000",
			QuoteVolume:             "0.00000000",
			WorkingTime:             1709393629467,
			SelfTradePreventionMode: "EXPIRE_MAKER",
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsUserDataServeOrderUpdateWithTradeOrder() {
	//Use trade order data
	data := []byte(`{
          "e":  "executionReport",
          "E":  1709393640518,
          "s":  "FDUSDUSDT",
          "c":  "ios_54d9b18d8e7a4caf9d149573e16480ba",
          "S":  "BUY",
          "o":  "LIMIT",
          "f":  "GTC",
          "q":  "100.00000000",
          "p":  "0.99750000",
          "P":  "0.00000000",
          "F":  "0.00000000",
          "g":  -1,
          "C":  "",
          "x":  "TRADE",
          "X":  "FILLED",
          "r":  "NONE",
          "i":  98534864,
          "l":  "100.00000000",
          "z":  "100.00000000",
          "L":  "0.99750000",
          "n":  "0.00000000",
          "N":  "BNB",
          "T":  1709393640518,
          "t":  73839999,
          "I":  270913853,
          "w":  false,
          "m":  false,
          "M":  true,
          "O":  1709393640518,
          "Z":  "99.75000000",
          "Y":  "99.75000000",
          "Q":  "0.00000000",
          "W":  1709393640518,
          "V":  "EXPIRE_MAKER"
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "executionReport",
		Time:  1709393640518,
		OrderUpdate: WsOrderUpdate{
			Symbol:                  "FDUSDUSDT",
			ClientOrderId:           "ios_54d9b18d8e7a4caf9d149573e16480ba",
			Side:                    "BUY",
			Type:                    "LIMIT",
			TimeInForce:             "GTC",
			Volume:                  "100.00000000",
			Price:                   "0.99750000",
			StopPrice:               "0.00000000",
			IceBergVolume:           "0.00000000",
			OrderListId:             -1,
			OrigCustomOrderId:       "",
			ExecutionType:           "TRADE",
			Status:                  "FILLED",
			RejectReason:            "NONE",
			Id:                      98534864,
			LatestVolume:            "100.00000000",
			FilledVolume:            "100.00000000",
			LatestPrice:             "0.99750000",
			FeeAsset:                "BNB",
			FeeCost:                 "0.00000000",
			TransactionTime:         1709393640518,
			TradeId:                 73839999,
			IgnoreI:                 270913853,
			IsInOrderBook:           false,
			IsMaker:                 false,
			IgnoreM:                 true,
			CreateTime:              1709393640518,
			FilledQuoteVolume:       "99.75000000",
			LatestQuoteVolume:       "99.75000000",
			QuoteVolume:             "0.00000000",
			WorkingTime:             1709393640518,
			SelfTradePreventionMode: "EXPIRE_MAKER",
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketServiceTestSuite) TestWsMarketStatServe() {
	data := []byte(`{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "BNBBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"x": "0.0009",
  		"c": "0.0025",
  		"Q": "10",
  		"b": "0.0024",
  		"B": "10",
  		"a": "0.0026",
  		"A": "100",
  		"o": "0.0010",
  		"h": "0.0026",
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

	doneC, stopC, err := WsMarketStatServe("BNBBTC", func(event *WsMarketStatEvent) {
		e := &WsMarketStatEvent{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BNBBTC",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			PrevClosePrice:     "0.0009",
			LastPrice:          "0.0025",
			CloseQty:           "10",
			BidPrice:           "0.0024",
			BidQty:             "10",
			AskPrice:           "0.0026",
			AskQty:             "100",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0026",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			Count:              18151,
		}
		s.assertWsMarketStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestWsCombinedMarketStatServe() {
	data := []byte(`{
	"stream":"bnbbtc@ticker",
	"data": {
		"e": "24hrTicker",
		"E": 123456789,
		"s": "BNBBTC",
		"p": "0.0015",
		"P": "250.00",
		"w": "0.0018",
		"x": "0.0009",
		"c": "0.0025",
		"Q": "10",
		"b": "0.0024",
		"B": "10",
		"a": "0.0026",
		"A": "100",
		"o": "0.0010",
		"h": "0.0026",
		"l": "0.0010",
		"v": "10000",
		"q": "18",
	  "O": 0,
		"C": 86400000,
		"F": 0,
		"L": 18150,
		"n": 18151
	}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedMarketStatServe([]string{"BNBBTC"}, func(event *WsMarketStatEvent) {
		e := &WsMarketStatEvent{
			Event:              "24hrTicker",
			Time:               123456789,
			Symbol:             "BNBBTC",
			PriceChange:        "0.0015",
			PriceChangePercent: "250.00",
			WeightedAvgPrice:   "0.0018",
			PrevClosePrice:     "0.0009",
			LastPrice:          "0.0025",
			CloseQty:           "10",
			BidPrice:           "0.0024",
			BidQty:             "10",
			AskPrice:           "0.0026",
			AskQty:             "100",
			OpenPrice:          "0.0010",
			HighPrice:          "0.0026",
			LowPrice:           "0.0010",
			BaseVolume:         "10000",
			QuoteVolume:        "18",
			OpenTime:           0,
			CloseTime:          86400000,
			FirstID:            0,
			LastID:             18150,
			Count:              18151,
		}
		s.assertWsMarketStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsMarketStatEventEqual(e, a *WsMarketStatEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.PrevClosePrice, a.PrevClosePrice, "PrevClosePrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.CloseQty, a.CloseQty, "CloseQty")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.BidQty, a.BidQty, "BidQty")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.AskQty, a.AskQty, "AskQty")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}

func (s *websocketServiceTestSuite) TestWsAllMarketsStatServe() {
	data := []byte(`[{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "BNBBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"x": "0.0009",
  		"c": "0.0025",
  		"Q": "10",
  		"b": "0.0024",
  		"B": "10",
  		"a": "0.0026",
  		"A": "100",
  		"o": "0.0010",
  		"h": "0.0026",
  		"l": "0.0010",
  		"v": "10000",
  		"q": "18",
 		"O": 0,
  		"C": 86400000,
  		"F": 0,
  		"L": 18150,
  		"n": 18151
	},{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "ETHBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"x": "0.0009",
  		"c": "0.0025",
  		"Q": "10",
  		"b": "0.0024",
  		"B": "10",
  		"a": "0.0026",
  		"A": "100",
  		"o": "0.0010",
  		"h": "0.0026",
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

	doneC, stopC, err := WsAllMarketsStatServe(func(event WsAllMarketsStatEvent) {
		e := WsAllMarketsStatEvent{
			&WsMarketStatEvent{
				Event:              "24hrTicker",
				Time:               123456789,
				Symbol:             "BNBBTC",
				PriceChange:        "0.0015",
				PriceChangePercent: "250.00",
				WeightedAvgPrice:   "0.0018",
				PrevClosePrice:     "0.0009",
				LastPrice:          "0.0025",
				CloseQty:           "10",
				BidPrice:           "0.0024",
				BidQty:             "10",
				AskPrice:           "0.0026",
				AskQty:             "100",
				OpenPrice:          "0.0010",
				HighPrice:          "0.0026",
				LowPrice:           "0.0010",
				BaseVolume:         "10000",
				QuoteVolume:        "18",
				OpenTime:           0,
				CloseTime:          86400000,
				FirstID:            0,
				LastID:             18150,
				Count:              18151,
			},
			&WsMarketStatEvent{
				Event:              "24hrTicker",
				Time:               123456789,
				Symbol:             "ETHBTC",
				PriceChange:        "0.0015",
				PriceChangePercent: "250.00",
				WeightedAvgPrice:   "0.0018",
				PrevClosePrice:     "0.0009",
				LastPrice:          "0.0025",
				CloseQty:           "10",
				BidPrice:           "0.0024",
				BidQty:             "10",
				AskPrice:           "0.0026",
				AskQty:             "100",
				OpenPrice:          "0.0010",
				HighPrice:          "0.0026",
				LowPrice:           "0.0010",
				BaseVolume:         "10000",
				QuoteVolume:        "18",
				OpenTime:           0,
				CloseTime:          86400000,
				FirstID:            0,
				LastID:             18150,
				Count:              18151,
			},
		}
		s.assertWsAllMarketsStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAllMarketsStatEventEqual(e, a WsAllMarketsStatEvent) {
	for i := range e {
		s.assertWsMarketStatEventEqual(e[i], a[i])
	}
}

func (s *websocketServiceTestSuite) TestWsTradeServe() {
	data := []byte(`{
        "e": "trade",
        "E": 123456789,
        "s": "BNBBTC",
        "t": 12345,
        "p": "0.001",
        "q": "100",
        "b": 88,
        "a": 50,
        "T": 123456785,
        "m": true,
        "M": true
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsTradeServe("BNBBTC", func(event *WsTradeEvent) {
		e := &WsTradeEvent{
			Event:         "trade",
			Time:          123456789,
			Symbol:        "BNBBTC",
			TradeID:       12345,
			Price:         "0.001",
			Quantity:      "100",
			BuyerOrderID:  88,
			SellerOrderID: 50,
			TradeTime:     123456785,
			IsBuyerMaker:  true,
		}
		s.assertWsTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsTradeEventEqual(e, a *WsTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TradeID, a.TradeID, "TradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.BuyerOrderID, a.BuyerOrderID, "BuyerOrderID")
	r.Equal(e.SellerOrderID, a.SellerOrderID, "SellerOrderID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *websocketServiceTestSuite) assertWsCombinedTradeEventEqual(e, a *WsCombinedTradeEvent) {
	r := s.r()
	r.Equal(e.Stream, a.Stream, "Stream")
	s.assertWsTradeEventEqual(&e.Data, &a.Data)
}

func (s *websocketServiceTestSuite) TestWsCombinedTradeServe() {
	data := []byte(`{
		"stream": "bnbbtc@trade",
		"data": {
			"e": "trade",
			"E": 123456789,
			"s": "BNBBTC",
			"t": 12345,
			"p": "0.001",
			"q": "100",
			"b": 88,
			"a": 50,
			"T": 123456785,
			"m": true,
			"M": true
		}
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsCombinedTradeServe([]string{"BNBBTC"}, func(event *WsCombinedTradeEvent) {
		e := &WsCombinedTradeEvent{
			Stream: "bnbbtc@trade",
			Data: WsTradeEvent{
				Event:         "trade",
				Time:          123456789,
				Symbol:        "BNBBTC",
				TradeID:       12345,
				Price:         "0.001",
				Quantity:      "100",
				BuyerOrderID:  88,
				SellerOrderID: 50,
				TradeTime:     123456785,
				IsBuyerMaker:  true,
			},
		}

		s.assertWsCombinedTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) TestWsAllMiniMarketsStatServe() {
	data := []byte(`[{
  		"e": "24hrMiniTicker",
    	"E": 1523658017154,
    	"s": "BNBBTC",
   	 	"c": "0.00175640",
    	"o": "0.00161200",
    	"h": "0.00176000",
    	"l": "0.00159370",
    	"v": "3479863.89000000",
    	"q": "5725.90587704"
	},{
  		"e": "24hrMiniTicker",
    	"E": 1523658017133,
    	"s": "BNBETH",
    	"c": "0.02827000",
    	"o": "0.02628100",
    	"h": "0.02830300",
    	"l": "0.02469400",
    	"v": "456266.78000000",
    	"q": "11873.11095682"
	}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllMiniMarketsStatServe(func(event WsAllMiniMarketsStatEvent) {
		e := WsAllMiniMarketsStatEvent{
			&WsMiniMarketsStatEvent{
				Event:       "24hrMiniTicker",
				Time:        1523658017154,
				Symbol:      "BNBBTC",
				LastPrice:   "0.00175640",
				OpenPrice:   "0.00161200",
				HighPrice:   "0.00176000",
				LowPrice:    "0.00159370",
				BaseVolume:  "3479863.89000000",
				QuoteVolume: "5725.90587704",
			},
			&WsMiniMarketsStatEvent{
				Event:       "24hrMiniTicker",
				Time:        1523658017133,
				Symbol:      "BNBETH",
				LastPrice:   "0.02827000",
				OpenPrice:   "0.02628100",
				HighPrice:   "0.02830300",
				LowPrice:    "0.02469400",
				BaseVolume:  "456266.78000000",
				QuoteVolume: "11873.11095682",
			},
		}
		s.assertWsAllMiniMarketsStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketServiceTestSuite) assertWsAllMiniMarketsStatEventEqual(e, a WsAllMiniMarketsStatEvent) {
	for i := range e {
		s.assertWsMiniMarketsStatEventEqual(e[i], a[i])
	}
}

func (s *websocketServiceTestSuite) assertWsMiniMarketsStatEventEqual(e, a *WsMiniMarketsStatEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
}

// https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-book-ticker-streams
func (s *websocketServiceTestSuite) TestBookTickerServe() {
	data := []byte(`{
  		"u":17242169,
  		"s":"BTCUSD_200626",
  		"b":"9548.1",
  		"B":"52",
  		"a":"9548.5",
  		"A":"11"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsBookTickerServe("BTCUSD_200626", func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			UpdateID:     17242169,
			Symbol:       "BTCUSD_200626",
			BestBidPrice: "9548.1",
			BestBidQty:   "52",
			BestAskPrice: "9548.5",
			BestAskQty:   "11",
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

// https://binance-docs.github.io/apidocs/spot/en/#all-book-tickers-stream
func (s *websocketServiceTestSuite) TestAllBookTickerServe() {
	data := []byte(`{
  		"u":17242169,
  		"s":"BTCUSD_200626",
  		"b":"9548.1",
  		"B":"52",
  		"a":"9548.5",
  		"A":"11"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := WsAllBookTickerServe(func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			UpdateID:     17242169,
			Symbol:       "BTCUSD_200626",
			BestBidPrice: "9548.1",
			BestBidQty:   "52",
			BestAskPrice: "9548.5",
			BestAskQty:   "11",
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
	r.Equal(e.UpdateID, a.UpdateID, "UpdateID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BestBidPrice, a.BestBidPrice, "BestBidPrice")
	r.Equal(e.BestBidQty, a.BestBidQty, "BestBidQty")
	r.Equal(e.BestAskPrice, a.BestAskPrice, "BestAskPrice")
	r.Equal(e.BestAskQty, a.BestAskQty, "BestAskQty")
}
