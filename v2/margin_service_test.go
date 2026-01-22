package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginTestSuite struct {
	baseTestSuite
}

func TestMarginAccountService(t *testing.T) {
	suite.Run(t, new(marginTestSuite))
}

func (s *marginTestSuite) TestTransfer() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	transferType := MarginTransferTypeToMargin
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
			"type":   transferType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginTransferService().Asset(asset).
		Amount(amount).Type(transferType).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) assertTransactionResponseEqual(a, e *TransactionResponse) {
	s.r().Equal(a.TranID, e.TranID, "TranID")
}

func (s *marginTestSuite) TestLoan() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginLoanService().Asset(asset).
		Amount(amount).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestRepay() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginRepayService().Asset(asset).
		Amount(amount).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestBorrowRepayBorrow() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	_type := MarginAccountBorrow
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":      asset,
			"amount":     amount,
			"isIsolated": false,
			"symbol":     "",
			"type":       string(_type),
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginBorrowRepayService().
		Asset(asset).
		Amount(amount).
		Type(_type).
		Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestBorrowRepayRepay() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	_type := MarginAccountRepay
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":      asset,
			"amount":     amount,
			"isIsolated": false,
			"symbol":     "",
			"type":       string(_type),
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginBorrowRepayService().
		Asset(asset).
		Amount(amount).
		Type(_type).
		Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestListBorrowRepay() {
	data := []byte(`{
		  "rows": [
			  {
				"type": "AUTO",
				"isolatedSymbol": "BNBUSDT",
				"amount": "14.00000000",
				"asset": "BNB",   
				"interest": "0.01866667",
				"principal": "13.98133333",
				"status": "CONFIRMED",
				"timestamp": 1563438204000,
				"txId": 2970933056
			  }
		  ],
		  "total": 1
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	isolatedSymbol := "BTCUSDT"
	txId := int64(100000001)
	startTime := int64(1555056425000)
	endTime := int64(1555056425001)
	current := int64(1)
	size := int64(10)
	_type := MarginAccountRepay
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":          asset,
			"isolatedSymbol": isolatedSymbol,
			"txId":           txId,
			"startTime":      startTime,
			"endTime":        endTime,
			"current":        current,
			"size":           size,
			"type":           _type,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListMarginBorrowRepayService().
		Asset(asset).
		IsolatedSymbol(isolatedSymbol).
		TxId(txId).
		StartTime(startTime).
		EndTime(endTime).
		Current(current).
		Size(size).
		Type(_type).
		Do(newContext())
	s.r().NoError(err)
	e := &MarginBorrowRepayResponse{
		Total: 1,
		Rows: []MarginBorrowRepay{
			{
				Type:           "AUTO",
				IsolatedSymbol: "BNBUSDT",
				Amount:         "14.00000000",
				Asset:          "BNB",
				Interest:       "0.01866667",
				Principal:      "13.98133333",
				Status:         "CONFIRMED",
				Timestamp:      1563438204000,
				TxID:           2970933056,
			},
		},
	}
	s.assertListBorrowRepayResponseEqual(e, res)
}

func (s *marginTestSuite) assertListBorrowRepayResponseEqual(e, a *MarginBorrowRepayResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(a.Rows, len(e.Rows), "Rows")
	for i := 0; i < len(e.Rows); i++ {
		s.assertListBorrowRepayItemEqual(&e.Rows[i], &a.Rows[i])
	}
}

func (s *marginTestSuite) assertListBorrowRepayItemEqual(e, a *MarginBorrowRepay) {
	r := s.r()
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.IsolatedSymbol, a.IsolatedSymbol, "IsolatedSymbol")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.Status, MarginAccountBorrowRepayStatusConfirmed, "Status")

	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.TxID, a.TxID, "TxID")
}

func (s *marginTestSuite) TestListMarginLoans() {
	data := []byte(`{
		"rows": [
		  {
			"asset": "BNB",
			"principal": "0.84624403",
			"timestamp": 1555056425000,
			"status": "CONFIRMED"
		  }
		],
		"total": 1
	  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BNB"
	txID := int64(1)
	startTime := int64(1555056425000)
	endTime := int64(1555056425001)
	current := int64(1)
	size := int64(10)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":     asset,
			"txId":      txID,
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListMarginLoansService().Asset(asset).
		TxID(txID).StartTime(startTime).EndTime(endTime).
		Current(current).Size(size).Do(newContext())
	s.r().NoError(err)
	e := &MarginLoanResponse{
		Rows: []MarginLoan{
			{
				Asset:     asset,
				Principal: "0.84624403",
				Timestamp: 1555056425000,
				Status:    MarginLoanStatusTypeConfirmed,
			},
		},
		Total: 1,
	}
	s.assertMarginLoanResponseEqual(e, res)
}

func (s *marginTestSuite) assertMarginLoanResponseEqual(e, a *MarginLoanResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(a.Rows, len(e.Rows), "Rows")
	for i := 0; i < len(e.Rows); i++ {
		s.assertMarginLoanEqual(&e.Rows[i], &a.Rows[i])
	}
}

func (s *marginTestSuite) assertMarginLoanEqual(e, a *MarginLoan) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Status, a.Status, "Status")
}

func (s *marginTestSuite) TestListMarginRepays() {
	data := []byte(`{
		"rows": [
			{
				"amount": "14.00000000",
				"asset": "BNB",
				"interest": "0.01866667",
				"principal": "13.98133333",
				"status": "CONFIRMED",
				"timestamp": 1563438204000,
				"txId": 2970933056
			}
		],
		"total": 1
   	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BNB"
	txID := int64(1)
	startTime := int64(1563438204000)
	endTime := int64(1563438204001)
	current := int64(1)
	size := int64(10)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":     asset,
			"txId":      txID,
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListMarginRepaysService().Asset(asset).
		TxID(txID).StartTime(startTime).EndTime(endTime).
		Current(current).Size(size).Do(newContext())
	s.r().NoError(err)
	e := &MarginRepayResponse{
		Rows: []MarginRepay{
			{
				Asset:     asset,
				Amount:    "14.00000000",
				Interest:  "0.01866667",
				Principal: "13.98133333",
				Timestamp: 1563438204000,
				Status:    MarginRepayStatusTypeConfirmed,
				TxID:      2970933056,
			},
		},
		Total: 1,
	}
	s.assertMarginRepayResponseEqual(e, res)
}

func (s *marginTestSuite) assertMarginRepayResponseEqual(e, a *MarginRepayResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(a.Rows, len(e.Rows), "Rows")
	for i := 0; i < len(e.Rows); i++ {
		s.assertMarginRepayEqual(&e.Rows[i], &a.Rows[i])
	}
}

func (s *marginTestSuite) assertMarginRepayEqual(e, a *MarginRepay) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TxID, a.TxID, "TxID")
}

func (s *marginTestSuite) TestGetMarginAccount() {
	data := []byte(`{
		"created": true,
		"borrowEnabled": true,
		"marginLevel": "11.64405625",
		"collateralMarginLevel": "0.64405625",
		"totalAssetOfBtc": "6.82728457",
		"totalLiabilityOfBtc": "0.58633215",
		"totalNetAssetOfBtc": "6.24095242",
		"totalCollateralValueInUSDT": "1234.33",
		"tradeEnabled": true,
		"transferInEnabled": true,
		"transferOutEnabled": true,
		"accountType": "MARGIN_1",
		"userAssets": [
			{
				"asset": "BTC",
				"borrowed": "0.00000000",
				"free": "0.00499500",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.00499500"
			},
			{
				"asset": "BNB",
				"borrowed": "201.66666672",
				"free": "2346.50000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "2144.83333328"
			},
			{
				"asset": "ETH",
				"borrowed": "0.00000000",
				"free": "0.00000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.00000000"
			},
			{
				"asset": "USDT",
				"borrowed": "0.00000000",
				"free": "0.00000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.00000000"
			}
		]
  	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetMarginAccountService().Do(newContext())
	s.r().NoError(err)
	e := &MarginAccount{
		Created:                    true,
		BorrowEnabled:              true,
		MarginLevel:                "11.64405625",
		CollateralMarginLevel:      "0.64405625",
		TotalAssetOfBTC:            "6.82728457",
		TotalLiabilityOfBTC:        "0.58633215",
		TotalNetAssetOfBTC:         "6.24095242",
		TotalCollateralValueInUSDT: "1234.33",
		TradeEnabled:               true,
		TransferInEnabled:          true,
		TransferOutEnabled:         true,
		AccountType:                "MARGIN_1",
		UserAssets: []UserAsset{
			{
				Asset:    "BTC",
				Borrowed: "0.00000000",
				Free:     "0.00499500",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00499500",
			},
			{
				Asset:    "BNB",
				Borrowed: "201.66666672",
				Free:     "2346.50000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "2144.83333328",
			},
			{
				Asset:    "ETH",
				Borrowed: "0.00000000",
				Free:     "0.00000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00000000",
			},
			{
				Asset:    "USDT",
				Borrowed: "0.00000000",
				Free:     "0.00000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00000000",
			},
		},
	}
	s.assertMarginAccountEqual(e, res)
}

func (s *marginTestSuite) TestGetIsolatedMarginAccount() {
	data := []byte(`{
      "assets": [
      	{
        "baseAsset": {
          "asset": "BTC",
          "borrowEnabled": true,
          "borrowed": "0.00000000",
          "free": "0.00000000",
          "interest": "0.00000000",
          "locked": "0.00000000",
          "netAsset": "0.00000000",
          "netAssetOfBtc": "0.00000000",
          "repayEnabled": true,
          "totalAsset": "0.00000000"
        },
        "quoteAsset": {
          "asset": "USDT",
          "borrowEnabled": true,
          "borrowed": "0.00000000",
          "free": "0.00000000",
          "interest": "0.00000000",
          "locked": "0.00000000",
          "netAsset": "0.00000000",
          "netAssetOfBtc": "0.00000000",
          "repayEnabled": true,
          "totalAsset": "0.00000000"
        },
        "symbol": "BTCUSDT",
        "isolatedCreated": true, 
        "enabled": true,
        "marginLevel": "0.00000000", 
        "marginLevelStatus": "EXCESSIVE",
        "marginRatio": "0.00000000",
        "indexPrice": "10000.00000000",
        "liquidatePrice": "1000.00000000",
        "liquidateRate": "1.00000000",
        "tradeEnabled": true
      }
    ],
    "totalAssetOfBtc": "0.00000000",
    "totalLiabilityOfBtc": "0.00000000",
    "totalNetAssetOfBtc": "0.00000000" 
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetIsolatedMarginAccountService().Do(newContext())
	s.r().NoError(err)

	e := &IsolatedMarginAccount{
		TotalAssetOfBTC:     "0.00000000",
		TotalLiabilityOfBTC: "0.00000000",
		TotalNetAssetOfBTC:  "0.00000000",
		Assets: []IsolatedMarginAsset{
			{
				Symbol:            "BTCUSDT",
				IsolatedCreated:   true,
				Enabled:           true,
				MarginLevel:       "0.00000000",
				MarginLevelStatus: "EXCESSIVE",
				MarginRatio:       "0.00000000",
				IndexPrice:        "10000.00000000",
				LiquidatePrice:    "1000.00000000",
				LiquidateRate:     "1.00000000",
				TradeEnabled:      true,
			},
		},
	}
	s.assertIsolatedMarginAccountEqual(e, res)
}

func (s *marginTestSuite) assertIsolatedMarginAccountEqual(e, a *IsolatedMarginAccount) {
	r := s.r()
	r.Equal(e.TotalAssetOfBTC, a.TotalAssetOfBTC, "TotalAssetOfBTC")
	r.Equal(e.TotalNetAssetOfBTC, a.TotalNetAssetOfBTC, "TotalNetAssetOfBTC")
	r.Equal(e.TotalLiabilityOfBTC, a.TotalLiabilityOfBTC, "TotalLiabilityOfBTC")
	for i := 0; i < len(a.Assets); i++ {
		s.assertIsolatedMarginAssetEqual(e.Assets[i], a.Assets[i])
	}
}

func (s *marginTestSuite) assertIsolatedMarginAssetEqual(e, a IsolatedMarginAsset) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.IsolatedCreated, a.IsolatedCreated, "IsolatedCreated")
	r.Equal(e.Enabled, a.Enabled, "Enabled")
	r.Equal(e.IndexPrice, a.IndexPrice, "IndexPrice")
	r.Equal(e.LiquidatePrice, a.LiquidatePrice, "LiquidatePrice")
	r.Equal(e.LiquidateRate, a.LiquidateRate, "LiquidateRate")
	r.Equal(e.MarginLevel, a.MarginLevel, "MarginLevel")
	r.Equal(e.MarginLevelStatus, a.MarginLevelStatus, "MarginLevelStatus")
	r.Equal(e.MarginRatio, a.MarginRatio, "MarginRatio")
	r.Equal(e.TradeEnabled, a.TradeEnabled, "TradeEnabled")
}

func (s *marginTestSuite) assertMarginAccountEqual(e, a *MarginAccount) {
	r := s.r()
	r.Equal(e.Created, a.Created, "Created")
	r.Equal(e.BorrowEnabled, a.BorrowEnabled, "BorrowEnabled")
	r.Equal(e.MarginLevel, a.MarginLevel, "MarginLevel")
	r.Equal(e.CollateralMarginLevel, a.CollateralMarginLevel, "CollateralMarginLevel")
	r.Equal(e.TotalAssetOfBTC, a.TotalAssetOfBTC, "TotalAssetOfBTC")
	r.Equal(e.TotalLiabilityOfBTC, a.TotalLiabilityOfBTC, "TotalLiabilityOfBTC")
	r.Equal(e.TotalNetAssetOfBTC, a.TotalNetAssetOfBTC, "TotalNetAssetOfBTC")
	r.Equal(e.TotalCollateralValueInUSDT, a.TotalCollateralValueInUSDT, "TotalCollateralValueInUSDT")
	r.Equal(e.TradeEnabled, a.TradeEnabled, "TradeEnabled")
	r.Equal(e.TransferInEnabled, a.TransferInEnabled, "TransferInEnabled")
	r.Equal(e.TransferOutEnabled, a.TransferOutEnabled, "TransferOutEnabled")
	r.Equal(e.AccountType, a.AccountType, "AccountType")
	r.Len(a.UserAssets, len(e.UserAssets), "UserAssets")
	for i := 0; i < len(a.UserAssets); i++ {
		s.assertUserAssetEqual(e.UserAssets[i], a.UserAssets[i])
	}
}

func (s *marginTestSuite) assertUserAssetEqual(e, a UserAsset) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Borrowed, a.Borrowed, "Borrowed")
	r.Equal(e.Free, a.Free, "Free")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.Locked, a.Locked, "Locked")
	r.Equal(e.NetAsset, a.NetAsset, "NetAsset")
}

func (s *marginTestSuite) TestGetMarginAsset() {
	data := []byte(`{
		"assetFullName": "Binance Coin",
		"assetName": "BNB",
		"isBorrowable": false,
		"isMortgageable": true,
		"userMinBorrow": "0.00000000",
		"userMinRepay": "0.00000000"
  	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BNB"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.setParam("asset", asset)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetMarginAssetService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	e := &MarginAsset{
		FullName:      "Binance Coin",
		Name:          asset,
		Borrowable:    false,
		Mortgageable:  true,
		UserMinBorrow: "0.00000000",
		UserMinRepay:  "0.00000000",
	}
	s.assertMarginAssetEqual(e, res)
}

func (s *marginTestSuite) assertMarginAssetEqual(e, a *MarginAsset) {
	r := s.r()
	r.Equal(e.FullName, a.FullName, "FullName")
	r.Equal(e.Name, a.Name, "Name")
	r.Equal(e.Borrowable, a.Borrowable, "Borrowable")
	r.Equal(e.Mortgageable, a.Mortgageable, "Mortgageable")
	r.Equal(e.UserMinBorrow, a.UserMinBorrow, "UserMinBorrow")
	r.Equal(e.UserMinRepay, a.UserMinRepay, "UserMinRepay")
}

func (s *marginTestSuite) TestGetMarginPair() {
	data := []byte(`{
		"id":323355778339572400,
		"symbol":"BTCUSDT",
		"base":"BTC",
		"quote":"USDT",
		"isMarginTrade":true,
		"isBuyAllowed":true,
		"isSellAllowed":true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetMarginPairService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	e := &MarginPair{
		ID:            323355778339572400,
		Symbol:        symbol,
		Base:          "BTC",
		Quote:         "USDT",
		IsMarginTrade: true,
		IsBuyAllowed:  true,
		IsSellAllowed: true,
	}
	s.assertMarginPairEqual(e, res)
}

func (s *marginTestSuite) assertMarginPairEqual(e, a *MarginPair) {
	r := s.r()
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Base, a.Base, "Base")
	r.Equal(e.Quote, a.Quote, "Quote")
	r.Equal(e.IsMarginTrade, a.IsMarginTrade, "IsMarginTrade")
	r.Equal(e.IsBuyAllowed, a.IsBuyAllowed, "IsBuyAllowed")
	r.Equal(e.IsSellAllowed, a.IsSellAllowed, "IsSellAllowed")
}

func (s *marginTestSuite) TestGetMarginAllPairs() {
	data := []byte(`[{
		"id":323355778339572400,
		"symbol":"BTCUSDT",
		"base":"BTC",
		"quote":"USDT",
		"isMarginTrade":true,
		"isBuyAllowed":true,
		"isSellAllowed":true
	},{
		"id":351637150141315861,
		"symbol":"BNBBTC",
		"base":"BNB",
		"quote":"BTC",
		"isMarginTrade":true,
		"isBuyAllowed":true,
		"isSellAllowed":true
	}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetMarginAllPairsService().
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(res, 2)
	e := []*MarginAllPair{
		{
			ID:            323355778339572400,
			Symbol:        "BTCUSDT",
			Base:          "BTC",
			Quote:         "USDT",
			IsMarginTrade: true,
			IsBuyAllowed:  true,
			IsSellAllowed: true,
		}, {
			ID:            351637150141315861,
			Symbol:        "BNBBTC",
			Base:          "BNB",
			Quote:         "BTC",
			IsMarginTrade: true,
			IsBuyAllowed:  true,
			IsSellAllowed: true,
		},
	}
	for i := 0; i < len(res); i++ {
		s.assertMarginAllPairsEqual(e[i], res[i])
	}
}

func (s *marginTestSuite) assertMarginAllPairsEqual(e, a *MarginAllPair) {
	r := s.r()
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Base, a.Base, "Base")
	r.Equal(e.Quote, a.Quote, "Quote")
	r.Equal(e.IsMarginTrade, a.IsMarginTrade, "IsMarginTrade")
	r.Equal(e.IsBuyAllowed, a.IsBuyAllowed, "IsBuyAllowed")
	r.Equal(e.IsSellAllowed, a.IsSellAllowed, "IsSellAllowed")
}

func (s *marginTestSuite) TestGetMarginPriceIndex() {
	data := []byte(`{
		"calcTime": 1562046418000,
		"price": "0.00333930",
		"symbol": "BNBBTC"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BNBBTC"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetMarginPriceIndexService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	e := &MarginPriceIndex{
		CalcTime: 1562046418000,
		Symbol:   symbol,
		Price:    "0.00333930",
	}
	s.assertMarginPriceIndexEqual(e, res)
}

func (s *marginTestSuite) assertMarginPriceIndexEqual(e, a *MarginPriceIndex) {
	r := s.r()
	r.Equal(e.CalcTime, a.CalcTime, "CalcTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
}

func (s *marginTestSuite) TestListMarginTrades() {
	data := []byte(`[{
		"commission": "0.00006000",
		"commissionAsset": "BTC",
		"id": 34,
		"isBestMatch": true,
		"isBuyer": false,
		"isMaker": false,
		"orderId": 39324,
		"price": "0.02000000",
		"qty": "3.00000000",
		"symbol": "BNBBTC",
		"time": 1561973357171
	},{
		"commission": "0.00002950",
		"commissionAsset": "BTC",
		"id": 32,
		"isBestMatch": true,
		"isBuyer": false,
		"isMaker": true,
		"orderId": 39319,
		"price": "0.00590000",
		"qty": "5.00000000",
		"symbol": "BNBBTC",
		"time": 1561964645345
	}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	limit := 3
	fromID := int64(1)
	startTime := int64(1499865549590)
	endTime := int64(1499865549590)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
			"fromId":    fromID,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewListMarginTradesService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).
		Limit(limit).FromID(fromID).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 2)
	e := []*TradeV3{
		{
			ID:              34,
			Symbol:          "BNBBTC",
			OrderID:         39324,
			Price:           "0.02000000",
			Quantity:        "3.00000000",
			Commission:      "0.00006000",
			CommissionAsset: "BTC",
			Time:            1561973357171,
			IsBuyer:         false,
			IsMaker:         false,
			IsBestMatch:     true,
		},
		{
			ID:              32,
			Symbol:          "BNBBTC",
			OrderID:         39319,
			Price:           "0.00590000",
			Quantity:        "5.00000000",
			Commission:      "0.00002950",
			CommissionAsset: "BTC",
			Time:            1561964645345,
			IsBuyer:         false,
			IsMaker:         true,
			IsBestMatch:     true,
		},
	}
	for i := 0; i < len(trades); i++ {
		s.assertTradeV3Equal(e[i], trades[i])
	}
}

func (s *marginTestSuite) TestGetMaxBorrowable() {
	data := []byte(`{
		"amount": "1.69248805",
		"borrowLimit": "1000"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "BNBBTC",
		})
		s.assertRequestEqual(e, r)
	})

	borrowable, err := s.client.NewGetMaxBorrowableService().
		Asset("BNBBTC").Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &MaxBorrowable{
		Amount:      "1.69248805",
		BorrowLimit: "1000",
	}
	s.assertMaxBorrowableEqual(e, borrowable)
}

func (s *marginTestSuite) assertMaxBorrowableEqual(e, a *MaxBorrowable) {
	s.r().Equal(e.Amount, a.Amount, "Amount")
	s.r().Equal(e.BorrowLimit, a.BorrowLimit, "BorrowLimit")
}

func (s *marginTestSuite) TestGetMaxTransferable() {
	data := []byte(`{
		"amount": "3.59498107"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "BNBBTC",
		})
		s.assertRequestEqual(e, r)
	})

	transferable, err := s.client.NewGetMaxTransferableService().
		Asset("BNBBTC").Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &MaxTransferable{
		Amount: "3.59498107",
	}
	s.assertMaxTransferableEqual(e, transferable)
}

func (s *marginTestSuite) assertMaxTransferableEqual(e, a *MaxTransferable) {
	s.r().Equal(e.Amount, a.Amount, "Amount")
}

func (s *marginTestSuite) TestStartMarginUserStream() {
	data := []byte(`{
        "listenKey": "T3ee22BIYuWqmvne0HNq2A2WsFlEtLhvWCtItw6ffhhdmjifQ2tRbuKkTHhr"
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		s.assertRequestEqual(newRequest(), r)
	})

	listenKey, err := s.client.NewStartMarginUserStreamService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("T3ee22BIYuWqmvne0HNq2A2WsFlEtLhvWCtItw6ffhhdmjifQ2tRbuKkTHhr", listenKey)
}

func (s *marginTestSuite) TestKeepaliveMarginUserStream() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey := "dummykey"
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newRequest().setFormParam("listenKey", listenKey), r)
	})

	err := s.client.NewKeepaliveMarginUserStreamService().ListenKey(listenKey).Do(newContext())
	s.r().NoError(err)
}

func (s *marginTestSuite) TestCloseMarginUserStream() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey := "dummykey"
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newRequest().setParam("listenKey", listenKey), r)
	})

	err := s.client.NewCloseMarginUserStreamService().ListenKey(listenKey).Do(newContext())
	s.r().NoError(err)
}

func (s *marginTestSuite) TestGetIsolatedMarginAllPairs() {
	data := []byte(`[{
        "base": "BNB",
        "isBuyAllowed": true,
        "isMarginTrade": true,
        "isSellAllowed": true,
        "quote": "BTC",
        "symbol": "BNBBTC"     
    },
    {
        "base": "TRX",
        "isBuyAllowed": true,
        "isMarginTrade": true,
        "isSellAllowed": true,
        "quote": "BTC",
        "symbol": "TRXBTC"    
    }]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetIsolatedMarginAllPairsService().
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(res, 2)
	e := []*IsolatedMarginAllPair{
		{
			Symbol:        "BNBBTC",
			Base:          "BNB",
			Quote:         "BTC",
			IsMarginTrade: true,
			IsBuyAllowed:  true,
			IsSellAllowed: true,
		}, {
			Symbol:        "TRXBTC",
			Base:          "TRX",
			Quote:         "BTC",
			IsMarginTrade: true,
			IsBuyAllowed:  true,
			IsSellAllowed: true,
		},
	}

	for i := 0; i < len(res); i++ {
		s.assertIsolatedMarginAllPairsEqual(e[i], res[i])
	}
}

func (s *marginTestSuite) assertIsolatedMarginAllPairsEqual(e, a *IsolatedMarginAllPair) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Base, a.Base, "Base")
	r.Equal(e.Quote, a.Quote, "Quote")
	r.Equal(e.IsMarginTrade, a.IsMarginTrade, "IsMarginTrade")
	r.Equal(e.IsBuyAllowed, a.IsBuyAllowed, "IsBuyAllowed")
	r.Equal(e.IsSellAllowed, a.IsSellAllowed, "IsSellAllowed")
}

func (s *marginTestSuite) TestIsolatedMarginTransferService() {
	data := []byte(`{"tranId": 100000001}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	var (
		asset     = "BTC"
		symbol    = "BTCBUSD"
		transFrom = AccountTypeIsolatedMargin
		transTo   = AccountTypeSpot
		amount    = "1"
	)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":     asset,
			"symbol":    symbol,
			"transFrom": transFrom,
			"transTo":   transTo,
			"amount":    amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewIsolatedMarginTransferService().
		Asset(asset).Symbol(symbol).TransFrom(transFrom).TransTo(transTo).Amount(amount).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &TransactionResponse{TranID: 100000001}
	s.r().Equal(res, e)
}
