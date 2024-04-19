-- name: CreateChatRoom :one
INSERT INTO chat_rooms (
name
) VALUES (
$1
) RETURNING *;

-- name: GetChatRoom :one
SELECT * FROM chat_rooms 
WHERE id = $1 LIMIT 1;

-- name: GetChatRooms :many
SELECT * FROM chat_rooms
LIMIT $1 OFFSET $2;

-- name: UpdateChatRoom :one
UPDATE chat_rooms SET
name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteChatRoom :exec
DELETE FROM chat_rooms WHERE id = $1;
