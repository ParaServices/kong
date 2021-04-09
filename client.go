package kong

import (
	"net/http"
	"net/url"
	"time"

	"github.com/ParaServices/kong/version"
)

const (
	// DefaultMaxIdleConnections ...
	DefaultMaxIdleConnections = 10
	// DefaultRequestTimeout ...
	DefaultRequestTimeOut = 5
)

type Client struct {
	client *http.Client
	// BaseURL ...
	BaseURL *url.URL
}

// this must be used in favor of client.Do because this will append headers
// internal to thsi package
func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", version.UserAgent())
	return c.client.Do(req)
}

// createHTTPClient for connection re-use
func createHTTPClient(maxIdleConnections, requestTimeOut int) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeOut) * time.Second,
	}

	return client
}

// NewClient ...
func NewClient(maxIdleConnections, requestTimeOut int, baseURL *url.URL) *Client {
	if maxIdleConnections == 0 {
		maxIdleConnections = DefaultMaxIdleConnections
	}

	if requestTimeOut == 0 {
		requestTimeOut = DefaultRequestTimeOut
	}

	return &Client{
		client:  createHTTPClient(maxIdleConnections, requestTimeOut),
		BaseURL: baseURL,
	}
}
