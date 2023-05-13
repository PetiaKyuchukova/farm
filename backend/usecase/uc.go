package usecase

import (
	"context"
	"farm/backend/domain"
	"time"
)

type Interactor interface {
	UpsertCow(ctx context.Context, id string, birthdate time.Time, colour string, motherNum string) error
	DeleteCow(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]domain.Cow, error)
	GetCowById(ctx context.Context, id string) (*domain.Cow, error)
}
