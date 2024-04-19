package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)
	arg := CreatePostParams{
		Title:       sql.NullString{String: "Hello", Valid: true},
		Body:        sql.NullString{String: "Hello", Valid: true},
		LikesIds:    []int32{},
		CommentsIds: []int32{},
		UserID:      user.ID,
		Status:      "active",
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	return post

}
func TestAddPost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := createRandomPost(t)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Body, post2.Body)
	require.Equal(t, post1.UserID, post2.UserID)
}

func TestGetPosts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}

	arg := GetPostsParams{
		Limit:  5,
		Offset: 5,
	}
	posts, err := testQueries.GetPosts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}

func TestUpdatePost(t *testing.T) {
	post1 := createRandomPost(t)
	arg := UpdatePostParams{
		ID: post1.ID,
		// //*string value
		// Title: "Hello",
		// Body:  "Hello",
	}
	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, arg.Title, post2.Title)
	require.Equal(t, arg.Body, post2.Body)
}

func TestDeletePost(t *testing.T) {
	post1 := createRandomPost(t)
	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}
