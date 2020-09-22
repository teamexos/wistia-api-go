package wistia

import (
	"fmt"
)

type (
	// RequestError is used for errors this package will return
	RequestError struct {
		StatusCode int
		Message    string
	}

	wistiaError struct {
		Error string `json:"error"`
	}
)

// Error satisfies the contract with the error interface
func (e *RequestError) Error() string {
	return fmt.Sprintf("statusCode: %d, error: %s", e.StatusCode, e.Message)
}

// NewError returns a RequestError
func NewError(s int, m string) *RequestError {
	return &RequestError{
		StatusCode: s,
		Message:    m,
	}
}
