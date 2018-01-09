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

	"github.com/ParaServices/errgo"
)

// CreateConsumerResponse ...
type CreateConsumerResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username,omitempty"`
	CustomID  string `json:"custom_id,omitempty"`
	CreatedAt int64  `json:"created_at"`
}

// CreateConsumer creates a consumer for the KONG API gateway.
func (c *Client) CreateConsumer(username string) (*CreateConsumerResponse, *errgo.Error) {
	form := url.Values{}
	form.Add("username", username)

	c.BaseURL.Path = path.Join(c.BaseURL.Path, "consumers")

	req, reqErr := http.NewRequest("POST", c.BaseURL.String(), bytes.NewBufferString(form.Encode()))
	if reqErr != nil {
		return nil, errgo.New(reqErr)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, errgo.New(doErr)
	}

	if resp.StatusCode != 201 {
		return nil, errgo.New(fmt.Errorf("KONG returned a status not equal to 201, status: %s", resp.Status))
	}

	b, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, errgo.New(readErr)
	}

	response := &CreateConsumerResponse{}
	unMarshalErr := json.Unmarshal(b, &response)
	if unMarshalErr != nil {
		return nil, errgo.New(unMarshalErr)
	}

	return response, nil
}
