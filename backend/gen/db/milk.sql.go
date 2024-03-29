// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: milk.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getMilkInTimeframe = `-- name: GetMilkInTimeframe :many
SELECT date, liters, price FROM milk
where date BETWEEN $1 AND $2
`

type GetMilkInTimeframeParams struct {
	Date   time.Time `json:"date"`
	Date_2 time.Time `json:"date_2"`
}

func (q *Queries) GetMilkInTimeframe(ctx context.Context, arg GetMilkInTimeframeParams) ([]Milk, error) {
	rows, err := q.query(ctx, q.getMilkInTimeframeStmt, getMilkInTimeframe, arg.Date, arg.Date_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Milk{}
	for rows.Next() {
		var i Milk
		if err := rows.Scan(&i.Date, &i.Liters, &i.Price); err != nil {
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

const upsertMilk = `-- name: UpsertMilk :exec
INSERT INTO milk(date,liters,price) VALUES ($1, $2, $3)
    ON CONFLICT(date)
    DO UPDATE SET
         date = $1,
        liters = $2,
        price = $3
`

type UpsertMilkParams struct {
	Date   time.Time       `json:"date"`
	Liters sql.NullFloat64 `json:"liters"`
	Price  sql.NullFloat64 `json:"price"`
}

func (q *Queries) UpsertMilk(ctx context.Context, arg UpsertMilkParams) error {
	_, err := q.exec(ctx, q.upsertMilkStmt, upsertMilk, arg.Date, arg.Liters, arg.Price)
	return err
}
