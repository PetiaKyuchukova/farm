// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.deleteCowStmt, err = db.PrepareContext(ctx, deleteCow); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCow: %w", err)
	}
	if q.deleteNotificationStmt, err = db.PrepareContext(ctx, deleteNotification); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteNotification: %w", err)
	}
	if q.getAllCowsStmt, err = db.PrepareContext(ctx, getAllCows); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllCows: %w", err)
	}
	if q.getAllNotificationStmt, err = db.PrepareContext(ctx, getAllNotification); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllNotification: %w", err)
	}
	if q.getCowByIdStmt, err = db.PrepareContext(ctx, getCowById); err != nil {
		return nil, fmt.Errorf("error preparing query GetCowById: %w", err)
	}
	if q.getNotificationByCowIdStmt, err = db.PrepareContext(ctx, getNotificationByCowId); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotificationByCowId: %w", err)
	}
	if q.getNotificationsByDateStmt, err = db.PrepareContext(ctx, getNotificationsByDate); err != nil {
		return nil, fmt.Errorf("error preparing query GetNotificationsByDate: %w", err)
	}
	if q.upsertCowStmt, err = db.PrepareContext(ctx, upsertCow); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertCow: %w", err)
	}
	if q.upsertNotificationStmt, err = db.PrepareContext(ctx, upsertNotification); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertNotification: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteCowStmt != nil {
		if cerr := q.deleteCowStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCowStmt: %w", cerr)
		}
	}
	if q.deleteNotificationStmt != nil {
		if cerr := q.deleteNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteNotificationStmt: %w", cerr)
		}
	}
	if q.getAllCowsStmt != nil {
		if cerr := q.getAllCowsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllCowsStmt: %w", cerr)
		}
	}
	if q.getAllNotificationStmt != nil {
		if cerr := q.getAllNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllNotificationStmt: %w", cerr)
		}
	}
	if q.getCowByIdStmt != nil {
		if cerr := q.getCowByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCowByIdStmt: %w", cerr)
		}
	}
	if q.getNotificationByCowIdStmt != nil {
		if cerr := q.getNotificationByCowIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotificationByCowIdStmt: %w", cerr)
		}
	}
	if q.getNotificationsByDateStmt != nil {
		if cerr := q.getNotificationsByDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNotificationsByDateStmt: %w", cerr)
		}
	}
	if q.upsertCowStmt != nil {
		if cerr := q.upsertCowStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertCowStmt: %w", cerr)
		}
	}
	if q.upsertNotificationStmt != nil {
		if cerr := q.upsertNotificationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertNotificationStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                         DBTX
	tx                         *sql.Tx
	deleteCowStmt              *sql.Stmt
	deleteNotificationStmt     *sql.Stmt
	getAllCowsStmt             *sql.Stmt
	getAllNotificationStmt     *sql.Stmt
	getCowByIdStmt             *sql.Stmt
	getNotificationByCowIdStmt *sql.Stmt
	getNotificationsByDateStmt *sql.Stmt
	upsertCowStmt              *sql.Stmt
	upsertNotificationStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		deleteCowStmt:              q.deleteCowStmt,
		deleteNotificationStmt:     q.deleteNotificationStmt,
		getAllCowsStmt:             q.getAllCowsStmt,
		getAllNotificationStmt:     q.getAllNotificationStmt,
		getCowByIdStmt:             q.getCowByIdStmt,
		getNotificationByCowIdStmt: q.getNotificationByCowIdStmt,
		getNotificationsByDateStmt: q.getNotificationsByDateStmt,
		upsertCowStmt:              q.upsertCowStmt,
		upsertNotificationStmt:     q.upsertNotificationStmt,
	}
}
