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
	FarmerId      string         `json:"FarmerId"`
	FatherBreed   string         `json:"fatherBreed"`
	IsPregnant    bool           `json:"isPregnant"`
	Ovulation     time.Time      `json:"ovulation"`
	Pregnancies   []Pregnancy    `json:"pregnancies"`
	Inseminations []Insemination `json:"inseminations"`
}

type Pregnancy struct {
	DetectedAt time.Time `json:"detectedAt"`
	FirstDay   time.Time `json:"firstDay"`
	LastDay    time.Time `json:"lastDay"`
}

type Insemination struct {
	Date         time.Time `json:"date"`
	Breed        string    `json:"breed"`
	IsArtificial bool      `json:"IsArtificial"`
}

type CowRepo interface {
	UpsertCow(ctx context.Context, cow Cow) error
	DeleteCow(ctx context.Context, id string) error
	GetAllCows(ctx context.Context) ([]Cow, error)
	GetCowById(ctx context.Context, id string) (*Cow, error)
}
