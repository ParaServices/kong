package kong

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// CreateJWTCredentialResponse ...
type CreateJWTCredentialResponse struct {
	Username  string `json:"username"`
	CustomID  string `json:"consumer_id"`
	CreatedAt int64  `json:"created_at"`
	ID        string `json:"id"`
	Key       string `json:"key"`
	Secret    string `json:"secret"`
}

// CreateJWTCredential ...
func (c *client) CreateJWTCredential(usernameOrCustomID, key, secret string) (*CreateJWTCredentialResponse, error) {
	form := url.Values{}
	form.Add("key", key)
	form.Add("secret", secret)

	rel, err := url.Parse(path.Join("consumers", usernameOrCustomID, "jwt"))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, reqErr := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, doErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return nil, NewErrKongResponse(http.StatusCreated, resp.StatusCode, u.String())
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
func (c *client) DeleteJWTCredential(usernameOrCustomID, jwtID string) error {
	rel, err := url.Parse(path.Join("consumers", usernameOrCustomID, "jwt", jwtID))
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, reqErr := http.NewRequest("DELETE", u.String(), nil)
	if reqErr != nil {
		return reqErr
	}

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return doErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return NewErrKongResponse(http.StatusNoContent, resp.StatusCode, u.String())
	}
	return nil
}
