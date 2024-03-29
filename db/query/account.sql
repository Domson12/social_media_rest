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

-- name: GetAccount :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE users SET
username = $2,
email = $3,
profile_picture = $4,
bio = $5
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM users WHERE id = $1;
