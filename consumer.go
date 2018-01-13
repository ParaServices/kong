// Package kong ...
package kong

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// CreateConsumerResponse ...
type CreateConsumerResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username,omitempty"`
	CustomID  string `json:"custom_id,omitempty"`
	CreatedAt int64  `json:"created_at"`
}

// CreateConsumer creates a consumer for the KONG API gateway.
func (c *client) CreateConsumer(username string) (*CreateConsumerResponse, error) {
	form := url.Values{}
	form.Add("username", username)

	c.BaseURL.Path = path.Join(c.BaseURL.Path, "consumers")

	req, reqErr := http.NewRequest("POST", c.BaseURL.String(), bytes.NewBufferString(form.Encode()))
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, doErr
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("KONG returned a status not equal to 201, status: %s, url: %s", resp.Status, c.BaseURL.String())
	}

	b, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	response := &CreateConsumerResponse{}
	unMarshalErr := json.Unmarshal(b, &response)
	if unMarshalErr != nil {
		return nil, unMarshalErr
	}

	return response, nil
}
