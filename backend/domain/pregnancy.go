package domain

import (
	"context"
	"time"
)

type Pregnancy struct {
	DetectedAt time.Time `json:"detectedAt"`
	FirstDay   time.Time `json:"firstDay"`
	LastDay    time.Time `json:"lastDay"`
}

type PregnancyRepo interface {
	UpsertPregnancy(ctx context.Context, pregnancy Pregnancy, cowId string) error
	GetPregnanciesByCowID(ctx context.Context, id string) ([]Pregnancy, error)
	DeletePregnancies(ctx context.Context, id string) error
}
