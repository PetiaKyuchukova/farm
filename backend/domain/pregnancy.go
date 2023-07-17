package domain

import (
	"context"
)

type Pregnancy struct {
	DetectedAt CustomTime `json:"detectedAt"`
	FirstDay   CustomTime `json:"firstDay"`
	LastDay    CustomTime `json:"lastDay"`
}

type PregnancyRepo interface {
	UpsertPregnancy(ctx context.Context, pregnancy Pregnancy, cowId string) error
	GetPregnanciesByCowID(ctx context.Context, id string) ([]Pregnancy, error)
	DeletePregnancies(ctx context.Context, id string) error
}
