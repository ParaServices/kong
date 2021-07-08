package kong

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/object"
	"github.com/ParaServices/paratils"
)

// CreateConsumer creates a consumer for the KONG API gateway.
func (c *Client) CreateConsumer(getter object.ConsumerGetter) (*object.Consumer, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("consumer is nil")
	}
	customID := getter.GetCustomID()
	username := getter.GetUsername()
	if paratils.StringIsEmpty(customID) && paratils.StringIsEmpty(username) {
		return nil, errgo.NewF("custom ID and username is empty")
	}
	form := url.Values{}
	form.Add("custom_id", customID)
	form.Add("username", username)

	rel, err := url.Parse("consumers")
	if err != nil {
		return nil, errgo.New(err)
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, errgo.New(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
	}
	defer resp.Body.Close()

	expCodes := expectedCodes{http.StatusCreated}
	if !expCodes.codeMatched(resp.StatusCode) {
		return nil, NewErrKongResponse(expCodes, resp)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errgo.New(err)
	}

	consumer := &object.Consumer{}
	err = json.Unmarshal(b, &consumer)
	if err != nil {
		return nil, errgo.New(err)
	}

	return consumer, nil
}

// Delete Consumer requires usernameOrCustomID or ID to delete consumer via Kong API
// https://docs.konghq.com/0.14.x/admin-api/#delete-consumer
func (c *Client) DeleteConsumer(getter object.ConsumerGetter) error {
	rel, err := url.Parse(path.Join("consumers", getter.GetCustomID()))
	if err != nil {
		return errgo.New(err)
	}
	u := c.baseURL.ResolveReference(rel)
	// Create Request
	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return errgo.New(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.doRequest(req)
	if err != nil {
		return errgo.New(err)
	}
	defer resp.Body.Close()

	expCodes := expectedCodes{http.StatusNoContent}
	if !expCodes.codeMatched(resp.StatusCode) {
		return NewErrKongResponse(expCodes, resp)
	}
	return nil
}
