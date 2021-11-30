package kong

import (
	"testing"

	"github.com/ParaServices/kong/object"
	"github.com/ParaServices/kong/plugin"
	"github.com/stretchr/testify/require"
)

func TestClient_EnablePlugin(t *testing.T) {
	client, err := NewClient(kongURL(t))
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		t.Run("config not given", func(t *testing.T) {
			plugin := []string{
				"cors",
				"jwt",
			}

			for i := range plugin {
				service := generateService(t)
				enabledPlugin := &object.Plugin{
					Name: object.Name{
						Name: plugin[i],
					},
					Service: &object.KongID{
						ID: service.ID,
					},
				}
				enabledPlugin, err := client.EnablePlugin(enabledPlugin)
				require.NoError(t, err)
				require.NotNil(t, enabledPlugin)
			}
		})

		t.Run("config given", func(t *testing.T) {
			origins := []string{
				"localhost",
				"test.com",
			}
			config := &plugin.CORSConfig{
				Origins: origins,
			}
			service := generateService(t)
			plugin := &object.Plugin{
				Name: object.Name{
					Name: "cors",
				},
				Service: &object.KongID{
					ID: service.ID,
				},
			}
			err := plugin.CopyConfig(config)
			require.NoError(t, err)
			plugin, err = client.EnablePlugin(plugin)
			require.NoError(t, err)
			require.Equal(t, origins, config.Origins)
		})
	})
}
