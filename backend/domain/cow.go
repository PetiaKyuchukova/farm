package domain

import (
	"context"
	"time"
)

type Cow struct {
	ID                string    `json:"id"`
	Birthdate         time.Time `json:"birthdate"`
	Colour            string    `json:"colour"`
	MotherId          string    `json:"motherId"`
	FarmerId          string    `json:"FarmerId"`
	IsPregnant        bool      `json:"isPregnant"`
	LastOvulation     time.Time `json:"lastOvulation"`
	LastBirth         time.Time `json:"lastBirth"`
	LastFertilization time.Time `json:"LastFertilization"`
}

type CowRepo interface {
	UpsertCow(ctx context.Context, id string, birthdate time.Time, colour string, motherNum string) error
	DeleteCow(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]Cow, error)
	GetCowById(ctx context.Context, id string) (*Cow, error)
}
