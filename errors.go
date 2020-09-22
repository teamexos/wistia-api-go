package wistia

import (
	"fmt"
)

type (
	// ResponseError is used for errors this package will return
	ResponseError struct {
		StatusCode int
		Message    string
	}

	wistiaError struct {
		Error string `json:"error"`
	}
)

// Error satisfies the contract with the error interface
func (e *ResponseError) Error() string {
	return fmt.Sprintf("statusCode: %d, error: %s", e.StatusCode, e.Message)
}

// NewResponseError returns a RequestError
func NewResponseError(s int, m string) *ResponseError {
	return &ResponseError{
		StatusCode: s,
		Message:    m,
	}
}
