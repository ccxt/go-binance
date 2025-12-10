package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginAvailableInventoryServiceTestSuite struct {
	baseTestSuite
}

func TestMarginAvailableInventoryService(t *testing.T) {
	suite.Run(t, new(marginAvailableInventoryServiceTestSuite))
}

func (s *marginAvailableInventoryServiceTestSuite) TestMarginAvailableInventory() {
	data := []byte(`
	{
		"assets": {
			"MATIC": "100000000",
			"STPT": "100000000",
			"TVK": "100000000",
			"SHIB": "97409653"
		},
		"updateTime": 1699272487
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	marginType := "MARGIN"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"type": marginType,
		})
		s.assertRequestEqual(e, r)
	})

	inventory, err := s.client.NewMarginAvailableInventoryService().
		MarginType(marginType).
		Do(context.Background())
	r := s.r()
	r.NoError(err)

	s.Equal(int64(1699272487), inventory.UpdateTime)

	matic := inventory.Assets["MATIC"]
	s.Equal("100000000", matic)
}
