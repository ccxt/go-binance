package binance

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type userUniversalTransferTestSuite struct {
	baseTestSuite
}

func TestUserUniversalTransferService(t *testing.T) {
	suite.Run(t, new(userUniversalTransferTestSuite))
}

func (s *userUniversalTransferTestSuite) TestUserUniversalTransfer() {
	data := []byte(`
	{
		"tranId":13526853623
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	types := UserUniversalTransferTypeMainToUmFutures
	asset := "USDT"
	amount := "0.1"
	fromSymbol := "USDT"
	toSymbol := "USDT"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"type":       types,
			"asset":      asset,
			"amount":     amount,
			"fromSymbol": fromSymbol,
			"toSymbol":   toSymbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUserUniversalTransferService().
		Type(types).
		Asset(asset).
		Amount(amount).
		FromSymbol(fromSymbol).
		ToSymbol(toSymbol).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(int64(13526853623), res.ID)
}

func (s *userUniversalTransferTestSuite) TestListUserUniversalTransfer() {
	data := []byte(`{
		"total":2,
		"rows":[
			{
				"asset":"USDT",
				"amount":"1",
				"type":"MAIN_UMFUTURE",
				"status": "CONFIRMED",
				"tranId": 11415955596,
				"timestamp":1544433328000
			},
			{
				"asset":"USDT",
				"amount":"2",
				"type":"MAIN_UMFUTURE",
				"status": "CONFIRMED",
				"tranId": 11366865406,
				"timestamp":1544433328000
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	types := UserUniversalTransferTypeMainToUmFutures
	startTime := time.Date(2018, 12, 10, 0, 0, 0, 0, time.UTC).UnixMilli()
	endTime := time.Date(2018, 12, 11, 0, 0, 0, 0, time.UTC).UnixMilli()
	current := 1
	size := 10

	s.assertReq(func(r *request) {
		fmt.Println(r.query)
		e := newSignedRequest().setParams(params{
			"type":      string(types),
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListUserUniversalTransferService().
		Type(types).
		StartTime(startTime).
		EndTime(endTime).
		Current(current).
		Size(size).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(res.Total, int64(2))
	results := res.Results
	r.Equal(int64(11415955596), results[0].TranId)
	r.Equal("USDT", results[0].Asset)
	r.Equal("1", results[0].Amount)
	r.Equal(UserUniversalTransferTypeMainToUmFutures, results[0].Type)
	r.Equal(UserUniversalTransferStatusTypeConfirmed, results[0].Status)
	r.Equal(int64(1544433328000), results[0].Timestamp)
}
