package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateWithdrawService submits a withdraw request.
//
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw
type CreateWithdrawService struct {
	c                  *Client
	coin               string
	withdrawOrderID    *string
	network            *string
	address            string
	addressTag         *string
	amount             string
	transactionFeeFlag *bool
	name               *string
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
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw-history-supporting-network-user_data
type ListWithdrawsService struct {
	c               *Client
	coin            *string
	withdrawOrderId *string
	status          *int
	startTime       *int64
	endTime         *int64
	offset          *int
	limit           *int
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
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderId"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	ConfirmNo       int32  `json:"confirmNo"`
	Info            string `json:"info"`
	TxID            string `json:"txId"`
}
