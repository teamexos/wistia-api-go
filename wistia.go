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

// MediasShow returns information about a specific piece of media
func (c *Client) MediasShow(ctx context.Context, id string) (*Media, error) {
	endpoint := fmt.Sprintf("%s/medias/%s.json?access_token=%s", c.BaseURL, id, c.accessToken)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Media{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectsList returns a list of projects from Wistia
func (c *Client) ProjectsList(ctx context.Context, options *PaginationOptions) (*Projects, error) {
	paginationOpts := getPaginationOptions(options)
	endpoint := fmt.Sprintf("%s/projects.json?%s&access_token=%s",
		c.BaseURL,
		paginationOpts,
		c.accessToken)

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
func (c *Client) ProjectShow(ctx context.Context, id string) (*Project, error) {
	endpoint := fmt.Sprintf("%s/projects/%s.json?access_token=%s", c.BaseURL, id, c.accessToken)
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

func getPaginationOptions(options *PaginationOptions) string {
	page := 1
	perPage := 100
	sortBy := "name"
	sortDirection := 1
	if options != nil {
		page = options.Page
		perPage = options.PerPage
		sortBy = options.SortBy
		sortDirection = options.SortDirection
	}

	return fmt.Sprintf("page=%d&per_page=%d&sort_by=%s&sort_direction=%d",
		page,
		perPage,
		sortBy,
		sortDirection)
}
