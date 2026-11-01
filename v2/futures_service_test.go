package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type futuresTransferTestSuite struct {
	baseTestSuite
}

func TestFuturesTransferService(t *testing.T) {
	suite.Run(t, new(futuresTransferTestSuite))
}

func (s *futuresTransferTestSuite) TestTransfer() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	transferType := FuturesTransferTypeToFutures
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
			"type":   transferType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewFuturesTransferService().Asset(asset).
		Amount(amount).Type(transferType).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *futuresTransferTestSuite) assertTransactionResponseEqual(a, e *TransactionResponse) {
	s.r().Equal(a.TranID, e.TranID, "TranID")
}

func (s *futuresTransferTestSuite) TestListFuturesTransfer() {
	data := []byte(`{
		"rows": [
		  {
			"asset": "USDT",
			"tranId": 100000001,
			"amount": "40.84624400",
			"type": 1,
			"timestamp": 1555056425000,
			"status": "CONFIRMED"
		  }
		],
		"total": 1
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "USDT"
	startTime := int64(1555056425000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":     asset,
			"startTime": startTime,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListFuturesTransferService().Asset(asset).
		StartTime(startTime).Do(newContext())
	s.r().NoError(err)
	e := &FuturesTransferHistory{
		Rows: []FuturesTransfer{
			{
				Asset:     asset,
				TranID:    int64(100000001),
				Amount:    "40.84624400",
				Type:      1,
				Timestamp: int64(1555056425000),
				Status:    FuturesTransferStatusTypeConfirmed,
			},
		},
		Total: 1,
	}
	s.assertFuturesTransferHistoryEqual(e, res)
}

func (s *futuresTransferTestSuite) assertFuturesTransferHistoryEqual(e, a *FuturesTransferHistory) {
	s.r().Equal(e.Total, a.Total, "Total")
	s.r().Len(a.Rows, len(e.Rows))
	for i := range a.Rows {
		s.assertFuturesTransferEqual(e.Rows[i], a.Rows[i])
	}
}

func (s *futuresTransferTestSuite) assertFuturesTransferEqual(e, a FuturesTransfer) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.TranID, a.TranID, "TranID")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Status, a.Status, "Status")
}

type futuresOrderBookHistoryTestSuite struct {
	baseTestSuite
}

func TestFuturesOrderBookHistoryService(t *testing.T) {
	suite.Run(t, new(futuresOrderBookHistoryTestSuite))
}

func (s *futuresOrderBookHistoryTestSuite) TestFuturesOrderBookHistory() {
	data := []byte(`{
		"data": [
        {
            "day": "2023-06-30",
            "url": "https://bin-prod-user-rebate-bucket.s3.ap-northeast-1.amazonaws.com/future-data-symbol-update/2023-06-30/BTCUSDT_T_DEPTH_2023-06-30.tar.gz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20230925T025710Z&X-Amz-SignedHeaders=host&X-Amz-Expires=86399&X-Amz-Credential=AKIAVL364M5ZNFZ74IPP%2F20230925%2Fap-northeast-1%2Fs3%2Faws4_request&X-Amz-Signature=5fffcb390d10f34d71615726f81f99e42d80a11532edeac77b858c51a88cbf59"
        }
    	]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSDT"
	dataType := FuturesOrderBookHistoryDataTypeTDepth
	startTime := int64(1625040000000)
	endTime := int64(1625126399999)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"dataType":  "T_DEPTH",
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewFuturesOrderBookHistoryService().Symbol(symbol).
		DataType(dataType).StartTime(startTime).EndTime(endTime).Do(newContext())
	s.r().NoError(err)
	e := &FuturesOrderBookHistory{
		Data: []*FuturesOrderBookHistoryItem{
			{
				Day: "2023-06-30",
				Url: "https://bin-prod-user-rebate-bucket.s3.ap-northeast-1.amazonaws.com/future-data-symbol-update/2023-06-30/BTCUSDT_T_DEPTH_2023-06-30.tar.gz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20230925T025710Z&X-Amz-SignedHeaders=host&X-Amz-Expires=86399&X-Amz-Credential=AKIAVL364M5ZNFZ74IPP%2F20230925%2Fap-northeast-1%2Fs3%2Faws4_request&X-Amz-Signature=5fffcb390d10f34d71615726f81f99e42d80a11532edeac77b858c51a88cbf59",
			},
		},
	}
	s.assertFuturesOrderBookHistoryEqual(e, res)
}

func (s *futuresOrderBookHistoryTestSuite) assertFuturesOrderBookHistoryEqual(a, e *FuturesOrderBookHistory) {
	for index, v := range a.Data {
		v.Day = e.Data[index].Day
		v.Url = e.Data[index].Url
	}
}
