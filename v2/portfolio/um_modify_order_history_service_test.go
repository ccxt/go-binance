package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umModifyOrderHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestUMModifyOrderHistoryService(t *testing.T) {
	suite.Run(t, new(umModifyOrderHistoryServiceTestSuite))
}

func (s *umModifyOrderHistoryServiceTestSuite) TestModifyOrderHistory() {
	data := []byte(`[
		{
			"amendmentId": 5363,
			"symbol": "BTCUSDT",
			"pair": "BTCUSDT",
			"orderId": 20072994037,
			"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
			"time": 1629184560899,
			"amendment": {
				"price": {
					"before": "30004",
					"after": "30003.2"
				},
				"origQty": {
					"before": "1",
					"after": "1"
				},
				"count": 3
			},
			"priceMatch": "NONE"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	orderID := int64(20072994037)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":  symbol,
			"orderId": orderID,
		})
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewUMModifyOrderHistoryService().
		Symbol(symbol).
		OrderID(orderID).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(history, 1)
	s.r().Equal(int64(5363), history[0].AmendmentID)
	s.r().Equal("BTCUSDT", history[0].Symbol)
	s.r().Equal(int64(20072994037), history[0].OrderID)
	s.r().Equal("30004", history[0].Amendment.Price.Before)
	s.r().Equal("30003.2", history[0].Amendment.Price.After)
	s.r().Equal(3, history[0].Amendment.Count)
}

func (s *umModifyOrderHistoryServiceTestSuite) TestModifyOrderHistoryWithAllParams() {
	data := []byte(`[
		{
			"amendmentId": 5363,
			"symbol": "BTCUSDT",
			"pair": "BTCUSDT",
			"orderId": 20072994037,
			"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
			"time": 1629184560899,
			"amendment": {
				"price": {
					"before": "30004",
					"after": "30003.2"
				},
				"origQty": {
					"before": "1",
					"after": "1"
				},
				"count": 3
			},
			"priceMatch": "NONE"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	origClientOrderID := "LJ9R4QZDihCaS8UAOOLpgW"
	startTime := int64(1629184560000)
	endTime := int64(1629184560999)
	limit := 500
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":            symbol,
			"origClientOrderId": origClientOrderID,
			"startTime":         startTime,
			"endTime":           endTime,
			"limit":             limit,
		})
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewUMModifyOrderHistoryService().
		Symbol(symbol).
		OrigClientOrderID(origClientOrderID).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(history, 1)
}
