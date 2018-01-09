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

	client, mux, server := setup()
	defer teardown(server)

	p := path.Join("/consumers", consumerID, "jwt")

	setupHandleFunc(t, mux, p, "POST", http.StatusCreated, respBody)
	createResponse, createErr := client.CreateJWTCredential(consumerID, key, secret)
	assert.Nil(t, createErr, "no error")
	assert.NotNil(t, createResponse, "received response")
	assert.NotNil(t, createResponse.ID, "received ID from response body")
	assert.Equal(t, createResponse.ConsumerID, consumerID, "consumerID equal")
	assert.Equal(t, createResponse.Key, key, "key equal")
	assert.Equal(t, createResponse.Secret, secret, "secret equal")
	// close server so we can create a new one for the next test
	server.Close()

	client, mux, server = setup()
	setupHandleFunc(t, mux, p, "POST", http.StatusOK, respBody)
	createResponse, createErr = client.CreateJWTCredential(consumerID, key, secret)
	assert.NotNil(t, createErr, "has error")
	assert.Nil(t, createResponse, "received no response")
}

func TestDeleteJWTCredential(t *testing.T) {
	consumerID := "manbearpig"
	jwtID := consumerID

	client, mux, server := setup()
	defer teardown(server)

	p := path.Join("/consumers", consumerID, "jwt", jwtID)

	setupHandleFunc(t, mux, p, "DELETE", http.StatusNoContent, nil)
	deleteErr := client.DeleteJWTCredential(consumerID, jwtID)
	assert.Nil(t, deleteErr, "no error")
	// close server so we can create a new one for the next test
	server.Close()

	client, mux, server = setup()
	setupHandleFunc(t, mux, p, "DELETE", http.StatusOK, nil)
	deleteErr = client.DeleteJWTCredential(consumerID, jwtID)
	assert.NotNil(t, deleteErr, "has error")
}
