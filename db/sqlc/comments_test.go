package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) Comment {
	user := createRandomUser(t)
	post := createRandomPost(t)

	arg := AddCommentParams{
		PostID: post.ID,
		UserID: user.ID,
		Text:   "Hello",
	}

	comment, err := testQueries.AddComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.PostID, comment.PostID)
	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.Text, comment.Text)

	require.NotZero(t, comment.ID)
	require.NotZero(t, comment.CreatedAt)

	return comment
}

func TestAddComment(t *testing.T) {
	createRandomComment(t)
}
