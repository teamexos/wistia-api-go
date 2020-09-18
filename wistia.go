package wistia

import (
	"context"
	"encoding/json"
	"fmt"
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

// ProjectsList returns a list of projects from Wistia
func (c *Client) ProjectsList(ctx context.Context) (*Projects, error) {
	endpoint := fmt.Sprintf("%s/projects.json?access_token=%s", c.BaseURL, c.accessToken)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Projects{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectShow returns a project from Wistia
func (c *Client) ProjectShow(ctx context.Context, id int) (*Project, error) {
	endpoint := fmt.Sprintf("%s/projects/%d.json?access_token=%s", c.BaseURL, id, c.accessToken)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Project{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
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
