package token

import (
	"testing"
	"time"

	"github.com/Domson12/social_media_rest/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	t.Run("TestCreateToken", func(t *testing.T) {
		maker, err :=
			NewPasetoMaker(util.RandomString(32))
		if err != nil {
			t.Error(err)
		}

		userID := int32(util.RandomInt(1, 1000))
		duration := time.Minute

		token, err := maker.CreateToken(userID, duration)
		require.NoError(t, err)
		require.NotEmpty(t, token)
	})

	t.Run("TestVerifyToken", func(t *testing.T) {

		maker, err :=
			NewPasetoMaker(util.RandomString(32))
		if err != nil {
			t.Error(err)
		}

		userID := int32(util.RandomInt(1, 1000))
		duration := time.Minute

		token, err := maker.CreateToken(userID, duration)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		payload, err := maker.VerifyToken(token)
		require.NoError(t, err)
		require.NotEmpty(t, payload)
	})
}
