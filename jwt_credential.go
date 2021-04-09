package kong

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type JWTCredential struct {
	ID           string    `json:"id,omtiempty"`
	Consumer     *Consumer `json:"consumer,omitempty"`
	CreatedAt    int64     `json:"created_at,omitempty"`
	Key          string    `json:"key,omitemtpy"`
	Secret       string    `json:"secret,omptempty"`
	RSAPublicKey string    `json:"rsa_public_key,omitempty"`
	Algorithm    string    `json:"algorithm,omitempty"`
	Tags
}

// CreateJWTCredentialResponse ...
type CreateJWTCredentialResponse struct {
	*JWTCredential
}

// CreateJWTCredential ...
func (c *Client) CreateJWTCredential(usernameOrCustomID, key, secret string) (*CreateJWTCredentialResponse, error) {
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

	resp, doErr := c.doRequest(req)
	if doErr != nil {
		return nil, doErr
	}
	defer resp.Body.Close()

	expCodes := expectedCodes{http.StatusCreated}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}

	b, rErr := ioutil.ReadAll(resp.Body)
	if rErr != nil {
		return nil, rErr
	}

	response := &CreateJWTCredentialResponse{}
	mErr := json.Unmarshal(b, &response)
	if mErr != nil {
		return nil, mErr
	}
	return response, nil
}

// DeleteJWTCredential ...
func (c *Client) DeleteJWTCredential(usernameOrCustomID, jwtID string) error {
	rel, err := url.Parse(path.Join("consumers", usernameOrCustomID, "jwt", jwtID))
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, reqErr := http.NewRequest("DELETE", u.String(), nil)
	if reqErr != nil {
		return reqErr
	}

	resp, doErr := c.doRequest(req)
	if doErr != nil {
		return doErr
	}
	defer resp.Body.Close()

	expCodes := expectedCodes{http.StatusNoContent}
	if !expCodes.codeMatched(resp.StatusCode) {
		return NewErrKongResponse(expCodes, resp)
	}

	return nil
}
