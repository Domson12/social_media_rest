package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Domson12/social_media_rest/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	username := util.RandomOwner()
	bio := util.RandomString(6)
	arg := CreateAccountParams{
		Username:       sql.NullString{String: username, Valid: true},
		Email:          util.RandomEmail(),
		ProfilePicture: sql.NullString{String: "profile_picture", Valid: true},
		Bio:            sql.NullString{String: bio, Valid: true},
		Role:           "user",
		LastActivityAt: time.Now(),
	}
	user, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.ProfilePicture, user.ProfilePicture)
	require.Equal(t, arg.Bio, user.Bio)
	require.Equal(t, arg.Role, user.Role)

	return user
}

func TestCreateAccount(t *testing.T) {
	createRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetAccount(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ProfilePicture, user2.ProfilePicture)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.Role, user2.Role)
}

func TestGetAccountByUsername(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetAccountByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ProfilePicture, user2.ProfilePicture)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.Role, user2.Role)
}

func TestGetUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := GetUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.GetUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateAccount(t *testing.T) {

	user1 := createRandomUser(t)

	arg := UpdateAccountParams{
		ID:             user1.ID,
		Username:       sql.NullString{String: "updated_username", Valid: true},
		Email:          util.RandomEmail(),
		ProfilePicture: sql.NullString{String: "profile_picture", Valid: true},
		Bio:            sql.NullString{String: util.RandomString(6), Valid: true},
	}

	user2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.ProfilePicture, user2.ProfilePicture)
	require.Equal(t, arg.Bio, user2.Bio)
}

func TestDeleteAccount(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteAccount(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetAccount(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
