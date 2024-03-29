// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: comments.sql

package db

import (
	"context"
)

const addComment = `-- name: AddComment :one
INSERT INTO comments (
post_id,
user_id,
text
) VALUES (
$1, $2, $3
) RETURNING id, post_id, user_id, text, created_at
`

type AddCommentParams struct {
	PostID int32  `json:"post_id"`
	UserID int32  `json:"user_id"`
	Text   string `json:"text"`
}

func (q *Queries) AddComment(ctx context.Context, arg AddCommentParams) (Comment, error) {
	row := q.queryRow(ctx, q.addCommentStmt, addComment, arg.PostID, arg.UserID, arg.Text)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Text,
		&i.CreatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteCommentStmt, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, post_id, user_id, text, created_at FROM comments 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetComment(ctx context.Context, id int32) (Comment, error) {
	row := q.queryRow(ctx, q.getCommentStmt, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Text,
		&i.CreatedAt,
	)
	return i, err
}

const getComments = `-- name: GetComments :many
SELECT id, post_id, user_id, text, created_at FROM comments
LIMIT $1 OFFSET $2
`

type GetCommentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetComments(ctx context.Context, arg GetCommentsParams) ([]Comment, error) {
	rows, err := q.query(ctx, q.getCommentsStmt, getComments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.Text,
			&i.CreatedAt,
		); err != nil {
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

const updateComment = `-- name: UpdateComment :one
UPDATE comments SET
text = $2
WHERE id = $1
RETURNING id, post_id, user_id, text, created_at
`

type UpdateCommentParams struct {
	ID   int32  `json:"id"`
	Text string `json:"text"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.queryRow(ctx, q.updateCommentStmt, updateComment, arg.ID, arg.Text)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Text,
		&i.CreatedAt,
	)
	return i, err
}