// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: inseminations.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getInseminationsByCowId = `-- name: GetInseminationsByCowId :many
SELECT cowid, date, breed, isartificial FROM inseminations
where cowId =$1
`

func (q *Queries) GetInseminationsByCowId(ctx context.Context, cowid string) ([]Insemination, error) {
	rows, err := q.query(ctx, q.getInseminationsByCowIdStmt, getInseminationsByCowId, cowid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Insemination{}
	for rows.Next() {
		var i Insemination
		if err := rows.Scan(
			&i.Cowid,
			&i.Date,
			&i.Breed,
			&i.Isartificial,
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

const upsertInsemination = `-- name: UpsertInsemination :exec
INSERT INTO inseminations(cowID,date,breed, isArtificial) VALUES ($1, $2, $3,$4)
`

type UpsertInseminationParams struct {
	Cowid        string         `json:"cowid"`
	Date         time.Time      `json:"date"`
	Breed        sql.NullString `json:"breed"`
	Isartificial sql.NullBool   `json:"isartificial"`
}

func (q *Queries) UpsertInsemination(ctx context.Context, arg UpsertInseminationParams) error {
	_, err := q.exec(ctx, q.upsertInseminationStmt, upsertInsemination,
		arg.Cowid,
		arg.Date,
		arg.Breed,
		arg.Isartificial,
	)
	return err
}
