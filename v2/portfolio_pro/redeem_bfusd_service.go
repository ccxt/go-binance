package portfolio_pro

import (
	"context"
	"encoding/json"
	"net/http"
)

type RedeemBFUSDService struct {
	c *Client

	fromAsset   string
	targetAsset string
	amount      string
}

func (s *RedeemBFUSDService) Amount(amount string) *RedeemBFUSDService {
	s.amount = amount
	return s
}

func (s *RedeemBFUSDService) FromAsset(fromAsset string) *RedeemBFUSDService {
	s.fromAsset = fromAsset
	return s
}

func (s *RedeemBFUSDService) TargetAsset(targetAsset string) *RedeemBFUSDService {
	s.targetAsset = targetAsset
	return s
}

func (s *RedeemBFUSDService) Do(ctx context.Context, opts ...RequestOption) (*RedeemBFUSDResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/portfolio/redeem",
		secType:  secTypeSigned,
	}
	r.setParam("fromAsset", s.fromAsset)
	r.setParam("targetAsset", s.targetAsset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(RedeemBFUSDResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Example:
// {"fromAssetQty":"10","targetAssetQty":"9.9983733","redeemRate":"0.99983733","fromAsset":"BFUSD","targetAsset":"USDC"}
type RedeemBFUSDResponse struct {
	FromAsset      string `json:"fromAsset"`
	TargetAsset    string `json:"targetAsset"`
	FromAssetQty   string `json:"fromAssetQty"`
	TargetAssetQty string `json:"targetAssetQty"`
	RedeemRate     string `json:"redeemRate"`
}
