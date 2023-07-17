package gen

import (
	"context"
	"farm/backend/domain"
	db2 "farm/backend/gen/db"
	"time"
)

type MIlkRepo struct {
	querier db2.Querier
}

func NewMilkRepo(querier db2.Querier) *MIlkRepo {
	return &MIlkRepo{querier: querier}
}

func (m *MIlkRepo) UpsertMIlk(ctx context.Context, milk domain.Milk) error {
	err := m.querier.UpsertMilk(ctx, db2.UpsertMilkParams{
		Date:   milk.Date.Time,
		Liters: makeNullFloat(&milk.Liters),
		Price:  makeNullFloat(&milk.Price),
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *MIlkRepo) GetMilkInTimeframe(ctx context.Context, from, to time.Time) ([]domain.Milk, error) {
	rows, err := m.querier.GetMilkInTimeframe(ctx, db2.GetMilkInTimeframeParams{
		Date:   from,
		Date_2: to,
	})
	if err != nil {
		return nil, err
	}

	milkEntries := []domain.Milk{}

	for _, row := range rows {
		milkEntries = append(milkEntries, domain.Milk{
			Date:   domain.CustomTime{row.Date},
			Liters: row.Liters.Float64,
			Price:  row.Price.Float64,
		})
	}

	return milkEntries, nil
}
