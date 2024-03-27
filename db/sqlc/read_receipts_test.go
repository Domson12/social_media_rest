package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomReadReceipt(t *testing.T) ReadReceipt {
	message := CreateRandomMessage(t)

	arg := CreateReadReceiptParams{
		MessageID: message.ID,
		UserID:    message.ReceiverUserID,
		ReadAt:    sql.NullTime{Time: time.Now(), Valid: true},
	}
	readReceipt, err := testQueries.CreateReadReceipt(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, readReceipt)

	return readReceipt
}

func TestCreateReadReceipt(t *testing.T) {
	readReceipt := CreateRandomReadReceipt(t)
	require.NotEmpty(t, readReceipt)

}

func TestGetReadReceipt(t *testing.T) {
	readReceipt1 := CreateRandomReadReceipt(t)
	readReceipt2, err := testQueries.GetReadReceipt(context.Background(), readReceipt1.MessageID)
	require.NoError(t, err)
	require.NotEmpty(t, readReceipt2)

	require.Equal(t, readReceipt1.MessageID, readReceipt2.MessageID)
	require.Equal(t, readReceipt1.UserID, readReceipt2.UserID)
}
