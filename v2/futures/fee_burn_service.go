package futures

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// FeeBurnService set fee burn info
type FeeBurnService struct {
	c       *Client
	feeBurn string
}

func (f *FeeBurnService) Enable() *FeeBurnService {
	f.feeBurn = "true"
	return f
}

func (f *FeeBurnService) Disable() *FeeBurnService {
	f.feeBurn = "false"
	return f
}

// Do send request
func (s *FeeBurnService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/feeBurn",
		secType:  secTypeSigned,
	}
	m := params{
		"feeBurn": s.feeBurn,
	}
	r.setFormParams(m)
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	j, err := newJSON(data)
	if err != nil {
		return err
	}
	msg, err := j.Get("msg").String()
	if err != nil {
		return err
	}
	if msg != "success" {
		code, _ := j.Get("code").Int()
		return fmt.Errorf("code: %d, msg: %s", code, msg)
	}
	return nil
}

// GetFeeBurnService get fee burn info
type GetFeeBurnService struct {
	c *Client
}

// Do send request
func (s *GetFeeBurnService) Do(ctx context.Context, opts ...RequestOption) (res *FeeBurn, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/feeBurn",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FeeBurn)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type FeeBurn struct {
	FeeBurn bool `json:"feeBurn"`
}
