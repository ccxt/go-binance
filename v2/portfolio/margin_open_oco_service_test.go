package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginOpenOCOServiceTestSuite struct {
	baseTestSuite
}

func TestMarginOpenOCOService(t *testing.T) {
	suite.Run(t, new(marginOpenOCOServiceTestSuite))
}

func (s *marginOpenOCOServiceTestSuite) TestGetOpenOCO() {
	data := []byte(`[
		{
			"orderListId": 31,
			"contingencyType": "OCO",
			"listStatusType": "EXEC_STARTED",
			"listOrderStatus": "EXECUTING",
			"listClientOrderId": "wuB13fmulKj3YjdqWEcsnp",
			"transactionTime": 1565246080644,
			"symbol": "LTCBTC",
			"orders": [
				{
					"symbol": "LTCBTC",
					"orderId": 4,
					"clientOrderId": "r3EH2N76dHfLoSZWIUw1bT"
				},
				{
					"symbol": "LTCBTC",
					"orderId": 5,
					"clientOrderId": "Cv1SnyPD3qhqpbjpYEHbd2"
				}
			]
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewMarginOpenOCOService().Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 1)
	s.r().Equal(int64(31), orders[0].OrderListID)
	s.r().Equal("OCO", orders[0].ContingencyType)
	s.r().Equal("EXEC_STARTED", orders[0].ListStatusType)
	s.r().Equal("EXECUTING", orders[0].ListOrderStatus)
	s.r().Equal("wuB13fmulKj3YjdqWEcsnp", orders[0].ListClientOrderID)
	s.r().Equal(int64(1565246080644), orders[0].TransactionTime)
	s.r().Equal("LTCBTC", orders[0].Symbol)
	s.r().Len(orders[0].Orders, 2)
	s.r().Equal("LTCBTC", orders[0].Orders[0].Symbol)
	s.r().Equal(int64(4), orders[0].Orders[0].OrderID)
	s.r().Equal("r3EH2N76dHfLoSZWIUw1bT", orders[0].Orders[0].ClientOrderID)
}

func (s *marginOpenOCOServiceTestSuite) TestGetOpenOCOWithRecvWindow() {
	data := []byte(`[]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	recvWindow := int64(1000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"recvWindow": recvWindow,
		})
		s.assertRequestEqual(e, r)
	})

	orders, err := s.client.NewMarginOpenOCOService().
		RecvWindow(recvWindow).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(orders, 0)
}
