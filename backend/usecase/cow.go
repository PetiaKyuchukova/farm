package usecase

import (
	"context"
	"farm/backend/domain"
	"time"
)

type CowsUC struct {
	repo domain.CowRepo
}

func NewCowUC(repo domain.CowRepo) CowsUC {
	return CowsUC{repo: repo}
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
