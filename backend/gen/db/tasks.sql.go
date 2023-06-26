// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: tasks.sql

package db

import (
	"context"
	"time"
)

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
where id =$1
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, id)
	return err
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT id, cowid, date, type, text FROM tasks
ORDER BY id ASC, type ASC
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.query(ctx, q.getAllTasksStmt, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
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

const getTaskByCowId = `-- name: GetTaskByCowId :one
SELECT id, cowid, date, type, text FROM tasks
where cowID =$1
`

func (q *Queries) GetTaskByCowId(ctx context.Context, cowid string) (Task, error) {
	row := q.queryRow(ctx, q.getTaskByCowIdStmt, getTaskByCowId, cowid)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Cowid,
		&i.Date,
		&i.Type,
		&i.Text,
	)
	return i, err
}

const getTasksByDate = `-- name: GetTasksByDate :many
SELECT id, cowid, date, type, text FROM tasks
where date =$1
`

func (q *Queries) GetTasksByDate(ctx context.Context, date time.Time) ([]Task, error) {
	rows, err := q.query(ctx, q.getTasksByDateStmt, getTasksByDate, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
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

const upsertTasks = `-- name: UpsertTasks :exec
INSERT INTO tasks(cowID,date,type, text) VALUES ($1, $2, $3,$4)
`

type UpsertTasksParams struct {
	Cowid string    `json:"cowid"`
	Date  time.Time `json:"date"`
	Type  string    `json:"type"`
	Text  string    `json:"text"`
}

func (q *Queries) UpsertTasks(ctx context.Context, arg UpsertTasksParams) error {
	_, err := q.exec(ctx, q.upsertTasksStmt, upsertTasks,
		arg.Cowid,
		arg.Date,
		arg.Type,
		arg.Text,
	)
	return err
}
