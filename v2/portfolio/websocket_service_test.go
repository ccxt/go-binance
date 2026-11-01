package portfolio

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

func (s *websocketServiceTestSuite) testWsUserDataServe(data []byte, expectedEvent *WsUserDataEvent) {
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	// Create a handler that implements WsUserDataHandler interface
	handler := &testWsUserDataHandler{
		suite:         s,
		expectedEvent: expectedEvent,
	}

	doneC, stopC, err := WsUserDataServe("fakeListenKey", handler,
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

// testWsUserDataHandler implements WsUserDataHandler interface
type testWsUserDataHandler struct {
	suite         *websocketServiceTestSuite
	expectedEvent *WsUserDataEvent
}

func (h *testWsUserDataHandler) HandleListenKeyExpired(event *WsListenKeyExpired) {
}

func (h *testWsUserDataHandler) HandleMarginBalanceUpdate(event *WsMarginBalanceUpdate) {
}

func (h *testWsUserDataHandler) HandleRiskLevelChange(event *WsRiskLevelChange) {
	// Implement if needed
}

func (h *testWsUserDataHandler) HandleFuturesAccountConfigUpdate(event *WsFuturesAccountConfigUpdate) {
	if h.expectedEvent.Event == "ACCOUNT_CONFIG_UPDATE" {
		h.suite.assertAccountConfigUpdate(h.expectedEvent.WsUserDataAccountConfigUpdate.AccountConfigUpdate, event.AccountConfig)
	}
}

func (h *testWsUserDataHandler) HandleFuturesAccountUpdate(event *WsFuturesAccountUpdate) {
	if h.expectedEvent.Event == "ACCOUNT_UPDATE" {
		h.suite.assertAccountUpdate(h.expectedEvent.WsUserDataAccountUpdate.AccountUpdate, *event)
	}
}

func (h *testWsUserDataHandler) HandleFuturesOrderUpdate(event *WsFuturesOrderUpdate) {
	if h.expectedEvent.Event == "ORDER_TRADE_UPDATE" {
		h.suite.assertOrderTradeUpdate(h.expectedEvent.WsUserDataOrderTradeUpdate.OrderTradeUpdate, event.Order)
	}
}

func (h *testWsUserDataHandler) HandleMarginOrderUpdate(event *WsMarginOrderUpdate) {
	// Implement if needed
}

func (h *testWsUserDataHandler) HandleLiabilityUpdate(event *WsLiabilityUpdate) {
	// Implement if needed
}

func (h *testWsUserDataHandler) HandleMarginAccountUpdate(event *WsMarginAccountUpdate) {
	// Implement if needed
}

func (h *testWsUserDataHandler) HandleOpenOrderLossUpdate(event *WsOpenOrderLossUpdate) {
	// Implement if needed
}

func (h *testWsUserDataHandler) HandleConditionalOrderTradeUpdate(event *WsConditionalOrderTradeUpdate) {
	// Implement if needed
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
				"ep":"187.17127",
				"cr":"0",
				"pa":"1.327",
				"up":"-1.166074",
				"bep":152.909412
			}
		]
	}`)
	expectedEvent := &WsUserDataEvent{
		Event: "MARGIN_CALL",
		Time:  1587727187525,
		WsUserDataMarginCall: WsUserDataMarginCall{CrossWalletBalance: "3.16812045",
			MarginCallPositions: []WsPosition{
				{
					Symbol:              "ETHUSDT",
					Side:                "LONG",
					EntryPrice:          "187.17127",
					AccumulatedRealized: "0",
					Amount:              "1.327",
					UnrealizedPnL:       "-1.166074",
					BreakEvenPrice:      152.909412,
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
						Side:                "BOTH",
					},
					{
						Symbol:              "BTCUSDT",
						Amount:              "20",
						EntryPrice:          "6563.66500",
						AccumulatedRealized: "0",
						UnrealizedPnL:       "2850.21200",
						Side:                "LONG",
					},
					{
						Symbol:              "BTCUSDT",
						Amount:              "-10",
						EntryPrice:          "6563.86000",
						AccumulatedRealized: "-45.04000000",
						UnrealizedPnL:       "-1423.15600",
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
				PositionSide:         "LONG",
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

func (s *websocketServiceTestSuite) assertPosition(e WsPosition, a WsFuturesPosition) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Side, a.PositionSide, "Side")
	r.Equal(e.Amount, a.PositionAmount, "Amount")
	r.Equal(e.EntryPrice, a.EntryPrice, "EntryPrice")
	r.Equal(e.UnrealizedPnL, a.UnrealizedPnL, "UnrealizedPnL")
	r.Equal(e.AccumulatedRealized, a.AccumulatedRealized, "AccumulatedRealized")
}

func (s *websocketServiceTestSuite) assertAccountUpdate(e WsAccountUpdate, a WsFuturesAccountUpdate) {
	r := s.r()
	r.Equal(e.Reason, UserDataEventReasonType(a.AccountData.ReasonType), "Reason")
	for i, e := range e.Balances {
		a := a.AccountData.Balances[i]
		r.Equal(e.Asset, a.Asset, "Asset")
		r.Equal(e.Balance, a.WalletBalance, "Balance")
		r.Equal(e.CrossWalletBalance, a.CrossWalletBalance, "CrossWalletBalance")
	}
	for i, e := range e.Positions {
		a := a.AccountData.Positions[i]
		s.assertPosition(e, a)
	}
}

func (s *websocketServiceTestSuite) assertOrderTradeUpdate(e WsOrderTradeUpdate, a WsFuturesOrderData) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.ClientOrderID, a.ClientOrderID, "ClientOrderID")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.Type, a.OrderType, "Type")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.OriginalQty, a.OriginalQty, "OriginalQty")
	r.Equal(e.OriginalPrice, a.OriginalPrice, "OriginalPrice")
	r.Equal(e.AveragePrice, a.AveragePrice, "AveragePrice")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.ExecutionType, a.ExecutionType, "ExecutionType")
	r.Equal(e.Status, a.OrderStatus, "Status")
	r.Equal(e.ID, a.OrderID, "ID")
	r.Equal(e.LastFilledQty, a.LastFilledQty, "LastFilledQty")
	r.Equal(e.AccumulatedFilledQty, a.FilledAccumQty, "AccumulatedFilledQty")
	r.Equal(e.LastFilledPrice, a.LastFilledPrice, "LastFilledPrice")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.TradeID, a.TradeID, "TradeID")
	r.Equal(e.BidsNotional, a.BidsNotional, "BidsNotional")
	r.Equal(e.AsksNotional, a.AskNotional, "AsksNotional")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.IsReduceOnly, a.IsReduceOnly, "IsReduceOnly")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
}

func (s *websocketServiceTestSuite) assertAccountConfigUpdate(e, a WsAccountConfigUpdate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Leverage, a.Leverage, "Leverage")
}
