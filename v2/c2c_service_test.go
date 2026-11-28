package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type c2cServiceTestSuite struct {
	baseTestSuite
}

func TestC2CService(t *testing.T) {
	suite.Run(t, new(c2cServiceTestSuite))
}

func (s *c2cServiceTestSuite) TestC2CTradeHistory() {
	data := []byte(`{
		"code": "000000",
   		"message": "success",
   		"data": [
   		{
      		"orderNumber":"20219644646554779648",
      		"advNo": "11218246497340923904",
      		"tradeType": "SELL",  
      		"asset": "BUSD", 
      		"fiat": "CNY",  
      		"fiatSymbol": "￥",
      		"amount": "5000.00000000",
      		"totalPrice": "33400.00000000",
      		"unitPrice": "6.68",
      		"orderStatus": "COMPLETED",
      		"createTime": 1619361369000,
      		"commission": "0",
      		"counterPartNickName": "ab***",
      		"advertisementRole": "TAKER"        
     	}
   		],
   		"total": 1,
   		"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	tradeType := SideTypeSell
	startTimestamp := int64(1619308800000)
	endTimestamp := int64(1619395200000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"tradeType":      tradeType,
			"startTimestamp": startTimestamp,
			"endTimestamp":   endTimestamp,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewC2CTradeHistoryService().
		TradeType(tradeType).
		StartTimestamp(startTimestamp).
		EndTime(endTimestamp).
		Do(newContext())
	s.r().NoError(err)
	e := &C2CTradeHistory{
		Code:    "000000",
		Message: "success",
		Data: []C2CRecord{
			{
				OrderNumber:         "20219644646554779648",
				AdvNo:               "11218246497340923904",
				TradeType:           "SELL",
				Asset:               "BUSD",
				Fiat:                "CNY",
				FiatSymbol:          "￥",
				Amount:              "5000.00000000",
				TotalPrice:          "33400.00000000",
				UnitPrice:           "6.68",
				OrderStatus:         "COMPLETED",
				CreateTime:          1619361369000,
				Commission:          "0",
				CounterPartNickName: "ab***",
				AdvertisementRole:   "TAKER",
			},
		},
		Total:   1,
		Success: true,
	}
	s.assertC2CTradeHistoryEqual(e, res)
}

func (s *c2cServiceTestSuite) assertC2CTradeHistoryEqual(e, a *C2CTradeHistory) {
	r := s.r()
	r.Equal(e.Code, a.Code, "Code")
	r.Equal(e.Message, a.Message, "Message")
	r.Equal(e.Total, a.Total, "Total")
	r.Equal(e.Success, a.Success, "Success")

	r.Len(a.Data, len(e.Data))
	for i := 0; i < len(a.Data); i++ {
		s.assertC2CTradeHistoryRecordEqual(&e.Data[i], &a.Data[i])
	}
}

func (s *c2cServiceTestSuite) assertC2CTradeHistoryRecordEqual(e, a *C2CRecord) {
	r := s.r()
	r.Equal(e.OrderNumber, a.OrderNumber, "OrderNumber")
	r.Equal(e.AdvNo, a.AdvNo, "AdvNo")
	r.Equal(e.TradeType, a.TradeType, "TradeType")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Fiat, a.Fiat, "Fiat")
	r.Equal(e.FiatSymbol, a.FiatSymbol, "FiatSymbol")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.TotalPrice, a.TotalPrice, "TotalPrice")
	r.Equal(e.UnitPrice, a.UnitPrice, "UnitPrice")
	r.Equal(e.OrderStatus, a.OrderStatus, "OrderStatus")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CounterPartNickName, a.CounterPartNickName, "CounterPartNickName")
	r.Equal(e.AdvertisementRole, a.AdvertisementRole, "AdvertisementRole")
}
