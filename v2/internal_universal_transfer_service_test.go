package binance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type internalUniversalTransferServiceTestSuite struct {
	baseTestSuite
}

func TestInternalUniversalTransferService(t *testing.T) {
	suite.Run(t, new(internalUniversalTransferServiceTestSuite))
}

func (s *internalUniversalTransferServiceTestSuite) TestInternalUniversalTransfer() {
	data := []byte(`
	{
		"tranId":11945860693,
		"clientTranId":"testID"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromEmail := "sub1@gmail.com"
	toEmail := "sub2@gmail.com"
	fromAccountType := "USDT_FUTURE"
	toAccountType := "SPOT"
	asset := "USDT"
	amount := "100.0"
	clientTranId := "testID"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":           asset,
			"amount":          amount,
			"fromEmail":       fromEmail,
			"toEmail":         toEmail,
			"fromAccountType": fromAccountType,
			"toAccountType":   toAccountType,
			"clientTranId":    clientTranId,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewInternalUniversalTransferService().
		FromEmail(fromEmail).
		ToEmail(toEmail).
		FromAccountType(fromAccountType).
		ToAccountType(toAccountType).
		Asset(asset).
		Amount(amount).
		ClientTranId(clientTranId).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(int64(11945860693), res.ID)
	r.Equal(clientTranId, res.ClientTranID)
}

func (s *internalUniversalTransferServiceTestSuite) TestInternalUniversalTransferHistory() {
	data := []byte(`
	{
		"result": [
			{
				"tranId": 92275823339,
				"fromEmail": "sub1@gmail.com",
				"toEmail": "sub2@gmail.com",
				"asset": "USDT",
				"amount": "100.0",
				"createTimeStamp": 1640317374000,
				"fromAccountType": "USDT_FUTURE",
				"toAccountType": "SPOT",
				"status": "SUCCESS",
				"clientTranId": "testID"
			}
		],
		"totalCount": 1
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromEmail := "sub1@gmail.com"
	toEmail := "sub2@gmail.com"
	clientTranId := "testID"
	endTime := time.Now().UnixNano() / 1000 / 1000
	startTime := endTime - 3600*1000
	page := 1
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromEmail":    fromEmail,
			"toEmail":      toEmail,
			"startTime":    startTime,
			"endTime":      endTime,
			"clientTranId": clientTranId,
			"page":         1,
			"limit":        10,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewInternalUniversalTransferHistoryService().
		FromEmail(fromEmail).
		ToEmail(toEmail).
		StartTime(startTime).
		EndTime(endTime).
		Page(page).
		Limit(limit).
		ClientTranId(clientTranId).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	s.assertInternalUniversalTransferEqual(&InternalUniversalTransfer{
		FromEmail:       fromEmail,
		ToEmail:         toEmail,
		FromAccountType: "USDT_FUTURE",
		ToAccountType:   "SPOT",
		TranId:          92275823339,
		ClientTranId:    clientTranId,
		Asset:           "USDT",
		Amount:          "100.0",
		Status:          "SUCCESS",
		CreateTimeStamp: 1640317374000,
	}, res.Result[0])
	r.Equal(1, res.TotalCount, "TotalCount")
}

func (s *internalUniversalTransferServiceTestSuite) assertInternalUniversalTransferEqual(e, a *InternalUniversalTransfer) {
	r := s.r()
	r.Equal(e.FromEmail, a.FromEmail, "FromEmail")
	r.Equal(e.ToEmail, a.ToEmail, "ToEmail")
	r.Equal(e.FromAccountType, a.FromAccountType, "FromAccountType")
	r.Equal(e.ToAccountType, a.ToAccountType, "ToAccountType")
	r.Equal(e.TranId, a.TranId, "TranId")
	r.Equal(e.ClientTranId, a.ClientTranId, "ClientTranId")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.CreateTimeStamp, a.CreateTimeStamp, "CreateTimeStamp")
}
