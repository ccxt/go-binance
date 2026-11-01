package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginAvailableInventoryService fetches the margin interest history
type MarginAvailableInventoryService struct {
	c          *Client
	marginType *string
}

// MarginType sets the marginType parameter.
func (s *MarginAvailableInventoryService) MarginType(marginType string) *MarginAvailableInventoryService {
	s.marginType = &marginType
	return s
}

// Do sends the request.
func (s *MarginAvailableInventoryService) Do(ctx context.Context) (*marginAvailableInventory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/available-inventory",
		secType:  secTypeSigned,
	}
	if s.marginType != nil {
		r.setParam("type", *s.marginType)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(marginAvailableInventory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// marginAvailableInventory represents the response
type marginAvailableInventory struct {
	Assets     map[string]string `json:"assets"`
	UpdateTime int64             `json:"updateTime"`
}
