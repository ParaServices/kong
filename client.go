package kong

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/paratils"
)

const (
	libraryVersion            = "0.4.0"
	defaultMaxIdleConnections = 10
	defaultRequestTimeOut     = time.Second * 10
)

type ClientFuncSetter func(c *Client) error

func defaultHTTPClient(c *Client) error {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: defaultMaxIdleConnections,
		},
		Timeout: defaultRequestTimeOut,
	}
	c.client = client
	return nil
}

func userAgent() string {
	var b strings.Builder
	defer b.Reset()

	fmt.Fprintf(&b, "kong-client/%s ", libraryVersion)
	fmt.Fprintf(&b, "(%s; %s) ", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(&b, "%s", runtime.Version())

	s := b.String()
	return s
}

func setBaseURL(u *url.URL) ClientFuncSetter {
	return func(c *Client) error {
		return c.SetBaseURL(u)
	}
}

func defaultUserAgent(c *Client) error {
	if c.userAgent == "" {
		c.userAgent = userAgent()
	}
	return nil
}

type Client struct {
	client    *http.Client
	baseURL   *url.URL
	userAgent string
}

func (c *Client) SetClient(client *http.Client) error {
	c.client = client
	return nil
}

func (c *Client) SetBaseURL(u *url.URL) error {
	c.baseURL = u
	return nil
}

func (c *Client) SetUserAgent(userAgent string) error {
	c.userAgent = userAgent
	return nil
}

// this must be used in favor of client.Do because this will append headers
// internal to thsi package
func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", c.userAgent)
	return c.client.Do(req)
}

// NewClient ...
func NewClient(baseURL *url.URL, setterFuncs ...ClientFuncSetter) (*Client, error) {
	if paratils.IsNil(baseURL) {
		return nil, errgo.NewF("base url can not be nil")
	}

	c := &Client{}

	defSetters := []ClientFuncSetter{
		setBaseURL(baseURL),
		defaultHTTPClient,
		defaultUserAgent,
	}

	for i := range defSetters {
		if err := defSetters[i](c); err != nil {
			return nil, errgo.New(err)
		}
	}

	for i := range setterFuncs {
		if err := setterFuncs[i](c); err != nil {
			return nil, errgo.New(err)
		}
	}

	return c, nil
}
