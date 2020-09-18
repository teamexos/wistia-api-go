package wistia

import (
	"net/http"
	"time"
)

const (
	// BaseURLv1 is the base URL for v1 of the Wistia API
	BaseURLv1 = "https://api.wistia.com/v1"
)

// Client is the base client for access Wistia APIs
type Client struct {
	// BaseURL is the base URL for requests
	BaseURL string

	// accessToken is the token used to authenticate requests
	accessToken string

	// HTTPClient is the client that makes the HTTP requests
	HTTPClient *http.Client
}

// NewClient returns a pointer to Client
func NewClient(accessToken string) *Client {
	return &Client{
		BaseURL:     BaseURLv1,
		accessToken: accessToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
