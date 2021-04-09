package kong

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// Route ...
type Route struct {
	Service       *Service `json:"service,omitempty"`
	ID            string   `json:"id,omitempty"`
	Protocols     []string `json:"protocols,omitempty"`
	Paths         []string `json:"paths,omitempty"`
	Methods       []string `json:"methods,omitempty"`
	Hosts         []string `json:"hosts,omitempty"`
	PreserveHost  bool     `json:"preserve_host,omitempty"`
	StripPath     bool     `json:"strip_path,omtiempty"`
	RegexPriority int      `json:"regex_priority,omitempty"`
	CreatedAt     int64    `json:"created_at,omitempty"`
	UpdatedAt     int64    `json:"updated_at,omitempty"`
}

// RoutesList ...
type RoutesList struct {
	Data []Route `json:"data,omitempty"`
}

// AddRoute ...
func (c *Client) AddRoute(route *Route) (*Route, error) {
	rel, err := url.Parse("routes")
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	b, err := json.Marshal(route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	expCodes := expectedCodes{http.StatusCreated}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Route{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListRoutesForService ...
func (c *Client) ListRoutesForService(serviceNameorID string) (*RoutesList, error) {
	rel, err := url.Parse(path.Join("services", serviceNameorID, "routes"))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	expCodes := expectedCodes{http.StatusOK}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &RoutesList{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteRoute ...
func (c *Client) DeleteRoute(routeID string) error {
	rel, err := url.Parse(path.Join("routes", routeID))
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}
	expCodes := expectedCodes{http.StatusNoContent}
	if !expCodes.codeMatched(resp.StatusCode) {
		return NewErrKongResponse(expCodes, resp)
	}
	return nil
}
