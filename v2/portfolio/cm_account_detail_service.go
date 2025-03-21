package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetCMAccountDetailService get current CM account asset and position information
type GetCMAccountDetailService struct {
	c *Client
}

// Do send request
func (s *GetCMAccountDetailService) Do(ctx context.Context, opts ...RequestOption) (*CMAccountDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/account",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(CMAccountDetail)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CMAccountDetail define CM account detail
type CMAccountDetail struct {
	Assets    []CMAsset    `json:"assets"`
	Positions []CMPosition `json:"positions"`
}

// CMAsset define CM asset info
type CMAsset struct {
	Asset                  string `json:"asset"`                  // asset name
	CrossWalletBalance     string `json:"crossWalletBalance"`     // total wallet balance
	CrossUnPnl             string `json:"crossUnPnl"`             // unrealized profit or loss
	MaintMargin            string `json:"maintMargin"`            // maintenance margin
	InitialMargin          string `json:"initialMargin"`          // total initial margin required with the latest mark price
	PositionInitialMargin  string `json:"positionInitialMargin"`  // positions' margin required with the latest mark price
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // open orders' initial margin required with the latest mark price
	UpdateTime             int64  `json:"updateTime"`             // last update time
}
