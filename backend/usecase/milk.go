package usecase

import (
	"context"
	"farm/backend/domain"
	"fmt"
	"sort"
	"time"
)

type MilkUC struct {
	repo domain.MilkRepo
}

func NewMilkUC(repo domain.MilkRepo) MilkUC {
	return MilkUC{repo: repo}
}

func (m *MilkUC) UpsertMilk(ctx context.Context, milk domain.Milk) error {
	return m.repo.UpsertMIlk(ctx, milk)
}

func (m *MilkUC) GetMilkInTimeframe(ctx context.Context, from, to time.Time) ([]domain.Milk, error) {
	milkSeries, err := m.repo.GetMilkInTimeframe(ctx, from, to)
	if err != nil {
		fmt.Errorf("error getting milk series in timeframe: %w", err)
		return nil, err
	}
	sort.Slice(milkSeries, func(aa, ab int) bool {
		return milkSeries[aa].Date.Time.Before(milkSeries[ab].Date.Time)
	})
	return milkSeries, nil
}
