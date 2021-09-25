package binance

import (
	"context"
	"encoding/json"
)

// ListSavingsFlexibleProductsService get
type ListSavingsFlexibleProductsService struct {
	c        *Client
	status   string
	featured string
	current  int64
	size     int64
}

func (s *ListSavingsFlexibleProductsService) Status(status string) *ListSavingsFlexibleProductsService {
	s.status = status
	return s
}

func (s *ListSavingsFlexibleProductsService) Featured(featured string) *ListSavingsFlexibleProductsService {
	s.featured = featured
	return s
}

func (s *ListSavingsFlexibleProductsService) Current(current int64) *ListSavingsFlexibleProductsService {
	s.current = current
	return s
}

func (s *ListSavingsFlexibleProductsService) Size(size int64) *ListSavingsFlexibleProductsService {
	s.size = size
	return s
}

// Do send request
func (s *ListSavingsFlexibleProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*SavingsFlexibleProduct, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/lending/daily/product/list",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.status != "" {
		m["status"] = s.status
	}
	if s.featured != "" {
		m["featured"] = s.featured
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*SavingsFlexibleProduct
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SavingsFlexibleProduct define a flexible product (Savings)
type SavingsFlexibleProduct struct {
	Asset                    string `json:"asset"`
	AvgAnnualInterestRate    string `json:"avgAnnualInterestRate"`
	CanPurchase              bool   `json:"canPurchase"`
	CanRedeem                bool   `json:"canRedeem"`
	DailyInterestPerThousand string `json:"dailyInterestPerThousand"`
	Featured                 bool   `json:"featured"`
	MinPurchaseAmount        string `json:"minPurchaseAmount"`
	ProductId                string `json:"productId"`
	PurchasedAmount          string `json:"purchasedAmount"`
	Status                   string `json:"status"`
	UpLimit                  string `json:"upLimit"`
	UpLimitPerUser           string `json:"upLimitPerUser"`
}

// ListSavingsFixedAndActivityProductsService https://binance-docs.github.io/apidocs/spot/en/#get-fixed-and-activity-project-list-user_data
type ListSavingsFixedAndActivityProductsService struct {
	c           *Client
	asset       string
	projectType string
	status      string
	isSortAsc   bool
	sortBy      string
	current     int64
	size        int64
}

func (s *ListSavingsFixedAndActivityProductsService) Asset(asset string) *ListSavingsFixedAndActivityProductsService {
	s.asset = asset
	return s
}

func (s *ListSavingsFixedAndActivityProductsService) IsSortAsc(inSortAsc bool) *ListSavingsFixedAndActivityProductsService {
	s.isSortAsc = inSortAsc
	return s
}

func (s *ListSavingsFixedAndActivityProductsService) SortBy(sortBy string) *ListSavingsFixedAndActivityProductsService {
	s.sortBy = sortBy
	return s
}

func (s *ListSavingsFixedAndActivityProductsService) Current(current int64) *ListSavingsFixedAndActivityProductsService {
	s.current = current
	return s
}

func (s *ListSavingsFixedAndActivityProductsService) Size(size int64) *ListSavingsFixedAndActivityProductsService {
	s.size = size
	return s
}

func (s *ListSavingsFixedAndActivityProductsService) Status(status string) *ListSavingsFixedAndActivityProductsService {
	s.status = status
	return s
}

// Type set project type ("ACTIVITY", "CUSTOMIZED_FIXED")
func (s *ListSavingsFixedAndActivityProductsService) Type(projectType string) *ListSavingsFixedAndActivityProductsService {
	s.projectType = projectType
	return s
}

// Do send request
func (s *ListSavingsFixedAndActivityProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*SavingsFixedProduct, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/lending/project/list",
		secType:  secTypeSigned,
	}
	m := params{
		"type": s.projectType,
	}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.status != "" {
		m["status"] = s.status
	}
	if s.isSortAsc != true {
		m["isSortAsc"] = s.isSortAsc
	}
	if s.sortBy != "" {
		m["sortBy"] = s.sortBy
	}
	if s.current != 1 {
		m["current"] = s.current
	}
	if s.size != 10 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*SavingsFixedProduct
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SavingsFixedProduct define a fixed product (Savings)
type SavingsFixedProduct struct {
	Asset              string `json:"asset"`
	DisplayPriority    int    `json:"displayPriority"`
	Duration           int    `json:"duration"`
	InterestPerLot     string `json:"interestPerLot"`
	InterestRate       string `json:"interestRate"`
	LotSize            string `json:"lotSize"`
	LotsLowLimit       int    `json:"lotsLowLimit"`
	LotsPurchased      int    `json:"lotsPurchased"`
	LotsUpLimit        int    `json:"lotsUpLimit"`
	MaxLotsPerUser     int    `json:"maxLotsPerUser"`
	NeedKyc            bool   `json:"needKyc"`
	ProjectId          string `json:"projectId"`
	ProjectName        string `json:"projectName"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	WithAreaLimitation bool   `json:"withAreaLimitation"`
}
