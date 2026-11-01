package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type convertServiceTestSuite struct {
	baseTestSuite
}

func TestConvertService(t *testing.T) {
	suite.Run(t, new(convertServiceTestSuite))
}

func (s *convertServiceTestSuite) TestListConvertExchangeInfo() {
	data := []byte(`[
  {
    "fromAsset":"BTC",
    "toAsset":"USDT",
    "fromAssetMinAmount":"0.0004",
    "fromAssetMaxAmount":"50",
    "toAssetMinAmount":"20",
    "toAssetMaxAmount":"2500000"
  }
]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromAsset := "BTC"
	toAsset := "USDT"
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"fromAsset": fromAsset,
			"toAsset":   toAsset,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewListConvertExchangeInfoService().FromAsset(fromAsset).ToAsset(toAsset).Do(newContext())
	s.r().NoError(err)
	e := []ConvertExchangeInfo{{
		FromAsset:          fromAsset,
		ToAsset:            toAsset,
		FromAssetMinAmount: "0.0004",
		FromAssetMaxAmount: "50",
		ToAssetMinAmount:   "20",
		ToAssetMaxAmount:   "2500000",
	}}
	s.assertConvertExchangeInfoEqual(e, res)
}

func (s *convertServiceTestSuite) TestCreateConvertQuote() {
	data := []byte(`{
   "quoteId":"12415572564",
   "ratio":"38163.7",
   "inverseRatio":"0.0000262",
   "validTimestamp":1623319461670,
   "toAmount":"3816.37",
   "fromAmount":"0.1"
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromAsset := "BTC"
	toAsset := "USDT"
	fromAmount := "0.0004"
	toAmount := "20"
	res, err := s.client.NewCreateConvertQuoteService().FromAsset(fromAsset).ToAsset(toAsset).FromAmount(fromAmount).ToAmount(toAmount).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("12415572564", res.QuoteId)
	s.r().Equal("38163.7", res.Ratio)
	s.r().Equal("0.0000262", res.InverseRatio)
	s.r().Equal(int64(1623319461670), res.ValidTimestamp)
	s.r().Equal("3816.37", res.ToAmount)
	s.r().Equal("0.1", res.FromAmount)
}

func (s *convertServiceTestSuite) TestAcceptQuote() {
	data := []byte(`{
  "orderId":"933256278426274426",
  "createTime":1623381330472,
  "orderStatus":"PROCESS" 
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	quoteId := "12415572564"
	res, err := s.client.NewConvertAcceptService().QuoteId(quoteId).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("933256278426274426", res.OrderId)
	s.r().Equal(int64(1623381330472), res.CreateTime)
	s.r().Equal(ConvertAcceptStatusProcess, res.OrderStatus)
}

func (s *convertServiceTestSuite) TestGetConvertStatus() {
	data := []byte(`{
  "orderId":"933256278426274426",
  "orderStatus":"SUCCESS",
  "fromAsset":"BTC",
  "fromAmount":"0.00054414",
  "toAsset":"USDT",
  "toAmount":"20",
  "ratio":"36755",
  "inverseRatio":"0.00002721",
  "createTime":1623381330472
}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orderId := "933256278426274426"
	res, err := s.client.NewGetConvertStatusService().OrderId(orderId).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("933256278426274426", res.OrderId)
	s.r().Equal(ConvertAcceptStatusSuccess, res.OrderStatus)
	s.r().Equal("BTC", res.FromAsset)
	s.r().Equal("0.00054414", res.FromAmount)
	s.r().Equal("USDT", res.ToAsset)
	s.r().Equal("20", res.ToAmount)
	s.r().Equal("36755", res.Ratio)
	s.r().Equal("0.00002721", res.InverseRatio)
	s.r().Equal(int64(1623381330472), res.CreateTime)

}

func (s *convertServiceTestSuite) assertConvertExchangeInfoEqual(e, a []ConvertExchangeInfo) {
	r := s.r()
	r.Len(a, len(e))
	for i := range e {
		r.Equal(e[i].FromAsset, a[i].FromAsset, "FromAsset")
		r.Equal(e[i].ToAsset, a[i].ToAsset, "ToAsset")
		r.Equal(e[i].FromAssetMinAmount, a[i].FromAssetMinAmount, "FromAssetMinAmount")
		r.Equal(e[i].FromAssetMaxAmount, a[i].FromAssetMaxAmount, "FromAssetMaxAmount")
		r.Equal(e[i].ToAssetMinAmount, a[i].ToAssetMinAmount, "ToAssetMinAmount")
		r.Equal(e[i].ToAssetMaxAmount, a[i].ToAssetMaxAmount, "ToAssetMaxAmount")
	}
}
