package gen

import (
	"context"
	"database/sql"
	"farm/backend/domain"
	db2 "farm/backend/gen/db"
	"time"
)

type FarmRepo struct {
	querier db2.Querier
}

func NewFarmRepo(querier db2.Querier) *FarmRepo {
	return &FarmRepo{querier: querier}
}
func makeNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

func (r *FarmRepo) UpsertCow(ctx context.Context, id string, birthdate time.Time, colour string, motherNum string) error {
	err := r.querier.UpsertCow(ctx, db2.UpsertCowParams{
		ID:        id,
		Birthdate: birthdate,
		Colour:    makeNullString(colour),
		Motherid:  makeNullString(motherNum),
	})
	return err
}

func (r *FarmRepo) DeleteCow(ctx context.Context, id string) error {
	err := r.querier.DeleteCow(ctx, id)
	return err
}
func (r *FarmRepo) GetAllCows(ctx context.Context) ([]domain.Cow, error) {
	rows, err := r.querier.GetAllCows(ctx)
	cows := []domain.Cow{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		cows = append(cows, domain.Cow{
			ID:        row.ID,
			Birthdate: row.Birthdate,
			Colour:    row.Colour.String,
			MotherId:  row.Motherid.String,
		})
	}
	return cows, err
}
func (r *FarmRepo) GetCowById(ctx context.Context, id string) (*domain.Cow, error) {
	row, err := r.querier.GetCowById(ctx, id)
	if err != nil {
		return nil, err
	}
	cow := domain.Cow{
		ID:        row.ID,
		Birthdate: row.Birthdate,
		Colour:    row.Colour.String,
		MotherId:  row.Motherid.String,
	}

	return &cow, err
}
