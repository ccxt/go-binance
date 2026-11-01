package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// AssetDividendService fetches the saving purchases
type AssetDividendService struct {
	c         *Client
	asset     *string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Asset sets the asset parameter.
func (s *AssetDividendService) Asset(asset string) *AssetDividendService {
	s.asset = &asset
	return s
}

// Limit sets the limit parameter.
func (s *AssetDividendService) Limit(limit int) *AssetDividendService {
	s.limit = &limit
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *AssetDividendService) StartTime(startTime int64) *AssetDividendService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *AssetDividendService) EndTime(endTime int64) *AssetDividendService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *AssetDividendService) Do(ctx context.Context) (*DividendResponseWrapper, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/assetDividend",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	} else {
		r.setParam("limit", 20)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(DividendResponseWrapper)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DividendResponseWrapper represents a wrapper around a AssetDividendService.
type DividendResponseWrapper struct {
	Rows  []*DividendResponse `json:"rows"`
	Total int32               `json:"total"`
}

// DividendResponse represents a response from AssetDividendService.
type DividendResponse struct {
	ID     int64  `json:"id"`
	Amount string `json:"amount"`
	Asset  string `json:"asset"`
	Info   string `json:"enInfo"`
	Time   int64  `json:"divTime"`
	TranID int64  `json:"tranId"`
}
