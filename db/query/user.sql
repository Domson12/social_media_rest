-- name: CreateUser :one
INSERT INTO users (
username,
email,
password,
profile_picture,
bio,
role,
last_activity_at
) VALUES (
$1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET
username = $2,
email = $3,
profile_picture = $4,
bio = $5,
password = $6
WHERE id = $1
RETURNING *;

-- name: UpdateUserUsername :one
UPDATE users SET
username = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users SET
email = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserProfilePicture :one
UPDATE users SET
profile_picture = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserBio :one
UPDATE users SET
bio = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users SET
password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
