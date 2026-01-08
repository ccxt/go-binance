package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type VipLoanService struct {
	c *Client
}

func (s *VipLoanService) InterestRate() *VipLoanInterestRateService {
	return &VipLoanInterestRateService{c: s.c}
}

func (s *VipLoanService) InterestRateHistory() *VipLoanInterestRateHistoryService {
	return &VipLoanInterestRateHistoryService{c: s.c}
}

func (s *VipLoanService) LoanableAssetData() *VipLoanLoanableAssetDataService {
	return &VipLoanLoanableAssetDataService{c: s.c}
}

func (s *VipLoanService) CollateralAssetData() *VipLoanCollateralAssetDataService {
	return &VipLoanCollateralAssetDataService{c: s.c}
}

func (s *VipLoanService) OngoingOrders() *VipLoanOngoingOrdersService {
	return &VipLoanOngoingOrdersService{c: s.c}
}

func (s *VipLoanService) RepaymentHistory() *VipLoanRepaymentHistoryService {
	return &VipLoanRepaymentHistoryService{c: s.c}
}

func (s *VipLoanService) AccruedInterest() *VipLoanAccruedInterestService {
	return &VipLoanAccruedInterestService{c: s.c}
}

func (s *VipLoanService) CollateralAccount() *VipLoanCollateralAccountService {
	return &VipLoanCollateralAccountService{c: s.c}
}

func (s *VipLoanService) ApplicationStatus() *VipLoanApplicationStatusService {
	return &VipLoanApplicationStatusService{c: s.c}
}

func (s *VipLoanService) Renew() *VipLoanRenewService {
	return &VipLoanRenewService{c: s.c}
}

func (s *VipLoanService) Repay() *VipLoanRepayService {
	return &VipLoanRepayService{c: s.c}
}

func (s *VipLoanService) Borrow() *VipLoanBorrowService {
	return &VipLoanBorrowService{c: s.c}
}

type VipLoanInterestRateService struct {
	c        *Client
	loanCoin string // Mandatory: YES, max 10 assets split by comma
}

func (s *VipLoanInterestRateService) LoanCoin(loanCoin string) *VipLoanInterestRateService {
	s.loanCoin = loanCoin
	return s
}

func (s *VipLoanInterestRateService) Do(ctx context.Context) (*VipLoanInterestRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/request/interestRate",
		secType:  secTypeSigned,
	}

	r.setParam("loanCoin", s.loanCoin)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanInterestRate)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanInterestRate []*VipLoanInterestRateElement

type VipLoanInterestRateElement struct {
	Asset                      string `json:"asset"`
	FlexibleDailyInterestRate  string `json:"flexibleDailyInterestRate"`
	FlexibleYearlyInterestRate string `json:"flexibleYearlyInterestRate"`
	Time                       string `json:"time"`
}

type VipLoanInterestRateHistoryService struct {
	c         *Client
	coin      string
	startTime *int64
	endTime   *int64
	current   *int64 // Mandatory: NO, default 1
	limit     *int64 // Mandatory: NO, default 10, max 100
}

func (s *VipLoanInterestRateHistoryService) Coin(coin string) *VipLoanInterestRateHistoryService {
	s.coin = coin
	return s
}

func (s *VipLoanInterestRateHistoryService) StartTime(startTime int64) *VipLoanInterestRateHistoryService {
	s.startTime = &startTime
	return s
}

func (s *VipLoanInterestRateHistoryService) EndTime(endTime int64) *VipLoanInterestRateHistoryService {
	s.endTime = &endTime
	return s
}

func (s *VipLoanInterestRateHistoryService) Current(current int64) *VipLoanInterestRateHistoryService {
	s.current = &current
	return s
}

func (s *VipLoanInterestRateHistoryService) Limit(limit int64) *VipLoanInterestRateHistoryService {
	s.limit = &limit
	return s
}

func (s *VipLoanInterestRateHistoryService) Do(ctx context.Context) (*VipLoanInterestRateHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/interestRateHistory",
		secType:  secTypeSigned,
	}

	r.setParam("coin", s.coin)

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanInterestRateHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanInterestRateHistoryResponse struct {
	Rows  []VipLoanInterestRateHistory `json:"rows"`
	Total int64                        `json:"total"`
}

type VipLoanInterestRateHistory struct {
	Coin                   string `json:"coin"`
	AnnualizedInterestRate string `json:"annualizedInterestRate"`
	Time                   string `json:"time"`
}

type VipLoanLoanableAssetDataService struct {
	c        *Client
	loanCoin *string
	vipLevel *int64 // Mandatory: NO, default to user vip level
}

func (s *VipLoanLoanableAssetDataService) LoanCoin(loanCoin string) *VipLoanLoanableAssetDataService {
	s.loanCoin = &loanCoin
	return s
}

func (s *VipLoanLoanableAssetDataService) VipLevel(vipLevel int64) *VipLoanLoanableAssetDataService {
	s.vipLevel = &vipLevel
	return s
}

func (s *VipLoanLoanableAssetDataService) Do(ctx context.Context) (*VipLoanLoanableAssetDataResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/loanable/data",
		secType:  secTypeSigned,
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.vipLevel != nil {
		r.setParam("vipLevel", *s.vipLevel)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanLoanableAssetDataResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanLoanableAssetDataResponse struct {
	Rows  []VipLoanLoanableAssetData `json:"rows"`
	Total int                        `json:"total"`
}

type VipLoanLoanableAssetData struct {
	LoanCoin                    string `json:"loanCoin"`
	FlexibleDailyInterestRate   string `json:"_flexibleDailyInterestRate"`
	FlexibleYearlyInterestRate  string `json:"_flexibleYearlyInterestRate"`
	ThirtyDayDailyInterestRate  string `json:"_30dDailyInterestRate"`
	ThirtyDayYearlyInterestRate string `json:"_30dYearlyInterestRate"`
	SixtyDayDailyInterestRate   string `json:"_60dDailyInterestRate"`
	SixtyDayYearlyInterestRate  string `json:"_60dYearlyInterestRate"`
	MinLimit                    string `json:"minLimit"`
	MaxLimit                    string `json:"maxLimit"`
	VipLevel                    int    `json:"vipLevel"`
}

type VipLoanCollateralAssetDataService struct {
	c              *Client
	collateralCoin *string // Mandatory: NO
}

func (s *VipLoanCollateralAssetDataService) CollateralCoin(collateralCoin string) *VipLoanCollateralAssetDataService {
	s.collateralCoin = &collateralCoin
	return s
}

func (s *VipLoanCollateralAssetDataService) Do(ctx context.Context) (*VipLoanCollateralAssetDataResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/collateral/data",
		secType:  secTypeSigned,
	}
	if s.collateralCoin != nil {
		r.setParam("collateralCoin", *s.collateralCoin)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanCollateralAssetDataResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanCollateralAssetDataResponse struct {
	Rows  []VipLoanCollateralAssetData `json:"rows"`
	Total int                          `json:"total"`
}

type VipLoanCollateralAssetData struct {
	CollateralCoin         string `json:"collateralCoin"`
	FirstCollateralRatio   string `json:"_1stCollateralRatio"`
	FirstCollateralRange   string `json:"_1stCollateralRange"`
	SecondCollateralRatio  string `json:"_2ndCollateralRatio"`
	SecondCollateralRange  string `json:"_2ndCollateralRange"`
	ThirdCollateralRatio   string `json:"_3rdCollateralRatio"`
	ThirdCollateralRange   string `json:"_3rdCollateralRange"`
	FourthCollateralRatio  string `json:"_4thCollateralRatio"`
	FourthCollateralRange  string `json:"_4thCollateralRange"`
	FifthCollateralRatio   string `json:"_5thCollateralRatio"`
	FifthCollateralRange   string `json:"_5thCollateralRange"`
	SixthCollateralRatio   string `json:"_6thCollateralRatio"`
	SixthCollateralRange   string `json:"_6thCollateralRange"`
	SeventhCollateralRatio string `json:"_7thCollateralRatio"`
	SeventhCollateralRange string `json:"_7thCollateralRange"`
	EighthCollateralRatio  string `json:"_8thCollateralRatio"`
	EighthCollateralRange  string `json:"_8thCollateralRange"`
	NinthCollateralRatio   string `json:"_9thCollateralRatio"`
	NinthCollateralRange   string `json:"_9thCollateralRange"`
}

type VipLoanOngoingOrdersService struct {
	c                   *Client
	orderId             *int64  // Mandatory: NO
	collateralAccountId *int64  // Mandatory: NO
	loanCoin            *string // Mandatory: NO
	collateralCoin      *string // Mandatory: NO
	current             *int64  // Mandatory: NO, default 1, max 100
	limit               *int64  // Mandatory: NO, default 10, max 100
}

func (s *VipLoanOngoingOrdersService) OrderId(orderId int64) *VipLoanOngoingOrdersService {
	s.orderId = &orderId
	return s
}

func (s *VipLoanOngoingOrdersService) CollateralAccountId(collateralAccountId int64) *VipLoanOngoingOrdersService {
	s.collateralAccountId = &collateralAccountId
	return s
}

func (s *VipLoanOngoingOrdersService) LoanCoin(loanCoin string) *VipLoanOngoingOrdersService {
	s.loanCoin = &loanCoin
	return s
}

func (s *VipLoanOngoingOrdersService) CollateralCoin(collateralCoin string) *VipLoanOngoingOrdersService {
	s.collateralCoin = &collateralCoin
	return s
}

func (s *VipLoanOngoingOrdersService) Current(current int64) *VipLoanOngoingOrdersService {
	s.current = &current
	return s
}

func (s *VipLoanOngoingOrdersService) Limit(limit int64) *VipLoanOngoingOrdersService {
	s.limit = &limit
	return s
}

func (s *VipLoanOngoingOrdersService) Do(ctx context.Context) (*VipLoanOngoingOrderResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/ongoing/orders",
		secType:  secTypeSigned,
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.collateralAccountId != nil {
		r.setParam("collateralAccountId", *s.collateralAccountId)
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.collateralCoin != nil {
		r.setParam("collateralCoin", *s.collateralCoin)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanOngoingOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanOngoingOrderResponse struct {
	Rows  []VipLoanOngoingOrder `json:"rows"`
	Total int                   `json:"total"`
}

type VipLoanOngoingOrder struct {
	OrderId                          int    `json:"orderId"`
	LoanCoin                         string `json:"loanCoin"`
	TotalDebt                        string `json:"totalDebt"`
	LoanRate                         string `json:"loanRate"`
	ResidualInterest                 string `json:"residualInterest"`
	CollateralAccountId              string `json:"collateralAccountId"`
	CollateralCoin                   string `json:"collateralCoin"`
	TotalCollateralValueAfterHaircut string `json:"totalCollateralValueAfterHaircut"`
	LockedCollateralValue            string `json:"lockedCollateralValue"`
	CurrentLTV                       string `json:"currentLTV"`
	ExpirationTime                   int64  `json:"expirationTime"`
	LoanDate                         string `json:"loanDate"`
	LoanTerm                         string `json:"loanTerm"`
}

type VipLoanRepaymentHistoryService struct {
	c         *Client
	orderId   *int64
	loanCoin  *string
	startTime *int64
	endTime   *int64
	current   *int64
	limit     *int64
}

func (s *VipLoanRepaymentHistoryService) OrderId(orderId int64) *VipLoanRepaymentHistoryService {
	s.orderId = &orderId
	return s
}

func (s *VipLoanRepaymentHistoryService) LoanCoin(loanCoin string) *VipLoanRepaymentHistoryService {
	s.loanCoin = &loanCoin
	return s
}

func (s *VipLoanRepaymentHistoryService) StartTime(startTime int64) *VipLoanRepaymentHistoryService {
	s.startTime = &startTime
	return s
}

func (s *VipLoanRepaymentHistoryService) EndTime(endTime int64) *VipLoanRepaymentHistoryService {
	s.endTime = &endTime
	return s
}

func (s *VipLoanRepaymentHistoryService) Current(current int64) *VipLoanRepaymentHistoryService {
	s.current = &current
	return s
}

func (s *VipLoanRepaymentHistoryService) Limit(limit int64) *VipLoanRepaymentHistoryService {
	s.limit = &limit
	return s
}

func (s *VipLoanRepaymentHistoryService) Do(ctx context.Context) (*VipLoanRepaymentHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/repayment/history",
		secType:  secTypeSigned,
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanRepaymentHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanRepaymentHistoryResponse struct {
	Rows  []VipLoanRepaymentHistory `json:"rows"`
	Total int                       `json:"total"`
}

type VipLoanRepaymentHistory struct {
	LoanCoin       string `json:"loanCoin"`
	RepayAmount    string `json:"repayAmount"`
	CollateralCoin string `json:"collateralCoin"`
	RepayStatus    string `json:"repayStatus"`
	LoanDate       string `json:"loanDate"`
	RepayTime      string `json:"repayTime"`
	OrderId        string `json:"orderId"`
}

type VipLoanAccruedInterestService struct {
	c         *Client
	orderId   *int64
	loanCoin  *string
	startTime *int64
	endTime   *int64
	current   *int64
	limit     *int64
}

func (s *VipLoanAccruedInterestService) OrderId(orderId int64) *VipLoanAccruedInterestService {
	s.orderId = &orderId
	return s
}

func (s *VipLoanAccruedInterestService) LoanCoin(loanCoin string) *VipLoanAccruedInterestService {
	s.loanCoin = &loanCoin
	return s
}

func (s *VipLoanAccruedInterestService) StartTime(startTime int64) *VipLoanAccruedInterestService {
	s.startTime = &startTime
	return s
}

func (s *VipLoanAccruedInterestService) EndTime(endTime int64) *VipLoanAccruedInterestService {
	s.endTime = &endTime
	return s
}

func (s *VipLoanAccruedInterestService) Current(current int64) *VipLoanAccruedInterestService {
	s.current = &current
	return s
}

func (s *VipLoanAccruedInterestService) Limit(limit int64) *VipLoanAccruedInterestService {
	s.limit = &limit
	return s
}

func (s *VipLoanAccruedInterestService) Do(ctx context.Context) (*VipLoanAccruedInterestResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/accruedInterest",
		secType:  secTypeSigned,
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanAccruedInterestResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanAccruedInterestResponse struct {
	Rows  []VipLoanAccruedInterest `json:"rows"`
	Total int                      `json:"total"`
}

type VipLoanAccruedInterest struct {
	LoanCoin           string `json:"loanCoin"`
	PrincipalAmount    string `json:"principalAmount"`
	InterestAmount     string `json:"interestAmount"`
	AnnualInterestRate string `json:"annualInterestRate"`
	AccrualTime        int64  `json:"accrualTime"`
	OrderId            int64  `json:"orderId"`
}

type VipLoanCollateralAccountService struct {
	c                   *Client
	orderId             *int64
	collateralAccountId *int64
}

func (s *VipLoanCollateralAccountService) OrderId(orderId int64) *VipLoanCollateralAccountService {
	s.orderId = &orderId
	return s
}

func (s *VipLoanCollateralAccountService) CollateralAccountId(collateralAccountId int64) *VipLoanCollateralAccountService {
	s.collateralAccountId = &collateralAccountId
	return s
}

func (s *VipLoanCollateralAccountService) Do(ctx context.Context) (*VipLoanCollateralAccountResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/collateral/account",
		secType:  secTypeSigned,
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.collateralAccountId != nil {
		r.setParam("collateralAccountId", *s.collateralAccountId)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanCollateralAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanCollateralAccountResponse struct {
	Rows  []VipLoanCollateralAccount `json:"rows"`
	Total int                        `json:"total"`
}

type VipLoanCollateralAccount struct {
	CollateralAccountId string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
}

type VipLoanApplicationStatusService struct {
	c       *Client
	current *int64 // Mandatory: NO, default 1
	limit   *int64 // Mandatory: NO, default 10, Max 100
}

func (s *VipLoanApplicationStatusService) Current(current int64) *VipLoanApplicationStatusService {
	s.current = &current
	return s
}

func (s *VipLoanApplicationStatusService) Limit(limit int64) *VipLoanApplicationStatusService {
	s.limit = &limit
	return s
}

func (s *VipLoanApplicationStatusService) Do(ctx context.Context) (*VipLoanApplicationStatusResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/request/data",
		secType:  secTypeSigned,
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanApplicationStatusResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanApplicationStatusResponse struct {
	Rows  []VipLoanApplicationStatus `json:"rows"`
	Total int                        `json:"total"`
}

type VipLoanApplicationStatus struct {
	LoanAccountId       string `json:"loanAccountId"`
	OrderId             string `json:"orderId"`
	RequestId           string `json:"requestId"`
	LoanCoin            string `json:"loanCoin"`
	LoanAmount          string `json:"loanAmount"`
	CollateralAccountId string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
	LoanTerm            string `json:"loanTerm"`
	Status              string `json:"status"`
	LoanDate            string `json:"loanDate"`
}

type VipLoanRenewService struct {
	c        *Client
	orderId  int64 // Mandatory: YES
	loanTerm int64 // Mandatory: YES, 30 or 60 days
}

func (s *VipLoanRenewService) OrderId(orderId int64) *VipLoanRenewService {
	s.orderId = orderId
	return s
}

func (s *VipLoanRenewService) LoanTerm(loanTerm int64) *VipLoanRenewService {
	s.loanTerm = loanTerm
	return s
}

func (s *VipLoanRenewService) Do(ctx context.Context, opts ...RequestOption) (*VipLoanRenew, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/vip/renew",
		secType:  secTypeSigned,
	}
	m := params{
		"orderId":  s.orderId,
		"loanTerm": s.loanTerm,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanRenew)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanRenew struct {
	LoanAccountId       string `json:"loanAccountId"`
	LoanCoin            string `json:"loanCoin"`
	LoanAmount          string `json:"loanAmount"`
	CollateralAccountId string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
	LoanTerm            string `json:"loanTerm"`
}

type VipLoanRepayService struct {
	c       *Client
	orderId int64
	amount  float64
}

func (s *VipLoanRepayService) OrderId(orderId int64) *VipLoanRepayService {
	s.orderId = orderId
	return s
}

func (s *VipLoanRepayService) Amount(amount float64) *VipLoanRepayService {
	s.amount = amount
	return s
}

func (s *VipLoanRepayService) Do(ctx context.Context, opts ...RequestOption) (*VipLoanRepay, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/vip/repay",
		secType:  secTypeSigned,
	}
	m := params{
		"orderId": s.orderId,
		"amount":  s.amount,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanRepay)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanRepay struct {
	LoanCoin           string `json:"loanCoin"`
	RepayAmount        string `json:"repayAmount"`
	RemainingPrincipal string `json:"remainingPrincipal"`
	RemainingInterest  string `json:"remainingInterest"`
	CollateralCoin     string `json:"collateralCoin"`
	CurrentLTV         string `json:"currentLTV"`
	RepayStatus        string `json:"repayStatus"`
}

type VipLoanBorrowService struct {
	c                   *Client
	loanAccountId       int64
	loanCoin            string // Mandatory: YES
	loanAmount          float64
	collateralAccountId string // Mandatory: YES, accounts split by comma
	collateralCoin      string // Mandatory: YES, coins split by comma
	isFlexibleRate      bool   // Mandatory: YES, TRUE: flexible rate; FALSE: fixed rate
	loanTerm            *int64 // Mandatory: YES for fixed interest, No for floating interest, 30 or 60
}

func (s *VipLoanBorrowService) LoanAccountId(loanAccountId int64) *VipLoanBorrowService {
	s.loanAccountId = loanAccountId
	return s
}

func (s *VipLoanBorrowService) LoanCoin(loanCoin string) *VipLoanBorrowService {
	s.loanCoin = loanCoin
	return s
}

func (s *VipLoanBorrowService) LoanAmount(loanAmount float64) *VipLoanBorrowService {
	s.loanAmount = loanAmount
	return s
}

func (s *VipLoanBorrowService) CollateralAccountId(collateralAccountId string) *VipLoanBorrowService {
	s.collateralAccountId = collateralAccountId
	return s
}

func (s *VipLoanBorrowService) CollateralCoin(collateralCoin string) *VipLoanBorrowService {
	s.collateralCoin = collateralCoin
	return s
}

func (s *VipLoanBorrowService) IsFlexibleRate(isFlexibleRate bool) *VipLoanBorrowService {
	s.isFlexibleRate = isFlexibleRate
	return s
}

func (s *VipLoanBorrowService) LoanTerm(loanTerm int64) *VipLoanBorrowService {
	s.loanTerm = &loanTerm
	return s
}

func (s *VipLoanBorrowService) Do(ctx context.Context, opts ...RequestOption) (*VipLoanBorrow, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/vip/borrow",
		secType:  secTypeSigned,
	}

	r.setParam("loanAccountId", s.loanAccountId)

	r.setParam("loanCoin", s.loanCoin)

	r.setParam("loanAmount", s.loanAmount)

	r.setParam("collateralAccountId", s.collateralAccountId)

	r.setParam("collateralCoin", s.collateralCoin)

	r.setParam("isFlexibleRate", s.isFlexibleRate)

	if s.loanTerm != nil {
		r.setParam("loanTerm", *s.loanTerm)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(VipLoanBorrow)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type VipLoanBorrow struct {
	LoanAccountId       string `json:"loanAccountId"`
	RequestId           string `json:"requestId"`
	LoanCoin            string `json:"loanCoin"`
	IsFlexibleRate      string `json:"isFlexibleRate"`
	LoanAmount          string `json:"loanAmount"`
	CollateralAccountId string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
	LoanTerm            string `json:"loanTerm"`
}
