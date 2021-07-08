package kong

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/plugin"
	"github.com/ParaServices/paratils"
)

// CreateJWTCredential ...
func (c *Client) CreateJWTCredential(getter plugin.JWTCredentialGetter) (*plugin.JWTCredential, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("jwt credential is nil")
	}
	form := url.Values{}
	form.Add("key", getter.GetKey())
	form.Add("secret", getter.GetSecret())

	rel, err := url.Parse(path.Join("consumers", getter.GetConsumer().GetID(), "jwt"))
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, errgo.New(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
	}

	expCodes := expectedCodes{http.StatusCreated}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errgo.New(err)
	}

	jwtCred := &plugin.JWTCredential{}
	err = json.Unmarshal(b, &jwtCred)
	if err != nil {
		return nil, errgo.New(err)
	}
	return jwtCred, nil
}

// DeleteJWTCredential ...
func (c *Client) DeleteJWTCredential(getter plugin.JWTCredentialGetter) error {
	rel, err := url.Parse(path.Join("consumers", getter.GetConsumer().GetID(), "jwt", getter.GetID()))
	if err != nil {
		return errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return errgo.New(err)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return errgo.New(err)
	}

	expCodes := expectedCodes{http.StatusNoContent}
	if !expCodes.codeMatched(resp.StatusCode) {
		return NewErrKongResponse(expCodes, resp)
	}
	defer resp.Body.Close()

	return nil
}
