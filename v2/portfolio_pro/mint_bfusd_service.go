package portfolio_pro

import (
	"context"
	"encoding/json"
	"net/http"
)

type MintBFUSDService struct {
	c *Client

	fromAsset   string
	targetAsset string
	amount      string
}

func (s *MintBFUSDService) FromAsset(fromAsset string) *MintBFUSDService {
	s.fromAsset = fromAsset
	return s
}

func (s *MintBFUSDService) Amount(amount string) *MintBFUSDService {
	s.amount = amount
	return s
}

func (s *MintBFUSDService) TargetAsset(targetAsset string) *MintBFUSDService {
	s.targetAsset = targetAsset
	return s
}

func (s *MintBFUSDService) Do(ctx context.Context, opts ...RequestOption) (*MintBFUSDResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/portfolio/mint",
		secType:  secTypeSigned,
	}
	r.setParam("fromAsset", s.fromAsset)
	r.setParam("targetAsset", s.targetAsset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(MintBFUSDResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Example:
// {"fromAssetQty":"10","targetAssetQty":"9.9966156","mintRate":"0.99966156","fromAsset":"USDC","targetAsset":"BFUSD"}
type MintBFUSDResponse struct {
	FromAsset      string `json:"fromAsset"`
	TargetAsset    string `json:"targetAsset"`
	FromAssetQty   string `json:"fromAssetQty"`
	TargetAssetQty string `json:"targetAssetQty"`
	MintRate       string `json:"mintRate"`
}
