package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// FundCollectionByAssetService transfers specific asset from Futures Account to Margin account
type FundCollectionByAssetService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *FundCollectionByAssetService) Asset(asset string) *FundCollectionByAssetService {
	s.asset = asset
	return s
}

// Do send request
func (s *FundCollectionByAssetService) Do(ctx context.Context, opts ...RequestOption) (*SuccessResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/asset-collection",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(SuccessResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
