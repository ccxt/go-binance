package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginForceOrdersServiceTestSuite struct {
	baseTestSuite
}

func TestMarginForceOrdersService(t *testing.T) {
	suite.Run(t, new(marginForceOrdersServiceTestSuite))
}

func (s *marginForceOrdersServiceTestSuite) TestForceOrders() {
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

	size := int64(10)
	current := int64(1)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"size":    size,
			"current": current,
		})
		s.assertRequestEqual(e, r)
	})

	response, err := s.client.NewMarginForceOrdersService().
		Size(size).
		Current(current).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1), response.Total)
	s.r().Len(response.Rows, 1)
	order := response.Rows[0]
	s.r().Equal(int64(180015097), order.OrderID)
	s.r().Equal("BNBBTC", order.Symbol)
	s.r().Equal("SELL", order.Side)
	s.r().Equal("0.00388359", order.AvgPrice)
}

func (s *marginForceOrdersServiceTestSuite) TestForceOrdersWithAllParams() {
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

	startTime := int64(1558941374000)
	endTime := int64(1558941374999)
	size := int64(10)
	current := int64(1)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"startTime": startTime,
			"endTime":   endTime,
			"size":      size,
			"current":   current,
		})
		s.assertRequestEqual(e, r)
	})

	response, err := s.client.NewMarginForceOrdersService().
		StartTime(startTime).
		EndTime(endTime).
		Size(size).
		Current(current).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1), response.Total)
}
