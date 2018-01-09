package kong

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/ParaServices/errgo"
	"github.com/davecgh/go-spew/spew"
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
func (c *Client) CreateJWTCredential(consumerID, key, secret string) (*CreateJWTCredentialResponse, *errgo.Error) {
	form := url.Values{}
	form.Add("key", key)
	form.Add("secret", secret)

	c.BaseURL.Path = path.Join(c.BaseURL.Path, "consumers", consumerID, "jwt")
	spew.Dump(c.BaseURL.String())

	req, reqErr := http.NewRequest("POST", c.BaseURL.String(), strings.NewReader(form.Encode()))
	if reqErr != nil {
		return nil, errgo.New(reqErr)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, errgo.New(doErr)
	}

	if resp.StatusCode != 201 {
		return nil, errgo.New(fmt.Errorf("KONG returned a status not equal to 201, status: %s", resp.Status))
	}

	b, rErr := ioutil.ReadAll(resp.Body)
	if rErr != nil {
		return nil, errgo.New(rErr)
	}
	defer resp.Body.Close()
	response := &CreateJWTCredentialResponse{}
	mErr := json.Unmarshal(b, &response)
	if mErr != nil {
		return nil, errgo.New(mErr)
	}
	return response, nil
}

// DeleteJWTCredential ...
func (c *Client) DeleteJWTCredential(consumerID, jwtID string) *errgo.Error {
	c.BaseURL.Path = path.Join(c.BaseURL.Path, "consumers", consumerID, "jwt", jwtID)
	req, reqErr := http.NewRequest("DELETE", c.BaseURL.String(), nil)
	if reqErr != nil {
		return errgo.New(reqErr)
	}

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return errgo.New(doErr)
	}

	if resp.StatusCode != 204 {
		return errgo.New(fmt.Errorf("KONG returned a status not equal to 204, status: %s", resp.Status))
	}
	return nil
}
