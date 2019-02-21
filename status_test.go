package kong

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStatus(t *testing.T) {
	// respBody := []byte(`
	// {
	// 	"server": {
	// 		"total_requests": 3,
	// 		"connections_active": 1,
	// 		"connections_accepted": 1,
	// 		"connections_handled": 1,
	// 		"connections_reading": 0,
	// 		"connections_writing": 1,
	// 		"connections_waiting": 0
	// 	},
	// 	"database": {
	// 		"reachable": true
	// 	}
	// }
	// `)
	// // Setup
	// c, mux, server := setup()
	// defer teardown(server)
	// p := path.Join("/status")
	// setupHandleFunc(t, mux, p, "GET", http.StatusOK, respBody)
	// /
	// Send Status Request
	u, err := url.Parse(kongURL())
	require.NoError(t, err)
	client := NewClient(1, 1, u)
	statusResp, err := client.GetStatus()
	require.NoError(t, err, "Get Status Should Not Return Error")
	require.NotNil(t, statusResp, "Status Response should not be nil")
	require.NotEmpty(t, statusResp.Database.Reachable)
	require.NotEmpty(t, statusResp.Server.TotalRequests)
}
