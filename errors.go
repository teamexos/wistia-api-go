package wistia

import (
	"fmt"
	"net/http"
)

const (
	errCodeRequestDecodeFailed = "failed to decode response"
	errCodeRequestDoFailed     = "failed to make request"
	errCodeRequestSetupFailed  = "failed to setup request"
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

var (
	wistiaErrorRequestDecode = NewResponseError(http.StatusInternalServerError, errCodeRequestDecodeFailed)
	wistiaErrorRequestDo     = NewResponseError(http.StatusInternalServerError, errCodeRequestDoFailed)
	wistiaErrorRequestSetup  = NewResponseError(http.StatusInternalServerError, errCodeRequestSetupFailed)
)

// Error satisfies the contract with the error interface
func (e *ResponseError) Error() string {
	return fmt.Sprintf("statusCode: %d, error: %s", e.StatusCode, e.Message)
}

// NewResponseError returns a ResponseError
func NewResponseError(s int, m string) *ResponseError {
	return &ResponseError{
		StatusCode: s,
		Message:    m,
	}
}
