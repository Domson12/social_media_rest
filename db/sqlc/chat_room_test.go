package db

import (
	"context"
	"testing"

	"github.com/Domson12/social_media_rest/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomChatRoom(t *testing.T) ChatRoom {
	arg := util.RandomString(6)
	chatRoom, err := testQueries.CreateChatRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, chatRoom)

	return chatRoom

}

func TestAddChatRoom(t *testing.T) {
	chatRoom1 := CreateRandomChatRoom(t)
	chatRoom2, err := testQueries.GetChatRoom(context.Background(), chatRoom1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, chatRoom2)

	require.Equal(t, chatRoom1.Name, chatRoom2.Name)
}

func TestGetChatRooms(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomChatRoom(t)
	}

	arg := GetChatRoomsParams{
		Limit:  5,
		Offset: 5,
	}
	chatRooms, err := testQueries.GetChatRooms(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, chatRooms, 5)

	for _, chatRoom := range chatRooms {
		require.NotEmpty(t, chatRoom)
	}
}

func TestUpdateChatRoom(t *testing.T) {
	chatRoom1 := CreateRandomChatRoom(t)
	arg := UpdateChatRoomParams{
		ID:   chatRoom1.ID,
		Name: util.RandomString(6),
	}
	chatRoom2, err := testQueries.UpdateChatRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, chatRoom2)

	require.Equal(t, chatRoom1.ID, chatRoom2.ID)
	require.Equal(t, arg.Name, chatRoom2.Name)
}

func TestDeleteChatRoom(t *testing.T) {
	chatRoom1 := CreateRandomChatRoom(t)
	err := testQueries.DeleteChatRoom(context.Background(), chatRoom1.ID)
	require.NoError(t, err)

	chatRoom2, err := testQueries.GetChatRoom(context.Background(), chatRoom1.ID)
	require.Error(t, err)
	require.Empty(t, chatRoom2)
}
