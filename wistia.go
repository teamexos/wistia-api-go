package wistia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// baseURLv1 is the base URL for v1 of the Wistia API
	baseURLv1 = "https://api.wistia.com/v1"
)

// Client is the base client for access Wistia APIs
type Client struct {
	// baseURL is the base URL for requests
	baseURL string

	// accessToken is the token used to authenticate requests
	accessToken string

	// httpClient is the client that makes the HTTP requests
	httpClient HTTPClient
}

// NewClient returns a pointer to Client
func NewClient(httpClient HTTPClient, accessToken string) *Client {
	return &Client{
		baseURL:     baseURLv1,
		accessToken: accessToken,
		httpClient:  httpClient,
	}
}

// MediasShow returns information about a specific piece of media
func (c *Client) MediasShow(ctx context.Context,
	id string,
	options *PaginationOptions) (*Media, *ResponseError) {

	opts := c.getOpts(options)
	endpoint := fmt.Sprintf("%s/medias/%s.json?%s",
		c.baseURL,
		id,
		opts)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, wistiaErrorRequestSetup
	}

	req = req.WithContext(ctx)

	res := Media{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectsList returns a list of projects from Wistia
func (c *Client) ProjectsList(ctx context.Context,
	options *PaginationOptions) (*Projects, *ResponseError) {

	opts := c.getOpts(options)
	endpoint := fmt.Sprintf("%s/projects.json?%s",
		c.baseURL,
		opts)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, wistiaErrorRequestSetup
	}

	req = req.WithContext(ctx)

	res := Projects{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectsShow returns a project from Wistia
func (c *Client) ProjectsShow(ctx context.Context,
	id string,
	options *PaginationOptions) (*Project, *ResponseError) {

	opts := c.getOpts(options)
	endpoint := fmt.Sprintf("%s/projects/%s.json?%s",
		c.baseURL,
		id,
		opts)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, wistiaErrorRequestSetup
	}

	req = req.WithContext(ctx)

	res := Project{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) getOpts(paginationOpts *PaginationOptions) string {
	page := 1
	perPage := 100
	sortBy := "name"
	sortDirection := 1
	if paginationOpts != nil {
		page = paginationOpts.Page
		perPage = paginationOpts.PerPage
		sortBy = paginationOpts.SortBy
		sortDirection = paginationOpts.SortDirection
	}

	return fmt.Sprintf("page=%d&per_page=%d&sort_by=%s&sort_direction=%d&access_token=%s",
		page,
		perPage,
		sortBy,
		sortDirection,
		c.accessToken)
}

func (c *Client) sendRequest(req *http.Request, v interface{}) *ResponseError {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return wistiaErrorRequestDo
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg := wistiaError{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)

		if err := json.Unmarshal([]byte(buf.String()), &msg); err != nil {
			msg.Error = "could not unmarshal response"
		}
		return NewResponseError(res.StatusCode, msg.Error)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return wistiaErrorRequestDecode
	}

	return nil
}
