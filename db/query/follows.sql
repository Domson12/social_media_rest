-- name: AddFollow :one
-- Add a follow relationship between two users
INSERT INTO follows (following_user_id, followed_user_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveFollow :one
-- Remove a follow relationship between two users
DELETE FROM follows
WHERE following_user_id = $1 AND followed_user_id = $2 RETURNING *;