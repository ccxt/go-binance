package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMFeeBurnService service to toggle BNB burn on UM futures trade
type UMFeeBurnService struct {
	c          *Client
	feeBurn    bool
	recvWindow *int64
}

// FeeBurn set feeBurn status
func (s *UMFeeBurnService) FeeBurn(feeBurn bool) *UMFeeBurnService {
	s.feeBurn = feeBurn
	return s
}

// RecvWindow set recvWindow
func (s *UMFeeBurnService) RecvWindow(recvWindow int64) *UMFeeBurnService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMFeeBurnService) Do(ctx context.Context) (*UMFeeBurnResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/feeBurn",
		secType:  secTypeSigned,
	}

	r.setParam("feeBurn", s.feeBurn)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(UMFeeBurnResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMFeeBurnResponse define response for toggling BNB burn
type UMFeeBurnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
