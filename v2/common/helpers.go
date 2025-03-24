package common

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AmountToLotSize convert amount to lot size
func AmountToLotSize(amount, minQty, stepSize string, precision int) string {
	amountDec := decimal.RequireFromString(amount)
	minQtyDec := decimal.RequireFromString(minQty)
	baseAmountDec := amountDec.Sub(minQtyDec)
	if baseAmountDec.LessThan(decimal.Zero) {
		return "0"
	}
	stepSizeDec := decimal.RequireFromString(stepSize)
	baseAmountDec = baseAmountDec.Div(stepSizeDec).Truncate(0).Mul(stepSizeDec)
	return baseAmountDec.Add(minQtyDec).Truncate(int32(precision)).String()
}

// ToJSONList convert v to json list if v is a map
func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}

func ToInt(digit interface{}) (i int, err error) {
	if intVal, ok := digit.(int); ok {
		return int(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}

func ToInt64(digit interface{}) (i int64, err error) {
	if intVal, ok := digit.(int); ok {
		return int64(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int64(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}

const (
	SPOT_ORDER_PREFIX     = "x-HNA2TXFJ"
	CONTRACT_ORDER_PREFIX = "x-Cb7ytekJ"
)

func BaseUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func Uuid22() string {
	return BaseUID()[:22]
}

func GenerateSpotId() string {
	return SPOT_ORDER_PREFIX + Uuid22()
}

func GenerateSwapId() string {
	return CONTRACT_ORDER_PREFIX + Uuid22()
}
