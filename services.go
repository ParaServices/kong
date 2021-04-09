package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// Service ...
type Service struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Host           string `json:"host,omitempty"`
	Path           string `json:"path,omitempty"`
	Port           int    `json:"port,omitempty"`
	Retries        int    `json:"retries,omitempty"`
	ConnectTimeout int    `json:"connect_timeout,omitempty"`
	WriteTimeout   int    `json:"write_timeout,omitempty"`
	ReadTimeout    int    `json:"read_timeout,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
}

// ServicesList ...
type ServicesList struct {
	Data []Service `json:"data"`
}

// AddService ...
func (c *Client) AddService(service *Service) (*Service, error) {
	rel, err := url.Parse("services")
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	b, err := json.Marshal(service)
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

	response := &Service{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateService ...
func (c *Client) UpdateService(service *Service) (*Service, error) {
	if service.ID == "" {
		return nil, errors.New("Service ID must be defined when updating a service")
	}
	rel, err := url.Parse(path.Join("services", service.ID))
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	// remove service ID from the request body
	service.ID = ""
	b, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	expCodes := expectedCodes{http.StatusOK}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Service{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListServices ...
func (c *Client) ListServices() (*ServicesList, error) {
	rel, err := url.Parse("services")
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
	response := &ServicesList{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveService ...
func (c *Client) RetrieveService(nameOrID string) (*Service, error) {
	rel, err := url.Parse(path.Join("services", nameOrID))
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
	expCodes := expectedCodes{http.StatusOK, http.StatusNotFound}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Service{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
