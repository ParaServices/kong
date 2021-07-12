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

// EnablePlugin ...
func (c *Client) EnablePlugin(getter object.PluginGetter) (*object.Plugin, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("plugin is nil")
	}
	rel, err := url.Parse("plugins")
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

	response := &object.Plugin{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	return response, nil
}

// PluginList ..
type PluginList struct {
	Total int            `json:"total,omitempty"`
	Next  string         `json:"next,omitempty"`
	Data  object.Plugins `json:"data,omitempty"`
}

// ListPlugins ...
func (c *Client) ListPluginsForService(serviceID string) (*PluginList, error) {
	rel, err := url.Parse(path.Join("services", serviceID, "plugins"))
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errgo.New(err)
	}
	response := &PluginList{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errgo.New(err)
	}

	return response, nil
}
