-- name: CreatePost :one
INSERT INTO posts (
title,
body,
user_id,
likes_count,
comments_count,
status
) VALUES (
$1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts 
WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts
LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts SET
title = $2,
body = $3
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;