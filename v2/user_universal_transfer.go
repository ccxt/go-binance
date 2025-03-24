/*******************************************************************************
** @Author:					Alessandro Maioli <Nicemine>
** @Email:					alessandro.maioli@gmail.com
** @Date:					Wednesday 29 September 2021 - 18:51:45
** @Filename:				user_universal_transfer.go
**
** @Last modified by:
*******************************************************************************/

package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateUserUniversalTransferService submits a transfer request.
//
// See https://developers.binance.com/docs/wallet/asset/user-universal-transfer
type CreateUserUniversalTransferService struct {
	c          *Client
	types      UserUniversalTransferType
	asset      string
	amount     string // type: DECIMAL
	fromSymbol *string
	toSymbol   *string
}

// Coin sets the coin parameter (MANDATORY).
func (s *CreateUserUniversalTransferService) Type(v UserUniversalTransferType) *CreateUserUniversalTransferService {
	s.types = v
	return s
}

// Asset sets the Asset parameter (MANDATORY).
func (s *CreateUserUniversalTransferService) Asset(v string) *CreateUserUniversalTransferService {
	s.asset = v
	return s
}

// Amount sets the Amount parameter (MANDATORY).
func (s *CreateUserUniversalTransferService) Amount(v string) *CreateUserUniversalTransferService {
	s.amount = v
	return s
}

// FromSymbol sets the fromSymbol parameter
func (s *CreateUserUniversalTransferService) FromSymbol(v string) *CreateUserUniversalTransferService {
	s.fromSymbol = &v
	return s
}

// ToSymbol sets the toSymbol parameter
func (s *CreateUserUniversalTransferService) ToSymbol(v string) *CreateUserUniversalTransferService {
	s.toSymbol = &v
	return s
}

// Do sends the request.
func (s *CreateUserUniversalTransferService) Do(ctx context.Context) (*CreateUserUniversalTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}

	r.setParam("type", s.types)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if v := s.fromSymbol; v != nil {
		r.setParam("fromSymbol", *v)
	}
	if v := s.toSymbol; v != nil {
		r.setParam("toSymbol", *v)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateUserUniversalTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateUserUniversalTransferResponse represents a response from CreateUserUniversalTransferResponse.
type CreateUserUniversalTransferResponse struct {
	ID int64 `json:"tranId"`
}

// ListUserUniversalTransfer fetches transfer history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#query-user-universal-transfer-history-user_data
type ListUserUniversalTransferService struct {
	c            *Client
	transferType UserUniversalTransferType
	startTime    *int64
	endTime      *int64
	current      *int
	size         *int
	fromSymbol   *string
	toSymbol     *string
}

// Type sets the type parameter.
func (s *ListUserUniversalTransferService) Type(v UserUniversalTransferType) *ListUserUniversalTransferService {
	s.transferType = v
	return s
}

// StartTime sets the startTime parameter.
func (s *ListUserUniversalTransferService) StartTime(v int64) *ListUserUniversalTransferService {
	s.startTime = &v
	return s
}

// EndTime sets the startTime parameter.
func (s *ListUserUniversalTransferService) EndTime(v int64) *ListUserUniversalTransferService {
	s.endTime = &v
	return s
}

// Current sets the current parameter.
func (s *ListUserUniversalTransferService) Current(v int) *ListUserUniversalTransferService {
	s.current = &v
	return s
}

// Size sets the size parameter.
func (s *ListUserUniversalTransferService) Size(v int) *ListUserUniversalTransferService {
	s.size = &v
	return s
}

// FromSymbol set fromSymbol
func (s *ListUserUniversalTransferService) FromSymbol(v string) *ListUserUniversalTransferService {
	s.fromSymbol = &v
	return s
}

// ToSymbol set toSymbol
func (s *ListUserUniversalTransferService) ToSymbol(v string) *ListUserUniversalTransferService {
	s.toSymbol = &v
	return s
}

// // Do sends the request.
func (s *ListUserUniversalTransferService) Do(ctx context.Context) (res *UserUniversalTransferResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}
	r.setParam("type", s.transferType)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	if s.fromSymbol != nil {
		r.setParam("fromSymbol", *s.fromSymbol)
	}
	if s.toSymbol != nil {
		r.setParam("toSymbol", *s.toSymbol)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = new(UserUniversalTransferResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// // Withdraw represents a single withdraw entry.
type UserUniversalTransferResponse struct {
	Total   int64                    `json:"total"`
	Results []*UserUniversalTransfer `json:"rows"`
}

type UserUniversalTransfer struct {
	Asset     string                          `json:"asset"`
	Amount    string                          `json:"amount"`
	Type      UserUniversalTransferType       `json:"type"`
	Status    UserUniversalTransferStatusType `json:"status"`
	TranId    int64                           `json:"tranId"`
	Timestamp int64                           `json:"timestamp"`
}
