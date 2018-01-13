package kong

import (
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	CreateConsumer(string) (*CreateConsumerResponse, error)
	CreateJWTCredential(string, string, string) (*CreateJWTCredentialResponse, error)
	DeleteJWTCredential(string, string) error
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)
}

// Client ...
type client struct {
	client httpClient
	// BaseURL ...
	BaseURL *url.URL
}

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
