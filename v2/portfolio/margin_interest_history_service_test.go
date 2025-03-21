package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginInterestHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestMarginInterestHistoryService(t *testing.T) {
	suite.Run(t, new(marginInterestHistoryServiceTestSuite))
}

func (s *marginInterestHistoryServiceTestSuite) TestGetMarginInterestHistory() {
	data := []byte(`{
		"rows": [
			{
				"txId": 1352286576452864727,
				"interestAccuredTime": 1672160400000,
				"asset": "USDT",
				"rawAsset": "USDT",
				"principal": "45.3313",
				"interest": "0.00024995",
				"interestRate": "0.00013233",
				"type": "ON_BORROW"
			}
		],
		"total": 1
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "USDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		s.assertRequestEqual(e, r)
	})

	history, err := s.client.NewGetMarginInterestHistoryService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1), history.Total)
	s.r().Len(history.Rows, 1)

	interest := history.Rows[0]
	s.r().Equal(int64(1352286576452864727), interest.TxID)
	s.r().Equal(int64(1672160400000), interest.InterestAccuredTime)
	s.r().Equal("USDT", interest.Asset)
	s.r().Equal("USDT", interest.RawAsset)
	s.r().Equal("45.3313", interest.Principal)
	s.r().Equal("0.00024995", interest.Interest)
	s.r().Equal("0.00013233", interest.InterestRate)
	s.r().Equal("ON_BORROW", interest.Type)
}
