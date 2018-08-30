package kong

import (
	"io"
	"net/http"
	"net/url"
)

// Client ...
type Client interface {
	CreateConsumer(string) (*CreateConsumerResponse, error)
	// DeleteConsumer Username or ID as string to delete consumer
	DeleteConsumer(string) error
	CreateJWTCredential(string, string, string) (*CreateJWTCredentialResponse, error)
	DeleteJWTCredential(string, string) error
	GetStatus() (*StatusResponse, error)
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)
}

type client struct {
	client httpClient
	// BaseURL ...
	BaseURL *url.URL
}

// NewClient ...
func NewClient(hc httpClient, baseURL *url.URL) Client {
	if hc == nil {
		hc = &http.Client{}
	}

	c := &client{
		client:  hc,
		BaseURL: baseURL,
	}

	return c
}
