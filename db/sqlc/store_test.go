package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddChatRoomTx(t *testing.T) {
	store := NewStore(testDB)

	chatRoom := CreateRandomChatRoom(t)

	user := createRandomUser(t)

	arg := CreateChatRoomTxParams{
		ChatRoomName: chatRoom.Name,
		UserID:       user.ID,
	}

	n := 5
	errs := make(chan error)
	results := make(chan ChatRoom)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateChatRoomTx(context.Background(), arg)
			errs <- err
			results <- result.ChatRoom

		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.Equal(t, chatRoom.Name, result.Name)
		require.NotZero(t, result.ID)
		require.NotZero(t, result.CreatedAt)

	}
}
