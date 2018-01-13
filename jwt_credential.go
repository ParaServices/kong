package kong

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// CreateJWTCredentialResponse ...
type CreateJWTCredentialResponse struct {
	ConsumerID string `json:"consumer_id"`
	CreatedAt  int64  `json:"created_at"`
	ID         string `json:"id"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
}

// CreateJWTCredential ...
func (c *client) CreateJWTCredential(consumerID, key, secret string) (*CreateJWTCredentialResponse, error) {
	form := url.Values{}
	form.Add("key", key)
	form.Add("secret", secret)

	c.BaseURL.Path = path.Join(c.BaseURL.Path, "consumers", consumerID, "jwt")

	req, reqErr := http.NewRequest("POST", c.BaseURL.String(), strings.NewReader(form.Encode()))
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, doErr
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("KONG returned a status not equal to 201, status: %s, url: %s", resp.Status, c.BaseURL.String())
	}

	b, rErr := ioutil.ReadAll(resp.Body)
	if rErr != nil {
		return nil, rErr
	}
	defer resp.Body.Close()
	response := &CreateJWTCredentialResponse{}
	mErr := json.Unmarshal(b, &response)
	if mErr != nil {
		return nil, mErr
	}
	return response, nil
}

// DeleteJWTCredential ...
func (c *client) DeleteJWTCredential(consumerID, jwtID string) error {
	c.BaseURL.Path = path.Join(c.BaseURL.Path, "consumers", consumerID, "jwt", jwtID)
	req, reqErr := http.NewRequest("DELETE", c.BaseURL.String(), nil)
	if reqErr != nil {
		return reqErr
	}

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return doErr
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("KONG returned a status not equal to 204, status: %s, url: %s", resp.Status, c.BaseURL.String())
	}
	return nil
}
