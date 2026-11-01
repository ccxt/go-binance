package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type SpotRebateHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
	page      *int32
}

// BeginTime set beginTime
func (s *SpotRebateHistoryService) StartTime(startTime int64) *SpotRebateHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *SpotRebateHistoryService) EndTime(endTime int64) *SpotRebateHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *SpotRebateHistoryService) Page(page int32) *SpotRebateHistoryService {
	s.page = &page
	return s
}

// Do send request
func (s *SpotRebateHistoryService) Do(ctx context.Context, opts ...RequestOption) (*SpotRebateHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/rebate/taxQuery",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := SpotRebateHistory{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SpotRebateHistory define the spot rebate history
type SpotRebateHistory struct {
	Status string                `json:"status"`
	Type   string                `json:"type"`
	Code   string                `json:"code"`
	Data   SpotRebateHistoryData `json:"data"`
}

// SpotRebateHistoryData define the data part of the spot rebate history
type SpotRebateHistoryData struct {
	Page         int32                       `json:"page"`
	TotalRecords int32                       `json:"totalRecords"`
	TotalPageNum int32                       `json:"totalPageNum"`
	Data         []SpotRebateHistoryDataItem `json:"data"`
}

// SpotRebateHistoryDataItem define a spot rebate history data item
type SpotRebateHistoryDataItem struct {
	Asset      string `json:"asset"`
	Type       int32  `json:"type"`
	Amount     string `json:"amount"`
	UpdateTime int64  `json:"updateTime"`
}
