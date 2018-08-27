package kong

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStatus(t *testing.T) {
	respBody := []byte(`
	{
		"server": {
			"total_requests": 3,
			"connections_active": 1,
			"connections_accepted": 1,
			"connections_handled": 1,
			"connections_reading": 0,
			"connections_writing": 1,
			"connections_waiting": 0
		},
		"database": {
			"reachable": true
		}
	}
	`)
	// Setup
	c, mux, server := setup()
	defer teardown(server)
	p := path.Join("/status")
	setupHandleFunc(t, mux, p, "GET", http.StatusOK, respBody)
	// Send Status Request
	statusResp, err := c.GetStatus()
	require.NoError(t, err, "Get Status Should Not Return Error")
	assert.NotNil(t, statusResp, "Status Response should not be nil")
	assert.Equal(t, true, statusResp.Database.Reachable)
	assert.Equal(t, 3, statusResp.Server.TotalRequests)
}
