-- name: createEntries :one
INSERT INTO
    entries ("account_id", "amount") VALUES ($1, $2) RETURNING *;

-- name: getEntries :one
SELECT * FROM entries WHERE id=$1;

-- name: getEntriesByAccountId :many
SELECT * FROM entries WHERE account_id=$1;

-- name: getAllEntries :many
SELECT * FROM entries;