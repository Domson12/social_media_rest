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
