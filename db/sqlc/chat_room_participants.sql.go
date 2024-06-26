// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: chat_room_participants.sql

package db

import (
	"context"
)

const addChatRoomParticipant = `-- name: AddChatRoomParticipant :one
INSERT INTO chat_room_participants (
chat_room_id,
user_id
) VALUES (
$1, $2
) RETURNING chat_room_id, user_id
`

type AddChatRoomParticipantParams struct {
	ChatRoomID int32 `json:"chat_room_id"`
	UserID     int32 `json:"user_id"`
}

func (q *Queries) AddChatRoomParticipant(ctx context.Context, arg AddChatRoomParticipantParams) (ChatRoomParticipant, error) {
	row := q.queryRow(ctx, q.addChatRoomParticipantStmt, addChatRoomParticipant, arg.ChatRoomID, arg.UserID)
	var i ChatRoomParticipant
	err := row.Scan(&i.ChatRoomID, &i.UserID)
	return i, err
}

const deleteChatRoomParticipant = `-- name: DeleteChatRoomParticipant :exec
DELETE FROM chat_room_participants WHERE user_id = $1
`

func (q *Queries) DeleteChatRoomParticipant(ctx context.Context, userID int32) error {
	_, err := q.exec(ctx, q.deleteChatRoomParticipantStmt, deleteChatRoomParticipant, userID)
	return err
}

const getChatRoomParticipant = `-- name: GetChatRoomParticipant :one
SELECT chat_room_id, user_id FROM chat_room_participants 
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetChatRoomParticipant(ctx context.Context, userID int32) (ChatRoomParticipant, error) {
	row := q.queryRow(ctx, q.getChatRoomParticipantStmt, getChatRoomParticipant, userID)
	var i ChatRoomParticipant
	err := row.Scan(&i.ChatRoomID, &i.UserID)
	return i, err
}

const getChatRoomParticipants = `-- name: GetChatRoomParticipants :many
SELECT chat_room_id, user_id FROM chat_room_participants
LIMIT $1 OFFSET $2
`

type GetChatRoomParticipantsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetChatRoomParticipants(ctx context.Context, arg GetChatRoomParticipantsParams) ([]ChatRoomParticipant, error) {
	rows, err := q.query(ctx, q.getChatRoomParticipantsStmt, getChatRoomParticipants, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ChatRoomParticipant{}
	for rows.Next() {
		var i ChatRoomParticipant
		if err := rows.Scan(&i.ChatRoomID, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
