package common

import (
	"fmt"
)

// APIError define API error when response status is 4xx or 5xx
type APIError struct {
	Code     int64  `json:"code"`
	Message  string `json:"msg"`
	Response []byte `json:"-"` // Assign the body value when the Code and Message fields are invalid.
}

// Error return error code and message
func (e APIError) Error() string {
	if e.IsValid() {
		return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("<APIError> rsp=%s", string(e.Response))
}

func (e APIError) IsValid() bool {
	return e.Code != 0 || e.Message != ""
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
