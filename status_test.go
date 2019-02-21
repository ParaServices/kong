package kong

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStatus(t *testing.T) {
	u, err := url.Parse(kongURL())
	require.NoError(t, err)
	client := NewClient(1, 1, u)
	statusResp, err := client.GetStatus()
	require.NoError(t, err, "Get Status Should Not Return Error")
	require.NotNil(t, statusResp, "Status Response should not be nil")
	require.NotEmpty(t, statusResp.Database.Reachable)
	require.NotEmpty(t, statusResp.Server.TotalRequests)
}
