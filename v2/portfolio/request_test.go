package portfolio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithExtraForm(t *testing.T) {
	tests := []struct {
		name        string
		r           *request
		m           map[string]any
		wantRequest *request
	}{
		{
			name: "place order use extra priceMatch and goodTillDate",
			r: &request{
				form: map[string][]string{
					"symbol":  {"BTCUSDT"},
					"orderId": {"1"},
				},
			},
			m: map[string]any{
				"priceMatch":   "QUEUE",
				"goodTillDate": 1697796587000,
			},
			wantRequest: &request{
				form: map[string][]string{
					"symbol":       {"BTCUSDT"},
					"orderId":      {"1"},
					"priceMatch":   {"QUEUE"},
					"goodTillDate": {"1697796587000"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			opt := WithExtraForm(tt.m)
			opt(tt.r)

			assert.Equal(t, tt.wantRequest, tt.r)
		})
	}
}
