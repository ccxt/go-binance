package portfolio

import (
	"context"
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
		endpoint: "/papi/v1/ping",
	}
	_, _, err = s.c.callAPI(ctx, r, opts...)
	return err
}
