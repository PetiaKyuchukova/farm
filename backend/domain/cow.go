package domain

import (
	"context"
	"time"
)

type Cow struct {
	ID            string         `json:"id"`
	Birthdate     time.Time      `json:"birthdate"`
	Colour        string         `json:"colour"`
	Gender        string         `json:"gender"`
	Breed         string         `json:"breed"`
	MotherId      string         `json:"motherId"`
	FarmerId      string         `json:"farmerId"`
	FatherBreed   string         `json:"fatherBreed"`
	IsPregnant    bool           `json:"isPregnant"`
	Ovulation     time.Time      `json:"ovulation"`
	Pregnancies   []Pregnancy    `json:"pregnancies"`
	Inseminations []Insemination `json:"inseminations"`
}

type CowRepo interface {
	UpsertCow(ctx context.Context, cow Cow) error
	DeleteCow(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]Cow, error)
	GetCowById(ctx context.Context, id string) (*Cow, error)
}
