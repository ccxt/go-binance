package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAllSwapPoolService get swap pool market data
type GetAllLiquidityPoolService struct {
	c *Client
}

// Do send request
func (s *GetAllLiquidityPoolService) Do(ctx context.Context, opts ...RequestOption) ([]*LiquidityPool, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := []*LiquidityPool{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type LiquidityPool struct {
	PoolId   int64    `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}

// GetLiquidityPoolDetailService get swap pool detail by pool id
type GetLiquidityPoolDetailService struct {
	c      *Client
	poolId *int64
}

// PoolId set poolId
func (s *GetLiquidityPoolDetailService) PoolId(poolId int64) *GetLiquidityPoolDetailService {
	s.poolId = &poolId
	return s
}

type LiquidityPoolDetail struct {
	PoolId     int64                 `json:"poolId"`
	PoolName   string                `json:"poolName"`
	UpdateTime int64                 `json:"updateTime"`
	Liquidity  map[string]string     `json:"liquidity"`
	Share      *PoolShareInformation `json:"share"`
}

type PoolShareInformation struct {
	ShareAmount     string            `json:"shareAmount"`
	SharePercentage string            `json:"sharePercentage"`
	Assets          map[string]string `json:"asset"`
}

// Do sends the request.
func (s *GetLiquidityPoolDetailService) Do(ctx context.Context) ([]*LiquidityPoolDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/liquidity",
		secType:  secTypeSigned,
	}
	if s.poolId != nil {
		r.setParam("poolId", *s.poolId)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := []*LiquidityPoolDetail{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddLiquidityPreviewService to preview the quote/base qty needed when adding assets to a liquidity pool with an estimated share after adding the liquidity
type AddLiquidityPreviewService struct {
	c             *Client
	poolId        *int64
	operationType *LiquidityOperationType
	quoteAsset    *string
	quoteQty      *float64
}

// PoolId set poolId
func (s *AddLiquidityPreviewService) PoolId(poolId int64) *AddLiquidityPreviewService {
	s.poolId = &poolId
	return s
}

// QuoteAsset set quoteAsset
func (s *AddLiquidityPreviewService) QuoteAsset(quoteAsset string) *AddLiquidityPreviewService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *AddLiquidityPreviewService) QuoteQty(quoteQty float64) *AddLiquidityPreviewService {
	s.quoteQty = &quoteQty
	return s
}

// OperationType set operationType
func (s *AddLiquidityPreviewService) OperationType(operationType LiquidityOperationType) *AddLiquidityPreviewService {
	s.operationType = &operationType
	return s
}

type AddLiquidityPreviewResponse struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"` // only existed when type is COMBINATION
	QuoteAmt   string `json:"quoteAmt"`
	BaseAmt    string `json:"baseAmt"` // only existed when type is COMBINATION
	Price      string `json:"price"`
	Share      string `json:"share"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

// Do sends the request.
func (s *AddLiquidityPreviewService) Do(ctx context.Context) (*AddLiquidityPreviewResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/addLiquidityPreview",
		secType:  secTypeSigned,
	}

	r.setParam("poolId", *s.poolId)
	r.setParam("type", *s.operationType)
	r.setParam("quoteAsset", *s.quoteAsset)
	r.setParam("quoteQty", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &AddLiquidityPreviewResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSwapQuoteService get the quote
type GetSwapQuoteService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *float64
}

// QuoteAsset set quoteAsset
func (s *GetSwapQuoteService) QuoteAsset(quoteAsset string) *GetSwapQuoteService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *GetSwapQuoteService) QuoteQty(quoteQty float64) *GetSwapQuoteService {
	s.quoteQty = &quoteQty
	return s
}

// BaseAsset set baseAsset
func (s *GetSwapQuoteService) BaseAsset(baseAsset string) *GetSwapQuoteService {
	s.baseAsset = &baseAsset
	return s
}

type GetSwapQuoteResponse struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"`
	QuoteQty   string `json:"quoteQty"`
	BaseQty    string `json:"baseQty"`
	Price      string `json:"price"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

// Do sends the request.
func (s *GetSwapQuoteService) Do(ctx context.Context) (*GetSwapQuoteResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/quote",
		secType:  secTypeSigned,
	}

	r.setParam("quoteAsset", *s.quoteAsset)
	r.setParam("baseAsset", *s.baseAsset)
	r.setParam("quoteQty", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &GetSwapQuoteResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SwapService swap tokens in liquidity pool
type SwapService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *float64
}

// QuoteAsset set quoteAsset
func (s *SwapService) QuoteAsset(quoteAsset string) *SwapService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *SwapService) QuoteQty(quoteQty float64) *SwapService {
	s.quoteQty = &quoteQty
	return s
}

// BaseAsset set baseAsset
func (s *SwapService) BaseAsset(baseAsset string) *SwapService {
	s.baseAsset = &baseAsset
	return s
}

type SwapResponse struct {
	SwapId int64 `json:"swapId"`
}

// Do sends the request.
func (s *SwapService) Do(ctx context.Context) (*SwapResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}

	r.setParam("quoteAsset", *s.quoteAsset)
	r.setParam("baseAsset", *s.baseAsset)
	r.setParam("quoteQty", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &SwapResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetUserSwapRecordsService struct {
	c          *Client
	swapId     *int64
	startTime  *int64
	endTime    *int64
	status     *SwappingStatus
	quoteAsset *string
	baseAsset  *string
	resultSize *int64
}

type SwapRecord struct {
	SwapId     int64          `json:"swapId"`
	SwapTime   int64          `json:"swapTime"`
	Status     SwappingStatus `json:"status"`
	QuoteAsset string         `json:"quoteAsset"`
	BaseAsset  string         `json:"baseAsset"`
	QuoteQty   string         `json:"quoteQty"`
	BaseQty    string         `json:"baseQty"`
	Price      string         `json:"price"`
	Fee        string         `json:"fee"`
}

// SwapId set swapId
func (s *GetUserSwapRecordsService) SwapId(swapId int64) *GetUserSwapRecordsService {
	s.swapId = &swapId
	return s
}

// StartTime set start time when swapping
func (s *GetUserSwapRecordsService) StartTime(startTime int64) *GetUserSwapRecordsService {
	s.startTime = &startTime
	return s
}

// EndTime set end time when swapping
func (s *GetUserSwapRecordsService) EndTime(endTime int64) *GetUserSwapRecordsService {
	s.endTime = &endTime
	return s
}

// Status set status we are query for
func (s *GetUserSwapRecordsService) Status(status SwappingStatus) *GetUserSwapRecordsService {
	s.status = &status
	return s
}

// QuoteAsset set quote asset
func (s *GetUserSwapRecordsService) QuoteAsset(quoteAsset string) *GetUserSwapRecordsService {
	s.quoteAsset = &quoteAsset
	return s
}

// SwapId set base asset
func (s *GetUserSwapRecordsService) BaseAsset(baseAsset string) *GetUserSwapRecordsService {
	s.baseAsset = &baseAsset
	return s
}

// ResultSize set the size will be returned, max to 100
func (s *GetUserSwapRecordsService) ResultSize(resultSize int64) *GetUserSwapRecordsService {
	s.resultSize = &resultSize
	return s
}

// Do sends the request.
func (s *GetUserSwapRecordsService) Do(ctx context.Context) ([]*SwapRecord, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}

	if s.swapId != nil {
		r.setParam("swapId", *s.swapId)
	}
	if s.quoteAsset != nil {
		r.setParam("quoteAsset", *s.quoteAsset)
	}
	if s.baseAsset != nil {
		r.setParam("baseAsset", *s.baseAsset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.resultSize != nil {
		r.setParam("limit", *s.resultSize)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := []*SwapRecord{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddLiquidityService to add liquidity
type AddLiquidityService struct {
	c             *Client
	poolId        *int64
	operationType *LiquidityOperationType
	quoteAsset    *string
	quoteQty      *float64
}

// PoolId set poolId
func (s *AddLiquidityService) PoolId(poolId int64) *AddLiquidityService {
	s.poolId = &poolId
	return s
}

// QuoteAsset set quoteAsset
func (s *AddLiquidityService) QuoteAsset(quoteAsset string) *AddLiquidityService {
	s.quoteAsset = &quoteAsset
	return s
}

// QuoteQty set quoteQty
func (s *AddLiquidityService) QuoteQty(quoteQty float64) *AddLiquidityService {
	s.quoteQty = &quoteQty
	return s
}

// OperationType set operationType
func (s *AddLiquidityService) OperationType(operationType LiquidityOperationType) *AddLiquidityService {
	s.operationType = &operationType
	return s
}

type AddLiquidityResponse struct {
	OperationId int64 `json:"operationId"`
}

// Do sends the request.
func (s *AddLiquidityService) Do(ctx context.Context) (*AddLiquidityResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/liquidityAdd",
		secType:  secTypeSigned,
	}

	r.setParam("poolId", *s.poolId)
	r.setParam("type", *s.operationType)
	r.setParam("asset", *s.quoteAsset)
	r.setParam("quantity", *s.quoteQty)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &AddLiquidityResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RemoveLiquidityService to remove liquidity
type RemoveLiquidityService struct {
	c             *Client
	poolId        *int64
	operationType *LiquidityOperationType
	assets        []string
	shareAmount   *float64
}

// PoolId set poolId
func (s *RemoveLiquidityService) PoolId(poolId int64) *RemoveLiquidityService {
	s.poolId = &poolId
	return s
}

// ShareAmount set shareAmount
func (s *RemoveLiquidityService) ShareAmount(amt float64) *RemoveLiquidityService {
	s.shareAmount = &amt
	return s
}

// QuoteQty set quoteQty
func (s *RemoveLiquidityService) AddAesst(asset string) *RemoveLiquidityService {
	s.assets = append(s.assets, asset)
	return s
}

// OperationType set operationType
func (s *RemoveLiquidityService) OperationType(operationType LiquidityOperationType) *RemoveLiquidityService {
	s.operationType = &operationType
	return s
}

type RemoveLiquidityResponse struct {
	OperationId int64 `json:"operationId"`
}

// Do sends the request.
func (s *RemoveLiquidityService) Do(ctx context.Context) (*RemoveLiquidityResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/liquidityRemove",
		secType:  secTypeSigned,
	}

	r.setParam("poolId", *s.poolId)
	r.setParam("type", *s.operationType)
	if len(s.assets) > 0 {
		r.setParam("asset", s.assets)
	}
	r.setParam("shareAmount", *s.shareAmount)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &RemoveLiquidityResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ClaimRewardService to claim reward
type ClaimRewardService struct {
	c          *Client
	rewardType *LiquidityRewardType
}

// RewardType set rewardType
func (s *ClaimRewardService) RewardType(t LiquidityRewardType) *ClaimRewardService {
	s.rewardType = &t
	return s
}

type ClaimRewardResponse struct {
	Success bool `json:"success"`
}

// Do sends the request.
func (s *ClaimRewardService) Do(ctx context.Context) (*ClaimRewardResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/claimRewards",
		secType:  secTypeSigned,
	}
	if s.rewardType != nil {
		r.setParam("type", *s.rewardType)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ClaimRewardResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryClaimedRewardHistoryService to query history of user claimed reward
type QueryClaimedRewardHistoryService struct {
	c            *Client
	poolId       *int64
	assetRewards *string
	rewardType   *LiquidityRewardType
	startTime    *int64
	endTime      *int64
	resultSize   *int64
}

// RewardType set rewardType
func (s *QueryClaimedRewardHistoryService) RewardType(t LiquidityRewardType) *QueryClaimedRewardHistoryService {
	s.rewardType = &t
	return s
}

// PoolId set pool id
func (s *QueryClaimedRewardHistoryService) PoolId(poolId int64) *QueryClaimedRewardHistoryService {
	s.poolId = &poolId
	return s
}

// AssetRewards set expected rewarded asset name
func (s *QueryClaimedRewardHistoryService) AssetRewards(assetRewards string) *QueryClaimedRewardHistoryService {
	s.assetRewards = &assetRewards
	return s
}

// StartTime set start time when reward
func (s *QueryClaimedRewardHistoryService) StartTime(startTime int64) *QueryClaimedRewardHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set end time when reward
func (s *QueryClaimedRewardHistoryService) EndTime(endTime int64) *QueryClaimedRewardHistoryService {
	s.endTime = &endTime
	return s
}

// ResultSize set the size will be returned, max to 100
func (s *QueryClaimedRewardHistoryService) ResultSize(resultSize int64) *QueryClaimedRewardHistoryService {
	s.resultSize = &resultSize
	return s
}

type ClaimedRewardHistory struct {
	PoolId        int               `json:"poolId"`
	PoolName      string            `json:"poolName"`
	AssetRewards  string            `json:"assetRewards"`
	ClaimedAt     int64             `json:"claimedTime"`
	ClaimedAmount string            `json:"claimAmount"`
	Status        RewardClaimStatus `json:"status"`
}

// Do sends the request.
func (s *QueryClaimedRewardHistoryService) Do(ctx context.Context) ([]*ClaimedRewardHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/claimedHistory",
		secType:  secTypeSigned,
	}
	if s.rewardType != nil {
		r.setParam("type", *s.rewardType)
	}
	if s.poolId != nil {
		r.setParam("poolId", *s.poolId)
	}
	if s.assetRewards != nil {
		r.setParam("assetRewards", *s.assetRewards)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.resultSize != nil {
		r.setParam("limit", *s.resultSize)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := []*ClaimedRewardHistory{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
