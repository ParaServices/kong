package kong

import (
	"net/http"
	"sync"
	"testing"

	"github.com/ParaServices/kong/object"
	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateConsumer(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		consumerID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		newConsumer := object.NewConsumer(consumerID, consumerID)
		resp, err := client.CreateConsumer(newConsumer)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, consumerID, resp.CustomID)

		// test concurrent creates
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(t *testing.T, client *Client) {
				defer wg.Done()
				consumerID, err := tg.RandGen(20, tg.Digit, "", "")
				require.NoError(t, err)
				username, err := tg.RandGen(20, tg.LowerUpper, "", "")
				require.NoError(t, err)
				newConsumer := object.NewConsumer(username, consumerID)
				createdConsumer, err := client.CreateConsumer(newConsumer)
				require.NoError(t, err)
				require.NotNil(t, createdConsumer)
				require.Equal(t, username, createdConsumer.GetUsername())
				require.Equal(t, consumerID, createdConsumer.GetCustomID())
			}(t, client)
		}
		wg.Wait()
	})

	t.Run("unique violation", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		consumerID, err := tg.RandGen(20, tg.Digit, "", "")
		require.NoError(t, err)
		newConsumer := object.NewConsumer(consumerID, consumerID)
		resp, err := client.CreateConsumer(newConsumer)
		require.NoError(t, err)
		require.NotNil(t, resp)

		resp, err = client.CreateConsumer(newConsumer)
		require.Error(t, err)
		errx := err.(KongError)
		require.Equal(t, http.StatusConflict, errx.ResponseCode())
		require.Nil(t, resp)
	})
}

func TestClient_DeleteConsumer(t *testing.T) {
	t.Run("delete success", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		customID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		username, err := tg.RandGen(32, tg.LowerUpper, "", "")
		require.NoError(t, err)
		consumer := object.NewConsumer(username, customID)
		createdConsumer, err := client.CreateConsumer(consumer)
		require.NoError(t, err)
		require.NotNil(t, createdConsumer)
		require.Equal(t, username, createdConsumer.GetUsername())
		require.Equal(t, customID, createdConsumer.GetCustomID())

		err = client.DeleteConsumer(consumer)
		require.NoError(t, err)
	})

	t.Run("does not exist", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		customID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		username, err := tg.RandGen(32, tg.LowerUpper, "", "")
		require.NoError(t, err)
		consumer := object.NewConsumer(username, customID)

		err = client.DeleteConsumer(consumer)
		require.NoError(t, err)
	})
}
