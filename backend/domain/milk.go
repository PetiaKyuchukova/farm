package domain

import (
	"context"
	"time"
)

type Milk struct {
	Date   CustomTime
	Liters float64
	Price  float64
}

type MilkRepo interface {
	UpsertMIlk(ctx context.Context, milk Milk) error
	GetMilkInTimeframe(ctx context.Context, from, to time.Time) ([]Milk, error)
}
