package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type getMarginForceOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestGetMarginForceOrdersService(t *testing.T) {
	suite.Run(t, new(getMarginForceOrdersServiceTestSuite))
}

func (s *getMarginForceOrdersServiceTestSuite) TestGetForceOrders() {
	data := []byte(`{
		"rows": [
			{
				"avgPrice": "0.00388359",
				"executedQty": "31.39000000",
				"orderId": 180015097,
				"price": "0.00388110",
				"qty": "31.39000000",
				"side": "SELL",
				"symbol": "BNBBTC",
				"timeInForce": "GTC",
				"updatedTime": 1558941374745
			}
		],
		"total": 1
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	startTime := int64(1558941374745)
	endTime := int64(1558941374746)
	current := int64(1)
	size := int64(10)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetMarginForceOrdersService().
		StartTime(startTime).
		EndTime(endTime).
		Current(current).
		Size(size).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1), res.Total)
	s.r().Len(res.Rows, 1)
	order := res.Rows[0]
	s.r().Equal("BNBBTC", order.Symbol)
	s.r().Equal(int64(180015097), order.OrderID)
	s.r().Equal("SELL", order.Side)
	s.r().Equal("GTC", order.TimeInForce)
	s.r().Equal("0.00388359", order.AvgPrice)
	s.r().Equal("31.39000000", order.ExecutedQty)
}
