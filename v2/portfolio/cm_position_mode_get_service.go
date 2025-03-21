package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMPositionModeService get user's position mode on EVERY symbol in CM
type GetCMPositionModeService struct {
	c *Client
}

// Do send request
func (s *GetCMPositionModeService) Do(ctx context.Context, opts ...RequestOption) (res *PositionMode, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/positionSide/dual",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(PositionMode)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
