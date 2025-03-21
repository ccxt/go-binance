package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMFeeBurnStatusService service to get BNB burn status on UM futures trade
type UMFeeBurnStatusService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *UMFeeBurnStatusService) RecvWindow(recvWindow int64) *UMFeeBurnStatusService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMFeeBurnStatusService) Do(ctx context.Context) (*UMFeeBurnStatusResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/feeBurn",
		secType:  secTypeSigned,
	}

	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMFeeBurnStatusResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMFeeBurnStatusResponse define response for getting BNB burn status
type UMFeeBurnStatusResponse struct {
	FeeBurn bool `json:"feeBurn"`
}
