package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmADLQuantileServiceTestSuite struct {
	baseTestSuite
}

func TestCMADLQuantileService(t *testing.T) {
	suite.Run(t, new(cmADLQuantileServiceTestSuite))
}

func (s *cmADLQuantileServiceTestSuite) TestADLQuantileWithHedgeMode() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_200925",
			"adlQuantile": {
				"LONG": 3,
				"SHORT": 3,
				"HEDGE": 0
			}
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMADLQuantileService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	s.r().Equal("BTCUSD_200925", res[0].Symbol)
	s.r().Equal(3, res[0].ADLQuantile.LONG)
	s.r().Equal(3, res[0].ADLQuantile.SHORT)
}

func (s *cmADLQuantileServiceTestSuite) TestADLQuantileWithOneWayMode() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_201225",
			"adlQuantile": {
				"LONG": 1,
				"SHORT": 2,
				"BOTH": 0
			}
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_201225"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCMADLQuantileService().
		Symbol(symbol).
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 1)
	s.r().Equal("BTCUSD_201225", res[0].Symbol)
	s.r().Equal(1, res[0].ADLQuantile.LONG)
	s.r().Equal(2, res[0].ADLQuantile.SHORT)
}

func (s *cmADLQuantileServiceTestSuite) TestADLQuantileForAllSymbols() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_200925",
			"adlQuantile": {
				"LONG": 3,
				"SHORT": 3,
				"HEDGE": 0
			}
		},
		{
			"symbol": "BTCUSD_201225",
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

	res, err := s.client.NewCMADLQuantileService().
		Do(newContext())
	s.r().NoError(err)
	s.r().Len(res, 2)
}
