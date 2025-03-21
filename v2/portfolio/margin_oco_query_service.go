package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// MarginOCOQueryService service to query margin OCO orders
type MarginOCOQueryService struct {
	c                 *Client
	orderListID       *int64
	origClientOrderID *string
}

// OrderListID set orderListID
func (s *MarginOCOQueryService) OrderListID(orderListID int64) *MarginOCOQueryService {
	s.orderListID = &orderListID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *MarginOCOQueryService) OrigClientOrderID(origClientOrderID string) *MarginOCOQueryService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *MarginOCOQueryService) Do(ctx context.Context, opts ...RequestOption) (res *MarginOCOResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/margin/orderList",
		secType:  secTypeSigned,
	}

	if s.orderListID != nil {
		r.setParam("orderListId", *s.orderListID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginOCOResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
