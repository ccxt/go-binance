package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginOCOQueryServiceTestSuite struct {
	baseTestSuite
}

func TestMarginOCOQueryService(t *testing.T) {
	suite.Run(t, new(marginOCOQueryServiceTestSuite))
}

func (s *marginOCOQueryServiceTestSuite) TestMarginOCOQuery() {
	data := []byte(`{
        "orderListId": 27,
        "contingencyType": "OCO",
        "listStatusType": "EXEC_STARTED",
        "listOrderStatus": "EXECUTING",
        "listClientOrderId": "h2USkA5YQpaXHPIrkd96xE",
        "transactionTime": 1565245656253,
        "symbol": "LTCBTC",
        "orders": [
            {
                "symbol": "LTCBTC",
                "orderId": 4,
                "clientOrderId": "qD1gy3kc3Gx0rihm9Y3xwS"
            },
            {
                "symbol": "LTCBTC",
                "orderId": 5,
                "clientOrderId": "ARzZ9I00CPM8i3NhmU9Ega"
            }
        ]
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orderListID := int64(27)

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("orderListId", orderListID)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewMarginOCOQueryService().
		OrderListID(orderListID).
		Do(newContext())

	s.r().NoError(err)
	e := &MarginOCOResponse{
		OrderListID:       27,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "EXECUTING",
		ListClientOrderID: "h2USkA5YQpaXHPIrkd96xE",
		TransactionTime:   1565245656253,
		Symbol:            "LTCBTC",
		Orders: []MarginOCOOrder{
			{
				Symbol:        "LTCBTC",
				OrderID:       4,
				ClientOrderID: "qD1gy3kc3Gx0rihm9Y3xwS",
			},
			{
				Symbol:        "LTCBTC",
				OrderID:       5,
				ClientOrderID: "ARzZ9I00CPM8i3NhmU9Ega",
			},
		},
	}
	s.assertOCOResponseEqual(e, res)
}

func (s *marginOCOQueryServiceTestSuite) assertOCOResponseEqual(e, a *MarginOCOResponse) {
	r := s.r()
	r.Equal(e.OrderListID, a.OrderListID, "OrderListID")
	r.Equal(e.ContingencyType, a.ContingencyType, "ContingencyType")
	r.Equal(e.ListStatusType, a.ListStatusType, "ListStatusType")
	r.Equal(e.ListOrderStatus, a.ListOrderStatus, "ListOrderStatus")
	r.Equal(e.ListClientOrderID, a.ListClientOrderID, "ListClientOrderID")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Len(a.Orders, len(e.Orders))
	for idx := range e.Orders {
		r.Equal(e.Orders[idx].Symbol, a.Orders[idx].Symbol, "Order.Symbol")
		r.Equal(e.Orders[idx].OrderID, a.Orders[idx].OrderID, "Order.OrderID")
		r.Equal(e.Orders[idx].ClientOrderID, a.Orders[idx].ClientOrderID, "Order.ClientOrderID")
	}
}
