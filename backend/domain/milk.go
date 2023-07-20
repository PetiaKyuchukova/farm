package domain

import (
	"context"
	"time"
)

type Milk struct {
	Date   CustomTime `json:"date"`
	Liters float64    `json:"liters"`
	Price  float64    `json:"price"`
}

type MilkRepo interface {
	UpsertMIlk(ctx context.Context, milk Milk) error
	GetMilkInTimeframe(ctx context.Context, from, to time.Time) ([]Milk, error)
}
