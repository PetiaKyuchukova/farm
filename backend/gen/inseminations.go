package gen

import (
	"context"
	"farm/backend/domain"
	db2 "farm/backend/gen/db"
)

type InseminationRepo struct {
	querier db2.Querier
}

func NewInseminationRepo(querier db2.Querier) *InseminationRepo {
	return &InseminationRepo{querier: querier}
}

func (r *InseminationRepo) UpsertInsemination(ctx context.Context, insemination domain.Insemination, cowId string) error {
	err := r.querier.UpsertInsemination(ctx, db2.UpsertInseminationParams{
		Cowid:        cowId,
		Date:         insemination.Date,
		Breed:        makeNullString(insemination.Breed),
		Isartificial: makeNullBool(&insemination.IsArtificial),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *InseminationRepo) GetInseminationsByCowID(ctx context.Context, id string) ([]domain.Insemination, error) {
	inseminationsRow, err := r.querier.GetInseminationsByCowId(ctx, id)
	if err != nil {
		return nil, err
	}

	inseminations := []domain.Insemination{}
	for _, insemination := range inseminationsRow {
		inseminations = append(inseminations, domain.Insemination{
			Date:         insemination.Date,
			Breed:        insemination.Breed.String,
			IsArtificial: insemination.Isartificial.Bool,
		})
	}

	return inseminations, nil
}

func (r *InseminationRepo) DeleteInseminations(ctx context.Context, id string) error {
	err := r.querier.DeleteInsemination(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
