package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginAllOCOServiceTestSuite struct {
	baseTestSuite
}

func TestMarginAllOCOService(t *testing.T) {
	suite.Run(t, new(marginAllOCOServiceTestSuite))
}

func (s *marginAllOCOServiceTestSuite) TestMarginAllOCO() {
	data := []byte(`[
		{
			"orderListId": 29,
			"contingencyType": "OCO",
			"listStatusType": "EXEC_STARTED",
			"listOrderStatus": "EXECUTING",
			"listClientOrderId": "amEEAXryFzFwYF1FeRpUoZ",
			"transactionTime": 1565245913483,
			"symbol": "LTCBTC",
			"orders": [
				{
					"symbol": "LTCBTC",
					"orderId": 4,
					"clientOrderId": "oD7aesZqjEGlZrbtRpy5zB"
				},
				{
					"symbol": "LTCBTC",
					"orderId": 5,
					"clientOrderId": "Jr1h6xirOxgeJOUuYQS7V3"
				}
			]
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromID := int64(1)
	startTime := int64(1565245913483)
	endTime := int64(1565245913484)
	limit := 500
	recvWindow := int64(1000)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromId":     fromID,
			"startTime":  startTime,
			"endTime":    endTime,
			"limit":      limit,
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewMarginAllOCOService().
		FromID(fromID).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		RecvWindow(recvWindow).
		Do(newContext())

	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(29), orders[0].OrderListID)
	s.r().Equal("OCO", orders[0].ContingencyType)
	s.r().Equal("EXEC_STARTED", orders[0].ListStatusType)
	s.r().Equal("EXECUTING", orders[0].ListOrderStatus)
	s.r().Equal("amEEAXryFzFwYF1FeRpUoZ", orders[0].ListClientOrderID)
	s.r().Equal(int64(1565245913483), orders[0].TransactionTime)
	s.r().Equal("LTCBTC", orders[0].Symbol)
	s.r().Len(orders[0].Orders, 2)
}
