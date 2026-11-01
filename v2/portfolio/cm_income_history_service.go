package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMIncomeHistoryService get CM income history
type GetCMIncomeHistoryService struct {
	c          *Client
	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	page       *int
	limit      *int
}

// Symbol set symbol
func (s *GetCMIncomeHistoryService) Symbol(symbol string) *GetCMIncomeHistoryService {
	s.symbol = &symbol
	return s
}

// IncomeType set income type
func (s *GetCMIncomeHistoryService) IncomeType(incomeType string) *GetCMIncomeHistoryService {
	s.incomeType = &incomeType
	return s
}

// StartTime set startTime
func (s *GetCMIncomeHistoryService) StartTime(startTime int64) *GetCMIncomeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetCMIncomeHistoryService) EndTime(endTime int64) *GetCMIncomeHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *GetCMIncomeHistoryService) Page(page int) *GetCMIncomeHistoryService {
	s.page = &page
	return s
}

// Limit set limit
func (s *GetCMIncomeHistoryService) Limit(limit int) *GetCMIncomeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetCMIncomeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*Income, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/income",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.incomeType != nil {
		r.setParam("incomeType", *s.incomeType)
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

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Income{}, err
	}
	res = make([]*Income, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Income{}, err
	}
	return res, nil
}

// Constants for CM income types
const (
	CMIncomeTypeTransfer            = "TRANSFER"
	CMIncomeTypeWelcomeBonus        = "WELCOME_BONUS"
	CMIncomeTypeFundingFee          = "FUNDING_FEE"
	CMIncomeTypeRealizedPNL         = "REALIZED_PNL"
	CMIncomeTypeCommission          = "COMMISSION"
	CMIncomeTypeInsuranceClear      = "INSURANCE_CLEAR"
	CMIncomeTypeDeliveredSettlement = "DELIVERED_SETTELMENT"
)
