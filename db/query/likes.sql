-- name: LikePost :exec
INSERT INTO likes (post_id, user_id) 
VALUES ($1, $2);

-- name: UnlikePost :exec
DELETE FROM likes 
WHERE post_id = $1 
AND user_id = $2;

-- name: GetLikes :many
SELECT * FROM likes;

-- name: GetLikesByPostId :many
SELECT * FROM likes 
WHERE post_id = $1;

-- name: GetLikesByUserId :many
SELECT * FROM likes
WHERE user_id = $1;
