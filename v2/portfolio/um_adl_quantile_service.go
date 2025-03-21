package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMADLQuantileService service to get UM position ADL quantile estimation
type UMADLQuantileService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *UMADLQuantileService) Symbol(symbol string) *UMADLQuantileService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *UMADLQuantileService) RecvWindow(recvWindow int64) *UMADLQuantileService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMADLQuantileService) Do(ctx context.Context) ([]*UMADLQuantileResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/adlQuantile",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}

	data, _, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var res []*UMADLQuantileResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
