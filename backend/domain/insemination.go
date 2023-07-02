package domain

import (
	"context"
	"time"
)

type Insemination struct {
	Date         time.Time `json:"date"`
	Breed        string    `json:"breed"`
	IsArtificial bool      `json:"IsArtificial"`
}

type InseminationRepo interface {
	UpsertInsemination(ctx context.Context, insemination Insemination, cowId string) error
	GetInseminationsByCowID(ctx context.Context, id string) ([]Insemination, error)
	DeleteInseminations(ctx context.Context, id string) error
}
