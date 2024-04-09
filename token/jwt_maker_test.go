package token

import (
	"testing"
	"time"

	"github.com/Domson12/social_media_rest/util"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Error(err)
	}

	userID := util.RandomInt(1, 1000)
	duration := time.Minute

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(int32(userID), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiresAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJwtToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Error(err)
	}

	token, err := maker.CreateToken(1, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "error parsing JWT token: token has invalid claims: token is expired")
	require.Nil(t, payload)
}

func TestInvalidJwtToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	if err != nil {
		t.Error(err)
	}

	token, err := maker.CreateToken(1, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Alter the token
	runes := []rune(token)
	runes[20] = 'a'
	invalidToken := string(runes)

	payload, err := maker.VerifyToken(invalidToken)
	require.Error(t, err)
	require.EqualError(t, err, "error parsing JWT token: token contains an invalid number of segments")
	require.Nil(t, payload)
}
