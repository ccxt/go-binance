package options

import (
	"context"
	"encoding/json"
	"net/http"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/ping",
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	d := map[string]string{}
	err = json.Unmarshal(data, &d)
	return err
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/eapi/v1/time",
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}
	j, err := newJSON(data)
	if err != nil {
		return 0, err
	}
	serverTime = j.Get("serverTime").MustInt64()
	return serverTime, nil
}
