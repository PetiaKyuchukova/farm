package usecase

import (
	"context"
	"farm/backend/domain"
	"fmt"
	"time"
)

type MilkUC struct {
	repo domain.MilkRepo
}

func NewMilkUC(repo domain.MilkRepo) MilkUC {
	return MilkUC{repo: repo}
}

func (m *MilkUC) UpsertMilk(ctx context.Context, milk domain.Milk) error {
	err := m.repo.UpsertMIlk(ctx, milk)
	if err != nil {
		fmt.Errorf("error upserting milk:%w", err)
		return err
	}

	return nil
}

func (m *MilkUC) GetMilkInTimeframe(ctx context.Context, from, to time.Time) ([]domain.Milk, error) {
	milkSeries, err := m.repo.GetMilkInTimeframe(ctx, from, to)
	if err != nil {
		fmt.Errorf("error getting milk series in timeframe: %w", err)
		return nil, err
	}

	return milkSeries, nil
}
