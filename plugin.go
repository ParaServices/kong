package kong

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// PluginConfig is an interface that needs to be implemented if configuration
// for the given plugin is needed.
type PluginConfig interface {
	Marshal() ([]byte, error)
}

// PluginService ...
type PluginService struct {
	ID string `json:"id"`
}

// PluginRoute ...
type PluginRoute struct {
	ID string `json:"id"`
}

// PluginConsumer ...
type PluginConsumer struct {
	ID string `json:"id,omitempty"`
}

// Plugin ...
type Plugin struct {
	Name      string          `json:"name,omitempty"`
	Service   *PluginService  `json:"service,omitempty"`
	Route     *PluginRoute    `json:"route,omitempty"`
	Consumer  *PluginConsumer `json:"consumer,omitempty"`
	Enabled   *bool           `json:"enabled,omitempty"`
	Config    PluginConfig    `json:"-"`
	configRaw *json.RawMessage
}

// UnmarshalJSON ...
func (p *Plugin) UnmarshalJSON(b []byte) error {
	type Alias Plugin

	aux := &struct {
		Config *json.RawMessage `json:"config,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.Config != nil {
		p.configRaw = aux.Config
		if p.Config != nil {
			if err := json.Unmarshal([]byte(*aux.Config), &p.Config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Plugin) MarshalJSON() ([]byte, error) {
	type Alias Plugin
	if p.Config != nil {
		b, err := p.Config.Marshal()
		if err != nil {
			return nil, err
		}
		jsonRawMessage := json.RawMessage(b)
		p.configRaw = &jsonRawMessage
	}
	return json.Marshal(&struct {
		Config *json.RawMessage `json:"config,omitempty"`
		*Alias
	}{
		Config: p.RawConfig(),
		Alias:  (*Alias)(p),
	})
}

// RawConfig returns the configRaw. This is useful when the p.Config is not
// set.
func (p *Plugin) RawConfig() *json.RawMessage {
	return p.configRaw
}

// EnablePlugin ...
func (c *Client) EnablePlugin(plugin *Plugin) (*Plugin, error) {
	rel, err := url.Parse("plugins")
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	type Alias struct {
		*Plugin
		Config *json.RawMessage `json:"config,omitempty"`
	}

	a := Alias{
		Plugin: plugin,
	}

	if plugin.Config != nil {
		configByte, err := plugin.Config.Marshal()
		if err != nil {
			return nil, err
		}
		jRawMessage := json.RawMessage(configByte)
		a.Config = &jRawMessage
	}

	b, err := json.Marshal(a)
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

	response := &Plugin{}
	if plugin.Config != nil {
		response.Config = plugin.Config
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Plugins ...
type Plugins []Plugin

// HasPlugin ...
func (p Plugins) HasPlugin(plugin string) bool {
	for i := range p {
		if plugin == p[i].Name {
			return true
		}
	}
	return false
}

// PluginList ..
type PluginList struct {
	Total int     `json:"total,omitempty"`
	Next  string  `json:"next,omitempty"`
	Data  Plugins `json:"data,omitempty"`
}

// ListPlugins ...
func (c *Client) ListPluginsForService(serviceID string) (*PluginList, error) {
	rel, err := url.Parse(path.Join("services", serviceID, "plugins"))
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
	response := &PluginList{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
