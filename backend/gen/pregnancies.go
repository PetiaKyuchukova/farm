package gen

import (
	"context"
	"farm/backend/domain"
	db2 "farm/backend/gen/db"
)

type PregnancyRepo struct {
	querier db2.Querier
}

func NewPregnancyRepo(querier db2.Querier) *PregnancyRepo {
	return &PregnancyRepo{querier: querier}
}

func (r *PregnancyRepo) UpsertPregnancy(ctx context.Context, pregnancy domain.Pregnancy, cowId string) error {
	err := r.querier.UpsertPregnancy(ctx, db2.UpsertPregnancyParams{
		Cowid:      cowId,
		Detectedat: pregnancy.DetectedAt.Time,
		Firstday:   makeNullTime(pregnancy.FirstDay.Time),
		Lastday:    makeNullTime(pregnancy.LastDay.Time),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *PregnancyRepo) GetPregnanciesByCowID(ctx context.Context, id string) ([]domain.Pregnancy, error) {
	pregnanciesRow, err := r.querier.GetPregnanciesByCowId(ctx, id)
	if err != nil {
		return nil, err
	}

	pregnancies := []domain.Pregnancy{}
	for _, pregnancy := range pregnanciesRow {
		pregnancies = append(pregnancies, domain.Pregnancy{
			DetectedAt: domain.CustomTime{pregnancy.Detectedat},
			FirstDay:   domain.CustomTime{pregnancy.Firstday.Time},
			LastDay:    domain.CustomTime{pregnancy.Lastday.Time},
		})
	}

	return pregnancies, nil
}

func (r *PregnancyRepo) DeletePregnancies(ctx context.Context, id string) error {
	err := r.querier.DeletePregnancy(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
