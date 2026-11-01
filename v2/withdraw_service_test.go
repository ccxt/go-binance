package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type withdrawServiceTestSuite struct {
	baseTestSuite
}

func TestWithdrawService(t *testing.T) {
	suite.Run(t, new(withdrawServiceTestSuite))
}

func (s *withdrawServiceTestSuite) TestCreateWithdraw() {
	data := []byte(`
	{
		"id":"7213fea8e94b4a5593d507237e5a555b"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "USDT"
	withdrawOrderID := "testID"
	network := "ETH"
	address := "myaddress"
	addressTag := "xyz"
	amount := "0.01"
	transactionFeeFlag := true
	name := "eth"
	walletType := 0
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":               coin,
			"withdrawOrderId":    withdrawOrderID,
			"network":            network,
			"address":            address,
			"addressTag":         addressTag,
			"amount":             amount,
			"transactionFeeFlag": transactionFeeFlag,
			"name":               name,
			"walletType":         walletType,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateWithdrawService().
		Coin(coin).
		WithdrawOrderID(withdrawOrderID).
		Network(network).
		Address(address).
		AddressTag(addressTag).
		Amount(amount).
		TransactionFeeFlag(transactionFeeFlag).
		Name(name).
		WalletType(walletType).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("7213fea8e94b4a5593d507237e5a555b", res.ID)
}

func (s *withdrawServiceTestSuite) TestListWithdraws() {
	data := []byte(`[
    {
        "address": "0x94df8b352de7f46f64b01d3666bf6e936e44ce60",
        "amount": "8.91000000",
        "applyTime": "2019-10-12 11:12:02",
        "coin": "USDT",
        "id": "b6ae22b3aa844210a7041aee7589627c",
        "withdrawOrderId": "WITHDRAWtest123",
        "network": "ETH", 
        "transferType": 0,
        "status": 6,
        "transactionFee": "0.004",
		"confirmNo":3,
		"info":"The address is not valid. Please confirm with the recipient",
        "txId": "0xb5ef8c13b968a406cc62a93a8bd80f9e9a906ef1b3fcf20a2e48573c17659268",
		"txKey": "",
		"completeTime": "2025-03-06 00:00:00"
    },
    {
        "address": "1FZdVHtiBqMrWdjPyRPULCUceZPJ2WLCsB",
        "amount": "0.00150000",
        "applyTime": "2019-09-24 12:43:45",
        "coin": "BTC",
        "id": "156ec387f49b41df8724fa744fa82719",
        "network": "BTC",
        "status": 6,
        "transactionFee": "0.004",
        "transferType": 0,
		"confirmNo":2,
		"info":"",
        "txId": "60fd9007ebfddc753455f95fafa808c4302c836e4d1eebc5a132c36c1d8ac354",
		"txKey": "",
		"completeTime": "2025-03-06 00:00:00"
    }
]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "ETH"
	status := 0
	startTime := int64(1508198532000)
	endTime := int64(1508198532001)
	offset := 0
	limit := 1000
	idList := "1,22,333"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":      coin,
			"status":    status,
			"startTime": startTime,
			"endTime":   endTime,
			"offset":    offset,
			"limit":     limit,
			"idList":    "1,22,333",
		})
		s.assertRequestEqual(e, r)
	})

	withdraws, err := s.client.NewListWithdrawsService().
		Coin(coin).
		Status(status).
		StartTime(startTime).
		EndTime(endTime).
		Offset(offset).
		Limit(limit).
		IdList(idList).
		Do(newContext())
	r := s.r()
	r.NoError(err)

	s.Len(withdraws, 2)
	s.assertWithdrawEqual(&Withdraw{
		Address:         "0x94df8b352de7f46f64b01d3666bf6e936e44ce60",
		Amount:          "8.91000000",
		ApplyTime:       "2019-10-12 11:12:02",
		Coin:            "USDT",
		ID:              "b6ae22b3aa844210a7041aee7589627c",
		WithdrawOrderID: "WITHDRAWtest123",
		Network:         "ETH",
		TransferType:    0,
		Status:          6,
		TransactionFee:  "0.004",
		ConfirmNo:       3,
		Info:            "The address is not valid. Please confirm with the recipient",
		TxID:            "0xb5ef8c13b968a406cc62a93a8bd80f9e9a906ef1b3fcf20a2e48573c17659268",
		TxKey:           "",
		CompleteTime:    "2025-03-06 00:00:00",
	}, withdraws[0])
	s.assertWithdrawEqual(&Withdraw{
		Address:         "1FZdVHtiBqMrWdjPyRPULCUceZPJ2WLCsB",
		Amount:          "0.00150000",
		ApplyTime:       "2019-09-24 12:43:45",
		Coin:            "BTC",
		ID:              "156ec387f49b41df8724fa744fa82719",
		WithdrawOrderID: "",
		Network:         "BTC",
		TransferType:    0,
		Status:          6,
		TransactionFee:  "0.004",
		ConfirmNo:       2,
		Info:            "",
		TxID:            "60fd9007ebfddc753455f95fafa808c4302c836e4d1eebc5a132c36c1d8ac354",
		TxKey:           "",
		CompleteTime:    "2025-03-06 00:00:00",
	}, withdraws[1])
}

func (s *withdrawServiceTestSuite) assertWithdrawEqual(e, a *Withdraw) {
	r := s.r()
	r.Equal(e.Address, a.Address, "Address")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.ApplyTime, a.ApplyTime, "ApplyTime")
	r.Equal(e.Coin, a.Coin, "Coin")
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.WithdrawOrderID, a.WithdrawOrderID, "WithdrawOrderID")
	r.Equal(e.Network, a.Network, "Network")
	r.Equal(e.TransferType, a.TransferType, "TransferType")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TransactionFee, a.TransactionFee, "TransactionFee")
	r.Equal(e.ConfirmNo, a.ConfirmNo, "ConfirmNo")
	r.Equal(e.Info, a.Info, "Info")
	r.Equal(e.TxID, a.TxID, "TxID")
}
