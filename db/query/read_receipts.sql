-- name: CreateReadReceipt :one
INSERT INTO read_receipts (
message_id,
user_id,
read_at
) VALUES (
$1, $2, $3
) RETURNING *;

-- name: GetReadReceipt :one
SELECT * FROM read_receipts 
WHERE message_id = $1 LIMIT 1;