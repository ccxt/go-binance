package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// BNBTransferService transfer BNB in and out of UM
type BNBTransferService struct {
	c            *Client
	amount       string
	transferSide string
}

// Amount set amount
func (s *BNBTransferService) Amount(amount string) *BNBTransferService {
	s.amount = amount
	return s
}

// TransferSide set transfer side
func (s *BNBTransferService) TransferSide(transferSide string) *BNBTransferService {
	s.transferSide = transferSide
	return s
}

// Do send request
func (s *BNBTransferService) Do(ctx context.Context, opts ...RequestOption) (*BNBTransferResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/bnb-transfer",
		secType:  secTypeSigned,
	}
	r.setParam("amount", s.amount)
	r.setParam("transferSide", s.transferSide)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(BNBTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BNBTransferResponse define bnb transfer response
type BNBTransferResponse struct {
	TranID int64 `json:"tranId"` // transaction id
}

// Constants for transfer side
const (
	TransferSideToUM   = "TO_UM"
	TransferSideFromUM = "FROM_UM"
)
