// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transaction.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTransaction = `-- name: createTransaction :one
INSERT INTO
    transfers ("from_account_id", "to_account_id", "amount") VALUES ($1, $2, $3) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type createTransactionParams struct {
	FromAccountID sql.NullInt64
	ToAccountID   sql.NullInt64
	Amount        string
}

func (q *Queries) createTransaction(ctx context.Context, arg createTransactionParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransaction, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: deleteTransfer :exec
DELETE FROM transfers WHERE id=$1
`

func (q *Queries) deleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: getTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers WHERE id=$1
`

func (q *Queries) getTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listAllTransger = `-- name: listAllTransger :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers LIMIT $1 OFFSET $2
`

type listAllTransgerParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) listAllTransger(ctx context.Context, arg listAllTransgerParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listAllTransger, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: updateTransfer :one
UPDATE transfers SET amount=$2 WHERE id=$1 RETURNING id, from_account_id, to_account_id, amount, created_at
`

type updateTransferParams struct {
	ID     int64
	Amount string
}

func (q *Queries) updateTransfer(ctx context.Context, arg updateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer, arg.ID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}