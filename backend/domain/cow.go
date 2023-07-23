package domain

import (
	"context"
)

type Cow struct {
	ID            string         `json:"id"`
	Birthdate     CustomTime     `json:"birthdate"`
	Colour        string         `json:"colour"`
	Gender        string         `json:"gender"`
	Breed         string         `json:"breed"`
	MotherId      string         `json:"motherId"`
	MotherBreed   string         `json:"motherBreed"`
	FarmerId      string         `json:"fatherId"`
	FatherBreed   string         `json:"fatherBreed"`
	IsPregnant    bool           `json:"isPregnant"`
	Ovulation     CustomTime     `json:"ovulation"`
	Pregnancies   []Pregnancy    `json:"pregnancies"`
	Inseminations []Insemination `json:"inseminations"`
}

type CowRepo interface {
	UpsertCow(ctx context.Context, cow Cow) error
	DeleteCow(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]Cow, error)
	GetCowById(ctx context.Context, id string) (*Cow, error)
}
