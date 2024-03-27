-- name: CreateMessage :one
INSERT INTO messages (
sender_user_id,
receiver_user_id,
text,
status
) VALUES (
$1, $2, $3, $4
) RETURNING *;

-- name: GetMessage :one
SELECT * FROM messages 
WHERE id = $1 LIMIT 1;

-- name: GetMessages :many
SELECT * FROM messages
LIMIT $1 OFFSET $2;

-- name: UpdateMessage :one
UPDATE messages SET
text = $1,
status = $2
WHERE id = $3 RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages WHERE id = $1;


