-- name: CreateAccount :one
INSERT INTO users (
username,
email,
profile_picture,
bio,
role,
last_activity_at
) VALUES (
$1, $2, $3, $4, $5, $6
) RETURNING *;