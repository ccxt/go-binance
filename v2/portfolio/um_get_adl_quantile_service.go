package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// UMGetADLQuantileService service to get UM position ADL quantile estimation
type UMGetADLQuantileService struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

// Symbol set symbol
func (s *UMGetADLQuantileService) Symbol(symbol string) *UMGetADLQuantileService {
	s.symbol = &symbol
	return s
}

// RecvWindow set recvWindow
func (s *UMGetADLQuantileService) RecvWindow(recvWindow int64) *UMGetADLQuantileService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *UMGetADLQuantileService) Do(ctx context.Context) ([]*UMADLQuantileResponse, error) {
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

// ADLQuantile define ADL quantile values
type ADLQuantile struct {
	LONG  int  `json:"LONG"`
	SHORT int  `json:"SHORT"`
	BOTH  *int `json:"BOTH,omitempty"`
	HEDGE *int `json:"HEDGE,omitempty"`
}

// UMADLQuantileResponse define UM ADL quantile response
type UMADLQuantileResponse struct {
	Symbol      string      `json:"symbol"`
	ADLQuantile ADLQuantile `json:"adlQuantile"`
}
