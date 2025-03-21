package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeAutoRepayFuturesStatusService change auto-repay-futures status
type ChangeAutoRepayFuturesStatusService struct {
	c         *Client
	autoRepay bool
}

// AutoRepay set auto repay status
func (s *ChangeAutoRepayFuturesStatusService) AutoRepay(autoRepay bool) *ChangeAutoRepayFuturesStatusService {
	s.autoRepay = autoRepay
	return s
}

// Do send request
func (s *ChangeAutoRepayFuturesStatusService) Do(ctx context.Context, opts ...RequestOption) (*SuccessResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/repay-futures-switch",
		secType:  secTypeSigned,
	}
	r.setParam("autoRepay", s.autoRepay)

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

// SuccessResponse define success response
type SuccessResponse struct {
	Msg string `json:"msg"`
}
