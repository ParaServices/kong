package kong

import (
	"net/http"
	"net/url"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client ...
type Client struct {
	client httpClient
	// BaseURL ...
	BaseURL *url.URL
}
