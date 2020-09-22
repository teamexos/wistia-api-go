package wistia

import (
	"fmt"
	"net/http"
)

// RequestError is used for errors this package will return
type RequestError struct {
	StatusCode int
	Message    string
}

// Error satisfies the contract with the error interface
func (e *RequestError) Error() string {
	return fmt.Sprintf("statusCode: %d, error: %s", e.StatusCode, e.Message)
}

// NewError returns a RequestError
func NewError(s int) *RequestError {
	return &RequestError{
		StatusCode: s,
		Message:    http.StatusText(s),
	}
}
