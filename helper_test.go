package kong

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/ParaServices/kong/object"
	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/require"
)

func kongURL(t *testing.T) *url.URL {
	v := os.Getenv("KONG_URL")
	if v == "" {
		v = "http://localhost:8001"
	}
	u, err := url.ParseRequestURI(v)
	require.NoError(t, err)

	return u
}

func generateService(t *testing.T) *object.Service {
	client, err := NewClient(kongURL(t))
	require.NoError(t, err)
	svcName, err := tg.RandGen(20, tg.LowerUpper, "", "")
	require.NoError(t, err)
	svcHost, err := tg.RandGen(10, tg.Lower, "", "")
	require.NoError(t, err)
	svc := &object.Service{
		Name: object.Name{
			Name: svcName,
		},
		Host: fmt.Sprintf("%s.com", svcHost),
	}

	svcResp, err := client.AddService(svc)
	require.NoError(t, err)
	require.NotNil(t, svcResp)
	require.NotEmpty(t, svcResp.ID)
	require.Equal(t, svc.Name, svcResp.Name)
	require.Equal(t, svc.Host, svcResp.Host)
	return svcResp
}
