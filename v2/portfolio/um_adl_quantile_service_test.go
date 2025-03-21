package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type umADLQuantileServiceTestSuite struct {
	baseTestSuite
}

func TestUMADLQuantileService(t *testing.T) {
	suite.Run(t, new(umADLQuantileServiceTestSuite))
}

func (s *umADLQuantileServiceTestSuite) TestADLQuantileWithHedgeMode() {
	data := []byte(`[
		{
			"symbol": "ETHUSDT",
			"adlQuantile": {
				"LONG": 3,
				"SHORT": 3,
				"HEDGE": 0
			}
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "ETHUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMADLQuantileService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	s.r().Equal("ETHUSDT", res[0].Symbol)
	s.r().Equal(3, res[0].ADLQuantile.LONG)
	s.r().Equal(3, res[0].ADLQuantile.SHORT)
}

func (s *umADLQuantileServiceTestSuite) TestADLQuantileWithOneWayMode() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"adlQuantile": {
				"LONG": 1,
				"SHORT": 2,
				"BOTH": 0
			}
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMADLQuantileService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	s.r().Equal("BTCUSDT", res[0].Symbol)
	s.r().Equal(1, res[0].ADLQuantile.LONG)
	s.r().Equal(2, res[0].ADLQuantile.SHORT)
}

func (s *umADLQuantileServiceTestSuite) TestADLQuantileForAllSymbols() {
	data := []byte(`[
		{
			"symbol": "ETHUSDT",
			"adlQuantile": {
				"LONG": 3,
				"SHORT": 3,
				"HEDGE": 0
			}
		},
		{
			"symbol": "BTCUSDT",
			"adlQuantile": {
				"LONG": 1,
				"SHORT": 2,
				"BOTH": 0
			}
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUMADLQuantileService().
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 2)
}
