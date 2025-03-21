package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeUMPositionModeService change user's position mode on EVERY symbol in UM
type ChangeUMPositionModeService struct {
	c                *Client
	dualSidePosition bool
}

// DualSidePosition set position mode
func (s *ChangeUMPositionModeService) DualSidePosition(dualSidePosition bool) *ChangeUMPositionModeService {
	s.dualSidePosition = dualSidePosition
	return s
}

// Do send request
func (s *ChangeUMPositionModeService) Do(ctx context.Context, opts ...RequestOption) (res *APIResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setParam("dualSidePosition", s.dualSidePosition)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(APIResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// APIResponse define API response
type APIResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
