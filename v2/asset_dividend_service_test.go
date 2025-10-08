package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type assetDividendServiceTestSuite struct {
	baseTestSuite
}

func TestAssetDividendService(t *testing.T) {
	suite.Run(t, new(assetDividendServiceTestSuite))
}

func (s *assetDividendServiceTestSuite) TestListAssetDividend() {
	data := []byte(`
	{
		"rows":[
			{
	            "id":1637366104,
				"amount":"10.00000000",
				"asset":"BHFT",
				"divTime":1563189166000,
				"enInfo":"BHFT distribution",
				"tranId":2968885920
			},
			{
	            "id":1631750237,
				"amount":"10.00000000",
				"asset":"BHFT",
				"divTime":1563189165000,
				"enInfo":"BHFT distribution",
				"tranId":2968885920
			}
		],
		"total":2
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	asset := `BHFT`
	startTime := int64(1508198532000)
	endTime := int64(1508198532001)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			`asset`:     asset,
			`limit`:     2,
			`startTime`: startTime,
			`endTime`:   endTime,
		})
		s.assertRequestEqual(e, r)
	})

	dividend, err := s.client.NewAssetDividendService().
		Asset(asset).
		StartTime(startTime).
		EndTime(endTime).
		Limit(2).
		Do(context.Background())
	r := s.r()
	r.NoError(err)
	rows := dividend.Rows

	s.Len(rows, 2)
	s.assertDividendEqual(&DividendResponse{
		ID:     1637366104,
		Amount: `10.00000000`,
		Asset:  `BHFT`,
		Time:   1563189166000,
		Info:   `BHFT distribution`,
		TranID: 2968885920,
	}, rows[0])
	s.assertDividendEqual(&DividendResponse{
		ID:     1631750237,
		Amount: `10.00000000`,
		Asset:  `BHFT`,
		Time:   1563189165000,
		Info:   `BHFT distribution`,
		TranID: 2968885920,
	}, rows[1])
	s.Equal(int32(2), dividend.Total, `Total`)
}

func (s *assetDividendServiceTestSuite) assertDividendEqual(e, a *DividendResponse) {
	r := s.r()
	r.Equal(e.ID, a.ID, `ID`)
	r.Equal(e.Amount, a.Amount, `Amount`)
	r.Equal(e.Amount, a.Amount, `Amount`)
	r.Equal(e.Info, a.Info, `Info`)
	r.Equal(e.Asset, a.Asset, `Asset`)
	r.Equal(e.Time, a.Time, `Time`)
	r.Equal(e.TranID, a.TranID, `TranID`)
}
