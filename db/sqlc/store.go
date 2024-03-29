package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateChatRoomTxParams struct {
	ChatRoomName string `json:"name"`
	UserID       int32  `json:"user_id"`
}

type CreateChatRoomTxResult struct {
	ChatRoom ChatRoom `json:"chat_room"`
}

type CreateMessageTxParams struct {
	SenderUserID   int32  `json:"sender_user_id"`
	ReceiverUserID int32  `json:"receiver_user_id"`
	ChatRoomID     int32  `json:"chat_room_id"`
	Text           string `json:"text"`
}

type CreateMessageTxResult struct {
	Message Message `json:"message"`
}

// create a chat room with a participant
func (store *Store) CreateChatRoomTx(ctx context.Context, arg CreateChatRoomTxParams) (CreateChatRoomTxResult, error) {

	var result CreateChatRoomTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		chatRoom, errTx := q.CreateChatRoom(ctx, arg.ChatRoomName)
		if errTx != nil {
			return errTx
		}

		_, errTx = q.AddChatRoomParticipant(ctx, AddChatRoomParticipantParams{
			ChatRoomID: chatRoom.ID,
			UserID:     arg.UserID,
		})
		if errTx != nil {
			return errTx
		}

		result.ChatRoom = chatRoom

		return nil
	})

	return result, err
}

// create a message with a read receipt
func (store *Store) CreateMessageTx(ctx context.Context, arg CreateMessageTxParams) (CreateMessageTxResult, error) {

	var result CreateMessageTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		message, errTx := q.CreateMessage(ctx, CreateMessageParams{
			SenderUserID:   arg.SenderUserID,
			ReceiverUserID: arg.ReceiverUserID,
			ChatRoomID:     arg.ChatRoomID,
			Text:           sql.NullString{String: arg.Text, Valid: true},
			Status:         "sent",
		})
		if errTx != nil {
			return errTx
		}

		_, errTx = q.CreateReadReceipt(ctx, CreateReadReceiptParams{
			MessageID: message.ID,
			UserID:    arg.ReceiverUserID,
			ReadAt:    sql.NullTime{Valid: false},
		})
		if errTx != nil {
			return errTx
		}

		result.Message = message

		return nil
	})

	return result, err
}
