// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: tasks.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
where cowID =$1
`

func (q *Queries) DeleteTask(ctx context.Context, cowid string) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, cowid)
	return err
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT cowid, date, type, text, done FROM tasks
ORDER BY cowID ASC, date ASC
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
			&i.Cowid,
			&i.Date,
			&i.Type,
			&i.Text,
			&i.Done,
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

const getTasksByDate = `-- name: GetTasksByDate :many
SELECT cowid, date, type, text, done FROM tasks
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
			&i.Cowid,
			&i.Date,
			&i.Type,
			&i.Text,
			&i.Done,
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

const updateTaskStatus = `-- name: UpdateTaskStatus :exec
UPDATE tasks
SET done = $1
where cowID = $2 and date = $3
`

type UpdateTaskStatusParams struct {
	Done  sql.NullBool `json:"done"`
	Cowid string       `json:"cowid"`
	Date  time.Time    `json:"date"`
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) error {
	_, err := q.exec(ctx, q.updateTaskStatusStmt, updateTaskStatus, arg.Done, arg.Cowid, arg.Date)
	return err
}

const upsertTasks = `-- name: UpsertTasks :exec
INSERT INTO tasks(cowID,date,type, text,done) VALUES ($1, $2, $3,$4, $5)
`

type UpsertTasksParams struct {
	Cowid string       `json:"cowid"`
	Date  time.Time    `json:"date"`
	Type  string       `json:"type"`
	Text  string       `json:"text"`
	Done  sql.NullBool `json:"done"`
}

func (q *Queries) UpsertTasks(ctx context.Context, arg UpsertTasksParams) error {
	_, err := q.exec(ctx, q.upsertTasksStmt, upsertTasks,
		arg.Cowid,
		arg.Date,
		arg.Type,
		arg.Text,
		arg.Done,
	)
	return err
}
