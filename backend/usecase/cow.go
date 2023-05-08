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

type CowsUC struct {
	repo domain.Repo
}

func NewCowUC(repo domain.Repo) *CowsUC {
	return &CowsUC{repo: repo}
}

func (c *CowsUC) UpsertCow(ctx context.Context, id string, birthdate time.Time, colour string, motherId string) error {
	return c.repo.UpsertCow(ctx, id, birthdate, colour, motherId)
}

func (c *CowsUC) DeleteCow(ctx context.Context, id string) error {
	return c.repo.DeleteCow(ctx, id)
}

func (c *CowsUC) GetAllCows(ctx context.Context) ([]domain.Cow, error) {
	return c.repo.GetAllCows(ctx)
}

func (c *CowsUC) GetCowById(ctx context.Context, id string) (*domain.Cow, error) {
	return c.repo.GetCowById(ctx, id)
}
