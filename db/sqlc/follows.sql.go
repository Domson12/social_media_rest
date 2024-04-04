// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: follows.sql

package db

import (
	"context"
)

const addFollow = `-- name: AddFollow :one
INSERT INTO follows (following_user_id, followed_user_id)
VALUES ($1, $2)
RETURNING following_user_id, followed_user_id, created_at
`

type AddFollowParams struct {
	FollowingUserID int32 `json:"following_user_id"`
	FollowedUserID  int32 `json:"followed_user_id"`
}

// Add a follow relationship between two users
func (q *Queries) AddFollow(ctx context.Context, arg AddFollowParams) (Follow, error) {
	row := q.queryRow(ctx, q.addFollowStmt, addFollow, arg.FollowingUserID, arg.FollowedUserID)
	var i Follow
	err := row.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt)
	return i, err
}

const removeFollow = `-- name: RemoveFollow :one
DELETE FROM follows
WHERE following_user_id = $1 AND followed_user_id = $2 RETURNING following_user_id, followed_user_id, created_at
`

type RemoveFollowParams struct {
	FollowingUserID int32 `json:"following_user_id"`
	FollowedUserID  int32 `json:"followed_user_id"`
}

// Remove a follow relationship between two users
func (q *Queries) RemoveFollow(ctx context.Context, arg RemoveFollowParams) (Follow, error) {
	row := q.queryRow(ctx, q.removeFollowStmt, removeFollow, arg.FollowingUserID, arg.FollowedUserID)
	var i Follow
	err := row.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt)
	return i, err
}
