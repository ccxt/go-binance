package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListDepositsService fetches deposit history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-history-user_data
type ListDepositsService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int
	limit     *int
	txId      *string
}

// Coin sets the coin parameter.
func (s *ListDepositsService) Coin(coin string) *ListDepositsService {
	s.coin = &coin
	return s
}

// Status sets the status parameter.
func (s *ListDepositsService) Status(status int) *ListDepositsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsService) StartTime(startTime int64) *ListDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsService) EndTime(endTime int64) *ListDepositsService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *ListDepositsService) Offset(offset int) *ListDepositsService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *ListDepositsService) Limit(limit int) *ListDepositsService {
	s.limit = &limit
	return s
}

func (s *ListDepositsService) TxID(id string) *ListDepositsService {
	s.txId = &id
	return s
}

// Do sends the request.
func (s *ListDepositsService) Do(ctx context.Context, opts ...RequestOption) (res []*Deposit, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
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
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*Deposit, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// Deposit represents a single deposit entry.
type Deposit struct {
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxID          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int64  `json:"transferType"`
	UnlockConfirm int64  `json:"unlockConfirm"`
	ConfirmTimes  string `json:"confirmTimes"`
}

// GetDepositsAddressService retrieves the details of a deposit address.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-address-supporting-network-user_data
type GetDepositsAddressService struct {
	c       *Client
	coin    string
	network *string
}

// Coin sets the coin parameter (MANDATORY).
func (s *GetDepositsAddressService) Coin(coin string) *GetDepositsAddressService {
	s.coin = coin
	return s
}

// Network sets the network parameter.
func (s *GetDepositsAddressService) Network(network string) *GetDepositsAddressService {
	s.network = &network
	return s
}

// Do sends the request.
func (s *GetDepositsAddressService) Do(ctx context.Context, opts ...RequestOption) (*GetDepositAddressResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/address",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &GetDepositAddressResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetDepositAddressResponse represents a response from GetDepositsAddressService.
type GetDepositAddressResponse struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Coin    string `json:"coin"`
	URL     string `json:"url"`
}
