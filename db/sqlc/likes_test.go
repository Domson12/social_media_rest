package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLikePost(t *testing.T) {
	// Create a random user
	user := createRandomUser(t)

	// Create a random post
	post := createRandomPost(t)

	// Like the post
	arg := LikePostParams{
		PostID: post.ID,
		UserID: user.ID,
	}
	err := testQueries.LikePost(context.Background(), arg)
	require.NoError(t, err)
}

func TestUnlikePost(t *testing.T) {
	// Create a random user
	user := createRandomUser(t)

	// Create a random post
	post := createRandomPost(t)

	// Like the post
	arg1 := LikePostParams{
		PostID: post.ID,
		UserID: user.ID,
	}
	err := testQueries.LikePost(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := UnlikePostParams{
		PostID: post.ID,
		UserID: user.ID,
	}

	// Unlike the post
	err = testQueries.UnlikePost(context.Background(), arg2)
	require.NoError(t, err)
}

func TestGetPostLikes(t *testing.T) {
	// Create a random user
	user := createRandomUser(t)

	// Create a random post
	post := createRandomPost(t)

	// Like the post
	arg := LikePostParams{
		PostID: post.ID,
		UserID: user.ID,
	}
	err := testQueries.LikePost(context.Background(), arg)
	require.NoError(t, err)

	// Get post likes
	likes, err := testQueries.GetLikesByUserId(context.Background(), arg.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, likes)
}

func TestGetLikesByPostId(t *testing.T) {
	// Create a random user
	user := createRandomUser(t)

	// Create a random post
	post := createRandomPost(t)

	// Like the post
	arg := LikePostParams{
		PostID: post.ID,
		UserID: user.ID,
	}
	err := testQueries.LikePost(context.Background(), arg)
	require.NoError(t, err)

	// Get post likes count
	count, err := testQueries.GetLikesByPostId(context.Background(), arg.PostID)
	require.NoError(t, err)
	require.NotZero(t, count)
}

func TestGetLikesByUserId(t *testing.T) {
	// Create a random user
	user := createRandomUser(t)

	// Create a random post
	post := createRandomPost(t)

	// Like the post
	arg := LikePostParams{
		PostID: post.ID,
		UserID: user.ID,
	}
	err := testQueries.LikePost(context.Background(), arg)
	require.NoError(t, err)

	// Get user likes
	likes, err := testQueries.GetLikesByUserId(context.Background(), arg.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, likes)
}
