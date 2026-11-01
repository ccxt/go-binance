package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type getMarginRepayServiceTestSuite struct {
	baseTestSuite
}

func TestGetMarginRepayService(t *testing.T) {
	suite.Run(t, new(getMarginRepayServiceTestSuite))
}

func (s *getMarginRepayServiceTestSuite) TestGetMarginRepay() {
	data := []byte(`{
		"rows": [
			{
				"amount": "14.00000000",
				"asset": "BNB",
				"interest": "0.01866667",
				"principal": "13.98133333",
				"status": "CONFIRMED",
				"timestamp": 1563438204000,
				"txId": 2970933056
			}
		],
		"total": 1
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := "BNB"
	txID := int64(2970933056)
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("asset", asset)
		e.setParam("txId", txID)
		s.assertRequestEqual(e, r)
	})

	repays, err := s.client.NewGetMarginRepayService().Asset(asset).TxID(txID).Do(newContext())
	s.r().NoError(err)
	s.r().Equal(int64(1), repays.Total)
	s.r().Len(repays.Rows, 1)

	repay := repays.Rows[0]
	s.r().Equal("14.00000000", repay.Amount)
	s.r().Equal("BNB", repay.Asset)
	s.r().Equal("0.01866667", repay.Interest)
	s.r().Equal("13.98133333", repay.Principal)
	s.r().Equal("CONFIRMED", repay.Status)
	s.r().Equal(int64(1563438204000), repay.Timestamp)
	s.r().Equal(int64(2970933056), repay.TxID)
}
