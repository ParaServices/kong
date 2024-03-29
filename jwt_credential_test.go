package kong

import (
	"net/http"
	"testing"

	"github.com/ParaServices/kong/object"
	"github.com/ParaServices/kong/plugin"
	"github.com/magicalbanana/tg"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateJWTCredential(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		username, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		customID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		consumer, err := object.NewConsumer(
			object.SetConsumerCustomID(customID),
			object.SetConsumerUsername(username),
		)
		require.NoError(t, err)
		createdConsumer, err := client.CreateConsumer(consumer)
		require.NoError(t, err)
		require.NotNil(t, createdConsumer)

		key, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		secret, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)

		jwtCred, err := plugin.NewJWTCredential(
			plugin.SetJWTCredentialConsumer(
				createdConsumer,
			),
			plugin.SetJWTCredentialKey(key),
			plugin.SetJWTCredentialSecret(secret),
		)
		require.NoError(t, err)

		resp, err := client.CreateJWTCredential(jwtCred)
		require.NoError(t, err)
		require.NotNil(t, resp)
		// the ID that's returned is not the custom ID that was used to create
		// the consumer but instead the ID generated by Kong. The
		// given JWtCredential that has Consumer data is also
		// marshaled into the new one. But it is expected that the
		// consumer.id is equal to the created consumer ID.
		require.Equal(t, createdConsumer.GetID(), resp.GetConsumer().GetID())
		require.Equal(t, createdConsumer.GetUsername(), resp.GetConsumer().GetUsername())
		require.Equal(t, createdConsumer.GetCustomID(), resp.GetConsumer().GetCustomID())
		require.Equal(t, createdConsumer.GetCreatedAt(), resp.GetConsumer().GetCreatedAt())
		require.NotEmpty(t, resp.ID)
	})

	t.Run("consumer does not exist", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		jwtCred, err := plugin.NewJWTCredential(
			plugin.SetJWTCredentialNewConsumer(
				object.NewConsumer,
				object.SetConsumerID("1"),
				object.SetConsumerCustomID("1"),
				object.SetConsumerUsername("a"),
			),
			plugin.SetJWTCredentialKey("key"),
			plugin.SetJWTCredentialSecret("secret"),
		)
		require.NoError(t, err)

		resp, err := client.CreateJWTCredential(jwtCred)
		require.Error(t, err)
		errx := err.(KongError)
		require.Equal(t, http.StatusNotFound, errx.ResponseCode())
		require.Nil(t, resp)
	})
}

func TestClient_DeleteJWTCredential(t *testing.T) {
	t.Run("delete success", func(t *testing.T) {
		client, err := NewClient(kongURL(t))
		require.NoError(t, err)

		username, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		customID, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		consumer, err := object.NewConsumer(
			object.SetConsumerUsername(username),
			object.SetConsumerCustomID(customID),
		)
		require.NoError(t, err)
		createdConsumer, err := client.CreateConsumer(consumer)
		require.NoError(t, err)
		require.NotNil(t, createdConsumer)

		key, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)
		secret, err := tg.RandGen(10, tg.Digit, "", "")
		require.NoError(t, err)

		jwtCred, err := plugin.NewJWTCredential(
			plugin.SetJWTCredentialConsumer(
				createdConsumer,
			),
			plugin.SetJWTCredentialKey(key),
			plugin.SetJWTCredentialSecret(secret),
		)
		require.NoError(t, err)

		createdJWTCred, err := client.CreateJWTCredential(jwtCred)
		require.NoError(t, err)
		require.NotNil(t, createdJWTCred)
		require.Equal(t, createdConsumer.GetID(), createdJWTCred.GetConsumer().GetID())
		require.NotEmpty(t, createdJWTCred.ID)

		err = client.DeleteJWTCredential(createdJWTCred)
		require.NoError(t, err)
	})
}
