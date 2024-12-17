-- name: createTransaction :one
INSERT INTO
    transfers ("from_account_id", "to_account_id", "amount") VALUES ($1, $2, $3) RETURNING *;


-- name: getTransfer :one
SELECT * FROM transfers WHERE id=$1;

-- name: listAllTransger :many
SELECT * FROM transfers LIMIT $1 OFFSET $2;

-- name: updateTransfer :one
UPDATE transfers SET amount=$2 WHERE id=$1 RETURNING *;

-- name: deleteTransfer :exec
DELETE FROM transfers WHERE id=$1;
