package kong

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJWTCredential(t *testing.T) {
	consumerID := "manbearpig"
	key := consumerID
	secret := consumerID
	respBody := []byte(`
		{
			"consumer_id": "manbearpig",
			"created_at": 1442426001000,
			"id": "bcbfb45d-e391-42bf-c2ed-94e32946753a",
			"key": "manbearpig",
			"secret": "manbearpig"
		}
	`)

	c, mux, server := setup()
	defer teardown(server)

	p := path.Join("/consumers", consumerID, "jwt")

	setupHandleFunc(t, mux, p, "POST", http.StatusCreated, respBody)
	createResponse, createErr := c.CreateJWTCredential(consumerID, key, secret)
	assert.NoError(t, createErr, "no error")
	assert.NotNil(t, createResponse, "did not receive response")
	assert.NotNil(t, createResponse.ID, "did not receive ID from response body")
	assert.Equal(t, consumerID, createResponse.ConsumerID, "consumerID not equal")
	assert.Equal(t, key, createResponse.Key, "key not equal")
	assert.Equal(t, secret, createResponse.Secret, "secret not equal")
	// close server so we can create a new one for the next test
	server.Close()

	c, mux, server = setup()
	setupHandleFunc(t, mux, p, "POST", http.StatusOK, respBody)
	createResponse, createErr = c.CreateJWTCredential(consumerID, key, secret)
	assert.Error(t, createErr, "should return error")
	assert.Nil(t, createResponse, "received response")
}

func TestDeleteJWTCredential(t *testing.T) {
	consumerID := "manbearpig"
	jwtID := consumerID

	c, mux, server := setup()
	defer teardown(server)

	p := path.Join("/consumers", consumerID, "jwt", jwtID)

	setupHandleFunc(t, mux, p, "DELETE", http.StatusNoContent, nil)
	deleteErr := c.DeleteJWTCredential(consumerID, jwtID)
	assert.NoError(t, deleteErr, "should not return error")
	// close server so we can create a new one for the next test
	server.Close()

	c, mux, server = setup()
	setupHandleFunc(t, mux, p, "DELETE", http.StatusOK, nil)
	deleteErr = c.DeleteJWTCredential(consumerID, jwtID)
	assert.Error(t, deleteErr, "should return error")
}
