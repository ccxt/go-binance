package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountToLotSize(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		minQty    string
		stepSize  string
		amount    string
		precision int
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "test with lot of zero and invalid amount",
			args: args{
				minQty:    "0.01",
				stepSize:  "0.01",
				amount:    "0.001",
				precision: 8,
			},
			expect: "0",
		},
		{
			name: "test with lot",
			args: args{
				minQty:    "0.01",
				stepSize:  "0.01",
				amount:    "1.39",
				precision: 8,
			},
			expect: "1.39",
		},
		{
			name: "test with exact precision",
			args: args{
				minQty:    "0.01",
				stepSize:  "0.01",
				amount:    "1.39",
				precision: 2,
			},
			expect: "1.39",
		},
		{
			name: "test with small precision",
			args: args{
				minQty:    "0.01",
				stepSize:  "0.01",
				amount:    "1.39",
				precision: 1,
			},
			expect: "1.3",
		},
		{
			name: "test with zero precision",
			args: args{
				minQty:    "0.01",
				stepSize:  "0.01",
				amount:    "1.39",
				precision: 0,
			},
			expect: "1",
		},
		{
			name: "test with big decimal",
			args: args{
				minQty:    "0.01",
				stepSize:  "0.02",
				amount:    "11.31232419283240912834434",
				precision: 8,
			},
			expect: "11.31",
		},
		{
			name: "test with big number",
			args: args{
				minQty:    "0.0001",
				stepSize:  "0.02",
				amount:    "11232821093480213.31232419283240912834434",
				precision: 8,
			},
			expect: "11232821093480213.3001",
		},
		{
			name: "test with small decimal",
			args: args{
				minQty:    "0.0000010",
				stepSize:  "0.0000010",
				amount:    "0.003923153000000002",
				precision: 7,
			},
			expect: "0.003923",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.expect, AmountToLotSize(tt.args.amount, tt.args.minQty, tt.args.stepSize, tt.args.precision))
		})
	}
}
