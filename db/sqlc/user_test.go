package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Domson12/social_media_rest/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func createRandomUser(t *testing.T) User {
	username := util.RandomOwner()
	bio := util.RandomString(6)

	// Generate a random password
	password := util.RandomString(6)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       sql.NullString{String: username, Valid: true},
		Email:          util.RandomEmail(),
		Password:       string(hashedPassword),
		ProfilePicture: sql.NullString{String: "profile_picture", Valid: true},
		Bio:            sql.NullString{String: bio, Valid: true},
		Role:           "user",
		LastActivityAt: time.Now(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.ProfilePicture, user.ProfilePicture)
	require.Equal(t, arg.Bio, user.Bio)
	require.Equal(t, arg.Role, user.Role)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ProfilePicture, user2.ProfilePicture)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.Role, user2.Role)
}

func TestGetUserByUsername(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)
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

func TestUpdateUser(t *testing.T) {

	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:             user1.ID,
		Username:       sql.NullString{String: "updated_username", Valid: true},
		Email:          util.RandomEmail(),
		ProfilePicture: sql.NullString{String: "profile_picture", Valid: true},
		Bio:            sql.NullString{String: util.RandomString(6), Valid: true},
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.ProfilePicture, user2.ProfilePicture)
	require.Equal(t, arg.Bio, user2.Bio)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
