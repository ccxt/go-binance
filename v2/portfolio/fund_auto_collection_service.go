package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// FundAutoCollectionService fund auto-collection for Portfolio Margin
type FundAutoCollectionService struct {
	c *Client
}

// Do send request
func (s *FundAutoCollectionService) Do(ctx context.Context, opts ...RequestOption) (*SuccessResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/auto-collection",
		secType:  secTypeSigned,
	}

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
