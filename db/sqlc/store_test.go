package db

import (
	"context"
	"fmt"
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

func TestCreateMessageTx(t *testing.T) {
	store := NewStore(testDB)

	chatRoom := CreateRandomChatRoom(t)
	fmt.Printf("Chat Room ID: %v\n", chatRoom.ID)

	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateMessageTxParams{
		ChatRoomID:     chatRoom.ID,
		SenderUserID:   user1.ID,
		ReceiverUserID: user2.ID,
		Text:           "Hello",
	}

	n := 5
	errs := make(chan error)
	results := make(chan Message)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateMessageTx(context.Background(), arg)
			errs <- err
			results <- result.Message

		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.Equal(t, arg.SenderUserID, result.SenderUserID)
		require.Equal(t, arg.ReceiverUserID, result.ReceiverUserID)
		require.Equal(t, arg.Text, result.Text.String)
		require.NotZero(t, result.ID)
		require.NotZero(t, result.CreatedAt)
	}
}
