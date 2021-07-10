package kong

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/object"
	"github.com/ParaServices/paratils"
)

// AddService ...
func (c *Client) AddService(getter object.ServiceGetter) (*object.Service, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("service is nil")
	}
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
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("service is nil")
	}
	if paratils.StringIsEmpty(getter.GetID()) {
		return nil, errgo.NewF("service ID is empty")
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
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("service is nil")
	}
	if paratils.StringIsEmpty(getter.GetID()) {
		return nil, errgo.NewF("service ID is empty")
	}
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
