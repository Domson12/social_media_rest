-- name: AddComment :one
INSERT INTO comments (
post_id,
user_id,
text
) VALUES (
$1, $2, $3
) RETURNING *;

-- name: GetComment :one
SELECT * FROM comments 
WHERE id = $1 LIMIT 1;

-- name: GetComments :many
SELECT * FROM comments
LIMIT $1 OFFSET $2;

-- name: UpdateComment :one
UPDATE comments SET
text = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;


