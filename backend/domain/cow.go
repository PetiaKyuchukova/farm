package domain

import (
	"context"
	"time"
)

type Cow struct {
	ID        string    `json:"id"`
	Birthdate time.Time `json:"birthdate"`
	Colour    string    `json:"colour"`
	MotherId  string    `json:"motherId"`
}
type Repo interface {
	UpsertCow(ctx context.Context, id string, birthdate time.Time, colour string, motherNum string) error
	DeleteCow(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]Cow, error)
	GetCowById(ctx context.Context, id string) (*Cow, error)
}
