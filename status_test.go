package kong

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStatus(t *testing.T) {
	client, err := NewClient(kongURL(t))
	require.NoError(t, err)
	statusResp, err := client.GetStatus()
	require.NoError(t, err, "Get Status Should Not Return Error")
	require.NotNil(t, statusResp, "Status Response should not be nil")
	require.NotEmpty(t, statusResp.Database.Reachable)
	require.NotEmpty(t, statusResp.Server.TotalRequests)
}
