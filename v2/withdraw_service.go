package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateWithdrawService submits a withdraw request.
//
// See https://developers.binance.com/docs/wallet/capital/withdraw
type CreateWithdrawService struct {
	c                  *Client
	coin               string
	withdrawOrderID    *string
	network            *string
	address            string
	addressTag         *string
	amount             string
	transactionFeeFlag *bool   // When making internal transfer, true for returning the fee to the destination account; false for returning the fee back to the departure account. Default false.
	name               *string // Description of the address. Address book cap is 200, space in name should be encoded into %20
	walletType         *int    // The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current "selected wallet" under wallet->Fiat and Spot/Funding->Deposit
}

// Coin sets the coin parameter (MANDATORY).
func (s *CreateWithdrawService) Coin(v string) *CreateWithdrawService {
	s.coin = v
	return s
}

// WithdrawOrderID sets the withdrawOrderID parameter.
func (s *CreateWithdrawService) WithdrawOrderID(v string) *CreateWithdrawService {
	s.withdrawOrderID = &v
	return s
}

// Network sets the network parameter.
func (s *CreateWithdrawService) Network(v string) *CreateWithdrawService {
	s.network = &v
	return s
}

// Address sets the address parameter (MANDATORY).
func (s *CreateWithdrawService) Address(v string) *CreateWithdrawService {
	s.address = v
	return s
}

// AddressTag sets the addressTag parameter.
func (s *CreateWithdrawService) AddressTag(v string) *CreateWithdrawService {
	s.addressTag = &v
	return s
}

// Amount sets the amount parameter (MANDATORY).
func (s *CreateWithdrawService) Amount(v string) *CreateWithdrawService {
	s.amount = v
	return s
}

// TransactionFeeFlag sets the transactionFeeFlag parameter.
func (s *CreateWithdrawService) TransactionFeeFlag(v bool) *CreateWithdrawService {
	s.transactionFeeFlag = &v
	return s
}

// Name sets the name parameter.
func (s *CreateWithdrawService) Name(v string) *CreateWithdrawService {
	s.name = &v
	return s
}

func (s *CreateWithdrawService) WalletType(walletType int) *CreateWithdrawService {
	s.walletType = &walletType
	return s
}

// Do sends the request.
func (s *CreateWithdrawService) Do(ctx context.Context, opts ...RequestOption) (*CreateWithdrawResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/capital/withdraw/apply",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if v := s.withdrawOrderID; v != nil {
		r.setParam("withdrawOrderId", *v)
	}
	if v := s.network; v != nil {
		r.setParam("network", *v)
	}
	if v := s.addressTag; v != nil {
		r.setParam("addressTag", *v)
	}
	if v := s.transactionFeeFlag; v != nil {
		r.setParam("transactionFeeFlag", *v)
	}
	if v := s.name; v != nil {
		r.setParam("name", *v)
	}
	if s.walletType != nil {
		r.setParam("walletType", *s.walletType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &CreateWithdrawResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateWithdrawResponse represents a response from CreateWithdrawService.
type CreateWithdrawResponse struct {
	ID string `json:"id"`
}

// ListWithdrawsService fetches withdraw history.
//
// See https://developers.binance.com/docs/wallet/capital/withdraw-history
//   - network may not be in the response for old withdraw.
//   - Please notice the default startTime and endTime to make sure that time interval is within 0-90 days.
//   - If both startTime and endTimeare sent, time between startTimeand endTimemust be less than 90 days.
//   - If withdrawOrderId is sent, time between startTime and endTime must be less than 7 days.
//   - If withdrawOrderId is sent, startTime and endTime are not sent, will return last 7 days records by default.
//   - Maximum support idList number is 45.
type ListWithdrawsService struct {
	c               *Client
	coin            *string
	withdrawOrderId *string
	status          *int   // 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed)
	startTime       *int64 // Default: 90 days from current timestamp
	endTime         *int64 // Default: present timestamp
	offset          *int
	limit           *int    // Default: 1000, Max: 1000
	idList          *string // id list returned in the response of POST /sapi/v1/capital/withdraw/apply, separated by ,
}

// Coin sets the coin parameter.
func (s *ListWithdrawsService) Coin(coin string) *ListWithdrawsService {
	s.coin = &coin
	return s
}

// WithdrawOrderId sets the withdrawOrderId parameter.
func (s *ListWithdrawsService) WithdrawOrderId(withdrawOrderId string) *ListWithdrawsService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// Status sets the status parameter.
func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *ListWithdrawsService) Offset(offset int) *ListWithdrawsService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *ListWithdrawsService) Limit(limit int) *ListWithdrawsService {
	s.limit = &limit
	return s
}

// IdList id list returned in the response of POST /sapi/v1/capital/withdraw/apply, separated by ,
func (s *ListWithdrawsService) IdList(ids string) *ListWithdrawsService {
	s.idList = &ids
	return s
}

// Do sends the request.
func (s *ListWithdrawsService) Do(ctx context.Context, opts ...RequestOption) (res []*Withdraw, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/withdraw/history",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.withdrawOrderId != nil {
		r.setParam("withdrawOrderId", *s.withdrawOrderId)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.idList != nil {
		r.setParam("idList", *s.idList)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*Withdraw, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// Withdraw represents a single withdraw entry.
type Withdraw struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"` // Withdrawal id in Binance
	WithdrawOrderID string `json:"withdrawOrderId"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"` // 1 for internal transfer, 0 for external transfer
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"` // transaction fee
	ConfirmNo       int32  `json:"confirmNo"`
	Info            string `json:"info"` // reason for withdrawal failure
	TxID            string `json:"txId"` // withdrawal transaction id
	TxKey           string `json:"txKey"`
	CompleteTime    string `json:"completeTime"` // complete UTC time when user's asset is deduct from withdrawing, only if status =  6(success)
}
