package portfolio_pro

import (
	"context"
	"encoding/json"
	"net/http"
)

type RedeemBFUSDService struct {
	c *Client

	amount      string
	targetAsset string
}

func (s *RedeemBFUSDService) Amount(amount string) *RedeemBFUSDService {
	s.amount = amount
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
	r.setParam("fromAsset", "BFUSD")
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

type RedeemBFUSDResponse struct {
	BFUSDResponse
	RedeemRate string `json:"redeemRate"`
}
