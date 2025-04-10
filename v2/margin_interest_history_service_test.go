package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginInterestHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestMarginInterestHistoryService(t *testing.T) {
	suite.Run(t, new(marginInterestHistoryServiceTestSuite))
}

func (s *marginInterestHistoryServiceTestSuite) TestMarginInterestHistory() {
	data := []byte(`
	{
		"rows": [
			{            
				"txId": 1352286576452864727,           
				"interestAccuredTime": 1672160400000,            
				"asset": "USDT",
				"rawAsset": "USDT",           
				"principal": "45.3313",            
				"interest": "0.00024995",            
				"interestRate": "0.00013233",            
				"type": "ON_BORROW",           
				"isolatedSymbol": "BNBUSDT"     
			}
		],
		"total": 1
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "USDT"
	symbol := "BNBUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset":          asset,
			"isolatedSymbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewMarginInterestHistoryService().
		Asset(asset).
		IsolatedSymbol(symbol).
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	s.Len(history.Rows, 1)
	s.Equal(int64(1), history.Total)

	item := history.Rows[0]
	s.Equal(int64(1352286576452864727), item.TxId)
	s.Equal(int64(1672160400000), item.InterestAccuredTime)
	s.Equal("USDT", item.Asset)
	s.Equal("USDT", item.RawAsset)
	s.Equal("45.3313", item.Principal)
	s.Equal("0.00024995", item.Interest)
	s.Equal("0.00013233", item.InterestRate)
	s.Equal("ON_BORROW", item.Type)
	s.Equal("BNBUSDT", item.IsolatedSymbol)
}
