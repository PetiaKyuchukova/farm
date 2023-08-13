package unitofwork

import (
	"context"
	"database/sql"
	"farm/backend/domain"
	"farm/backend/domain/unitofwork"
	"farm/backend/gen"
	"farm/backend/gen/db"
	"fmt"
)

type uowCow struct {
	cowRepo          domain.CowRepo
	pregnancyRepo    domain.PregnancyRepo
	inseminationRepo domain.InseminationRepo
	taskRepo         domain.TaskRepo
}

func (u *uowCow) CowRepo() domain.CowRepo {
	return u.cowRepo
}

func (u *uowCow) PregnancyRepo() domain.PregnancyRepo {
	return u.pregnancyRepo
}

func (u *uowCow) InseminationRepo() domain.InseminationRepo {
	return u.inseminationRepo
}

func (u *uowCow) TaskRepo() domain.TaskRepo {
	return u.taskRepo
}

type UnitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (uow *UnitOfWork) Do(ctx context.Context, fn unitofwork.Callback) error {
	tx, err := uow.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	querierWithTx := db.New(uow.db).WithTx(tx)

	err = fn(ctx, &uowCow{
		cowRepo:          gen.NewCowRepo(querierWithTx),
		pregnancyRepo:    gen.NewPregnancyRepo(querierWithTx),
		inseminationRepo: gen.NewInseminationRepo(querierWithTx),
		taskRepo:         gen.NewTaskRepo(querierWithTx),
	})

	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
