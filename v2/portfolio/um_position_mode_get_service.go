package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMPositionModeService get user's position mode on EVERY symbol in UM
type GetUMPositionModeService struct {
	c *Client
}

// Do send request
func (s *GetUMPositionModeService) Do(ctx context.Context, opts ...RequestOption) (res *PositionMode, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/positionSide/dual",
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

// PositionMode define position mode info
type PositionMode struct {
	DualSidePosition bool `json:"dualSidePosition"` // true: Hedge Mode; false: One-way Mode
}
