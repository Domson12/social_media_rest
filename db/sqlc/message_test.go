package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomMessage(t *testing.T) Message {

	user1 := createRandomUser(t)
	chatRoom := CreateRandomChatRoom(t)

	arg := CreateMessageParams{
		SenderUserID: user1.ID,
		ChatRoomID:   chatRoom.ID,
		Text:         sql.NullString{String: "Hello", Valid: true},
		Status:       "sent",
	}

	message, err := testQueries.CreateMessage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, message)

	require.Equal(t, arg.SenderUserID, message.SenderUserID)
	require.Equal(t, arg.Text, message.Text)
	require.Equal(t, arg.Status, message.Status)

	return message
}

func TestAddMessage(t *testing.T) {
	CreateRandomMessage(t)
}

func TestGetMessage(t *testing.T) {
	message1 := CreateRandomMessage(t)
	message2, err := testQueries.GetMessage(context.Background(), message1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, message2)

	require.Equal(t, message1.SenderUserID, message2.SenderUserID)
	require.Equal(t, message1.Text, message2.Text)
	require.Equal(t, message1.Status, message2.Status)
}

func TestGetMessages(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomMessage(t)
	}

	arg := GetMessagesParams{
		Limit:  5,
		Offset: 5,
	}

	messages, err := testQueries.GetMessages(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, messages, 5)

	for _, message := range messages {
		require.NotEmpty(t, message)
	}
}

func TestUpdateMessage(t *testing.T) {
	message1 := CreateRandomMessage(t)

	arg := UpdateMessageParams{
		ID:     message1.ID,
		Text:   sql.NullString{String: "Hello World", Valid: true},
		Status: "delivered",
	}

	message2, err := testQueries.UpdateMessage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, message2)

	require.Equal(t, message1.ID, message2.ID)
	require.Equal(t, arg.Text, message2.Text)
	require.Equal(t, arg.Status, message2.Status)
}

func TestDeleteMessage(t *testing.T) {
	message1 := CreateRandomMessage(t)
	err := testQueries.DeleteMessage(context.Background(), message1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, message1)

	message2, err := testQueries.GetMessage(context.Background(), message1.ID)
	require.Error(t, err)
	require.Empty(t, message2)
}
