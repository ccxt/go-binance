package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SymbolConfigServiceTestSuite struct {
	baseTestSuite
}

func TestSymbolConfigService(t *testing.T) {
	suite.Run(t, new(SymbolConfigServiceTestSuite))
}

func (s *SymbolConfigServiceTestSuite) TestGetSymbolConfig() {
	data := []byte(`[
        {
            "symbol": "BTCUSDT",
            "marginType": "CROSSED",
            "isAutoAddMargin": false,
            "leverage": 21,
            "maxNotionalValue": "1000000"
        }
    ]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	configs, err := s.client.NewGetSymbolConfigService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	s.r().Len(configs, 1)
	s.r().Equal("BTCUSDT", configs[0].Symbol)
	s.r().Equal("CROSSED", configs[0].MarginType)
	s.r().Equal(false, configs[0].IsAutoAddMargin)
	s.r().Equal(21, configs[0].Leverage)
	s.r().Equal("1000000", configs[0].MaxNotionalValue)
}
