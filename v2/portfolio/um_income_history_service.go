package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMIncomeHistoryService get UM income history
type GetUMIncomeHistoryService struct {
	c          *Client
	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	page       *int
	limit      *int
}

// Symbol set symbol
func (s *GetUMIncomeHistoryService) Symbol(symbol string) *GetUMIncomeHistoryService {
	s.symbol = &symbol
	return s
}

// IncomeType set income type
func (s *GetUMIncomeHistoryService) IncomeType(incomeType string) *GetUMIncomeHistoryService {
	s.incomeType = &incomeType
	return s
}

// StartTime set startTime
func (s *GetUMIncomeHistoryService) StartTime(startTime int64) *GetUMIncomeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetUMIncomeHistoryService) EndTime(endTime int64) *GetUMIncomeHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *GetUMIncomeHistoryService) Page(page int) *GetUMIncomeHistoryService {
	s.page = &page
	return s
}

// Limit set limit
func (s *GetUMIncomeHistoryService) Limit(limit int) *GetUMIncomeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetUMIncomeHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*Income, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/income",
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

// Income define income info
type Income struct {
	Symbol     string `json:"symbol"`     // trade symbol, if existing
	IncomeType string `json:"incomeType"` // income type
	Income     string `json:"income"`     // income amount
	Asset      string `json:"asset"`      // income asset
	Info       string `json:"info"`       // extra information
	Time       int64  `json:"time"`
	TranID     int64  `json:"tranId"`  // transaction id
	TradeID    string `json:"tradeId"` // trade id, if existing
}

// Constants for income types
const (
	IncomeTypeTransfer                 = "TRANSFER"
	IncomeTypeWelcomeBonus             = "WELCOME_BONUS"
	IncomeTypeRealizedPNL              = "REALIZED_PNL"
	IncomeTypeFundingFee               = "FUNDING_FEE"
	IncomeTypeCommission               = "COMMISSION"
	IncomeTypeInsuranceClear           = "INSURANCE_CLEAR"
	IncomeTypeReferralKickback         = "REFERRAL_KICKBACK"
	IncomeTypeCommissionRebate         = "COMMISSION_REBATE"
	IncomeTypeAPIRebate                = "API_REBATE"
	IncomeTypeContestReward            = "CONTEST_REWARD"
	IncomeTypeCrossCollateralTransfer  = "CROSS_COLLATERAL_TRANSFER"
	IncomeTypeOptionsPremiumFee        = "OPTIONS_PREMIUM_FEE"
	IncomeTypeOptionsSettleProfit      = "OPTIONS_SETTLE_PROFIT"
	IncomeTypeInternalTransfer         = "INTERNAL_TRANSFER"
	IncomeTypeAutoExchange             = "AUTO_EXCHANGE"
	IncomeTypeDeliveredSettlement      = "DELIVERED_SETTELMENT"
	IncomeTypeCoinSwapDeposit          = "COIN_SWAP_DEPOSIT"
	IncomeTypeCoinSwapWithdraw         = "COIN_SWAP_WITHDRAW"
	IncomeTypePositionLimitIncreaseFee = "POSITION_LIMIT_INCREASE_FEE"
)
