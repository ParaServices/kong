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
)

// AddRoute ...
func (c *Client) AddRoute(getter object.RouteGetter) (*object.Route, error) {
	rel, err := url.Parse("routes")
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

	response := &object.Route{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	return response, nil
}

// ListRoutesForService ...
func (c *Client) ListRoutesForService(getter object.KongIDGetter) (object.Routes, error) {
	rel, err := url.Parse(path.Join("services", getter.GetID(), "routes"))
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
		Data object.Routes `json:"data,omitempty"`
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

// DeleteRoute ...
func (c *Client) DeleteRoute(getter object.KongIDGetter) error {
	rel, err := url.Parse(path.Join("routes", getter.GetID()))
	if err != nil {
		return err
	}

	u := c.baseURL.ResolveReference(rel)

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
