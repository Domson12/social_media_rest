
-- name: AddChatRoomParticipant :one
INSERT INTO chat_room_participants (
chat_room_id,
user_id
) VALUES (
$1, $2
) RETURNING *;

-- name: GetChatRoomParticipant :one
SELECT * FROM chat_room_participants 
WHERE user_id = $1 LIMIT 1;

-- name: GetChatRoomParticipants :many
SELECT * FROM chat_room_participants
LIMIT $1 OFFSET $2;

-- name: DeleteChatRoomParticipant :exec
DELETE FROM chat_room_participants WHERE user_id = $1;
