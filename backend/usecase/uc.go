package usecase

import (
	"context"
	"farm/backend/domain"
	"time"
)

type CowInteractor interface {
	UpsertCow(ctx context.Context, cow domain.Cow) error
	DeleteCowEntry(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]domain.Cow, error)
	GetCowEntryById(ctx context.Context, id string) (*domain.Cow, error)
}

type MilkInteractor interface {
	GetMilkInTimeframe(ctx context.Context, from, to time.Time) ([]domain.Milk, error)
	UpsertMilk(ctx context.Context, milk domain.Milk) error
}

type TaskInteractor interface {
	UpsertTask(ctx context.Context, notification domain.Task) error
	DeleteTask(ctx context.Context, id string) error
	FetchAllTasks(ctx context.Context) ([]domain.Task, error)
	FetchTasksByDate(ctx context.Context, date time.Time) ([]domain.Task, error)
}
