package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type deliveryPriceServiceTestSuite struct {
	baseTestSuite
}

func TestDeliveryPriceService(t *testing.T) {
	suite.Run(t, new(deliveryPriceServiceTestSuite))
}

func (s *deliveryPriceServiceTestSuite) TestDeliveryPrice() {
	data := []byte(`[
 {
  "deliveryTime": 1711670400000,
  "deliveryPrice": 70151.5
 },
 {
  "deliveryTime": 1703808000000,
  "deliveryPrice": 42388.0
 }]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var pair string = "BTCUSDT"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("pair", pair)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewDeliveryPriceService().Pair(pair).Do(newContext())
	s.r().NoError(err)

	e := []*DeliveryPrice{
		{
			DeliveryTime:  1711670400000,
			DeliveryPrice: 70151.5},
		{
			DeliveryTime:  1703808000000,
			DeliveryPrice: 42388.0}}

	s.assertDeliveryPricesEqual(e, res)
}

func (s *deliveryPriceServiceTestSuite) assertDeliveryPriceEqual(e, a *DeliveryPrice) {
	r := s.r()
	r.Equal(e.DeliveryTime, a.DeliveryTime, "DeliveryTime")
	r.Equal(e.DeliveryPrice, a.DeliveryPrice, "DeliveryPrice")
}

func (s *deliveryPriceServiceTestSuite) assertDeliveryPricesEqual(e, a []*DeliveryPrice) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertDeliveryPriceEqual(e[i], a[i])
	}
}

type topLongShortAccountRatioServiceTestSuite struct {
	baseTestSuite
}

func TestTopLongShortAccountRatioService(t *testing.T) {
	suite.Run(t, new(topLongShortAccountRatioServiceTestSuite))
}

func (s *topLongShortAccountRatioServiceTestSuite) TestTopLongShortAccountRatio() {
	data := []byte(`[
 {
  "symbol": "BTCUSDT",
  "longAccount": "0.7106",
  "longShortRatio": "2.4554",
  "shortAccount": "0.2894",
  "timestamp": 1719517500000
 },
 {
  "symbol": "BTCUSDT",
  "longAccount": "0.7117",
  "longShortRatio": "2.4686",
  "shortAccount": "0.2883",
  "timestamp": 1719518400000
 }]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSDT"
	var period string = "15m"
	var limit uint32 = 30
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).setParam("period", period).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewTopLongShortAccountRatioService().Symbol(symbol).Period(period).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := []*TopLongShortAccountRatio{
		{
			Symbol:         "BTCUSDT",
			LongAccount:    "0.7106",
			LongShortRatio: "2.4554",
			ShortAccount:   "0.2894",
			Timestamp:      1719517500000},
		{
			Symbol:         "BTCUSDT",
			LongAccount:    "0.7117",
			LongShortRatio: "2.4686",
			ShortAccount:   "0.2883",
			Timestamp:      1719518400000}}

	s.assertTopLongShortAccountRatiosEqual(e, res)
}

func (s *topLongShortAccountRatioServiceTestSuite) assertTopLongShortAccountRatioEqual(e, a *TopLongShortAccountRatio) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LongShortRatio, a.LongShortRatio, "LongShortRatio")
	r.Equal(e.LongAccount, a.LongAccount, "LongAccount")
	r.Equal(e.ShortAccount, a.ShortAccount, "ShortAccount")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
}

func (s *topLongShortAccountRatioServiceTestSuite) assertTopLongShortAccountRatiosEqual(e, a []*TopLongShortAccountRatio) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertTopLongShortAccountRatioEqual(e[i], a[i])
	}
}

type topLongShortPositionRatioServiceTestSuite struct {
	baseTestSuite
}

func TestTopLongShortPositionRatioService(t *testing.T) {
	suite.Run(t, new(topLongShortPositionRatioServiceTestSuite))
}

func (s *topLongShortPositionRatioServiceTestSuite) TestTopLongShortPositionRatio() {
	data := []byte(`[
 {
  "symbol": "BTCUSDT",
  "longAccount": "0.6237",
  "longShortRatio": "1.6575",
  "shortAccount": "0.3763",
  "timestamp": 1719517500000
 },
 {
  "symbol": "BTCUSDT",
  "longAccount": "0.6251",
  "longShortRatio": "1.6676",
  "shortAccount": "0.3749",
  "timestamp": 1719518400000
 }]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSDT"
	var period string = "15m"
	var limit uint32 = 30
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).setParam("period", period).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewTopLongShortPositionRatioService().Symbol(symbol).Period(period).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := []*TopLongShortPositionRatio{
		{
			Symbol:         "BTCUSDT",
			LongAccount:    "0.6237",
			LongShortRatio: "1.6575",
			ShortAccount:   "0.3763",
			Timestamp:      1719517500000},
		{
			Symbol:         "BTCUSDT",
			LongAccount:    "0.6251",
			LongShortRatio: "1.6676",
			ShortAccount:   "0.3749",
			Timestamp:      1719518400000}}

	s.assertTopLongShortPositionRatiosEqual(e, res)
}

func (s *topLongShortPositionRatioServiceTestSuite) assertTopLongShortPositionRatioEqual(e, a *TopLongShortPositionRatio) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LongShortRatio, a.LongShortRatio, "LongShortRatio")
	r.Equal(e.LongAccount, a.LongAccount, "LongAccount")
	r.Equal(e.ShortAccount, a.ShortAccount, "ShortAccount")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
}

func (s *topLongShortPositionRatioServiceTestSuite) assertTopLongShortPositionRatiosEqual(e, a []*TopLongShortPositionRatio) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertTopLongShortPositionRatioEqual(e[i], a[i])
	}
}

type takerLongShortRatioServiceTestSuite struct {
	baseTestSuite
}

func TestTakerLongShortRatioService(t *testing.T) {
	suite.Run(t, new(takerLongShortRatioServiceTestSuite))
}

func (s *takerLongShortRatioServiceTestSuite) TestTakerLongShortRatio() {
	data := []byte(`[
 {
  "buySellRatio": "0.9500",
  "sellVol": "939.8980",
  "buyVol": "892.9140",
  "timestamp": 1719516600000
 },
 {
  "buySellRatio": "0.8046",
  "sellVol": "1146.2270",
  "buyVol": "922.2120",
  "timestamp": 1719517500000
 }]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var symbol string = "BTCUSDT"
	var period string = "15m"
	var limit uint32 = 30
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol).setParam("period", period).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewTakerLongShortRatioService().Symbol(symbol).Period(period).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := []*TakerLongShortRatio{
		{
			BuySellRatio: "0.9500",
			SellVol:      "939.8980",
			BuyVol:       "892.9140",
			Timestamp:    1719516600000},
		{
			BuySellRatio: "0.8046",
			SellVol:      "1146.2270",
			BuyVol:       "922.2120",
			Timestamp:    1719517500000}}

	s.assertTakerLongShortRatiosEqual(e, res)
}

func (s *takerLongShortRatioServiceTestSuite) assertTakerLongShortRatioEqual(e, a *TakerLongShortRatio) {
	r := s.r()
	r.Equal(e.BuySellRatio, a.BuySellRatio, "BuySellRatio")
	r.Equal(e.BuyVol, a.BuyVol, "BuyVol")
	r.Equal(e.SellVol, a.SellVol, "SellVol")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
}

func (s *takerLongShortRatioServiceTestSuite) assertTakerLongShortRatiosEqual(e, a []*TakerLongShortRatio) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertTakerLongShortRatioEqual(e[i], a[i])
	}
}

type basisServiceTestSuite struct {
	baseTestSuite
}

func TestBasisService(t *testing.T) {
	suite.Run(t, new(basisServiceTestSuite))
}

func (s *basisServiceTestSuite) TestBasis() {
	data := []byte(`[
 {
  "indexPrice": "62059.81936170",
  "contractType": "CURRENT_QUARTER",
  "basisRate": "-0.0002",
  "futuresPrice": "62046.6",
  "annualizedBasisRate": "-0.1016",
  "basis": "-13.21936170",
  "pair": "BTCUSDT",
  "timestamp": 1719499500000
 },
 {
  "indexPrice": "61734.27489362",
  "contractType": "CURRENT_QUARTER",
  "basisRate": "0.0000",
  "futuresPrice": "61733.5",
  "annualizedBasisRate": "0.0000",
  "basis": "-0.77489362",
  "pair": "BTCUSDT",
  "timestamp": 1719500400000
 }]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	var pair string = "BTCUSDT"
	var contractType string = "CURRENT_QUARTER"
	var period string = "15m"
	var limit uint32 = 30
	s.assertReq(func(r *request) {
		e := newRequest().setParam("pair", pair).setParam("contractType", contractType).setParam("period", period).setParam("limit", limit)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewBasisService().Pair(pair).ContractType(contractType).Period(period).Limit(limit).Do(newContext())
	s.r().NoError(err)

	e := []*Basis{
		{
			IndexPrice:          "62059.81936170",
			ContractType:        "CURRENT_QUARTER",
			BasisRate:           "-0.0002",
			FuturesPrice:        "62046.6",
			AnnualizedBasisRate: "-0.1016",
			Basis:               "-13.21936170",
			Pair:                "BTCUSDT",
			Timestamp:           1719499500000},
		{
			IndexPrice:          "61734.27489362",
			ContractType:        "CURRENT_QUARTER",
			BasisRate:           "0.0000",
			FuturesPrice:        "61733.5",
			AnnualizedBasisRate: "0.0000",
			Basis:               "-0.77489362",
			Pair:                "BTCUSDT",
			Timestamp:           1719500400000}}

	s.assertBasissEqual(e, res)
}

func (s *basisServiceTestSuite) assertBasisEqual(e, a *Basis) {
	r := s.r()
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.IndexPrice, a.IndexPrice, "IndexPrice")
	r.Equal(e.ContractType, a.ContractType, "ContractType")
	r.Equal(e.BasisRate, a.BasisRate, "BasisRate")
	r.Equal(e.FuturesPrice, a.FuturesPrice, "FuturesPrice")
	r.Equal(e.AnnualizedBasisRate, a.AnnualizedBasisRate, "AnnualizedBasisRate")
	r.Equal(e.Basis, a.Basis, "Basis")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
}

func (s *basisServiceTestSuite) assertBasissEqual(e, a []*Basis) {
	s.r().Len(e, len(a))
	for i := range e {
		s.assertBasisEqual(e[i], a[i])
	}
}
