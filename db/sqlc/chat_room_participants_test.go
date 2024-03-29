package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomChatRoomParticipant(t *testing.T) ChatRoomParticipant {
	chatRoom := CreateRandomChatRoom(t)
	user := createRandomUser(t)
	arg := AddChatRoomParticipantParams{
		ChatRoomID: chatRoom.ID,
		UserID:     user.ID,
	}
	chatRoomParticipant, err := testQueries.AddChatRoomParticipant(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, chatRoomParticipant)

	require.Equal(t, arg.ChatRoomID, chatRoomParticipant.ChatRoomID)
	require.Equal(t, arg.UserID, chatRoomParticipant.UserID)

	return chatRoomParticipant
}

func TestAddChatRoomParticipant(t *testing.T) {
	CreateRandomChatRoomParticipant(t)
}

func TestGetChatRoomParticipant(t *testing.T) {
	chatRoomParticipant1 := CreateRandomChatRoomParticipant(t)
	chatRoomParticipant2, err := testQueries.GetChatRoomParticipant(context.Background(), chatRoomParticipant1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, chatRoomParticipant2)

	require.Equal(t, chatRoomParticipant1.ChatRoomID, chatRoomParticipant2.ChatRoomID)
	require.Equal(t, chatRoomParticipant1.UserID, chatRoomParticipant2.UserID)
}

func TestGetChatRoomParticipants(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomChatRoomParticipant(t)
	}

	arg := GetChatRoomParticipantsParams{
		Limit:  5,
		Offset: 5,
	}
	chatRoomParticipants, err := testQueries.GetChatRoomParticipants(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, chatRoomParticipants, 5)

	for _, chatRoomParticipant := range chatRoomParticipants {
		require.NotEmpty(t, chatRoomParticipant)
	}
}

func TestRemoveChatRoomParticipant(t *testing.T) {
	chatRoomParticipant1 := CreateRandomChatRoomParticipant(t)
	err := testQueries.DeleteChatRoomParticipant(context.Background(), chatRoomParticipant1.UserID)
	require.NoError(t, err)

	chatRoomParticipant2, err := testQueries.GetChatRoomParticipant(context.Background(), chatRoomParticipant1.UserID)
	require.Error(t, err)
	require.Empty(t, chatRoomParticipant2)
}
