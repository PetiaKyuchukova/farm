package gen

import (
	"context"
	"database/sql"
	"farm/backend/domain"
	db2 "farm/backend/gen/db"
	"time"
)

type CowRepo struct {
	querier db2.Querier
}

func NewCowRepo(querier db2.Querier) *CowRepo {
	return &CowRepo{querier: querier}
}

func makeNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func (r *CowRepo) UpsertCow(ctx context.Context, cow domain.Cow) error {
	err := r.querier.UpsertCow(ctx, db2.UpsertCowParams{
		ID:          cow.ID,
		Birthdate:   cow.Birthdate.Time,
		Gender:      makeNullString(cow.Gender),
		Breed:       makeNullString(cow.Breed),
		Colour:      makeNullString(cow.Colour),
		Motherid:    makeNullString(cow.MotherId),
		Fatherid:    makeNullString(cow.FarmerId),
		Fatherbreed: makeNullString(cow.Breed),
		Ispregnant:  makeNullBool(&cow.IsPregnant),
		Ovulation:   makeNullTime(cow.Ovulation.Time),
	})
	return err
}

func (r *CowRepo) DeleteCow(ctx context.Context, id string) error {
	err := r.querier.DeleteCow(ctx, id)
	return err
}
func (r *CowRepo) GetAllCows(ctx context.Context) ([]domain.Cow, error) {
	rows, err := r.querier.GetAllCows(ctx)
	cows := []domain.Cow{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		cows = append(cows, domain.Cow{
			ID:            row.ID,
			Birthdate:     domain.CustomTime{row.Birthdate},
			Colour:        row.Colour.String,
			Gender:        row.Gender.String,
			Breed:         row.Breed.String,
			MotherId:      row.Motherid.String,
			FarmerId:      row.Fatherid.String,
			FatherBreed:   row.Fatherbreed.String,
			IsPregnant:    row.Ispregnant.Bool,
			Ovulation:     domain.CustomTime{row.Ovulation.Time},
			Pregnancies:   nil,
			Inseminations: nil,
		})
	}
	return cows, err
}
func (r *CowRepo) GetCowById(ctx context.Context, id string) (*domain.Cow, error) {
	row, err := r.querier.GetCowById(ctx, id)
	if err != nil {
		return nil, err
	}
	cow := domain.Cow{
		ID:            row.ID,
		Birthdate:     domain.CustomTime{row.Birthdate},
		Colour:        row.Colour.String,
		Gender:        row.Gender.String,
		Breed:         row.Breed.String,
		MotherId:      row.Motherid.String,
		FarmerId:      row.Fatherid.String,
		FatherBreed:   row.Fatherbreed.String,
		IsPregnant:    row.Ispregnant.Bool,
		Ovulation:     domain.CustomTime{row.Ovulation.Time},
		Inseminations: nil,
		Pregnancies:   nil,
	}

	inseminationsRow, err := r.querier.GetInseminationsByCowId(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, insemination := range inseminationsRow {
		cow.Inseminations = append(cow.Inseminations, domain.Insemination{
			Date:         insemination.Date,
			Breed:        insemination.Breed.String,
			IsArtificial: insemination.Isartificial.Bool,
		})
	}

	return &cow, err
}
