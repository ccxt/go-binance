package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMOrderHistoryDownloadIDService get download id for UM futures order history
type GetUMOrderHistoryDownloadIDService struct {
	c          *Client
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// StartTime set startTime
func (s *GetUMOrderHistoryDownloadIDService) StartTime(startTime int64) *GetUMOrderHistoryDownloadIDService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetUMOrderHistoryDownloadIDService) EndTime(endTime int64) *GetUMOrderHistoryDownloadIDService {
	s.endTime = &endTime
	return s
}

// RecvWindow set recvWindow
func (s *GetUMOrderHistoryDownloadIDService) RecvWindow(recvWindow int64) *GetUMOrderHistoryDownloadIDService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetUMOrderHistoryDownloadIDService) Do(ctx context.Context, opts ...RequestOption) (*UMOrderHistoryDownloadID, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order/asyn",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(UMOrderHistoryDownloadID)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMOrderHistoryDownloadID define download id response
type UMOrderHistoryDownloadID struct {
	AvgCostTimestampOfLast30d int64  `json:"avgCostTimestampOfLast30d"`
	DownloadID                string `json:"downloadId"`
}
