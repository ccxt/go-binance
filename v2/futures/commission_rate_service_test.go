package futures

import (
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
)

type commissionRateServiceTestSuite struct {
	baseTestSuite
}

func TestCommissionRateService(t *testing.T) {
	suite.Run(t, new(commissionRateServiceTestSuite))
}

func (commissionRateService *commissionRateServiceTestSuite) TestCommissionRate() {
	data := []byte(`
		{
			"symbol": "BTCUSDT",
			"makerCommissionRate": "0.001",
			"takerCommissionRate": "0.001"
		}
	`)
	commissionRateService.mockDo(data, nil)
	defer commissionRateService.assertDo()

	symbol := "BTCUSDT"
	commissionRateService.assertReq(func(request *request) {
		requestParams := newSignedRequest().setParam("symbol", symbol)
		commissionRateService.assertRequestEqual(requestParams, request)
	})
	res, _, err := commissionRateService.client.NewCommissionRateService().Symbol(symbol).Do(newContext())
	commissionRateService.r().NoError(err)
	expectation := &CommissionRate{
		Symbol:              symbol,
		MakerCommissionRate: "0.001",
		TakerCommissionRate: "0.001",
	}
	commissionRateService.assertCommissionRateResponseEqual(expectation, res)
}

func (commissionRateServiceTestSuite *commissionRateServiceTestSuite) assertCommissionRateResponseEqual(expectation, assertedData *CommissionRate) {
	assertion := commissionRateServiceTestSuite.r()
	log.Printf("TEST 1 %#v\n", expectation)
	log.Printf("TEST 2 %#v\n", assertedData)
	log.Printf("TEST 3 %#v\n", assertion)
	assertion.Equal(expectation.Symbol, assertedData.Symbol, "Symbol")
	assertion.Equal(expectation.MakerCommissionRate, assertedData.MakerCommissionRate, "MakerCommissionRate")
	assertion.Equal(expectation.TakerCommissionRate, assertedData.TakerCommissionRate, "TakerCommissionRate")
}
