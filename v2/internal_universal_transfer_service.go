package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// InternalUniversalTransferService Universal Transfer (For Master Account)
// https://developers.binance.com/docs/sub_account/asset-management/Universal-Transfer
type InternalUniversalTransferService struct {
	c               *Client
	fromEmail       *string // Mandatory No. Transfer from master account by default if fromEmail is not sent.
	toEmail         *string // Mandatory No. Transfer to master account by default if toEmail is not sent.
	fromAccountType string  // Mandatory Yes. "SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	toAccountType   string  // Mandatory Yes. "SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	clientTranId    *string
	symbol          *string
	asset           string
	amount          string
}

func (s *InternalUniversalTransferService) FromEmail(v string) *InternalUniversalTransferService {
	s.fromEmail = &v
	return s
}

func (s *InternalUniversalTransferService) ToEmail(v string) *InternalUniversalTransferService {
	s.toEmail = &v
	return s
}

func (s *InternalUniversalTransferService) FromAccountType(v string) *InternalUniversalTransferService {
	s.fromAccountType = v
	return s
}

func (s *InternalUniversalTransferService) ToAccountType(v string) *InternalUniversalTransferService {
	s.toAccountType = v
	return s
}

func (s *InternalUniversalTransferService) Symbol(v string) *InternalUniversalTransferService {
	s.symbol = &v
	return s
}

func (s *InternalUniversalTransferService) Asset(v string) *InternalUniversalTransferService {
	s.asset = v
	return s
}

func (s *InternalUniversalTransferService) Amount(v string) *InternalUniversalTransferService {
	s.amount = v
	return s
}

func (s *InternalUniversalTransferService) ClientTranId(v string) *InternalUniversalTransferService {
	s.clientTranId = &v
	return s
}

func (s *InternalUniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (*InternalUniversalTransferResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if v := s.fromEmail; v != nil {
		r.setParam("fromEmail", *v)
	}
	if v := s.toEmail; v != nil {
		r.setParam("toEmail", *v)
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	r.setParam("fromAccountType", s.fromAccountType)
	r.setParam("toAccountType", s.toAccountType)
	if v := s.clientTranId; v != nil {
		r.setParam("clientTranId", *v)
	}
	if v := s.symbol; v != nil {
		r.setParam("symbol", *v)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &InternalUniversalTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type InternalUniversalTransferResponse struct {
	ID           int64  `json:"tranId"`
	ClientTranID string `json:"clientTranId"`
}

// InternalUniversalTransferHistoryService Query Universal Transfer History (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#query-universal-transfer-history-for-master-account
type InternalUniversalTransferHistoryService struct {
	c            *Client
	fromEmail    *string
	toEmail      *string
	clientTranId *string
	startTime    *int64
	endTime      *int64
	page         *int
	limit        *int
}

func (s *InternalUniversalTransferHistoryService) FromEmail(v string) *InternalUniversalTransferHistoryService {
	s.fromEmail = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) ToEmail(v string) *InternalUniversalTransferHistoryService {
	s.toEmail = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) StartTime(v int64) *InternalUniversalTransferHistoryService {
	s.startTime = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) EndTime(v int64) *InternalUniversalTransferHistoryService {
	s.endTime = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) Page(v int) *InternalUniversalTransferHistoryService {
	s.page = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) Limit(v int) *InternalUniversalTransferHistoryService {
	s.limit = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) ClientTranId(v string) *InternalUniversalTransferHistoryService {
	s.clientTranId = &v
	return s
}

func (s *InternalUniversalTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res InternalUniversalTransferHistoryResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if v := s.fromEmail; v != nil {
		r.setParam("fromEmail", *v)
	}
	if v := s.toEmail; v != nil {
		r.setParam("toEmail", *v)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if v := s.clientTranId; v != nil {
		r.setParam("clientTranId", *v)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res.Result = make([]*InternalUniversalTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

type InternalUniversalTransferHistoryResponse struct {
	Result     []*InternalUniversalTransfer `json:"result"`
	TotalCount int                          `json:"totalCount"`
}

type InternalUniversalTransfer struct {
	TranId          int64  `json:"tranId"`
	ClientTranId    string `json:"clientTranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	CreateTimeStamp uint64 `json:"createTimeStamp"`
}
