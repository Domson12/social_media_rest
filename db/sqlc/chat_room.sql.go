// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: chat_room.sql

package db

import (
	"context"
)

const createChatRoom = `-- name: CreateChatRoom :one
INSERT INTO chat_rooms (
name
) VALUES (
$1
) RETURNING id, name, created_at
`

func (q *Queries) CreateChatRoom(ctx context.Context, name string) (ChatRoom, error) {
	row := q.queryRow(ctx, q.createChatRoomStmt, createChatRoom, name)
	var i ChatRoom
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteChatRoom = `-- name: DeleteChatRoom :exec
DELETE FROM chat_rooms WHERE id = $1
`

func (q *Queries) DeleteChatRoom(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteChatRoomStmt, deleteChatRoom, id)
	return err
}

const getChatRoom = `-- name: GetChatRoom :one
SELECT id, name, created_at FROM chat_rooms 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetChatRoom(ctx context.Context, id int32) (ChatRoom, error) {
	row := q.queryRow(ctx, q.getChatRoomStmt, getChatRoom, id)
	var i ChatRoom
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getChatRooms = `-- name: GetChatRooms :many
SELECT id, name, created_at FROM chat_rooms
LIMIT $1 OFFSET $2
`

type GetChatRoomsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetChatRooms(ctx context.Context, arg GetChatRoomsParams) ([]ChatRoom, error) {
	rows, err := q.query(ctx, q.getChatRoomsStmt, getChatRooms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ChatRoom{}
	for rows.Next() {
		var i ChatRoom
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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

const updateChatRoom = `-- name: UpdateChatRoom :one
UPDATE chat_rooms SET
name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateChatRoomParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateChatRoom(ctx context.Context, arg UpdateChatRoomParams) (ChatRoom, error) {
	row := q.queryRow(ctx, q.updateChatRoomStmt, updateChatRoom, arg.ID, arg.Name)
	var i ChatRoom
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
