// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: entries.sql

package tutorial

import (
	"context"
	"database/sql"
)

const createEntries = `-- name: createEntries :one
INSERT INTO
    entries ("account_id", "amount") VALUES ($1, $2) RETURNING id, account_id, amount, created_at
`

type createEntriesParams struct {
	AccountID sql.NullInt64
	Amount    string
}

func (q *Queries) createEntries(ctx context.Context, arg createEntriesParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntries, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getAllEntries = `-- name: getAllEntries :many
SELECT id, account_id, amount, created_at FROM entries
`

func (q *Queries) getAllEntries(ctx context.Context) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getAllEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
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

const getEntries = `-- name: getEntries :one
SELECT id, account_id, amount, created_at FROM entries WHERE id=$1
`

func (q *Queries) getEntries(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntries, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getEntriesByAccountId = `-- name: getEntriesByAccountId :many
SELECT id, account_id, amount, created_at FROM entries WHERE account_id=$1
`

func (q *Queries) getEntriesByAccountId(ctx context.Context, accountID sql.NullInt64) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntriesByAccountId, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
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
