package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type DualInvestmentService struct {
	c *Client
}

func (s DualInvestmentService) ListProductService() *ListDualInvestmentProductService {
	return &ListDualInvestmentProductService{c: s.c}
}

func (s DualInvestmentService) ListPositionService() *ListDualInvestmentPositionService {
	return &ListDualInvestmentPositionService{c: s.c}
}

func (s DualInvestmentService) GetAccountService() *GetDualInvestmentAccountsService {
	return &GetDualInvestmentAccountsService{c: s.c}
}

func (s DualInvestmentService) SubscribeService() *SubscribeDualInvestmentService {
	return &SubscribeDualInvestmentService{c: s.c}
}

func (s DualInvestmentService) EditAutoCompoundStatusService() *EditAutoCompoundStatusDualInvestmentService {
	return &EditAutoCompoundStatusDualInvestmentService{c: s.c}
}

// ---------------------------------------------------------------
type DualInvestmentOptionType string

const (
	DualInvestmentOptionTypeCall DualInvestmentOptionType = "CALL"
	DualInvestmentOptionTypePut  DualInvestmentOptionType = "PUT"
)

type ListDualInvestmentProductService struct {
	c             *Client
	optionType    DualInvestmentOptionType
	exercisedCoin string
	investCoin    string

	pageSize  *int
	pageIndex *int
}

type DualInvestmentProductListResponse struct {
	Total int                     `json:"total"`
	List  []DualInvestmentProduct `json:"list"`
}

type DualInvestmentProduct struct {
	ID                   string                   `json:"id"`
	InvestCoin           string                   `json:"investCoin"`
	ExercisedCoin        string                   `json:"exercisedCoin"`
	StrikePrice          string                   `json:"strikePrice"`
	Duration             int                      `json:"duration"`
	SettleDate           int64                    `json:"settleDate"`
	PurchaseDecimal      int                      `json:"purchaseDecimal"`
	PurchaseEndTime      int64                    `json:"purchaseEndTime"`
	CanPurchase          bool                     `json:"canPurchase"`
	APR                  string                   `json:"apr"`
	OrderID              int64                    `json:"orderId"`
	MinAmount            string                   `json:"minAmount"`
	MaxAmount            string                   `json:"maxAmount"`
	CreateTimestamp      int64                    `json:"createTimestamp"`
	OptionType           DualInvestmentOptionType `json:"optionType"`
	IsAutoCompoundEnable bool                     `json:"isAutoCompoundEnable"`
	AutoCompoundPlanList []string                 `json:"autoCompoundPlanList"`
}

func (s *ListDualInvestmentProductService) OptionType(optionType DualInvestmentOptionType) *ListDualInvestmentProductService {
	s.optionType = optionType
	return s
}
func (s *ListDualInvestmentProductService) ExercisedCoin(exercisedCoin string) *ListDualInvestmentProductService {
	s.exercisedCoin = exercisedCoin
	return s
}
func (s *ListDualInvestmentProductService) InvestCoin(investCoin string) *ListDualInvestmentProductService {
	s.investCoin = investCoin
	return s
}
func (s *ListDualInvestmentProductService) PageSize(pageSize int) *ListDualInvestmentProductService {
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	s.pageSize = &pageSize
	return s
}
func (s *ListDualInvestmentProductService) PageIndex(pageIndex int) *ListDualInvestmentProductService {
	if pageIndex < 1 {
		pageIndex = 1
	}
	s.pageIndex = &pageIndex
	return s
}

// https://developers.binance.com/docs/dual_investment/quick-start
func (s *ListDualInvestmentProductService) Do(ctx context.Context, opts ...RequestOption) (res *DualInvestmentProductListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/dci/product/list",
		secType:  secTypeSigned,
	}
	r.setParam("optionType", s.optionType)
	r.setParam("exercisedCoin", s.exercisedCoin)
	r.setParam("investCoin", s.investCoin)
	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}
	if s.pageIndex != nil {
		r.setParam("pageIndex", *s.pageIndex)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(DualInvestmentProductListResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ListDualInvestmentPositionStatus string

const (
	ListDualInvestmentPositionStatusPending         ListDualInvestmentPositionStatus = "PENDING"
	ListDualInvestmentPositionStatusPurchaseSuccess ListDualInvestmentPositionStatus = "PURCHASE_SUCCESS"
	ListDualInvestmentPositionStatusSettled         ListDualInvestmentPositionStatus = "SETTLED"
	ListDualInvestmentPositionStatusPurchaseFail    ListDualInvestmentPositionStatus = "PURCHASE_FAIL"
	ListDualInvestmentPositionStatusRefunding       ListDualInvestmentPositionStatus = "REFUNDING"
	ListDualInvestmentPositionStatusRefundSuccess   ListDualInvestmentPositionStatus = "REFUND_SUCCESS"
	ListDualInvestmentPositionStatusSettling        ListDualInvestmentPositionStatus = "SETTLING"
)

type ListDualInvestmentPositionService struct {
	c         *Client
	status    *ListDualInvestmentPositionStatus
	pageSize  *int
	pageIndex *int
}

type ListDualInvestmentPositionResponse struct {
	Total int                          `json:"total"`
	List  []ListDualInvestmentPosition `json:"list"`
}
type ListDualInvestmentPosition struct {
	ID                 string                           `json:"id"`
	InvestCoin         string                           `json:"investCoin"`
	ExercisedCoin      string                           `json:"exercisedCoin"`
	SubscriptionAmount string                           `json:"subscriptionAmount"`
	StrikePrice        string                           `json:"strikePrice"`
	Duration           int                              `json:"duration"`
	SettleDate         int64                            `json:"settleDate"`
	PurchaseStatus     ListDualInvestmentPositionStatus `json:"purchaseStatus"`
	APR                string                           `json:"apr"`
	OrderID            int64                            `json:"orderId"`
	PurchaseEndTime    int64                            `json:"purchaseEndTime"`
	OptionType         DualInvestmentOptionType         `json:"optionType"`
	AutoCompoundPlan   DualInvestmentCompoundPlan       `json:"autoCompoundPlan"`

	// Settlement fields (only present for SETTLED positions)
	SettlePrice  *string `json:"settlePrice,omitempty"`  // Actual market price at settlement
	SettleAmount *string `json:"settleAmount,omitempty"` // Amount received after settlement
	SettleAsset  *string `json:"settleAsset,omitempty"`  // Coin received (investCoin or exercisedCoin)
	IsExercised  *bool   `json:"isExercised,omitempty"`  // Whether the option was exercised
}

func (s *ListDualInvestmentPositionService) Status(status ListDualInvestmentPositionStatus) *ListDualInvestmentPositionService {
	s.status = &status
	return s
}

func (s *ListDualInvestmentPositionService) PageSize(pageSize int) *ListDualInvestmentPositionService {
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	s.pageSize = &pageSize
	return s
}
func (s *ListDualInvestmentPositionService) PageIndex(pageIndex int) *ListDualInvestmentPositionService {
	if pageIndex < 1 {
		pageIndex = 1
	}
	s.pageIndex = &pageIndex
	return s
}

func (s *ListDualInvestmentPositionService) Do(ctx context.Context, opts ...RequestOption) (res *ListDualInvestmentPositionResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/dci/product/positions",
		secType:  secTypeSigned,
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}
	if s.pageIndex != nil {
		r.setParam("pageIndex", *s.pageIndex)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ListDualInvestmentPositionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetDualInvestmentAccountsService struct {
	c *Client
}

type GetDualInvestmentAccountsResp struct {
	TotalAmountInBTC  string `json:"totalAmountInBTC"`
	TotalAmountInUSDT string `json:"totalAmountInUSDT"`
}

func (s *GetDualInvestmentAccountsService) Do(ctx context.Context, opts ...RequestOption) (res *GetDualInvestmentAccountsResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/dci/product/accounts",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetDualInvestmentAccountsResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DualInvestmentCompoundPlan string

const (
	DualInvestmentCompoundPlanNone     DualInvestmentCompoundPlan = "NONE"     // disable auto compound
	DualInvestmentCompoundPlanStandard DualInvestmentCompoundPlan = "STANDARD" // basic plan
	DualInvestmentCompoundPlanAdvanced DualInvestmentCompoundPlan = "ADVANCED" // advance plan
)

type SubscribeDualInvestmentService struct {
	c                *Client
	id               string
	orderId          int64
	depositAmount    string
	autoCompoundPlan DualInvestmentCompoundPlan
}

type SubscribeDualInvestmentResp struct {
	PositionID         int64                            `json:"positionId"`
	InvestCoin         string                           `json:"investCoin"`
	ExercisedCoin      string                           `json:"exercisedCoin"`
	SubscriptionAmount string                           `json:"subscriptionAmount"`
	Duration           int                              `json:"duration"`
	AutoCompoundPlan   DualInvestmentCompoundPlan       `json:"autoCompoundPlan"` // 自动复投计划，当为NONE时不显示
	StrikePrice        string                           `json:"strikePrice"`
	SettleDate         int64                            `json:"settleDate"`
	PurchaseStatus     ListDualInvestmentPositionStatus `json:"purchaseStatus"`
	APR                string                           `json:"apr"`
	OrderID            int64                            `json:"orderId"`
	PurchaseTime       int64                            `json:"purchaseTime"`
	OptionType         DualInvestmentOptionType         `json:"optionType"`
}

func (s *SubscribeDualInvestmentService) ID(id string) *SubscribeDualInvestmentService {
	s.id = id
	return s
}

func (s *SubscribeDualInvestmentService) OrderID(orderId int64) *SubscribeDualInvestmentService {
	s.orderId = orderId
	return s
}

func (s *SubscribeDualInvestmentService) DepositAmount(depositAmount string) *SubscribeDualInvestmentService {
	s.depositAmount = depositAmount
	return s
}

func (s *SubscribeDualInvestmentService) AutoCompoundPlan(autoCompoundPlan DualInvestmentCompoundPlan) *SubscribeDualInvestmentService {
	s.autoCompoundPlan = autoCompoundPlan
	return s
}

func (s *SubscribeDualInvestmentService) Do(ctx context.Context, opts ...RequestOption) (res *SubscribeDualInvestmentResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/dci/product/subscribe",
		secType:  secTypeSigned,
	}
	r.setParam("id", s.id)
	r.setParam("orderId", s.orderId)
	r.setParam("depositAmount", s.depositAmount)
	r.setParam("autoCompoundPlan", s.autoCompoundPlan)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubscribeDualInvestmentResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EditAutoCompoundStatusDualInvestmentService struct {
	c                *Client
	positionId       string
	autoCompoundPlan DualInvestmentCompoundPlan
}
type EditAutoCompoundStatusDualInvestmentResp struct {
	PositionID       int64                      `json:"positionId"`
	AutoCompoundPlan DualInvestmentCompoundPlan `json:"autoCompoundPlan"`
}

func (s *EditAutoCompoundStatusDualInvestmentService) PositionID(positionID string) *EditAutoCompoundStatusDualInvestmentService {
	s.positionId = positionID
	return s
}
func (s *EditAutoCompoundStatusDualInvestmentService) AutoCompoundPlan(plan DualInvestmentCompoundPlan) *EditAutoCompoundStatusDualInvestmentService {
	s.autoCompoundPlan = plan
	return s
}

func (s *EditAutoCompoundStatusDualInvestmentService) Do(ctx context.Context, opts ...RequestOption) (res *EditAutoCompoundStatusDualInvestmentResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/dci/product/auto_compound/edit-status",
		secType:  secTypeSigned,
	}
	r.setParam("positionId", s.positionId)
	r.setParam("AutoCompoundPlan", s.autoCompoundPlan)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EditAutoCompoundStatusDualInvestmentResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
