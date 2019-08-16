// Package kong ...
package kong

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// CreateConsumerResponse ...
type CreateConsumerResponse struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	CustomID  string `json:"custom_id,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

// CreateConsumer creates a consumer for the KONG API gateway.
func (c *client) CreateConsumer(usernameOrCustomID string) (*CreateConsumerResponse, error) {
	form := url.Values{}
	form.Add("username", usernameOrCustomID)
	form.Add("custom_id", usernameOrCustomID)

	rel, err := url.Parse("consumers")
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, reqErr := http.NewRequest("POST", u.String(), bytes.NewBufferString(form.Encode()))
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, doErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return nil, NewErrKongResponse(http.StatusCreated, resp.StatusCode, u.String())
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

// Delete Consumer requires usernameOrCustomID or ID to delete consumer via Kong API
// https://docs.konghq.com/0.14.x/admin-api/#delete-consumer
func (c *client) DeleteConsumer(usernameOrCustomID string) error {
	// Build URL
	rel, err := url.Parse(path.Join("consumers", usernameOrCustomID))
	if err != nil {
		return err
	}
	u := c.BaseURL.ResolveReference(rel)
	// Create Request
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Send Request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Check Response
	if resp.StatusCode != http.StatusNoContent {
		return NewErrKongResponse(http.StatusCreated, resp.StatusCode, u.String())
	}
	return nil
}
