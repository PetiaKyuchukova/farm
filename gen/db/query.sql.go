// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: inseminations.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteCow = `-- name: DeleteCow :exec
DELETE FROM cows.sql
where id =$1
`

func (q *Queries) DeleteCow(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteCowStmt, deleteCow, id)
	return err
}

const deleteNotification = `-- name: DeleteNotification :exec
DELETE FROM notifications
where id =$1
`

func (q *Queries) DeleteNotification(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteNotificationStmt, deleteNotification, id)
	return err
}

const getAllCows = `-- name: GetAllCows :many
SELECT id, birthdate, farmerid, colour, motherid, lastovulation, lastbirth, ispregnant, fertilization, givingbirthdate FROM cows.sql
ORDER BY id ASC, birthdate ASC
`

func (q *Queries) GetAllCows(ctx context.Context) ([]Cow, error) {
	rows, err := q.query(ctx, q.getAllCowsStmt, getAllCows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Cow{}
	for rows.Next() {
		var i Cow
		if err := rows.Scan(
			&i.ID,
			&i.Birthdate,
			&i.Farmerid,
			&i.Colour,
			&i.Motherid,
			&i.Lastovulation,
			&i.Lastbirth,
			&i.Ispregnant,
			&i.Fertilization,
			&i.Givingbirthdate,
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

const getAllNotification = `-- name: GetAllNotification :many
SELECT id, cowid, date, type, text FROM notifications
ORDER BY id ASC, type ASC
`

func (q *Queries) GetAllNotification(ctx context.Context) ([]Notification, error) {
	rows, err := q.query(ctx, q.getAllNotificationStmt, getAllNotification)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Notification{}
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.Cowid,
			&i.Date,
			&i.Type,
			&i.Text,
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

const getCowById = `-- name: GetCowById :one
SELECT id, birthdate, farmerid, colour, motherid, lastovulation, lastbirth, ispregnant, fertilization, givingbirthdate FROM cows.sql
where id =$1
`

func (q *Queries) GetCowById(ctx context.Context, id string) (Cow, error) {
	row := q.queryRow(ctx, q.getCowByIdStmt, getCowById, id)
	var i Cow
	err := row.Scan(
		&i.ID,
		&i.Birthdate,
		&i.Farmerid,
		&i.Colour,
		&i.Motherid,
		&i.Lastovulation,
		&i.Lastbirth,
		&i.Ispregnant,
		&i.Fertilization,
		&i.Givingbirthdate,
	)
	return i, err
}

const getNotificationByCowId = `-- name: GetNotificationByCowId :one
SELECT id, cowid, date, type, text FROM notifications
where cowID =$1
`

func (q *Queries) GetNotificationByCowId(ctx context.Context, cowid string) (Notification, error) {
	row := q.queryRow(ctx, q.getNotificationByCowIdStmt, getNotificationByCowId, cowid)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Cowid,
		&i.Date,
		&i.Type,
		&i.Text,
	)
	return i, err
}

const upsertCow = `-- name: UpsertCow :exec
INSERT INTO cows.sql (id, birthdate,colour,motherid) VALUES ($1,$2, $3, $4)
ON CONFLICT(id,birthDate)
    DO UPDATE SET
    colour = $3,
           motherid = $4
`

type UpsertCowParams struct {
	ID        string         `json:"id"`
	Birthdate time.Time      `json:"birthdate"`
	Colour    sql.NullString `json:"colour"`
	Motherid  sql.NullString `json:"motherid"`
}

func (q *Queries) UpsertCow(ctx context.Context, arg UpsertCowParams) error {
	_, err := q.exec(ctx, q.upsertCowStmt, upsertCow,
		arg.ID,
		arg.Birthdate,
		arg.Colour,
		arg.Motherid,
	)
	return err
}

const upsertNotification = `-- name: UpsertNotification :exec
INSERT INTO notifications(id, cowID,date,type, text) VALUES ($1,$2, $3, $4,$5)
`

type UpsertNotificationParams struct {
	ID    string    `json:"id"`
	Cowid string    `json:"cowid"`
	Date  time.Time `json:"date"`
	Type  string    `json:"type"`
	Text  string    `json:"text"`
}

func (q *Queries) UpsertNotification(ctx context.Context, arg UpsertNotificationParams) error {
	_, err := q.exec(ctx, q.upsertNotificationStmt, upsertNotification,
		arg.ID,
		arg.Cowid,
		arg.Date,
		arg.Type,
		arg.Text,
	)
	return err
}
