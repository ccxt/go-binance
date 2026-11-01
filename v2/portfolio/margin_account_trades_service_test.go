package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginAccountTradesServiceTestSuite struct {
	baseTestSuite
}

func TestMarginAccountTradesService(t *testing.T) {
	suite.Run(t, new(marginAccountTradesServiceTestSuite))
}

func (s *marginAccountTradesServiceTestSuite) TestGetTrades() {
	data := []byte(`[
		{
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
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	orderID := int64(39324)
	limit := 500
	fromID := int64(34)
	startTime := int64(1561973357171)
	endTime := int64(1561973357172)

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"orderId":   orderID,
			"limit":     limit,
			"fromId":    fromID,
			"startTime": startTime,
			"endTime":   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewMarginAccountTradesService().
		Symbol(symbol).
		OrderID(orderID).
		Limit(limit).
		FromID(fromID).
		StartTime(startTime).
		EndTime(endTime).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(trades, 1)
	s.r().Equal("BNBBTC", trades[0].Symbol)
	s.r().Equal(int64(34), trades[0].ID)
	s.r().Equal(int64(39324), trades[0].OrderID)
	s.r().Equal("0.00006000", trades[0].Commission)
	s.r().Equal("BTC", trades[0].CommissionAsset)
	s.r().Equal("0.02000000", trades[0].Price)
	s.r().Equal("3.00000000", trades[0].Qty)
	s.r().False(trades[0].IsBuyer)
	s.r().False(trades[0].IsMaker)
	s.r().True(trades[0].IsBestMatch)
}
