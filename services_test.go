package kong

import (
	"net/http"
	"testing"

	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/require"
)

func TestClient_AddService(t *testing.T) {
	client := NewClient(1, 1, kongURL(t))

	t.Run("success", func(t *testing.T) {
		t.Run("default values for attributes", func(t *testing.T) {
			svcName, err := tg.RandGen(20, tg.LowerUpper, "", "")
			require.NoError(t, err)
			svc := &Service{
				Name: svcName,
				Host: "service.com",
			}

			svcResp, err := client.AddService(svc)
			require.NoError(t, err)
			require.NotNil(t, svcResp)
			require.NotEmpty(t, svcResp.ID)
			require.Equal(t, svc.Name, svcResp.Name)
			require.Equal(t, svc.Host, svcResp.Host)
		})

		t.Run("defined values", func(t *testing.T) {
			svcName, err := tg.RandGen(20, tg.LowerUpper, "", "")
			require.NoError(t, err)
			svc := &Service{
				Name:           svcName,
				Protocol:       "https",
				Host:           "service.com",
				Port:           9999,
				Path:           "/api",
				Retries:        10,
				ConnectTimeout: 10000,
				WriteTimeout:   10000,
				ReadTimeout:    10000,
			}

			svcResp, err := client.AddService(svc)
			require.NoError(t, err)
			require.NotNil(t, svcResp)
			require.NotEmpty(t, svcResp.ID)
			require.Equal(t, svc.Name, svcResp.Name)
			require.Equal(t, svc.Protocol, svcResp.Protocol)
			require.Equal(t, svc.Host, svcResp.Host)
			require.Equal(t, svc.Port, svcResp.Port)
			require.Equal(t, svc.Path, svcResp.Path)
			require.Equal(t, svc.Retries, svcResp.Retries)
			require.Equal(t, svc.ConnectTimeout, svcResp.ConnectTimeout)
			require.Equal(t, svc.WriteTimeout, svcResp.WriteTimeout)
			require.Equal(t, svc.ReadTimeout, svcResp.ReadTimeout)
		})
	})

	t.Run("fail", func(t *testing.T) {
		t.Run("unique constraint violation", func(t *testing.T) {
			svcName, err := tg.RandGen(20, tg.LowerUpper, "", "")
			require.NoError(t, err)
			svc := &Service{
				Name: svcName,
				Host: "service.com",
			}

			svcResp, err := client.AddService(svc)
			require.NoError(t, err)
			require.NotNil(t, svc)
			require.NotEmpty(t, svcResp.ID)
			require.Equal(t, svc.Name, svcResp.Name)
			require.Equal(t, svc.Host, svcResp.Host)

			svcResp, err = client.AddService(svc)
			require.Error(t, err)
			require.Nil(t, svcResp)
			kongErr, ok := err.(KongError)
			require.True(t, ok)
			require.NotNil(t, kongErr)
			require.Equal(t, http.StatusConflict, kongErr.Response().StatusCode)
		})
	})
}

func TestClient_UpdateService(t *testing.T) {
	client := NewClient(1, 1, kongURL(t))

	t.Run("success", func(t *testing.T) {
		t.Run("new values for attributes", func(t *testing.T) {
			svcName, err := tg.RandGen(20, tg.LowerUpper, "", "")
			require.NoError(t, err)
			svc := &Service{
				Name: svcName,
				Host: "service.com",
			}

			svcResp, err := client.AddService(svc)
			require.NoError(t, err)
			require.NotNil(t, svcResp)
			require.NotEmpty(t, svcResp.ID)
			require.Equal(t, svc.Name, svcResp.Name)
			require.Equal(t, svc.Host, svcResp.Host)

			svcNameUpdate, err := tg.RandGen(20, tg.LowerUpper, "", "")
			require.NoError(t, err)
			svcResp.Name = svcNameUpdate
			newHost := "newhost.com"
			svcResp.Host = newHost

			svcRespUpdate, err := client.UpdateService(svcResp)
			require.NoError(t, err)
			require.NotNil(t, svcResp)
			require.NotEmpty(t, svcRespUpdate.ID)
			require.Equal(t, svcNameUpdate, svcRespUpdate.Name)
			require.Equal(t, newHost, svcRespUpdate.Host)
		})
	})
}
