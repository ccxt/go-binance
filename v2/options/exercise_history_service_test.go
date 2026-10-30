package options

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExerciseHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestExerciseHistoryService(t *testing.T) {
	suite.Run(t, new(ExerciseHistoryServiceTestSuite))
}

func (s *ExerciseHistoryServiceTestSuite) TestExerciseHistory() {
	data := []byte(`[
		{
		 "symbol": "BTC-240529-67500-C",
		 "strikePrice": "67500",
		 "realStrikePrice": "68154.65503404",
		 "expiryDate": 1716969600000,
		 "strikeResult": "REALISTIC_VALUE_STRICKEN"
		}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"underlying": "BTCUSDT",
			"startTime":  uint64(1716940800000),
			"endTime":    uint64(1716969600000),
			"limit":      uint64(10),
		})
		s.assertRequestEqual(e, r)
	})
	ETs, err := s.client.NewExerciseHistoryService().Underlying("BTCUSDT").
		StartTime(1716940800000).EndTime(1716969600000).Limit(10).Do(newContext())
	targetETs := []*ExerciseHistory{
		{
			Symbol:          "BTC-240529-67500-C",
			StrikePrice:     "67500",
			RealStrikePrice: "68154.65503404",
			ExpiryDate:      1716969600000,
			StrikeResult:    "REALISTIC_VALUE_STRICKEN",
		},
	}

	s.r().Equal(err, nil, "err != nil")
	for i := range ETs {
		r := s.r()
		r.Equal(ETs[i].Symbol, targetETs[i].Symbol, "Symbol")
		r.Equal(ETs[i].StrikePrice, targetETs[i].StrikePrice, "StrikePrice")
		r.Equal(ETs[i].RealStrikePrice, targetETs[i].RealStrikePrice, "RealStrikePrice")
		r.Equal(ETs[i].ExpiryDate, targetETs[i].ExpiryDate, "ExpiryDate")
		r.Equal(ETs[i].StrikeResult, targetETs[i].StrikeResult, "StrikeResult")
	}
}
