package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// CMADLQuantileService service to get CM position ADL quantile estimation
type CMADLQuantileService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *CMADLQuantileService) Symbol(symbol string) *CMADLQuantileService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *CMADLQuantileService) RecvWindow(recvWindow int64) *CMADLQuantileService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CMADLQuantileService) Do(ctx context.Context) ([]*CMADLQuantileResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/adlQuantile",
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
	var res []*CMADLQuantileResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMADLQuantileResponse define CM ADL quantile response
type CMADLQuantileResponse struct {
	Symbol      string      `json:"symbol"`
	ADLQuantile ADLQuantile `json:"adlQuantile"`
}
