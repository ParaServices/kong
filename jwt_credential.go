package kong

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/plugin"
	"github.com/ParaServices/paratils"
)

// CreateJWTCredential makes a request to kong's admin API to createa JWT
// credential from the given getter.
func (c *Client) CreateJWTCredential(getter plugin.JWTCredentialGetter) (*plugin.JWTCredential, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("jwt credential is nil")
	}
	if !getter.HasConsumerID() {
		return nil, errgo.NewF("consumer ID is empty")
	}
	if paratils.StringIsEmpty(getter.GetKey()) {
		return nil, errgo.NewF("key is empty")
	}

	rel, err := url.Parse(path.Join("consumers", getter.GetConsumer().GetID(), "jwt"))
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	b, err := json.Marshal(getter)
	if err != nil {
		return nil, errgo.New(err)
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(b))
	if err != nil {
		return nil, errgo.New(err)
	}
	req.Header.Set("Content-Type", "application/json")

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

	jwtCred := &plugin.JWTCredential{}
	if err := plugin.CopyJWTCredential(getter, jwtCred); err != nil {
		return nil, errgo.New(err)
	}
	err = json.Unmarshal(b, &jwtCred)
	if err != nil {
		return nil, errgo.New(err)
	}
	return jwtCred, nil
}

// DeleteJWTCredential ...
func (c *Client) DeleteJWTCredential(getter plugin.JWTCredentialGetter) error {
	if paratils.IsNil(getter) {
		return errgo.NewF("jwt credential is nil")
	}
	if !getter.HasConsumerID() {
		return errgo.NewF("consumer ID is empty")
	}
	rel, err := url.Parse(path.Join("consumers", getter.GetConsumer().GetID(), "jwt", getter.GetID()))
	if err != nil {
		return errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return errgo.New(err)
	}
	req.Header.Add("Content-Type", "application/json")

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
