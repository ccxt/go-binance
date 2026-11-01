package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeCMPositionModeService change user's position mode on EVERY symbol in CM
type ChangeCMPositionModeService struct {
	c                *Client
	dualSidePosition bool
}

// DualSidePosition set position mode
func (s *ChangeCMPositionModeService) DualSidePosition(dualSidePosition bool) *ChangeCMPositionModeService {
	s.dualSidePosition = dualSidePosition
	return s
}

// Do send request
func (s *ChangeCMPositionModeService) Do(ctx context.Context, opts ...RequestOption) (res *APIResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/cm/positionSide/dual",
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
