package kong

import (
	"testing"

	"github.com/ParaServices/kong/plugins"
	"github.com/stretchr/testify/require"
)

func TestClient_EnablePlugin(t *testing.T) {
	client := NewClient(1, 1, kongURL(t))

	t.Run("success", func(t *testing.T) {
		t.Run("config not given", func(t *testing.T) {
			plugins := []string{
				"cors",
				"jwt",
			}

			for i := range plugins {
				service := generateService(t)
				plugin := &Plugin{
					Name: plugins[i],
					Service: &PluginService{
						ID: service.ID,
					},
				}
				plugin, err := client.EnablePlugin(plugin)
				require.NoError(t, err)
			}
		})

		t.Run("config given", func(t *testing.T) {
			origins := []string{
				"localhost",
				"test.com",
			}
			config := &plugins.CORSConfig{
				Origins: origins,
			}
			service := generateService(t)
			plugin := &Plugin{
				Name: "cors",
				Service: &PluginService{
					ID: service.ID,
				},
				Config: config,
			}
			plugin, err := client.EnablePlugin(plugin)
			require.NoError(t, err)
			require.Equal(t, origins, config.Origins)
		})
	})
}
