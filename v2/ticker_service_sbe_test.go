package binance

import (
	"bytes"
	"math"
	"testing"

	sbe "github.com/adshao/go-binance/v2/sbe/spot_3_1"
)

func TestConvertSBEPrice(t *testing.T) {
	tests := []struct {
		name     string
		mantissa int64
		exponent int8
		expected string
	}{
		{
			name:     "Positive with negative exponent",
			mantissa: 123456,
			exponent: -2,
			expected: "1234.56",
		},
		{
			name:     "Positive with zero exponent",
			mantissa: 12345,
			exponent: 0,
			expected: "12345",
		},
		{
			name:     "Null value",
			mantissa: math.MinInt64,
			exponent: -2,
			expected: "",
		},
		{
			name:     "Large number",
			mantissa: 9876543210,
			exponent: -8,
			expected: "98.76543210",
		},
		{
			name:     "Small number",
			mantissa: 1,
			exponent: -8,
			expected: "0.00000001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertSBEPrice(tt.mantissa, tt.exponent)
			if result != tt.expected {
				t.Errorf("convertSBEPrice(%d, %d) = %s; want %s",
					tt.mantissa, tt.exponent, result, tt.expected)
			}
		})
	}
}

func TestConvertSBEString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "String with null terminator",
			input:    []byte{'B', 'T', 'C', 0, 0, 0, 0, 0},
			expected: "BTC",
		},
		{
			name:     "Full string without null",
			input:    []byte{'E', 'T', 'H', 'U', 'S', 'D', 'T'},
			expected: "ETHUSDT",
		},
		{
			name:     "Empty string",
			input:    []byte{0, 0, 0, 0},
			expected: "",
		},
		{
			name:     "String at start",
			input:    []byte{'A', 0, 0, 0, 0, 0, 0, 0},
			expected: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertSBEString(tt.input)
			if result != tt.expected {
				t.Errorf("convertSBEString(%v) = %s; want %s",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestSBEMessageHeaderDecoding(t *testing.T) {
	// Create a mock SBE message header
	marshaller := sbe.NewSbeGoMarshaller()
	buf := new(bytes.Buffer)

	// Encode a header
	header := sbe.MessageHeader{
		BlockLength: 34,
		TemplateId:  211, // BookTickerSymbolResponse
		SchemaId:    3,
		Version:     1,
	}

	err := header.Encode(marshaller, buf)
	if err != nil {
		t.Fatalf("Failed to encode header: %v", err)
	}

	// Decode the header
	var decodedHeader sbe.MessageHeader
	reader := bytes.NewReader(buf.Bytes())
	err = decodedHeader.Decode(marshaller, reader, 0)
	if err != nil {
		t.Fatalf("Failed to decode header: %v", err)
	}

	// Verify decoded values
	if decodedHeader.BlockLength != header.BlockLength {
		t.Errorf("BlockLength = %d; want %d", decodedHeader.BlockLength, header.BlockLength)
	}
	if decodedHeader.TemplateId != header.TemplateId {
		t.Errorf("TemplateId = %d; want %d", decodedHeader.TemplateId, header.TemplateId)
	}
	if decodedHeader.SchemaId != header.SchemaId {
		t.Errorf("SchemaId = %d; want %d", decodedHeader.SchemaId, header.SchemaId)
	}
	if decodedHeader.Version != header.Version {
		t.Errorf("Version = %d; want %d", decodedHeader.Version, header.Version)
	}
}

func TestWithSBERequestOption(t *testing.T) {
	r := &request{}

	// Apply the SBE option
	opt := WithSBE(3, 1)
	opt(r)

	// Verify headers are set correctly
	if r.header == nil {
		t.Fatal("Expected headers to be set")
	}

	acceptHeader := r.header.Get("Accept")
	if acceptHeader != "application/sbe" {
		t.Errorf("Accept header = %s; want application/sbe", acceptHeader)
	}

	sbeHeader := r.header.Get("X-MBX-SBE")
	if sbeHeader != "3:1" {
		t.Errorf("X-MBX-SBE header = %s; want 3:1", sbeHeader)
	}
}

// Test the centralized SBE decoder
func TestSBEDecoder(t *testing.T) {
	decoder := NewSBEDecoder()
	if decoder == nil {
		t.Fatal("NewSBEDecoder() returned nil")
	}
	if decoder.marshaller == nil {
		t.Fatal("SBEDecoder marshaller is nil")
	}
}
