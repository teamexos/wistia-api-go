package wistia

import (
	"encoding/json"
	"fmt"
)

// RequestError is used for errors this package will return
type RequestError struct {
	StatusCode int
	Message    string
}

type wistiaError struct {
	Error string `json:"error"`
}

// Error satisfies the contract with the error interface
func (e *RequestError) Error() string {
	return fmt.Sprintf("statusCode: %d, error: %s", e.StatusCode, e.Message)
}

// NewError returns a RequestError
func NewError(s int, m *string) *RequestError {
	msg := wistiaError{}
	if err := json.Unmarshal([]byte(*m), &msg); err != nil {
		msg.Error = "Could not unmarshal response"
	}

	return &RequestError{
		StatusCode: s,
		Message:    msg.Error,
	}
}
