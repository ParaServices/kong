package kong

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConsumer(t *testing.T) {
	username := "manbearpig"
	respBody := []byte(`
		{
		  "username": "manbearpig",
		  "created_at": 1428555626000,
		  "id": "bbdf1c48-19dc-4ab7-cae0-ff4f59d87dc9"
		}
	`)

	client, mux, server := setup()
	defer teardown(server)

	setupHandleFunc(t, mux, "/consumers", "POST", http.StatusCreated, respBody)
	createResponse, createErr := client.CreateConsumer(username)
	assert.Nil(t, createErr, "no error")
	assert.NotNil(t, createResponse, "received response")
	assert.Equal(t, createResponse.Username, username, "username equal")
	// close server so we can create a new one for the next test
	server.Close()

	client, mux, server = setup()
	setupHandleFunc(t, mux, "/consumers", "POST", http.StatusOK, respBody)
	createResponse, createErr = client.CreateConsumer(username)
	assert.NotNil(t, createErr, "has error")
	assert.Nil(t, createResponse, "received no response")
}
