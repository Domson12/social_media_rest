// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read_receipts.sql

package db

import (
	"context"
	"database/sql"
)

const createReadReceipt = `-- name: CreateReadReceipt :one
INSERT INTO read_receipts (
message_id,
user_id,
read_at
) VALUES (
$1, $2, $3
) RETURNING message_id, user_id, read_at
`

type CreateReadReceiptParams struct {
	MessageID int32        `json:"message_id"`
	UserID    int32        `json:"user_id"`
	ReadAt    sql.NullTime `json:"read_at"`
}

func (q *Queries) CreateReadReceipt(ctx context.Context, arg CreateReadReceiptParams) (ReadReceipt, error) {
	row := q.queryRow(ctx, q.createReadReceiptStmt, createReadReceipt, arg.MessageID, arg.UserID, arg.ReadAt)
	var i ReadReceipt
	err := row.Scan(&i.MessageID, &i.UserID, &i.ReadAt)
	return i, err
}

const getReadReceipt = `-- name: GetReadReceipt :one
SELECT message_id, user_id, read_at FROM read_receipts 
WHERE message_id = $1 LIMIT 1
`

func (q *Queries) GetReadReceipt(ctx context.Context, messageID int32) (ReadReceipt, error) {
	row := q.queryRow(ctx, q.getReadReceiptStmt, getReadReceipt, messageID)
	var i ReadReceipt
	err := row.Scan(&i.MessageID, &i.UserID, &i.ReadAt)
	return i, err
}
