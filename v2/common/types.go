package common

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

type UsedWeight struct {
	Used   int64
	Used1M int64 // used in last 1 minute
}

func (u *UsedWeight) UpdateByHeader(header http.Header) {
	usedWeight := header.Get("X-Mbx-Used-Weight")
	if usedWeight != "" {
		if used, err := strconv.ParseInt(usedWeight, 10, 64); err == nil {
			atomic.StoreInt64(&u.Used, used)
		}
	}
	usedWeight1M := header.Get("X-Mbx-Used-Weight-1m")
	if usedWeight1M != "" {
		if used, err := strconv.ParseInt(usedWeight1M, 10, 64); err == nil {
			atomic.StoreInt64(&u.Used1M, used)
		}
	}
}

type OrderCount struct {
	Count10s int64
	Count1d  int64
}

func (o *OrderCount) UpdateByHeader(header http.Header) {
	orderCount10s := header.Get("X-Mbx-Order-Count-10s")
	if orderCount10s != "" {
		if count, err := strconv.ParseInt(orderCount10s, 10, 64); err == nil {
			atomic.StoreInt64(&o.Count10s, count)
		}
	}

	orderCount1d := header.Get("X-Mbx-Order-Count-1d")
	if orderCount1d != "" {
		if count, err := strconv.ParseInt(orderCount1d, 10, 64); err == nil {
			atomic.StoreInt64(&o.Count1d, count)
		}
	}
}
