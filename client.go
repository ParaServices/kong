package kong

import (
	"net/url"
)

// Client ...
type Client struct {
	client httpClient
	// BaseURL ...
	BaseURL *url.URL
}
