package domain

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Cow struct {
	ID            string         `json:"id"`
	Birthdate     CustomTime     `json:"birthdate"`
	Colour        string         `json:"colour"`
	Gender        string         `json:"gender"`
	Breed         string         `json:"breed"`
	MotherId      string         `json:"motherId"`
	FarmerId      string         `json:"farmerId"`
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

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
