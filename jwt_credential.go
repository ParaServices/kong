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

// CreateJWTCredential makes a request to kong's admin API to createa JWT
// credential from the given getter.
func (c *Client) CreateJWTCredential(getter plugin.JWTCredentialGetter) (*plugin.JWTCredential, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("jwt credential is nil")
	}
	if paratils.StringIsEmpty(getter.GetConsumer().GetID()) {
		return nil, errgo.NewF("consumer ID is empty")
	}
	if paratils.StringIsEmpty(getter.GetKey()) {
		return nil, errgo.NewF("key is empty")
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

	// marshal so that the given getter's consumer is carried over. This
	// means that since we have omitempty on the JSON tags for the
	// consumer fields, if it's not returned by the server, the marshaled
	// fields will remain the same.
	jwtCred := &plugin.JWTCredential{}
	if err := plugin.MarshalJWTCredential(getter, jwtCred); err != nil {
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
	if paratils.StringIsEmpty(getter.GetConsumer().GetID()) {
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
