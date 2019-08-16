package kong

import (
	"net/http"
	"net/url"
	"sync"
	"testing"

	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateConsumer(t *testing.T) {
	u, err := url.Parse(kongURL())
	require.NoError(t, err)

	t.Run("create success", func(t *testing.T) {
		client := NewClient(1, 1, u)

		usernameOrCustomID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		resp, err := client.CreateConsumer(usernameOrCustomID)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, usernameOrCustomID, resp.CustomID)

		// test concurrent creates
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(t *testing.T, client Client) {
				defer wg.Done()
				usernameOrCustomID, err := tg.RandGen(20, tg.Digit, "", "")
				require.NoError(t, err)
				resp, err := client.CreateConsumer(usernameOrCustomID)
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, usernameOrCustomID, resp.CustomID)
			}(t, client)
		}
		wg.Wait()
	})

	t.Run("unique violation", func(t *testing.T) {
		client := NewClient(1, 1, u)

		usernameOrCustomID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		resp, err := client.CreateConsumer(usernameOrCustomID)
		require.NoError(t, err)
		require.NotNil(t, resp)

		resp, err = client.CreateConsumer(usernameOrCustomID)
		require.Error(t, err)
		errx := err.(Error)
		require.Equal(t, http.StatusConflict, errx.ResponseCode())
		require.Nil(t, resp)
	})
}
