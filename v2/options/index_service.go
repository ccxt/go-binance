package options

import (
	"context"
	"encoding/json"
	"net/http"
)

type Index struct {
	Time       uint64 `json:"time"`
	IndexPrice string `json:"indexPrice"`
}

// underlying: Spot trading pairs such as BTCUSDT
type IndexService struct {
	c          *Client
	underlying string
}

// Underlying set underlying
func (s *IndexService) Underlying(underlying string) *IndexService {
	s.underlying = underlying
	return s
}

// Do send request
func (s *IndexService) Do(ctx context.Context, opts ...RequestOption) (res *Index, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/index",
	}
	r.setParam("underlying", s.underlying)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Index{}, err
	}
	res = new(Index)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &Index{}, err
	}
	return res, nil
}
