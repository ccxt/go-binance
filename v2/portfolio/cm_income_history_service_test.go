package portfolio

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cmIncomeHistoryServiceTestSuite struct {
	baseTestSuite
}

func TestCMIncomeHistoryService(t *testing.T) {
	suite.Run(t, new(cmIncomeHistoryServiceTestSuite))
}

func (s *cmIncomeHistoryServiceTestSuite) TestGetCMIncomeHistory() {
	data := []byte(`[
		{
			"symbol": "",
			"incomeType": "TRANSFER",
			"income": "-0.37500000",
			"asset": "BTC",
			"info": "WITHDRAW",
			"time": 1570608000000,
			"tranId": 9689322392,
			"tradeId": ""
		},
		{
			"symbol": "BTCUSD_200925",
			"incomeType": "COMMISSION",
			"income": "-0.01000000",
			"asset": "BTC",
			"info": "",
			"time": 1570636800000,
			"tranId": 9689322392,
			"tradeId": "2059192"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200925"
	incomeType := CMIncomeTypeCommission
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", symbol)
		e.setParam("incomeType", incomeType)
		s.assertRequestEqual(e, r)
	})

	incomes, err := s.client.NewGetCMIncomeHistoryService().
		Symbol(symbol).
		IncomeType(incomeType).
		Do(newContext())

	s.r().NoError(err)
	s.r().Len(incomes, 2)

	s.r().Equal("", incomes[0].Symbol)
	s.r().Equal(CMIncomeTypeTransfer, incomes[0].IncomeType)
	s.r().Equal("-0.37500000", incomes[0].Income)
	s.r().Equal("BTC", incomes[0].Asset)
	s.r().Equal("WITHDRAW", incomes[0].Info)
	s.r().Equal(int64(1570608000000), incomes[0].Time)
	s.r().Equal(int64(9689322392), incomes[0].TranID)
	s.r().Equal("", incomes[0].TradeID)

	s.r().Equal("BTCUSD_200925", incomes[1].Symbol)
	s.r().Equal(CMIncomeTypeCommission, incomes[1].IncomeType)
	s.r().Equal("-0.01000000", incomes[1].Income)
	s.r().Equal("BTC", incomes[1].Asset)
	s.r().Equal("", incomes[1].Info)
	s.r().Equal(int64(1570636800000), incomes[1].Time)
	s.r().Equal(int64(9689322392), incomes[1].TranID)
	s.r().Equal("2059192", incomes[1].TradeID)
}
