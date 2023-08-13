package unitofwork

import (
	"context"
	"farm/backend/domain"
)

type Callback func(ctx context.Context, uow Cow) error

type Executor interface {
	Do(ctx context.Context, fn Callback) error
}

type Cow interface {
	CowRepo() domain.CowRepo
	PregnancyRepo() domain.PregnancyRepo
	InseminationRepo() domain.InseminationRepo
	TaskRepo() domain.TaskRepo
}
