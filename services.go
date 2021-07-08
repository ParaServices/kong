package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/object"
)

// AddService ...
func (c *Client) AddService(getter object.ServiceGetter) (*object.Service, error) {
	rel, err := url.Parse("services")
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	b, err := json.Marshal(getter)
	if err != nil {
		return nil, errgo.New(err)
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, errgo.New(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
	}

	expCodes := expectedCodes{http.StatusCreated}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errgo.New(err)
	}

	response := &object.Service{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	return response, nil
}

// UpdateService ...
func (c *Client) UpdateService(getter object.ServiceGetter) (*object.Service, error) {
	if getter.GetID() == "" {
		return nil, errors.New("Service ID must be defined when updating a service")
	}
	rel, err := url.Parse(path.Join("services", getter.GetID()))
	if err != nil {
		return nil, errgo.New(err)
	}
	u := c.baseURL.ResolveReference(rel)

	b, err := json.Marshal(getter)
	if err != nil {
		return nil, errgo.New(err)
	}

	req, err := http.NewRequest(http.MethodPatch, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, errgo.New(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
	}

	expCodes := expectedCodes{http.StatusOK}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errgo.New(err)
	}

	response := &object.Service{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	return response, nil
}

// ListServices ...
func (c *Client) ListServices() (object.Services, error) {
	rel, err := url.Parse("services")
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errgo.New(err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
	}
	expCodes := expectedCodes{http.StatusOK}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errgo.New(err)
	}
	response := &struct {
		Data object.Services `json:"data,omitempty"`
	}{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	if response.Data.GetLength() < 1 {
		return nil, nil
	}

	return response.Data, nil
}

// RetrieveService ...
func (c *Client) RetrieveService(getter object.KongIDGetter) (*object.Service, error) {
	rel, err := url.Parse(path.Join("services", getter.GetID()))
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errgo.New(err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
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
		return nil, errgo.New(err)
	}

	response := &object.Service{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	return response, nil
}
