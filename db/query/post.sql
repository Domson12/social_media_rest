-- name: CreatePost :one
INSERT INTO posts (
title,
body,
user_id,
likes_ids,
comments_ids,
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

-- name: UpdatePostTitle :one
UPDATE posts SET
title = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePostBody :one
UPDATE posts SET
body = $2
WHERE id = $1
RETURNING *;

-- name: AddLikeToPost :exec
UPDATE posts SET likes_ids = array_append(likes_ids, $2)
WHERE id = $1;

-- name: RemoveLikeFromPost :exec
UPDATE posts SET likes_ids = array_remove(likes_ids, $2)
WHERE id = $1;

-- name: AddCommentToPost :exec
UPDATE posts SET comments_ids = array_append(comments_ids, $2)
WHERE id = $1;

-- name: RemoveCommentFromPost :exec
UPDATE posts SET comments_ids = array_remove(comments_ids, $2)
WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;

-- name: DeleteUserPosts :exec
DELETE FROM posts WHERE user_id = $1;

-- name: ListPostsByUserId :many
SELECT * FROM posts
WHERE user_id = $1;