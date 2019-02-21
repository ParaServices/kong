package kong

import (
	"net/http"
	"net/url"
	"time"
)

// Client ...
type Client interface {
	CreateConsumer(usernameOrCustomID string) (*CreateConsumerResponse, error)
	// DeleteConsumer Username or ID as string to delete consumer
	DeleteConsumer(string) error
	CreateJWTCredential(usernameOrCustomID, key, secret string) (*CreateJWTCredentialResponse, error)
	DeleteJWTCredential(usernameOrCustomID, jwtID string) error
	GetStatus() (*StatusResponse, error)
}

const (
	// DefaultMaxIdleConnections ...
	DefaultMaxIdleConnections = 10
	// DefaultRequestTimeout ...
	DefaultRequestTimeOut = 5
)

type client struct {
	client *http.Client
	// BaseURL ...
	BaseURL *url.URL
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
func NewClient(maxIdleConnections, requestTimeOut int, baseURL *url.URL) Client {
	if maxIdleConnections == 0 {
		maxIdleConnections = DefaultMaxIdleConnections
	}

	if requestTimeOut == 0 {
		requestTimeOut = DefaultRequestTimeOut
	}

	return &client{
		client:  createHTTPClient(maxIdleConnections, requestTimeOut),
		BaseURL: baseURL,
	}
}
